package model

import "time"

// ===== Database Models (matching SCHEMA_EXTENSIONS.sql) =====

// DatabaseFarm - database representation of a farm
type DatabaseFarm struct {
	ID             int64     `json:"id" db:"id"`
	FarmKey        string    `json:"farm_key" db:"farm_key"` // owner/repository
	OwnerID        string    `json:"owner_id" db:"owner_id"` // GitHub username
	RepositoryName string    `json:"repository_name" db:"repository_name"`
	RepositoryURL  string    `json:"repository_url" db:"repository_url"`
	PlantedAt      time.Time `json:"planted_at" db:"planted_at"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// DatabaseFarmFood - database representation of food in a farm
type DatabaseFarmFood struct {
	ID                int64      `json:"id" db:"id"`
	FarmID            int64      `json:"farm_id" db:"farm_id"`
	FoodID            string     `json:"food_id" db:"food_id"`     // unique food identifier
	FoodType          string     `json:"food_type" db:"food_type"` // cookie, donut, apple, gene
	Emoji             string     `json:"emoji" db:"emoji"`
	CurrentQuantity   int        `json:"current_quantity" db:"current_quantity"`
	MaxQuantity       int        `json:"max_quantity" db:"max_quantity"`
	RegenerationHours int        `json:"regeneration_hours" db:"regeneration_hours"`
	LastEatenAt       *time.Time `json:"last_eaten_at" db:"last_eaten_at"`
	SeedHash          string     `json:"seed_hash" db:"seed_hash"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
}

// DatabaseCookie - database representation of a cookie
type DatabaseCookie struct {
	ID          int64     `json:"id" db:"id"`
	CookieID    string    `json:"cookie_id" db:"cookie_id"`     // 0x... format
	CookieType  string    `json:"cookie_type" db:"cookie_type"` // cookie, donut, apple, gene
	Emoji       string    `json:"emoji" db:"emoji"`
	SourceFile  string    `json:"source_file" db:"source_file"`
	GeneratorID string    `json:"generator_id" db:"generator_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// DatabaseCookieClaim - database representation of a cookie claim
type DatabaseCookieClaim struct {
	ID           int64     `json:"id" db:"id"`
	CookieID     string    `json:"cookie_id" db:"cookie_id"`
	PlayerID     string    `json:"player_id" db:"player_id"`
	ClaimedAt    time.Time `json:"claimed_at" db:"claimed_at"`
	ExpReward    int       `json:"exp_reward" db:"exp_reward"`
	EnergyReward int       `json:"energy_reward" db:"energy_reward"`
}

// DatabaseEgg - database representation of an egg
type DatabaseEgg struct {
	ID              int64      `json:"id" db:"id"`
	EggID           string     `json:"egg_id" db:"egg_id"`     // unique egg identifier
	OwnerID         string     `json:"owner_id" db:"owner_id"` // GitHub username
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	IncubationHours int        `json:"incubation_hours" db:"incubation_hours"`
	Stage           int        `json:"stage" db:"stage"` // 0-3
	HatchedAt       *time.Time `json:"hatched_at" db:"hatched_at"`
	HatchedPetID    *string    `json:"hatched_pet_id" db:"hatched_pet_id"`
	Attributes      string     `json:"attributes" db:"attributes"` // JSON string
}

// DatabaseShopItem - database representation of a shop item
type DatabaseShopItem struct {
	ID          int64     `json:"id" db:"id"`
	ItemID      string    `json:"item_id" db:"item_id"`
	ItemName    string    `json:"item_name" db:"item_name"`
	ItemType    string    `json:"item_type" db:"item_type"`
	Description string    `json:"description" db:"description"`
	Price       int       `json:"price" db:"price"`
	Emoji       string    `json:"emoji" db:"emoji"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// DatabaseShopInventory - database representation of shop inventory
type DatabaseShopInventory struct {
	ID         int64     `json:"id" db:"id"`
	ShopItemID int64     `json:"shop_item_id" db:"shop_item_id"`
	StockLevel int       `json:"stock_level" db:"stock_level"`
	MinStock   int       `json:"min_stock" db:"min_stock"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// DatabaseShopTransaction - database representation of a shop transaction
type DatabaseShopTransaction struct {
	ID           int64     `json:"id" db:"id"`
	ShopItemID   int64     `json:"shop_item_id" db:"shop_item_id"`
	PlayerID     string    `json:"player_id" db:"player_id"`
	TransType    string    `json:"trans_type" db:"trans_type"` // buy, sell
	Quantity     int       `json:"quantity" db:"quantity"`
	TotalPrice   int       `json:"total_price" db:"total_price"`
	TransactedAt time.Time `json:"transacted_at" db:"transacted_at"`
}

// DatabaseCookieFragment - database representation of a cookie fragment
type DatabaseCookieFragment struct {
	ID           int64     `json:"id" db:"id"`
	FragmentID   string    `json:"fragment_id" db:"fragment_id"`
	FragmentType string    `json:"fragment_type" db:"fragment_type"`
	Emoji        string    `json:"emoji" db:"emoji"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// ===== API Request/Response Models =====

// APIFarm is the API response model for a farm
type APIFarm struct {
	ID             int64     `json:"id"`
	FarmKey        string    `json:"farm_key"`
	OwnerID        string    `json:"owner_id"`
	RepositoryName string    `json:"repository_name"`
	RepositoryURL  string    `json:"repository_url"`
	PlantedAt      time.Time `json:"planted_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// APIFarmFood is the API response model for food in a farm
type APIFarmFood struct {
	ID                int64      `json:"id"`
	FarmID            int64      `json:"farm_id"`
	FoodID            string     `json:"food_id"`
	FoodType          string     `json:"food_type"`
	Emoji             string     `json:"emoji"`
	CurrentQuantity   int        `json:"current_quantity"`
	MaxQuantity       int        `json:"max_quantity"`
	RegenerationHours int        `json:"regeneration_hours"`
	LastEatenAt       *time.Time `json:"last_eaten_at"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

// CreateFarmRequest is the request model for creating a farm
type CreateFarmRequest struct {
	OwnerID        string `json:"owner_id" binding:"required"`
	RepositoryName string `json:"repository_name" binding:"required"`
	RepositoryURL  string `json:"repository_url" binding:"required"`
}

// AddFoodRequest is the request model for adding food to a farm
type AddFoodRequest struct {
	FoodID            string `json:"food_id" binding:"required"`
	FoodType          string `json:"food_type" binding:"required"`
	Quantity          int    `json:"quantity" binding:"required,gt=0"`
	MaxQuantity       int    `json:"max_quantity"`
	RegenerationHours int    `json:"regeneration_hours"`
	Emoji             string `json:"emoji"`
}

// ConsumeFoodRequest is the request model for consuming food
type ConsumeFoodRequest struct {
	FoodID     string `json:"food_id" binding:"required"`
	EaterID    string `json:"eater_id" binding:"required"`
	EaterPetID string `json:"eater_pet_id" binding:"required"`
}

// FarmResponse is the generic API response for farm operations
type FarmResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// FarmListResponse is the API response for listing farms
type FarmListResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message,omitempty"`
	Data    []APIFarm `json:"data,omitempty"`
	Total   int       `json:"total"`
	Error   string    `json:"error,omitempty"`
}

// FarmStatistics holds statistics for a farm
type FarmStatistics struct {
	FarmID            int64          `json:"farm_id"`
	TotalFoods        int            `json:"total_foods"`
	TotalQuantity     int            `json:"total_quantity"`
	TotalConsumptions int            `json:"total_consumptions"`
	LastConsumption   *time.Time     `json:"last_consumption"`
	FoodBreakdown     map[string]int `json:"food_breakdown"`
}
