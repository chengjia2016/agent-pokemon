package handler

import (
	"net/http"
	"strconv"
)

// GetPokemonSpecies handles GET /api/pokemons/{id}
func (h *Handler) GetPokemonSpecies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Extract pokemon_id from query parameter or path
	pokemonID := r.URL.Query().Get("id")
	if pokemonID == "" {
		// Try path parameter
		pokemonID = r.URL.Query().Get("pokemon_id")
	}

	if pokemonID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "pokemon_id parameter required",
		})
		return
	}

	species, err := h.db.GetPokemonSpecies(pokemonID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Pokemon species not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    species,
	})
}

// ListPokemonSpecies handles GET /api/pokemons
func (h *Handler) ListPokemonSpecies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Get pagination parameters
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 20
	offset := 0

	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
		limit = l
	}
	if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
		offset = o
	}

	species, err := h.db.GetAllPokemonSpecies(limit, offset)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get Pokemon species: " + err.Error(),
		})
		return
	}

	count, _ := h.db.GetPokemonCount()

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    species,
		"total":   count,
		"limit":   limit,
		"offset":  offset,
	})
}

// SearchPokemon handles GET /api/pokemons/search
func (h *Handler) SearchPokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	keyword := r.URL.Query().Get("q")
	if keyword == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "search keyword required",
		})
		return
	}

	species, err := h.db.SearchPokemon(keyword)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Search failed: " + err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    species,
		"total":   len(species),
	})
}

// GetPokemonPokedex handles GET /api/pokemons/pokedex
func (h *Handler) GetPokemonPokedex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// Get all pokemons (limit to 150 for pokedex view)
	species, err := h.db.GetAllPokemonSpecies(150, 0)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get pokedex: " + err.Error(),
		})
		return
	}

	count, _ := h.db.GetPokemonCount()

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"pokedex": species,
		"total":   count,
	})
}
