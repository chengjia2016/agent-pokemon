package service

import (
	"judge-server/internal/model"
	"sort"
)

// ==================== 招式策略系统 ====================
// BattleStrategySystem 提供智能的战斗建议和策略
type BattleStrategySystem struct {
	typeMatchupSystem *TypeMatchupSystem
	effectsSystem     *BattleEffectsSystem
}

// TypeMatchupSystem 类型相克系统
type TypeMatchupSystem struct {
	matchups map[string]map[string]float32
}

// NewBattleStrategySystem 创建战斗策略系统
func NewBattleStrategySystem() *BattleStrategySystem {
	return &BattleStrategySystem{
		typeMatchupSystem: NewTypeMatchupSystem(),
		effectsSystem:     NewBattleEffectsSystem(),
	}
}

// NewTypeMatchupSystem 创建类型相克系统
func NewTypeMatchupSystem() *TypeMatchupSystem {
	return &TypeMatchupSystem{
		matchups: initializeTypeMatchups(),
	}
}

// ==================== 类型相克表 ====================

func initializeTypeMatchups() map[string]map[string]float32 {
	return map[string]map[string]float32{
		// 普通型
		"normal": {
			"rock":  0.5,
			"steel": 0.5,
			"ghost": 0.0,
		},
		// 火型
		"fire": {
			"grass":  2.0,
			"bug":    2.0,
			"steel":  2.0,
			"ice":    2.0,
			"fairy":  2.0,
			"water":  0.5,
			"fire":   0.5,
			"rock":   0.5,
			"dragon": 0.5,
		},
		// 水型
		"water": {
			"fire":   2.0,
			"rock":   2.0,
			"ground": 2.0,
			"grass":  0.5,
			"water":  0.5,
			"ice":    0.5,
		},
		// 草型
		"grass": {
			"water":  2.0,
			"rock":   2.0,
			"ground": 2.0,
			"fire":   0.5,
			"grass":  0.5,
			"poison": 0.5,
			"flying": 0.5,
			"bug":    0.5,
		},
		// 电型
		"electric": {
			"water":    2.0,
			"flying":   2.0,
			"grass":    0.5,
			"electric": 0.5,
			"dragon":   0.5,
		},
		// 冰型
		"ice": {
			"grass":  2.0,
			"flying": 2.0,
			"ground": 2.0,
			"dragon": 2.0,
			"fire":   0.5,
			"water":  0.5,
			"ice":    0.5,
		},
		// 格斗型
		"fighting": {
			"normal":  2.0,
			"rock":    2.0,
			"steel":   2.0,
			"ice":     2.0,
			"dark":    2.0,
			"flying":  0.5,
			"poison":  0.5,
			"psychic": 0.5,
			"bug":     0.5,
			"fairy":   0.5,
		},
		// 毒型
		"poison": {
			"grass":  2.0,
			"fairy":  2.0,
			"poison": 0.5,
			"rock":   0.5,
			"ghost":  0.5,
			"steel":  0.0,
		},
		// 地面型
		"ground": {
			"fire":   2.0,
			"poison": 2.0,
			"rock":   2.0,
			"steel":  2.0,
			"grass":  0.5,
			"bug":    0.5,
			"flying": 0.0,
		},
		// 飞行型
		"flying": {
			"grass":    2.0,
			"fighting": 2.0,
			"bug":      2.0,
			"rock":     0.5,
			"steel":    0.5,
			"electric": 0.5,
		},
		// 超能力型
		"psychic": {
			"fighting": 2.0,
			"poison":   2.0,
			"psychic":  0.5,
			"steel":    0.5,
			"dark":     0.0,
		},
		// 虫型
		"bug": {
			"grass":    2.0,
			"psychic":  2.0,
			"dark":     2.0,
			"fire":     0.5,
			"fighting": 0.5,
			"poison":   0.5,
			"flying":   0.5,
			"ghost":    0.5,
			"steel":    0.5,
			"fairy":    0.5,
		},
		// 岩石型
		"rock": {
			"flying":   2.0,
			"bug":      2.0,
			"fire":     2.0,
			"ice":      2.0,
			"fighting": 0.5,
			"ground":   0.5,
			"steel":    0.5,
		},
		// 幽灵型
		"ghost": {
			"ghost":   2.0,
			"psychic": 2.0,
			"normal":  0.0,
			"dark":    0.5,
		},
		// 龙型
		"dragon": {
			"dragon": 2.0,
			"fairy":  0.5,
			"steel":  0.5,
		},
		// 恶型
		"dark": {
			"ghost":    2.0,
			"psychic":  2.0,
			"fighting": 0.5,
			"dark":     0.5,
			"fairy":    0.5,
		},
		// 钢型
		"steel": {
			"rock":     2.0,
			"ice":      2.0,
			"fairy":    2.0,
			"fire":     0.5,
			"water":    0.5,
			"electric": 0.5,
			"grass":    0.5,
			"psychic":  0.5,
			"bug":      0.5,
			"flying":   0.5,
			"dragon":   0.5,
		},
		// 妖精型
		"fairy": {
			"fighting": 2.0,
			"dragon":   2.0,
			"dark":     2.0,
			"poison":   0.5,
			"steel":    0.5,
		},
	}
}

// ==================== 招式策略分析 ====================

// AnalyzeMoveEffectiveness 分析招式对目标的有效性
func (bss *BattleStrategySystem) AnalyzeMoveEffectiveness(
	move *model.PokemonMove,
	targetTypes []string,
) *model.TypeMatchupInfo {
	matchupInfo := &model.TypeMatchupInfo{
		Effective:    []string{},
		NotEffective: []string{},
		Resistant:    []string{},
		Weak:         []string{},
		Immune:       []string{},
	}

	moveType := move.Type
	matchups := bss.typeMatchupSystem.matchups[moveType]

	for _, targetType := range targetTypes {
		multiplier, exists := matchups[targetType]
		if !exists {
			multiplier = 1.0
		}

		switch {
		case multiplier >= 2.0:
			matchupInfo.Effective = append(matchupInfo.Effective, targetType)
		case multiplier == 0.0:
			matchupInfo.Immune = append(matchupInfo.Immune, targetType)
		case multiplier < 1.0 && multiplier > 0:
			matchupInfo.NotEffective = append(matchupInfo.NotEffective, targetType)
		}
	}

	// 计算对目标的总体伤害倍数
	totalMultiplier := float32(1.0)
	for _, targetType := range targetTypes {
		if mult, exists := matchups[targetType]; exists {
			totalMultiplier *= mult
		}
	}
	matchupInfo.Multiplier = totalMultiplier

	return matchupInfo
}

// AnalyzeDefenderWeakness 分析防守方的弱点
func (bss *BattleStrategySystem) AnalyzeDefenderWeakness(
	defenderTypes []string,
	defenderConditions []string,
) *model.TypeWeaknessData {
	weakness := &model.TypeWeaknessData{
		Weak:      make(map[string]float32),
		Resistant: make(map[string]float32),
		Immune:    []string{},
	}

	// 遍历所有类型，找出对目标有利的
	for attackType := range bss.typeMatchupSystem.matchups {
		matchups := bss.typeMatchupSystem.matchups[attackType]

		for _, defType := range defenderTypes {
			if mult, exists := matchups[defType]; exists {
				if mult >= 2.0 {
					weakness.Weak[attackType] = mult
				} else if mult == 0.0 {
					weakness.Immune = append(weakness.Immune, attackType)
				} else if mult < 1.0 && mult > 0 {
					weakness.Resistant[attackType] = mult
				}
			}
		}
	}

	return weakness
}

// ==================== 招式推荐系统 ====================

// MoveRecommendationScore 招式推荐评分
type MoveRecommendationScore struct {
	Move     *model.PokemonMove
	Score    float32
	Reasons  []string
	Priority int
	Icon     string
}

// RecommendBestMove 推荐最佳招式
func (bss *BattleStrategySystem) RecommendBestMove(
	attacker *model.BattlePokemon,
	defender *model.BattlePokemon,
	defenderTypes []string,
	defenderWeakness *model.TypeWeaknessData,
	weatherBoost string,
) *MoveRecommendationScore {

	scores := make([]*MoveRecommendationScore, 0)

	for _, move := range attacker.Moves {
		score := &MoveRecommendationScore{
			Move:    move,
			Score:   0,
			Reasons: []string{},
		}

		// 检查招式是否可用（有PP）
		if move.Power == 0 && move.Category == "status" {
			// 变化类招式
			score.Score += 10
			score.Reasons = append(score.Reasons, "可用的变化招式")
			score.Icon = "🎯"
		} else if move.Power > 0 {
			// 伤害招式
			// 1. 类型匹配加分
			matchupInfo := bss.AnalyzeMoveEffectiveness(move, defenderTypes)
			if matchupInfo.Multiplier >= 4.0 {
				score.Score += 100
				score.Reasons = append(score.Reasons, "超级有效！")
				score.Icon = "🔴"
				score.Priority = 10
			} else if matchupInfo.Multiplier >= 2.0 {
				score.Score += 50
				score.Reasons = append(score.Reasons, "有效")
				score.Icon = "🟠"
				score.Priority = 8
			} else if matchupInfo.Multiplier == 0.0 {
				score.Score -= 100
				score.Reasons = append(score.Reasons, "无效（免疫）")
				score.Icon = "❌"
				score.Priority = 1
				continue
			} else if matchupInfo.Multiplier < 1.0 {
				score.Score -= 20
				score.Reasons = append(score.Reasons, "效果一般")
				score.Icon = "🔵"
				score.Priority = 3
			} else {
				score.Score += 20
				score.Icon = "⚪"
				score.Priority = 5
			}

			// 2. 威力和精准度加分
			score.Score += float32(move.Power) * 0.1
			if move.Accuracy < 100 {
				score.Score -= float32(100-move.Accuracy) * 0.2
			}

			// 3. 会心率高的招式额外加分
			if move.Priority > 0 || move.ID == "slash" || move.ID == "psycho_cut" {
				score.Score += 15
				score.Reasons = append(score.Reasons, "会心率高")
			}

			// 4. 天气加成
			if weatherBoost != "" {
				weatherMultiplier := getWeatherMultiplier(move.Type, weatherBoost)
				if weatherMultiplier > 1.0 {
					score.Score += 25
					score.Reasons = append(score.Reasons, "天气加成")
				}
			}

			// 5. 目标HP较低时优先攻击
			targetHPPercentage := float32(defender.CurrentHP) / float32(defender.MaxHP)
			if targetHPPercentage < 0.3 {
				score.Score += 30
				score.Reasons = append(score.Reasons, "目标HP较低")
				score.Priority = 10
			}
		}

		scores = append(scores, score)
	}

	// 排序分数
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	if len(scores) > 0 {
		return scores[0]
	}
	return nil
}

// ==================== 宝可梦换出建议 ====================

// EvaluateSwitchOption 评估换出选项
func (bss *BattleStrategySystem) EvaluateSwitchOption(
	currentPokemon *model.BattlePokemon,
	switchCandidate *model.BattlePokemon,
	switchCandidateTypes []string,
	opponent *model.BattlePokemon,
	opponentTypes []string,
) *model.SwitchOption {
	option := &model.SwitchOption{
		PokemonID:    switchCandidate.PokemonID,
		Name:         switchCandidate.NameEn,
		Level:        switchCandidate.Level,
		CurrentHP:    switchCandidate.CurrentHP,
		MaxHP:        switchCandidate.MaxHP,
		HPPercentage: float32(switchCandidate.CurrentHP) / float32(switchCandidate.MaxHP),
	}

	// 转换为状态条件ID列表
	for _, cond := range switchCandidate.Conditions {
		option.Status = append(option.Status, cond)
	}

	// 分析优势/劣势
	candidateWeakness := bss.AnalyzeDefenderWeakness(switchCandidateTypes, switchCandidate.Conditions)

	// 检查是否有高伤害的招式对对手有效
	hasSuperEffective := false
	for _, move := range switchCandidate.Moves {
		if move.Power > 0 {
			matchup := bss.AnalyzeMoveEffectiveness(move, opponentTypes)
			if matchup.Multiplier >= 2.0 {
				hasSuperEffective = true
				break
			}
		}
	}

	if hasSuperEffective && len(candidateWeakness.Weak) == 0 {
		option.Advantage = "有利"
		option.AdvantageReason = "有超级有效的招式且无明显弱点"
	} else if len(candidateWeakness.Weak) > 2 && option.HPPercentage < 0.5 {
		option.Advantage = "劣势"
		option.AdvantageReason = "弱点多且HP较低"
	} else {
		option.Advantage = "平衡"
	}

	return option
}

// ==================== 状态效果分析 ====================

// AnalyzeStatusThreat 分析状态异常威胁
func (bss *BattleStrategySystem) AnalyzeStatusThreat(
	pokemon *model.BattlePokemon,
	incomingStatus string,
) *model.StatusAlert {
	alert := &model.StatusAlert{
		HasStatus: len(pokemon.Conditions) > 0,
		Status:    []*model.StatusEffectDisplay{},
	}

	// 获取状态异常效果数据
	statusData := bss.effectsSystem.GetStatusEffectData(incomingStatus)
	if statusData != nil {
		// 判断紧急程度
		switch incomingStatus {
		case "status_freeze", "status_sleep":
			alert.IsUrgent = true
			alert.Impact = "无法行动"
			alert.SuggestedAction = "立即更换宝可梦或使用恢复道具"
		case "status_paralysis":
			alert.IsUrgent = false
			alert.Impact = "速度下降，可能无法行动"
			alert.SuggestedAction = "保持警惕"
		case "status_burn", "status_poison":
			alert.IsUrgent = false
			alert.Impact = "持续伤害"
			alert.SuggestedAction = "及时治疗"
		}
	}

	return alert
}

// ==================== 战斗建议生成 ====================

// GenerateBattleStrategy 生成整体战斗策略
func (bss *BattleStrategySystem) GenerateBattleStrategy(
	attacker *model.BattlePokemon,
	attackerTypes []string,
	defender *model.BattlePokemon,
	defenderTypes []string,
	weather string,
	availablePokemon []*model.BattlePokemon,
) *BattleStrategyRecommendation {
	strategy := &BattleStrategyRecommendation{}

	// 推荐最佳招式
	defenderWeakness := bss.AnalyzeDefenderWeakness(defenderTypes, defender.Conditions)
	strategy.BestMove = bss.RecommendBestMove(attacker, defender, defenderTypes, defenderWeakness, weather)

	// 分析是否需要换宝可梦
	if len(availablePokemon) > 0 {
		attackerWeakness := bss.AnalyzeDefenderWeakness(attackerTypes, attacker.Conditions)

		// 如果当前宝可梦有很多弱点且HP较低，建议换出
		hpRatio := float32(attacker.CurrentHP) / float32(attacker.MaxHP)
		if len(attackerWeakness.Weak) > 2 && hpRatio < 0.4 {
			strategy.ShouldSwitch = true
			strategy.SwitchRecommendation = bss.EvaluateSwitchOption(
				attacker,
				availablePokemon[0],
				attackerTypes, // switch candidate types would need to be passed
				defender,
				defenderTypes,
			)
			strategy.SwitchReason = "当前宝可梦多个弱点且HP较低"
		}
	}

	// 分析状态威胁
	strategy.StatusThreats = make([]*model.StatusAlert, 0)
	for _, condition := range defender.Conditions {
		alert := bss.AnalyzeStatusThreat(attacker, condition)
		if alert.IsUrgent {
			strategy.StatusThreats = append(strategy.StatusThreats, alert)
		}
	}

	// 生成整体提示
	if strategy.BestMove != nil && strategy.BestMove.Priority >= 8 {
		strategy.OverallTip = "这一回合有强力的招式可用，应该马上使用！"
	} else if strategy.ShouldSwitch {
		strategy.OverallTip = "当前宝可梦形势不利，建议更换宝可梦。"
	} else {
		strategy.OverallTip = "稳健应对，继续注意对手的属性克制。"
	}

	return strategy
}

// ==================== 战斗策略推荐结构 ====================

// BattleStrategyRecommendation 战斗策略推荐
type BattleStrategyRecommendation struct {
	BestMove             *MoveRecommendationScore
	ShouldSwitch         bool
	SwitchRecommendation *model.SwitchOption
	SwitchReason         string
	StatusThreats        []*model.StatusAlert
	OverallTip           string
	ConfidenceLevel      float32 // 0.0-1.0
}
