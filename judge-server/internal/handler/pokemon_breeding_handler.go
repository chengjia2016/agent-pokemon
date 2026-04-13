package handler

import (
	"encoding/json"
	"judge-server/internal/model"
	"judge-server/internal/service"
	"net/http"
	"strconv"
)

// AddEffortValuesRequest 请求体
type AddEffortValuesRequest struct {
	UserPokemonID int `json:"user_pokemon_id"`
	HPEffort      int `json:"hp_effort"`
	AtkEffort     int `json:"atk_effort"`
	DefEffort     int `json:"def_effort"`
	SpaEffort     int `json:"spa_effort"`
	SpdEffort     int `json:"spd_effort"`
	SpeEffort     int `json:"spe_effort"`
}

// ChangePokemonNatureRequest 请求体
type ChangePokemonNatureRequest struct {
	UserPokemonID int    `json:"user_pokemon_id"`
	NatureID      string `json:"nature_id"`
}

// AddEffortValues 增加宝可梦努力值
// POST /api/pokemon/effort-values
func (h *Handler) AddEffortValues(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
			"code":    "METHOD_NOT_ALLOWED",
		})
		return
	}

	var req AddEffortValuesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
			"code":    "INVALID_REQUEST",
		})
		return
	}

	if req.UserPokemonID <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid user_pokemon_id (must be positive integer)",
			"code":    "INVALID_POKEMON_ID",
		})
		return
	}

	// Validate effort values (each should be >= 0)
	if req.HPEffort < 0 || req.AtkEffort < 0 || req.DefEffort < 0 ||
		req.SpaEffort < 0 || req.SpdEffort < 0 || req.SpeEffort < 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Effort values cannot be negative",
			"code":    "INVALID_EFFORT_VALUES",
		})
		return
	}

	// Validate total effort values (max 510 in Pokemon games)
	totalEffort := req.HPEffort + req.AtkEffort + req.DefEffort + req.SpaEffort + req.SpdEffort + req.SpeEffort
	if totalEffort > 510 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Total effort values cannot exceed 510",
			"code":    "EFFORT_LIMIT_EXCEEDED",
		})
		return
	}

	// Initialize breeding service if not exists
	breedingService := service.NewPokemonBreedingService(h.db)

	// Build EV bonus map
	evBonus := map[string]int{
		"hp":      req.HPEffort,
		"attack":  req.AtkEffort,
		"defense": req.DefEffort,
		"sp_atk":  req.SpaEffort,
		"sp_def":  req.SpdEffort,
		"speed":   req.SpeEffort,
	}

	// Add effort values
	err := breedingService.AddEffortValues(req.UserPokemonID, evBonus)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"code":    "ADD_EFFORT_FAILED",
		})
		return
	}

	// Get updated effort stats
	stats, err := h.db.GetEffortStats(req.UserPokemonID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to retrieve updated stats",
			"code":    "RETRIEVE_STATS_FAILED",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    stats,
	})
}

// ChangePokemonNature 改变宝可梦性格
// POST /api/pokemon/nature
func (h *Handler) ChangePokemonNature(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
			"code":    "METHOD_NOT_ALLOWED",
		})
		return
	}

	var req ChangePokemonNatureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
			"code":    "INVALID_REQUEST",
		})
		return
	}

	if req.UserPokemonID <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid user_pokemon_id (must be positive integer)",
			"code":    "INVALID_POKEMON_ID",
		})
		return
	}

	if req.NatureID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "nature_id cannot be empty",
			"code":    "INVALID_NATURE_ID",
		})
		return
	}

	// Verify nature exists before updating
	nature, err := h.db.GetNature(req.NatureID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Nature not found",
			"code":    "NATURE_NOT_FOUND",
		})
		return
	}

	// Update pokemon nature
	err = h.db.UpdatePokemonNature(req.UserPokemonID, req.NatureID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"code":    "UPDATE_NATURE_FAILED",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    nature,
	})
}

// GetEffortStats 获取宝可梦努力值
// GET /api/pokemon/effort-values?user_pokemon_id=123
func (h *Handler) GetEffortStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	userPokemonIDStr := r.URL.Query().Get("user_pokemon_id")
	if userPokemonIDStr == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "user_pokemon_id parameter required",
		})
		return
	}

	userPokemonID, err := strconv.Atoi(userPokemonIDStr)
	if err != nil || userPokemonID <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid user_pokemon_id",
		})
		return
	}

	stats, err := h.db.GetEffortStats(userPokemonID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Effort stats not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    stats,
	})
}

// CalculatePokemonStats 计算宝可梦能力值
// POST /api/pokemon/calculate-stats
type CalculateStatsRequest struct {
	UserPokemonID    int                           `json:"user_pokemon_id"`
	PokemonSpeciesID string                        `json:"pokemon_species_id"`
	Level            int                           `json:"level"`
	BaseStats        *model.PokemonBaseStats       `json:"base_stats"`
	IndividualStats  *model.PokemonIndividualStats `json:"individual_stats"`
	EffortStats      *model.PokemonEffortStats     `json:"effort_stats"`
	Nature           *model.PokemonNature          `json:"nature"`
}

func (h *Handler) CalculatePokemonStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
			"code":    "METHOD_NOT_ALLOWED",
		})
		return
	}

	var req CalculateStatsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
			"code":    "INVALID_REQUEST",
		})
		return
	}

	// Validate user_pokemon_id
	if req.UserPokemonID <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "user_pokemon_id is required and must be positive",
			"code":    "INVALID_POKEMON_ID",
		})
		return
	}

	// Validate level
	if req.Level <= 0 || req.Level > 100 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid level (must be 1-100)",
			"code":    "INVALID_LEVEL",
		})
		return
	}

	breedingService := service.NewPokemonBreedingService(h.db)

	// Calculate stats (use empty stats if not provided)
	baseStats := req.BaseStats
	individualStats := req.IndividualStats
	effortStats := req.EffortStats
	nature := req.Nature

	stats := breedingService.CalculatePokemonStats(
		baseStats,
		individualStats,
		effortStats,
		nature,
		req.Level,
	)

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    stats,
	})
}

// AttemptCapture 尝试捕捉宝可梦
// POST /api/pokemon/capture
type CaptureRequest struct {
	UserID          int    `json:"user_id"`
	WildPokemonID   string `json:"wild_pokemon_id"`
	PokemonLevel    int    `json:"pokemon_level"`
	CurrentHP       int    `json:"current_hp"`
	MaxHP           int    `json:"max_hp"`
	StatusCondition string `json:"status_condition"`
	BallType        string `json:"ball_type"`
}

type CaptureResponse struct {
	Success     bool   `json:"success"`
	Captured    bool   `json:"captured"`
	CaptureRate int    `json:"capture_rate"`
	Message     string `json:"message"`
}

func (h *Handler) AttemptCapture(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
			"code":    "METHOD_NOT_ALLOWED",
		})
		return
	}

	var req CaptureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
			"code":    "INVALID_REQUEST",
		})
		return
	}

	// Validate user_id
	if req.UserID <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "user_id is required and must be positive",
			"code":    "INVALID_USER_ID",
		})
		return
	}

	// Validate wild_pokemon_id
	if req.WildPokemonID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "wild_pokemon_id is required",
			"code":    "INVALID_WILD_POKEMON",
		})
		return
	}

	// Validate pokemon_level
	if req.PokemonLevel <= 0 || req.PokemonLevel > 100 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "pokemon_level must be between 1 and 100",
			"code":    "INVALID_LEVEL",
		})
		return
	}

	// Validate HP values
	if req.CurrentHP < 0 || req.MaxHP <= 0 || req.CurrentHP > req.MaxHP {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid HP values (current_hp must be between 0 and max_hp)",
			"code":    "INVALID_HP",
		})
		return
	}

	// Validate ball type
	validBallTypes := map[string]bool{
		"pokeball": true, "greatball": true, "ultraball": true, "masterball": true,
	}
	if req.BallType == "" || !validBallTypes[req.BallType] {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid ball_type (must be: pokeball, greatball, ultraball, masterball)",
			"code":    "INVALID_BALL_TYPE",
		})
		return
	}

	// Get wild pokemon species
	wildPokemon, err := h.db.GetPokemonSpecies(req.WildPokemonID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Wild pokemon not found",
			"code":    "POKEMON_NOT_FOUND",
		})
		return
	}

	// Create capture attempt
	breedingService := service.NewPokemonBreedingService(h.db)
	captureSystem := service.NewCaptureSystem(breedingService, h.db)

	attempt := &service.CaptureAttempt{
		WildPokemonID:   req.WildPokemonID,
		PokemonLevel:    req.PokemonLevel,
		CurrentHP:       req.CurrentHP,
		MaxHP:           req.MaxHP,
		StatusCondition: req.StatusCondition,
		BallType:        req.BallType,
	}

	captured, captureRate, err := captureSystem.AttemptCapture(req.UserID, attempt, wildPokemon)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"code":    "CAPTURE_FAILED",
		})
		return
	}

	message := "Capture attempt failed"
	if captured {
		message = "Pokemon captured successfully!"
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":      true,
		"captured":     captured,
		"capture_rate": captureRate,
		"message":      message,
	})
}
