package handler

import (
	"context"
	"net/http"
	"strings"
)

// AuthMiddleware is a middleware that validates access tokens
type AuthMiddleware struct {
	authHandler *AuthHandler
}

// NewAuthMiddleware creates a new authentication middleware
func NewAuthMiddleware(authHandler *AuthHandler) *AuthMiddleware {
	return &AuthMiddleware{authHandler: authHandler}
}

// Middleware wraps an http handler with authentication
func (am *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		var token string

		if authHeader != "" {
			// Support both "Bearer <token>" and direct token formats
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				token = parts[1]
			} else if len(parts) == 1 {
				token = parts[0]
			}
		}

		// Also check for token in query parameter (for some clients)
		if token == "" {
			token = r.URL.Query().Get("token")
		}

		// Set token in context
		if token != "" {
			ctx := context.WithValue(r.Context(), "access_token", token)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

// RequireAuth wraps an http handler with required authentication
func (am *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		var token string

		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				token = parts[1]
			} else if len(parts) == 1 {
				token = parts[0]
			}
		}

		if token == "" {
			token = r.URL.Query().Get("token")
		}

		if token == "" {
			writeJSON(w, http.StatusUnauthorized, map[string]interface{}{
				"success": false,
				"error":   "Missing authentication token",
			})
			return
		}

		// Store token and player info in context
		ctx := context.WithValue(r.Context(), "access_token", token)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
