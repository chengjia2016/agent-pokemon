package service

import (
	"judge-server/internal/db"
	"judge-server/internal/model"
	"time"
)

const (
	DailyEatLimit       = 5
	RepoVisitDailyLimit = 10
)

type FoodService struct {
	database *db.Database
}

func NewFoodService(database *db.Database) *FoodService {
	return &FoodService{
		database: database,
	}
}

// ValidateFood validates a food consumption request
func (s *FoodService) ValidateFood(req *model.FoodValidationRequest, farm *model.Farm) *model.FoodValidationResult {
	result := &model.FoodValidationResult{
		IsValid:       true,
		CanEat:        true,
		Reasons:       []string{},
		NutritionGain: make(map[string]int),
	}

	// Find the food in the farm
	var targetFood *model.FarmFood
	for i := range farm.Foods {
		if farm.Foods[i].ID == req.FoodID {
			targetFood = &farm.Foods[i]
			break
		}
	}

	if targetFood == nil {
		result.CanEat = false
		result.Reasons = append(result.Reasons, "food not found on farm")
		return result
	}

	// Check if food has any quantity available
	if targetFood.Quantity <= 0 {
		result.CanEat = false
		result.Reasons = append(result.Reasons, "food quantity is 0, waiting for regeneration")
		result.RegenerationReady = false

		// Calculate next available time
		if targetFood.LastEatenAt != nil {
			regenerationHours := time.Duration(targetFood.RegenerationHours) * time.Hour
			nextAvailable := targetFood.LastEatenAt.Add(regenerationHours)
			result.NextAvailableEat = nextAvailable
		}
		return result
	}

	// Get today's eat count for this pet
	dailyEatCount, err := s.database.GetDailyEatCount(req.EaterPetID, req.Timestamp)
	if err != nil {
		result.IsValid = false
		result.Reasons = append(result.Reasons, "failed to check daily eat count")
		return result
	}

	result.DailyEatCount = dailyEatCount
	result.DailyEatLimit = DailyEatLimit

	// Check if pet has exceeded daily eat limit
	if dailyEatCount >= DailyEatLimit {
		result.CanEat = false
		result.Reasons = append(result.Reasons, "daily eat limit reached")

		// Calculate next available time (next day at 00:00)
		today := time.Now().Truncate(24 * time.Hour)
		result.NextAvailableEat = today.AddDate(0, 0, 1)
		return result
	}

	// Get today's visit count for this repository
	repoVisitCount, err := s.database.GetRepoVisitCount(req.FarmOwner, req.FarmRepo, req.Timestamp)
	if err != nil {
		result.IsValid = false
		result.Reasons = append(result.Reasons, "failed to check repo visit count")
		return result
	}

	result.RepoVisitCount = repoVisitCount
	result.RepoVisitLimit = RepoVisitDailyLimit

	// Check if repository has exceeded daily visit limit
	if repoVisitCount >= RepoVisitDailyLimit {
		result.CanEat = false
		result.Reasons = append(result.Reasons, "repository daily visit limit reached")

		// Calculate next available time (next day at 00:00)
		today := time.Now().Truncate(24 * time.Hour)
		result.NextAvailableEat = today.AddDate(0, 0, 1)
		return result
	}

	// Calculate regeneration status
	result.RegenerationReady = s.isRegenerationReady(targetFood)

	// If all checks pass, calculate nutrition gain based on food type
	result.FoodStatus = map[string]interface{}{
		"id":           targetFood.ID,
		"type":         targetFood.Type,
		"quantity":     targetFood.Quantity,
		"max_quantity": targetFood.MaxQuantity,
	}

	// Set nutrition gain based on food type
	s.setNutritionGain(targetFood.Type, result.NutritionGain)

	return result
}

// RecordFoodConsumption records a food consumption transaction
func (s *FoodService) RecordFoodConsumption(req *model.FoodValidationRequest, farm *model.Farm) error {
	// Find the food
	var targetFood *model.FarmFood
	for i := range farm.Foods {
		if farm.Foods[i].ID == req.FoodID {
			targetFood = &farm.Foods[i]
			break
		}
	}

	if targetFood == nil {
		return nil // Already handled in validation
	}

	// Record the transaction
	now := time.Now()
	return s.database.SaveFoodTransaction(&model.FoodTransaction{
		ID:        req.ID,
		PlayerID:  req.EaterID,
		PetID:     req.EaterPetID,
		FarmOwner: req.FarmOwner,
		FarmRepo:  req.FarmRepo,
		FoodID:    req.FoodID,
		FoodType:  targetFood.Type,
		Quantity:  1,
		Timestamp: now,
		IsValid:   true,
	})
}

// isRegenerationReady checks if food is ready to regenerate
func (s *FoodService) isRegenerationReady(food *model.FarmFood) bool {
	if food.Quantity >= food.MaxQuantity {
		return false // Already at max
	}

	if food.LastEatenAt == nil {
		return true // Never eaten, or waiting after first consumption
	}

	regenerationHours := time.Duration(food.RegenerationHours) * time.Hour
	nextAvailable := food.LastEatenAt.Add(regenerationHours)
	return time.Now().After(nextAvailable)
}

// setNutritionGain sets nutrition gains based on food type
func (s *FoodService) setNutritionGain(foodType string, gains map[string]int) {
	switch model.FoodType(foodType) {
	case model.FoodTypeCookie:
		gains["exp"] = 10
	case model.FoodTypeDonut:
		gains["energy"] = 50
	case model.FoodTypeApple:
		gains["hp"] = 5
		gains["attack"] = 5
		gains["defense"] = 5
		gains["speed"] = 5
		gains["sp_atk"] = 5
		gains["sp_def"] = 5
	case model.FoodTypeGene:
		gains["gene_mutation"] = 1
	}
}
