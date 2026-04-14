package middleware

import (
	"judge-server/internal/db"
	"net/http"
)

// APIKeyMiddleware validates API key from request headers
type APIKeyMiddleware struct {
	database *db.Database
}

// NewAPIKeyMiddleware creates a new API key middleware
func NewAPIKeyMiddleware(database *db.Database) *APIKeyMiddleware {
	return &APIKeyMiddleware{
		database: database,
	}
}

// ValidateAPIKey validates the API key from the request
// Returns (githubID, error)
func (m *APIKeyMiddleware) ValidateAPIKey(r *http.Request) (int, error) {
	// Get API key from header
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		// Try from Authorization header with Bearer token
		auth := r.Header.Get("Authorization")
		if auth != "" && len(auth) > 7 && auth[:7] == "Bearer " {
			apiKey = auth[7:]
		}
	}

	if apiKey == "" {
		return 0, NewAPIError("missing api key", http.StatusUnauthorized)
	}

	// Verify API key in database
	user, err := m.database.GetUserByAPIKey(apiKey)
	if err != nil {
		return 0, NewAPIError("invalid api key", http.StatusUnauthorized)
	}

	return user.GithubID, nil
}

// APIError represents an API error
type APIError struct {
	Message    string
	StatusCode int
}

func NewAPIError(message string, statusCode int) *APIError {
	return &APIError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func (e *APIError) Error() string {
	return e.Message
}
