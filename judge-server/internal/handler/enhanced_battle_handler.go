package handler

import (
	"encoding/json"
	"judge-server/internal/model"
	"judge-server/internal/service"
	"net/http"
	"time"
)

// ==================== 增强战斗处理器 ====================

// EnhancedBattleHandler 增强的战斗处理器，集成所有新的战斗系统
type EnhancedBattleHandler struct {
	enhancedEngine *service.EnhancedBattleEngine
	effectsSystem  *service.BattleEffectsSystem
	strategySystem *service.BattleStrategySystem
	battleEngine   *service.BattleEngineV2
}

// NewEnhancedBattleHandler 创建增强的战斗处理器
func NewEnhancedBattleHandler(
	ee *service.EnhancedBattleEngine,
	es *service.BattleEffectsSystem,
	bs *service.BattleStrategySystem,
	be *service.BattleEngineV2,
) *EnhancedBattleHandler {
	return &EnhancedBattleHandler{
		enhancedEngine: ee,
		effectsSystem:  es,
		strategySystem: bs,
		battleEngine:   be,
	}
}

// ==================== 获取增强的战斗状态 ====================

// GetEnhancedBattleState 获取增强的战斗状态（包含所有UI元素）
// GET /api/battles/{battleId}/enhanced
func (ebh *EnhancedBattleHandler) GetEnhancedBattleState(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// 这里应该从数据库获取战斗信息，此处为示例结构
	enhancedState := &model.EnhancedBattleState{
		BattleID: "battle_example",
		Status:   "ongoing",
		Turn:     5,
		Weather:  "harsh_sunlight",
		Terrain:  "grassy_terrain",
		AttackerState: &model.EnhancedPokemonState{
			PokemonID:    "pikachu",
			Name:         "Pikachu",
			Level:        50,
			CurrentHP:    120,
			MaxHP:        150,
			HPPercentage: 0.8,
			HPBarColor:   "green",
		},
		DefenderState: &model.EnhancedPokemonState{
			PokemonID:    "charizard",
			Name:         "Charizard",
			Level:        55,
			CurrentHP:    80,
			MaxHP:        160,
			HPPercentage: 0.5,
			HPBarColor:   "yellow",
		},
		BattleLog:     []*model.DetailedBattleLogEntry{},
		ActionOptions: &model.BattleActionOptions{},
		UiHints:       &model.BattleUIHints{},
		Statistics:    &model.BattleStatistics{},
		UpdatedAt:     time.Now(),
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"battle":  enhancedState,
	})
}

// ==================== 获取招式动画和效果 ====================

// GetMoveAnimation 获取招式的动画数据
// GET /api/moves/{moveId}/animation
func (ebh *EnhancedBattleHandler) GetMoveAnimation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	// 从URL参数获取招式ID
	moveID := r.URL.Query().Get("move_id")
	if moveID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "move_id is required",
		})
		return
	}

	animation := ebh.effectsSystem.GetMoveAnimation(moveID)

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":   true,
		"animation": animation,
	})
}

// ==================== 获取状态异常效果 ====================

// GetStatusEffectData 获取状态异常的效果数据
// GET /api/status/{conditionId}/effects
func (ebh *EnhancedBattleHandler) GetStatusEffectData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	conditionID := r.URL.Query().Get("condition_id")
	if conditionID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "condition_id is required",
		})
		return
	}

	effectData := ebh.effectsSystem.GetStatusEffectData(conditionID)
	if effectData == nil {
		writeJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "condition not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"effect":  effectData,
	})
}

// ==================== 获取战斗建议 ====================

// GetBattleRecommendation 获取战斗建议
// POST /api/battles/recommendation
type RecommendationRequest struct {
	AttackerPokemon  *model.BattlePokemon   `json:"attacker"`
	AttackerTypes    []string               `json:"attacker_types"`
	DefenderPokemon  *model.BattlePokemon   `json:"defender"`
	DefenderTypes    []string               `json:"defender_types"`
	Weather          string                 `json:"weather"`
	AvailablePokemon []*model.BattlePokemon `json:"available_pokemon,omitempty"`
}

func (ebh *EnhancedBattleHandler) GetBattleRecommendation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req RecommendationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// 生成战斗策略
	strategy := ebh.strategySystem.GenerateBattleStrategy(
		req.AttackerPokemon,
		req.AttackerTypes,
		req.DefenderPokemon,
		req.DefenderTypes,
		req.Weather,
		req.AvailablePokemon,
	)

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":  true,
		"strategy": strategy,
	})
}

// ==================== 计算增强的伤害 ====================

// CalculateDamageWithEffects 计算包含所有效果的伤害
// POST /api/battles/damage-calculation
type DamageCalculationRequest struct {
	Attacker *model.BattlePokemon `json:"attacker"`
	Defender *model.BattlePokemon `json:"defender"`
	Move     *model.PokemonMove   `json:"move"`
	Weather  string               `json:"weather,omitempty"`
	Terrain  string               `json:"terrain,omitempty"`
}

type DamageCalculationResponse struct {
	BaseDamage        int                `json:"base_damage"`
	FinalDamage       int                `json:"final_damage"`
	MinDamage         int                `json:"min_damage"`
	MaxDamage         int                `json:"max_damage"`
	IsCritical        bool               `json:"is_critical"`
	CriticalChance    float32            `json:"critical_chance"`
	Effectiveness     float32            `json:"effectiveness"`
	TypeMatchup       string             `json:"type_matchup"`
	Multipliers       map[string]float32 `json:"multipliers"`
	DefenderHPAfter   int                `json:"defender_hp_after"`
	DefenderHPPercent float32            `json:"defender_hp_percent"`
	IsFatal           bool               `json:"is_fatal"`
}

func (ebh *EnhancedBattleHandler) CalculateDamageWithEffects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	var req DamageCalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	if req.Attacker == nil || req.Defender == nil || req.Move == nil {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "attacker, defender, and move are required",
		})
		return
	}

	// 计算会心一击概率
	isCritical := ebh.battleEngine.CalculateCriticalHit(req.Move, req.Attacker.Ability)

	// 计算增强的伤害
	damage, effectiveness, details := ebh.enhancedEngine.CalculateEnhancedDamage(
		req.Attacker,
		req.Defender,
		req.Move,
		req.Weather,
		isCritical,
		req.Terrain,
	)

	// 计算范围伤害（85-100%）
	minDamage := int(float32(damage) * 0.85)
	maxDamage := damage

	// 计算防御者HP变化
	defenderHPAfter := req.Defender.CurrentHP - damage
	if defenderHPAfter < 0 {
		defenderHPAfter = 0
	}
	defenderHPPercent := float32(defenderHPAfter) / float32(req.Defender.MaxHP)

	// 获取类型相克视觉信息
	typeMatchupVisual := ebh.effectsSystem.GetTypeMatchupVisual(effectiveness)

	response := DamageCalculationResponse{
		BaseDamage:        details["base_damage"].(int),
		FinalDamage:       damage,
		MinDamage:         minDamage,
		MaxDamage:         maxDamage,
		IsCritical:        isCritical,
		CriticalChance:    12.5, // TODO: 从特性和招式计算
		Effectiveness:     effectiveness,
		TypeMatchup:       typeMatchupVisual["description"].(string),
		Multipliers:       map[string]float32{},
		DefenderHPAfter:   defenderHPAfter,
		DefenderHPPercent: defenderHPPercent,
		IsFatal:           defenderHPAfter == 0,
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success":            true,
		"damage_calculation": response,
	})
}

// ==================== 获取类型相克信息 ====================

// GetTypeMatchupInfo 获取类型相克信息
// GET /api/types/matchup
type TypeMatchupQuery struct {
	AttackerTypes []string `json:"attacker_types"`
	DefenderTypes []string `json:"defender_types"`
}

func (ebh *EnhancedBattleHandler) GetTypeMatchupInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	attackType := r.URL.Query().Get("attack_type")
	defendType := r.URL.Query().Get("defend_type")

	if attackType == "" || defendType == "" {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "attack_type and defend_type are required",
		})
		return
	}

	// 这里应该从类型系统获取相克信息
	// 示例响应
	response := map[string]interface{}{
		"attack_type":   attackType,
		"defend_type":   defendType,
		"effectiveness": 1.0,
		"description":   "normal",
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"matchup": response,
	})
}

// ==================== 注册路由 ====================

// RegisterEnhancedBattleRoutes 注册增强战斗相关的路由
func RegisterEnhancedBattleRoutes(mux *http.ServeMux, handler *EnhancedBattleHandler) {
	mux.HandleFunc("/api/battles/enhanced", handler.GetEnhancedBattleState)
	mux.HandleFunc("/api/moves/animation", handler.GetMoveAnimation)
	mux.HandleFunc("/api/status/effects", handler.GetStatusEffectData)
	mux.HandleFunc("/api/battles/recommendation", handler.GetBattleRecommendation)
	mux.HandleFunc("/api/battles/damage-calculation", handler.CalculateDamageWithEffects)
	mux.HandleFunc("/api/types/matchup", handler.GetTypeMatchupInfo)
}
