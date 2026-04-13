package model

import (
	"time"
)

// ==================== 队伍管理模型 ====================

// UserTeam 用户的战斗队伍
type UserTeam struct {
	ID            int              `json:"id"`
	GitHubID      int              `json:"github_id"`
	TeamID        string           `json:"team_id"`
	TeamName      string           `json:"team_name"`
	Description   string           `json:"description"`
	IsDefenseTeam bool             `json:"is_defense_team"`
	MaxMembers    int              `json:"max_members"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	Members       []UserTeamMember `json:"members,omitempty"`
}

// UserTeamMember 队伍中的成员（宠物）
type UserTeamMember struct {
	ID           int       `json:"id"`
	TeamID       string    `json:"team_id"`
	PetID        string    `json:"pet_id"`
	SlotPosition int       `json:"slot_position"`
	JoinedAt     time.Time `json:"joined_at"`
	Pet          *Pet      `json:"pet,omitempty"`
}

// UserOwnedPokemon 用户拥有的宠物
type UserOwnedPokemon struct {
	ID         int        `json:"id"`
	GitHubID   int        `json:"github_id"`
	PetID      string     `json:"pet_id"`
	IsCaptured bool       `json:"is_captured"`
	Status     string     `json:"status"` // active, fainted, training
	CurrentHP  int        `json:"current_hp"`
	MaxHP      int        `json:"max_hp"`
	Level      int        `json:"level"`
	Experience int        `json:"experience"`
	CapturedAt *time.Time `json:"captured_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	PetDetail  *Pet       `json:"pet_detail,omitempty"`
}

// ==================== 防守基地模型 ====================

// UserBase 用户的防守基地
type UserBase struct {
	ID            int       `json:"id"`
	GitHubID      int       `json:"github_id"`
	BaseID        string    `json:"base_id"`
	RepositoryURL string    `json:"repository_url"`
	DefenseTeamID *string   `json:"defense_team_id"`
	Level         int       `json:"level"`
	Prestige      int       `json:"prestige"`
	DefenseWins   int       `json:"defense_wins"`
	DefenseLosses int       `json:"defense_losses"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DefenseTeam   *UserTeam `json:"defense_team,omitempty"`
}

// BaseDefenseRecord 基地防守记录
type BaseDefenseRecord struct {
	ID               int       `json:"id"`
	BaseID           string    `json:"base_id"`
	AttackerGitHubID int       `json:"attacker_github_id"`
	BattleID         string    `json:"battle_id"`
	Result           string    `json:"result"` // win or lose
	AttackerLevel    int       `json:"attacker_level"`
	DefenderLevel    int       `json:"defender_level"`
	CoinsGained      float64   `json:"coins_gained"`
	Timestamp        time.Time `json:"timestamp"`
}

// ==================== 扩展战斗模型 ====================

// BattleV2 扩展的战斗记录（与现有Battle共存）
type BattleV2 struct {
	ID               int           `json:"id"`
	BattleID         string        `json:"battle_id"`
	AttackerGitHubID int           `json:"attacker_github_id"`
	DefenderGitHubID int           `json:"defender_github_id"`
	AttackerTeamID   string        `json:"attacker_team_id"`
	DefenderTeamID   string        `json:"defender_team_id"`
	BattleType       string        `json:"battle_type"` // pvp or base_defense
	WinnerGitHubID   *int          `json:"winner_github_id,omitempty"`
	Status           string        `json:"status"` // ongoing, completed, abandoned
	RewardCoins      float64       `json:"reward_coins"`
	StartedAt        time.Time     `json:"started_at"`
	EndedAt          *time.Time    `json:"ended_at,omitempty"`
	Rounds           []BattleRound `json:"rounds,omitempty"`
}

// BattleRound 战斗回合
type BattleRound struct {
	ID                   int       `json:"id"`
	BattleID             string    `json:"battle_id"`
	RoundNumber          int       `json:"round_number"`
	AttackerPetID        string    `json:"attacker_pet_id"`
	DefenderPetID        string    `json:"defender_pet_id"`
	AttackerDamage       int       `json:"attacker_damage"`
	DefenderDamage       int       `json:"defender_damage"`
	AttackerPetCurrentHP int       `json:"attacker_pet_current_hp"`
	DefenderPetCurrentHP int       `json:"defender_pet_current_hp"`
	RoundWinner          string    `json:"round_winner"` // attacker, defender, draw
	CreatedAt            time.Time `json:"created_at"`
}

// PetBattleStats 宠物战斗统计
type PetBattleStats struct {
	ID               int        `json:"id"`
	PetID            string     `json:"pet_id"`
	TotalBattles     int        `json:"total_battles"`
	Wins             int        `json:"wins"`
	Losses           int        `json:"losses"`
	Draws            int        `json:"draws"`
	TotalDamageDealt int        `json:"total_damage_dealt"`
	TotalDamageTaken int        `json:"total_damage_taken"`
	LastBattleAt     *time.Time `json:"last_battle_at,omitempty"`
}

// BattleReward 战斗奖励
type BattleReward struct {
	ID               int       `json:"id"`
	BattleID         string    `json:"battle_id"`
	WinnerGitHubID   int       `json:"winner_github_id"`
	LoserGitHubID    int       `json:"loser_github_id"`
	CoinsTransferred float64   `json:"coins_transferred"`
	Percentage       float64   `json:"percentage"` // 1-5%
	TransactionID    string    `json:"transaction_id"`
	CompletedAt      time.Time `json:"completed_at"`
}

// ==================== 野生精灵模型 ====================

// WildPokemon 地图上的野生精灵（NPC）
type WildPokemon struct {
	ID                 int             `json:"id"`
	WildID             string          `json:"wild_id"`
	LocationID         string          `json:"location_id"`
	PokemonSpeciesID   string          `json:"pokemon_species_id"`
	Level              int             `json:"level"`
	Status             string          `json:"status"` // active, captured, defeated
	CurrentHP          int             `json:"current_hp"`
	MaxHP              int             `json:"max_hp"`
	CaptureDifficulty  int             `json:"capture_difficulty"` // 1-100
	CaughtByGitHubID   *int            `json:"caught_by_github_id,omitempty"`
	DefeatedByGitHubID *int            `json:"defeated_by_github_id,omitempty"`
	CreatedAt          time.Time       `json:"created_at"`
	CapturedAt         *time.Time      `json:"captured_at,omitempty"`
	DefeatedAt         *time.Time      `json:"defeated_at,omitempty"`
	UpdatedAt          time.Time       `json:"updated_at"`
	SpeciesDetail      *PokemonSpecies `json:"species_detail,omitempty"`
}

// CaptureHistory 捕获历史
type CaptureHistory struct {
	ID            int       `json:"id"`
	GitHubID      int       `json:"github_id"`
	PetID         string    `json:"pet_id"`
	WildPokemonID *string   `json:"wild_pokemon_id,omitempty"`
	CaptureType   string    `json:"capture_type"` // wild, hatched, traded
	Success       bool      `json:"success"`
	AttemptCount  int       `json:"attempt_count"`
	CapturedAt    time.Time `json:"captured_at"`
}

// ==================== 请求/响应模型 ====================

// CreateTeamRequest 创建队伍请求
type CreateTeamRequest struct {
	TeamName      string `json:"team_name"`
	Description   string `json:"description"`
	IsDefenseTeam bool   `json:"is_defense_team"`
}

// AddPetToTeamRequest 将宠物加入队伍请求
type AddPetToTeamRequest struct {
	TeamID       string `json:"team_id"`
	PetID        string `json:"pet_id"`
	SlotPosition int    `json:"slot_position"`
}

// StartBattleRequest 开始战斗请求
type StartBattleRequest struct {
	AttackerID     int    `json:"attacker_id"`
	DefenderID     int    `json:"defender_id"`
	AttackerTeamID string `json:"attacker_team_id"`
	DefenderTeamID string `json:"defender_team_id"`
	BattleType     string `json:"battle_type"` // pvp or base_defense
}

// CapturePokemonRequest 捕获精灵请求
type CapturePokemonRequest struct {
	WildPokemonID string `json:"wild_pokemon_id"`
}

// CreateBaseRequest 创建防守基地请求
type CreateBaseRequest struct {
	RepositoryURL string `json:"repository_url"`
}

// ==================== 状态常量 ====================

const (
	// 宠物状态
	PokemonStatusActive   = "active"
	PokemonStatusFainted  = "fainted"
	PokemonStatusTraining = "training"

	// 战斗类型
	BattleTypePVP         = "pvp"
	BattleTypeBaseDefense = "base_defense"

	// 战斗状态
	BattleStatusOngoing   = "ongoing"
	BattleStatusCompleted = "completed"
	BattleStatusAbandoned = "abandoned"

	// 防守基地结果
	DefenseResultWin  = "win"
	DefenseResultLose = "lose"

	// 野生精灵状态
	WildStatusActive   = "active"
	WildStatusCaptured = "captured"
	WildStatusDefeated = "defeated"

	// 捕获类型
	CaptureTypeWild    = "wild"
	CaptureTypeHatched = "hatched"
	CaptureTypeTraded  = "traded"

	// 回合赢家
	RoundWinnerAttacker = "attacker"
	RoundWinnerDefender = "defender"
	RoundWinnerDraw     = "draw"
)
