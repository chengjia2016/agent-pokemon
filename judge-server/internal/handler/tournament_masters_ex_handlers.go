package handler

import (
	"encoding/json"
	"judge-server/internal/service"
	"net/http"
	"strconv"
	"time"
)

// TournamentMastersExHandlers handles HTTP requests for Tournament System
type TournamentMastersExHandlers struct {
	tournamentService *service.TournamentService
}

// NewTournamentMastersExHandlers creates a new instance of TournamentMastersExHandlers
func NewTournamentMastersExHandlers(tournamentService *service.TournamentService) *TournamentMastersExHandlers {
	return &TournamentMastersExHandlers{
		tournamentService: tournamentService,
	}
}

// CreateSeasonRequest represents a request to create a tournament season
type CreateSeasonRequest struct {
	SeasonName    string          `json:"season_name"`
	SeasonNameZh  *string         `json:"season_name_zh,omitempty"`
	SeasonNumber  int             `json:"season_number"`
	Description   *string         `json:"description,omitempty"`
	DescriptionZh *string         `json:"description_zh,omitempty"`
	StartDate     time.Time       `json:"start_date"`
	EndDate       time.Time       `json:"end_date"`
	StartRank     int             `json:"start_rank"`
	MaxRankPoints int             `json:"max_rank_points"`
	RewardPool    json.RawMessage `json:"reward_pool,omitempty"`
}

// HandleCreateTournamentSeason creates a new tournament season
func (h *TournamentMastersExHandlers) HandleCreateTournamentSeason(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateSeasonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	season, err := h.tournamentService.CreateTournamentSeason(req.SeasonName, req.SeasonNameZh,
		req.SeasonNumber, req.Description, req.DescriptionZh, req.StartDate, req.EndDate,
		req.StartRank, req.MaxRankPoints, req.RewardPool)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(season)
}

// HandleGetTournamentSeason retrieves a tournament season
func (h *TournamentMastersExHandlers) HandleGetTournamentSeason(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	seasonID := r.URL.Query().Get("season_id")
	if seasonID == "" {
		http.Error(w, "season_id is required", http.StatusBadRequest)
		return
	}

	season, err := h.tournamentService.GetTournamentSeason(seasonID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(season)
}

// ActivateSeasonRequest represents a request to activate a season
type ActivateSeasonRequest struct {
	SeasonID string `json:"season_id"`
}

// HandleActivateTournamentSeason activates a tournament season
func (h *TournamentMastersExHandlers) HandleActivateTournamentSeason(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ActivateSeasonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.tournamentService.ActivateTournamentSeason(req.SeasonID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Season activated successfully"})
}

// RegisterPlayerRequest represents a request to register a player
type RegisterPlayerRequest struct {
	SeasonID string `json:"season_id"`
	PlayerID string `json:"player_id"`
}

// HandleRegisterPlayerInTournament registers a player in a tournament season
func (h *TournamentMastersExHandlers) HandleRegisterPlayerInTournament(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterPlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ranking, err := h.tournamentService.RegisterPlayerInTournament(req.SeasonID, req.PlayerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ranking)
}

// HandleGetPlayerRanking retrieves a player's ranking in a tournament
func (h *TournamentMastersExHandlers) HandleGetPlayerRanking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	seasonID := r.URL.Query().Get("season_id")
	playerID := r.URL.Query().Get("player_id")

	if seasonID == "" || playerID == "" {
		http.Error(w, "season_id and player_id are required", http.StatusBadRequest)
		return
	}

	ranking, err := h.tournamentService.GetPlayerRanking(seasonID, playerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ranking)
}

// HandleGetSeasonLeaderboard retrieves the top players in a tournament season
func (h *TournamentMastersExHandlers) HandleGetSeasonLeaderboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	seasonID := r.URL.Query().Get("season_id")
	if seasonID == "" {
		http.Error(w, "season_id is required", http.StatusBadRequest)
		return
	}

	limit := 100
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	rankings, err := h.tournamentService.GetSeasonLeaderboard(seasonID, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rankings)
}

// UpdatePlayerStatsRequest represents a request to update player tournament stats
type UpdatePlayerStatsRequest struct {
	SeasonID     string `json:"season_id"`
	PlayerID     string `json:"player_id"`
	PointsChange int    `json:"points_change"`
	IsWin        bool   `json:"is_win"`
}

// HandleUpdatePlayerTournamentStats updates player's tournament statistics
func (h *TournamentMastersExHandlers) HandleUpdatePlayerTournamentStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UpdatePlayerStatsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ranking, err := h.tournamentService.UpdatePlayerTournamentStats(req.SeasonID, req.PlayerID, req.PointsChange, req.IsWin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ranking)
}

// AwardRewardRequest represents a request to award a tournament reward
type AwardRewardRequest struct {
	SeasonID     string  `json:"season_id"`
	PlayerID     string  `json:"player_id"`
	RewardType   string  `json:"reward_type"`
	RewardName   *string `json:"reward_name,omitempty"`
	RewardAmount int     `json:"reward_amount"`
	RankEarnedAt int     `json:"rank_earned_at"`
}

// HandleAwardTournamentReward awards a reward to a player
func (h *TournamentMastersExHandlers) HandleAwardTournamentReward(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req AwardRewardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	reward, err := h.tournamentService.AwardTournamentReward(req.SeasonID, req.PlayerID, req.RewardType,
		req.RewardName, req.RewardAmount, req.RankEarnedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reward)
}

// HandleClaimTournamentReward claims a tournament reward
func (h *TournamentMastersExHandlers) HandleClaimTournamentReward(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rewardID := r.URL.Query().Get("reward_id")
	if rewardID == "" {
		http.Error(w, "reward_id is required", http.StatusBadRequest)
		return
	}

	reward, err := h.tournamentService.ClaimTournamentReward(rewardID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reward)
}
