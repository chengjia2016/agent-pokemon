package handler

import (
	"encoding/json"
	"fmt"
	"judge-server/internal/model"
	"judge-server/internal/service"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

// BattleHandler 战斗处理器
type BattleHandler struct {
	battleEngine     *service.BattleEngine
	rewardCalculator *service.RewardCalculator
	pokemonManager   *service.PokemonManager
}

func NewBattleHandler(be *service.BattleEngine, rc *service.RewardCalculator, pm *service.PokemonManager) *BattleHandler {
	return &BattleHandler{
		battleEngine:     be,
		rewardCalculator: rc,
		pokemonManager:   pm,
	}
}

// ==================== 战斗开始 ====================

// StartBattle 开始新的战斗
// POST /api/battles/start
func (bh *BattleHandler) StartBattle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req model.StartBattleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// 从请求中获取用户ID
	attackerID := req.AttackerID
	defenderID := req.DefenderID

	// 验证用户ID
	if attackerID == 0 || defenderID == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "attacker_id and defender_id are required",
		})
		return
	}

	// 验证战斗请求
	if err := bh.battleEngine.ValidateBattleRequest(attackerID, defenderID, req.AttackerTeamID, req.DefenderTeamID); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// 创建战斗
	battleID := fmt.Sprintf("battle_%s", uuid.New().String())
	battle, err := bh.battleEngine.DB.CreateBattle(battleID, attackerID, defenderID, req.AttackerTeamID, req.DefenderTeamID, req.BattleType)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Failed to create battle",
		})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"battle":  battle,
	})
}

// GetBattle 获取战斗信息
// GET /api/battles/{battleId}
func (bh *BattleHandler) GetBattle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// TODO: 从URL路径提取battleId
	battleID := r.URL.Query().Get("battle_id")
	if battleID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "battle_id parameter required",
		})
		return
	}

	// TODO: 从数据库获取战斗信息
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    nil,
	})
}

// ==================== 战斗回合 ====================

// ExecuteBattleRound 执行战斗回合
// POST /api/battles/{battleId}/round
func (bh *BattleHandler) ExecuteBattleRound(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	battleID := r.URL.Query().Get("battle_id")
	if battleID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "battle_id parameter required",
		})
		return
	}

	// TODO: 实现战斗回合逻辑
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Round executed",
	})
}

// ==================== 战斗结束 ====================

// EndBattle 结束战斗
// POST /api/battles/{battleId}/end
func (bh *BattleHandler) EndBattle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	battleID := r.URL.Query().Get("battle_id")
	if battleID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "battle_id parameter required",
		})
		return
	}

	// TODO: 实现战斗结束逻辑
	// 1. 确定赢家
	// 2. 计算和转账奖励
	// 3. 更新战斗状态

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Battle ended",
	})
}

// ==================== 战斗列表 ====================

// GetUserBattles 获取用户的战斗列表
// GET /api/battles
func (bh *BattleHandler) GetUserBattles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// TODO: 从认证token获取用户ID
	// TODO: 从数据库查询用户的战斗列表
	// 支持分页和过滤

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

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    []interface{}{},
		"limit":   limit,
		"offset":  offset,
		"total":   0,
	})
}

// ==================== 战斗统计 ====================

// GetBattleStats 获取战斗统计
// GET /api/battles/stats
func (bh *BattleHandler) GetBattleStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// TODO: 从认证token获取用户ID
	// TODO: 从数据库计算统计数据

	stats := map[string]interface{}{
		"total_battles":      0,
		"wins":               0,
		"losses":             0,
		"draws":              0,
		"win_rate":           0.0,
		"total_coins_gained": 0.0,
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"stats":   stats,
	})
}
