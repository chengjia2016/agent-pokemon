package handler

import (
	"judge-server/internal/service"
	"net/http"
	"strconv"
)

// PokemonManagementHandler 宠物管理处理器
type PokemonManagementHandler struct {
	pokemonManager *service.PokemonManager
}

func NewPokemonManagementHandler(pm *service.PokemonManager) *PokemonManagementHandler {
	return &PokemonManagementHandler{
		pokemonManager: pm,
	}
}

// ==================== 宠物列表 ====================

// ListUserPokemon 列出用户的宠物
// GET /api/pokemon
func (pmh *PokemonManagementHandler) ListUserPokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// TODO: 从认证token获取用户ID
	githubID := 1

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

	// 获取用户的宠物
	pokemons, err := pmh.pokemonManager.GetUserPokemons(githubID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get pokemons",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    pokemons,
		"limit":   limit,
		"offset":  offset,
		"total":   len(pokemons),
	})
}

// ==================== 宠物详情 ====================

// GetPokemonDetails 获取单个宠物的详情
// GET /api/pokemon/{petId}
func (pmh *PokemonManagementHandler) GetPokemonDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// TODO: 从URL路径提取petId
	petID := r.URL.Query().Get("pet_id")
	if petID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "pet_id parameter required",
		})
		return
	}

	// TODO: 从数据库获取单个宠物详情
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    nil,
	})
}

// ==================== 宠物释放 ====================

// ReleasePokemon 释放宠物
// DELETE /api/pokemon/{petId}
func (pmh *PokemonManagementHandler) ReleasePokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// TODO: 从URL路径提取petId
	petID := r.URL.Query().Get("pet_id")
	if petID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "pet_id parameter required",
		})
		return
	}

	// TODO: 实现释放宠物逻辑
	// 1. 检查宠物是否属于当前用户
	// 2. 从队伍中移除宠物（如果存在）
	// 3. 删除宠物记录

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Pokemon released successfully",
	})
}

// ==================== 宠物恢复 ====================

// GetRecoveryStatus 获取宠物恢复状态
// GET /api/pokemon/recovery-status
func (pmh *PokemonManagementHandler) GetRecoveryStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// TODO: 从认证token获取用户ID
	githubID := 1

	// 获取用户的宠物恢复状态
	status, err := pmh.pokemonManager.GetRecoveryStatus(githubID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get recovery status",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    status,
	})
}

// ==================== 宠物训练 ====================

// TrainPokemon 训练宠物
// POST /api/pokemon/{petId}/train
func (pmh *PokemonManagementHandler) TrainPokemon(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// TODO: 从URL路径提取petId
	petID := r.URL.Query().Get("pet_id")
	if petID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "pet_id parameter required",
		})
		return
	}

	// TODO: 实现训练逻辑
	// 1. 检查宠物是否属于当前用户
	// 2. 计算经验值增加
	// 3. 检查是否升级
	// 4. 更新宠物属性

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Pokemon trained successfully",
	})
}
