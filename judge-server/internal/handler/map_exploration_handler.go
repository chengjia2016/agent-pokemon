package handler

import (
	"encoding/json"
	"judge-server/internal/service"
	"net/http"
	"strconv"
)

// EnterGrassAreaRequest 进入草地请求
type EnterGrassAreaRequest struct {
	AreaID string `json:"area_id"`
	UserID int    `json:"user_id"`
}

// ExitGrassAreaRequest 离开草地请求
type ExitGrassAreaRequest struct {
	AreaID string `json:"area_id"`
	UserID int    `json:"user_id"`
}

// EnterGrassArea 进入草地进行探索
// POST /api/map/enter-grass
func (h *Handler) EnterGrassArea(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
			"code":    "METHOD_NOT_ALLOWED",
		})
		return
	}

	var req EnterGrassAreaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
			"code":    "INVALID_REQUEST",
		})
		return
	}

	if req.AreaID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "area_id is required",
			"code":    "INVALID_AREA_ID",
		})
		return
	}

	if req.UserID <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "user_id must be a positive integer",
			"code":    "INVALID_USER_ID",
		})
		return
	}

	// Create map exploration system
	explorationSystem := service.NewMapExplorationSystem(h.db)

	// Enter grass area with empty spawns (would be fetched from DB in real implementation)
	var spawns []service.WildPokemonSpawn
	result := explorationSystem.EnterGrassArea(req.AreaID, spawns)

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":   true,
		"encounter": result.Encounter,
		"data":      result,
	})
}

// ExitGrassArea 离开草地
// POST /api/map/exit-grass
func (h *Handler) ExitGrassArea(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
			"code":    "METHOD_NOT_ALLOWED",
		})
		return
	}

	var req ExitGrassAreaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
			"code":    "INVALID_REQUEST",
		})
		return
	}

	if req.AreaID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "area_id is required",
			"code":    "INVALID_AREA_ID",
		})
		return
	}

	if req.UserID <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "user_id must be a positive integer",
			"code":    "INVALID_USER_ID",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Successfully exited grass area",
		"code":    "EXIT_SUCCESS",
	})
}

// GetRegionsResponse 地区响应
type GetRegionsResponse struct {
	ID            string `json:"id"`
	NameEn        string `json:"name_en"`
	NameZh        string `json:"name_zh"`
	Description   string `json:"description"`
	LevelRangeMin int    `json:"level_range_min"`
	LevelRangeMax int    `json:"level_range_max"`
}

// GetRegions 获取所有地区
// GET /api/map/regions
func (h *Handler) GetRegions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
			"code":    "METHOD_NOT_ALLOWED",
		})
		return
	}

	explorationSystem := service.NewMapExplorationSystem(h.db)
	regions := explorationSystem.GetAllRegions()

	// Convert to response format
	var regionList []GetRegionsResponse
	for _, region := range regions {
		regionList = append(regionList, GetRegionsResponse{
			ID:            region.ID,
			NameEn:        region.NameEn,
			NameZh:        region.NameZh,
			Description:   region.Description,
			LevelRangeMin: region.LevelRangeMin,
			LevelRangeMax: region.LevelRangeMax,
		})
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    regionList,
		"code":    "GET_REGIONS_SUCCESS",
	})
}

// GetGrassAreasResponse 草地响应
type GetGrassAreasResponse struct {
	ID         string `json:"id"`
	RegionID   string `json:"region_id"`
	NameEn     string `json:"name_en"`
	NameZh     string `json:"name_zh"`
	XCoord     int    `json:"x_coord"`
	YCoord     int    `json:"y_coord"`
	Difficulty int    `json:"difficulty"`
}

// GetGrassAreas 获取特定地区的草地
// GET /api/map/grass-areas?region_id=region001
func (h *Handler) GetGrassAreas(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
			"code":    "METHOD_NOT_ALLOWED",
		})
		return
	}

	regionID := r.URL.Query().Get("region_id")
	if regionID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "region_id parameter required",
			"code":    "INVALID_REGION_ID",
		})
		return
	}

	explorationSystem := service.NewMapExplorationSystem(h.db)
	areas := explorationSystem.GetGrassAreasForRegion(regionID)

	// Convert to response format
	var areaList []GetGrassAreasResponse
	for _, area := range areas {
		areaList = append(areaList, GetGrassAreasResponse{
			ID:         area.ID,
			RegionID:   area.RegionID,
			NameEn:     area.NameEn,
			NameZh:     area.NameZh,
			XCoord:     area.XCoord,
			YCoord:     area.YCoord,
			Difficulty: area.Difficulty,
		})
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    areaList,
		"code":    "GET_GRASS_AREAS_SUCCESS",
	})
}

// GetExplorationLog 获取用户的探索日志
// GET /api/map/exploration-log?user_id=123
func (h *Handler) GetExplorationLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
			"code":    "METHOD_NOT_ALLOWED",
		})
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "user_id parameter required",
			"code":    "INVALID_USER_ID",
		})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "user_id must be a positive integer",
			"code":    "INVALID_USER_ID",
		})
		return
	}

	// In real implementation, retrieve from database
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"user_id": userID,
		"data":    []interface{}{},
		"code":    "GET_LOG_SUCCESS",
	})
}

// GetMapStatus 获取地图状态
// GET /api/map/status
func (h *Handler) GetMapStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
			"code":    "METHOD_NOT_ALLOWED",
		})
		return
	}

	explorationSystem := service.NewMapExplorationSystem(h.db)

	// Get summary of regions and grass areas
	regions := explorationSystem.GetAllRegions()

	totalAreas := 0
	for _, region := range regions {
		areas := explorationSystem.GetGrassAreasForRegion(region.ID)
		totalAreas += len(areas)
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"total_regions":     len(regions),
			"total_grass_areas": totalAreas,
		},
		"code": "GET_STATUS_SUCCESS",
	})
}
