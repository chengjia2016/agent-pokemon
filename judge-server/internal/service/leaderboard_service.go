package service

import (
	"judge-server/internal/db"
	"judge-server/internal/model"
	"sort"
	"time"
)

type LeaderboardService struct {
	DB *db.Database
}

func NewLeaderboardService(database *db.Database) *LeaderboardService {
	return &LeaderboardService{DB: database}
}

func (s *LeaderboardService) GenerateDailyLeaderboard(date string) (*model.DailyLeaderboard, error) {
	entries, err := s.DB.GetLeaderboard(100)
	if err != nil {
		return nil, err
	}

	for i := range entries {
		entries[i].Rank = i + 1
	}

	daily := &model.DailyLeaderboard{
		Date:      date,
		Entries:   entries,
		Generated: time.Now(),
	}

	if err := s.DB.SaveDailyLeaderboard(daily); err != nil {
		return nil, err
	}

	return daily, nil
}

func (s *LeaderboardService) UpdateLeaderboard(pet *model.Pet, battleResult string) error {
	entry, err := s.getOrCreateEntry(pet)
	if err != nil {
		return err
	}

	entry.TotalBattles++
	if battleResult == "win" {
		entry.Wins++
		entry.Rating += 15
	} else if battleResult == "lose" {
		entry.Losses++
		entry.Rating = max(0, entry.Rating-10)
	}

	entry.Level = pet.Level
	entry.UpdatedAt = time.Now()

	return s.saveEntry(entry)
}

func (s *LeaderboardService) getOrCreateEntry(pet *model.Pet) (*model.LeaderboardEntry, error) {
	entries, err := s.DB.GetLeaderboard(1000)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if e.Player == pet.Owner {
			return &e, nil
		}
	}

	newEntry := &model.LeaderboardEntry{
		ID:           "lb_" + pet.Owner + "_" + time.Now().Format("20060102"),
		Player:       pet.Owner,
		PetName:      pet.Name,
		Level:        pet.Level,
		TotalBattles: 0,
		Wins:         0,
		Losses:       0,
		Rating:       1000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return newEntry, nil
}

func (s *LeaderboardService) saveEntry(entry *model.LeaderboardEntry) error {
	query := `
		INSERT INTO leaderboard (id, player, pet_name, level, total_battles, wins, losses, rating, rank, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (id) DO UPDATE SET
			pet_name = $3, level = $4, total_battles = $5, wins = $6, losses = $7,
			rating = $8, rank = $9, updated_at = $11
	`
	_, err := s.DB.Exec(query,
		entry.ID, entry.Player, entry.PetName, entry.Level, entry.TotalBattles,
		entry.Wins, entry.Losses, entry.Rating, entry.Rank, entry.CreatedAt, entry.UpdatedAt)
	return err
}

func (s *LeaderboardService) GetTopPlayers(limit int) ([]model.LeaderboardEntry, error) {
	entries, err := s.DB.GetLeaderboard(limit)
	if err != nil {
		return nil, err
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Rating > entries[j].Rating
	})

	for i := range entries {
		entries[i].Rank = i + 1
	}

	return entries, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
