package service

import (
	"errors"
	"fmt"
	"judge-server/internal/model"
	"math/rand"
)

// BattleEngineV2 改进的战斗引擎（支持完整数值系统）
type BattleEngineV2 struct {
	breedingService *PokemonBreedingService
	db              Database
}

// NewBattleEngineV2 创建新的战斗引擎
func NewBattleEngineV2(bs *PokemonBreedingService, db Database) *BattleEngineV2 {
	return &BattleEngineV2{
		breedingService: bs,
		db:              db,
	}
}

// ==================== 伤害计算 ====================

// CalculateDamage 计算战斗中的伤害
// 宝可梦伤害计算公式（简化版）：
// ((((2 * AttackStat / 5 + 2) * Power * DefenseStat / 50) / 50) + 2) * Modifier
func (be *BattleEngineV2) CalculateDamage(
	attacker *model.BattlePokemon,
	defender *model.BattlePokemon,
	move *model.PokemonMove,
	weather string,
	isCritical bool,
) (int, float32, error) {
	if attacker == nil || defender == nil || move == nil {
		return 0, 0, errors.New("invalid battle pokemon or move")
	}

	if move.Power == 0 {
		return 0, 1.0, nil // 变化类技能不造成伤害
	}

	// 基础伤害计算
	var attackStat int
	var defenseStat int

	if move.Category == "physical" {
		attackStat = attacker.Stats.Attack
		defenseStat = defender.Stats.Defense
	} else if move.Category == "special" {
		attackStat = attacker.Stats.SpAtk
		defenseStat = defender.Stats.SpDef
	} else {
		return 0, 1.0, nil
	}

	// 计算基础伤害
	baseDamage := ((((2*attackStat/5 + 2) * move.Power * defenseStat / 50) / 50) + 2)

	// 获取类型克制倍数
	effectiveness := be.breedingService.GetTypeMatchup(move.Type, defender.PokemonID)
	if effectiveness == 0 {
		return 0, 0, nil // 完全免疫
	}

	// 应用会心一击加成
	criticalMultiplier := float32(1.0)
	if isCritical {
		criticalMultiplier = 1.5
	}

	// 应用天气加成
	weatherMultiplier := be.getWeatherMultiplier(move.Type, weather)

	// 应用随机浮动（85-100%）
	randomMultiplier := float32(85+rand.Intn(16)) / 100.0

	// 计算最终伤害
	damage := int(float32(baseDamage) * effectiveness * criticalMultiplier * weatherMultiplier * randomMultiplier)

	// 确保伤害至少为1
	if damage < 1 {
		damage = 1
	}

	return damage, effectiveness, nil
}

// getWeatherMultiplier 根据天气获取威力倍数
func (be *BattleEngineV2) getWeatherMultiplier(moveType, weather string) float32 {
	weatherBoosts := map[string]map[string]float32{
		"rain": {
			"water": 1.5,
			"fire":  0.5,
		},
		"harsh_sunlight": {
			"fire":  1.5,
			"water": 0.5,
			"grass": 1.5,
		},
		"hail": {
			"ice": 1.5,
		},
		"sandstorm": {
			"rock":   1.5,
			"ground": 1.5,
			"steel":  1.5,
		},
	}

	if boost, exists := weatherBoosts[weather][moveType]; exists {
		return boost
	}
	return 1.0
}

// ==================== 优先级和速度计算 ====================

// DetermineMoveOrder 确定招式执行顺序
// 优先级高的招式先执行，相同优先级按速度排序
func (be *BattleEngineV2) DetermineMoveOrder(
	move1 *model.BattleMove,
	move2 *model.BattleMove,
	attacker1 *model.BattlePokemon,
	attacker2 *model.BattlePokemon,
	move1Data *model.PokemonMove,
	move2Data *model.PokemonMove,
) int {
	// 比较优先级
	if move1Data.Priority != move2Data.Priority {
		return move1Data.Priority - move2Data.Priority
	}

	// 优先级相同，比较速度
	if attacker1.Stats.Speed != attacker2.Stats.Speed {
		return attacker1.Stats.Speed - attacker2.Stats.Speed
	}

	// 速度相同，随机决定
	if rand.Intn(2) == 0 {
		return 1
	}
	return -1
}

// ==================== 会心一击 ====================

// CalculateCriticalHit 计算会心一击
// 大多数招式会心率为12.5%
// 有些招式或特性会增加会心率
func (be *BattleEngineV2) CalculateCriticalHit(
	move *model.PokemonMove,
	ability *model.PokemonAbility,
) bool {
	criticalRate := 12.5 // 默认12.5%

	// 某些招式有更高的会心率
	highCritMoves := map[string]bool{
		"slash":             true,
		"psycho_cut":        true,
		"night_slash":       true,
		"leaf_blade":        true,
		"stone_edge":        true,
		"air_slash":         true,
		"cross_poison":      true,
		"shadow_claw":       true,
		"cross_chop":        true,
		"x_scissor":         true,
		"aqua_jet":          true,
		"dragon_rush":       true,
		"poison_powder_cut": true,
	}

	if highCritMoves[move.ID] {
		criticalRate = 50.0
	}

	// 某些特性增加会心率
	if ability != nil {
		switch ability.AbilityID {
		case "super_luck":
			criticalRate *= 1.5
		case "scope_lens":
			criticalRate *= 2.0
		}
	}

	// 计算是否发生会心一击
	randomValue := rand.Float64() * 100
	return randomValue < criticalRate
}

// ==================== 状态效果 ====================

// ApplyStatusCondition 应用异常状态
func (be *BattleEngineV2) ApplyStatusCondition(
	pokemon *model.BattlePokemon,
	condition string,
	durationTurns int,
) error {
	// 检查宝可梦是否已有该状态
	for _, existingCondition := range pokemon.Conditions {
		if existingCondition == condition {
			return errors.New("pokemon already has this condition")
		}
	}

	// 检查宝可梦能否被施加此状态
	if !be.canApplyCondition(pokemon, condition) {
		return errors.New("pokemon is immune to this condition")
	}

	pokemon.Conditions = append(pokemon.Conditions, condition)
	return nil
}

// canApplyCondition 检查宝可梦是否能被施加条件
func (be *BattleEngineV2) canApplyCondition(pokemon *model.BattlePokemon, condition string) bool {
	// 钢系宝可梦免疫中毒
	if condition == "status_poison" {
		// TODO: 检查宝可梦类型
	}

	// 冰系宝可梦免疫冰冻
	if condition == "status_freeze" {
		// TODO: 检查宝可梦类型
	}

	// 火系宝可梦免疫灼伤
	if condition == "status_burn" {
		// TODO: 检查宝可梦类型
	}

	return true
}

// ApplyStatusEffect 在战斗中应用状态效果
func (be *BattleEngineV2) ApplyStatusEffect(
	pokemon *model.BattlePokemon,
	condition string,
	attacker *model.BattlePokemon,
) (bool, string) {
	switch condition {
	case "status_paralysis":
		// 25%概率无法行动
		if rand.Intn(100) < 25 {
			return false, "pokemon is paralyzed and cannot move"
		}
		// 速度降低25%
		pokemon.Stats.Speed = int(float32(pokemon.Stats.Speed) * 0.75)

	case "status_sleep":
		return false, "pokemon is asleep"

	case "status_confusion":
		// 33%概率无法行动并自伤
		if rand.Intn(100) < 33 {
			damage := int(float32(pokemon.MaxHP) * 0.25)
			pokemon.CurrentHP -= damage
			return false, fmt.Sprintf("pokemon is confused and hurt itself for %d damage", damage)
		}

	case "status_burn":
		// 每回合损失1/8最大HP
		damage := pokemon.MaxHP / 8
		if damage < 1 {
			damage = 1
		}
		pokemon.CurrentHP -= damage
		// 物理攻击力减半
		pokemon.Stats.Attack = int(float32(pokemon.Stats.Attack) * 0.5)

	case "status_poison":
		// 每回合损失1/8最大HP
		damage := pokemon.MaxHP / 8
		if damage < 1 {
			damage = 1
		}
		pokemon.CurrentHP -= damage
	}

	return true, ""
}

// ==================== 战斗模拟 ====================

// SimulateBattle 模拟一场战斗
type BattleSimulation struct {
	Pokemon1    *model.BattlePokemon
	Pokemon2    *model.BattlePokemon
	Winner      *model.BattlePokemon
	TotalRounds int
	BattleLog   []string
	ExpGained   int
	EVsGained   map[string]int
}

// SimulateBattleRound 模拟单一回合
func (be *BattleEngineV2) SimulateBattleRound(
	attacker, defender *model.BattlePokemon,
	move *model.PokemonMove,
	weather string,
) (int, float32, error) {
	// 应用攻击者的状态效果
	canAttack, msg := be.ApplyStatusEffect(attacker, "", nil)
	if !canAttack {
		return 0, 0, errors.New(msg)
	}

	// 计算是否会心一击
	isCritical := be.CalculateCriticalHit(move, attacker.Ability)

	// 计算伤害
	damage, effectiveness, err := be.CalculateDamage(
		attacker,
		defender,
		move,
		weather,
		isCritical,
	)

	if err != nil {
		return 0, 0, err
	}

	// 应用伤害
	defender.CurrentHP -= damage
	if defender.CurrentHP < 0 {
		defender.CurrentHP = 0
	}

	return damage, effectiveness, nil
}

// ==================== 经验和努力值 ====================

// CalculateExperience 计算获得的经验值
func (be *BattleEngineV2) CalculateExperience(
	opponentLevel int,
	isBossBattle bool,
) int {
	baseExp := opponentLevel * 10
	if isBossBattle {
		baseExp *= 2
	}
	return baseExp
}

// GetEffortValueReward 获取宝可梦被击败后的努力值奖励
func (be *BattleEngineV2) GetEffortValueReward(pokemonID string) map[string]int {
	// 不同宝可梦给予不同的EV奖励
	evRewards := map[string]map[string]int{
		"pikachu": {
			"attack": 1,
			"speed":  1,
		},
		"charizard": {
			"sp_atk": 3,
		},
		"blastoise": {
			"defense": 2,
			"sp_def":  1,
		},
		// TODO: 添加更多宝可梦的EV奖励
	}

	if reward, exists := evRewards[pokemonID]; exists {
		return reward
	}

	// 默认奖励
	return map[string]int{
		"hp": 1,
	}
}
