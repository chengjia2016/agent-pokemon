package handler

import (
	"encoding/json"
	"judge-server/internal/service"
	"net/http"
	"strconv"
)

// PokemonCaptureHandler 宠物捕获处理器
type PokemonCaptureHandler struct {
	pokemonManager *service.PokemonManager
	worldService   *service.WorldService
}

func NewPokemonCaptureHandler(pm *service.PokemonManager) *PokemonCaptureHandler {
	return &PokemonCaptureHandler{
		pokemonManager: pm,
	}
}

func NewPokemonCaptureHandlerWithWorldService(pm *service.PokemonManager, ws *service.WorldService) *PokemonCaptureHandler {
	return &PokemonCaptureHandler{
		pokemonManager: pm,
		worldService:   ws,
	}
}

// ==================== 野生宠物列表 ====================

// ListWildPokemon 列出野生宠物
// GET /api/wild-pokemon?zone_code=ZONE_001
func (pch *PokemonCaptureHandler) ListWildPokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// 获取 zone_code 参数
	zoneCode := r.URL.Query().Get("zone_code")
	if zoneCode == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "zone_code parameter required",
		})
		return
	}

	// 检查是否有 worldService
	if pch.worldService == nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "world service not available",
		})
		return
	}

	// 获取该地区的野生 Pokemon
	wildPokemons, err := pch.worldService.GetWildPokemonByZoneCode(zoneCode)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    wildPokemons,
		"total":   len(wildPokemons),
	})
}

// ==================== 捕获宠物 ====================

// CapturePokemon 捕获野生宠物
// POST /api/wild-pokemon/capture
// 请求体: {"github_id": 274799269, "wild_pokemon_id": "wild_zone001_001", "ball_type": "poke_ball"}
func (pch *PokemonCaptureHandler) CapturePokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// 获取用户 ID - 支持 github_id 或 API key 验证
	var githubID int
	var wildPokemonID string
	var ballType = "poke_ball"

	// 首先尝试从 github_id 参数获取（向后兼容）
	if gid, ok := req["github_id"].(float64); ok {
		githubID = int(gid)
	} else {
		// 如果没有 github_id，必须返回错误
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "github_id parameter required",
		})
		return
	}

	if wpid, ok := req["wild_pokemon_id"].(string); ok {
		wildPokemonID = wpid
	} else {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "wild_pokemon_id parameter required",
		})
		return
	}

	if bt, ok := req["ball_type"].(string); ok && bt != "" {
		ballType = bt
	}

	// 检查是否有 worldService
	if pch.worldService == nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "world service not available",
		})
		return
	}

	// 尝试捕捉 Pokemon
	success, userPokemon, err := pch.worldService.CaptureWildPokemonAPI(githubID, wildPokemonID, ballType)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if success {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success":          true,
			"captured":         true,
			"message":          "Pokemon captured successfully!",
			"captured_pokemon": userPokemon,
		})
	} else {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success":  true,
			"captured": false,
			"message":  "Pokemon escaped from the Pokéball!",
		})
	}
}

// ==================== 捕获尝试历史 ====================

// GetCaptureHistory 获取捕获尝试历史
// GET /api/wild-pokemon/capture-history
func (pch *PokemonCaptureHandler) GetCaptureHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// 解析分页参数
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

	// TODO: 从数据库获取捕获历史
	// 当前返回空列表作为占位符
	history := []interface{}{}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    history,
		"limit":   limit,
		"offset":  offset,
		"total":   0,
	})
}
