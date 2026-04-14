package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateAPIKey generates a random API key for user authentication
func GenerateAPIKey() (string, error) {
	// Generate 32 random bytes (256 bits)
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Convert to hex string
	apiKey := hex.EncodeToString(randomBytes)
	return apiKey, nil
}

// ValidateAPIKey validates the format of an API key
func ValidateAPIKey(apiKey string) bool {
	// Should be 64 characters (32 bytes in hex)
	if len(apiKey) != 64 {
		return false
	}

	// Should only contain hex characters
	_, err := hex.DecodeString(apiKey)
	return err == nil
}

// GetDefaultAPIKeyExpiration returns the default expiration time for API keys (90 days)
func GetDefaultAPIKeyExpiration() time.Time {
	return time.Now().AddDate(0, 0, 90)
}

// IsAPIKeyExpired checks if an API key has expired
func IsAPIKeyExpired(expiresAt *time.Time) bool {
	if expiresAt == nil {
		// No expiration time set
		return false
	}
	return time.Now().After(*expiresAt)
}
