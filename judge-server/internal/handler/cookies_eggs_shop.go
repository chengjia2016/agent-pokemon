package handler

import (
	"encoding/json"
	"fmt"
	"judge-server/internal/model"
	"net/http"
	"strconv"
	"strings"
)

// ===== Cookie API Handlers =====

// RegisterCookie handles POST /api/cookies/register
func (h *Handler) RegisterCookie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		CookieID    string `json:"cookie_id" binding:"required"`
		CookieType  string `json:"cookie_type" binding:"required"`
		Emoji       string `json:"emoji"`
		SourceFile  string `json:"source_file"`
		GeneratorID string `json:"generator_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	cookie := &model.DatabaseCookie{
		CookieID:    req.CookieID,
		CookieType:  req.CookieType,
		Emoji:       req.Emoji,
		SourceFile:  req.SourceFile,
		GeneratorID: req.GeneratorID,
	}

	createdCookie, err := h.db.RegisterCookie(cookie)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to register cookie",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    createdCookie,
	})
}

// ClaimCookie handles POST /api/cookies/claim
func (h *Handler) ClaimCookie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		CookieID     string `json:"cookie_id" binding:"required"`
		PlayerID     string `json:"player_id" binding:"required"`
		ExpReward    int    `json:"exp_reward"`
		EnergyReward int    `json:"energy_reward"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	if req.ExpReward == 0 {
		req.ExpReward = 10
	}
	if req.EnergyReward == 0 {
		req.EnergyReward = 5
	}

	claim, err := h.db.ClaimCookie(req.CookieID, req.PlayerID, req.ExpReward, req.EnergyReward)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    claim,
	})
}

// GetCookieStatistics handles GET /api/cookies/statistics
func (h *Handler) GetCookieStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stats, err := h.db.GetCookieStats()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get statistics",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    stats,
	})
}

// ScanCookies handles GET /api/cookies/scan
func (h *Handler) ScanCookies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	playerID := r.URL.Query().Get("player_id")
	if playerID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing player_id parameter",
		})
		return
	}

	// Get player's fragment balance
	balance, err := h.db.GetPlayerFragmentBalance(playerID)
	if err != nil {
		fmt.Printf("ScanCookies error for playerID %s: %v\n", playerID, err)
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to scan cookies: " + err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"player_id": playerID,
			"fragments": balance,
		},
	})
}

// ===== Egg API Handlers =====

// CreateEgg handles POST /api/eggs/create
func (h *Handler) CreateEgg(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		EggID           string `json:"egg_id" binding:"required"`
		OwnerID         string `json:"owner_id" binding:"required"`
		IncubationHours int    `json:"incubation_hours"`
		Attributes      string `json:"attributes"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	if req.IncubationHours == 0 {
		req.IncubationHours = 72
	}

	egg := &model.DatabaseEgg{
		EggID:           req.EggID,
		OwnerID:         req.OwnerID,
		IncubationHours: req.IncubationHours,
		Stage:           0,
		Attributes:      req.Attributes,
	}

	createdEgg, err := h.db.CreateEgg(egg)
	if err != nil {
		// Log error for debugging
		fmt.Printf("CreateEgg error: %v\n", err)
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to create egg: " + err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    createdEgg,
	})
}

// GetEgg handles GET /api/eggs/:id
func (h *Handler) GetEgg(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing egg ID",
		})
		return
	}

	eggID := parts[3]
	egg, err := h.db.GetEggByID(eggID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Egg not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    egg,
	})
}

// HatchEgg handles POST /api/eggs/:id/hatch
func (h *Handler) HatchEgg(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing egg ID",
		})
		return
	}

	eggID := parts[3]

	var req struct {
		PetID   string `json:"pet_id"`
		PetName string `json:"pet_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// Get egg details to find owner
	egg, err := h.db.GetEggByID(eggID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Egg not found",
		})
		return
	}

	// Generate unique pet_id if not provided
	if req.PetID == "" {
		req.PetID = fmt.Sprintf("pet_%s_%d", egg.OwnerID, egg.ID)
	}

	// Hatch the egg
	err = h.db.HatchEgg(eggID, req.PetID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Parse owner ID as integer
	ownerID, err := strconv.Atoi(egg.OwnerID)
	if err != nil {
		// Log error but continue - egg was already hatched
		fmt.Printf("Warning: Invalid owner ID format: %v\n", err)
	} else {
		// Set default pet name if not provided
		petName := req.PetName
		if petName == "" {
			petName = "Pokemon"
		}

		// Create pet in user_pokemons table
		pokemon := &model.UserPokemon{
			GithubID: ownerID,
			PetID:    req.PetID,
			PetName:  petName,
			Level:    1,
			Species:  "Unknown",
		}

		err = h.db.AddUserPokemon(pokemon)
		if err != nil {
			// Log error but don't fail - egg was already hatched
			fmt.Printf("Warning: Failed to create pet in user inventory: %v\n", err)
		}
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Egg hatched successfully",
		"pet_id":  req.PetID,
	})
}

// GetEggStatistics handles GET /api/eggs/statistics
func (h *Handler) GetEggStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stats, err := h.db.GetEggStats()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get statistics",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    stats,
	})
}

// ===== Shop API Handlers =====

// ListShopItems handles GET /api/shop/items
func (h *Handler) ListShopItems(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	items, err := h.db.GetAllShopItems()
	if err != nil {
		fmt.Printf("GetAllShopItems error: %v\n", err)
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get items",
			"details": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    items,
		"total":   len(items),
	})
}

// BuyItem handles POST /api/shop/buy
func (h *Handler) BuyItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ItemID   string `json:"item_id" binding:"required"`
		PlayerID string `json:"player_id" binding:"required"`
		Quantity int    `json:"quantity" binding:"required,gt=0"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// Get item details
	item, err := h.db.GetShopItem(req.ItemID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Item not found",
		})
		return
	}

	totalPrice := item.Price * req.Quantity

	// Parse player ID and check balance before processing
	playerIDInt, err := strconv.Atoi(req.PlayerID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid player ID format",
		})
		return
	}

	// Check user balance
	currentBalance, err := h.db.GetUserBalance(playerIDInt)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Failed to check balance: " + err.Error(),
		})
		return
	}

	// Validate sufficient balance
	if currentBalance < float64(totalPrice) {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Insufficient balance. Required: %d, Available: %.0f", totalPrice, currentBalance),
		})
		return
	}

	// Process transaction
	transaction, err := h.db.BuyItem(item.ID, req.PlayerID, req.Quantity, totalPrice)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Deduct coins from player balance
	err = h.db.UpdateUserBalance(playerIDInt, float64(-totalPrice), fmt.Sprintf("Purchase: %d x %s", req.Quantity, item.ItemName))
	if err != nil {
		// Log the error but don't fail the response since transaction already recorded
		fmt.Printf("Warning: Failed to update user balance: %v\n", err)
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"transaction": transaction,
			"item":        item,
			"total_price": totalPrice,
		},
	})
}

// GetShopStatistics handles GET /api/shop/statistics
func (h *Handler) GetShopStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stats, err := h.db.GetShopStats()
	if err != nil {
		fmt.Printf("GetShopStats error: %v\n", err)
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get statistics",
			"details": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    stats,
	})
}

// GetTransactionHistory handles GET /api/shop/transactions?player_id=...
func (h *Handler) GetTransactionHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	playerID := r.URL.Query().Get("player_id")
	if playerID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing player_id parameter",
		})
		return
	}

	limit := 50
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	transactions, err := h.db.GetTransactionHistory(playerID, limit)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get transactions",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    transactions,
		"total":   len(transactions),
	})
}
