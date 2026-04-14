package handler

import (
	"encoding/json"
	"judge-server/internal/service"
	"net/http"
	"time"
)

// SeasonalEventMastersExHandlers handles HTTP requests for Seasonal Events
type SeasonalEventMastersExHandlers struct {
	eventService *service.SeasonalEventService
}

// NewSeasonalEventMastersExHandlers creates a new instance of SeasonalEventMastersExHandlers
func NewSeasonalEventMastersExHandlers(eventService *service.SeasonalEventService) *SeasonalEventMastersExHandlers {
	return &SeasonalEventMastersExHandlers{
		eventService: eventService,
	}
}

// CreateSeasonalEventRequest represents a request to create a seasonal event
type CreateSeasonalEventRequest struct {
	EventName         string          `json:"event_name"`
	EventNameZh       string          `json:"event_name_zh"`
	EventType         string          `json:"event_type"` // story, challenge, time-attack, score-attack
	Season            string          `json:"season"`     // spring, summer, autumn, winter
	Description       *string         `json:"description,omitempty"`
	DescriptionZh     *string         `json:"description_zh,omitempty"`
	StartDate         time.Time       `json:"start_date"`
	EndDate           time.Time       `json:"end_date"`
	FeaturedSyncPairs json.RawMessage `json:"featured_sync_pairs,omitempty"`
	DifficultyLevels  json.RawMessage `json:"difficulty_levels,omitempty"`
	Rewards           json.RawMessage `json:"rewards,omitempty"`
	IsLimited         bool            `json:"is_limited"`
}

// HandleCreateSeasonalEvent creates a new seasonal event
func (h *SeasonalEventMastersExHandlers) HandleCreateSeasonalEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateSeasonalEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	event, err := h.eventService.CreateSeasonalEvent(req.EventName, req.EventNameZh, req.EventType,
		req.Season, req.Description, req.DescriptionZh, req.StartDate, req.EndDate,
		req.FeaturedSyncPairs, req.DifficultyLevels, req.Rewards, req.IsLimited)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

// HandleGetSeasonalEvent retrieves a seasonal event
func (h *SeasonalEventMastersExHandlers) HandleGetSeasonalEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	eventID := r.URL.Query().Get("event_id")
	if eventID == "" {
		http.Error(w, "event_id is required", http.StatusBadRequest)
		return
	}

	event, err := h.eventService.GetSeasonalEvent(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

// HandleGetActiveEvents retrieves all active seasonal events
func (h *SeasonalEventMastersExHandlers) HandleGetActiveEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	events, err := h.eventService.GetActiveEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// ActivateEventRequest represents a request to activate an event
type ActivateEventRequest struct {
	EventID string `json:"event_id"`
}

// HandleActivateSeasonalEvent activates a seasonal event
func (h *SeasonalEventMastersExHandlers) HandleActivateSeasonalEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ActivateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.eventService.ActivateSeasonalEvent(req.EventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Event activated successfully"})
}

// HandleDeactivateSeasonalEvent deactivates a seasonal event
func (h *SeasonalEventMastersExHandlers) HandleDeactivateSeasonalEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ActivateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.eventService.DeactivateSeasonalEvent(req.EventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Event deactivated successfully"})
}

// JoinEventRequest represents a request to join an event
type JoinEventRequest struct {
	EventID  string `json:"event_id"`
	PlayerID string `json:"player_id"`
}

// HandleJoinEvent registers a player in an event
func (h *SeasonalEventMastersExHandlers) HandleJoinEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req JoinEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	progress, err := h.eventService.JoinEvent(req.EventID, req.PlayerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(progress)
}

// HandleGetEventProgress retrieves a player's progress in an event
func (h *SeasonalEventMastersExHandlers) HandleGetEventProgress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	eventID := r.URL.Query().Get("event_id")
	playerID := r.URL.Query().Get("player_id")

	if eventID == "" || playerID == "" {
		http.Error(w, "event_id and player_id are required", http.StatusBadRequest)
		return
	}

	progress, err := h.eventService.GetEventProgress(eventID, playerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(progress)
}

// UpdateEventProgressRequest represents a request to update event progress
type UpdateEventProgressRequest struct {
	EventID            string  `json:"event_id"`
	PlayerID           string  `json:"player_id"`
	StageCompleted     int     `json:"stage_completed"`
	ProgressPercentage float64 `json:"progress_percentage"`
	Score              int64   `json:"score"`
	TimeSpent          int64   `json:"time_spent"`
}

// HandleUpdateEventProgress updates a player's progress in an event
func (h *SeasonalEventMastersExHandlers) HandleUpdateEventProgress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UpdateEventProgressRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	progress, err := h.eventService.UpdateEventProgress(req.EventID, req.PlayerID,
		req.StageCompleted, req.ProgressPercentage, req.Score, req.TimeSpent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(progress)
}

// ClaimRewardsRequest represents a request to claim event rewards
type ClaimRewardsRequest struct {
	EventID  string          `json:"event_id"`
	PlayerID string          `json:"player_id"`
	Rewards  json.RawMessage `json:"rewards"`
}

// HandleClaimEventRewards claims rewards for event progress
func (h *SeasonalEventMastersExHandlers) HandleClaimEventRewards(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ClaimRewardsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	progress, err := h.eventService.ClaimEventRewards(req.EventID, req.PlayerID, req.Rewards)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(progress)
}

// HandleGetEventItems retrieves all items for a seasonal event
func (h *SeasonalEventMastersExHandlers) HandleGetEventItems(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	eventID := r.URL.Query().Get("event_id")
	if eventID == "" {
		http.Error(w, "event_id is required", http.StatusBadRequest)
		return
	}

	items, err := h.eventService.GetEventItems(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
