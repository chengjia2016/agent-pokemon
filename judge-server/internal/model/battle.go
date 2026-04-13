package model

import "time"

type Battle struct {
	ID            string    `json:"id"`
	AttackerID    string    `json:"attacker_id"`
	DefenderID    string    `json:"defender_id"`
	AttackerName  string    `json:"attacker_name"`
	DefenderName  string    `json:"defender_name"`
	Winner        string    `json:"winner"`
	Turns         int       `json:"turns"`
	AttackStack   []string  `json:"attack_stack"`
	DefenseStack  []string  `json:"defense_stack"`
	BattleLog     []string  `json:"battle_log"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	IsValid       bool      `json:"is_valid"`
	ValidationMsg string    `json:"validation_msg"`
}

type BattleValidationResult struct {
	IsValid bool     `json:"is_valid"`
	Errors  []string `json:"errors"`
}

type BattleResult struct {
	Winner    string   `json:"winner"`
	Turns     int      `json:"turns"`
	BattleLog []string `json:"battle_log"`
}
