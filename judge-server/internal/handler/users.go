package handler

import (
	"encoding/json"
	"judge-server/internal/auth"
	"judge-server/internal/model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CreateUserAccount creates a new user account
func (h *Handler) CreateUserAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req struct {
		GithubID    int     `json:"github_id"`
		GithubLogin string  `json:"github_login"`
		Email       string  `json:"email"`
		AvatarURL   string  `json:"avatar_url"`
		Balance     float64 `json:"balance"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body: " + err.Error(),
		})
		return
	}

	// Auto-add starter pack bonus (500 coins) for new accounts
	const STARTER_PACK_COINS = 500.0
	finalBalance := STARTER_PACK_COINS
	if req.Balance > 0 {
		finalBalance = req.Balance // Allow override if explicitly provided
	}

	account := &model.UserAccount{
		GithubID:    req.GithubID,
		GithubLogin: req.GithubLogin,
		Email:       req.Email,
		AvatarURL:   req.AvatarURL,
		Balance:     finalBalance,
	}

	if err := h.db.CreateUserAccount(account); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to create user account: " + err.Error(),
		})
		return
	}

	// Generate API key
	apiKey, err := auth.GenerateAPIKey()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to generate API key: " + err.Error(),
		})
		return
	}

	// Save API key to database
	if err := h.db.SetUserAPIKey(req.GithubID, apiKey); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to save API key: " + err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"success":            true,
		"message":            "User account created with starter pack bonus",
		"user":               account,
		"api_key":            apiKey,
		"starter_pack_bonus": STARTER_PACK_COINS,
		"bonus_reason":       "Automatic starter pack reward for new players",
	})
}

// GetUserAccount retrieves a user account
func (h *Handler) GetUserAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract github_id from URL path
	parts := strings.Split(r.URL.Path, "/")
	githubIDStr := parts[len(parts)-1]
	githubID, err := strconv.Atoi(githubIDStr)
	if err != nil {
		http.Error(w, "Invalid github_id", http.StatusBadRequest)
		return
	}

	account, err := h.db.GetUserAccount(githubID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

// GetUserBalance retrieves user balance
func (h *Handler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		http.Error(w, "github_id parameter required", http.StatusBadRequest)
		return
	}

	githubID, err := strconv.Atoi(githubIDStr)
	if err != nil {
		http.Error(w, "Invalid github_id", http.StatusBadRequest)
		return
	}

	balance, err := h.db.GetUserBalance(githubID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"github_id": githubID,
		"balance":   balance,
	})
}

// UpdateUserBalance updates user balance (INTERNAL USE ONLY)
// This endpoint should only be called from internal server processes
// DO NOT expose this for external client calls - coins should be added through game mechanics only
func (h *Handler) UpdateUserBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// INTERNAL TOKEN CHECK - Verify this is an internal call
	// For now, we check if the call is from localhost
	// In production, use proper authentication/authorization
	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = strings.Split(r.RemoteAddr, ":")[0]
	}

	// Only allow calls from localhost or internal network
	// Remove this for production and implement proper auth
	if clientIP != "localhost" && clientIP != "127.0.0.1" && !strings.HasPrefix(clientIP, "172.") && !strings.HasPrefix(clientIP, "192.168.") {
		writeJSON(w, http.StatusForbidden, map[string]interface{}{
			"success": false,
			"error":   "This endpoint is for internal use only. Balance updates must happen through game mechanics (quests, battles, etc.)",
		})
		return
	}

	var req struct {
		GithubID    int     `json:"github_id"`
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := h.db.UpdateUserBalance(req.GithubID, req.Amount, req.Description); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Get new balance
	balance, _ := h.db.GetUserBalance(req.GithubID)

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Balance updated (internal operation)",
		"amount":  req.Amount,
		"balance": balance,
	})
}

// GetUserPokemons retrieves all pokemons of a user
func (h *Handler) GetUserPokemons(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		http.Error(w, "github_id parameter required", http.StatusBadRequest)
		return
	}

	githubID, err := strconv.Atoi(githubIDStr)
	if err != nil {
		http.Error(w, "Invalid github_id", http.StatusBadRequest)
		return
	}

	pokemons, err := h.db.GetUserPokemons(githubID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"pokemons": pokemons,
		"total":    len(pokemons),
	})
}

// AddUserPokemon adds a pokemon to user's collection
func (h *Handler) AddUserPokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		GithubID int    `json:"github_id"`
		PetID    string `json:"pet_id"`
		PetName  string `json:"pet_name"`
		Level    int    `json:"level"`
		Species  string `json:"species"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pokemon := &model.UserPokemon{
		GithubID: req.GithubID,
		PetID:    req.PetID,
		PetName:  req.PetName,
		Level:    req.Level,
		Species:  req.Species,
	}

	if err := h.db.AddUserPokemon(pokemon); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Pokemon added",
		"pokemon": pokemon,
	})
}

// GetUserInventory retrieves user's inventory
func (h *Handler) GetUserInventory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		http.Error(w, "github_id parameter required", http.StatusBadRequest)
		return
	}

	githubID, err := strconv.Atoi(githubIDStr)
	if err != nil {
		http.Error(w, "Invalid github_id", http.StatusBadRequest)
		return
	}

	items, err := h.db.GetUserInventory(githubID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": items,
		"total": len(items),
	})
}

// AddUserItem adds an item to user's inventory
func (h *Handler) AddUserItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		GithubID int    `json:"github_id"`
		ItemID   string `json:"item_id"`
		ItemName string `json:"item_name"`
		Quantity int    `json:"quantity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.db.AddUserItem(req.GithubID, req.ItemID, req.ItemName, req.Quantity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Item added",
	})
}

// GetUserTransactions retrieves transaction history
func (h *Handler) GetUserTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		http.Error(w, "github_id parameter required", http.StatusBadRequest)
		return
	}

	githubID, err := strconv.Atoi(githubIDStr)
	if err != nil {
		http.Error(w, "Invalid github_id", http.StatusBadRequest)
		return
	}

	limit := 50
	if l := r.URL.Query().Get("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	transactions, err := h.db.GetUserTransactions(githubID, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"transactions": transactions,
		"total":        len(transactions),
	})
}

// InitUserConfig returns user info and API key for local configuration
func (h *Handler) InitUserConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req struct {
		GithubID int `json:"github_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body: " + err.Error(),
		})
		return
	}

	if req.GithubID == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "github_id parameter required",
		})
		return
	}

	// Get user account
	account, err := h.db.GetUserAccount(req.GithubID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "User not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"config": map[string]interface{}{
			"github_id":    account.GithubID,
			"github_login": account.GithubLogin,
			"api_key":      account.APIKey,
			"email":        account.Email,
		},
		"message": "User config ready - save this to ~/.config/petskill/settings.json",
	})
}

// RotateAPIKey rotates the user's API key
func (h *Handler) RotateAPIKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req struct {
		GithubID int    `json:"github_id"`
		Reason   string `json:"reason"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body: " + err.Error(),
		})
		return
	}

	if req.GithubID == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "github_id parameter required",
		})
		return
	}

	// Generate new API key
	newAPIKey, err := auth.GenerateAPIKey()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to generate new API key: " + err.Error(),
		})
		return
	}

	// Set expiration time to 90 days from now
	expiresAt := time.Now().AddDate(0, 0, 90)

	// Rotate the API key in database
	history, err := h.db.RotateAPIKey(req.GithubID, newAPIKey, expiresAt, req.Reason)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to rotate API key: " + err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":          true,
		"message":          "API key rotated successfully",
		"new_api_key":      newAPIKey,
		"expires_at":       expiresAt,
		"rotation_history": history,
	})
}

// GetAPIKeyHistory retrieves the API key rotation history
func (h *Handler) GetAPIKeyHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "github_id parameter required",
		})
		return
	}

	githubID, err := strconv.Atoi(githubIDStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid github_id",
		})
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	history, err := h.db.GetAPIKeyHistory(githubID, limit)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"history": history,
		"total":   len(history),
	})
}
