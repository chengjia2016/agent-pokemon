package db

import (
	"database/sql"
	"errors"
	"fmt"
	"judge-server/internal/model"
)

// ===== Farm Operations =====

// CreateFarm creates a new farm in the database
func (d *Database) CreateFarm(farm *model.DatabaseFarm) (*model.DatabaseFarm, error) {
	farm.FarmKey = farm.OwnerID + "/" + farm.RepositoryName

	query := `
		INSERT INTO farms (farm_key, owner_id, repository_name, repository_url, planted_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, farm_key, owner_id, repository_name, repository_url, planted_at, created_at, updated_at
	`

	err := d.conn.QueryRow(query,
		farm.FarmKey, farm.OwnerID, farm.RepositoryName, farm.RepositoryURL).
		Scan(&farm.ID, &farm.FarmKey, &farm.OwnerID, &farm.RepositoryName, &farm.RepositoryURL,
			&farm.PlantedAt, &farm.CreatedAt, &farm.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return farm, nil
}

// GetFarmByID retrieves a farm by ID
func (d *Database) GetFarmByID(farmID int64) (*model.DatabaseFarm, error) {
	query := `
		SELECT id, farm_key, owner_id, repository_name, repository_url, planted_at, created_at, updated_at
		FROM farms
		WHERE id = $1
	`

	farm := &model.DatabaseFarm{}
	err := d.conn.QueryRow(query, farmID).Scan(
		&farm.ID, &farm.FarmKey, &farm.OwnerID, &farm.RepositoryName, &farm.RepositoryURL,
		&farm.PlantedAt, &farm.CreatedAt, &farm.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("farm not found")
	}

	return farm, err
}

// SearchFarms searches farms by owner or repository
func (d *Database) SearchFarms(owner string, repository string) ([]model.DatabaseFarm, error) {
	query := "SELECT id, farm_key, owner_id, repository_name, repository_url, planted_at, created_at, updated_at FROM farms WHERE 1=1"
	var args []interface{}
	argCount := 1

	if owner != "" {
		query += fmt.Sprintf(" AND owner_id ILIKE $%d", argCount)
		args = append(args, "%"+owner+"%")
		argCount++
	}

	if repository != "" {
		query += fmt.Sprintf(" AND repository_name ILIKE $%d", argCount)
		args = append(args, "%"+repository+"%")
		argCount++
	}

	query += " ORDER BY created_at DESC LIMIT 100"

	rows, err := d.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var farms []model.DatabaseFarm
	for rows.Next() {
		farm := model.DatabaseFarm{}
		err := rows.Scan(
			&farm.ID, &farm.FarmKey, &farm.OwnerID, &farm.RepositoryName, &farm.RepositoryURL,
			&farm.PlantedAt, &farm.CreatedAt, &farm.UpdatedAt)
		if err != nil {
			return nil, err
		}
		farms = append(farms, farm)
	}

	return farms, rows.Err()
}

// DeleteFarm deletes a farm
func (d *Database) DeleteFarm(farmID int64) error {
	query := "DELETE FROM farms WHERE id = $1"
	result, err := d.conn.Exec(query, farmID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("farm not found")
	}

	return nil
}

// ===== Farm Food Operations =====

// AddFoodToFarm adds food to a farm
func (d *Database) AddFoodToFarm(farmFood *model.DatabaseFarmFood) (*model.DatabaseFarmFood, error) {
	query := `
		INSERT INTO farm_foods (farm_id, food_id, food_type, emoji, current_quantity, max_quantity, regeneration_hours, seed_hash)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, farm_id, food_id, food_type, emoji, current_quantity, max_quantity, regeneration_hours, last_eaten_at, seed_hash, created_at, updated_at
	`

	err := d.conn.QueryRow(query,
		farmFood.FarmID, farmFood.FoodID, farmFood.FoodType, farmFood.Emoji,
		farmFood.CurrentQuantity, farmFood.MaxQuantity, farmFood.RegenerationHours, farmFood.SeedHash).
		Scan(&farmFood.ID, &farmFood.FarmID, &farmFood.FoodID, &farmFood.FoodType, &farmFood.Emoji,
			&farmFood.CurrentQuantity, &farmFood.MaxQuantity, &farmFood.RegenerationHours,
			&farmFood.LastEatenAt, &farmFood.SeedHash, &farmFood.CreatedAt, &farmFood.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return farmFood, nil
}

// GetFoodByID retrieves food information by ID
func (d *Database) GetFoodByID(foodID string) (*model.DatabaseFarmFood, error) {
	query := `
		SELECT id, farm_id, food_id, food_type, emoji, current_quantity, max_quantity, regeneration_hours, last_eaten_at, seed_hash, created_at, updated_at
		FROM farm_foods
		WHERE food_id = $1
	`

	food := &model.DatabaseFarmFood{}
	err := d.conn.QueryRow(query, foodID).Scan(
		&food.ID, &food.FarmID, &food.FoodID, &food.FoodType, &food.Emoji,
		&food.CurrentQuantity, &food.MaxQuantity, &food.RegenerationHours,
		&food.LastEatenAt, &food.SeedHash, &food.CreatedAt, &food.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("food not found")
	}

	return food, err
}

// ConsumeFood records food consumption
func (d *Database) ConsumeFood(farmID int64, foodID string, eaterID string, eaterPetID string) error {
	tx, err := d.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Update food quantity
	updateQuery := `
		UPDATE farm_foods 
		SET current_quantity = current_quantity - 1, last_eaten_at = NOW(), updated_at = NOW()
		WHERE food_id = $1 AND farm_id = $2 AND current_quantity > 0
	`
	result, err := tx.Exec(updateQuery, foodID, farmID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("food not found or out of stock")
	}

	// Record consumption in history
	historyQuery := `
		INSERT INTO food_consumption_history (food_id, eater_player_id, eater_pet_id, consumed_at, nutrition_exp, nutrition_energy)
		VALUES ($1, $2, $3, NOW(), $4, $5)
	`
	_, err = tx.Exec(historyQuery, foodID, eaterID, eaterPetID, 10, 5)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// GetFarmStatistics retrieves farm statistics
func (d *Database) GetFarmStatistics(farmID int64) (*model.FarmStatistics, error) {
	stats := &model.FarmStatistics{
		FarmID:        farmID,
		FoodBreakdown: make(map[string]int),
	}

	// Get total foods and quantities
	query := `
		SELECT 
			COUNT(*) as total_foods,
			COALESCE(SUM(current_quantity), 0) as total_quantity
		FROM farm_foods
		WHERE farm_id = $1
	`

	err := d.conn.QueryRow(query, farmID).Scan(&stats.TotalFoods, &stats.TotalQuantity)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Get consumption count
	consumeQuery := `
		SELECT COALESCE(COUNT(*), 0) as total_consumptions, MAX(consumed_at) as last_consumption
		FROM food_consumption_history
		WHERE food_id IN (SELECT food_id FROM farm_foods WHERE farm_id = $1)
	`

	err = d.conn.QueryRow(consumeQuery, farmID).Scan(&stats.TotalConsumptions, &stats.LastConsumption)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Get food breakdown by type
	breakdownQuery := `
		SELECT food_type, SUM(current_quantity) as qty
		FROM farm_foods
		WHERE farm_id = $1
		GROUP BY food_type
	`

	rows, err := d.conn.Query(breakdownQuery, farmID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var foodType string
		var qty int
		if err := rows.Scan(&foodType, &qty); err != nil {
			return nil, err
		}
		stats.FoodBreakdown[foodType] = qty
	}

	return stats, rows.Err()
}

// ===== Helper Functions =====

// Helper to convert string argument number for SQL placeholders
func argNum(n int) string {
	return "$" + string(rune(n+48))
}

// GetFarmFoods retrieves all foods in a farm
func (d *Database) GetFarmFoods(farmID int64) ([]model.DatabaseFarmFood, error) {
	query := `
		SELECT id, farm_id, food_id, food_type, emoji, current_quantity, max_quantity, regeneration_hours, last_eaten_at, seed_hash, created_at, updated_at
		FROM farm_foods
		WHERE farm_id = $1
		ORDER BY created_at DESC
	`

	rows, err := d.conn.Query(query, farmID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var foods []model.DatabaseFarmFood
	for rows.Next() {
		food := model.DatabaseFarmFood{}
		err := rows.Scan(
			&food.ID, &food.FarmID, &food.FoodID, &food.FoodType, &food.Emoji,
			&food.CurrentQuantity, &food.MaxQuantity, &food.RegenerationHours,
			&food.LastEatenAt, &food.SeedHash, &food.CreatedAt, &food.UpdatedAt)
		if err != nil {
			return nil, err
		}
		foods = append(foods, food)
	}

	return foods, rows.Err()
}
