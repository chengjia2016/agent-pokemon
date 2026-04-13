package handler

import (
	"encoding/json"
	"judge-server/internal/model"
	"judge-server/internal/service"
	"net/http"
	"strconv"
)

// PokemonCaptureHandler 宠物捕获处理器
type PokemonCaptureHandler struct {
	pokemonManager *service.PokemonManager
}

func NewPokemonCaptureHandler(pm *service.PokemonManager) *PokemonCaptureHandler {
	return &PokemonCaptureHandler{
		pokemonManager: pm,
	}
}

// ==================== 野生宠物列表 ====================

// ListWildPokemon 列出野生宠物
// GET /api/wild-pokemon
func (pch *PokemonCaptureHandler) ListWildPokemon(w http.ResponseWriter, r *http.Request) {
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

	// TODO: 从数据库获取野生宠物列表
	// 当前返回空列表作为占位符
	wildPokemons := []interface{}{}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    wildPokemons,
		"limit":   limit,
		"offset":  offset,
		"total":   0,
	})
}

// ==================== 捕获宠物 ====================

// CapturePokemon 捕获野生宠物
// POST /api/wild-pokemon/capture
func (pch *PokemonCaptureHandler) CapturePokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req model.CapturePokemonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// TODO: 从认证token获取用户ID
	githubID := 1

	// 检查是否可以添加宠物（限制10只）
	canAdd, err := pch.pokemonManager.CanAddPokemon(githubID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to check pokemon limit",
		})
		return
	}

	if !canAdd {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Pokemon limit reached (max 10)",
		})
		return
	}

	// TODO: 实现捕获逻辑
	// 1. 获取野生宠物
	// 2. 更新捕获记录
	// 3. 将宠物添加到用户账户

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Pokemon captured successfully",
	})
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
