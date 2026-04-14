package service

import (
	"judge-server/internal/model"
	"math"
	"math/rand"
)

// ==================== 战斗效果系统 ====================
// BattleEffectsSystem 提供增强的战斗效果和动画数据
type BattleEffectsSystem struct {
	moveAnimations     map[string]*MoveAnimation
	statusEffects      map[string]*StatusEffectData
	typeMatchupVisuals map[string]map[string]string
}

// NewBattleEffectsSystem 创建新的战斗效果系统
func NewBattleEffectsSystem() *BattleEffectsSystem {
	return &BattleEffectsSystem{
		moveAnimations:     initializeMoveAnimations(),
		statusEffects:      initializeStatusEffects(),
		typeMatchupVisuals: initializeTypeMatchupVisuals(),
	}
}

// ==================== 招式动画 ====================

// MoveAnimation 招式动画数据
type MoveAnimation struct {
	MoveID         string          `json:"move_id"`
	MoveName       string          `json:"move_name"`
	AnimationType  string          `json:"animation_type"` // "projectile", "melee", "beam", "status", "particle"
	Duration       int             `json:"duration"`       // 毫秒
	Color          string          `json:"color"`          // 十六进制颜色
	ParticleEffect *ParticleEffect `json:"particle_effect"`
	SoundEffect    string          `json:"sound_effect"`
	ScreenEffect   *ScreenEffect   `json:"screen_effect"`
	Description    string          `json:"description"`
	CriticalEffect *CriticalEffect `json:"critical_effect"`
}

// ParticleEffect 粒子效果
type ParticleEffect struct {
	Type       string  `json:"type"`     // "spark", "explosion", "wave", "aura", "slash"
	Count      int     `json:"count"`    // 粒子数量
	Speed      float32 `json:"speed"`    // 粒子速度
	Lifetime   int     `json:"lifetime"` // 粒子生命周期(毫秒)
	Spread     float32 `json:"spread"`   // 粒子散布角度
	ColorStart string  `json:"color_start"`
	ColorEnd   string  `json:"color_end"`
}

// ScreenEffect 屏幕效果
type ScreenEffect struct {
	Effect    string  `json:"effect"`    // "shake", "flash", "darken", "bloom"
	Intensity float32 `json:"intensity"` // 0.0-1.0
	Duration  int     `json:"duration"`  // 毫秒
}

// CriticalEffect 会心一击效果
type CriticalEffect struct {
	Icon       string  `json:"icon"`        // "⚡", "✨", "🔥"
	FlashColor string  `json:"flash_color"` // 闪光颜色
	Scale      float32 `json:"scale"`       // 缩放倍数
	Duration   int     `json:"duration"`    // 毫秒
}

// ==================== 状态异常效果 ====================

// StatusEffectData 状态异常效果数据
type StatusEffectData struct {
	ConditionID   string             `json:"condition_id"`
	ConditionName string             `json:"condition_name"`
	Description   string             `json:"description"`
	Icon          string             `json:"icon"`
	BadgeColor    string             `json:"badge_color"`
	OverlayColor  string             `json:"overlay_color"` // HP条叠加色
	Animation     string             `json:"animation"`
	SoundEffect   string             `json:"sound_effect"`
	StatModifiers map[string]float32 `json:"stat_modifiers"`
	MaxDuration   int                `json:"max_duration"`
	CanStack      bool               `json:"can_stack"`
	Priority      int                `json:"priority"` // 优先级（冻结优先级最高）
}

// ==================== 类型相克提示 ====================

// TypeMatchupVisual 类型相克视觉提示
type TypeMatchupVisual struct {
	Advantage      string  `json:"advantage"`       // "强势", "有效", 等
	Multiplier     float32 `json:"multiplier"`      // 伤害倍数
	Icon           string  `json:"icon"`            // 图标
	TextColor      string  `json:"text_color"`      // 文本颜色
	HighlightColor string  `json:"highlight_color"` // 高亮颜色
	Description    string  `json:"description"`
}

// ==================== 初始化函数 ====================

func initializeMoveAnimations() map[string]*MoveAnimation {
	return map[string]*MoveAnimation{
		// 物理招式
		"tackle": {
			MoveID:        "tackle",
			MoveName:      "撞击",
			AnimationType: "melee",
			Duration:      500,
			Color:         "#FFFFFF",
			ParticleEffect: &ParticleEffect{
				Type:       "spark",
				Count:      8,
				Speed:      2.0,
				Lifetime:   400,
				Spread:     45,
				ColorStart: "#CCCCCC",
				ColorEnd:   "#666666",
			},
			SoundEffect: "tackle.mp3",
			ScreenEffect: &ScreenEffect{
				Effect:    "shake",
				Intensity: 0.3,
				Duration:  200,
			},
			Description: "用身体猛烈撞向对方",
		},
		"slash": {
			MoveID:        "slash",
			MoveName:      "劈开",
			AnimationType: "melee",
			Duration:      600,
			Color:         "#FF6B6B",
			ParticleEffect: &ParticleEffect{
				Type:       "slash",
				Count:      3,
				Speed:      3.0,
				Lifetime:   500,
				Spread:     30,
				ColorStart: "#FF6B6B",
				ColorEnd:   "#FF0000",
			},
			SoundEffect: "slash.mp3",
			ScreenEffect: &ScreenEffect{
				Effect:    "flash",
				Intensity: 0.4,
				Duration:  150,
			},
			CriticalEffect: &CriticalEffect{
				Icon:       "⚡",
				FlashColor: "#FFD700",
				Scale:      1.3,
				Duration:   300,
			},
			Description: "用爪子或刀刃猛烈斩击",
		},
		"thunderbolt": {
			MoveID:        "thunderbolt",
			MoveName:      "十万伏特",
			AnimationType: "beam",
			Duration:      800,
			Color:         "#FFD700",
			ParticleEffect: &ParticleEffect{
				Type:       "aura",
				Count:      20,
				Speed:      4.0,
				Lifetime:   600,
				Spread:     60,
				ColorStart: "#FFD700",
				ColorEnd:   "#FFA500",
			},
			SoundEffect: "thunderbolt.mp3",
			ScreenEffect: &ScreenEffect{
				Effect:    "bloom",
				Intensity: 0.8,
				Duration:  400,
			},
			Description: "释放强大的电击",
		},
		"flame_burst": {
			MoveID:        "flame_burst",
			MoveName:      "火焰爆发",
			AnimationType: "particle",
			Duration:      700,
			Color:         "#FF4500",
			ParticleEffect: &ParticleEffect{
				Type:       "explosion",
				Count:      30,
				Speed:      2.5,
				Lifetime:   700,
				Spread:     90,
				ColorStart: "#FF4500",
				ColorEnd:   "#FFA500",
			},
			SoundEffect: "explosion.mp3",
			ScreenEffect: &ScreenEffect{
				Effect:    "shake",
				Intensity: 0.5,
				Duration:  300,
			},
			Description: "释放猛烈的火焰",
		},
		"water_pulse": {
			MoveID:        "water_pulse",
			MoveName:      "水之波动",
			AnimationType: "wave",
			Duration:      600,
			Color:         "#4A90E2",
			ParticleEffect: &ParticleEffect{
				Type:       "wave",
				Count:      15,
				Speed:      1.5,
				Lifetime:   500,
				Spread:     60,
				ColorStart: "#4A90E2",
				ColorEnd:   "#2E5C8A",
			},
			SoundEffect: "water.mp3",
			ScreenEffect: &ScreenEffect{
				Effect:    "darken",
				Intensity: 0.2,
				Duration:  250,
			},
			Description: "发动水之能量的脉冲",
		},
		"solar_beam": {
			MoveID:        "solar_beam",
			MoveName:      "日光束",
			AnimationType: "beam",
			Duration:      1200,
			Color:         "#FFD700",
			ParticleEffect: &ParticleEffect{
				Type:       "aura",
				Count:      25,
				Speed:      3.0,
				Lifetime:   800,
				Spread:     45,
				ColorStart: "#FFD700",
				ColorEnd:   "#FF8C00",
			},
			SoundEffect: "solar_beam.mp3",
			ScreenEffect: &ScreenEffect{
				Effect:    "bloom",
				Intensity: 1.0,
				Duration:  600,
			},
			Description: "收集日光发射强大射线",
		},
	}
}

func initializeStatusEffects() map[string]*StatusEffectData {
	return map[string]*StatusEffectData{
		"status_paralysis": {
			ConditionID:   "status_paralysis",
			ConditionName: "麻痹",
			Description:   "速度下降，可能无法行动",
			Icon:          "⚡",
			BadgeColor:    "#FFD700",
			OverlayColor:  "#FFFF99",
			Animation:     "spark_loop",
			SoundEffect:   "paralysis.mp3",
			StatModifiers: map[string]float32{
				"speed": 0.75,
			},
			MaxDuration: 999, // 永久直到被治愈
			CanStack:    false,
			Priority:    5,
		},
		"status_sleep": {
			ConditionID:   "status_sleep",
			ConditionName: "睡眠",
			Description:   "无法行动直到苏醒",
			Icon:          "😴",
			BadgeColor:    "#7B68EE",
			OverlayColor:  "#9DB4FF",
			Animation:     "z_loop",
			SoundEffect:   "sleep.mp3",
			StatModifiers: map[string]float32{},
			MaxDuration:   999,
			CanStack:      false,
			Priority:      8,
		},
		"status_burn": {
			ConditionID:   "status_burn",
			ConditionName: "灼伤",
			Description:   "每回合损失HP，物攻下降",
			Icon:          "🔥",
			BadgeColor:    "#FF4500",
			OverlayColor:  "#FFB6C1",
			Animation:     "flame_loop",
			SoundEffect:   "burn.mp3",
			StatModifiers: map[string]float32{
				"attack": 0.5,
			},
			MaxDuration: 999,
			CanStack:    false,
			Priority:    3,
		},
		"status_poison": {
			ConditionID:   "status_poison",
			ConditionName: "中毒",
			Description:   "每回合损失HP",
			Icon:          "☠️",
			BadgeColor:    "#9932CC",
			OverlayColor:  "#DDA0DD",
			Animation:     "poison_loop",
			SoundEffect:   "poison.mp3",
			StatModifiers: map[string]float32{},
			MaxDuration:   999,
			CanStack:      false,
			Priority:      2,
		},
		"status_freeze": {
			ConditionID:   "status_freeze",
			ConditionName: "冻结",
			Description:   "无法行动",
			Icon:          "❄️",
			BadgeColor:    "#87CEEB",
			OverlayColor:  "#B0E0E6",
			Animation:     "ice_loop",
			SoundEffect:   "freeze.mp3",
			StatModifiers: map[string]float32{},
			MaxDuration:   999,
			CanStack:      false,
			Priority:      9, // 最高优先级
		},
		"status_confusion": {
			ConditionID:   "status_confusion",
			ConditionName: "混乱",
			Description:   "可能无法行动或自伤",
			Icon:          "😵",
			BadgeColor:    "#FFB6C1",
			OverlayColor:  "#FFE4E1",
			Animation:     "spiral_loop",
			SoundEffect:   "confusion.mp3",
			StatModifiers: map[string]float32{},
			MaxDuration:   4,
			CanStack:      false,
			Priority:      4,
		},
	}
}

func initializeTypeMatchupVisuals() map[string]map[string]string {
	return map[string]map[string]string{
		"4x": {
			"text_color":      "#FF0000",
			"highlight_color": "#FF6B6B",
			"icon":            "🔴",
			"description":     "超级有效",
		},
		"2x": {
			"text_color":      "#FFA500",
			"highlight_color": "#FFD700",
			"icon":            "🟠",
			"description":     "有效",
		},
		"1x": {
			"text_color":      "#FFFFFF",
			"highlight_color": "#CCCCCC",
			"icon":            "⚪",
			"description":     "普通",
		},
		"0.5x": {
			"text_color":      "#87CEEB",
			"highlight_color": "#ADD8E6",
			"icon":            "🔵",
			"description":     "不太有效",
		},
		"0.25x": {
			"text_color":      "#4169E1",
			"highlight_color": "#1E90FF",
			"icon":            "🟦",
			"description":     "效果很差",
		},
		"0x": {
			"text_color":      "#000000",
			"highlight_color": "#808080",
			"icon":            "❌",
			"description":     "无效",
		},
	}
}

// ==================== 获取效果数据的方法 ====================

// GetMoveAnimation 获取招式动画数据
func (bes *BattleEffectsSystem) GetMoveAnimation(moveID string) *MoveAnimation {
	if anim, exists := bes.moveAnimations[moveID]; exists {
		return anim
	}
	// 返回默认动画
	return &MoveAnimation{
		MoveID:        moveID,
		MoveName:      "未知招式",
		AnimationType: "particle",
		Duration:      500,
		Color:         "#FFFFFF",
		Description:   "使用招式",
	}
}

// GetStatusEffectData 获取状态异常效果数据
func (bes *BattleEffectsSystem) GetStatusEffectData(conditionID string) *StatusEffectData {
	if effect, exists := bes.statusEffects[conditionID]; exists {
		return effect
	}
	return nil
}

// GetTypeMatchupVisual 获取类型相克视觉提示
func (bes *BattleEffectsSystem) GetTypeMatchupVisual(multiplier float32) map[string]interface{} {
	var key string
	switch {
	case multiplier >= 4.0:
		key = "4x"
	case multiplier >= 2.0:
		key = "2x"
	case multiplier > 0.5 && multiplier < 1.0:
		key = "0.5x"
	case multiplier > 0.0 && multiplier < 0.5:
		key = "0.25x"
	case multiplier == 0.0:
		key = "0x"
	default:
		key = "1x"
	}

	visual := bes.typeMatchupVisuals[key]
	return map[string]interface{}{
		"multiplier":      multiplier,
		"text_color":      visual["text_color"],
		"highlight_color": visual["highlight_color"],
		"icon":            visual["icon"],
		"description":     visual["description"],
	}
}

// ==================== 增强型战斗机制 ====================

// EnhancedBattleEngine 增强的战斗引擎
type EnhancedBattleEngine struct {
	effectsSystem    *BattleEffectsSystem
	statGrowthSystem *StatGrowthSystem
	abilitySystem    *AbilitySystem
	itemBonusSystem  *ItemBonusSystem
}

// NewEnhancedBattleEngine 创建增强战斗引擎
func NewEnhancedBattleEngine() *EnhancedBattleEngine {
	return &EnhancedBattleEngine{
		effectsSystem:    NewBattleEffectsSystem(),
		statGrowthSystem: NewStatGrowthSystem(),
		abilitySystem:    NewAbilitySystem(),
		itemBonusSystem:  NewItemBonusSystem(),
	}
}

// ==================== 属性成长系统 ====================

// StatGrowthSystem 属性成长系统
type StatGrowthSystem struct {
	baseGrowthRates map[string]float32     // 基础成长率
	levelUpBonuses  map[int]map[string]int // 等级奖励
}

// NewStatGrowthSystem 创建属性成长系统
func NewStatGrowthSystem() *StatGrowthSystem {
	return &StatGrowthSystem{
		baseGrowthRates: map[string]float32{
			"hp":      1.1,
			"attack":  1.05,
			"defense": 1.05,
			"sp_atk":  1.08,
			"sp_def":  1.08,
			"speed":   1.07,
		},
		levelUpBonuses: initializeLevelUpBonuses(),
	}
}

// CalculateStatAtLevel 计算某等级下的属性值
func (sgs *StatGrowthSystem) CalculateStatAtLevel(baseStat int, currentLevel int, statName string) int {
	if currentLevel < 1 {
		currentLevel = 1
	}

	growthRate := sgs.baseGrowthRates[statName]
	if growthRate == 0 {
		growthRate = 1.05 // 默认成长率
	}

	// 基础计算: 基础值 * (成长率 ^ (等级-1))
	stat := float32(baseStat) * float32(math.Pow(float64(growthRate), float64(currentLevel-1)))

	// 添加等级奖励
	if bonus, exists := sgs.levelUpBonuses[currentLevel][statName]; exists {
		stat += float32(bonus)
	}

	return int(stat)
}

func initializeLevelUpBonuses() map[int]map[string]int {
	bonuses := make(map[int]map[string]int)
	// 每10级获得额外奖励
	for level := 10; level <= 100; level += 10 {
		bonuses[level] = map[string]int{
			"hp":      level / 2,
			"attack":  level / 3,
			"defense": level / 3,
			"sp_atk":  level / 4,
			"sp_def":  level / 4,
			"speed":   level / 5,
		}
	}
	return bonuses
}

// ==================== 特性系统 ====================

// AbilitySystem 特性系统
type AbilitySystem struct {
	abilities map[string]*AbilityData
}

// AbilityData 特性数据
type AbilityData struct {
	AbilityID   string  `json:"ability_id"`
	AbilityName string  `json:"ability_name"`
	Description string  `json:"description"`
	Effect      string  `json:"effect"`    // "stat_boost", "damage_reduction", "healing", "status_immunity"
	Magnitude   float32 `json:"magnitude"` // 效果幅度
	Trigger     string  `json:"trigger"`   // "passive", "on_hit", "on_damage"
	Icon        string  `json:"icon"`
	Probability float32 `json:"probability"` // 触发概率
}

// NewAbilitySystem 创建特性系统
func NewAbilitySystem() *AbilitySystem {
	return &AbilitySystem{
		abilities: initializeAbilities(),
	}
}

func initializeAbilities() map[string]*AbilityData {
	return map[string]*AbilityData{
		"static": {
			AbilityID:   "static",
			AbilityName: "静电",
			Description: "有概率使接触者陷入麻痹状态",
			Effect:      "status_application",
			Magnitude:   0.3,
			Trigger:     "on_hit",
			Icon:        "⚡",
			Probability: 0.3,
		},
		"speed_boost": {
			AbilityID:   "speed_boost",
			AbilityName: "加速",
			Description: "每回合速度提升",
			Effect:      "stat_boost",
			Magnitude:   1.1,
			Trigger:     "passive",
			Icon:        "🚀",
			Probability: 1.0,
		},
		"filter": {
			AbilityID:   "filter",
			AbilityName: "过滤",
			Description: "减轻属性的弱点",
			Effect:      "damage_reduction",
			Magnitude:   0.5,
			Trigger:     "passive",
			Icon:        "🛡️",
			Probability: 1.0,
		},
		"natural_cure": {
			AbilityID:   "natural_cure",
			AbilityName: "自然恢复",
			Description: "更换精灵时治愈异常状态",
			Effect:      "healing",
			Magnitude:   1.0,
			Trigger:     "on_switch",
			Icon:        "💚",
			Probability: 1.0,
		},
	}
}

// GetAbility 获取特性数据
func (as *AbilitySystem) GetAbility(abilityID string) *AbilityData {
	if ability, exists := as.abilities[abilityID]; exists {
		return ability
	}
	return nil
}

// ApplyAbilityEffect 应用特性效果
func (as *AbilitySystem) ApplyAbilityEffect(ability *AbilityData, pokemon *model.BattlePokemon) {
	if ability == nil || pokemon == nil {
		return
	}

	switch ability.Effect {
	case "stat_boost":
		pokemon.Stats.Speed = int(float32(pokemon.Stats.Speed) * ability.Magnitude)
	case "damage_reduction":
		// 在伤害计算时应用
	case "healing":
		pokemon.CurrentHP = int(float32(pokemon.CurrentHP) * ability.Magnitude)
		if pokemon.CurrentHP > pokemon.MaxHP {
			pokemon.CurrentHP = pokemon.MaxHP
		}
	}
}

// ==================== 物品加成系统 ====================

// ItemBonusSystem 物品加成系统
type ItemBonusSystem struct {
	items map[string]*ItemBonusData
}

// ItemBonusData 物品加成数据
type ItemBonusData struct {
	ItemID      string  `json:"item_id"`
	ItemName    string  `json:"item_name"`
	BonusType   string  `json:"bonus_type"`  // "stat_boost", "type_boost", "healing"
	TargetStat  string  `json:"target_stat"` // 目标属性
	Magnitude   float32 `json:"magnitude"`   // 加成倍数
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
}

// NewItemBonusSystem 创建物品加成系统
func NewItemBonusSystem() *ItemBonusSystem {
	return &ItemBonusSystem{
		items: initializeItemBonuses(),
	}
}

func initializeItemBonuses() map[string]*ItemBonusData {
	return map[string]*ItemBonusData{
		"choice_band": {
			ItemID:      "choice_band",
			ItemName:    "选择围巾",
			BonusType:   "stat_boost",
			TargetStat:  "attack",
			Magnitude:   1.5,
			Description: "物攻提升50%，但只能使用同一招式",
			Icon:        "🧣",
		},
		"life_orb": {
			ItemID:      "life_orb",
			ItemName:    "生命宝珠",
			BonusType:   "stat_boost",
			TargetStat:  "sp_atk",
			Magnitude:   1.3,
			Description: "特攻提升30%，但每回合损失1/10最大HP",
			Icon:        "🔴",
		},
		"assault_vest": {
			ItemID:      "assault_vest",
			ItemName:    "突击背心",
			BonusType:   "stat_boost",
			TargetStat:  "sp_def",
			Magnitude:   1.5,
			Description: "特防提升50%",
			Icon:        "🦺",
		},
	}
}

// GetItemBonus 获取物品加成数据
func (ibs *ItemBonusSystem) GetItemBonus(itemID string) *ItemBonusData {
	if item, exists := ibs.items[itemID]; exists {
		return item
	}
	return nil
}

// ApplyItemBonus 应用物品加成
func (ibs *ItemBonusSystem) ApplyItemBonus(itemBonus *ItemBonusData, pokemon *model.BattlePokemon) {
	if itemBonus == nil || pokemon == nil {
		return
	}

	switch itemBonus.TargetStat {
	case "attack":
		pokemon.Stats.Attack = int(float32(pokemon.Stats.Attack) * itemBonus.Magnitude)
	case "sp_atk":
		pokemon.Stats.SpAtk = int(float32(pokemon.Stats.SpAtk) * itemBonus.Magnitude)
	case "defense":
		pokemon.Stats.Defense = int(float32(pokemon.Stats.Defense) * itemBonus.Magnitude)
	case "sp_def":
		pokemon.Stats.SpDef = int(float32(pokemon.Stats.SpDef) * itemBonus.Magnitude)
	case "speed":
		pokemon.Stats.Speed = int(float32(pokemon.Stats.Speed) * itemBonus.Magnitude)
	}
}

// ==================== 增强伤害计算 ====================

// CalculateEnhancedDamage 计算增强的伤害（包含所有加成）
func (ebe *EnhancedBattleEngine) CalculateEnhancedDamage(
	attacker *model.BattlePokemon,
	defender *model.BattlePokemon,
	move *model.PokemonMove,
	weather string,
	isCritical bool,
	terrain string, // 场地类型
) (int, float32, map[string]interface{}) {
	details := make(map[string]interface{})

	if move.Power == 0 {
		return 0, 1.0, details
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
		return 0, 1.0, details
	}

	// 应用特性加成
	if attacker.Ability != nil {
		abilityData := ebe.abilitySystem.GetAbility(attacker.Ability.AbilityID)
		if abilityData != nil {
			ebe.abilitySystem.ApplyAbilityEffect(abilityData, attacker)
		}
	}

	// 计算基础伤害
	baseDamage := ((((2*attackStat/5 + 2) * move.Power * defenseStat / 50) / 50) + 2)

	// 获取类型克制倍数
	effectiveness := float32(1.0) // 需要从类型克制系统获取

	// 应用会心一击加成
	criticalMultiplier := float32(1.0)
	if isCritical {
		criticalMultiplier = 1.5
		details["critical"] = true
	}

	// 应用天气加成
	weatherMultiplier := getWeatherMultiplier(move.Type, weather)
	if weatherMultiplier != 1.0 {
		details["weather_boost"] = weatherMultiplier
	}

	// 应用场地加成
	terrainMultiplier := getTerrainMultiplier(move.Type, terrain)
	if terrainMultiplier != 1.0 {
		details["terrain_boost"] = terrainMultiplier
	}

	// 应用随机浮动（85-100%）
	randomMultiplier := float32(85+rand.Intn(16)) / 100.0

	// 计算最终伤害
	damage := int(float32(baseDamage) * effectiveness * criticalMultiplier * weatherMultiplier * terrainMultiplier * randomMultiplier)

	// 确保伤害至少为1
	if damage < 1 {
		damage = 1
	}

	details["base_damage"] = baseDamage
	details["effectiveness"] = effectiveness
	details["final_damage"] = damage
	details["multiplier"] = criticalMultiplier * weatherMultiplier * terrainMultiplier

	return damage, effectiveness, details
}

func getWeatherMultiplier(moveType, weather string) float32 {
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

func getTerrainMultiplier(moveType, terrain string) float32 {
	terrainBoosts := map[string]map[string]float32{
		"grassy_terrain": {
			"grass":  1.5,
			"ground": 0.5,
		},
		"electric_terrain": {
			"electric": 1.5,
		},
		"psychic_terrain": {
			"psychic": 1.5,
		},
		"misty_terrain": {
			"dragon": 0.5,
		},
	}

	if boost, exists := terrainBoosts[terrain][moveType]; exists {
		return boost
	}
	return 1.0
}
