package handler

import (
	"encoding/json"
	"judge-server/internal/model"
	"judge-server/internal/service"
	"net/http"
	"strconv"
)

// DefenseHandler 防守基地处理器
type DefenseHandler struct {
	defenseSystem *service.DefenseSystem
}

func NewDefenseHandler(ds *service.DefenseSystem) *DefenseHandler {
	return &DefenseHandler{
		defenseSystem: ds,
	}
}

// ==================== 基地创建 ====================

// CreateBase 创建防守基地
// POST /api/defense/base
func (dh *DefenseHandler) CreateBase(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req model.CreateBaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// TODO: 从认证token获取用户ID
	githubID := 1

	// 创建基地
	base, err := dh.defenseSystem.CreateBase(githubID, req.RepositoryURL)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"base":    base,
	})
}

// GetBase 获取防守基地信息
// GET /api/defense/base
func (dh *DefenseHandler) GetBase(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// TODO: 从认证token获取用户ID
	githubID := 1

	// 获取基地
	base, err := dh.defenseSystem.DB.GetBase(githubID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get base",
		})
		return
	}

	if base == nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Base not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"base":    base,
	})
}

// ==================== 防守者列表 ====================

// GetDefenders 获取基地的防守者（队伍中的宠物）
// GET /api/defense/defenders
func (dh *DefenseHandler) GetDefenders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// TODO: 从URL参数获取基地ID或用户ID
	githubID := 1

	// 获取基地
	base, err := dh.defenseSystem.DB.GetBase(githubID)
	if err != nil || base == nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Base not found",
		})
		return
	}

	// 获取防守队伍成员
	defenders, err := dh.defenseSystem.DB.GetTeamMembers(*base.DefenseTeamID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to get defenders",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":   true,
		"defenders": defenders,
	})
}

// ==================== 防守历史 ====================

// GetDefenseHistory 获取基地的防守历史
// GET /api/defense/history
func (dh *DefenseHandler) GetDefenseHistory(w http.ResponseWriter, r *http.Request) {
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

	// 获取基地
	base, err := dh.defenseSystem.DB.GetBase(githubID)
	if err != nil || base == nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Base not found",
		})
		return
	}

	// TODO: 从数据库获取防守历史记录
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

// ==================== 防守统计 ====================

// GetDefenseStats 获取防守统计
// GET /api/defense/stats
func (dh *DefenseHandler) GetDefenseStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// TODO: 从认证token获取用户ID
	githubID := 1

	// 获取基地
	base, err := dh.defenseSystem.DB.GetBase(githubID)
	if err != nil || base == nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "Base not found",
		})
		return
	}

	// 获取防守统计
	stats := dh.defenseSystem.GetDefenseStats(base)

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"stats":   stats,
	})
}
