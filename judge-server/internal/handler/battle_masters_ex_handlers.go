package handler

import (
	"encoding/json"
	"judge-server/internal/service"
	"net/http"
	"strconv"
)

// BattleMastersExHandlers handles HTTP requests for Battle System
type BattleMastersExHandlers struct {
	battleService *service.BattleMastersExService
}

// NewBattleMastersExHandlers creates a new instance of BattleMastersExHandlers
func NewBattleMastersExHandlers(battleService *service.BattleMastersExService) *BattleMastersExHandlers {
	return &BattleMastersExHandlers{
		battleService: battleService,
	}
}

// StartBattleRequest represents a request to start a battle
type StartBattleRequest struct {
	Player1ID   string          `json:"player_1_id"`
	Player2ID   string          `json:"player_2_id"`
	Player3ID   *string         `json:"player_3_id,omitempty"`
	Player1Team json.RawMessage `json:"player_1_team"` // JSON array of sync pair IDs
	Player2Team json.RawMessage `json:"player_2_team"` // JSON array of sync pair IDs
	Player3Team json.RawMessage `json:"player_3_team,omitempty"`
}

// HandleStartBattle starts a new battle session
func (h *BattleMastersExHandlers) HandleStartBattle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req StartBattleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	battleSession, err := h.battleService.StartBattle(req.Player1ID, req.Player2ID, req.Player3ID,
		req.Player1Team, req.Player2Team, req.Player3Team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(battleSession)
}

// HandleGetBattleSession retrieves a battle session
func (h *BattleMastersExHandlers) HandleGetBattleSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sessionID := r.URL.Query().Get("session_id")
	if sessionID == "" {
		http.Error(w, "session_id is required", http.StatusBadRequest)
		return
	}

	battleSession, err := h.battleService.GetBattleSession(sessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(battleSession)
}

// ExecuteBattleActionRequest represents a request to execute an action during battle
type ExecuteBattleActionRequest struct {
	SessionID  string          `json:"session_id"`
	RoundNum   int             `json:"round_number"`
	PlayerID   string          `json:"player_id"`
	ActionType string          `json:"action_type"` // move, switch, mega-evolve, sync-move
	ActionData json.RawMessage `json:"action_data"`
}

// HandleExecuteBattleAction executes an action during battle
func (h *BattleMastersExHandlers) HandleExecuteBattleAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ExecuteBattleActionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	action, err := h.battleService.ExecuteBattleAction(req.SessionID, req.RoundNum, req.PlayerID, req.ActionType, req.ActionData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(action)
}

// UpdateBattleRoundRequest represents a request to update battle round
type UpdateBattleRoundRequest struct {
	SessionID    string          `json:"session_id"`
	CurrentRound int             `json:"current_round"`
	Player1Score int             `json:"player_1_score"`
	Player2Score int             `json:"player_2_score"`
	Player3Score int             `json:"player_3_score,omitempty"`
	Player1HP    json.RawMessage `json:"player_1_hp"`
	Player2HP    json.RawMessage `json:"player_2_hp"`
	Player3HP    json.RawMessage `json:"player_3_hp,omitempty"`
}

// HandleUpdateBattleRound updates battle round information
func (h *BattleMastersExHandlers) HandleUpdateBattleRound(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UpdateBattleRoundRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	battleSession, err := h.battleService.UpdateBattleRound(req.SessionID, req.CurrentRound,
		req.Player1Score, req.Player2Score, req.Player3Score,
		req.Player1HP, req.Player2HP, req.Player3HP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(battleSession)
}

// EndBattleRequest represents a request to end a battle
type EndBattleRequest struct {
	SessionID  string          `json:"session_id"`
	WinnerID   string          `json:"winner_id"`
	WinnerTeam json.RawMessage `json:"winner_team"`
}

// HandleEndBattle ends a battle session
func (h *BattleMastersExHandlers) HandleEndBattle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req EndBattleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	battleSession, err := h.battleService.EndBattle(req.SessionID, req.WinnerID, req.WinnerTeam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(battleSession)
}

// HandleGetBattleStatistics retrieves battle statistics for a player
func (h *BattleMastersExHandlers) HandleGetBattleStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	playerID := r.URL.Query().Get("player_id")
	if playerID == "" {
		http.Error(w, "player_id is required", http.StatusBadRequest)
		return
	}

	stats, err := h.battleService.GetBattleStatistics(playerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

// HandleGetBattleHistory retrieves battle history for a player
func (h *BattleMastersExHandlers) HandleGetBattleHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	playerID := r.URL.Query().Get("player_id")
	if playerID == "" {
		http.Error(w, "player_id is required", http.StatusBadRequest)
		return
	}

	limit := 10
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	battles, err := h.battleService.GetBattleHistory(playerID, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(battles)
}
