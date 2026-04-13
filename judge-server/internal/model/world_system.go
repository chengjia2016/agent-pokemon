package model

import "time"

// ==================== NPC 系统 ====================

// NPC 代表世界中的一个 NPC
type NPC struct {
	ID        int       `json:"id"`
	TownID    string    `json:"town_id"`
	NameEn    string    `json:"name_en"`
	NameZh    string    `json:"name_zh"`
	Type      string    `json:"type"` // trainer, merchant, quest_giver, gym_leader, professor
	Role      string    `json:"role"` // 描述性角色
	CoordX    float64   `json:"coord_x"`
	CoordY    float64   `json:"coord_y"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NPCDialogue NPC对话
type NPCDialogue struct {
	ID           int    `json:"id"`
	NPCId        int    `json:"npc_id"`
	DialogueText string `json:"dialogue_text"`
	DialogueType string `json:"dialogue_type"` // greeting, question, trade, quest_offer, farewell
	Condition    string `json:"condition"`     // none, has_item, completed_quest, etc.
	ConditionVal string `json:"condition_val"`
	Order        int    `json:"order"`
}

// NPCInteraction 记录用户与NPC的交互
type NPCInteraction struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	NPCId           int       `json:"npc_id"`
	InteractionType string    `json:"interaction_type"` // talk, trade, battle, receive_quest
	InteractionData string    `json:"interaction_data"` // JSON encoded
	InteractedAt    time.Time `json:"interacted_at"`
}

// ==================== 任务系统 ====================

// Quest 任务定义
type Quest struct {
	ID          int       `json:"id"`
	QuestCode   string    `json:"quest_code"` // 唯一编码
	NameEn      string    `json:"name_en"`
	NameZh      *string   `json:"name_zh"`
	Description *string   `json:"description"`
	Type        *string   `json:"type"`       // main, side, daily
	Difficulty  *int      `json:"difficulty"` // 1-5
	GiverNPCId  *int      `json:"giver_npc_id"`
	RewardExp   *int      `json:"reward_exp"`
	RewardCoins *int      `json:"reward_coins"`
	RewardItem  *string   `json:"reward_item"` // item_code or pokemon_id
	LevelReq    *int      `json:"level_req"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserQuest 用户任务进度
type UserQuest struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	QuestID     int        `json:"quest_id"`
	Status      string     `json:"status"`   // not_started, in_progress, completed, failed
	Progress    int        `json:"progress"` // 0-100%
	StartedAt   time.Time  `json:"started_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

// QuestStep 任务步骤
type QuestStep struct {
	ID         int    `json:"id"`
	QuestID    int    `json:"quest_id"`
	StepOrder  int    `json:"step_order"`
	StepType   string `json:"step_type"`   // defeat_pokemon, catch_pokemon, visit_location, talk_to_npc, collect_item
	StepTarget string `json:"step_target"` // pokemon_id, location_id, npc_id, item_id
	StepCount  int    `json:"step_count"`  // 需要完成的数量
}

// ==================== 地牢副本系统 ====================

// Dungeon 地牢定义
type Dungeon struct {
	ID            int       `json:"id"`
	DungeonCode   string    `json:"dungeon_code"`
	NameEn        string    `json:"name_en"`
	NameZh        string    `json:"name_zh"`
	RegionID      int       `json:"region_id"`
	Difficulty    int       `json:"difficulty"` // 1-5
	FloorCount    int       `json:"floor_count"`
	BossPokemonID *string   `json:"boss_pokemon_id"`
	RewardExp     int       `json:"reward_exp"`
	RewardCoins   int       `json:"reward_coins"`
	CreatedAt     time.Time `json:"created_at"`
}

// DungeonFloor 地牢楼层
type DungeonFloor struct {
	ID           int    `json:"id"`
	DungeonID    int    `json:"dungeon_id"`
	FloorNumber  int    `json:"floor_number"`
	EnemyData    string `json:"enemy_data"`    // JSON: {pokemon_ids, levels}
	TrapData     string `json:"trap_data"`     // JSON: {trap_types, trap_positions}
	TreasureData string `json:"treasure_data"` // JSON: {items, rate}
	BossFlag     bool   `json:"boss_flag"`
}

// UserDungeonProgress 用户地牢进度
type UserDungeonProgress struct {
	ID               int        `json:"id"`
	UserID           int        `json:"user_id"`
	DungeonID        int        `json:"dungeon_id"`
	CurrentFloor     int        `json:"current_floor"`
	Status           string     `json:"status"`            // not_started, in_progress, completed, failed
	PartyComposition string     `json:"party_composition"` // JSON: {pokemon_ids}
	StartedAt        time.Time  `json:"started_at"`
	CompletedAt      *time.Time `json:"completed_at"`
}

// ==================== 体操馆系统 ====================

// Gym 体操馆
type Gym struct {
	ID          int       `json:"id"`
	GymCode     string    `json:"gym_code"`
	NameEn      string    `json:"name_en"`
	NameZh      *string   `json:"name_zh"`
	TownID      string    `json:"town_id"`
	LeaderNPCId *int      `json:"leader_npc_id"`
	TypeFocus   *string   `json:"type_focus"` // 属性: fire, water, grass, etc.
	BadgeName   *string   `json:"badge_name"`
	BadgeIcon   *string   `json:"badge_icon"`
	CreatedAt   time.Time `json:"created_at"`
}

// GymTeam 体操馆馆主的队伍
type GymTeam struct {
	ID              int    `json:"id"`
	GymID           int    `json:"gym_id"`
	LeaderPokemonID string `json:"leader_pokemon_id"`
	TeamSize        int    `json:"team_size"`
	PokemonIDs      string `json:"pokemon_ids"` // JSON array
}

// UserGymBadge 用户获得的体操馆徽章
type UserGymBadge struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	GymID         int       `json:"gym_id"`
	BadgeEarnedAt time.Time `json:"badge_earned_at"`
}

// ==================== 地区和区域系统 ====================

// Region 地区
type Region struct {
	ID          int    `json:"id"`
	NameEn      string `json:"name_en"`
	NameZh      string `json:"name_zh"`
	LevelMin    int    `json:"level_min"`
	LevelMax    int    `json:"level_max"`
	IslandCount int    `json:"island_count"`
}

// GrassArea 草地区域
type GrassArea struct {
	ID          int     `json:"id"`
	RegionID    int     `json:"region_id"`
	NameEn      string  `json:"name_en"`
	NameZh      string  `json:"name_zh"`
	CoordX      float64 `json:"coord_x"`
	CoordY      float64 `json:"coord_y"`
	Difficulty  int     `json:"difficulty"`   // 1-5
	TerrainType string  `json:"terrain_type"` // grass, cave, forest, mountain, water
}

// WildPokemonSpawn 野生宝可梦生成配置
type WildPokemonSpawn struct {
	ID               int    `json:"id"`
	AreaID           int    `json:"area_id"`
	PokemonSpeciesID string `json:"pokemon_species_id"`
	LevelMin         int    `json:"level_min"`
	LevelMax         int    `json:"level_max"`
	EncounterRate    int    `json:"encounter_rate"` // 百分比 1-100
}

// ExplorationHistory 探索历史
type ExplorationHistory struct {
	ID                 int       `json:"id"`
	UserID             int       `json:"user_id"`
	AreaID             int       `json:"area_id"`
	PokemonEncountered string    `json:"pokemon_encountered"` // JSON array
	PokemonCaptured    string    `json:"pokemon_captured"`    // JSON array
	LastVisitedAt      time.Time `json:"last_visited_at"`
}

// ==================== 地图关卡系统 ====================

// MapZone 地图区域（细化的地图单元）
type MapZone struct {
	ID             int       `json:"id"`
	ZoneCode       string    `json:"zone_code"` // 如 "kanto_route_1"
	NameEn         string    `json:"name_en"`
	NameZh         *string   `json:"name_zh"`
	Description    *string   `json:"description"`
	Type           *string   `json:"type"`  // route, cave, city, building, water
	Level          *int      `json:"level"` // 推荐等级
	IslandID       *string   `json:"island_id"`
	TownID         *string   `json:"town_id"`         // 可选，有些区域属于城镇
	ConnectedZones []string  `json:"connected_zones"` // Zone IDs connected to this zone
	CreatedAt      time.Time `json:"created_at"`
}

// Level 关卡定义
type Level struct {
	ID           int       `json:"id"`
	LevelCode    string    `json:"level_code"` // 如 "kanto_1_1"
	ZoneID       int       `json:"zone_id"`
	NameEn       string    `json:"name_en"`
	NameZh       string    `json:"name_zh"`
	LevelNum     int       `json:"level_num"`
	Type         string    `json:"type"` // normal, boss, special
	Difficulty   int       `json:"difficulty"`
	EnemyData    string    `json:"enemy_data"` // JSON
	TrapData     string    `json:"trap_data"`
	TreasureData string    `json:"treasure_data"`
	Rewards      string    `json:"rewards"` // JSON
	CreatedAt    time.Time `json:"created_at"`
}

// UserLevelProgress 用户关卡进度
type UserLevelProgress struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	LevelID     int        `json:"level_id"`
	Status      string     `json:"status"` // locked, unlocked, in_progress, completed, failed
	Score       int        `json:"score"`  // 0-3 stars
	CompletedAt *time.Time `json:"completed_at"`
	Attempts    int        `json:"attempts"`
}
