package handler

import (
	"encoding/json"
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/github"
	"judge-server/internal/model"
	"judge-server/internal/service"
	"net/http"
	"sync"
	"time"
)

type Handler struct {
	db                 *db.Database
	petValidator       *service.PetValidator
	battleValidator    *service.BattleValidator
	leaderboardService *service.LeaderboardService
	githubSync         *github.SyncService
	batchSync          *service.BatchSyncService
	eggService         *service.EggService
	captureService     *service.CaptureService
	foodService        *service.FoodService
	// Battle System Services
	battleEngine     *service.BattleEngine
	rewardCalculator *service.RewardCalculator
	defenseSystem    *service.DefenseSystem
	pokemonManager   *service.PokemonManager
	// World System Service
	worldService *service.WorldService
	// Language Service
	languageService *service.LanguageService
	// Masters EX System Services
	syncPairService        *service.SyncPairService
	battleMastersExService *service.BattleMastersExService
	tournamentService      *service.TournamentService
	seasonalEventService   *service.SeasonalEventService
	// Map Cache
	mapCache      map[string]*MapCacheEntry
	mapCacheMutex sync.RWMutex
	mapCacheTTL   time.Duration // Time to live for cached maps
}

// MapCacheEntry represents a cached map with timestamp
type MapCacheEntry struct {
	Data      *MapData
	ExpiresAt time.Time
}

func NewHandler(database *db.Database, syncService *github.SyncService, batchSync *service.BatchSyncService) *Handler {
	return &Handler{
		db:                 database,
		petValidator:       service.NewPetValidator(),
		battleValidator:    service.NewBattleValidator(),
		leaderboardService: service.NewLeaderboardService(database),
		githubSync:         syncService,
		batchSync:          batchSync,
		eggService:         service.NewEggService(),
		captureService:     service.NewCaptureService(),
		foodService:        service.NewFoodService(database),
		// Battle System Services
		battleEngine:     service.NewBattleEngine(database),
		rewardCalculator: service.NewRewardCalculator(database),
		defenseSystem:    service.NewDefenseSystem(database),
		pokemonManager:   service.NewPokemonManager(database),
		// World System Service
		worldService: service.NewWorldService(database),
		// Language Service
		languageService: service.NewLanguageService(database),
		// Masters EX System Services
		syncPairService:        service.NewSyncPairService(database.GetDB()),
		battleMastersExService: service.NewBattleMastersExService(database.GetDB()),
		tournamentService:      service.NewTournamentService(database.GetDB()),
		seasonalEventService:   service.NewSeasonalEventService(database.GetDB()),
		// Map Cache (30-minute TTL)
		mapCache:    make(map[string]*MapCacheEntry),
		mapCacheTTL: 30 * time.Minute,
	}
}

func (h *Handler) ValidatePet(w http.ResponseWriter, r *http.Request) {
	var pet model.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := h.petValidator.Validate(&pet)
	pet.IsValid = result.IsValid
	pet.ValidationMsg = joinStrings(result.Errors)

	if pet.IsValid {
		h.db.SavePet(&pet)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) ValidateBattle(w http.ResponseWriter, r *http.Request) {
	var battle model.Battle
	if err := json.NewDecoder(r.Body).Decode(&battle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := h.battleValidator.Validate(&battle)
	battle.IsValid = result.IsValid
	battle.ValidationMsg = joinStrings(result.Errors)

	if battle.IsValid {
		if err := h.db.SaveBattle(&battle); err != nil {
			fmt.Printf("Error saving battle: %v\n", err)
		}

		attacker := &model.Pet{ID: battle.AttackerID, Name: battle.AttackerName}
		defender := &model.Pet{ID: battle.DefenderID, Name: battle.DefenderName}

		battleResult := "lose"
		if battle.Winner == battle.AttackerID {
			battleResult = "win"
		}

		h.leaderboardService.UpdateLeaderboard(attacker, battleResult)
		h.leaderboardService.UpdateLeaderboard(defender, "lose")

		if h.batchSync != nil {
			h.batchSync.AddBattle(battle)
		} else if h.githubSync != nil {
			h.githubSync.SyncBattleReportToIssue(&battle)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	entries, err := h.leaderboardService.GetTopPlayers(100)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}

func (h *Handler) RecordFood(w http.ResponseWriter, r *http.Request) {
	var tx model.FoodTransaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx.ID = "ft_" + time.Now().Format("20060102150405")
	tx.Timestamp = time.Now()
	tx.IsValid = true

	h.db.SaveFoodTransaction(&tx)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func (h *Handler) RecordGrowth(w http.ResponseWriter, r *http.Request) {
	var gr model.GrowthRecord
	if err := json.NewDecoder(r.Body).Decode(&gr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gr.ID = "gr_" + time.Now().Format("20060102150405")
	gr.Timestamp = time.Now()
	gr.IsValid = true

	h.db.SaveGrowthRecord(&gr)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func (h *Handler) GenerateDailyLeaderboard(w http.ResponseWriter, r *http.Request) {
	date := time.Now().Format("2006-01-02")

	daily, err := h.leaderboardService.GenerateDailyLeaderboard(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if h.githubSync != nil {
		h.githubSync.SyncLeaderboardToIssue(date, daily.Entries)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(daily)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func joinStrings(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += "; "
		}
		result += s
	}
	return result
}

func (h *Handler) ValidateEggIncubation(w http.ResponseWriter, r *http.Request) {
	var egg model.Egg
	if err := json.NewDecoder(r.Body).Decode(&egg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := h.eggService.ValidateIncubation(&egg)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) ValidateCapture(w http.ResponseWriter, r *http.Request) {
	var capture model.CaptureAttempt
	if err := json.NewDecoder(r.Body).Decode(&capture); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := h.captureService.ValidateCapture(&capture)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) ValidateFood(w http.ResponseWriter, r *http.Request) {
	var foodReq model.FoodValidationRequest
	if err := json.NewDecoder(r.Body).Decode(&foodReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// For now, return a simple validation response
	// In production, this would fetch the farm from GitHub and validate
	// This is a placeholder that can be enhanced later
	result := &model.FoodValidationResult{
		IsValid:       true,
		CanEat:        true,
		Reasons:       []string{},
		NutritionGain: make(map[string]int),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// writeJSON writes a JSON response to the ResponseWriter with the given status code
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// validateAPIKey checks the API key from request headers and returns the github ID
// Returns (githubID, error)
func (h *Handler) validateAPIKey(r *http.Request) (int, error) {
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
		return 0, fmt.Errorf("missing api key")
	}

	// Verify API key in database
	user, err := h.db.GetUserByAPIKey(apiKey)
	if err != nil {
		return 0, fmt.Errorf("invalid api key")
	}

	// Check if API key has expired
	if user.APIKeyExpiresAt != nil && time.Now().After(*user.APIKeyExpiresAt) {
		return 0, fmt.Errorf("api key has expired")
	}

	return user.GithubID, nil
}

// ==================== Battle System Delegation Methods ====================

// Battle endpoints
func (h *Handler) StartBattle(w http.ResponseWriter, r *http.Request) {
	bh := NewBattleHandler(h.battleEngine, h.rewardCalculator, h.pokemonManager)
	bh.StartBattle(w, r)
}

func (h *Handler) GetBattle(w http.ResponseWriter, r *http.Request) {
	bh := NewBattleHandler(h.battleEngine, h.rewardCalculator, h.pokemonManager)
	bh.GetBattle(w, r)
}

func (h *Handler) ExecuteBattleRound(w http.ResponseWriter, r *http.Request) {
	bh := NewBattleHandler(h.battleEngine, h.rewardCalculator, h.pokemonManager)
	bh.ExecuteBattleRound(w, r)
}

func (h *Handler) EndBattle(w http.ResponseWriter, r *http.Request) {
	bh := NewBattleHandler(h.battleEngine, h.rewardCalculator, h.pokemonManager)
	bh.EndBattle(w, r)
}

func (h *Handler) GetUserBattles(w http.ResponseWriter, r *http.Request) {
	bh := NewBattleHandler(h.battleEngine, h.rewardCalculator, h.pokemonManager)
	bh.GetUserBattles(w, r)
}

func (h *Handler) GetBattleStats(w http.ResponseWriter, r *http.Request) {
	bh := NewBattleHandler(h.battleEngine, h.rewardCalculator, h.pokemonManager)
	bh.GetBattleStats(w, r)
}

// Defense endpoints
func (h *Handler) CreateBase(w http.ResponseWriter, r *http.Request) {
	dh := NewDefenseHandler(h.defenseSystem)
	dh.CreateBase(w, r)
}

func (h *Handler) GetBase(w http.ResponseWriter, r *http.Request) {
	dh := NewDefenseHandler(h.defenseSystem)
	dh.GetBase(w, r)
}

func (h *Handler) GetDefenders(w http.ResponseWriter, r *http.Request) {
	dh := NewDefenseHandler(h.defenseSystem)
	dh.GetDefenders(w, r)
}

func (h *Handler) GetDefenseHistory(w http.ResponseWriter, r *http.Request) {
	dh := NewDefenseHandler(h.defenseSystem)
	dh.GetDefenseHistory(w, r)
}

func (h *Handler) GetDefenseStats(w http.ResponseWriter, r *http.Request) {
	dh := NewDefenseHandler(h.defenseSystem)
	dh.GetDefenseStats(w, r)
}

// Pokemon Capture endpoints
func (h *Handler) ListWildPokemon(w http.ResponseWriter, r *http.Request) {
	pch := NewPokemonCaptureHandlerWithWorldService(h.pokemonManager, h.worldService)
	pch.ListWildPokemon(w, r)
}

func (h *Handler) CapturePokemon(w http.ResponseWriter, r *http.Request) {
	pch := NewPokemonCaptureHandlerWithWorldService(h.pokemonManager, h.worldService)
	pch.CapturePokemon(w, r)
}

func (h *Handler) GetCaptureHistory(w http.ResponseWriter, r *http.Request) {
	pch := NewPokemonCaptureHandler(h.pokemonManager)
	pch.GetCaptureHistory(w, r)
}

// Pokemon Management endpoints
func (h *Handler) ListUserPokemon(w http.ResponseWriter, r *http.Request) {
	pmh := NewPokemonManagementHandler(h.pokemonManager)
	pmh.ListUserPokemon(w, r)
}

func (h *Handler) GetPokemonDetails(w http.ResponseWriter, r *http.Request) {
	pmh := NewPokemonManagementHandler(h.pokemonManager)
	pmh.GetPokemonDetails(w, r)
}

func (h *Handler) ReleasePokemon(w http.ResponseWriter, r *http.Request) {
	pmh := NewPokemonManagementHandler(h.pokemonManager)
	pmh.ReleasePokemon(w, r)
}

func (h *Handler) GetPokemonRecoveryStatus(w http.ResponseWriter, r *http.Request) {
	pmh := NewPokemonManagementHandler(h.pokemonManager)
	pmh.GetRecoveryStatus(w, r)
}

func (h *Handler) TrainPokemon(w http.ResponseWriter, r *http.Request) {
	pmh := NewPokemonManagementHandler(h.pokemonManager)
	pmh.TrainPokemon(w, r)
}

// ==================== World System Handlers ====================

func (h *Handler) HandleWorldNPCs(w http.ResponseWriter, r *http.Request) {
	wh := NewWorldHTTPHandler(h.worldService)
	wh.HandleNPCs(w, r)
}

func (h *Handler) HandleWorldNPCTalk(w http.ResponseWriter, r *http.Request) {
	wh := NewWorldHTTPHandler(h.worldService)
	wh.HandleNPCTalk(w, r)
}

func (h *Handler) HandleWorldQuests(w http.ResponseWriter, r *http.Request) {
	wh := NewWorldHTTPHandler(h.worldService)
	wh.HandleQuests(w, r)
}

func (h *Handler) HandleWorldUserQuests(w http.ResponseWriter, r *http.Request) {
	wh := NewWorldHTTPHandler(h.worldService)
	wh.HandleUserQuests(w, r)
}

func (h *Handler) HandleWorldDungeons(w http.ResponseWriter, r *http.Request) {
	wh := NewWorldHTTPHandler(h.worldService)
	wh.HandleDungeons(w, r)
}

func (h *Handler) HandleWorldGyms(w http.ResponseWriter, r *http.Request) {
	wh := NewWorldHTTPHandler(h.worldService)
	wh.HandleGyms(w, r)
}

func (h *Handler) HandleWorldMapZones(w http.ResponseWriter, r *http.Request) {
	wh := NewWorldHTTPHandler(h.worldService)
	wh.HandleMapZones(w, r)
}

func (h *Handler) HandleWorldLevels(w http.ResponseWriter, r *http.Request) {
	wh := NewWorldHTTPHandler(h.worldService)
	wh.HandleLevels(w, r)
}

func (h *Handler) HandleWorldUserLevelProgress(w http.ResponseWriter, r *http.Request) {
	wh := NewWorldHTTPHandler(h.worldService)
	wh.HandleUserLevelProgress(w, r)
}

func (h *Handler) HandleCompleteQuest(w http.ResponseWriter, r *http.Request) {
	wh := NewWorldHTTPHandler(h.worldService)
	wh.HandleCompleteQuest(w, r)
}

// ==================== Masters EX System Handlers ====================

// HandleCreateSyncPair creates a new sync pair
func (h *Handler) HandleCreateSyncPair(w http.ResponseWriter, r *http.Request) {
	spHandler := NewSyncPairHandlers(h.syncPairService)
	spHandler.HandleCreateSyncPair(w, r)
}

// HandleGetSyncPair retrieves a sync pair
func (h *Handler) HandleGetSyncPair(w http.ResponseWriter, r *http.Request) {
	spHandler := NewSyncPairHandlers(h.syncPairService)
	spHandler.HandleGetSyncPair(w, r)
}

// HandleUpdateSyncPair levels up a sync pair
func (h *Handler) HandleUpdateSyncPair(w http.ResponseWriter, r *http.Request) {
	spHandler := NewSyncPairHandlers(h.syncPairService)
	spHandler.HandleLevelUpSyncPair(w, r)
}

// HandleGetSyncPairStats retrieves player game stats
func (h *Handler) HandleGetSyncPairStats(w http.ResponseWriter, r *http.Request) {
	spHandler := NewSyncPairHandlers(h.syncPairService)
	spHandler.HandleGetPlayerGameStats(w, r)
}

// HandleUseSyncMove uses a sync move
func (h *Handler) HandleUseSyncMove(w http.ResponseWriter, r *http.Request) {
	spHandler := NewSyncPairHandlers(h.syncPairService)
	spHandler.HandleUseSyncMove(w, r)
}

// HandleGetSyncPairDex retrieves player sync pairs (dex)
func (h *Handler) HandleGetSyncPairDex(w http.ResponseWriter, r *http.Request) {
	spHandler := NewSyncPairHandlers(h.syncPairService)
	spHandler.HandleGetPlayerSyncPairs(w, r)
}

// HandleCreateOrUpdateTeam creates or gets a battle team
func (h *Handler) HandleCreateOrUpdateTeam(w http.ResponseWriter, r *http.Request) {
	spHandler := NewSyncPairHandlers(h.syncPairService)
	spHandler.HandleCreateTeam(w, r)
}

// HandleGetTeam retrieves a battle team
func (h *Handler) HandleGetTeam(w http.ResponseWriter, r *http.Request) {
	spHandler := NewSyncPairHandlers(h.syncPairService)
	spHandler.HandleGetTeam(w, r)
}

// ==================== Battle System Handlers ====================

// HandleStartBattle starts a new 3v3 battle
func (h *Handler) HandleStartBattle(w http.ResponseWriter, r *http.Request) {
	bHandler := NewBattleMastersExHandlers(h.battleMastersExService)
	bHandler.HandleStartBattle(w, r)
}

// HandleGetBattleSession retrieves a battle session
func (h *Handler) HandleGetBattleSession(w http.ResponseWriter, r *http.Request) {
	bHandler := NewBattleMastersExHandlers(h.battleMastersExService)
	bHandler.HandleGetBattleSession(w, r)
}

// HandleExecuteBattleRound executes a battle round (updated to use ExecuteBattleAction)
func (h *Handler) HandleExecuteBattleRound(w http.ResponseWriter, r *http.Request) {
	bHandler := NewBattleMastersExHandlers(h.battleMastersExService)
	bHandler.HandleExecuteBattleAction(w, r)
}

// HandleEndBattle ends a battle
func (h *Handler) HandleEndBattle(w http.ResponseWriter, r *http.Request) {
	bHandler := NewBattleMastersExHandlers(h.battleMastersExService)
	bHandler.HandleEndBattle(w, r)
}

// HandleGetPlayerBattleStats retrieves player battle stats
func (h *Handler) HandleGetPlayerBattleStats(w http.ResponseWriter, r *http.Request) {
	bHandler := NewBattleMastersExHandlers(h.battleMastersExService)
	bHandler.HandleGetBattleStatistics(w, r)
}

// HandleGetPlayerBattles retrieves player's battles
func (h *Handler) HandleGetPlayerBattles(w http.ResponseWriter, r *http.Request) {
	bHandler := NewBattleMastersExHandlers(h.battleMastersExService)
	bHandler.HandleGetBattleHistory(w, r)
}

// ==================== Tournament System Handlers ====================

// HandleCreateTournamentSeason creates a new season
func (h *Handler) HandleCreateTournamentSeason(w http.ResponseWriter, r *http.Request) {
	tHandler := NewTournamentMastersExHandlers(h.tournamentService)
	tHandler.HandleCreateTournamentSeason(w, r)
}

// HandleGetTournamentSeason retrieves a season
func (h *Handler) HandleGetTournamentSeason(w http.ResponseWriter, r *http.Request) {
	tHandler := NewTournamentMastersExHandlers(h.tournamentService)
	tHandler.HandleGetTournamentSeason(w, r)
}

// HandleRegisterTournamentPlayer registers a player
func (h *Handler) HandleRegisterTournamentPlayer(w http.ResponseWriter, r *http.Request) {
	tHandler := NewTournamentMastersExHandlers(h.tournamentService)
	tHandler.HandleRegisterPlayerInTournament(w, r)
}

// HandleGetTournamentPlayerRanking gets player ranking
func (h *Handler) HandleGetTournamentPlayerRanking(w http.ResponseWriter, r *http.Request) {
	tHandler := NewTournamentMastersExHandlers(h.tournamentService)
	tHandler.HandleGetPlayerRanking(w, r)
}

// HandleGetTournamentLeaderboard gets leaderboard
func (h *Handler) HandleGetTournamentLeaderboard(w http.ResponseWriter, r *http.Request) {
	tHandler := NewTournamentMastersExHandlers(h.tournamentService)
	tHandler.HandleGetSeasonLeaderboard(w, r)
}

// HandleGetTournamentSeasonRewards gets rewards
func (h *Handler) HandleGetTournamentSeasonRewards(w http.ResponseWriter, r *http.Request) {
	tHandler := NewTournamentMastersExHandlers(h.tournamentService)
	tHandler.HandleClaimTournamentReward(w, r)
}

// HandleActivateTournamentSeason activates a tournament season
func (h *Handler) HandleActivateTournamentSeason(w http.ResponseWriter, r *http.Request) {
	tHandler := NewTournamentMastersExHandlers(h.tournamentService)
	tHandler.HandleActivateTournamentSeason(w, r)
}

// HandleGetActiveTournamentSeason gets active season
func (h *Handler) HandleGetActiveTournamentSeason(w http.ResponseWriter, r *http.Request) {
	tHandler := NewTournamentMastersExHandlers(h.tournamentService)
	tHandler.HandleGetTournamentSeason(w, r)
}

// ==================== Seasonal Event System Handlers ====================

// HandleCreateSeasonalEvent creates a new event
func (h *Handler) HandleCreateSeasonalEvent(w http.ResponseWriter, r *http.Request) {
	eHandler := NewSeasonalEventMastersExHandlers(h.seasonalEventService)
	eHandler.HandleCreateSeasonalEvent(w, r)
}

// HandleGetSeasonalEvent retrieves an event
func (h *Handler) HandleGetSeasonalEvent(w http.ResponseWriter, r *http.Request) {
	eHandler := NewSeasonalEventMastersExHandlers(h.seasonalEventService)
	eHandler.HandleGetSeasonalEvent(w, r)
}

// HandleJoinSeasonalEvent adds player to event
func (h *Handler) HandleJoinSeasonalEvent(w http.ResponseWriter, r *http.Request) {
	eHandler := NewSeasonalEventMastersExHandlers(h.seasonalEventService)
	eHandler.HandleJoinEvent(w, r)
}

// HandleGetEventProgress gets player event progress
func (h *Handler) HandleGetEventProgress(w http.ResponseWriter, r *http.Request) {
	eHandler := NewSeasonalEventMastersExHandlers(h.seasonalEventService)
	eHandler.HandleGetEventProgress(w, r)
}

// HandleUpdateEventProgress updates player progress
func (h *Handler) HandleUpdateEventProgress(w http.ResponseWriter, r *http.Request) {
	eHandler := NewSeasonalEventMastersExHandlers(h.seasonalEventService)
	eHandler.HandleUpdateEventProgress(w, r)
}

// HandleGetEventItems gets event items
func (h *Handler) HandleGetEventItems(w http.ResponseWriter, r *http.Request) {
	eHandler := NewSeasonalEventMastersExHandlers(h.seasonalEventService)
	eHandler.HandleGetEventItems(w, r)
}

// HandleBuyEventItem buys an event item (mapped to claim rewards)
func (h *Handler) HandleBuyEventItem(w http.ResponseWriter, r *http.Request) {
	eHandler := NewSeasonalEventMastersExHandlers(h.seasonalEventService)
	eHandler.HandleClaimEventRewards(w, r)
}

// HandleGetActiveSeasonalEvents gets active events
func (h *Handler) HandleGetActiveSeasonalEvents(w http.ResponseWriter, r *http.Request) {
	eHandler := NewSeasonalEventMastersExHandlers(h.seasonalEventService)
	eHandler.HandleGetActiveEvents(w, r)
}

// HandleActivateSeasonalEvent activates a seasonal event
func (h *Handler) HandleActivateSeasonalEvent(w http.ResponseWriter, r *http.Request) {
	eHandler := NewSeasonalEventMastersExHandlers(h.seasonalEventService)
	eHandler.HandleActivateSeasonalEvent(w, r)
}
