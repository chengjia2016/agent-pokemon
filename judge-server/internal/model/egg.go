package model

import "time"

type Egg struct {
	ID                 string       `json:"id"`
	Owner              string       `json:"owner"`
	StartTime          time.Time    `json:"start_time"`
	HatchTime          time.Time    `json:"hatch_time"`
	Status             string       `json:"status"` // "incubating", "hatched", "claimed"
	IncubationDuration int64        `json:"incubation_duration_seconds"`
	EnergyAbsorbed     int          `json:"energy_absorbed"`
	GeneChanges        []GeneChange `json:"gene_changes"`
}

type GeneChange struct {
	GeneType  string  `json:"gene_type"`
	OldWeight float64 `json:"old_weight"`
	NewWeight float64 `json:"new_weight"`
}

type EggValidationResult struct {
	IsValid     bool         `json:"is_valid"`
	Errors      []string     `json:"errors"`
	Warnings    []string     `json:"warnings"`
	HatchResult *HatchResult `json:"hatch_result"`
}

type HatchResult struct {
	PokemonID     string       `json:"pokemon_id"`
	PokemonName   string       `json:"pokemon_name"`
	GeneModifiers []GeneChange `json:"gene_modifiers"`
	TotalEnergy   int          `json:"total_energy"`
	Rarity        string       `json:"rarity"`
}

type CaptureAttempt struct {
	ID          string    `json:"id"`
	CaptureID   string    `json:"capture_id"`
	BattleID    string    `json:"battle_id"`
	AttackerID  string    `json:"attacker_id"`
	TargetID    string    `json:"target_id"`
	TargetHP    float64   `json:"target_hp"`
	MaxHP       float64   `json:"max_hp"`
	CaptureRate float64   `json:"capture_rate"`
	Success     bool      `json:"success"`
	ThrowTime   time.Time `json:"throw_time"`
}

type CaptureResult struct {
	IsValid     bool     `json:"is_valid"`
	Success     bool     `json:"success"`
	CaptureRate float64  `json:"capture_rate"`
	Errors      []string `json:"errors"`
	CapturedPet *Pet     `json:"captured_pet"`
}
