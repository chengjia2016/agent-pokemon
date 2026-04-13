package model

import (
	"time"
)

type Pet struct {
	ID             string                `json:"monster_id"`
	Name           string                `json:"name"`
	Owner          string                `json:"owner"`
	Species        string                `json:"species"`
	Generation     int                   `json:"generation"`
	EvolutionStage int                   `json:"evolution_stage"`
	BaseStats      map[string]int        `json:"base_stats"`
	IVs            map[string]int        `json:"ivs"`
	EVs            map[string]int        `json:"evs"`
	EXP            int                   `json:"exp"`
	Level          int                   `json:"level"`
	Genes          map[string]GeneWeight `json:"genes"`
	BattleHistory  []BattleRecord        `json:"battle_history"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
	Signature      Signature             `json:"signature"`
	IsValid        bool                  `json:"is_valid"`
	ValidationMsg  string                `json:"validation_msg"`
}

type GeneWeight struct {
	Weight        float64  `json:"weight"`
	SourceCommits []string `json:"source_commits"`
}

type BattleRecord struct {
	Date      string `json:"date"`
	Opponent  string `json:"opponent"`
	Result    string `json:"result"`
	BattleLog string `json:"battle_log"`
}

type Signature struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
	KeyID     string `json:"keyid"`
}

type PetValidationResult struct {
	IsValid  bool     `json:"is_valid"`
	Errors   []string `json:"errors"`
	Warnings []string `json:"warnings"`
}
