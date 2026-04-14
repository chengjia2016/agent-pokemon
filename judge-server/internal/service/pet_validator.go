package service

import (
	"judge-server/internal/model"
	"math"
	"time"
)

type PetValidator struct {
	maxLevel       int
	maxStatValue   int
	maxIVValue     int
	maxEVValue     int
	maxGeneWeight  float64
	validSpecies   []string
	validStatNames []string
}

func NewPetValidator() *PetValidator {
	return &PetValidator{
		maxLevel:       100,
		maxStatValue:   255,
		maxIVValue:     31,
		maxEVValue:     252,
		maxGeneWeight:  1.0,
		validSpecies:   []string{"Hybrid", "Duck", "Dragon", "Phoenix", "Wolf", "Cat", "Custom"},
		validStatNames: []string{"hp", "attack", "defense", "speed", "armor", "quota"},
	}
}

func (v *PetValidator) Validate(pet *model.Pet) *model.PetValidationResult {
	result := &model.PetValidationResult{
		IsValid:  true,
		Errors:   []string{},
		Warnings: []string{},
	}

	if pet.ID == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing monster_id")
	}

	if pet.Name == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing name")
	}

	if pet.Owner == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing owner")
	}

	if pet.Species != "" && !v.isValidSpecies(pet.Species) {
		result.Warnings = append(result.Warnings, "unknown species: "+pet.Species)
	}

	if pet.Generation < 1 {
		result.IsValid = false
		result.Errors = append(result.Errors, "generation must be >= 1")
	}

	if pet.EvolutionStage < 1 || pet.EvolutionStage > 3 {
		result.IsValid = false
		result.Errors = append(result.Errors, "evolution_stage must be between 1 and 3")
	}

	if pet.Level < 1 || pet.Level > v.maxLevel {
		result.IsValid = false
		result.Errors = append(result.Errors, "level must be between 1 and "+string(rune(v.maxLevel)))
	}

	if pet.EXP < 0 {
		result.IsValid = false
		result.Errors = append(result.Errors, "exp cannot be negative")
	}

	maxExp := v.calculateMaxExp(pet.Level)
	if pet.EXP > maxExp {
		result.IsValid = false
		result.Errors = append(result.Errors, "exp exceeds maximum for level "+string(rune(pet.Level)))
	}

	if err := v.validateStats(pet); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, err.Error())
	}

	if err := v.validateGenes(pet); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, err.Error())
	}

	if err := v.validateSignature(pet); err != nil {
		result.Warnings = append(result.Warnings, err.Error())
	}

	return result
}

func (v *PetValidator) isValidSpecies(species string) bool {
	for _, s := range v.validSpecies {
		if s == species {
			return true
		}
	}
	return false
}

func (v *PetValidator) calculateMaxExp(level int) int {
	return int(math.Pow(float64(level), 3) * 10)
}

func (v *PetValidator) validateStats(pet *model.Pet) error {
	if pet.BaseStats == nil {
		return nil
	}

	for statName, value := range pet.BaseStats {
		if !v.isValidStatName(statName) {
			return nil
		}
		if value < 1 || value > v.maxStatValue {
			return nil
		}
	}

	if pet.IVs != nil {
		for _, iv := range pet.IVs {
			if iv < 0 || iv > v.maxIVValue {
				return nil
			}
		}
	}

	if pet.EVs != nil {
		totalEV := 0
		for _, ev := range pet.EVs {
			if ev < 0 || ev > v.maxEVValue {
				return nil
			}
			totalEV += ev
		}
		if totalEV > 510 {
			return nil
		}
	}

	return nil
}

func (v *PetValidator) isValidStatName(name string) bool {
	for _, s := range v.validStatNames {
		if s == name {
			return true
		}
	}
	return false
}

func (v *PetValidator) validateGenes(pet *model.Pet) error {
	if pet.Genes == nil {
		return nil
	}

	totalWeight := 0.0
	for _, gene := range pet.Genes {
		if gene.Weight < 0 || gene.Weight > v.maxGeneWeight {
			return nil
		}
		totalWeight += gene.Weight
	}

	if totalWeight > 1.01 {
		return nil
	}

	return nil
}

func (v *PetValidator) validateSignature(pet *model.Pet) error {
	if pet.Signature.Algorithm == "" {
		return nil
	}

	if pet.Signature.Value == "" {
		return nil
	}

	return nil
}

func CalculateLevel(exp int) int {
	level := 1
	for {
		maxExp := int(math.Pow(float64(level+1), 3) * 10)
		if exp < maxExp {
			return level
		}
		level++
		if level > 100 {
			return 100
		}
	}
}

func CalculateStat(base, iv, ev, level int) int {
	hp := base*2 + iv + ev/4
	stat := (hp * level / 100) + 5
	return stat
}

func CreateGrowthRecord(petID string, oldLevel, newLevel, oldExp, newExp int, source string) *model.GrowthRecord {
	return &model.GrowthRecord{
		ID:        "gr_" + time.Now().Format("20060102150405"),
		PetID:     petID,
		OldLevel:  oldLevel,
		NewLevel:  newLevel,
		OldEXP:    oldExp,
		NewEXP:    newExp,
		Source:    source,
		Timestamp: time.Now(),
		IsValid:   true,
	}
}
