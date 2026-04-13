package model

import "time"

// UserAccount represents a user account in the system
type UserAccount struct {
	ID          int       `json:"id"`
	GithubID    int       `json:"github_id"`
	GithubLogin string    `json:"github_login"`
	Email       string    `json:"email"`
	AvatarURL   string    `json:"avatar_url"`
	Balance     float64   `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserPokemon represents a pokemon owned by a user
type UserPokemon struct {
	ID        int       `json:"id"`
	GithubID  int       `json:"github_id"`
	PetID     string    `json:"pet_id"`
	PetName   string    `json:"pet_name"`
	Level     int       `json:"level"`
	Species   string    `json:"species"`
	CreatedAt time.Time `json:"created_at"`
}

// InventoryItem represents an item in user's inventory
type InventoryItem struct {
	ID         int       `json:"id"`
	GithubID   int       `json:"github_id"`
	ItemID     string    `json:"item_id"`
	ItemName   string    `json:"item_name"`
	Quantity   int       `json:"quantity"`
	AcquiredAt time.Time `json:"acquired_at"`
}

// Transaction represents a balance transaction
type Transaction struct {
	ID            int       `json:"id"`
	GithubID      int       `json:"github_id"`
	Type          string    `json:"type"`
	Amount        float64   `json:"amount"`
	Description   string    `json:"description"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedAt     time.Time `json:"created_at"`
}
