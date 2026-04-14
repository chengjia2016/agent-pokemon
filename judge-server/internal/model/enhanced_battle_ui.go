package model

import "time"

// ==================== 增强的战斗UI响应 ====================

// EnhancedBattleState 增强的战斗状态响应
type EnhancedBattleState struct {
	BattleID      string                    `json:"battle_id"`
	Status        string                    `json:"status"` // "ongoing", "won", "lost", "draw"
	Turn          int                       `json:"turn"`
	Weather       string                    `json:"weather"`
	Terrain       string                    `json:"terrain"`
	AttackerState *EnhancedPokemonState     `json:"attacker_state"`
	DefenderState *EnhancedPokemonState     `json:"defender_state"`
	BattleLog     []*DetailedBattleLogEntry `json:"battle_log"`
	ActionOptions *BattleActionOptions      `json:"action_options"`
	UiHints       *BattleUIHints            `json:"ui_hints"`
	Statistics    *BattleStatistics         `json:"statistics"`
	UpdatedAt     time.Time                 `json:"updated_at"`
}

// ==================== 宝可梦状态显示 ====================

// EnhancedPokemonState 增强的宝可梦状态
type EnhancedPokemonState struct {
	PokemonID      string                 `json:"pokemon_id"`
	Name           string                 `json:"name"`
	Level          int                    `json:"level"`
	CurrentHP      int                    `json:"current_hp"`
	MaxHP          int                    `json:"max_hp"`
	HPPercentage   float32                `json:"hp_percentage"`
	HPBarColor     string                 `json:"hp_bar_color"` // "green", "yellow", "red"
	Stats          *EnhancedStatsDisplay  `json:"stats"`
	Moves          []*MoveOption          `json:"moves"`
	Status         []*StatusEffectDisplay `json:"status"`
	Ability        *AbilityDisplay        `json:"ability"`
	HeldItem       *ItemDisplay           `json:"held_item,omitempty"`
	Experience     *ExperienceDisplay     `json:"experience,omitempty"`
	Typing         []string               `json:"typing"` // 属性类型
	TypeWeaknesses *TypeWeaknessData      `json:"type_weaknesses"`
	Terrain        string                 `json:"terrain,omitempty"` // 当前站立的场地
}

// EnhancedStatsDisplay 增强的属性显示
type EnhancedStatsDisplay struct {
	Attack  *StatDisplay `json:"attack"`
	Defense *StatDisplay `json:"defense"`
	SpAtk   *StatDisplay `json:"sp_atk"`
	SpDef   *StatDisplay `json:"sp_def"`
	Speed   *StatDisplay `json:"speed"`
	HP      *StatDisplay `json:"hp"`
}

// StatDisplay 单个属性显示
type StatDisplay struct {
	Value          int     `json:"value"`
	BaseValue      int     `json:"base_value"`
	Modifier       float32 `json:"modifier"`        // 当前修正倍率
	ModifierReason string  `json:"modifier_reason"` // "ability", "item", "status", "terrain"
	IsModified     bool    `json:"is_modified"`
	BarColor       string  `json:"bar_color"` // 属性条颜色
	Icon           string  `json:"icon"`      // 属性图标
}

// MoveOption 招式选项
type MoveOption struct {
	MoveID         string                `json:"move_id"`
	MoveName       string                `json:"move_name"`
	Type           string                `json:"type"`
	Category       string                `json:"category"` // "physical", "special", "status"
	Power          int                   `json:"power"`
	Accuracy       int                   `json:"accuracy"`
	PP             int                   `json:"pp"`
	PPMax          int                   `json:"pp_max"`
	Priority       int                   `json:"priority"`
	Animation      *MoveAnimationDisplay `json:"animation,omitempty"`
	Effect         string                `json:"effect"` // 招式效果描述
	Description    string                `json:"description"`
	TypeMatchup    *TypeMatchupInfo      `json:"type_matchup"` // 对手的类型相克信息
	IsDisabled     bool                  `json:"is_disabled"`
	DisableReason  string                `json:"disable_reason,omitempty"`
	Recommendation *MoveRecommendation   `json:"recommendation,omitempty"`
}

// MoveAnimationDisplay 招式动画显示数据
type MoveAnimationDisplay struct {
	Type            string   `json:"type"` // "projectile", "melee", "beam", "status", "particle"
	Duration        int      `json:"duration"`
	Color           string   `json:"color"`
	SoundEffect     string   `json:"sound_effect"`
	ParticleEffects []string `json:"particle_effects"`
	ScreenEffect    string   `json:"screen_effect"`
}

// TypeMatchupInfo 类型相克信息
type TypeMatchupInfo struct {
	Effective       []string `json:"effective"`        // 有效的类型
	NotEffective    []string `json:"not_effective"`    // 无效的类型
	Resistant       []string `json:"resistant"`        // 抗性
	Weak            []string `json:"weak"`             // 弱点
	Immune          []string `json:"immune"`           // 免疫
	Multiplier      float32  `json:"multiplier"`       // 对目标的伤害倍数
	VisualIndicator string   `json:"visual_indicator"` // "🔴" "🟠" "⚪" "🔵" "🟦" "❌"
}

// MoveRecommendation 招式推荐
type MoveRecommendation struct {
	Reason   string `json:"reason"`   // "super_effective", "low_hp", "status_prevention"
	Priority int    `json:"priority"` // 推荐优先级 1-10
	Icon     string `json:"icon"`     // "⭐", "🔥", "💡"
}

// StatusEffectDisplay 状态异常显示
type StatusEffectDisplay struct {
	ConditionID   string             `json:"condition_id"`
	ConditionName string             `json:"condition_name"`
	Icon          string             `json:"icon"`
	BadgeColor    string             `json:"badge_color"`
	OverlayColor  string             `json:"overlay_color"`
	Description   string             `json:"description"`
	Duration      int                `json:"duration,omitempty"` // 剩余回合数
	DamagePerTurn int                `json:"damage_per_turn,omitempty"`
	StatModifiers map[string]float32 `json:"stat_modifiers"`
	Animation     string             `json:"animation"`
	IsCurable     bool               `json:"is_curable"`
}

// AbilityDisplay 特性显示
type AbilityDisplay struct {
	AbilityID     string `json:"ability_id"`
	AbilityName   string `json:"ability_name"`
	Description   string `json:"description"`
	Icon          string `json:"icon"`
	IsActive      bool   `json:"is_active"`
	CurrentEffect string `json:"current_effect,omitempty"`
}

// ItemDisplay 物品显示
type ItemDisplay struct {
	ItemID      string `json:"item_id"`
	ItemName    string `json:"item_name"`
	Effect      string `json:"effect"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

// ExperienceDisplay 经验显示
type ExperienceDisplay struct {
	Current          int     `json:"current"`
	NextLevel        int     `json:"next_level"`
	Percentage       float32 `json:"percentage"`
	ToNextLevel      int     `json:"to_next_level"`
	GainedThisBattle int     `json:"gained_this_battle,omitempty"`
}

// TypeWeaknessData 属性弱点数据
type TypeWeaknessData struct {
	Weak      map[string]float32 `json:"weak"`      // 弱点及倍数
	Resistant map[string]float32 `json:"resistant"` // 抗性及倍数
	Immune    []string           `json:"immune"`    // 完全免疫
}

// ==================== 战斗日志 ====================

// DetailedBattleLogEntry 详细的战斗日志条目
type DetailedBattleLogEntry struct {
	Round        int    `json:"round"`
	Turn         int    `json:"turn"`
	ActorName    string `json:"actor_name"`
	ActionType   string `json:"action_type"` // "move", "switch", "item", "status_damage"
	ActionDetail string `json:"action_detail"`
	TargetName   string `json:"target_name,omitempty"`

	// 伤害相关
	Damage         int     `json:"damage,omitempty"`
	DamageReason   string  `json:"damage_reason,omitempty"` // "attack", "status_burn", "status_poison"
	EffectiveIndex float32 `json:"effective_index,omitempty"`
	IsCritical     bool    `json:"is_critical,omitempty"`
	CriticalNote   string  `json:"critical_note,omitempty"`

	// 效果相关
	StatusApplied string             `json:"status_applied,omitempty"`
	StatusRemoved string             `json:"status_removed,omitempty"`
	StatChanges   map[string]float32 `json:"stat_changes,omitempty"`

	// UI显示
	Icon      string              `json:"icon"`
	MessageZH string              `json:"message_zh"`
	MessageEN string              `json:"message_en"`
	Color     string              `json:"color"`
	Timestamp time.Time           `json:"timestamp"`
	Animation *LogAnimationEffect `json:"animation,omitempty"`
}

// LogAnimationEffect 日志动画效果
type LogAnimationEffect struct {
	Type      string  `json:"type"` // "damage_float", "heal_float", "text_pop", "icon_animate"
	Duration  int     `json:"duration"`
	Magnitude float32 `json:"magnitude"`
}

// ==================== 战斗操作选项 ====================

// BattleActionOptions 战斗操作选项
type BattleActionOptions struct {
	CanAttack         bool                  `json:"can_attack"`
	CanSwitch         bool                  `json:"can_switch"`
	CanUseItem        bool                  `json:"can_use_item"`
	CanFlee           bool                  `json:"can_flee"`
	AvailableMoves    []*MoveOption         `json:"available_moves"`
	AvailablePokemon  []*SwitchOption       `json:"available_pokemon,omitempty"`
	AvailableItems    []*ItemOption         `json:"available_items,omitempty"`
	RecommendedAction *ActionRecommendation `json:"recommended_action,omitempty"`
	TimeLimit         int                   `json:"time_limit"` // 秒数
}

// SwitchOption 换宝可梦选项
type SwitchOption struct {
	PokemonID       string   `json:"pokemon_id"`
	Name            string   `json:"name"`
	Level           int      `json:"level"`
	CurrentHP       int      `json:"current_hp"`
	MaxHP           int      `json:"max_hp"`
	HPPercentage    float32  `json:"hp_percentage"`
	Status          []string `json:"status"`
	Advantage       string   `json:"advantage,omitempty"` // "有利", "劣势", "平衡"
	AdvantageReason string   `json:"advantage_reason,omitempty"`
}

// ItemOption 物品选项
type ItemOption struct {
	ItemID    string `json:"item_id"`
	ItemName  string `json:"item_name"`
	Effect    string `json:"effect"`
	Quantity  int    `json:"quantity"`
	CanUse    bool   `json:"can_use"`
	UseReason string `json:"use_reason,omitempty"`
}

// ActionRecommendation 行动推荐
type ActionRecommendation struct {
	RecommendedAction string  `json:"recommended_action"` // "attack", "switch", "item", "flee"
	Reason            string  `json:"reason"`
	Confidence        float32 `json:"confidence"` // 0.0-1.0
	Priority          int     `json:"priority"`
	Icon              string  `json:"icon"`
	Description       string  `json:"description"`
}

// ==================== 战斗UI提示 ====================

// BattleUIHints 战斗UI提示
type BattleUIHints struct {
	TypingAdvantage  *TypingHint    `json:"typing_advantage"`
	WeaknessAlert    *WeaknessAlert `json:"weakness_alert,omitempty"`
	StatusAlert      *StatusAlert   `json:"status_alert,omitempty"`
	StrategyTip      string         `json:"strategy_tip,omitempty"`
	DamagePreview    *DamagePreview `json:"damage_preview,omitempty"`
	NextRoundWarning string         `json:"next_round_warning,omitempty"`
	TurnInfo         string         `json:"turn_info"`
}

// TypingHint 属性提示
type TypingHint struct {
	YourType     string  `json:"your_type"`
	OpponentType string  `json:"opponent_type"`
	Advantage    string  `json:"advantage"` // "你有优势", "对手有优势", "平衡"
	Multiplier   float32 `json:"multiplier"`
	Icon         string  `json:"icon"`
	Description  string  `json:"description"`
}

// WeaknessAlert 弱点警报
type WeaknessAlert struct {
	IsAlert         bool          `json:"is_alert"`
	YourWeakness    []string      `json:"your_weakness"`
	DangerLevel     string        `json:"danger_level"` // "low", "medium", "high", "critical"
	Recommendation  string        `json:"recommendation"`
	SuggestedSwitch *SwitchOption `json:"suggested_switch,omitempty"`
}

// StatusAlert 状态警报
type StatusAlert struct {
	HasStatus       bool                   `json:"has_status"`
	Status          []*StatusEffectDisplay `json:"status"`
	Impact          string                 `json:"impact"` // "伤害", "无法行动", "属性下降"
	IsUrgent        bool                   `json:"is_urgent"`
	SuggestedAction string                 `json:"suggested_action"`
}

// DamagePreview 伤害预览
type DamagePreview struct {
	EstimatedMinDamage int        `json:"estimated_min_damage"`
	EstimatedMaxDamage int        `json:"estimated_max_damage"`
	TargetHPAfter      *HPPreview `json:"target_hp_after"`
	IsFatal            bool       `json:"is_fatal"`
	FatalRound         int        `json:"fatal_round,omitempty"`
}

// HPPreview HP预览
type HPPreview struct {
	Min           int     `json:"min"`
	Max           int     `json:"max"`
	MinPercentage float32 `json:"min_percentage"`
	MaxPercentage float32 `json:"max_percentage"`
}

// ==================== 战斗统计 ====================

// BattleStatistics 战斗统计
type BattleStatistics struct {
	TotalTurns           int            `json:"total_turns"`
	AttackerDamageDealt  int            `json:"attacker_damage_dealt"`
	DefenderDamageDealt  int            `json:"defender_damage_dealt"`
	AttackerMovesUsed    map[string]int `json:"attacker_moves_used"`
	DefenderMovesUsed    map[string]int `json:"defender_moves_used"`
	CriticalHitsAttacker int            `json:"critical_hits_attacker"`
	CriticalHitsDefender int            `json:"critical_hits_defender"`
	StatusApplied        map[string]int `json:"status_applied"`
	AverageEffectiveness float32        `json:"average_effectiveness"`
	ExpGained            int            `json:"exp_gained,omitempty"`
	CoinsGained          int            `json:"coins_gained,omitempty"`
	ItemsObtained        []string       `json:"items_obtained,omitempty"`
}
