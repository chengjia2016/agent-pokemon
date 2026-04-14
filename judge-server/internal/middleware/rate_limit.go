package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// RateLimitConfig holds configuration for rate limiting
type RateLimitConfig struct {
	RequestsPerMinute int
	BurstSize         int
}

// RateLimiter implements token bucket rate limiting
type RateLimiter struct {
	config        RateLimitConfig
	buckets       map[string]*TokenBucket
	bucketMu      sync.RWMutex
	cleanupTicker *time.Ticker
	done          chan struct{}
}

// TokenBucket represents a token bucket for rate limiting
type TokenBucket struct {
	tokens     float64
	maxTokens  float64
	refillRate float64 // tokens per millisecond
	lastRefill time.Time
	mu         sync.Mutex
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(config RateLimitConfig) *RateLimiter {
	rl := &RateLimiter{
		config:  config,
		buckets: make(map[string]*TokenBucket),
		done:    make(chan struct{}),
	}

	// Start cleanup goroutine to remove old buckets (every 5 minutes)
	rl.cleanupTicker = time.NewTicker(5 * time.Minute)
	go rl.cleanupOldBuckets()

	return rl
}

// Stop gracefully shuts down the rate limiter
func (rl *RateLimiter) Stop() {
	close(rl.done)
	rl.cleanupTicker.Stop()
}

// AllowRequest checks if a request from the given identifier should be allowed
func (rl *RateLimiter) AllowRequest(identifier string) bool {
	bucket := rl.getOrCreateBucket(identifier)
	return bucket.TakeToken()
}

// getOrCreateBucket gets or creates a token bucket for an identifier
func (rl *RateLimiter) getOrCreateBucket(identifier string) *TokenBucket {
	rl.bucketMu.RLock()
	bucket, exists := rl.buckets[identifier]
	rl.bucketMu.RUnlock()

	if exists {
		return bucket
	}

	rl.bucketMu.Lock()
	defer rl.bucketMu.Unlock()

	// Double-check pattern
	if bucket, exists := rl.buckets[identifier]; exists {
		return bucket
	}

	// Create new bucket
	// Rate: requests per minute / 1000 milliseconds per second / 60 seconds per minute
	refillRate := float64(rl.config.RequestsPerMinute) / 60000.0 // tokens per millisecond
	bucket = &TokenBucket{
		tokens:     float64(rl.config.BurstSize),
		maxTokens:  float64(rl.config.BurstSize),
		refillRate: refillRate,
		lastRefill: time.Now(),
	}

	rl.buckets[identifier] = bucket
	return bucket
}

// TakeToken attempts to take a token from the bucket
func (tb *TokenBucket) TakeToken() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// Refill tokens based on time elapsed
	now := time.Now()
	elapsedMs := float64(now.Sub(tb.lastRefill).Milliseconds())
	tokensToAdd := elapsedMs * tb.refillRate

	tb.tokens = min(tb.maxTokens, tb.tokens+tokensToAdd)
	tb.lastRefill = now

	// Try to take a token
	if tb.tokens >= 1.0 {
		tb.tokens -= 1.0
		return true
	}

	return false
}

// cleanupOldBuckets removes buckets that haven't been used recently
func (rl *RateLimiter) cleanupOldBuckets() {
	for {
		select {
		case <-rl.done:
			return
		case <-rl.cleanupTicker.C:
			rl.bucketMu.Lock()
			now := time.Now()
			for identifier, bucket := range rl.buckets {
				bucket.mu.Lock()
				// Remove buckets not used for more than 10 minutes
				if now.Sub(bucket.lastRefill) > 10*time.Minute {
					delete(rl.buckets, identifier)
				}
				bucket.mu.Unlock()
			}
			rl.bucketMu.Unlock()
		}
	}
}

// GetRateLimitStatus returns the current status of rate limiting for an identifier
func (rl *RateLimiter) GetRateLimitStatus(identifier string) map[string]interface{} {
	bucket := rl.getOrCreateBucket(identifier)

	bucket.mu.Lock()
	defer bucket.mu.Unlock()

	// Calculate current tokens
	now := time.Now()
	elapsedMs := float64(now.Sub(bucket.lastRefill).Milliseconds())
	tokensToAdd := elapsedMs * bucket.refillRate
	currentTokens := min(bucket.maxTokens, bucket.tokens+tokensToAdd)

	return map[string]interface{}{
		"remaining": int(currentTokens),
		"limit":     int(bucket.maxTokens),
		"reset_at":  bucket.lastRefill.Add(time.Minute),
	}
}

// RateLimitMiddleware wraps an HTTP handler with rate limiting
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use API key or IP address as identifier
		identifier := r.Header.Get("X-API-Key")
		if identifier == "" {
			identifier = r.Header.Get("Authorization")
			if identifier != "" && len(identifier) > 7 && identifier[:7] == "Bearer " {
				identifier = identifier[7:]
			}
		}
		if identifier == "" {
			// Fall back to IP address
			identifier = r.RemoteAddr
		}

		// Check rate limit
		if !rl.AllowRequest(identifier) {
			w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", rl.config.RequestsPerMinute))
			w.Header().Set("X-RateLimit-Remaining", "0")
			w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", time.Now().Add(time.Minute).Unix()))
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// Add rate limit headers
		status := rl.GetRateLimitStatus(identifier)
		w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", rl.config.RequestsPerMinute))
		w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", status["remaining"]))
		w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", status["reset_at"].(time.Time).Unix()))

		next.ServeHTTP(w, r)
	})
}

// min returns the minimum of two float64 values
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
