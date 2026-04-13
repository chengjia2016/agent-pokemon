package handler

import (
	"encoding/json"
	"fmt"
	"judge-server/internal/model"
	"net/http"
	"strconv"
	"strings"
)

// ===== Farm API Handlers =====

// CreateFarm handles POST /api/farms/create
func (h *Handler) CreateFarm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.CreateFarmRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// Validate request
	if req.OwnerID == "" || req.RepositoryName == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing required fields: owner_id, repository_name",
		})
		return
	}

	// Create farm
	farm := &model.DatabaseFarm{
		OwnerID:        req.OwnerID,
		RepositoryName: req.RepositoryName,
		RepositoryURL:  req.RepositoryURL,
		FarmKey:        req.OwnerID + "/" + req.RepositoryName,
	}

	createdFarm, err := h.db.CreateFarm(farm)
	if err != nil {
		fmt.Printf("CreateFarm error: %v\n", err)
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to create farm",
			"details": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    createdFarm,
	})
}

// GetFarm handles GET /api/farms/:id
func (h *Handler) GetFarm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract farm ID from path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing farm ID",
		})
		return
	}

	farmID, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid farm ID",
		})
		return
	}

	farm, err := h.db.GetFarmByID(farmID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Farm not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    farm,
	})
}

// SearchFarms handles GET /api/farms/search?owner=&repository=
func (h *Handler) SearchFarms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	owner := r.URL.Query().Get("owner")
	repository := r.URL.Query().Get("repository")

	if owner == "" && repository == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "At least one search parameter required: owner, repository",
		})
		return
	}

	farms, err := h.db.SearchFarms(owner, repository)
	if err != nil {
		fmt.Printf("SearchFarms error: %v\n", err)
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to search farms",
			"details": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    farms,
	})
}

// AddFoodToFarm handles POST /api/farms/:id/foods
func (h *Handler) AddFoodToFarm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract farm ID from path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing farm ID",
		})
		return
	}

	farmID, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid farm ID",
		})
		return
	}

	var req model.AddFoodRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	if req.FoodID == "" || req.Quantity <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid food_id or quantity",
		})
		return
	}

	// Create farm food entry (no need to verify food exists first)
	farmFood := &model.DatabaseFarmFood{
		FarmID:            farmID,
		FoodID:            req.FoodID,
		FoodType:          req.FoodType,
		CurrentQuantity:   req.Quantity,
		MaxQuantity:       req.MaxQuantity,
		RegenerationHours: req.RegenerationHours,
		Emoji:             req.Emoji,
	}

	addedFood, err := h.db.AddFoodToFarm(farmFood)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to add food to farm: " + err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    addedFood,
	})
}

// ConsumeFood handles POST /api/farms/:id/foods/consume
func (h *Handler) ConsumeFood(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract farm ID from path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing farm ID",
		})
		return
	}

	farmID, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid farm ID",
		})
		return
	}

	var req model.ConsumeFoodRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	if req.FoodID == "" || req.EaterID == "" || req.EaterPetID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing required fields: food_id, eater_id, eater_pet_id",
		})
		return
	}

	// Record consumption
	err = h.db.ConsumeFood(farmID, req.FoodID, req.EaterID, req.EaterPetID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to consume food",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Food consumed successfully",
	})
}

// GetFarmStatistics handles GET /api/farms/:id/statistics
func (h *Handler) GetFarmStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract farm ID from path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing farm ID",
		})
		return
	}

	farmID, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid farm ID",
		})
		return
	}

	stats, err := h.db.GetFarmStatistics(farmID)
	if err != nil {
		fmt.Printf("GetFarmStatistics error for farmID %d: %v\n", farmID, err)
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get farm statistics: " + err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    stats,
	})
}

// DeleteFarm handles DELETE /api/farms/:id
func (h *Handler) DeleteFarm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract farm ID from path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Missing farm ID",
		})
		return
	}

	farmID, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid farm ID",
		})
		return
	}

	err = h.db.DeleteFarm(farmID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to delete farm",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Farm deleted successfully",
	})
}
