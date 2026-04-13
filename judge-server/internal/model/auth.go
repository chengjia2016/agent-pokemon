package model

import "time"

// PlayerProfile represents a player's account with GitHub binding and tokens
type PlayerProfile struct {
	ID            int        `json:"id"`
	PlayerID      string     `json:"player_id"`    // Unique player identifier
	GithubID      int        `json:"github_id"`    // GitHub user ID (optional for non-GitHub users)
	GithubLogin   string     `json:"github_login"` // GitHub username (optional)
	Email         string     `json:"email"`
	AvatarURL     string     `json:"avatar_url"`
	IsTestAccount bool       `json:"is_test_account"` // Flag for test accounts
	ServerToken   string     `json:"server_token"`    // Secret token for server binding
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	LastLoginAt   *time.Time `json:"last_login_at"`   // Nullable field
	TokenExpireAt *time.Time `json:"token_expire_at"` // Token expiration (optional)
}

// ClientSession represents a session on a specific client
type ClientSession struct {
	ID           int       `json:"id"`
	PlayerID     string    `json:"player_id"`
	ClientName   string    `json:"client_name"`  // "claude", "opencode", "openclaw", "gemini", etc.
	AccessToken  string    `json:"access_token"` // Client-specific access token
	ClientType   string    `json:"client_type"`  // "web", "cli", "agent", etc.
	LastActiveAt time.Time `json:"last_active_at"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// TestAccount represents a test account for development
type TestAccount struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`  // Test account username (e.g., "test_player_1", "test_player_2")
	Password  string    `json:"password"`  // Hashed password
	PlayerID  string    `json:"player_id"` // Reference to PlayerProfile
	Status    string    `json:"status"`    // "active", "inactive", "locked"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PlayerToken represents a token issued to bind GitHub account to server
type PlayerToken struct {
	ID          int        `json:"id"`
	PlayerID    *string    `json:"player_id"` // Nullable - not set until used
	GithubID    int        `json:"github_id"`
	GithubLogin string     `json:"github_login"`
	Token       string     `json:"token"`      // Unique binding token
	IsUsed      bool       `json:"is_used"`    // Whether token has been used for binding
	UsedAt      *time.Time `json:"used_at"`    // When token was used (nullable)
	ExpiresAt   time.Time  `json:"expires_at"` // Token expiration time
	CreatedAt   time.Time  `json:"created_at"`
}

// AuthRequest represents an authentication request
type AuthRequest struct {
	PlayerID    string `json:"player_id,omitempty"`    // For player ID based auth
	Username    string `json:"username,omitempty"`     // For test account login
	Password    string `json:"password,omitempty"`     // For test account login
	GithubToken string `json:"github_token,omitempty"` // GitHub PAT for GitHub-based auth
	ServerToken string `json:"server_token,omitempty"` // Server token for direct binding
	ClientName  string `json:"client_name,omitempty"`  // Client identification
}

// AuthResponse represents an authentication response
type AuthResponse struct {
	Success       bool           `json:"success"`
	PlayerID      string         `json:"player_id,omitempty"`
	ServerToken   string         `json:"server_token,omitempty"`
	AccessToken   string         `json:"access_token,omitempty"`
	ClientSession *ClientSession `json:"client_session,omitempty"`
	Message       string         `json:"message,omitempty"`
	Error         string         `json:"error,omitempty"`
}
