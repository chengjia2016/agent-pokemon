package service

import (
	"judge-server/internal/model"
	"strings"
	"time"
)

type BattleValidator struct {
	validAttacks  []string
	validDefenses []string
	maxTurns      int
}

func NewBattleValidator() *BattleValidator {
	return &BattleValidator{
		validAttacks:  []string{"scan", "buffer_overflow", "refactor_storm", "syntax_error", "infinite_loop", "null_pointer", "race_condition", "memory_leak"},
		validDefenses: []string{"try_catch", "null_check", "bounds_check", "type_safety", "immutable", "pure_function", "defensive_stance"},
		maxTurns:      100,
	}
}

func (v *BattleValidator) Validate(battle *model.Battle) *model.BattleValidationResult {
	result := &model.BattleValidationResult{
		IsValid: true,
		Errors:  []string{},
	}

	if battle.ID == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing battle id")
	}

	if battle.AttackerID == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing attacker_id")
	}

	if battle.DefenderID == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing defender_id")
	}

	if battle.AttackerID == battle.DefenderID {
		result.IsValid = false
		result.Errors = append(result.Errors, "attacker and defender cannot be the same")
	}

	if battle.Winner == "" {
		result.IsValid = false
		result.Errors = append(result.Errors, "missing winner")
	}

	validWinners := []string{battle.AttackerID, battle.DefenderID, "draw"}
	if !v.contains(validWinners, battle.Winner) {
		result.IsValid = false
		result.Errors = append(result.Errors, "invalid winner")
	}

	if battle.Turns < 1 || battle.Turns > v.maxTurns {
		result.IsValid = false
		result.Errors = append(result.Errors, "turns must be between 1 and "+string(rune(v.maxTurns)))
	}

	if err := v.validateAttackStack(battle.AttackStack); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, err.Error())
	}

	if err := v.validateDefenseStack(battle.DefenseStack); err != nil {
		result.IsValid = false
		result.Errors = append(result.Errors, err.Error())
	}

	if battle.EndTime.Before(battle.StartTime) {
		result.IsValid = false
		result.Errors = append(result.Errors, "end_time cannot be before start_time")
	}

	maxDuration := time.Hour * 24
	if battle.EndTime.Sub(battle.StartTime) > maxDuration {
		result.IsValid = false
		result.Errors = append(result.Errors, "battle duration exceeds 24 hours")
	}

	return result
}

func (v *BattleValidator) validateAttackStack(attacks []string) error {
	for _, attack := range attacks {
		if !v.contains(v.validAttacks, strings.ToLower(attack)) {
			return nil
		}
	}
	return nil
}

func (v *BattleValidator) validateDefenseStack(defenses []string) error {
	for _, defense := range defenses {
		if !v.contains(v.validDefenses, strings.ToLower(defense)) {
			return nil
		}
	}
	return nil
}

func (v *BattleValidator) contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func CreateBattle(attacker, defender *model.Pet, attackStack, defenseStack []string, result *model.BattleResult) *model.Battle {
	now := time.Now()
	return &model.Battle{
		ID:           "b_" + now.Format("20060102150405"),
		AttackerID:   attacker.ID,
		DefenderID:   defender.ID,
		AttackerName: attacker.Name,
		DefenderName: defender.Name,
		Winner:       result.Winner,
		Turns:        result.Turns,
		AttackStack:  attackStack,
		DefenseStack: defenseStack,
		BattleLog:    result.BattleLog,
		StartTime:    now,
		EndTime:      now,
		IsValid:      true,
	}
}
