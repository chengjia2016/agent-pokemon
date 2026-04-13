package service

import (
	"errors"
	"judge-server/internal/model"
	"math/rand"
	"time"
)

// CaptureSystem 宝可梦捕捉系统
type CaptureSystem struct {
	breedingService *PokemonBreedingService
	db              Database
}

// NewCaptureSystem 创建新的捕捉系统
func NewCaptureSystem(bs *PokemonBreedingService, db Database) *CaptureSystem {
	return &CaptureSystem{
		breedingService: bs,
		db:              db,
	}
}

// CaptureAttempt 捕捉尝试
type CaptureAttempt struct {
	WildPokemonID   string
	PokemonLevel    int
	CurrentHP       int
	MaxHP           int
	StatusCondition string
	BallType        string
}

// AttemptCapture 尝试捕捉宝可梦
func (cs *CaptureSystem) AttemptCapture(
	userID int,
	attempt *CaptureAttempt,
	wildPokemon *model.PokemonSpecies,
) (bool, int, error) {
	if wildPokemon == nil {
		return false, 0, errors.New("invalid wild pokemon")
	}

	// 获取捕捉率
	captureRate := cs.breedingService.CalculateCaptureRate(
		wildPokemon.CaptureRate,
		attempt.MaxHP,
		attempt.CurrentHP,
		[]string{attempt.StatusCondition},
		attempt.BallType,
	)

	// 进行捕捉尝试
	success := cs.breedingService.AttemptCapture(captureRate)

	// 记录捕捉尝试
	cs.logCaptureAttempt(userID, attempt, captureRate, success)

	return success, captureRate, nil
}

// logCaptureAttempt 记录捕捉尝试
func (cs *CaptureSystem) logCaptureAttempt(
	userID int,
	attempt *CaptureAttempt,
	captureRate int,
	success bool,
) {
	// TODO: 保存到数据库
	// INSERT INTO capture_attempts (user_id, wild_pokemon_id, ball_type, pokemon_hp_percent, status_condition, capture_rate_percent, success, attempted_at)
	// VALUES (userID, attempt.WildPokemonID, attempt.BallType, hpPercent, attempt.StatusCondition, captureRate, success, NOW())
}

// GetAllBallTypes 获取所有精灵球类型
func (cs *CaptureSystem) GetAllBallTypes() []map[string]interface{} {
	ballTypes := []map[string]interface{}{
		{
			"id":          "poke_ball",
			"name_en":     "Poké Ball",
			"name_zh":     "精灵球",
			"multiplier":  1.0,
			"description": "普通精灵球",
			"cost":        200,
		},
		{
			"id":          "great_ball",
			"name_en":     "Great Ball",
			"name_zh":     "超级球",
			"multiplier":  1.5,
			"description": "增强的精灵球",
			"cost":        600,
		},
		{
			"id":          "ultra_ball",
			"name_en":     "Ultra Ball",
			"name_zh":     "高级球",
			"multiplier":  2.0,
			"description": "强大的精灵球",
			"cost":        1000,
		},
		{
			"id":          "master_ball",
			"name_en":     "Master Ball",
			"name_zh":     "大师球",
			"multiplier":  100.0,
			"description": "必定捕获（传说宝可梦除外）",
			"cost":        0,
			"rare":        true,
		},
		{
			"id":          "fast_ball",
			"name_en":     "Fast Ball",
			"name_zh":     "快速球",
			"multiplier":  4.0,
			"description": "对速度高的宝可梦有效",
			"condition":   "target_speed_high",
		},
		{
			"id":          "level_ball",
			"name_en":     "Level Ball",
			"name_zh":     "等级球",
			"multiplier":  4.0,
			"description": "对等级低于己方宝可梦的目标有效",
			"condition":   "target_level_low",
		},
		{
			"id":          "lure_ball",
			"name_en":     "Lure Ball",
			"name_zh":     "鱼竿球",
			"multiplier":  3.0,
			"description": "对钓鱼遇到的宝可梦有效",
			"condition":   "fishing",
		},
		{
			"id":          "moon_ball",
			"name_en":     "Moon Ball",
			"name_zh":     "月亮球",
			"multiplier":  4.0,
			"description": "对用月石进化的宝可梦有效",
			"condition":   "moon_stone_evolution",
		},
		{
			"id":          "net_ball",
			"name_en":     "Net Ball",
			"name_zh":     "渔网球",
			"multiplier":  3.0,
			"description": "对水系和虫系宝可梦有效",
			"condition":   "water_or_bug_type",
		},
		{
			"id":          "dive_ball",
			"name_en":     "Dive Ball",
			"name_zh":     "潜水球",
			"multiplier":  3.5,
			"description": "对水中遇到的宝可梦有效",
			"condition":   "water_encounter",
		},
		{
			"id":          "nest_ball",
			"name_en":     "Nest Ball",
			"name_zh":     "巢穴球",
			"multiplier":  3.0,
			"description": "对低等级宝可梦有效",
			"condition":   "low_level",
		},
		{
			"id":          "repeat_ball",
			"name_en":     "Repeat Ball",
			"name_zh":     "重复球",
			"multiplier":  3.0,
			"description": "对已捕捉过的宝可梦有效",
			"condition":   "already_owned",
		},
		{
			"id":          "timer_ball",
			"name_en":     "Timer Ball",
			"name_zh":     "计时球",
			"multiplier":  4.0,
			"description": "战斗回合数越多效果越好",
			"condition":   "battle_duration",
		},
		{
			"id":          "heavy_ball",
			"name_en":     "Heavy Ball",
			"name_zh":     "重球",
			"multiplier":  4.0,
			"description": "对重量大的宝可梦有效",
			"condition":   "heavy_pokemon",
		},
		{
			"id":          "love_ball",
			"name_en":     "Love Ball",
			"name_zh":     "爱心球",
			"multiplier":  8.0,
			"description": "对与己方宝可梦同性别的异性有效",
			"condition":   "opposite_gender",
		},
	}

	return ballTypes
}

// ==================== 孵蛋系统 ====================

// EggHatchingSystem 蛋孵化系统
type EggHatchingSystem struct {
	db Database
}

// NewEggHatchingSystem 创建新的孵蛋系统
func NewEggHatchingSystem(db Database) *EggHatchingSystem {
	return &EggHatchingSystem{db: db}
}

// EggData 蛋数据
type EggData struct {
	ID          int       `json:"id"`
	OwnerID     int       `json:"owner_id"`
	FatherID    string    `json:"father_id"`    // 公宝可梦
	MotherID    string    `json:"mother_id"`    // 母宝可梦
	StepsNeeded int       `json:"steps_needed"` // 孵化所需步数
	StepsTaken  int       `json:"steps_taken"`  // 已走步数
	Status      string    `json:"status"`       // 孵化中/已孵化/已领取
	HatchedAt   time.Time `json:"hatched_at,omitempty"`
	ClaimedAt   time.Time `json:"claimed_at,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// ProduceEgg 产生蛋
func (ehs *EggHatchingSystem) ProduceEgg(
	ownerID int,
	fatherID, motherID string,
) (*EggData, error) {
	// TODO: 检查两只宝可梦是否可以繁殖（兼容的蛋组）

	// 产生的宝可梦等级随父母高
	egg := &EggData{
		OwnerID:     ownerID,
		FatherID:    fatherID,
		MotherID:    motherID,
		StepsNeeded: 10000, // 默认10000步孵化
		StepsTaken:  0,
		Status:      "incubating",
		CreatedAt:   time.Now(),
	}

	return egg, nil
}

// IncreaseHatchProgress 增加孵蛋进度
func (ehs *EggHatchingSystem) IncreaseHatchProgress(eggID int, steps int) (*EggData, error) {
	// TODO: 从数据库获取蛋数据
	// 增加进度
	// 如果达到所需步数，标记为已孵化

	return nil, nil
}

// HatchEgg 孵化蛋为宝可梦
func (ehs *EggHatchingSystem) HatchEgg(eggID int) (*model.UserPokemonFull, error) {
	// TODO: 生成新宝可梦，继承父母的一些属性
	// 新宝可梦应该有高个体值（因为是通过百变怪或良好组合孵化的）
	// 学会父母的某些招式

	return nil, nil
}

// ==================== 宝可梦池管理 ====================

// WildPokemonPool 野生宝可梦池
type WildPokemonPool struct {
	AreaID   string
	Pokemons []WildPokemonSpawnRate
}

// WildPokemonSpawnRate 野生宝可梦生成率
type WildPokemonSpawnRate struct {
	SpeciesID     string
	LevelMin      int
	LevelMax      int
	EncounterRate int    // 概率
	TimeOfDay     string // morning/day/evening/night
}

// GenerateWildPokemon 生成随机野生宝可梦
func (cs *CaptureSystem) GenerateWildPokemon(
	pool []WildPokemonSpawnRate,
) (string, int, error) {
	if len(pool) == 0 {
		return "", 0, errors.New("empty wild pokemon pool")
	}

	// 计算总概率
	totalRate := 0
	for _, spawn := range pool {
		totalRate += spawn.EncounterRate
	}

	// 随机选择宝可梦
	randomValue := rand.Intn(totalRate)
	currentRate := 0

	var selectedSpawn WildPokemonSpawnRate
	for _, spawn := range pool {
		currentRate += spawn.EncounterRate
		if randomValue < currentRate {
			selectedSpawn = spawn
			break
		}
	}

	// 生成随机等级
	level := selectedSpawn.LevelMin + rand.Intn(selectedSpawn.LevelMax-selectedSpawn.LevelMin+1)

	return selectedSpawn.SpeciesID, level, nil
}
