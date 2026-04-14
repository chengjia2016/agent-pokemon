package service

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// BattleSession represents a 3v3 or 1v1 battle session
// Supports multi-player battles with up to 3 players
type BattleSession struct {
	ID           int             `json:"id"`
	SessionID    string          `json:"session_id"`
	Player1ID    string          `json:"player_1_id"`
	Player2ID    string          `json:"player_2_id"`
	Player3ID    *string         `json:"player_3_id,omitempty"`
	BattleType   string          `json:"battle_type"`   // single, multi
	BattleStatus string          `json:"battle_status"` // waiting, active, finished
	Player1Team  json.RawMessage `json:"player_1_team"` // JSON array of 3 sync pair IDs
	Player2Team  json.RawMessage `json:"player_2_team"` // JSON array of 3 sync pair IDs
	Player3Team  json.RawMessage `json:"player_3_team,omitempty"`
	CurrentRound int             `json:"current_round"`
	MaxRounds    int             `json:"max_rounds"`
	Player1HP    json.RawMessage `json:"player_1_hp,omitempty"` // JSONB HP tracking
	Player2HP    json.RawMessage `json:"player_2_hp,omitempty"`
	Player3HP    json.RawMessage `json:"player_3_hp,omitempty"`
	Player1Score int             `json:"player_1_score"`
	Player2Score int             `json:"player_2_score"`
	Player3Score int             `json:"player_3_score,omitempty"`
	WinnerID     *string         `json:"winner_id,omitempty"`
	WinnerTeam   *string         `json:"winner_team,omitempty"`
	BattleLog    json.RawMessage `json:"battle_log,omitempty"`
	StartedAt    *time.Time      `json:"started_at,omitempty"`
	EndedAt      *time.Time      `json:"ended_at,omitempty"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

// BattleAction represents a single action in a battle
type BattleAction struct {
	ID         int             `json:"id"`
	SessionID  string          `json:"session_id"`
	RoundNum   int             `json:"round_number"`
	PlayerID   string          `json:"player_id"`
	ActionType string          `json:"action_type"` // move, switch, mega-evolve, sync-move
	ActionData json.RawMessage `json:"action_data"`
	Timestamp  time.Time       `json:"timestamp"`
}

// BattleStatistics represents battle performance stats for a player
type BattleStatistics struct {
	ID               int        `json:"id"`
	PlayerID         string     `json:"player_id"`
	TotalBattles     int        `json:"total_battles"`
	Wins             int        `json:"wins"`
	Losses           int        `json:"losses"`
	WinRate          float64    `json:"win_rate"`
	TotalDamageDealt int64      `json:"total_damage_dealt"`
	TotalDamageTaken int64      `json:"total_damage_taken"`
	SyncMovesUsed    int        `json:"sync_moves_used"`
	MostUsedPokemon  *string    `json:"most_used_pokemon,omitempty"`
	LastBattleAt     *time.Time `json:"last_battle_at,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// BattleMastersExService handles 3v3 battles for Masters EX
type BattleMastersExService struct {
	db *sql.DB
}

// NewBattleMastersExService creates a new battle service
func NewBattleMastersExService(db *sql.DB) *BattleMastersExService {
	return &BattleMastersExService{
		db: db,
	}
}

// StartBattle starts a new battle session (supports 1v1 or multi-player)
func (s *BattleMastersExService) StartBattle(player1ID, player2ID string, player3ID *string,
	player1Team, player2Team, player3Team json.RawMessage) (*BattleSession, error) {

	sessionID := fmt.Sprintf("battle_%s_vs_%s_%d", player1ID, player2ID, time.Now().Unix())

	// Determine battle type
	battleType := "single"
	if player3ID != nil && *player3ID != "" {
		battleType = "multi"
	}

	// Initialize HP for all players (3 sync pairs each, starting with full HP)
	player1HP := []map[string]interface{}{
		{"sync_pair_index": 0, "current_hp": 100, "max_hp": 100},
		{"sync_pair_index": 1, "current_hp": 100, "max_hp": 100},
		{"sync_pair_index": 2, "current_hp": 100, "max_hp": 100},
	}
	player2HP := []map[string]interface{}{
		{"sync_pair_index": 0, "current_hp": 100, "max_hp": 100},
		{"sync_pair_index": 1, "current_hp": 100, "max_hp": 100},
		{"sync_pair_index": 2, "current_hp": 100, "max_hp": 100},
	}
	var player3HP json.RawMessage
	if player3ID != nil && *player3ID != "" {
		p3HP := []map[string]interface{}{
			{"sync_pair_index": 0, "current_hp": 100, "max_hp": 100},
			{"sync_pair_index": 1, "current_hp": 100, "max_hp": 100},
			{"sync_pair_index": 2, "current_hp": 100, "max_hp": 100},
		}
		player3HP, _ = json.Marshal(p3HP)
	}
	player1HPJson, _ := json.Marshal(player1HP)
	player2HPJson, _ := json.Marshal(player2HP)

	// Initialize battle log
	battleLog := []map[string]interface{}{
		{
			"round":     0,
			"event":     "Battle started",
			"timestamp": time.Now().Unix(),
		},
	}
	battleLogJson, _ := json.Marshal(battleLog)

	now := time.Now()
	query := `
		INSERT INTO battle_sessions 
		(session_id, player_1_id, player_2_id, player_3_id, battle_type, battle_status,
		 player_1_team, player_2_team, player_3_team, current_round, max_rounds,
		 player_1_hp, player_2_hp, player_3_hp, player_1_score, player_2_score, player_3_score,
		 battle_log, started_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
		RETURNING id, created_at, updated_at`

	bs := &BattleSession{
		SessionID:    sessionID,
		Player1ID:    player1ID,
		Player2ID:    player2ID,
		Player3ID:    player3ID,
		BattleType:   battleType,
		BattleStatus: "active",
		Player1Team:  player1Team,
		Player2Team:  player2Team,
		Player3Team:  player3Team,
		CurrentRound: 0,
		MaxRounds:    20,
		Player1HP:    player1HPJson,
		Player2HP:    player2HPJson,
		Player3HP:    player3HP,
		Player1Score: 0,
		Player2Score: 0,
		BattleLog:    battleLogJson,
		StartedAt:    &now,
	}

	err := s.db.QueryRow(query, bs.SessionID, bs.Player1ID, bs.Player2ID, bs.Player3ID,
		bs.BattleType, bs.BattleStatus, bs.Player1Team, bs.Player2Team, bs.Player3Team,
		bs.CurrentRound, bs.MaxRounds, bs.Player1HP, bs.Player2HP, bs.Player3HP,
		bs.Player1Score, bs.Player2Score, bs.Player3Score, bs.BattleLog, bs.StartedAt, bs.CreatedAt, bs.UpdatedAt).Scan(
		&bs.ID, &bs.CreatedAt, &bs.UpdatedAt)

	if err != nil {
		log.Printf("Error starting battle: %v", err)
		return nil, err
	}

	return bs, nil
}

// GetBattleSession retrieves a battle session by session ID
func (s *BattleMastersExService) GetBattleSession(sessionID string) (*BattleSession, error) {
	bs := &BattleSession{}
	query := `
		SELECT id, session_id, player_1_id, player_2_id, player_3_id, battle_type, battle_status,
		       player_1_team, player_2_team, player_3_team, current_round, max_rounds,
		       player_1_hp, player_2_hp, player_3_hp, player_1_score, player_2_score, player_3_score,
		       winner_id, winner_team, battle_log, started_at, ended_at, created_at, updated_at
		FROM battle_sessions WHERE session_id = $1`

	var player3ID, winnerID, winnerTeam sql.NullString
	var player3HP, battleLog sql.NullString
	var startedAt, endedAt sql.NullTime

	err := s.db.QueryRow(query, sessionID).Scan(
		&bs.ID, &bs.SessionID, &bs.Player1ID, &bs.Player2ID, &player3ID, &bs.BattleType, &bs.BattleStatus,
		&bs.Player1Team, &bs.Player2Team, &player3HP, &bs.CurrentRound, &bs.MaxRounds,
		&bs.Player1HP, &bs.Player2HP, &player3HP, &bs.Player1Score, &bs.Player2Score, &bs.Player3Score,
		&winnerID, &winnerTeam, &battleLog, &startedAt, &endedAt, &bs.CreatedAt, &bs.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("battle session not found: %s", sessionID)
		}
		log.Printf("Error getting battle session: %v", err)
		return nil, err
	}

	if player3ID.Valid {
		bs.Player3ID = &player3ID.String
	}
	if winnerID.Valid {
		bs.WinnerID = &winnerID.String
	}
	if winnerTeam.Valid {
		bs.WinnerTeam = &winnerTeam.String
	}
	if battleLog.Valid {
		bs.BattleLog = json.RawMessage(battleLog.String)
	}
	if player3HP.Valid {
		bs.Player3HP = json.RawMessage(player3HP.String)
	}
	if startedAt.Valid {
		bs.StartedAt = &startedAt.Time
	}
	if endedAt.Valid {
		bs.EndedAt = &endedAt.Time
	}

	return bs, nil
}

// ExecuteBattleAction records an action taken during battle
func (s *BattleMastersExService) ExecuteBattleAction(sessionID string, roundNumber int, playerID, actionType string, actionData json.RawMessage) (*BattleAction, error) {
	action := &BattleAction{
		SessionID:  sessionID,
		RoundNum:   roundNumber,
		PlayerID:   playerID,
		ActionType: actionType,
		ActionData: actionData,
		Timestamp:  time.Now(),
	}

	query := `
		INSERT INTO battle_actions (session_id, round_number, player_id, action_type, action_data, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	err := s.db.QueryRow(query, action.SessionID, action.RoundNum, action.PlayerID, action.ActionType, action.ActionData, action.Timestamp).Scan(&action.ID)
	if err != nil {
		log.Printf("Error executing battle action: %v", err)
		return nil, err
	}

	return action, nil
}

// UpdateBattleRound updates the battle session with new round information
func (s *BattleMastersExService) UpdateBattleRound(sessionID string, currentRound int,
	player1Score, player2Score, player3Score int,
	player1HP, player2HP, player3HP json.RawMessage) (*BattleSession, error) {

	query := `
		UPDATE battle_sessions
		SET current_round = $1, player_1_score = $2, player_2_score = $3, player_3_score = $4,
		    player_1_hp = $5, player_2_hp = $6, player_3_hp = $7, updated_at = $8
		WHERE session_id = $9`

	_, err := s.db.Exec(query, currentRound, player1Score, player2Score, player3Score,
		player1HP, player2HP, player3HP, time.Now(), sessionID)

	if err != nil {
		log.Printf("Error updating battle round: %v", err)
		return nil, err
	}

	// Return updated battle session
	return s.GetBattleSession(sessionID)
}

// EndBattle ends a battle session and records the winner
func (s *BattleMastersExService) EndBattle(sessionID, winnerID string, winnerTeam json.RawMessage) (*BattleSession, error) {
	now := time.Now()
	query := `
		UPDATE battle_sessions
		SET battle_status = $1, winner_id = $2, winner_team = $3, ended_at = $4, updated_at = $5
		WHERE session_id = $6`

	_, err := s.db.Exec(query, "finished", winnerID, winnerTeam, now, now, sessionID)
	if err != nil {
		log.Printf("Error ending battle: %v", err)
		return nil, err
	}

	// Record the win/loss in battle statistics
	// Get both player IDs from session
	bs, err := s.GetBattleSession(sessionID)
	if err != nil {
		return nil, err
	}

	// Update winner stats
	s.updatePlayerStats(winnerID, true)

	// Update loser stats
	if winnerID == bs.Player1ID {
		s.updatePlayerStats(bs.Player2ID, false)
	} else {
		s.updatePlayerStats(bs.Player1ID, false)
	}

	return s.GetBattleSession(sessionID)
}

// GetBattleStatistics gets battle stats for a player
func (s *BattleMastersExService) GetBattleStatistics(playerID string) (*BattleStatistics, error) {
	stats := &BattleStatistics{
		PlayerID: playerID,
	}

	query := `
		SELECT id, player_id, total_battles, wins, losses, win_rate, total_damage_dealt,
		       total_damage_taken, sync_moves_used, most_used_pokemon, last_battle_at, created_at, updated_at
		FROM battle_statistics WHERE player_id = $1`

	var lastBattleAt, mostUsedPokemon sql.NullString

	err := s.db.QueryRow(query, playerID).Scan(
		&stats.ID, &stats.PlayerID, &stats.TotalBattles, &stats.Wins, &stats.Losses,
		&stats.WinRate, &stats.TotalDamageDealt, &stats.TotalDamageTaken, &stats.SyncMovesUsed,
		&mostUsedPokemon, &lastBattleAt, &stats.CreatedAt, &stats.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// Create new stats record if doesn't exist
			return s.CreateBattleStatistics(playerID)
		}
		log.Printf("Error getting battle statistics: %v", err)
		return nil, err
	}

	if mostUsedPokemon.Valid {
		stats.MostUsedPokemon = &mostUsedPokemon.String
	}
	if lastBattleAt.Valid {
		lastBattleTime, _ := time.Parse(time.RFC3339, lastBattleAt.String)
		stats.LastBattleAt = &lastBattleTime
	}

	return stats, nil
}

// CreateBattleStatistics creates a new battle statistics record
func (s *BattleMastersExService) CreateBattleStatistics(playerID string) (*BattleStatistics, error) {
	stats := &BattleStatistics{
		PlayerID:     playerID,
		TotalBattles: 0,
		Wins:         0,
		Losses:       0,
		WinRate:      0.0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query := `
		INSERT INTO battle_statistics (player_id, total_battles, wins, losses, win_rate, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at`

	err := s.db.QueryRow(query, stats.PlayerID, stats.TotalBattles, stats.Wins, stats.Losses,
		stats.WinRate, stats.CreatedAt, stats.UpdatedAt).Scan(&stats.ID, &stats.CreatedAt, &stats.UpdatedAt)

	if err != nil {
		log.Printf("Error creating battle statistics: %v", err)
		return nil, err
	}

	return stats, nil
}

// GetBattleHistory gets recent battles for a player
func (s *BattleMastersExService) GetBattleHistory(playerID string, limit int) ([]*BattleSession, error) {
	query := `
		SELECT id, session_id, player_1_id, player_2_id, player_3_id, battle_type, battle_status,
		       player_1_team, player_2_team, player_3_team, current_round, max_rounds,
		       player_1_hp, player_2_hp, player_3_hp, player_1_score, player_2_score, player_3_score,
		       winner_id, winner_team, battle_log, started_at, ended_at, created_at, updated_at
		FROM battle_sessions 
		WHERE (player_1_id = $1 OR player_2_id = $1 OR player_3_id = $1) AND battle_status = 'finished'
		ORDER BY ended_at DESC
		LIMIT $2`

	rows, err := s.db.Query(query, playerID, limit)
	if err != nil {
		log.Printf("Error getting battle history: %v", err)
		return nil, err
	}
	defer rows.Close()

	var battles []*BattleSession
	for rows.Next() {
		bs := &BattleSession{}
		var player3ID, winnerID, winnerTeam sql.NullString
		var player3HP, battleLog sql.NullString
		var startedAt, endedAt sql.NullTime

		err := rows.Scan(
			&bs.ID, &bs.SessionID, &bs.Player1ID, &bs.Player2ID, &player3ID, &bs.BattleType, &bs.BattleStatus,
			&bs.Player1Team, &bs.Player2Team, &player3HP, &bs.CurrentRound, &bs.MaxRounds,
			&bs.Player1HP, &bs.Player2HP, &player3HP, &bs.Player1Score, &bs.Player2Score, &bs.Player3Score,
			&winnerID, &winnerTeam, &battleLog, &startedAt, &endedAt, &bs.CreatedAt, &bs.UpdatedAt,
		)

		if err != nil {
			log.Printf("Error scanning battle history row: %v", err)
			continue
		}

		if player3ID.Valid {
			bs.Player3ID = &player3ID.String
		}
		if winnerID.Valid {
			bs.WinnerID = &winnerID.String
		}
		if winnerTeam.Valid {
			bs.WinnerTeam = &winnerTeam.String
		}
		if battleLog.Valid {
			bs.BattleLog = json.RawMessage(battleLog.String)
		}
		if player3HP.Valid {
			bs.Player3HP = json.RawMessage(player3HP.String)
		}
		if startedAt.Valid {
			bs.StartedAt = &startedAt.Time
		}
		if endedAt.Valid {
			bs.EndedAt = &endedAt.Time
		}

		battles = append(battles, bs)
	}

	return battles, nil
}

// updatePlayerStats helper method to update player's battle statistics
func (s *BattleMastersExService) updatePlayerStats(playerID string, isWin bool) error {
	stats, err := s.GetBattleStatistics(playerID)
	if err != nil {
		return err
	}

	stats.TotalBattles++
	if isWin {
		stats.Wins++
	} else {
		stats.Losses++
	}

	if stats.TotalBattles > 0 {
		stats.WinRate = float64(stats.Wins) / float64(stats.TotalBattles) * 100
	}

	query := `
		UPDATE battle_statistics
		SET total_battles = $1, wins = $2, losses = $3, win_rate = $4, 
		    last_battle_at = $5, updated_at = $6
		WHERE player_id = $7`

	_, err = s.db.Exec(query, stats.TotalBattles, stats.Wins, stats.Losses, stats.WinRate,
		time.Now(), time.Now(), playerID)

	return err
}

// Implement Value interface for JSON serialization
func (bs *BattleSession) Value() (driver.Value, error) {
	return json.Marshal(bs)
}
