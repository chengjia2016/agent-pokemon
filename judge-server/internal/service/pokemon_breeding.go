package service

import (
	"errors"
	"fmt"
	"judge-server/internal/model"
	"math"
	"math/rand"
	"time"
)

// PokemonBreedingService 宝可梦养成服务
type PokemonBreedingService struct {
	db Database
}

// NewPokemonBreedingService 创建新的养成服务
func NewPokemonBreedingService(db Database) *PokemonBreedingService {
	return &PokemonBreedingService{db: db}
}

// ==================== 能力值计算 ====================

// CalculatePokemonStats 计算宝可梦的实际能力值
// 公式：
// HP = ((2*base + iv + ev/4) * level) / 100 + level + 5
// 其他 = ((2*base + iv + ev/4) * level) / 100 + 5) * 性格倍数
func (pbs *PokemonBreedingService) CalculatePokemonStats(
	baseStats *model.PokemonBaseStats,
	ivStats *model.PokemonIndividualStats,
	evStats *model.PokemonEffortStats,
	nature *model.PokemonNature,
	level int,
) *model.CalculatedStats {

	if baseStats == nil || level <= 0 {
		return nil
	}

	// 创建默认的IV和EV如果未提供
	if ivStats == nil {
		ivStats = &model.PokemonIndividualStats{
			IVHP: 31, IVAttack: 31, IVDefense: 31,
			IVSpAtk: 31, IVSpDef: 31, IVSpeed: 31,
		}
	}
	if evStats == nil {
		evStats = &model.PokemonEffortStats{}
	}

	// 计算HP（公式略有不同）
	hp := int((float64(2*baseStats.HP+ivStats.IVHP+evStats.EVHP/4)*float64(level))/100) + level + 5

	// 计算其他能力值
	attack := pbs.calculateStat(baseStats.Attack, ivStats.IVAttack, evStats.EVAttack, level, nature, "attack")
	defense := pbs.calculateStat(baseStats.Defense, ivStats.IVDefense, evStats.EVDefense, level, nature, "defense")
	spAtk := pbs.calculateStat(baseStats.SpAtk, ivStats.IVSpAtk, evStats.EVSpAtk, level, nature, "sp_atk")
	spDef := pbs.calculateStat(baseStats.SpDef, ivStats.IVSpDef, evStats.EVSpDef, level, nature, "sp_def")
	speed := pbs.calculateStat(baseStats.Speed, ivStats.IVSpeed, evStats.EVSpeed, level, nature, "speed")

	return &model.CalculatedStats{
		HP:      hp,
		Attack:  attack,
		Defense: defense,
		SpAtk:   spAtk,
		SpDef:   spDef,
		Speed:   speed,
	}
}

// calculateStat 计算单个能力值
func (pbs *PokemonBreedingService) calculateStat(
	base, iv, ev, level int,
	nature *model.PokemonNature,
	statName string,
) int {
	// 基础计算
	stat := int((float64(2*base+iv+ev/4)*float64(level))/100) + 5

	// 应用性格倍数
	if nature != nil {
		multiplier := 1.0
		if nature.IncreasedStat == statName {
			multiplier = 1.1
		} else if nature.DecreasedStat == statName {
			multiplier = 0.9
		}
		stat = int(float64(stat) * multiplier)
	}

	return stat
}

// ==================== 宝可梦生成 ====================

// GenerateRandomIVs 生成随机个体值（IV）
func (pbs *PokemonBreedingService) GenerateRandomIVs() *model.PokemonIndividualStats {
	rand.Seed(time.Now().UnixNano())
	return &model.PokemonIndividualStats{
		IVHP:      rand.Intn(32),
		IVAttack:  rand.Intn(32),
		IVDefense: rand.Intn(32),
		IVSpAtk:   rand.Intn(32),
		IVSpDef:   rand.Intn(32),
		IVSpeed:   rand.Intn(32),
	}
}

// GenerateRandomNature 生成随机性格
func (pbs *PokemonBreedingService) GenerateRandomNature() string {
	natures := []string{
		"nature_hardy", "nature_lonely", "nature_brave", "nature_adamant", "nature_naughty",
		"nature_bold", "nature_docile", "nature_relaxed", "nature_impish", "nature_lax",
		"nature_timid", "nature_hasty", "nature_serious", "nature_jolly", "nature_naive",
		"nature_modest", "nature_mild", "nature_quiet", "nature_bashful", "nature_rash",
		"nature_calm", "nature_gentle", "nature_sassy", "nature_careful", "nature_quirky",
	}
	return natures[rand.Intn(len(natures))]
}

// GenerateRandomGender 生成随机性别
func (pbs *PokemonBreedingService) GenerateRandomGender(genderRatioPercent int) string {
	if genderRatioPercent < 0 || genderRatioPercent > 100 {
		genderRatioPercent = 50
	}
	if rand.Intn(100) < genderRatioPercent {
		return "male"
	}
	return "female"
}

// ==================== 努力值调整 ====================

// AddEffortValues 增加宝可梦的努力值
// 每击败一只宝可梦会获得该宝可梦对应的努力值
func (pbs *PokemonBreedingService) AddEffortValues(
	userPokemonID int,
	evBonus map[string]int,
) error {
	if evBonus == nil {
		return errors.New("ev_bonus cannot be nil")
	}

	// 获取当前的努力值
	currentEV, err := pbs.db.GetEffortStats(userPokemonID)
	if err != nil {
		return fmt.Errorf("failed to get effort stats: %w", err)
	}

	// 应用努力值加成
	currentEV.EVHP = pbs.capEV(currentEV.EVHP + evBonus["hp"])
	currentEV.EVAttack = pbs.capEV(currentEV.EVAttack + evBonus["attack"])
	currentEV.EVDefense = pbs.capEV(currentEV.EVDefense + evBonus["defense"])
	currentEV.EVSpAtk = pbs.capEV(currentEV.EVSpAtk + evBonus["sp_atk"])
	currentEV.EVSpDef = pbs.capEV(currentEV.EVSpDef + evBonus["sp_def"])
	currentEV.EVSpeed = pbs.capEV(currentEV.EVSpeed + evBonus["speed"])

	// 检查总EV是否超过510
	totalEV := currentEV.EVHP + currentEV.EVAttack + currentEV.EVDefense +
		currentEV.EVSpAtk + currentEV.EVSpDef + currentEV.EVSpeed
	if totalEV > 510 {
		return errors.New("total EV cannot exceed 510")
	}

	currentEV.TotalEV = totalEV
	currentEV.UpdatedAt = time.Now()

	// 保存更新
	return pbs.db.UpdateEffortStats(currentEV)
}

// capEV 限制努力值的最大值为252
func (pbs *PokemonBreedingService) capEV(ev int) int {
	if ev > 252 {
		return 252
	}
	if ev < 0 {
		return 0
	}
	return ev
}

// ReduceEffortWithBerry 使用树果减少特定能力值的努力值
// 树果会减少10点EV，但减少速度会增加不喜欢的口味
func (pbs *PokemonBreedingService) ReduceEffortWithBerry(
	userPokemonID int,
	berryID string,
	pokemon *model.UserPokemonFull,
) error {
	// TODO: 实现树果使用逻辑
	return nil
}

// ==================== 性格修改 ====================

// ChangeNature 使用性格薄荷修改宝可梦性格
func (pbs *PokemonBreedingService) ChangeNature(
	userPokemonID int,
	newNatureID string,
) error {
	// 验证性格存在
	nature, err := pbs.db.GetNature(newNatureID)
	if err != nil {
		return fmt.Errorf("invalid nature: %w", err)
	}

	// 更新宝可梦性格
	// TODO: 检查用户是否拥有性格薄荷
	return pbs.db.UpdatePokemonNature(userPokemonID, nature.NatureID)
}

// ==================== 捕捉相关 ====================

// CalculateCaptureRate 计算捕捉概率
// 基础公式受以下影响：
// 1. 宝可梦的基础捕捉率
// 2. 宝可梦当前HP百分比
// 3. 宝可梦的状态条件（降低能力值）
// 4. 精灵球的类型
func (pbs *PokemonBreedingService) CalculateCaptureRate(
	baseCapture int,
	maxHP, currentHP int,
	statusConditions []string,
	ballType string,
) int {
	// 基础捕捉率
	rate := float64(baseCapture)

	// HP因素（HP越低，捕捉率越高）
	hpPercent := float64(currentHP) / float64(maxHP)
	if hpPercent > 0.5 {
		rate *= 0.5
	} else if hpPercent > 0.25 {
		rate *= 1.0
	} else {
		rate *= 1.5
	}

	// 状态条件因素
	statusMultiplier := 1.0
	for _, condition := range statusConditions {
		switch condition {
		case "status_sleep", "status_freeze":
			statusMultiplier *= 2.0
		case "status_paralysis", "status_poison", "status_burn":
			statusMultiplier *= 1.5
		case "status_confusion":
			statusMultiplier *= 1.0
		}
	}
	rate *= statusMultiplier

	// 精灵球类型因素
	ballMultiplier := pbs.getBallMultiplier(ballType)
	rate *= ballMultiplier

	// 限制概率为0-100
	if rate > 100 {
		rate = 100
	}
	if rate < 1 {
		rate = 1
	}

	return int(rate)
}

// getBallMultiplier 获取精灵球的捕捉倍数
func (pbs *PokemonBreedingService) getBallMultiplier(ballType string) float64 {
	ballMultipliers := map[string]float64{
		"poke_ball":       1.0,
		"great_ball":      1.5,
		"ultra_ball":      2.0,
		"master_ball":     math.Inf(1), // 100%捕捉
		"fast_ball":       4.0,         // 速度高的宝可梦
		"level_ball":      4.0,         // 等级低的宝可梦
		"lure_ball":       3.0,         // 钓鱼遇到的宝可梦
		"moon_ball":       4.0,         // 用月石进化的宝可梦
		"net_ball":        3.0,         // 水系/虫系宝可梦
		"dive_ball":       3.5,         // 水中遇到的宝可梦
		"nest_ball":       3.0,         // 低等级宝可梦
		"repeat_ball":     3.0,         // 已捕捉过的宝可梦
		"timer_ball":      4.0,         // 战斗回合数多时有效
		"luxury_ball":     1.0,         // 普通球的高级版本
		"dive_ball_water": 3.5,
		"heavy_ball":      4.0, // 重量大的宝可梦
		"love_ball":       8.0, // 与宝可梦同性别时
		"moon_ball_dark":  4.0,
		"safari_ball":     1.5, // 狩猎地带
		"sport_ball":      1.5, // 比赛中
		"park_ball":       1.5, // 公园
	}

	if multiplier, exists := ballMultipliers[ballType]; exists {
		return multiplier
	}
	return 1.0 // 默认倍数
}

// AttemptCapture 尝试捕捉宝可梦
func (pbs *PokemonBreedingService) AttemptCapture(
	captureRate int,
) bool {
	// 捕捉公式：只要随机数小于捕捉率就成功
	randomValue := rand.Intn(100)
	return randomValue < captureRate
}

// ==================== 类型克制 ====================

// GetTypeMatchup 获取类型克制关系
func (pbs *PokemonBreedingService) GetTypeMatchup(attackType, defendType string) float32 {
	// 类型克制表
	matchups := map[string]map[string]float32{
		"normal": {
			"rock": 0.5, "ghost": 0, "steel": 0.5,
		},
		"fire": {
			"fire": 0.5, "water": 0.5, "grass": 2, "ice": 2, "bug": 2,
			"steel": 2, "fairy": 1, "rock": 0.5, "dragon": 0.5, "ground": 1,
		},
		"water": {
			"fire": 2, "water": 0.5, "grass": 0.5, "ground": 2, "rock": 2,
			"ice": 1, "steel": 1, "flying": 1, "psychic": 1, "normal": 1,
		},
		"grass": {
			"fire": 0.5, "water": 2, "grass": 0.5, "poison": 0.5, "ground": 2,
			"rock": 2, "bug": 0.5, "flying": 0.5, "ice": 0.5, "dragon": 0.5,
		},
		"electric": {
			"water": 2, "electric": 0.5, "grass": 0.5, "flying": 2, "ground": 0,
			"steel": 1, "fire": 1, "ice": 1, "dragon": 0.5,
		},
		"ice": {
			"fire": 0.5, "water": 0.5, "grass": 2, "ice": 0.5, "flying": 2,
			"ground": 2, "dragon": 2, "steel": 0.5, "rock": 1, "normal": 1,
		},
		"fighting": {
			"normal": 2, "ice": 2, "rock": 2, "dark": 2, "steel": 2,
			"flying": 0.5, "poison": 0.5, "psychic": 0.5, "bug": 0.5, "fairy": 0.5,
			"ghost": 0, "grass": 1, "water": 1, "electric": 1, "fire": 1,
		},
		"poison": {
			"grass": 2, "fairy": 2, "poison": 0.5, "ground": 0.5, "rock": 0.5,
			"ghost": 0.5, "steel": 0, "psychic": 1, "bug": 1, "fire": 1,
			"water": 1, "electric": 1, "ice": 1, "flying": 1, "dragon": 1,
		},
		"ground": {
			"fire": 2, "electric": 2, "poison": 2, "rock": 2, "steel": 2,
			"grass": 0.5, "bug": 0.5, "flying": 0, "water": 1, "ice": 1,
			"fighting": 1, "dragon": 1, "psychic": 1, "ghost": 1, "fairy": 1,
		},
		"flying": {
			"fighting": 2, "bug": 2, "grass": 2, "rock": 0.5, "steel": 0.5,
			"electric": 0.5, "psychic": 1, "ghost": 1, "dragon": 1, "dark": 1,
			"normal": 1, "fire": 1, "water": 1, "ice": 1, "poison": 1, "ground": 0,
		},
		"psychic": {
			"fighting": 2, "poison": 2, "psychic": 0.5, "dark": 0, "steel": 0.5,
			"bug": 1, "ghost": 1, "fire": 1, "water": 1, "grass": 1, "ice": 1,
			"ground": 1, "rock": 1, "flying": 1, "dragon": 1, "fairy": 1, "electric": 1,
		},
		"rock": {
			"fire": 2, "ice": 2, "flying": 2, "bug": 2, "water": 0.5,
			"grass": 0.5, "fighting": 0.5, "ground": 0.5, "steel": 0.5,
			"normal": 1, "poison": 1, "electric": 1, "psychic": 1, "ghost": 1,
			"dragon": 1, "dark": 1, "fairy": 1,
		},
		"ghost": {
			"psychic": 2, "ghost": 2, "dark": 0.5, "normal": 0, "fighting": 0,
			"poison": 1, "bug": 1, "fire": 1, "water": 1, "grass": 1, "ice": 1,
			"electric": 1, "ground": 1, "flying": 1, "rock": 1, "steel": 1, "fairy": 1,
			"dragon": 1,
		},
		"dragon": {
			"dragon": 2, "steel": 0.5, "fairy": 0, "fire": 1, "water": 1,
			"grass": 1, "ice": 1, "electric": 1, "poison": 1, "ground": 1,
			"flying": 1, "psychic": 1, "bug": 1, "rock": 1, "ghost": 1, "dark": 1,
			"normal": 1, "fighting": 1,
		},
		"dark": {
			"psychic": 2, "ghost": 2, "fighting": 0.5, "dark": 0.5, "fairy": 0.5,
			"normal": 1, "fire": 1, "water": 1, "grass": 1, "ice": 1, "electric": 1,
			"poison": 1, "ground": 1, "flying": 1, "rock": 1, "bug": 1, "steel": 1,
			"dragon": 1,
		},
		"steel": {
			"ice": 2, "rock": 2, "fairy": 2, "fire": 0.5, "water": 0.5,
			"grass": 0.5, "poison": 0, "flying": 1, "psychic": 1, "bug": 1,
			"ground": 1, "ghost": 1, "dragon": 1, "dark": 1, "steel": 0.5,
			"electric": 0.5, "normal": 1, "fighting": 1,
		},
		"fairy": {
			"fighting": 2, "dragon": 2, "dark": 2, "poison": 0.5, "steel": 0.5,
			"fire": 1, "water": 1, "grass": 1, "ice": 1, "electric": 1, "ghost": 1,
			"bug": 1, "rock": 1, "ground": 1, "flying": 1, "psychic": 1, "normal": 1,
		},
	}

	if defenderTypes, exists := matchups[attackType]; exists {
		if effectiveness, exists := defenderTypes[defendType]; exists {
			return effectiveness
		}
	}

	return 1.0 // 无克制关系
}

// ==================== 数据库接口 ====================

// Database 定义所需的数据库操作接口
type Database interface {
	GetEffortStats(userPokemonID int) (*model.PokemonEffortStats, error)
	UpdateEffortStats(stats *model.PokemonEffortStats) error
	GetNature(natureID string) (*model.PokemonNature, error)
	UpdatePokemonNature(userPokemonID int, natureID string) error
	GetPokemonSpecies(pokemonID string) (*model.PokemonSpecies, error)
}
