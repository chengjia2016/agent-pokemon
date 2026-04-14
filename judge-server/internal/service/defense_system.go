package service

import (
	"errors"
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/model"
	"strings"
	"time"
)

// DefenseSystem 防守基地系统
type DefenseSystem struct {
	DB *db.Database
}

func NewDefenseSystem(database *db.Database) *DefenseSystem {
	return &DefenseSystem{DB: database}
}

// ==================== 基地创建 ====================

// CanCreateBase 检查用户是否可以创建基地（需要Fork）
func (ds *DefenseSystem) CanCreateBase(githubID int, repositoryURL string) (bool, error) {
	// 检查用户是否已有基地
	base, err := ds.DB.GetBase(githubID)
	if err == nil && base != nil {
		return false, errors.New("user already has a base")
	}

	// TODO: 检查是否已Fork chengjia2016/agent-monster
	// 这需要调用GitHub API验证
	if !strings.Contains(repositoryURL, "agent-monster") {
		return false, errors.New("repository must be a fork of agent-monster")
	}

	return true, nil
}

// CreateBase 创建防守基地
func (ds *DefenseSystem) CreateBase(githubID int, repositoryURL string) (*model.UserBase, error) {
	// 检查是否可以创建基地
	canCreate, err := ds.CanCreateBase(githubID, repositoryURL)
	if err != nil {
		return nil, err
	}

	if !canCreate {
		return nil, errors.New("cannot create base")
	}

	// 获取用户最强的3只宠物
	strongestPokemons, err := ds.DB.GetStrongestPokemon(githubID, 3)
	if err != nil {
		return nil, fmt.Errorf("failed to get strongest pokemons: %w", err)
	}

	if len(strongestPokemons) < 1 {
		return nil, fmt.Errorf("user needs at least 1 pokemon to create a base (has %d)", len(strongestPokemons))
	}

	// 创建防守队伍
	defenseTeamID := fmt.Sprintf("defense_team_%d_%d", githubID, time.Now().Unix())
	defenseTeam, err := ds.DB.CreateTeam(githubID, defenseTeamID, "Defense Team", "Automatic defense team", true)
	if err != nil {
		return nil, fmt.Errorf("failed to create defense team: %w", err)
	}

	// 将最强的3只宠物添加到防守队伍
	for i, pokemon := range strongestPokemons {
		if i >= 3 {
			break
		}

		err = ds.DB.AddPetToTeam(defenseTeam.TeamID, pokemon.PetID, i+1)
		if err != nil {
			return nil, fmt.Errorf("failed to add pokemon to defense team: %w", err)
		}
	}

	// 创建基地
	baseID := fmt.Sprintf("base_%d_%d", githubID, time.Now().Unix())
	base, err := ds.DB.CreateBase(githubID, baseID, repositoryURL, defenseTeam.TeamID)
	if err != nil {
		return nil, fmt.Errorf("failed to create base: %w", err)
	}

	return base, nil
}

// ==================== 防守验证 ====================

// ValidateDefenseAttack 验证防守攻击请求
func (ds *DefenseSystem) ValidateDefenseAttack(attackerGitHubID int, base *model.UserBase) error {
	if base == nil {
		return errors.New("base not found")
	}

	// 不能攻击自己的基地
	if attackerGitHubID == base.GitHubID {
		return errors.New("cannot attack your own base")
	}

	// 检查基地防守队伍是否存在
	if base.DefenseTeamID == nil || *base.DefenseTeamID == "" {
		return errors.New("base has no defense team")
	}

	// 获取防守队伍
	defenseMembers, err := ds.DB.GetTeamMembers(*base.DefenseTeamID)
	if err != nil {
		return fmt.Errorf("failed to get defense team: %w", err)
	}

	if len(defenseMembers) < 1 {
		return fmt.Errorf("defense team must have at least 1 member, has %d", len(defenseMembers))
	}

	return nil
}

// ==================== 防守更新 ====================

// UpdateDefenseTeam 更新防守队伍（最强的3只宠物）
func (ds *DefenseSystem) UpdateDefenseTeam(githubID int, base *model.UserBase) error {
	// 获取用户最强的3只宠物
	strongestPokemons, err := ds.DB.GetStrongestPokemon(githubID, 3)
	if err != nil {
		return fmt.Errorf("failed to get strongest pokemons: %w", err)
	}

	if len(strongestPokemons) < 1 {
		return fmt.Errorf("user needs at least 1 pokemon (has %d)", len(strongestPokemons))
	}

	// 获取当前防守队伍成员
	currentMembers, err := ds.DB.GetTeamMembers(*base.DefenseTeamID)
	if err != nil {
		return fmt.Errorf("failed to get current defense team: %w", err)
	}

	// 比较是否需要更新
	needsUpdate := false
	if len(currentMembers) != len(strongestPokemons) {
		needsUpdate = true
	} else {
		// 检查宠物是否相同
		for i, member := range currentMembers {
			if i < len(strongestPokemons) && member.PetID != strongestPokemons[i].PetID {
				needsUpdate = true
				break
			}
		}
	}

	if !needsUpdate {
		return nil
	}

	// TODO: 实现更新队伍成员的逻辑
	// 需要删除旧成员并添加新成员
	return nil
}

// ==================== 防守记录 ====================

// GetDefenseHistory 获取基地的防守历史
func (ds *DefenseSystem) GetDefenseHistory(baseID string, limit int) ([]model.BaseDefenseRecord, error) {
	// TODO: 实现从数据库查询防守历史
	return []model.BaseDefenseRecord{}, errors.New("method not implemented")
}

// GetDefenseStats 获取基地防守统计
func (ds *DefenseSystem) GetDefenseStats(base *model.UserBase) map[string]interface{} {
	stats := map[string]interface{}{
		"defense_wins":   base.DefenseWins,
		"defense_losses": base.DefenseLosses,
		"prestige":       base.Prestige,
		"level":          base.Level,
		"win_rate":       0.0,
	}

	totalBattles := base.DefenseWins + base.DefenseLosses
	if totalBattles > 0 {
		stats["win_rate"] = float64(base.DefenseWins) / float64(totalBattles)
	}

	return stats
}

// ==================== 辅助函数 ====================

// GetDefenseTeamPokemons 获取防守队伍的宠物
func (ds *DefenseSystem) GetDefenseTeamPokemons(teamID string) ([]model.UserTeamMember, error) {
	members, err := ds.DB.GetTeamMembers(teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get defense team members: %w", err)
	}

	if len(members) < 1 {
		return nil, fmt.Errorf("defense team must have at least 1 member, has %d", len(members))
	}

	return members, nil
}

// CalculateBaseLevel 计算基地等级
func (ds *DefenseSystem) CalculateBaseLevel(prestige int) int {
	// 等级 = prestige / 10 + 1
	return prestige/10 + 1
}
