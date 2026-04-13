package service

import (
	"judge-server/internal/model"
	"math"
	"math/rand"
	"time"
)

type EggService struct {
	pokemonPool          []PokemonInfo
	minIncubationSeconds int64
	maxIncubationSeconds int64
}

type PokemonInfo struct {
	ID        string
	Name      string
	Rarity    string
	Types     []string
	Stats     map[string]int
	GeneType  string
	MinEnergy int
	MaxEnergy int
}

func NewEggService() *EggService {
	return &EggService{
		pokemonPool: []PokemonInfo{
			{"001", "Bulbasaur", "Common", []string{"Grass", "Poison"}, map[string]int{"hp": 45, "attack": 49, "defense": 49, "speed": 45, "sp_atk": 65, "sp_def": 65}, "logic", 0, 100},
			{"004", "Charmander", "Common", []string{"Fire"}, map[string]int{"hp": 39, "attack": 52, "defense": 43, "speed": 65, "sp_atk": 60, "sp_def": 50}, "creative", 0, 100},
			{"007", "Squirtle", "Common", []string{"Water"}, map[string]int{"hp": 44, "attack": 48, "defense": 65, "speed": 43, "sp_atk": 50, "sp_def": 64}, "logic", 0, 100},
			{"025", "Pikachu", "Rare", []string{"Electric"}, map[string]int{"hp": 35, "attack": 55, "defense": 40, "speed": 90, "sp_atk": 50, "sp_def": 50}, "speed", 50, 200},
			{"006", "Charizard", "Legendary", []string{"Fire", "Flying"}, map[string]int{"hp": 78, "attack": 84, "defense": 78, "speed": 100, "sp_atk": 109, "sp_def": 85}, "creative", 200, 500},
			{"009", "Blastoise", "Legendary", []string{"Water"}, map[string]int{"hp": 79, "attack": 83, "defense": 100, "speed": 78, "sp_atk": 85, "sp_def": 105}, "logic", 200, 500},
			{"003", "Venusaur", "Epic", []string{"Grass", "Poison"}, map[string]int{"hp": 80, "attack": 82, "defense": 83, "speed": 80, "sp_atk": 100, "sp_def": 100}, "logic", 100, 300},
			{"035", "Clefairy", "Rare", []string{"Fairy"}, map[string]int{"hp": 70, "attack": 45, "defense": 48, "speed": 35, "sp_atk": 60, "sp_def": 65}, "lucky", 50, 150},
			{"037", "Vulpix", "Uncommon", []string{"Fire"}, map[string]int{"hp": 38, "attack": 41, "defense": 40, "speed": 65, "sp_atk": 50, "sp_def": 65}, "creative", 20, 80},
			{"043", "Oddish", "Common", []string{"Grass", "Poison"}, map[string]int{"hp": 45, "attack": 50, "defense": 55, "speed": 30, "sp_atk": 75, "sp_def": 65}, "logic", 0, 60},
		},
		minIncubationSeconds: 60,
		maxIncubationSeconds: 259200, // 72 hours
	}
}

func (s *EggService) ValidateIncubation(egg *model.Egg) *model.EggValidationResult {
	result := &model.EggValidationResult{
		IsValid:  true,
		Errors:   []string{},
		Warnings: []string{},
	}

	if egg.ID == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing egg ID")
	}

	if egg.Owner == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing egg owner")
	}

	if egg.StartTime.IsZero() {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing start time")
	}

	elapsed := time.Since(egg.StartTime)
	egg.IncubationDuration = int64(elapsed.Seconds())

	if egg.IncubationDuration < s.minIncubationSeconds {
		result.Warnings = append(result.Warnings, "incubation too short, may affect results")
	}

	if egg.IncubationDuration > s.maxIncubationSeconds {
		result.Warnings = append(result.Warnings, "incubation exceeded max time, capping energy")
	}

	energy := s.calculateEnergy(elapsed)
	egg.EnergyAbsorbed = energy

	rarity := s.determineRarity(energy)
	pokemon := s.selectPokemon(rarity)

	geneModifiers := s.calculateGeneChanges(energy, pokemon)

	result.HatchResult = &model.HatchResult{
		PokemonID:     pokemon.ID,
		PokemonName:   pokemon.Name,
		GeneModifiers: geneModifiers,
		TotalEnergy:   energy,
		Rarity:        rarity,
	}

	return result
}

func (s *EggService) calculateEnergy(elapsed time.Duration) int {
	seconds := int64(elapsed.Seconds())
	if seconds <= 0 {
		return 10
	}
	if seconds > 259200 {
		seconds = 259200
	}
	energy := int(math.Sqrt(float64(seconds)) * 10)
	if energy > 500 {
		energy = 500
	}
	return energy
}

func (s *EggService) determineRarity(energy int) string {
	rand.Seed(time.Now().UnixNano())
	roll := rand.Float64()

	if energy >= 200 {
		if roll < 0.05 {
			return "Legendary"
		}
		if roll < 0.15 {
			return "Epic"
		}
		if roll < 0.30 {
			return "Rare"
		}
		return "Uncommon"
	}

	if energy >= 100 {
		if roll < 0.03 {
			return "Epic"
		}
		if roll < 0.12 {
			return "Rare"
		}
		if roll < 0.35 {
			return "Uncommon"
		}
		return "Common"
	}

	if energy >= 50 {
		if roll < 0.05 {
			return "Rare"
		}
		if roll < 0.20 {
			return "Uncommon"
		}
		return "Common"
	}

	if roll < 0.10 {
		return "Uncommon"
	}
	return "Common"
}

func (s *EggService) selectPokemon(rarity string) PokemonInfo {
	var candidates []PokemonInfo
	for _, p := range s.pokemonPool {
		if p.Rarity == rarity {
			candidates = append(candidates, p)
		}
	}

	if len(candidates) == 0 {
		candidates = s.pokemonPool[:3]
	}

	rand.Seed(time.Now().UnixNano())
	return candidates[rand.Intn(len(candidates))]
}

func (s *EggService) calculateGeneChanges(energy int, pokemon PokemonInfo) []model.GeneChange {
	rand.Seed(time.Now().UnixNano())

	modifier := float64(energy) / 500.0
	if modifier > 1.0 {
		modifier = 1.0
	}
	if modifier < 0.1 {
		modifier = 0.1
	}

	changes := []model.GeneChange{
		{
			GeneType:  pokemon.GeneType,
			OldWeight: 0.33,
			NewWeight: 0.33 + modifier*0.4,
		},
		{
			GeneType:  "lucky",
			OldWeight: 0.33,
			NewWeight: 0.33 - modifier*0.1,
		},
	}

	return changes
}

type CaptureService struct {
	baseCaptureRate float64
	bonusPerLowHP   float64
	minHPThreshold  float64
}

func NewCaptureService() *CaptureService {
	return &CaptureService{
		baseCaptureRate: 0.3,
		bonusPerLowHP:   0.5,
		minHPThreshold:  0.3,
	}
}

func (s *CaptureService) ValidateCapture(capture *model.CaptureAttempt) *model.CaptureResult {
	result := &model.CaptureResult{
		IsValid:     true,
		Success:     false,
		Errors:      []string{},
		CapturedPet: nil,
	}

	if capture.ID == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing capture ID")
	}

	if capture.BattleID == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing battle ID")
	}

	if capture.TargetID == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing target ID")
	}

	if capture.TargetHP <= 0 {
		result.IsValid = false
		result.Errors = append(result.Errors, "target already defeated")
	}

	hpPercent := capture.TargetHP / capture.MaxHP

	if hpPercent > s.minHPThreshold {
		result.Errors = append(result.Errors, "target HP too high, capture difficult")
		result.CaptureRate = s.baseCaptureRate * hpPercent
		return result
	}

	lowHPBonus := (s.minHPThreshold - hpPercent) * s.bonusPerLowHP
	captureRate := s.baseCaptureRate + lowHPBonus

	if captureRate > 0.9 {
		captureRate = 0.9
	}

	rand.Seed(time.Now().UnixNano())
	roll := rand.Float64()

	result.Success = roll < captureRate
	result.CaptureRate = captureRate

	if result.Success {
		result.CapturedPet = &model.Pet{
			ID:    "captured_" + capture.TargetID,
			Name:  capture.TargetID,
			Owner: capture.AttackerID,
			Level: 1,
			EXP:   0,
		}
	}

	return result
}
