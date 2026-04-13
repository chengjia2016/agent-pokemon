package db

import (
	"database/sql"
	"errors"
	"fmt"
	"judge-server/internal/model"
	"time"
)

// ==================== 队伍管理操作 ====================

// CreateTeam 创建用户的战斗队伍
func (d *Database) CreateTeam(githubID int, teamID, teamName, description string, isDefenseTeam bool) (*model.UserTeam, error) {
	team := &model.UserTeam{
		GitHubID:      githubID,
		TeamID:        teamID,
		TeamName:      teamName,
		Description:   description,
		IsDefenseTeam: isDefenseTeam,
		MaxMembers:    3,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	query := `INSERT INTO user_teams (github_id, team_id, team_name, description, is_defense_team, max_members, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := d.conn.QueryRow(query, githubID, teamID, teamName, description, isDefenseTeam, 3, time.Now(), time.Now()).Scan(&team.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create team: %w", err)
	}

	return team, nil
}

// AddPetToTeam 将宠物加入队伍
func (d *Database) AddPetToTeam(teamID, petID string, slotPosition int) error {
	if slotPosition < 1 || slotPosition > 3 {
		return errors.New("slot position must be between 1 and 3")
	}

	query := `INSERT INTO user_team_members (team_id, pet_id, slot_position, joined_at)
	          VALUES ($1, $2, $3, $4)`

	_, err := d.conn.Exec(query, teamID, petID, slotPosition, time.Now())
	if err != nil {
		return fmt.Errorf("failed to add pet to team: %w", err)
	}

	return nil
}

// GetTeamMembers 获取队伍的所有成员
func (d *Database) GetTeamMembers(teamID string) ([]model.UserTeamMember, error) {
	query := `SELECT id, team_id, pet_id, slot_position, joined_at 
	          FROM user_team_members WHERE team_id = $1 ORDER BY slot_position`

	rows, err := d.conn.Query(query, teamID)
	if err != nil {
		return nil, fmt.Errorf("failed to get team members: %w", err)
	}
	defer rows.Close()

	var members []model.UserTeamMember
	for rows.Next() {
		var member model.UserTeamMember
		err := rows.Scan(&member.ID, &member.TeamID, &member.PetID, &member.SlotPosition, &member.JoinedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan team member: %w", err)
		}
		members = append(members, member)
	}

	return members, nil
}

// ==================== 用户宠物所有权 ====================

// CountUserPokemon 统计用户拥有的宠物数量
func (d *Database) CountUserPokemon(githubID int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM user_owned_pokemon WHERE github_id = $1`
	err := d.conn.QueryRow(query, githubID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count user pokemon: %w", err)
	}
	return count, nil
}

// AddOwnedPokemon 添加用户的宠物
func (d *Database) AddOwnedPokemon(githubID int, petID string, isCaptured bool, maxHP int) (*model.UserOwnedPokemon, error) {
	// 检查是否已达到10个宠物限制
	count, err := d.CountUserPokemon(githubID)
	if err != nil {
		return nil, err
	}
	if count >= 10 {
		return nil, errors.New("user already owns 10 pokemon (maximum limit)")
	}

	owned := &model.UserOwnedPokemon{
		GitHubID:   githubID,
		PetID:      petID,
		IsCaptured: isCaptured,
		Status:     model.PokemonStatusActive,
		CurrentHP:  maxHP,
		MaxHP:      maxHP,
		Level:      1,
		Experience: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if isCaptured {
		now := time.Now()
		owned.CapturedAt = &now
	}

	query := `INSERT INTO user_owned_pokemon (github_id, pet_id, is_captured, status, current_hp, max_hp, level, experience, captured_at, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`

	err = d.conn.QueryRow(query, githubID, petID, isCaptured, owned.Status, owned.CurrentHP, owned.MaxHP, owned.Level, owned.Experience, owned.CapturedAt, owned.CreatedAt, owned.UpdatedAt).Scan(&owned.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to add owned pokemon: %w", err)
	}

	return owned, nil
}

// GetUserOwnedPokemons 获取用户拥有的宠物列表
func (d *Database) GetUserOwnedPokemons(githubID int) ([]model.UserOwnedPokemon, error) {
	query := `SELECT id, github_id, pet_id, is_captured, status, current_hp, max_hp, level, experience, captured_at, created_at, updated_at
	          FROM user_owned_pokemon WHERE github_id = $1 ORDER BY level DESC, created_at ASC`

	rows, err := d.conn.Query(query, githubID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user pokemon: %w", err)
	}
	defer rows.Close()

	var pokemons []model.UserOwnedPokemon
	for rows.Next() {
		var pokemon model.UserOwnedPokemon
		err := rows.Scan(&pokemon.ID, &pokemon.GitHubID, &pokemon.PetID, &pokemon.IsCaptured, &pokemon.Status,
			&pokemon.CurrentHP, &pokemon.MaxHP, &pokemon.Level, &pokemon.Experience, &pokemon.CapturedAt, &pokemon.CreatedAt, &pokemon.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user pokemon: %w", err)
		}
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

// UpdatePokemonHP 更新宠物HP
func (d *Database) UpdatePokemonHP(petID string, currentHP int) error {
	query := `UPDATE user_owned_pokemon SET current_hp = $1, updated_at = $2 WHERE pet_id = $3`
	_, err := d.conn.Exec(query, currentHP, time.Now(), petID)
	if err != nil {
		return fmt.Errorf("failed to update pokemon hp: %w", err)
	}
	return nil
}

// UpdatePokemonStatus 更新宠物状态
func (d *Database) UpdatePokemonStatus(petID, status string) error {
	query := `UPDATE user_owned_pokemon SET status = $1, updated_at = $2 WHERE pet_id = $3`
	_, err := d.conn.Exec(query, status, time.Now(), petID)
	if err != nil {
		return fmt.Errorf("failed to update pokemon status: %w", err)
	}
	return nil
}

// ==================== 防守基地操作 ====================

// CreateBase 创建用户的防守基地
func (d *Database) CreateBase(githubID int, baseID, repositoryURL, defenseTeamID string) (*model.UserBase, error) {
	var defenseTeamIDPtr *string
	if defenseTeamID != "" {
		defenseTeamIDPtr = &defenseTeamID
	}

	base := &model.UserBase{
		GitHubID:      githubID,
		BaseID:        baseID,
		RepositoryURL: repositoryURL,
		DefenseTeamID: defenseTeamIDPtr,
		Level:         1,
		Prestige:      0,
		DefenseWins:   0,
		DefenseLosses: 0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	query := `INSERT INTO user_bases (github_id, base_id, repository_url, defense_team_id, level, prestige, defense_wins, defense_losses, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	err := d.conn.QueryRow(query, githubID, baseID, repositoryURL, defenseTeamID, 1, 0, 0, 0, time.Now(), time.Now()).Scan(&base.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create base: %w", err)
	}

	return base, nil
}

// GetBase 获取用户的防守基地
func (d *Database) GetBase(githubID int) (*model.UserBase, error) {
	base := &model.UserBase{}
	query := `SELECT id, github_id, base_id, repository_url, defense_team_id, level, prestige, defense_wins, defense_losses, created_at, updated_at
	          FROM user_bases WHERE github_id = $1`

	err := d.conn.QueryRow(query, githubID).Scan(&base.ID, &base.GitHubID, &base.BaseID, &base.RepositoryURL, &base.DefenseTeamID,
		&base.Level, &base.Prestige, &base.DefenseWins, &base.DefenseLosses, &base.CreatedAt, &base.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("base not found")
		}
		return nil, fmt.Errorf("failed to get base: %w", err)
	}

	return base, nil
}

// RecordDefense 记录一次防守战斗
func (d *Database) RecordDefense(baseID string, attackerGitHubID int, battleID, result string, attackerLevel, defenderLevel int, coinsGained float64) error {
	query := `INSERT INTO base_defense_records (base_id, attacker_github_id, battle_id, result, attacker_level, defender_level, coins_gained, timestamp)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := d.conn.Exec(query, baseID, attackerGitHubID, battleID, result, attackerLevel, defenderLevel, coinsGained, time.Now())
	if err != nil {
		return fmt.Errorf("failed to record defense: %w", err)
	}

	// 更新基地统计
	if result == model.DefenseResultWin {
		d.conn.Exec(`UPDATE user_bases SET defense_wins = defense_wins + 1, prestige = prestige + 10, updated_at = $1 WHERE base_id = $2`, time.Now(), baseID)
	} else {
		d.conn.Exec(`UPDATE user_bases SET defense_losses = defense_losses + 1, updated_at = $1 WHERE base_id = $2`, time.Now(), baseID)
	}

	return nil
}

// ==================== 野生精灵操作 ====================

// CreateWildPokemon 在地图上创建野生精灵
func (d *Database) CreateWildPokemon(wildID, locationID, speciesID string, level, maxHP, captureDifficulty int) (*model.WildPokemon, error) {
	wild := &model.WildPokemon{
		WildID:            wildID,
		LocationID:        locationID,
		PokemonSpeciesID:  speciesID,
		Level:             level,
		Status:            model.WildStatusActive,
		CurrentHP:         maxHP,
		MaxHP:             maxHP,
		CaptureDifficulty: captureDifficulty,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	query := `INSERT INTO wild_pokemon (wild_id, location_id, pokemon_species_id, level, status, current_hp, max_hp, capture_difficulty, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	err := d.conn.QueryRow(query, wildID, locationID, speciesID, level, wild.Status, wild.CurrentHP, wild.MaxHP, captureDifficulty, time.Now(), time.Now()).Scan(&wild.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create wild pokemon: %w", err)
	}

	return wild, nil
}

// GetWildPokemon 获取野生精灵
func (d *Database) GetWildPokemon(wildID string) (*model.WildPokemon, error) {
	wild := &model.WildPokemon{}
	query := `SELECT id, wild_id, location_id, pokemon_species_id, level, status, current_hp, max_hp, capture_difficulty, 
	                 caught_by_github_id, defeated_by_github_id, created_at, captured_at, defeated_at, updated_at
	          FROM wild_pokemon WHERE wild_id = $1`

	err := d.conn.QueryRow(query, wildID).Scan(&wild.ID, &wild.WildID, &wild.LocationID, &wild.PokemonSpeciesID, &wild.Level,
		&wild.Status, &wild.CurrentHP, &wild.MaxHP, &wild.CaptureDifficulty, &wild.CaughtByGitHubID, &wild.DefeatedByGitHubID,
		&wild.CreatedAt, &wild.CapturedAt, &wild.DefeatedAt, &wild.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("wild pokemon not found")
		}
		return nil, fmt.Errorf("failed to get wild pokemon: %w", err)
	}

	return wild, nil
}

// CaptureWildPokemon 捕获野生精灵
func (d *Database) CaptureWildPokemon(wildID string, githubID int, petID string, success bool) error {
	query := `UPDATE wild_pokemon SET status = $1, caught_by_github_id = $2, captured_at = $3, updated_at = $4 WHERE wild_id = $5`

	status := model.WildStatusCaptured
	if !success {
		status = model.WildStatusDefeated
	}

	_, err := d.conn.Exec(query, status, githubID, time.Now(), time.Now(), wildID)
	if err != nil {
		return fmt.Errorf("failed to capture wild pokemon: %w", err)
	}

	// 记录捕获历史
	if success {
		captureQuery := `INSERT INTO capture_history (github_id, pet_id, wild_pokemon_id, capture_type, success, attempt_count, captured_at)
		                VALUES ($1, $2, $3, $4, $5, $6, $7)`
		_, err = d.conn.Exec(captureQuery, githubID, petID, wildID, model.CaptureTypeWild, true, 1, time.Now())
		if err != nil {
			return fmt.Errorf("failed to record capture history: %w", err)
		}
	}

	return nil
}

// ==================== 战斗操作 ====================

// CreateBattle 创建战斗记录
func (d *Database) CreateBattle(battleID string, attackerID, defenderID int, attackerTeamID, defenderTeamID, battleType string) (*model.BattleV2, error) {
	battle := &model.BattleV2{
		BattleID:         battleID,
		AttackerGitHubID: attackerID,
		DefenderGitHubID: defenderID,
		AttackerTeamID:   attackerTeamID,
		DefenderTeamID:   defenderTeamID,
		BattleType:       battleType,
		Status:           model.BattleStatusOngoing,
		StartedAt:        time.Now(),
	}

	query := `INSERT INTO battles_v2 (battle_id, attacker_github_id, defender_github_id, attacker_team_id, defender_team_id, battle_type, status, started_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := d.conn.QueryRow(query, battleID, attackerID, defenderID, attackerTeamID, defenderTeamID, battleType, model.BattleStatusOngoing, time.Now()).Scan(&battle.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create battle: %w", err)
	}

	return battle, nil
}

// ==================== 辅助函数 ====================

// GetStrongestPokemon 获取用户最强的3只宠物（用于防守队伍）
func (d *Database) GetStrongestPokemon(githubID int, limit int) ([]model.UserOwnedPokemon, error) {
	if limit > 3 {
		limit = 3
	}

	query := `SELECT id, github_id, pet_id, is_captured, status, current_hp, max_hp, level, experience, captured_at, created_at, updated_at
	          FROM user_owned_pokemon WHERE github_id = $1 AND status = $2
	          ORDER BY level DESC, experience DESC LIMIT $3`

	rows, err := d.conn.Query(query, githubID, model.PokemonStatusActive, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get strongest pokemon: %w", err)
	}
	defer rows.Close()

	var pokemons []model.UserOwnedPokemon
	for rows.Next() {
		var pokemon model.UserOwnedPokemon
		err := rows.Scan(&pokemon.ID, &pokemon.GitHubID, &pokemon.PetID, &pokemon.IsCaptured, &pokemon.Status,
			&pokemon.CurrentHP, &pokemon.MaxHP, &pokemon.Level, &pokemon.Experience, &pokemon.CapturedAt, &pokemon.CreatedAt, &pokemon.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan pokemon: %w", err)
		}
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}
