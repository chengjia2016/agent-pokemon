package service

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// TournamentSeason represents a tournament season
type TournamentSeason struct {
	ID            int             `json:"id"`
	SeasonID      string          `json:"season_id"`
	SeasonName    string          `json:"season_name"`
	SeasonNameZh  *string         `json:"season_name_zh,omitempty"`
	SeasonNumber  int             `json:"season_number"`
	Description   *string         `json:"description,omitempty"`
	DescriptionZh *string         `json:"description_zh,omitempty"`
	StartDate     time.Time       `json:"start_date"`
	EndDate       time.Time       `json:"end_date"`
	StartRank     int             `json:"start_rank"`
	MaxRankPoints int             `json:"max_rank_points"`
	IsActive      bool            `json:"is_active"`
	RewardPool    json.RawMessage `json:"reward_pool,omitempty"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

// TournamentRanking represents a player's ranking in a tournament season
type TournamentRanking struct {
	ID             int        `json:"id"`
	RankingID      string     `json:"ranking_id"`
	SeasonID       string     `json:"season_id"`
	PlayerID       string     `json:"player_id"`
	Rank           *int       `json:"rank,omitempty"`
	RankPoints     int        `json:"rank_points"`
	Wins           int        `json:"wins"`
	Losses         int        `json:"losses"`
	WinStreak      int        `json:"win_streak"`
	LossStreak     int        `json:"loss_streak"`
	MatchesPlayed  int        `json:"matches_played"`
	HighestRank    *int       `json:"highest_rank,omitempty"`
	HighestPoints  int        `json:"highest_points"`
	LastMatchAt    *time.Time `json:"last_match_at,omitempty"`
	LastRankUpdate *time.Time `json:"last_rank_update,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// TournamentReward represents a reward earned in tournament
type TournamentReward struct {
	ID           int        `json:"id"`
	RewardID     string     `json:"reward_id"`
	SeasonID     string     `json:"season_id"`
	PlayerID     string     `json:"player_id"`
	RewardType   string     `json:"reward_type"` // gems, coins, items, pokemon
	RewardName   *string    `json:"reward_name,omitempty"`
	RewardAmount int        `json:"reward_amount"`
	RankEarnedAt int        `json:"rank_earned_at"`
	Claimed      bool       `json:"claimed"`
	ClaimedAt    *time.Time `json:"claimed_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
}

// TournamentService handles tournament operations for Masters EX
type TournamentService struct {
	db *sql.DB
}

// NewTournamentService creates a new tournament service
func NewTournamentService(db *sql.DB) *TournamentService {
	return &TournamentService{
		db: db,
	}
}

// CreateTournamentSeason creates a new tournament season
func (s *TournamentService) CreateTournamentSeason(seasonName string, seasonNameZh *string,
	seasonNumber int, description, descriptionZh *string,
	startDate, endDate time.Time, startRank, maxRankPoints int, rewardPool json.RawMessage) (*TournamentSeason, error) {

	seasonID := fmt.Sprintf("season_%d_%d", seasonNumber, time.Now().Unix())

	season := &TournamentSeason{
		SeasonID:      seasonID,
		SeasonName:    seasonName,
		SeasonNameZh:  seasonNameZh,
		SeasonNumber:  seasonNumber,
		Description:   description,
		DescriptionZh: descriptionZh,
		StartDate:     startDate,
		EndDate:       endDate,
		StartRank:     startRank,
		MaxRankPoints: maxRankPoints,
		IsActive:      false,
		RewardPool:    rewardPool,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	query := `
		INSERT INTO tournament_seasons 
		(season_id, season_name, season_name_zh, season_number, description, description_zh,
		 start_date, end_date, start_rank, max_rank_points, is_active, reward_pool, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id, created_at, updated_at`

	err := s.db.QueryRow(query, season.SeasonID, season.SeasonName, season.SeasonNameZh, season.SeasonNumber,
		season.Description, season.DescriptionZh, season.StartDate, season.EndDate, season.StartRank,
		season.MaxRankPoints, season.IsActive, season.RewardPool, season.CreatedAt, season.UpdatedAt).Scan(
		&season.ID, &season.CreatedAt, &season.UpdatedAt)

	if err != nil {
		log.Printf("Error creating tournament season: %v", err)
		return nil, err
	}

	return season, nil
}

// GetTournamentSeason retrieves a tournament season by ID
func (s *TournamentService) GetTournamentSeason(seasonID string) (*TournamentSeason, error) {
	season := &TournamentSeason{}

	query := `
		SELECT id, season_id, season_name, season_name_zh, season_number, description, description_zh,
		       start_date, end_date, start_rank, max_rank_points, is_active, reward_pool, created_at, updated_at
		FROM tournament_seasons WHERE season_id = $1`

	var seasonNameZh, description, descriptionZh sql.NullString
	var rewardPool sql.NullString

	err := s.db.QueryRow(query, seasonID).Scan(
		&season.ID, &season.SeasonID, &season.SeasonName, &seasonNameZh, &season.SeasonNumber,
		&description, &descriptionZh, &season.StartDate, &season.EndDate, &season.StartRank,
		&season.MaxRankPoints, &season.IsActive, &rewardPool, &season.CreatedAt, &season.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tournament season not found: %s", seasonID)
		}
		log.Printf("Error getting tournament season: %v", err)
		return nil, err
	}

	if seasonNameZh.Valid {
		season.SeasonNameZh = &seasonNameZh.String
	}
	if description.Valid {
		season.Description = &description.String
	}
	if descriptionZh.Valid {
		season.DescriptionZh = &descriptionZh.String
	}
	if rewardPool.Valid {
		season.RewardPool = json.RawMessage(rewardPool.String)
	}

	return season, nil
}

// ActivateTournamentSeason activates a tournament season
func (s *TournamentService) ActivateTournamentSeason(seasonID string) error {
	query := `UPDATE tournament_seasons SET is_active = true, updated_at = $1 WHERE season_id = $2`

	_, err := s.db.Exec(query, time.Now(), seasonID)
	if err != nil {
		log.Printf("Error activating tournament season: %v", err)
		return err
	}

	return nil
}

// RegisterPlayerInTournament registers a player in a tournament season
func (s *TournamentService) RegisterPlayerInTournament(seasonID, playerID string) (*TournamentRanking, error) {
	rankingID := fmt.Sprintf("ranking_%s_%s_%d", seasonID, playerID, time.Now().Unix())

	season, err := s.GetTournamentSeason(seasonID)
	if err != nil {
		return nil, err
	}

	ranking := &TournamentRanking{
		RankingID:     rankingID,
		SeasonID:      seasonID,
		PlayerID:      playerID,
		RankPoints:    season.StartRank,
		HighestPoints: season.StartRank,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	query := `
		INSERT INTO tournament_rankings
		(ranking_id, season_id, player_id, rank_points, highest_points, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at`

	err = s.db.QueryRow(query, ranking.RankingID, ranking.SeasonID, ranking.PlayerID,
		ranking.RankPoints, ranking.HighestPoints, ranking.CreatedAt, ranking.UpdatedAt).Scan(
		&ranking.ID, &ranking.CreatedAt, &ranking.UpdatedAt)

	if err != nil {
		log.Printf("Error registering player in tournament: %v", err)
		return nil, err
	}

	return ranking, nil
}

// GetPlayerRanking gets a player's ranking in a tournament season
func (s *TournamentService) GetPlayerRanking(seasonID, playerID string) (*TournamentRanking, error) {
	ranking := &TournamentRanking{}

	query := `
		SELECT id, ranking_id, season_id, player_id, rank, rank_points, wins, losses,
		       win_streak, loss_streak, matches_played, highest_rank, highest_points,
		       last_match_at, last_rank_update, created_at, updated_at
		FROM tournament_rankings 
		WHERE season_id = $1 AND player_id = $2`

	var rank, highestRank sql.NullInt64
	var lastMatchAt, lastRankUpdate sql.NullTime

	err := s.db.QueryRow(query, seasonID, playerID).Scan(
		&ranking.ID, &ranking.RankingID, &ranking.SeasonID, &ranking.PlayerID, &rank,
		&ranking.RankPoints, &ranking.Wins, &ranking.Losses, &ranking.WinStreak, &ranking.LossStreak,
		&ranking.MatchesPlayed, &highestRank, &ranking.HighestPoints, &lastMatchAt, &lastRankUpdate,
		&ranking.CreatedAt, &ranking.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("player not registered in this tournament season")
		}
		log.Printf("Error getting player ranking: %v", err)
		return nil, err
	}

	if rank.Valid {
		rankInt := int(rank.Int64)
		ranking.Rank = &rankInt
	}
	if highestRank.Valid {
		highestInt := int(highestRank.Int64)
		ranking.HighestRank = &highestInt
	}
	if lastMatchAt.Valid {
		ranking.LastMatchAt = &lastMatchAt.Time
	}
	if lastRankUpdate.Valid {
		ranking.LastRankUpdate = &lastRankUpdate.Time
	}

	return ranking, nil
}

// GetSeasonLeaderboard gets the top N players in a tournament season
func (s *TournamentService) GetSeasonLeaderboard(seasonID string, limit int) ([]*TournamentRanking, error) {
	query := `
		SELECT id, ranking_id, season_id, player_id, rank, rank_points, wins, losses,
		       win_streak, loss_streak, matches_played, highest_rank, highest_points,
		       last_match_at, last_rank_update, created_at, updated_at
		FROM tournament_rankings 
		WHERE season_id = $1
		ORDER BY rank_points DESC, rank ASC
		LIMIT $2`

	rows, err := s.db.Query(query, seasonID, limit)
	if err != nil {
		log.Printf("Error getting season leaderboard: %v", err)
		return nil, err
	}
	defer rows.Close()

	var rankings []*TournamentRanking
	for rows.Next() {
		ranking := &TournamentRanking{}
		var rank, highestRank sql.NullInt64
		var lastMatchAt, lastRankUpdate sql.NullTime

		err := rows.Scan(
			&ranking.ID, &ranking.RankingID, &ranking.SeasonID, &ranking.PlayerID, &rank,
			&ranking.RankPoints, &ranking.Wins, &ranking.Losses, &ranking.WinStreak, &ranking.LossStreak,
			&ranking.MatchesPlayed, &highestRank, &ranking.HighestPoints, &lastMatchAt, &lastRankUpdate,
			&ranking.CreatedAt, &ranking.UpdatedAt)

		if err != nil {
			log.Printf("Error scanning leaderboard row: %v", err)
			continue
		}

		if rank.Valid {
			rankInt := int(rank.Int64)
			ranking.Rank = &rankInt
		}
		if highestRank.Valid {
			highestInt := int(highestRank.Int64)
			ranking.HighestRank = &highestInt
		}
		if lastMatchAt.Valid {
			ranking.LastMatchAt = &lastMatchAt.Time
		}
		if lastRankUpdate.Valid {
			ranking.LastRankUpdate = &lastRankUpdate.Time
		}

		rankings = append(rankings, ranking)
	}

	return rankings, nil
}

// UpdatePlayerTournamentStats updates player's tournament statistics after a match
func (s *TournamentService) UpdatePlayerTournamentStats(seasonID, playerID string,
	pointsChange int, isWin bool) (*TournamentRanking, error) {

	ranking, err := s.GetPlayerRanking(seasonID, playerID)
	if err != nil {
		return nil, err
	}

	ranking.RankPoints += pointsChange
	ranking.MatchesPlayed++

	if isWin {
		ranking.Wins++
		ranking.WinStreak++
		ranking.LossStreak = 0
		if ranking.HighestPoints < ranking.RankPoints {
			ranking.HighestPoints = ranking.RankPoints
		}
	} else {
		ranking.Losses++
		ranking.LossStreak++
		ranking.WinStreak = 0
	}

	now := time.Now()
	ranking.LastMatchAt = &now
	ranking.LastRankUpdate = &now

	query := `
		UPDATE tournament_rankings
		SET rank_points = $1, wins = $2, losses = $3, win_streak = $4, loss_streak = $5,
		    matches_played = $6, highest_points = $7, last_match_at = $8, last_rank_update = $9, updated_at = $10
		WHERE season_id = $11 AND player_id = $12`

	_, err = s.db.Exec(query, ranking.RankPoints, ranking.Wins, ranking.Losses, ranking.WinStreak,
		ranking.LossStreak, ranking.MatchesPlayed, ranking.HighestPoints, ranking.LastMatchAt,
		ranking.LastRankUpdate, time.Now(), seasonID, playerID)

	if err != nil {
		log.Printf("Error updating player tournament stats: %v", err)
		return nil, err
	}

	return ranking, nil
}

// AwardTournamentReward awards a reward to a player
func (s *TournamentService) AwardTournamentReward(seasonID, playerID, rewardType string,
	rewardName *string, rewardAmount, rankEarnedAt int) (*TournamentReward, error) {

	rewardID := fmt.Sprintf("reward_%s_%s_%d", seasonID, playerID, time.Now().Unix())

	reward := &TournamentReward{
		RewardID:     rewardID,
		SeasonID:     seasonID,
		PlayerID:     playerID,
		RewardType:   rewardType,
		RewardName:   rewardName,
		RewardAmount: rewardAmount,
		RankEarnedAt: rankEarnedAt,
		CreatedAt:    time.Now(),
	}

	query := `
		INSERT INTO tournament_rewards
		(reward_id, season_id, player_id, reward_type, reward_name, reward_amount, rank_earned_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	err := s.db.QueryRow(query, reward.RewardID, reward.SeasonID, reward.PlayerID, reward.RewardType,
		reward.RewardName, reward.RewardAmount, reward.RankEarnedAt, reward.CreatedAt).Scan(&reward.ID)

	if err != nil {
		log.Printf("Error awarding tournament reward: %v", err)
		return nil, err
	}

	return reward, nil
}

// ClaimTournamentReward claims a tournament reward
func (s *TournamentService) ClaimTournamentReward(rewardID string) (*TournamentReward, error) {
	reward := &TournamentReward{}
	now := time.Now()

	query := `
		UPDATE tournament_rewards
		SET claimed = true, claimed_at = $1
		WHERE reward_id = $2
		RETURNING id, reward_id, season_id, player_id, reward_type, reward_name, reward_amount,
		          rank_earned_at, claimed, claimed_at, created_at`

	var rewardName sql.NullString
	var claimedAt sql.NullTime

	err := s.db.QueryRow(query, now, rewardID).Scan(
		&reward.ID, &reward.RewardID, &reward.SeasonID, &reward.PlayerID, &reward.RewardType,
		&rewardName, &reward.RewardAmount, &reward.RankEarnedAt, &reward.Claimed, &claimedAt, &reward.CreatedAt)

	if err != nil {
		log.Printf("Error claiming tournament reward: %v", err)
		return nil, err
	}

	if rewardName.Valid {
		reward.RewardName = &rewardName.String
	}
	if claimedAt.Valid {
		reward.ClaimedAt = &claimedAt.Time
	}

	return reward, nil
}

// Implement Value interface for JSON serialization
func (ts *TournamentSeason) Value() (driver.Value, error) {
	return json.Marshal(ts)
}

func (tr *TournamentRanking) Value() (driver.Value, error) {
	return json.Marshal(tr)
}

func (tr *TournamentReward) Value() (driver.Value, error) {
	return json.Marshal(tr)
}
