package model

import "time"

// ===== 宝可梦能力值相关模型 =====

// PokemonBaseStats 宝可梦种族值（base stats）
type PokemonBaseStats struct {
	HP      int `json:"hp"`      // 生命值
	Attack  int `json:"attack"`  // 物理攻击
	Defense int `json:"defense"` // 物理防御
	SpAtk   int `json:"sp_atk"`  // 特殊攻击
	SpDef   int `json:"sp_def"`  // 特殊防御
	Speed   int `json:"speed"`   // 速度
	Total   int `json:"total"`   // 总值（通常是这6项之和）
}

// PokemonIndividualStats 宝可梦个体值（Individual Values - IV）
// 范围：0-31，每个宝可梦生成时的随机值
type PokemonIndividualStats struct {
	ID        int       `json:"id"`
	PokemonID string    `json:"pokemon_id"`
	IVHP      int       `json:"iv_hp"`      // 0-31
	IVAttack  int       `json:"iv_attack"`  // 0-31
	IVDefense int       `json:"iv_defense"` // 0-31
	IVSpAtk   int       `json:"iv_sp_atk"`  // 0-31
	IVSpDef   int       `json:"iv_sp_def"`  // 0-31
	IVSpeed   int       `json:"iv_speed"`   // 0-31
	CreatedAt time.Time `json:"created_at"`
}

// PokemonEffortStats 宝可梦努力值（Effort Values - EV）
// 用户可以通过训练来增加，每项最多252
type PokemonEffortStats struct {
	ID        int       `json:"id"`
	UserPkmID int       `json:"user_pokemon_id"`
	EVHP      int       `json:"ev_hp"`      // 0-252
	EVAttack  int       `json:"ev_attack"`  // 0-252
	EVDefense int       `json:"ev_defense"` // 0-252
	EVSpAtk   int       `json:"ev_sp_atk"`  // 0-252
	EVSpDef   int       `json:"ev_sp_def"`  // 0-252
	EVSpeed   int       `json:"ev_speed"`   // 0-252
	TotalEV   int       `json:"total_ev"`   // 0-510（所有EV的总和，最多510）
	UpdatedAt time.Time `json:"updated_at"`
}

// PokemonNature 宝可梦性格
// 性格会增加一个能力值10%，降低另一个10%
type PokemonNature struct {
	ID             int       `json:"id"`
	NatureID       string    `json:"nature_id"`
	NameEn         string    `json:"name_en"`
	NameZh         string    `json:"name_zh,omitempty"`
	IncreasedStat  string    `json:"increased_stat"`  // 增加的能力值名称
	DecreasedStat  string    `json:"decreased_stat"`  // 减少的能力值名称
	FavoriteFlavor string    `json:"favorite_flavor"` // 喜欢的口味（用于树果）
	DislikedFlavor string    `json:"disliked_flavor"` // 厌恶的口味
	CreatedAt      time.Time `json:"created_at"`
}

// CalculatedStats 计算后的实际能力值
// 计算公式：
// HP = ((2*base + iv + ev/4) * level) / 100 + level + 5
// 其他 = ((2*base + iv + ev/4) * level) / 100 + 5) * 性格倍数
type CalculatedStats struct {
	HP      int `json:"hp"`
	Attack  int `json:"attack"`
	Defense int `json:"defense"`
	SpAtk   int `json:"sp_atk"`
	SpDef   int `json:"sp_def"`
	Speed   int `json:"speed"`
}

// ===== 宝可梦特性和技能相关 =====

// PokemonAbility 宝可梦特性
type PokemonAbility struct {
	ID          int    `json:"id"`
	AbilityID   string `json:"ability_id"`
	NameEn      string `json:"name_en"`
	NameZh      string `json:"name_zh,omitempty"`
	Description string `json:"description,omitempty"`
	IsHidden    bool   `json:"is_hidden"` // 是否为隐藏特性
}

// PokemonMove 宝可梦招式
type PokemonMove struct {
	ID          string `json:"id"`
	NameEn      string `json:"name_en"`
	NameZh      string `json:"name_zh,omitempty"`
	Type        string `json:"type"`
	Category    string `json:"category"` // physical/special/status
	Power       int    `json:"power"`    // 可能为0（变化类技能）
	Accuracy    int    `json:"accuracy"` // 可能为0（必中技能）
	PP          int    `json:"pp"`       // 威力点数
	Priority    int    `json:"priority"` // 优先级，0为普通
	Description string `json:"description"`
}

// PokemonMoveSlot 宝可梦已学会的招式
type PokemonMoveSlot struct {
	ID            int       `json:"id"`
	UserPokemonID int       `json:"user_pokemon_id"`
	MoveID        string    `json:"move_id"`
	Slot          int       `json:"slot"` // 1-4
	LearnedAt     time.Time `json:"learned_at"`
}

// ===== 异常状态相关 =====

// PokemonStatusCondition 宝可梦异常状态
type PokemonStatusCondition struct {
	ID                int    `json:"id"`
	ConditionID       string `json:"condition_id"`
	NameEn            string `json:"name_en"`
	NameZh            string `json:"name_zh,omitempty"`
	EffectDescription string `json:"effect_description,omitempty"`
}

// PokemonCurrentCondition 宝可梦当前的异常状态
type PokemonCurrentCondition struct {
	ID            int       `json:"id"`
	UserPokemonID int       `json:"user_pokemon_id"`
	ConditionID   string    `json:"condition_id"`
	AppliedAt     time.Time `json:"applied_at"`
	DurationTurns int       `json:"duration_turns"` // 持续回合数
	IsActive      bool      `json:"is_active"`
}

// ===== 树果相关 =====

// Pokeberry 树果
type Pokeberry struct {
	ID           int    `json:"id"`
	BerryID      string `json:"berry_id"`
	NameEn       string `json:"name_en"`
	NameZh       string `json:"name_zh,omitempty"`
	BerryType    string `json:"berry_type"`    // 树果的类型
	StatAffected string `json:"stat_affected"` // 影响的能力值
	EffectAmount int    `json:"effect_amount"` // 效果大小
	GrowthTime   int    `json:"growth_time"`   // 生长时间
}

// ===== 宝可梦完整信息 =====

// UserPokemonFull 用户拥有的宝可梦的完整信息
type UserPokemonFull struct {
	ID         int             `json:"id"`
	UserID     int             `json:"user_id"`
	PokemonID  string          `json:"pokemon_id"`
	Nickname   string          `json:"nickname,omitempty"`
	Level      int             `json:"level"`
	Experience int             `json:"experience"`
	CurrentHP  int             `json:"current_hp"`
	MaxHP      int             `json:"max_hp"`
	Gender     string          `json:"gender"` // male/female/unknown
	IsShiny    bool            `json:"is_shiny"`
	Nature     *PokemonNature  `json:"nature,omitempty"`
	Ability    *PokemonAbility `json:"ability,omitempty"`
	HeldItem   string          `json:"held_item,omitempty"`

	// 能力值系统
	BaseStats       *PokemonBaseStats       `json:"base_stats,omitempty"`
	IndividualStats *PokemonIndividualStats `json:"individual_stats,omitempty"`
	EffortStats     *PokemonEffortStats     `json:"effort_stats,omitempty"`
	CalculatedStats *CalculatedStats        `json:"calculated_stats,omitempty"`

	// 招式和特性
	Moves      []*PokemonMoveSlot         `json:"moves,omitempty"`
	Conditions []*PokemonCurrentCondition `json:"conditions,omitempty"`

	CapturedAt time.Time `json:"captured_at,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ===== 战斗中使用的数据 =====

// BattlePokemon 战斗中的宝可梦状态
type BattlePokemon struct {
	UserPokemonID int              `json:"user_pokemon_id"`
	PokemonID     string           `json:"pokemon_id"`
	NameEn        string           `json:"name_en"`
	Level         int              `json:"level"`
	CurrentHP     int              `json:"current_hp"`
	MaxHP         int              `json:"max_hp"`
	Stats         *CalculatedStats `json:"stats"`
	Moves         []*PokemonMove   `json:"moves"`
	Ability       *PokemonAbility  `json:"ability"`
	Conditions    []string         `json:"conditions"` // 当前状态条件ID列表
}

// BattleMove 战斗中的招式选择
type BattleMove struct {
	MoveID   string `json:"move_id"`
	MoveSlot int    `json:"move_slot"`
	PPUsed   int    `json:"pp_used"`
}

// BattleAction 战斗中的一个动作
type BattleAction struct {
	ID            int       `json:"id"`
	BattleID      string    `json:"battle_id"`
	RoundNumber   int       `json:"round_number"`
	ActorPkmnID   int       `json:"actor_pokemon_id"`
	TargetPkmnID  int       `json:"target_pokemon_id,omitempty"`
	ActionType    string    `json:"action_type"` // attack/switch/item/flee
	MoveID        string    `json:"move_id,omitempty"`
	DamageDealt   int       `json:"damage_dealt,omitempty"`
	DamageTaken   int       `json:"damage_taken,omitempty"`
	Effectiveness float32   `json:"effectiveness,omitempty"` // 克制倍数 0.25/0.5/1/2/4
	CriticalHit   bool      `json:"critical_hit"`
	CreatedAt     time.Time `json:"created_at"`
}
