package handler

import (
	"encoding/json"
	"judge-server/internal/model"
	"judge-server/internal/service"
	"net/http"
)

// AuthHandler handles authentication-related requests
type AuthHandler struct {
	authManager *service.AuthManager
}

// NewAuthHandler creates a new authentication handler
func NewAuthHandler(authManager *service.AuthManager) *AuthHandler {
	return &AuthHandler{authManager: authManager}
}

// Login handles player login
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req model.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// Determine authentication method
	var profile *model.PlayerProfile
	var err error

	if req.Username != "" && req.Password != "" {
		// Test account login
		profile, err = ah.authManager.AuthenticateTestAccount(req.Username, req.Password)
	} else if req.ServerToken != "" {
		// Server token login
		profile, err = ah.authManager.AuthenticateWithServerToken(req.ServerToken)
	} else {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing authentication credentials",
		})
		return
	}

	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Create client session
	clientName := req.ClientName
	if clientName == "" {
		clientName = "unknown"
	}

	session, err := ah.authManager.CreateClientSession(profile.PlayerID, clientName, "agent")
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to create session",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":        true,
		"player_id":      profile.PlayerID,
		"server_token":   profile.ServerToken,
		"access_token":   session.AccessToken,
		"client_session": session,
		"message":        "Login successful",
	})
}

// RegisterTestAccount handles test account registration
func (ah *AuthHandler) RegisterTestAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	if req.Username == "" || req.Password == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Username and password are required",
		})
		return
	}

	profile, testAccount, err := ah.authManager.CreateTestAccount(req.Username, req.Password)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Create initial client session
	session, err := ah.authManager.CreateClientSession(profile.PlayerID, "registration", "agent")
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to create session",
		})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"success":      true,
		"player_id":    profile.PlayerID,
		"server_token": profile.ServerToken,
		"access_token": session.AccessToken,
		"username":     testAccount.Username,
		"message":      "Test account created successfully",
	})
}

// ListTestAccounts lists all test accounts (admin endpoint)
func (ah *AuthHandler) ListTestAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	accounts, err := ah.authManager.ListAllTestAccounts()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Sanitize passwords before returning
	for _, acc := range accounts {
		acc.Password = "***"
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":  true,
		"accounts": accounts,
		"total":    len(accounts),
	})
}

// BindGitHub handles GitHub account binding
func (ah *AuthHandler) BindGitHub(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req struct {
		PlayerToken string `json:"player_token"`
		ClientName  string `json:"client_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	if req.PlayerToken == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Player token is required",
		})
		return
	}

	if req.ClientName == "" {
		req.ClientName = "unknown"
	}

	resp, err := ah.authManager.UsePlayerTokenToBindGitHub(req.PlayerToken, req.ClientName)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// GetPlayerProfile retrieves the current player's profile
func (ah *AuthHandler) GetPlayerProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	profile, err := ah.authManager.AuthenticateWithAccessToken(r.Header.Get("Authorization"))
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"error":   "Unauthorized",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"profile": profile,
	})
}

// CreatePlayerTokenForGitHub creates a binding token for GitHub users
func (ah *AuthHandler) CreatePlayerTokenForGitHub(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req struct {
		GithubID    int    `json:"github_id"`
		GithubLogin string `json:"github_login"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	if req.GithubID == 0 || req.GithubLogin == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "GitHub ID and login are required",
		})
		return
	}

	token, err := ah.authManager.CreatePlayerTokenForGitHub(req.GithubID, req.GithubLogin)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"success":      true,
		"player_token": token.Token,
		"expires_at":   token.ExpiresAt,
		"message":      "Player token created. Use this token to bind your GitHub account.",
	})
}

// ValidateToken validates an access token
func (ah *AuthHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	if req.AccessToken == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Access token is required",
		})
		return
	}

	profile, err := ah.authManager.AuthenticateWithAccessToken(req.AccessToken)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":   true,
		"player_id": profile.PlayerID,
		"valid":     true,
	})
}
