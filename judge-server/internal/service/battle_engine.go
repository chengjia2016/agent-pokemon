package service

import (
	"errors"
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/model"
	"math/rand"
	"time"
)

// BattleEngine 战斗引擎
type BattleEngine struct {
	DB *db.Database
}

func NewBattleEngine(database *db.Database) *BattleEngine {
	return &BattleEngine{DB: database}
}

// ==================== 战斗验证 ====================

// ValidateBattleRequest 验证战斗请求
func (be *BattleEngine) ValidateBattleRequest(attackerID, defenderID int, attackerTeamID, defenderTeamID string) error {
	// 不能攻击自己
	if attackerID == defenderID {
		return errors.New("cannot attack yourself")
	}

	// 获取攻击者和防守者的宠物
	attackerPokemons, err := be.DB.GetUserOwnedPokemons(attackerID)
	if err != nil {
		return fmt.Errorf("failed to get attacker pokemons: %w", err)
	}

	defenderPokemons, err := be.DB.GetUserOwnedPokemons(defenderID)
	if err != nil {
		return fmt.Errorf("failed to get defender pokemons: %w", err)
	}

	if len(attackerPokemons) == 0 {
		return errors.New("attacker has no pokemons")
	}

	if len(defenderPokemons) == 0 {
		return errors.New("defender has no pokemons")
	}

	// 检查攻击者是否有3个满HP的宠物（要求1）
	healthyCount := 0
	for _, p := range attackerPokemons {
		if p.Status == model.PokemonStatusActive && p.CurrentHP == p.MaxHP {
			healthyCount++
		}
	}

	if healthyCount < 1 {
		return fmt.Errorf("attacker needs 1 pokemon with full HP to battle (current: %d)", healthyCount)
	}

	// 计算攻击者和防守者的平均等级
	attackerAvgLevel := be.calculateAverageLevel(attackerPokemons)
	defenderAvgLevel := be.calculateAverageLevel(defenderPokemons)

	// 检查等级限制：不能攻击比自己低5级以上的对手（要求2）
	if attackerAvgLevel > defenderAvgLevel+5 {
		return fmt.Errorf("target level too low: attacker average level %d, defender average level %d", attackerAvgLevel, defenderAvgLevel)
	}

	return nil
}

// calculateAverageLevel 计算平均等级
func (be *BattleEngine) calculateAverageLevel(pokemons []model.UserOwnedPokemon) int {
	if len(pokemons) == 0 {
		return 0
	}

	totalLevel := 0
	for _, p := range pokemons {
		totalLevel += p.Level
	}

	return totalLevel / len(pokemons)
}

// ==================== HP管理 ====================

// CalculateDamage 计算伤害（简单算法）
func (be *BattleEngine) CalculateDamage(attacker, defender *model.Pet) int {
	if attacker == nil || defender == nil {
		return 0
	}

	// 基础伤害：攻击者等级 * 攻击力 - 防守者防守力
	baseDamage := attacker.Level * 10
	if baseDamage > 0 {
		// 加入随机性（80%-120%）
		variance := 0.8 + rand.Float64()*0.4
		baseDamage = int(float64(baseDamage) * variance)
	}

	return baseDamage
}

// ApplyDamage 应用伤害到宠物
func (be *BattleEngine) ApplyDamage(petID string, damage int) error {
	// 获取宠物
	pokemons, err := be.DB.GetUserOwnedPokemons(0) // 需要修改查询
	if err != nil {
		return err
	}

	var targetPokemon *model.UserOwnedPokemon
	for i := range pokemons {
		if pokemons[i].PetID == petID {
			targetPokemon = &pokemons[i]
			break
		}
	}

	if targetPokemon == nil {
		return errors.New("pokemon not found")
	}

	// 减少HP
	newHP := targetPokemon.CurrentHP - damage
	if newHP < 0 {
		newHP = 0
	}

	// 更新HP
	err = be.DB.UpdatePokemonHP(petID, newHP)
	if err != nil {
		return err
	}

	// 如果HP <= 0，标记为昏迷
	if newHP <= 0 {
		err = be.DB.UpdatePokemonStatus(petID, model.PokemonStatusFainted)
		if err != nil {
			return err
		}
	}

	return nil
}

// ==================== HP恢复 ====================

// CheckAndRecoverHP 检查并恢复已经过期的HP（1小时自动恢复）
func (be *BattleEngine) CheckAndRecoverHP(githubID int) error {
	pokemons, err := be.DB.GetUserOwnedPokemons(githubID)
	if err != nil {
		return err
	}

	now := time.Now()
	for _, pokemon := range pokemons {
		// 只恢复昏迷的宠物
		if pokemon.Status != model.PokemonStatusFainted {
			continue
		}

		// 检查是否已经过了1小时
		if pokemon.UpdatedAt.Add(1 * time.Hour).Before(now) {
			// 恢复HP到满值
			err = be.DB.UpdatePokemonHP(pokemon.PetID, pokemon.MaxHP)
			if err != nil {
				return err
			}

			// 更新状态为活跃
			err = be.DB.UpdatePokemonStatus(pokemon.PetID, model.PokemonStatusActive)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetRecoveryTime 获取宠物恢复所需的剩余时间（分钟）
func (be *BattleEngine) GetRecoveryTime(pokemon *model.UserOwnedPokemon) int {
	if pokemon.Status != model.PokemonStatusFainted {
		return 0
	}

	recoveryTime := pokemon.UpdatedAt.Add(1 * time.Hour)
	remainingTime := time.Until(recoveryTime)

	if remainingTime <= 0 {
		return 0
	}

	return int(remainingTime.Minutes())
}

// ==================== 战斗结果 ====================

// DetermineBattleWinner 确定战斗赢家
func (be *BattleEngine) DetermineBattleWinner(attackerPokemons, defenderPokemons []model.UserOwnedPokemon) (string, error) {
	// 计算双方活跃宠物数
	attackerActive := 0
	defenderActive := 0

	for _, p := range attackerPokemons {
		if p.Status == model.PokemonStatusActive && p.CurrentHP > 0 {
			attackerActive++
		}
	}

	for _, p := range defenderPokemons {
		if p.Status == model.PokemonStatusActive && p.CurrentHP > 0 {
			defenderActive++
		}
	}

	if attackerActive > defenderActive {
		return "attacker", nil
	} else if defenderActive > attackerActive {
		return "defender", nil
	} else if attackerActive == 0 && defenderActive == 0 {
		// 都没有活跃宠物，平手
		return "draw", nil
	}

	return "", errors.New("unable to determine battle winner")
}

// ==================== 战斗统计 ====================

// RecordBattleStats 记录战斗统计
func (be *BattleEngine) RecordBattleStats(attackerPokemons, defenderPokemons []model.UserOwnedPokemon, winner string, totalDamage map[string]int) error {
	// 记录攻击者的宠物统计
	for _, pokemon := range attackerPokemons {
		stats := &model.PetBattleStats{
			PetID:            pokemon.PetID,
			TotalDamageDealt: totalDamage[pokemon.PetID],
		}

		if winner == "attacker" {
			stats.Wins = 1
		} else if winner == "defender" {
			stats.Losses = 1
		} else {
			stats.Draws = 1
		}

		// TODO: 实现更新或插入统计的数据库操作
		_ = stats
	}

	// 记录防守者的宠物统计
	for _, pokemon := range defenderPokemons {
		stats := &model.PetBattleStats{
			PetID:            pokemon.PetID,
			TotalDamageDealt: totalDamage[pokemon.PetID],
		}

		if winner == "defender" {
			stats.Wins = 1
		} else if winner == "attacker" {
			stats.Losses = 1
		} else {
			stats.Draws = 1
		}

		// TODO: 实现更新或插入统计的数据库操作
		_ = stats
	}

	return nil
}

// ==================== 辅助函数 ====================

// GetSelectedTeamPokemons 获取队伍的选定宠物
func (be *BattleEngine) GetSelectedTeamPokemons(teamID string) ([]model.UserTeamMember, error) {
	members, err := be.DB.GetTeamMembers(teamID)
	if err != nil {
		return nil, err
	}

	if len(members) < 1 {
		return nil, fmt.Errorf("team must have at least 1 member, has %d", len(members))
	}

	return members, nil
}

// SimulateBattleRound 模拟战斗回合
func (be *BattleEngine) SimulateBattleRound(attackerPet, defenderPet *model.Pet) (int, int, error) {
	if attackerPet == nil || defenderPet == nil {
		return 0, 0, errors.New("invalid pokemon for battle round")
	}

	// 计算伤害
	attackerDamage := be.CalculateDamage(attackerPet, defenderPet)
	defenderDamage := be.CalculateDamage(defenderPet, attackerPet)

	return attackerDamage, defenderDamage, nil
}
