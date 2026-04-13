package model

import "time"

// FoodType 食物类型
type FoodType string

const (
	FoodTypeCookie FoodType = "cookie"
	FoodTypeDonut  FoodType = "donut"
	FoodTypeApple  FoodType = "apple"
	FoodTypeGene   FoodType = "gene"
)

// FoodValidationRequest 食物验证请求
type FoodValidationRequest struct {
	ID         string    `json:"id"`
	EaterID    string    `json:"eater_id"`
	EaterPetID string    `json:"eater_pet_id"`
	FarmOwner  string    `json:"farm_owner"`
	FarmRepo   string    `json:"farm_repo"`
	FoodID     string    `json:"food_id"`
	Timestamp  time.Time `json:"timestamp"`
}

// FoodValidationResult 食物验证结果
type FoodValidationResult struct {
	IsValid           bool                   `json:"is_valid"`
	CanEat            bool                   `json:"can_eat"`
	Reasons           []string               `json:"reasons"`
	DailyEatCount     int                    `json:"daily_eat_count"`
	DailyEatLimit     int                    `json:"daily_eat_limit"`
	RepoVisitCount    int                    `json:"repo_visit_count"`
	RepoVisitLimit    int                    `json:"repo_visit_limit"`
	NextAvailableEat  time.Time              `json:"next_available_eat"`
	RegenerationReady bool                   `json:"regeneration_ready"`
	FoodStatus        map[string]interface{} `json:"food_status"`
	NutritionGain     map[string]int         `json:"nutrition_gain"`
}

// Farm 农场数据（来自 GitHub）
type Farm struct {
	Owner      string     `json:"owner"`
	Repository string     `json:"repository"`
	URL        string     `json:"url"`
	PlantedAt  time.Time  `json:"planted_at"`
	Foods      []FarmFood `json:"foods"`
}

// FarmFood 农场中的食物
type FarmFood struct {
	ID                string     `json:"id"`
	Type              string     `json:"type"`
	Emoji             string     `json:"emoji"`
	Quantity          int        `json:"quantity"`
	MaxQuantity       int        `json:"max_quantity"`
	RegenerationHours int        `json:"regeneration_hours"`
	LastEatenAt       *time.Time `json:"last_eaten_at"`
	Seed              string     `json:"seed"`
}

// FoodTransaction 食物消费事务记录
type FoodTransaction struct {
	ID        string    `json:"id"`
	PlayerID  string    `json:"player_id"`
	PetID     string    `json:"pet_id"`
	FarmOwner string    `json:"farm_owner"`
	FarmRepo  string    `json:"farm_repo"`
	FoodID    string    `json:"food_id"`
	FoodType  string    `json:"food_type"`
	Quantity  int       `json:"quantity"`
	Timestamp time.Time `json:"timestamp"`
	IsValid   bool      `json:"is_valid"`
}
