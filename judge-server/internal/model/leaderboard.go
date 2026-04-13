package model

import "time"

type LeaderboardEntry struct {
	ID           string    `json:"id"`
	Player       string    `json:"player"`
	PetName      string    `json:"pet_name"`
	Level        int       `json:"level"`
	TotalBattles int       `json:"total_battles"`
	Wins         int       `json:"wins"`
	Losses       int       `json:"losses"`
	Rating       int       `json:"rating"`
	Rank         int       `json:"rank"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DailyLeaderboard struct {
	Date      string             `json:"date"`
	Entries   []LeaderboardEntry `json:"entries"`
	Generated time.Time          `json:"generated_at"`
}

type GrowthRecord struct {
	ID        string    `json:"id"`
	PetID     string    `json:"pet_id"`
	OldLevel  int       `json:"old_level"`
	NewLevel  int       `json:"new_level"`
	OldEXP    int       `json:"old_exp"`
	NewEXP    int       `json:"new_exp"`
	Source    string    `json:"source"`
	Timestamp time.Time `json:"timestamp"`
	IsValid   bool      `json:"is_valid"`
}
