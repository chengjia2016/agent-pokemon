package handler

import (
	"encoding/json"
	"judge-server/internal/service"
	"net/http"
	"strconv"
)

// WorldHTTPHandler handles HTTP requests for world system endpoints
type WorldHTTPHandler struct {
	worldService *service.WorldService
}

// NewWorldHTTPHandler creates a new WorldHTTPHandler
func NewWorldHTTPHandler(worldService *service.WorldService) *WorldHTTPHandler {
	return &WorldHTTPHandler{
		worldService: worldService,
	}
}

// HandleNPCs handles GET /api/npcs
func (h *WorldHTTPHandler) HandleNPCs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	townID := r.URL.Query().Get("town_id")

	if townID != "" {
		npcs, err := h.worldService.GetNPCsByTown(townID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"npcs":    npcs,
			"total":   len(npcs),
		})
	} else {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "town_id parameter required",
		})
	}
}

// HandleNPCTalk handles POST /api/npcs/talk
func (h *WorldHTTPHandler) HandleNPCTalk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req struct {
		UserID int `json:"user_id"`
		NPCID  int `json:"npc_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	npc, err := h.worldService.GetNPC(req.NPCID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "NPC not found",
		})
		return
	}

	dialogues, err := h.worldService.GetNPCDialogues(req.NPCID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":   true,
		"npc":       npc,
		"dialogues": dialogues,
	})
}

// HandleQuests handles GET /api/quests
func (h *WorldHTTPHandler) HandleQuests(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	questType := r.URL.Query().Get("type")
	difficultyStr := r.URL.Query().Get("difficulty")
	difficulty := 0
	if difficultyStr != "" {
		if d, err := strconv.Atoi(difficultyStr); err == nil {
			difficulty = d
		}
	}

	quests, err := h.worldService.GetAllQuests(questType, difficulty)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"quests":  quests,
		"total":   len(quests),
		"filters": map[string]interface{}{
			"type":       questType,
			"difficulty": difficulty,
		},
	})
}

// HandleUserQuests handles GET /api/user/quests
func (h *WorldHTTPHandler) HandleUserQuests(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "user_id parameter required",
		})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid user_id",
		})
		return
	}

	quests, err := h.worldService.GetUserQuests(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"quests":  quests,
		"total":   len(quests),
	})
}

// HandleDungeons handles GET /api/dungeons
func (h *WorldHTTPHandler) HandleDungeons(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	difficultyStr := r.URL.Query().Get("difficulty")
	difficulty := 0
	if difficultyStr != "" {
		if d, err := strconv.Atoi(difficultyStr); err == nil {
			difficulty = d
		}
	}

	dungeons, err := h.worldService.GetAllDungeons(difficulty)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":  true,
		"dungeons": dungeons,
		"total":    len(dungeons),
		"filters": map[string]interface{}{
			"difficulty": difficulty,
		},
	})
}

// HandleGyms handles GET /api/gyms
func (h *WorldHTTPHandler) HandleGyms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	townID := r.URL.Query().Get("town_id")

	if townID != "" {
		gyms, err := h.worldService.GetGymsByTown(townID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"gyms":    gyms,
			"total":   len(gyms),
		})
	} else {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "town_id parameter required",
		})
	}
}

// HandleMapZones handles GET /api/map/zones
func (h *WorldHTTPHandler) HandleMapZones(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	islandID := r.URL.Query().Get("island_id")

	if islandID != "" {
		zones, err := h.worldService.GetMapZonesByIsland(islandID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"zones":   zones,
			"total":   len(zones),
		})
	} else {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "island_id parameter required",
		})
	}
}

// HandleLevels handles GET /api/levels
func (h *WorldHTTPHandler) HandleLevels(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	zoneIDStr := r.URL.Query().Get("zone_id")

	if zoneIDStr != "" {
		zoneID, err := strconv.Atoi(zoneIDStr)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]interface{}{
				"success": false,
				"error":   "Invalid zone_id",
			})
			return
		}

		levels, err := h.worldService.GetLevelsByZone(zoneID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"levels":  levels,
			"total":   len(levels),
		})
	} else {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"message": "Levels endpoint - provide zone_id to filter",
			"levels":  []interface{}{},
			"total":   0,
		})
	}
}

// HandleUserLevelProgress handles GET /api/user/levels/progress
func (h *WorldHTTPHandler) HandleUserLevelProgress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "user_id parameter required",
		})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid user_id",
		})
		return
	}

	progress, err := h.worldService.GetUserLevelProgress(userID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":  true,
		"progress": progress,
		"total":    len(progress),
	})
}
