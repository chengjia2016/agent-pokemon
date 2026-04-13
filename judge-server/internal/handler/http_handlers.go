package handler

import (
	"encoding/json"
	"fmt"
	"judge-server/internal/model"
	"judge-server/internal/service"
	"net/http"
	"strings"
)

// BaseHTTPHandler provides HTTP handlers for base management
type BaseHTTPHandler struct {
	baseService *service.BaseService
}

// NewBaseHTTPHandler creates a new BaseHTTPHandler
func NewBaseHTTPHandler(baseService *service.BaseService) *BaseHTTPHandler {
	return &BaseHTTPHandler{
		baseService: baseService,
	}
}

// HandleBase handles /api/base endpoints
func (h *BaseHTTPHandler) HandleBase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		h.createBase(w, r)
	case http.MethodGet:
		h.getBase(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *BaseHTTPHandler) createBase(w http.ResponseWriter, r *http.Request) {
	var req model.CreateBaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %v", err), http.StatusBadRequest)
		return
	}

	// In a real app, githubID would come from authentication middleware
	// For now, we'll require it as a query parameter
	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		http.Error(w, "github_id parameter required", http.StatusBadRequest)
		return
	}

	var githubID int
	_, err := fmt.Sscanf(githubIDStr, "%d", &githubID)
	if err != nil {
		http.Error(w, "invalid github_id", http.StatusBadRequest)
		return
	}

	base, err := h.baseService.CreateBaseForUser(githubID, &req)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create base: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"base":    base,
	})
}

func (h *BaseHTTPHandler) getBase(w http.ResponseWriter, r *http.Request) {
	// Extract githubID from query parameter or path
	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		http.Error(w, "github_id parameter required", http.StatusBadRequest)
		return
	}

	var githubID int
	_, err := fmt.Sscanf(githubIDStr, "%d", &githubID)
	if err != nil {
		http.Error(w, "invalid github_id", http.StatusBadRequest)
		return
	}

	base, err := h.baseService.GetUserBase(githubID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"base":    base,
	})
}

// MapHTTPHandler provides HTTP handlers for map management
type MapHTTPHandler struct {
	mapService *service.MapService
}

// NewMapHTTPHandler creates a new MapHTTPHandler
func NewMapHTTPHandler(mapService *service.MapService) *MapHTTPHandler {
	return &MapHTTPHandler{
		mapService: mapService,
	}
}

// HandleMapData handles /api/map/data endpoint
func (h *MapHTTPHandler) HandleMapData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mapData, err := h.mapService.GetMapData()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get map data: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapData)
}

// HandleIslands handles /api/map/islands endpoints
func (h *MapHTTPHandler) HandleIslands(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		h.createIsland(w, r)
	case http.MethodGet:
		h.listIslands(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *MapHTTPHandler) createIsland(w http.ResponseWriter, r *http.Request) {
	var req model.CreateIslandRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %v", err), http.StatusBadRequest)
		return
	}

	island := &model.Island{
		ID:          req.ID,
		NameEn:      req.NameEn,
		NameZh:      req.NameZh,
		Description: req.Description,
		Level:       req.Level,
		TerrainType: req.TerrainType,
		MaxBases:    req.MaxBases,
	}

	created, err := h.mapService.CreateIsland(island)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create island: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"island":  created,
	})
}

func (h *MapHTTPHandler) listIslands(w http.ResponseWriter, r *http.Request) {
	islands, err := h.mapService.ListIslands()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to list islands: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"islands": islands,
	})
}

// HandleTowns handles /api/map/towns endpoints
func (h *MapHTTPHandler) HandleTowns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		h.createTown(w, r)
	case http.MethodGet:
		h.listTowns(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *MapHTTPHandler) createTown(w http.ResponseWriter, r *http.Request) {
	var req model.CreateTownRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %v", err), http.StatusBadRequest)
		return
	}

	town := &model.Town{
		ID:               req.ID,
		IslandID:         req.IslandID,
		NameEn:           req.NameEn,
		NameZh:           req.NameZh,
		Description:      req.Description,
		Type:             req.Type,
		CoordX:           req.CoordX,
		CoordY:           req.CoordY,
		WildPokemonLevel: req.WildPokemonLevel,
	}

	created, err := h.mapService.CreateTown(town)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create town: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"town":    created,
	})
}

func (h *MapHTTPHandler) listTowns(w http.ResponseWriter, r *http.Request) {
	islandID := r.URL.Query().Get("island_id")

	if islandID == "" {
		http.Error(w, "island_id parameter required", http.StatusBadRequest)
		return
	}

	towns, err := h.mapService.GetTownsByIsland(islandID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to list towns: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"towns":   towns,
	})
}

// PokemonCreationHTTPHandler provides HTTP handlers for pokemon creation
type PokemonCreationHTTPHandler struct {
	pokemonService *service.PokemonService
}

// NewPokemonCreationHTTPHandler creates a new PokemonCreationHTTPHandler
func NewPokemonCreationHTTPHandler(pokemonService *service.PokemonService) *PokemonCreationHTTPHandler {
	return &PokemonCreationHTTPHandler{
		pokemonService: pokemonService,
	}
}

// HandlePokemonList handles /api/pokemon endpoints
func (h *PokemonCreationHTTPHandler) HandlePokemonList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		h.createPokemon(w, r)
	case http.MethodGet:
		h.listPokemons(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *PokemonCreationHTTPHandler) createPokemon(w http.ResponseWriter, r *http.Request) {
	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		http.Error(w, "github_id parameter required", http.StatusBadRequest)
		return
	}

	var githubID int
	_, err := fmt.Sscanf(githubIDStr, "%d", &githubID)
	if err != nil {
		http.Error(w, "invalid github_id", http.StatusBadRequest)
		return
	}

	var req struct {
		PetName string   `json:"pet_name"`
		Level   int      `json:"level"`
		Species string   `json:"species"`
		Skills  []string `json:"skills"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %v", err), http.StatusBadRequest)
		return
	}

	pokemon, err := h.pokemonService.CreatePokemon(githubID, req.Species, req.Level)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create pokemon: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"pokemon": pokemon,
	})
}

func (h *PokemonCreationHTTPHandler) listPokemons(w http.ResponseWriter, r *http.Request) {
	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		http.Error(w, "github_id parameter required", http.StatusBadRequest)
		return
	}

	var githubID int
	_, err := fmt.Sscanf(githubIDStr, "%d", &githubID)
	if err != nil {
		http.Error(w, "invalid github_id", http.StatusBadRequest)
		return
	}

	pokemons, err := h.pokemonService.ListUserPokemons(githubID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to list pokemons: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"pokemons": pokemons,
	})
}

// HandleEggList handles /api/pokemon/eggs endpoints
func (h *PokemonCreationHTTPHandler) HandleEggList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		h.createEgg(w, r)
	case http.MethodGet:
		h.listEggs(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *PokemonCreationHTTPHandler) createEgg(w http.ResponseWriter, r *http.Request) {
	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		http.Error(w, "github_id parameter required", http.StatusBadRequest)
		return
	}

	var githubID int
	_, err := fmt.Sscanf(githubIDStr, "%d", &githubID)
	if err != nil {
		http.Error(w, "invalid github_id", http.StatusBadRequest)
		return
	}

	var req struct {
		Species        string `json:"species"`
		EnergyRequired int64  `json:"energy_required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %v", err), http.StatusBadRequest)
		return
	}

	egg, err := h.pokemonService.CreateEgg(githubID, req.Species, req.EnergyRequired)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create egg: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"egg":     egg,
	})
}

func (h *PokemonCreationHTTPHandler) listEggs(w http.ResponseWriter, r *http.Request) {
	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		http.Error(w, "github_id parameter required", http.StatusBadRequest)
		return
	}

	var githubID int
	_, err := fmt.Sscanf(githubIDStr, "%d", &githubID)
	if err != nil {
		http.Error(w, "invalid github_id", http.StatusBadRequest)
		return
	}

	eggs, err := h.pokemonService.ListUserEggs(githubID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to list eggs: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"eggs":    eggs,
	})
}

// HandleEgg handles /api/pokemon/egg/{eggID} endpoints
func (h *PokemonCreationHTTPHandler) HandleEgg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract eggID from path
	path := strings.TrimPrefix(r.URL.Path, "/api/pokemon/egg/")
	eggID := strings.Split(path, "/")[0]

	if eggID == "" {
		http.Error(w, "egg_id required", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodGet {
		egg, err := h.pokemonService.GetEgg(eggID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"egg":     egg,
		})
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleStatus returns a status/health check
func (h *PokemonCreationHTTPHandler) HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"status":  "initialized",
	})
}
