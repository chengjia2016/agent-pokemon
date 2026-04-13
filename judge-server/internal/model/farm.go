package model

import (
	"time"
)

// ==================== 菜园模型 ====================

// Crop 作物类型定义
type Crop struct {
	ID              string
	Name            string
	RecoveryTime    int64  // 秒数
	HPRecovery      int
	BonusEffect     string // 如 "speed+1", "attack+1"
	Yield           int
	RequiredLevel   int
	SeedCost        int // 金币
}

// BaseFarmPlot 基地菜园地块
type BaseFarmPlot struct {
	ID         int        `json:"id"`
	BaseID     string     `json:"base_id"`
	GitHubID   int        `json:"github_id"`
	PlotNumber int        `json:"plot_number"` // 1-8
	CropType   string     `json:"crop_type"`
	PlantedAt  time.Time  `json:"planted_at"`
	ReadyAt    time.Time  `json:"ready_at"`
	Quantity   int        `json:"quantity"`
	Status     string     `json:"status"` // growing / ready / empty
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// FarmHarvest 菜园收获记录
type FarmHarvest struct {
	ID         int       `json:"id"`
	BaseID     string    `json:"base_id"`
	GitHubID   int       `json:"github_id"`
	CropType   string    `json:"crop_type"`
	Quantity   int       `json:"quantity"`
	HarvestedAt time.Time `json:"harvested_at"`
}

// ==================== 基地访问与战斗模型 ====================

// BaseVisit 基地访问记录
type BaseVisit struct {
	ID              int        `json:"id"`
	BaseID          string     `json:"base_id"`
	OwnerGitHubID   int        `json:"owner_github_id"`
	VisitorGitHubID int        `json:"visitor_github_id"`
	VisitType       string     `json:"visit_type"` // duel / attack
	VisitorPetID    string     `json:"visitor_pet_id"`
	DefenderPetID   string     `json:"defender_pet_id"`
	BattleResult    string     `json:"battle_result"` // win / lose / draw
	PointsGained    int        `json:"points_gained"`
	ReputationChange int       `json:"reputation_change"`
	CoinsLooted     int        `json:"coins_looted"`
	VisitedAt       time.Time  `json:"visited_at"`
}

// FarmEating 进食记录
type FarmEating struct {
	ID              int       `json:"id"`
	BaseID          string    `json:"base_id"`
	OwnerGitHubID   int       `json:"owner_github_id"`
	VisitorGitHubID int       `json:"visitor_github_id"`
	VisitorPetID    string    `json:"visitor_pet_id"`
	CropType        string    `json:"crop_type"`
	HPRecovered     int       `json:"hp_recovered"`
	EatenAt         time.Time `json:"eaten_at"`
}

// UserReputation 用户声望
type UserReputation struct {
	ID             int       `json:"id"`
	GitHubID       int       `json:"github_id"`
	TotalPoints    int       `json:"total_points"`
	TotalReputation int      `json:"total_reputation"`
	DefenseWins    int       `json:"defense_wins"`
	DefenseLosses  int       `json:"defense_losses"`
	AttackWins     int       `json:"attack_wins"`
	AttackLosses   int       `json:"attack_losses"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ==================== 请求/响应模型 ====================

// PlantCropRequest 种植作物请求
type PlantCropRequest struct {
	PlotNumber int    `json:"plot_number"` // 1-8
	CropType   string `json:"crop_type"`
	Quantity   int    `json:"quantity"`
}

// HarvestCropRequest 收获作物请求
type HarvestCropRequest struct {
	PlotNumber int `json:"plot_number"` // 1-8
}

// EatFarmRequest 进食请求
type EatFarmRequest struct {
	VisitorGitHubID int    `json:"visitor_github_id"`
	VisitorPetID    string `json:"visitor_pet_id"`
	CropType        string `json:"crop_type"`
}

// VisitBaseRequest 访问基地请求
type VisitBaseRequest struct {
	VisitType    string `json:"visit_type"` // duel / attack
	VisitorPetID string `json:"visitor_pet_id"`
}

// ==================== 常量 ====================

// 作物定义
var Crops = map[string]*Crop{
	"strawberry": {
		ID:            "strawberry",
		Name:          "草莓",
		RecoveryTime:  3600,
		HPRecovery:    20,
		BonusEffect:   "",
		Yield:         3,
		RequiredLevel: 1,
		SeedCost:      50,
	},
	"blueberry": {
		ID:            "blueberry",
		Name:          "蓝莓",
		RecoveryTime:  5400,
		HPRecovery:    35,
		BonusEffect:   "speed+1",
		Yield:         2,
		RequiredLevel: 3,
		SeedCost:      100,
	},
	"cherry": {
		ID:            "cherry",
		Name:          "樱桃",
		RecoveryTime:  7200,
		HPRecovery:    50,
		BonusEffect:   "attack+1",
		Yield:         2,
		RequiredLevel: 5,
		SeedCost:      150,
	},
	"tomato": {
		ID:            "tomato",
		Name:          "番茄",
		RecoveryTime:  10800,
		HPRecovery:    60,
		BonusEffect:   "defense+1",
		Yield:         2,
		RequiredLevel: 8,
		SeedCost:      200,
	},
	"golden_apple": {
		ID:            "golden_apple",
		Name:          "金苹果",
		RecoveryTime:  14400,
		HPRecovery:    100,
		BonusEffect:   "sp_attack+1",
		Yield:         1,
		RequiredLevel: 12,
		SeedCost:      300,
	},
}

// 声望等级
const (
	ReputationLegend     = 100
	ReputationRespected  = 50
	ReputationKnown      = 20
	ReputationNovice     = 0
	ReputationVillain    = -1
)

// 访问类型
const (
	VisitTypeDuel   = "duel"
	VisitTypeAttack = "attack"
)

// 战斗结果
const (
	BattleResultWin  = "win"
	BattleResultLose = "lose"
	BattleResultDraw = "draw"
)
