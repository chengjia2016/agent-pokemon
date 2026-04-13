package service

import (
	"errors"
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/model"
	"time"
)

// PokemonManager 宠物管理器
type PokemonManager struct {
	DB *db.Database
}

func NewPokemonManager(database *db.Database) *PokemonManager {
	return &PokemonManager{DB: database}
}

// ==================== 宠物所有权 ====================

// CanAddPokemon 检查是否可以添加宠物
func (pm *PokemonManager) CanAddPokemon(githubID int) (bool, error) {
	count, err := pm.DB.CountUserPokemon(githubID)
	if err != nil {
		return false, err
	}

	return count < 10, nil
}

// AddPokemon 添加宠物到用户账户
func (pm *PokemonManager) AddPokemon(githubID int, petID string, isCaptured bool, maxHP int) (*model.UserOwnedPokemon, error) {
	// 检查是否可以添加
	can, err := pm.CanAddPokemon(githubID)
	if err != nil {
		return nil, err
	}

	if !can {
		return nil, errors.New("user already owns 10 pokemon (maximum limit)")
	}

	// 添加宠物
	pokemon, err := pm.DB.AddOwnedPokemon(githubID, petID, isCaptured, maxHP)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

// GetUserPokemons 获取用户所有宠物
func (pm *PokemonManager) GetUserPokemons(githubID int) ([]model.UserOwnedPokemon, error) {
	pokemons, err := pm.DB.GetUserOwnedPokemons(githubID)
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

// GetPokemonCount 获取用户宠物数量
func (pm *PokemonManager) GetPokemonCount(githubID int) (int, error) {
	count, err := pm.DB.CountUserPokemon(githubID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// ==================== 宠物替换 ====================

// GetReplacementCandidates 获取可被替换的宠物列表
func (pm *PokemonManager) GetReplacementCandidates(githubID int) ([]model.UserOwnedPokemon, error) {
	pokemons, err := pm.GetUserPokemons(githubID)
	if err != nil {
		return nil, err
	}

	// 返回所有宠物，让用户选择
	return pokemons, nil
}

// ReplacePokemon 替换宠物
func (pm *PokemonManager) ReplacePokemon(githubID int, oldPetID string, newPetID string, isCaptured bool, maxHP int) error {
	// 检查旧宠物是否存在并属于该用户
	pokemons, err := pm.GetUserPokemons(githubID)
	if err != nil {
		return err
	}

	found := false
	for _, p := range pokemons {
		if p.PetID == oldPetID {
			found = true
			break
		}
	}

	if !found {
		return errors.New("pokemon not found in user's collection")
	}

	// TODO: 实现删除旧宠物和添加新宠物的逻辑
	// 这需要在数据库层添加删除宠物的方法

	return nil
}

// ==================== HP恢复 ====================

// GetRecoveryStatus 获取宠物恢复状态
func (pm *PokemonManager) GetRecoveryStatus(githubID int) (map[string]interface{}, error) {
	pokemons, err := pm.GetUserPokemons(githubID)
	if err != nil {
		return nil, err
	}

	status := make(map[string]interface{})
	faintedPokemons := []map[string]interface{}{}
	healthyPokemons := []map[string]interface{}{}
	recoveryCount := 0

	now := time.Now()
	for _, pokemon := range pokemons {
		info := map[string]interface{}{
			"pet_id": pokemon.PetID,
			"status": pokemon.Status,
			"hp":     pokemon.CurrentHP,
			"max_hp": pokemon.MaxHP,
			"level":  pokemon.Level,
		}

		if pokemon.Status == model.PokemonStatusFainted {
			// 计算恢复时间
			recoveryTime := pokemon.UpdatedAt.Add(1 * time.Hour)
			if recoveryTime.Before(now) {
				recoveryCount++
				info["recovery_status"] = "ready"
			} else {
				remainingMinutes := int(time.Until(recoveryTime).Minutes())
				info["recovery_status"] = "recovering"
				info["remaining_minutes"] = remainingMinutes
			}
			faintedPokemons = append(faintedPokemons, info)
		} else {
			healthyPokemons = append(healthyPokemons, info)
		}
	}

	status["healthy_pokemons"] = healthyPokemons
	status["fainted_pokemons"] = faintedPokemons
	status["ready_for_recovery"] = recoveryCount
	status["total_fainted"] = len(faintedPokemons)

	return status, nil
}

// ==================== 宠物状态 ====================

// UpdatePokemonLevel 更新宠物等级
func (pm *PokemonManager) UpdatePokemonLevel(petID string, newLevel int) error {
	// TODO: 实现宠物等级更新
	return errors.New("method not implemented")
}

// IsPokemonHealthy 检查宠物是否健康（满HP）
func (pm *PokemonManager) IsPokemonHealthy(pokemon *model.UserOwnedPokemon) bool {
	return pokemon.Status == model.PokemonStatusActive && pokemon.CurrentHP == pokemon.MaxHP
}

// CountHealthyPokemons 计算健康宠物数
func (pm *PokemonManager) CountHealthyPokemons(pokemons []model.UserOwnedPokemon) int {
	count := 0
	for _, p := range pokemons {
		if pm.IsPokemonHealthy(&p) {
			count++
		}
	}
	return count
}

// ==================== 宠物训练 ====================

// TrainPokemon 训练宠物
func (pm *PokemonManager) TrainPokemon(githubID int, petID string, trainingType string) error {
	// TODO: 实现宠物训练逻辑
	// 增加经验值或属性

	switch trainingType {
	case "strength":
		// 提升攻击力
	case "defense":
		// 提升防守力
	case "speed":
		// 提升速度
	default:
		return errors.New("invalid training type")
	}

	return nil
}

// ==================== 宠物释放 ====================

// ReleasePokemon 释放宠物
func (pm *PokemonManager) ReleasePokemon(githubID int, petID string) error {
	// TODO: 实现删除宠物的数据库操作
	// 可以标记为已释放或直接删除

	return errors.New("method not implemented")
}

// ==================== 辅助函数 ====================

// GetWeakestPokemon 获取最弱的宠物（用于替换）
func (pm *PokemonManager) GetWeakestPokemon(pokemons []model.UserOwnedPokemon) *model.UserOwnedPokemon {
	if len(pokemons) == 0 {
		return nil
	}

	weakest := &pokemons[0]
	for i := 1; i < len(pokemons); i++ {
		if pokemons[i].Level < weakest.Level ||
			(pokemons[i].Level == weakest.Level && pokemons[i].Experience < weakest.Experience) {
			weakest = &pokemons[i]
		}
	}

	return weakest
}

// SortPokemonsByStrength 按实力排序宠物
func (pm *PokemonManager) SortPokemonsByStrength(pokemons []model.UserOwnedPokemon) {
	// 简单的冒泡排序
	for i := 0; i < len(pokemons); i++ {
		for j := 0; j < len(pokemons)-1-i; j++ {
			if pokemons[j].Level < pokemons[j+1].Level {
				pokemons[j], pokemons[j+1] = pokemons[j+1], pokemons[j]
			}
		}
	}
}

// ValidatePokemonTeam 验证宠物队伍
func (pm *PokemonManager) ValidatePokemonTeam(pokemons []model.UserOwnedPokemon) error {
	if len(pokemons) < 1 {
		return fmt.Errorf("team must have at least 1 pokemon, has %d", len(pokemons))
	}

	for i, p := range pokemons {
		if p.Status != model.PokemonStatusActive {
			return fmt.Errorf("pokemon %d is not active (status: %s)", i+1, p.Status)
		}
		if p.CurrentHP <= 0 {
			return fmt.Errorf("pokemon %d has no HP", i+1)
		}
	}

	return nil
}
