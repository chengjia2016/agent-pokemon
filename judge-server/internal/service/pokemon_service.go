package service

import (
	"database/sql"
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/model"
	"math/rand"
	"strings"
	"time"
)

// PokemonService 宝可梦服务
type PokemonService struct {
	db *db.Database
}

// NewPokemonService 创建宝可梦服务
func NewPokemonService(database *db.Database) *PokemonService {
	return &PokemonService{
		db: database,
	}
}

// InitializePokemonTables 初始化宝可梦相关表
func (s *PokemonService) InitializePokemonTables() error {
	// 创建宝可梦蛋表
	eggTableQuery := `
		CREATE TABLE IF NOT EXISTS pokemon_eggs (
			id VARCHAR(100) PRIMARY KEY,
			owner_id INTEGER NOT NULL,
			species VARCHAR(100),
			start_time TIMESTAMP DEFAULT NOW(),
			hatch_time TIMESTAMP,
			status VARCHAR(50) DEFAULT 'incubating',
			incubation_duration_seconds BIGINT DEFAULT 86400,
			energy_absorbed INTEGER DEFAULT 0,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)
	`

	if _, err := s.db.Exec(eggTableQuery); err != nil {
		return fmt.Errorf("failed to create pokemon_eggs table: %w", err)
	}

	// 创建用户宝可梦表（如果尚未存在）
	userPokemonTableQuery := `
		CREATE TABLE IF NOT EXISTS user_pokemons (
			id SERIAL PRIMARY KEY,
			github_id INTEGER NOT NULL,
			pet_id VARCHAR(100) NOT NULL UNIQUE,
			pet_name VARCHAR(255) NOT NULL,
			level INTEGER DEFAULT 1,
			species VARCHAR(100),
			created_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (github_id) REFERENCES users(github_id) ON DELETE CASCADE
		)
	`

	if _, err := s.db.Exec(userPokemonTableQuery); err != nil {
		// 表可能已存在，这是可以接受的
		// return fmt.Errorf("failed to create user_pokemons table: %w", err)
	}

	return nil
}

// CreateEgg 创建宝可梦蛋
func (s *PokemonService) CreateEgg(ownerID int, species string, incubationDuration int64) (*model.Egg, error) {
	eggID := fmt.Sprintf("egg_%s_%d", strings.ToLower(species), time.Now().UnixNano())

	query := `
		INSERT INTO pokemon_eggs (id, owner_id, species, start_time, hatch_time, status, incubation_duration_seconds, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW() + ($4 * interval '1 second'), 'incubating', $4, NOW(), NOW())
		RETURNING id, owner_id, species, start_time, hatch_time, status, incubation_duration_seconds, energy_absorbed, created_at, updated_at
	`

	egg := &model.Egg{
		ID:                 eggID,
		Owner:              fmt.Sprintf("%d", ownerID),
		Status:             "incubating",
		IncubationDuration: incubationDuration,
		EnergyAbsorbed:     0,
	}

	// Temporary variables for created_at and updated_at
	var createdAt, updatedAt time.Time
	var scannedOwnerID int

	err := s.db.QueryRow(query,
		eggID, ownerID, species, incubationDuration,
	).Scan(
		&egg.ID, &scannedOwnerID, &egg.Owner, &egg.StartTime, &egg.HatchTime,
		&egg.Status, &egg.IncubationDuration, &egg.EnergyAbsorbed, &createdAt, &updatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create egg: %w", err)
	}

	return egg, nil
}

// GetEgg 获取蛋信息
func (s *PokemonService) GetEgg(eggID string) (*model.Egg, error) {
	query := `
		SELECT id, owner_id, species, start_time, hatch_time, status, incubation_duration_seconds, energy_absorbed
		FROM pokemon_eggs
		WHERE id = $1
	`

	egg := &model.Egg{}
	var ownerID int

	err := s.db.QueryRow(query, eggID).Scan(
		&egg.ID, &ownerID, &egg.Owner, &egg.StartTime, &egg.HatchTime,
		&egg.Status, &egg.IncubationDuration, &egg.EnergyAbsorbed,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("egg not found")
		}
		return nil, fmt.Errorf("failed to get egg: %w", err)
	}

	egg.Owner = fmt.Sprintf("%d", ownerID)
	return egg, nil
}

// ListUserEggs 列出用户的蛋
func (s *PokemonService) ListUserEggs(ownerID int) ([]*model.Egg, error) {
	query := `
		SELECT id, owner_id, species, start_time, hatch_time, status, incubation_duration_seconds, energy_absorbed
		FROM pokemon_eggs
		WHERE owner_id = $1
		ORDER BY created_at DESC
	`

	rows, err := s.db.Query(query, ownerID)
	if err != nil {
		return nil, fmt.Errorf("failed to list eggs: %w", err)
	}
	defer rows.Close()

	var eggs []*model.Egg
	for rows.Next() {
		egg := &model.Egg{}
		err := rows.Scan(
			&egg.ID, &ownerID, &egg.Owner, &egg.StartTime, &egg.HatchTime,
			&egg.Status, &egg.IncubationDuration, &egg.EnergyAbsorbed,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan egg: %w", err)
		}
		egg.Owner = fmt.Sprintf("%d", ownerID)
		eggs = append(eggs, egg)
	}

	return eggs, nil
}

// HatchEgg 孵化蛋变成宝可梦
func (s *PokemonService) HatchEgg(eggID string, ownerID int) (*model.UserPokemon, error) {
	// 先获取蛋的信息
	egg, err := s.GetEgg(eggID)
	if err != nil {
		return nil, fmt.Errorf("failed to get egg: %w", err)
	}

	// 检查孵化时间是否到达
	if time.Now().Before(egg.HatchTime) {
		return nil, fmt.Errorf("egg is not ready to hatch yet. Ready at: %v", egg.HatchTime)
	}

	// 检查蛋状态
	if egg.Status != "incubating" {
		return nil, fmt.Errorf("egg has already been hatched or is invalid")
	}

	// 生成新的宝可梦
	petID := fmt.Sprintf("pet_%s_%d", strings.ToLower(egg.Owner), time.Now().UnixNano())
	petName := fmt.Sprintf("%s_%s", egg.Owner, time.Now().Format("20060102150405")[:8])

	// 创建用户宝可梦记录
	userPokemon := &model.UserPokemon{
		GithubID:  ownerID,
		PetID:     petID,
		PetName:   petName,
		Level:     1,
		Species:   egg.Owner,
		CreatedAt: time.Now(),
	}

	// 在事务中更新蛋状态并创建宝可梦
	insertPokemonQuery := `
		INSERT INTO user_pokemons (github_id, pet_id, pet_name, level, species, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
		RETURNING id
	`

	err = s.db.QueryRow(insertPokemonQuery,
		ownerID, petID, petName, 1, egg.Owner,
	).Scan(&userPokemon.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to create pokemon: %w", err)
	}

	// 更新蛋的状态为已孵化
	updateEggQuery := `
		UPDATE pokemon_eggs
		SET status = 'hatched', updated_at = NOW()
		WHERE id = $1
	`

	_, err = s.db.Exec(updateEggQuery, eggID)
	if err != nil {
		return nil, fmt.Errorf("failed to update egg status: %w", err)
	}

	return userPokemon, nil
}

// CreatePokemon 直接创建宝可梦
func (s *PokemonService) CreatePokemon(ownerID int, species string, level int) (*model.UserPokemon, error) {
	petID := fmt.Sprintf("pet_%d_%d", ownerID, time.Now().UnixNano())
	petName := fmt.Sprintf("%s_%d", species, time.Now().Unix())

	if level < 1 {
		level = 1
	}

	userPokemon := &model.UserPokemon{
		GithubID:  ownerID,
		PetID:     petID,
		PetName:   petName,
		Level:     level,
		Species:   species,
		CreatedAt: time.Now(),
	}

	query := `
		INSERT INTO user_pokemons (github_id, pet_id, pet_name, level, species, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
		RETURNING id
	`

	err := s.db.QueryRow(query,
		ownerID, petID, petName, level, species,
	).Scan(&userPokemon.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to create pokemon: %w", err)
	}

	return userPokemon, nil
}

// GetUserPokemon 获取用户的宝可梦
func (s *PokemonService) GetUserPokemon(petID string) (*model.UserPokemon, error) {
	query := `
		SELECT id, github_id, pet_id, pet_name, level, species, created_at
		FROM user_pokemons
		WHERE pet_id = $1
	`

	pokemon := &model.UserPokemon{}
	err := s.db.QueryRow(query, petID).Scan(
		&pokemon.ID, &pokemon.GithubID, &pokemon.PetID, &pokemon.PetName,
		&pokemon.Level, &pokemon.Species, &pokemon.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("pokemon not found")
		}
		return nil, fmt.Errorf("failed to get pokemon: %w", err)
	}

	return pokemon, nil
}

// ListUserPokemons 列出用户的所有宝可梦
func (s *PokemonService) ListUserPokemons(ownerID int) ([]*model.UserPokemon, error) {
	query := `
		SELECT id, github_id, pet_id, pet_name, level, species, created_at
		FROM user_pokemons
		WHERE github_id = $1
		ORDER BY created_at DESC
	`

	rows, err := s.db.Query(query, ownerID)
	if err != nil {
		return nil, fmt.Errorf("failed to list pokemons: %w", err)
	}
	defer rows.Close()

	var pokemons []*model.UserPokemon
	for rows.Next() {
		pokemon := &model.UserPokemon{}
		err := rows.Scan(
			&pokemon.ID, &pokemon.GithubID, &pokemon.PetID, &pokemon.PetName,
			&pokemon.Level, &pokemon.Species, &pokemon.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan pokemon: %w", err)
		}
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

// UpdatePokemonLevel 更新宝可梦等级
func (s *PokemonService) UpdatePokemonLevel(petID string, newLevel int) error {
	query := `
		UPDATE user_pokemons
		SET level = $1
		WHERE pet_id = $2
	`

	result, err := s.db.Exec(query, newLevel, petID)
	if err != nil {
		return fmt.Errorf("failed to update pokemon level: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("pokemon not found")
	}

	return nil
}

// GenerateWildPokemon 生成野生宝可梦
func (s *PokemonService) GenerateWildPokemon(townID string, level int) (*model.UserPokemon, error) {
	// 定义一些宝可梦物种
	species := []string{
		"Pikachu", "Charmander", "Squirtle", "Bulbasaur", "Pidgeot",
		"Raticate", "Spearow", "Jigglypuff", "Zubat", "Golbat",
		"Oddish", "Gloom", "Vileplume", "Paras", "Parasect",
		"Venonat", "Venomoth", "Diglett", "Dugtrio", "Meowth",
		"Persian", "Psyduck", "Golduck", "Mankey", "Primeape",
	}

	// 随机选择物种
	selectedSpecies := species[rand.Intn(len(species))]

	// 生成等级，范围在 1 到 town level+5
	minLevel := 1
	maxLevel := level + 5
	if maxLevel < minLevel {
		maxLevel = minLevel
	}
	actualLevel := minLevel + rand.Intn(maxLevel-minLevel+1)

	// 创建临时的宝可梦ID用于野生宝可梦记录
	wildPetID := fmt.Sprintf("wild_%s_%d", townID, time.Now().UnixNano())

	wildPokemon := &model.UserPokemon{
		PetID:     wildPetID,
		PetName:   fmt.Sprintf("%s_Wild", selectedSpecies),
		Level:     actualLevel,
		Species:   selectedSpecies,
		CreatedAt: time.Now(),
	}

	return wildPokemon, nil
}

// DeleteUserPokemon 删除用户的宝可梦
func (s *PokemonService) DeleteUserPokemon(petID string) error {
	query := `
		DELETE FROM user_pokemons
		WHERE pet_id = $1
	`

	result, err := s.db.Exec(query, petID)
	if err != nil {
		return fmt.Errorf("failed to delete pokemon: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("pokemon not found")
	}

	return nil
}

// DeleteEgg 删除蛋
func (s *PokemonService) DeleteEgg(eggID string) error {
	query := `
		DELETE FROM pokemon_eggs
		WHERE id = $1
	`

	result, err := s.db.Exec(query, eggID)
	if err != nil {
		return fmt.Errorf("failed to delete egg: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("egg not found")
	}

	return nil
}

// AbsorbEnergyToEgg 向蛋吸收能量（用于加速孵化）
func (s *PokemonService) AbsorbEnergyToEgg(eggID string, energyAmount int) (*model.Egg, error) {
	egg, err := s.GetEgg(eggID)
	if err != nil {
		return nil, fmt.Errorf("failed to get egg: %w", err)
	}

	// 更新能量
	newEnergyAbsorbed := egg.EnergyAbsorbed + energyAmount

	// 根据吸收的能量减少孵化时间（每100能量减少1小时）
	hoursToReduce := int64(energyAmount / 100)
	newHatchTime := egg.HatchTime.Add(-time.Duration(hoursToReduce) * time.Hour)

	// 确保孵化时间不会早于当前时间加上最小孵化时间
	minHatchTime := time.Now().Add(time.Hour * 1)
	if newHatchTime.Before(minHatchTime) {
		newHatchTime = minHatchTime
	}

	query := `
		UPDATE pokemon_eggs
		SET energy_absorbed = $1, hatch_time = $2, updated_at = NOW()
		WHERE id = $3
	`

	_, err = s.db.Exec(query, newEnergyAbsorbed, newHatchTime, eggID)
	if err != nil {
		return nil, fmt.Errorf("failed to update egg energy: %w", err)
	}

	egg.EnergyAbsorbed = newEnergyAbsorbed
	egg.HatchTime = newHatchTime
	return egg, nil
}
