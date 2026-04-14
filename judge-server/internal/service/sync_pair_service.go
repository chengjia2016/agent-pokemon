package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// SyncPairService handles all Sync Pair related operations
type SyncPairService struct {
	db *sql.DB
}

// NewSyncPairService creates a new instance of SyncPairService
func NewSyncPairService(db *sql.DB) *SyncPairService {
	return &SyncPairService{
		db: db,
	}
}

// SyncPair represents a trainer and Pokemon combination
type SyncPair struct {
	ID                int        `json:"id"`
	SyncPairID        string     `json:"sync_pair_id"`
	TrainerID         string     `json:"trainer_id"`
	PokemonID         string     `json:"pokemon_id"`
	TrainerName       string     `json:"trainer_name"`
	PokemonName       string     `json:"pokemon_name"`
	PokemonType       string     `json:"pokemon_type"`
	SyncPairName      string     `json:"sync_pair_name"`
	RarityStars       int        `json:"rarity_stars"`
	Level             int        `json:"level"`
	Experience        int64      `json:"experience"`
	MaxLevel          int        `json:"max_level"`
	HP                int        `json:"hp"`
	Attack            int        `json:"attack"`
	Defense           int        `json:"defense"`
	SpAttack          int        `json:"sp_attack"`
	SpDefense         int        `json:"sp_defense"`
	Speed             int        `json:"speed"`
	Move1ID           *string    `json:"move_1_id"`
	Move2ID           *string    `json:"move_2_id"`
	Move3ID           *string    `json:"move_3_id"`
	SyncMoveID        *string    `json:"sync_move_id"`
	SyncMoveReady     int        `json:"sync_move_ready_percentage"`
	PotentialUnlocked int        `json:"potential_unlocked"`
	IsActive          bool       `json:"is_active"`
	ObtainedAt        time.Time  `json:"obtained_at"`
	LastUsedAt        *time.Time `json:"last_used_at,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

// CreateSyncPair creates a new Sync Pair in the database
func (s *SyncPairService) CreateSyncPair(trainerID, pokemonID, trainerName, pokemonName, pokemonType string, rarityStars int) (*SyncPair, error) {
	syncPairID := fmt.Sprintf("%s_%s", trainerID, pokemonID)

	sp := &SyncPair{
		SyncPairID:        syncPairID,
		TrainerID:         trainerID,
		PokemonID:         pokemonID,
		TrainerName:       trainerName,
		PokemonName:       pokemonName,
		PokemonType:       pokemonType,
		SyncPairName:      fmt.Sprintf("%s & %s", trainerName, pokemonName),
		RarityStars:       rarityStars,
		Level:             1,
		Experience:        0,
		MaxLevel:          130,
		HP:                100,
		Attack:            50,
		Defense:           50,
		SpAttack:          50,
		SpDefense:         50,
		Speed:             50,
		SyncMoveReady:     0,
		PotentialUnlocked: 0,
		IsActive:          true,
		ObtainedAt:        time.Now(),
	}

	query := `
		INSERT INTO sync_pairs 
		(sync_pair_id, trainer_id, pokemon_id, trainer_name, pokemon_name, pokemon_type, 
		 sync_pair_name, rarity_stars, level, experience, max_level, hp, attack, defense,
		 sp_attack, sp_defense, speed, is_active, obtained_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
		RETURNING id, created_at, updated_at`

	err := s.db.QueryRow(query, sp.SyncPairID, sp.TrainerID, sp.PokemonID, sp.TrainerName,
		sp.PokemonName, sp.PokemonType, sp.SyncPairName, sp.RarityStars, sp.Level,
		sp.Experience, sp.MaxLevel, sp.HP, sp.Attack, sp.Defense, sp.SpAttack,
		sp.SpDefense, sp.Speed, sp.IsActive, sp.ObtainedAt).Scan(&sp.ID, &sp.CreatedAt, &sp.UpdatedAt)

	if err != nil {
		log.Printf("Error creating sync pair: %v", err)
		return nil, err
	}

	return sp, nil
}

// GetSyncPair retrieves a Sync Pair by ID
func (s *SyncPairService) GetSyncPair(syncPairID string) (*SyncPair, error) {
	sp := &SyncPair{}
	query := `SELECT id, sync_pair_id, trainer_id, pokemon_id, trainer_name, pokemon_name, pokemon_type,
		sync_pair_name, rarity_stars, level, experience, max_level, hp, attack, defense, sp_attack,
		sp_defense, speed, move_1_id, move_2_id, move_3_id, sync_move_id, sync_move_ready_percentage,
		potential_unlocked, is_active, obtained_at, last_used_at, created_at, updated_at
		FROM sync_pairs WHERE sync_pair_id = $1`

	err := s.db.QueryRow(query, syncPairID).Scan(
		&sp.ID, &sp.SyncPairID, &sp.TrainerID, &sp.PokemonID, &sp.TrainerName, &sp.PokemonName,
		&sp.PokemonType, &sp.SyncPairName, &sp.RarityStars, &sp.Level, &sp.Experience, &sp.MaxLevel,
		&sp.HP, &sp.Attack, &sp.Defense, &sp.SpAttack, &sp.SpDefense, &sp.Speed,
		&sp.Move1ID, &sp.Move2ID, &sp.Move3ID, &sp.SyncMoveID, &sp.SyncMoveReady,
		&sp.PotentialUnlocked, &sp.IsActive, &sp.ObtainedAt, &sp.LastUsedAt, &sp.CreatedAt, &sp.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("sync pair not found: %s", syncPairID)
		}
		log.Printf("Error getting sync pair: %v", err)
		return nil, err
	}

	return sp, nil
}

// GetPlayerSyncPairs retrieves all Sync Pairs for a player
func (s *SyncPairService) GetPlayerSyncPairs(trainerID string) ([]SyncPair, error) {
	query := `SELECT id, sync_pair_id, trainer_id, pokemon_id, trainer_name, pokemon_name, pokemon_type,
		sync_pair_name, rarity_stars, level, experience, max_level, hp, attack, defense, sp_attack,
		sp_defense, speed, move_1_id, move_2_id, move_3_id, sync_move_id, sync_move_ready_percentage,
		potential_unlocked, is_active, obtained_at, last_used_at, created_at, updated_at
		FROM sync_pairs WHERE trainer_id = $1 AND is_active = true ORDER BY level DESC`

	rows, err := s.db.Query(query, trainerID)
	if err != nil {
		log.Printf("Error querying sync pairs: %v", err)
		return nil, err
	}
	defer rows.Close()

	var syncPairs []SyncPair
	for rows.Next() {
		sp := SyncPair{}
		err := rows.Scan(
			&sp.ID, &sp.SyncPairID, &sp.TrainerID, &sp.PokemonID, &sp.TrainerName, &sp.PokemonName,
			&sp.PokemonType, &sp.SyncPairName, &sp.RarityStars, &sp.Level, &sp.Experience, &sp.MaxLevel,
			&sp.HP, &sp.Attack, &sp.Defense, &sp.SpAttack, &sp.SpDefense, &sp.Speed,
			&sp.Move1ID, &sp.Move2ID, &sp.Move3ID, &sp.SyncMoveID, &sp.SyncMoveReady,
			&sp.PotentialUnlocked, &sp.IsActive, &sp.ObtainedAt, &sp.LastUsedAt, &sp.CreatedAt, &sp.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning sync pair: %v", err)
			continue
		}
		syncPairs = append(syncPairs, sp)
	}

	return syncPairs, nil
}

// LevelUpSyncPair levels up a Sync Pair
func (s *SyncPairService) LevelUpSyncPair(syncPairID string, levels int) (*SyncPair, error) {
	sp, err := s.GetSyncPair(syncPairID)
	if err != nil {
		return nil, err
	}

	newLevel := sp.Level + levels
	if newLevel > sp.MaxLevel {
		newLevel = sp.MaxLevel
	}

	// Calculate stat increases
	levelDiff := newLevel - sp.Level
	statIncrease := levelDiff * 5 // Each level increases stats by ~5

	query := `UPDATE sync_pairs SET level = $1, hp = hp + $2, attack = attack + $2,
		defense = defense + $2, sp_attack = sp_attack + $2, sp_defense = sp_defense + $2,
		speed = speed + $2, updated_at = NOW()
		WHERE sync_pair_id = $3 RETURNING *`

	updatedSP := &SyncPair{}
	err = s.db.QueryRow(query, newLevel, statIncrease, syncPairID).Scan(
		&updatedSP.ID, &updatedSP.SyncPairID, &updatedSP.TrainerID, &updatedSP.PokemonID,
		&updatedSP.TrainerName, &updatedSP.PokemonName, &updatedSP.PokemonType, &updatedSP.SyncPairName,
		&updatedSP.RarityStars, &updatedSP.Level, &updatedSP.Experience, &updatedSP.MaxLevel,
		&updatedSP.HP, &updatedSP.Attack, &updatedSP.Defense, &updatedSP.SpAttack, &updatedSP.SpDefense,
		&updatedSP.Speed, &updatedSP.Move1ID, &updatedSP.Move2ID, &updatedSP.Move3ID, &updatedSP.SyncMoveID,
		&updatedSP.SyncMoveReady, &updatedSP.PotentialUnlocked, &updatedSP.IsActive, &updatedSP.ObtainedAt,
		&updatedSP.LastUsedAt, &updatedSP.CreatedAt, &updatedSP.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error leveling up sync pair: %v", err)
		return nil, err
	}

	return updatedSP, nil
}

// UnlockPotential unlocks a potential tier for a Sync Pair
func (s *SyncPairService) UnlockPotential(syncPairID string) (*SyncPair, error) {
	sp, err := s.GetSyncPair(syncPairID)
	if err != nil {
		return nil, err
	}

	if sp.PotentialUnlocked >= 20 {
		return nil, fmt.Errorf("maximum potential already unlocked")
	}

	newPotential := sp.PotentialUnlocked + 1
	statBonus := newPotential * 10

	query := `UPDATE sync_pairs SET potential_unlocked = $1, hp = hp + $2, attack = attack + $2,
		defense = defense + $2, sp_attack = sp_attack + $2, sp_defense = sp_defense + $2,
		speed = speed + $2, max_level = max_level + 10, updated_at = NOW()
		WHERE sync_pair_id = $3 RETURNING *`

	updatedSP := &SyncPair{}
	err = s.db.QueryRow(query, newPotential, statBonus, syncPairID).Scan(
		&updatedSP.ID, &updatedSP.SyncPairID, &updatedSP.TrainerID, &updatedSP.PokemonID,
		&updatedSP.TrainerName, &updatedSP.PokemonName, &updatedSP.PokemonType, &updatedSP.SyncPairName,
		&updatedSP.RarityStars, &updatedSP.Level, &updatedSP.Experience, &updatedSP.MaxLevel,
		&updatedSP.HP, &updatedSP.Attack, &updatedSP.Defense, &updatedSP.SpAttack, &updatedSP.SpDefense,
		&updatedSP.Speed, &updatedSP.Move1ID, &updatedSP.Move2ID, &updatedSP.Move3ID, &updatedSP.SyncMoveID,
		&updatedSP.SyncMoveReady, &updatedSP.PotentialUnlocked, &updatedSP.IsActive, &updatedSP.ObtainedAt,
		&updatedSP.LastUsedAt, &updatedSP.CreatedAt, &updatedSP.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error unlocking potential: %v", err)
		return nil, err
	}

	return updatedSP, nil
}

// SetMoves sets the moves for a Sync Pair
func (s *SyncPairService) SetMoves(syncPairID string, move1, move2, move3, syncMove string) (*SyncPair, error) {
	query := `UPDATE sync_pairs SET move_1_id = $1, move_2_id = $2, move_3_id = $3,
		sync_move_id = $4, updated_at = NOW()
		WHERE sync_pair_id = $5 RETURNING *`

	updatedSP := &SyncPair{}
	err := s.db.QueryRow(query, move1, move2, move3, syncMove, syncPairID).Scan(
		&updatedSP.ID, &updatedSP.SyncPairID, &updatedSP.TrainerID, &updatedSP.PokemonID,
		&updatedSP.TrainerName, &updatedSP.PokemonName, &updatedSP.PokemonType, &updatedSP.SyncPairName,
		&updatedSP.RarityStars, &updatedSP.Level, &updatedSP.Experience, &updatedSP.MaxLevel,
		&updatedSP.HP, &updatedSP.Attack, &updatedSP.Defense, &updatedSP.SpAttack, &updatedSP.SpDefense,
		&updatedSP.Speed, &updatedSP.Move1ID, &updatedSP.Move2ID, &updatedSP.Move3ID, &updatedSP.SyncMoveID,
		&updatedSP.SyncMoveReady, &updatedSP.PotentialUnlocked, &updatedSP.IsActive, &updatedSP.ObtainedAt,
		&updatedSP.LastUsedAt, &updatedSP.CreatedAt, &updatedSP.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error setting moves: %v", err)
		return nil, err
	}

	return updatedSP, nil
}

// IncreaseSyncMoveReady increases the sync move ready percentage
func (s *SyncPairService) IncreaseSyncMoveReady(syncPairID string, percentage int) error {
	query := `UPDATE sync_pairs SET sync_move_ready_percentage = LEAST(sync_move_ready_percentage + $1, 100)
		WHERE sync_pair_id = $2`

	result, err := s.db.Exec(query, percentage, syncPairID)
	if err != nil {
		log.Printf("Error increasing sync move ready: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("sync pair not found: %s", syncPairID)
	}

	return nil
}

// UseSyncMove resets the sync move ready percentage to 0
func (s *SyncPairService) UseSyncMove(syncPairID string) error {
	query := `UPDATE sync_pairs SET sync_move_ready_percentage = 0, last_used_at = NOW()
		WHERE sync_pair_id = $1`

	result, err := s.db.Exec(query, syncPairID)
	if err != nil {
		log.Printf("Error using sync move: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("sync pair not found: %s", syncPairID)
	}

	return nil
}

// GetPlayerTeam retrieves a player's battle team
type BattleTeam struct {
	TeamID     string     `json:"team_id"`
	PlayerID   string     `json:"player_id"`
	TeamName   string     `json:"team_name"`
	SyncPairs  []SyncPair `json:"sync_pairs"`
	IsMainTeam bool       `json:"is_main_team"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// GetPlayerTeam retrieves a player's battle team
func (s *SyncPairService) GetPlayerTeam(teamID string) (*BattleTeam, error) {
	team := &BattleTeam{}
	query := `SELECT team_id, player_id, team_name, is_main_team, created_at, updated_at
		FROM player_battle_teams WHERE team_id = $1`

	err := s.db.QueryRow(query, teamID).Scan(
		&team.TeamID, &team.PlayerID, &team.TeamName, &team.IsMainTeam, &team.CreatedAt, &team.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("team not found: %s", teamID)
		}
		log.Printf("Error getting team: %v", err)
		return nil, err
	}

	// Now get the sync pairs for the team
	teamQuery := `SELECT sp.id, sp.sync_pair_id, sp.trainer_id, sp.pokemon_id, sp.trainer_name,
		sp.pokemon_name, sp.pokemon_type, sp.sync_pair_name, sp.rarity_stars, sp.level, sp.experience,
		sp.max_level, sp.hp, sp.attack, sp.defense, sp.sp_attack, sp.sp_defense, sp.speed,
		sp.move_1_id, sp.move_2_id, sp.move_3_id, sp.sync_move_id, sp.sync_move_ready_percentage,
		sp.potential_unlocked, sp.is_active, sp.obtained_at, sp.last_used_at, sp.created_at, sp.updated_at
		FROM sync_pairs sp
		JOIN (
			SELECT 1 as position, sync_pair_1_id as id FROM player_battle_teams WHERE team_id = $1
			UNION ALL
			SELECT 2, sync_pair_2_id FROM player_battle_teams WHERE team_id = $1
			UNION ALL
			SELECT 3, sync_pair_3_id FROM player_battle_teams WHERE team_id = $1
		) team_sp ON sp.sync_pair_id = team_sp.id
		ORDER BY team_sp.position`

	rows, err := s.db.Query(teamQuery, teamID)
	if err != nil {
		log.Printf("Error querying team sync pairs: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		sp := SyncPair{}
		err := rows.Scan(
			&sp.ID, &sp.SyncPairID, &sp.TrainerID, &sp.PokemonID, &sp.TrainerName, &sp.PokemonName,
			&sp.PokemonType, &sp.SyncPairName, &sp.RarityStars, &sp.Level, &sp.Experience, &sp.MaxLevel,
			&sp.HP, &sp.Attack, &sp.Defense, &sp.SpAttack, &sp.SpDefense, &sp.Speed,
			&sp.Move1ID, &sp.Move2ID, &sp.Move3ID, &sp.SyncMoveID, &sp.SyncMoveReady,
			&sp.PotentialUnlocked, &sp.IsActive, &sp.ObtainedAt, &sp.LastUsedAt, &sp.CreatedAt, &sp.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning sync pair: %v", err)
			continue
		}
		team.SyncPairs = append(team.SyncPairs, sp)
	}

	return team, nil
}

// CreateTeam creates a new battle team for a player
func (s *SyncPairService) CreateTeam(playerID, teamName string, syncPair1, syncPair2, syncPair3 string, isMainTeam bool) (*BattleTeam, error) {
	teamID := fmt.Sprintf("team_%s_%d", playerID, time.Now().Unix())

	query := `INSERT INTO player_battle_teams 
		(team_id, player_id, team_name, sync_pair_1_id, sync_pair_2_id, sync_pair_3_id, is_main_team)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := s.db.Exec(query, teamID, playerID, teamName, syncPair1, syncPair2, syncPair3, isMainTeam)
	if err != nil {
		log.Printf("Error creating team: %v", err)
		return nil, err
	}

	return s.GetPlayerTeam(teamID)
}

// GetPlayerTeams retrieves all battle teams for a player
func (s *SyncPairService) GetPlayerTeams(playerID string) ([]BattleTeam, error) {
	query := `SELECT team_id, player_id, team_name, is_main_team, created_at, updated_at
		FROM player_battle_teams WHERE player_id = $1 ORDER BY is_main_team DESC, created_at DESC`

	rows, err := s.db.Query(query, playerID)
	if err != nil {
		log.Printf("Error querying teams: %v", err)
		return nil, err
	}
	defer rows.Close()

	var teams []BattleTeam
	for rows.Next() {
		team := BattleTeam{}
		err := rows.Scan(&team.TeamID, &team.PlayerID, &team.TeamName, &team.IsMainTeam, &team.CreatedAt, &team.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning team: %v", err)
			continue
		}
		teams = append(teams, team)
	}

	return teams, nil
}

// GetPlayerGameStatistics retrieves comprehensive game statistics for a player
type PlayerGameStats struct {
	PlayerID              string  `json:"player_id"`
	TotalSyncPairs        int     `json:"total_sync_pairs"`
	TotalLevelSum         int64   `json:"total_level_sum"`
	AverageSyncPairLevel  float64 `json:"average_sync_pair_level"`
	SyncPairsWithMaxLevel int     `json:"sync_pairs_with_max_level"`
	HighestRarityObtained int     `json:"highest_rarity_obtained"`
	TotalPowerIndex       int64   `json:"total_power_index"`
}

// GetPlayerGameStatistics retrieves game statistics for a player
func (s *SyncPairService) GetPlayerGameStatistics(playerID string) (*PlayerGameStats, error) {
	stats := &PlayerGameStats{PlayerID: playerID}

	query := `SELECT 
		total_sync_pairs, total_level_sum, average_sync_pair_level, 
		sync_pairs_with_max_level, highest_rarity_obtained, total_power_index
		FROM player_game_statistics WHERE player_id = $1`

	err := s.db.QueryRow(query, playerID).Scan(
		&stats.TotalSyncPairs, &stats.TotalLevelSum, &stats.AverageSyncPairLevel,
		&stats.SyncPairsWithMaxLevel, &stats.HighestRarityObtained, &stats.TotalPowerIndex,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// Initialize statistics if not exist
			return s.InitializePlayerStatistics(playerID)
		}
		log.Printf("Error getting player statistics: %v", err)
		return nil, err
	}

	return stats, nil
}

// InitializePlayerStatistics initializes statistics for a new player
func (s *SyncPairService) InitializePlayerStatistics(playerID string) (*PlayerGameStats, error) {
	stats := &PlayerGameStats{
		PlayerID:              playerID,
		TotalSyncPairs:        0,
		TotalLevelSum:         0,
		AverageSyncPairLevel:  0,
		HighestRarityObtained: 3,
		TotalPowerIndex:       0,
	}

	query := `INSERT INTO player_game_statistics 
		(player_id, total_sync_pairs, total_level_sum, average_sync_pair_level, 
		 sync_pairs_with_max_level, highest_rarity_obtained, total_power_index)
		VALUES ($1, 0, 0, 0, 0, 3, 0)
		ON CONFLICT (player_id) DO NOTHING`

	_, err := s.db.Exec(query, playerID)
	if err != nil {
		log.Printf("Error initializing player statistics: %v", err)
		return nil, err
	}

	return stats, nil
}

// UpdatePlayerStatistics updates the comprehensive statistics for a player
func (s *SyncPairService) UpdatePlayerStatistics(playerID string) error {
	query := `
		WITH player_pairs AS (
			SELECT 
				COUNT(*) as total_pairs,
				SUM(level) as total_level,
				COUNT(CASE WHEN level >= max_level THEN 1 END) as max_level_count,
				MAX(rarity_stars) as max_rarity,
				SUM((hp + attack + defense + sp_attack + sp_defense + speed) * level) as power_index
			FROM sync_pairs
			WHERE trainer_id = $1 AND is_active = true
		)
		UPDATE player_game_statistics
		SET 
			total_sync_pairs = COALESCE(player_pairs.total_pairs, 0),
			total_level_sum = COALESCE(player_pairs.total_level, 0),
			average_sync_pair_level = CASE 
				WHEN COALESCE(player_pairs.total_pairs, 0) > 0 
				THEN ROUND(COALESCE(player_pairs.total_level, 0)::NUMERIC / player_pairs.total_pairs, 2)
				ELSE 0
			END,
			sync_pairs_with_max_level = COALESCE(player_pairs.max_level_count, 0),
			highest_rarity_obtained = COALESCE(player_pairs.max_rarity, 3),
			total_power_index = COALESCE(player_pairs.power_index, 0),
			last_updated = NOW()
		FROM player_pairs
		WHERE player_id = $1
	`

	_, err := s.db.Exec(query, playerID)
	if err != nil {
		log.Printf("Error updating player statistics: %v", err)
		return err
	}

	return nil
}

// ExportSyncPairData exports sync pair data to JSON
func (s *SyncPairService) ExportSyncPairData(playerID string) ([]byte, error) {
	syncPairs, err := s.GetPlayerSyncPairs(playerID)
	if err != nil {
		return nil, err
	}

	data, err := json.MarshalIndent(syncPairs, "", "  ")
	if err != nil {
		log.Printf("Error exporting sync pair data: %v", err)
		return nil, err
	}

	return data, nil
}
