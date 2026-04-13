package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"judge-server/internal/model"
)

// ===== Egg Operations =====

// CreateEgg creates a new egg
func (d *Database) CreateEgg(egg *model.DatabaseEgg) (*model.DatabaseEgg, error) {
	query := `
		INSERT INTO eggs (egg_id, owner_id, incubation_hours, stage, attributes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, egg_id, owner_id, created_at, incubation_hours, stage, hatched_at, hatched_pet_id, attributes
	`

	attrsJSON, _ := json.Marshal(map[string]interface{}{})
	if egg.Attributes != "" {
		attrsJSON = []byte(egg.Attributes)
	}

	err := d.conn.QueryRow(query,
		egg.EggID, egg.OwnerID, egg.IncubationHours, egg.Stage, string(attrsJSON)).
		Scan(&egg.ID, &egg.EggID, &egg.OwnerID, &egg.CreatedAt, &egg.IncubationHours,
			&egg.Stage, &egg.HatchedAt, &egg.HatchedPetID, &egg.Attributes)

	return egg, err
}

// GetEggByID retrieves an egg by ID
func (d *Database) GetEggByID(eggID string) (*model.DatabaseEgg, error) {
	query := `
		SELECT id, egg_id, owner_id, created_at, incubation_hours, stage, hatched_at, hatched_pet_id, attributes
		FROM eggs WHERE egg_id = $1
	`

	egg := &model.DatabaseEgg{}
	err := d.conn.QueryRow(query, eggID).Scan(
		&egg.ID, &egg.EggID, &egg.OwnerID, &egg.CreatedAt, &egg.IncubationHours,
		&egg.Stage, &egg.HatchedAt, &egg.HatchedPetID, &egg.Attributes)

	if err == sql.ErrNoRows {
		return nil, errors.New("egg not found")
	}

	return egg, err
}

// HatchEgg hatches an egg
func (d *Database) HatchEgg(eggID string, petID string) error {
	query := `
		UPDATE eggs 
		SET stage = 3, hatched_at = NOW(), hatched_pet_id = $2
		WHERE egg_id = $1
	`

	result, err := d.conn.Exec(query, eggID, petID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("egg not found")
	}

	return nil
}

// IncubateEgg updates egg incubation stage
func (d *Database) IncubateEgg(eggID string, newStage int) error {
	query := "UPDATE eggs SET stage = $1 WHERE egg_id = $2"
	result, err := d.conn.Exec(query, newStage, eggID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("egg not found")
	}

	return nil
}

// GetPlayerEggs retrieves all eggs for a player
func (d *Database) GetPlayerEggs(playerID string) ([]model.DatabaseEgg, error) {
	query := `
		SELECT id, egg_id, owner_id, created_at, incubation_hours, stage, hatched_at, hatched_pet_id, attributes
		FROM eggs
		WHERE owner_id = $1
		ORDER BY created_at DESC
	`

	rows, err := d.conn.Query(query, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eggs []model.DatabaseEgg
	for rows.Next() {
		egg := model.DatabaseEgg{}
		err := rows.Scan(
			&egg.ID, &egg.EggID, &egg.OwnerID, &egg.CreatedAt, &egg.IncubationHours,
			&egg.Stage, &egg.HatchedAt, &egg.HatchedPetID, &egg.Attributes)
		if err != nil {
			return nil, err
		}
		eggs = append(eggs, egg)
	}

	return eggs, rows.Err()
}

// GetEggStats retrieves egg statistics
func (d *Database) GetEggStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	var totalEggs, hatchedEggs int
	err := d.conn.QueryRow(`
		SELECT 
			COUNT(*) as total,
			COUNT(hatched_at) as hatched
		FROM eggs
	`).Scan(&totalEggs, &hatchedEggs)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	stats["total_eggs"] = totalEggs
	stats["hatched_eggs"] = hatchedEggs
	stats["unhatched_eggs"] = totalEggs - hatchedEggs

	return stats, nil
}
