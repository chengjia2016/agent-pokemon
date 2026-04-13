package db

import (
	"database/sql"
	"errors"
	"judge-server/internal/model"
)

// User Account Operations

func (d *Database) CreateUserAccount(account *model.UserAccount) error {
	query := `
		INSERT INTO user_accounts (github_id, github_login, email, avatar_url, balance)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	err := d.conn.QueryRow(query,
		account.GithubID, account.GithubLogin, account.Email, account.AvatarURL, account.Balance).
		Scan(&account.ID, &account.CreatedAt, &account.UpdatedAt)
	return err
}

func (d *Database) GetUserAccount(githubID int) (*model.UserAccount, error) {
	query := `
		SELECT id, github_id, github_login, email, avatar_url, balance, created_at, updated_at
		FROM user_accounts WHERE github_id = $1
	`
	account := &model.UserAccount{}
	err := d.conn.QueryRow(query, githubID).Scan(
		&account.ID, &account.GithubID, &account.GithubLogin, &account.Email,
		&account.AvatarURL, &account.Balance, &account.CreatedAt, &account.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return account, err
}

func (d *Database) GetUserBalance(githubID int) (float64, error) {
	query := `SELECT balance FROM user_accounts WHERE github_id = $1`
	var balance float64
	err := d.conn.QueryRow(query, githubID).Scan(&balance)
	if err == sql.ErrNoRows {
		return 0, errors.New("user not found")
	}
	return balance, err
}

func (d *Database) UpdateUserBalance(githubID int, amount float64, description string) error {
	tx, err := d.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Get current balance
	var oldBalance float64
	err = tx.QueryRow(`SELECT balance FROM user_accounts WHERE github_id = $1`, githubID).Scan(&oldBalance)
	if err != nil {
		return err
	}

	// Update balance
	newBalance := oldBalance + amount
	_, err = tx.Exec(`
		UPDATE user_accounts SET balance = $1, updated_at = NOW() WHERE github_id = $2
	`, newBalance, githubID)
	if err != nil {
		return err
	}

	// Record transaction
	_, err = tx.Exec(`
		INSERT INTO account_transactions (github_id, transaction_type, amount, description, balance_before, balance_after)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, githubID, "balance_update", amount, description, oldBalance, newBalance)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// User Pokemon Operations

func (d *Database) AddUserPokemon(pokemon *model.UserPokemon) error {
	query := `
		INSERT INTO user_pokemons (github_id, pet_id, pet_name, level, species)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	err := d.conn.QueryRow(query,
		pokemon.GithubID, pokemon.PetID, pokemon.PetName, pokemon.Level, pokemon.Species).
		Scan(&pokemon.ID, &pokemon.CreatedAt)
	return err
}

func (d *Database) GetUserPokemons(githubID int) ([]model.UserPokemon, error) {
	query := `
		SELECT id, github_id, pet_id, pet_name, level, species, created_at
		FROM user_pokemons WHERE github_id = $1 ORDER BY created_at DESC
	`
	rows, err := d.conn.Query(query, githubID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pokemons []model.UserPokemon
	for rows.Next() {
		var p model.UserPokemon
		err := rows.Scan(&p.ID, &p.GithubID, &p.PetID, &p.PetName, &p.Level, &p.Species, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		pokemons = append(pokemons, p)
	}
	return pokemons, rows.Err()
}

func (d *Database) GetUserPokemon(githubID int, petID string) (*model.UserPokemon, error) {
	query := `
		SELECT id, github_id, pet_id, pet_name, level, species, created_at
		FROM user_pokemons WHERE github_id = $1 AND pet_id = $2
	`
	pokemon := &model.UserPokemon{}
	err := d.conn.QueryRow(query, githubID, petID).Scan(
		&pokemon.ID, &pokemon.GithubID, &pokemon.PetID, &pokemon.PetName,
		&pokemon.Level, &pokemon.Species, &pokemon.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("pokemon not found")
	}
	return pokemon, err
}

// User Inventory Operations

func (d *Database) AddUserItem(githubID int, itemID, itemName string, quantity int) error {
	query := `
		INSERT INTO user_inventory (github_id, item_id, item_name, quantity)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (github_id, item_id) DO UPDATE SET quantity = user_inventory.quantity + $4
	`
	_, err := d.conn.Exec(query, githubID, itemID, itemName, quantity)
	return err
}

func (d *Database) GetUserInventory(githubID int) ([]model.InventoryItem, error) {
	query := `
		SELECT id, github_id, item_id, item_name, quantity, acquired_at
		FROM user_inventory WHERE github_id = $1 ORDER BY acquired_at DESC
	`
	rows, err := d.conn.Query(query, githubID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.InventoryItem
	for rows.Next() {
		var item model.InventoryItem
		err := rows.Scan(&item.ID, &item.GithubID, &item.ItemID, &item.ItemName,
			&item.Quantity, &item.AcquiredAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (d *Database) UpdateUserItemQuantity(githubID int, itemID string, quantity int) error {
	query := `
		UPDATE user_inventory SET quantity = $1 WHERE github_id = $2 AND item_id = $3
	`
	result, err := d.conn.Exec(query, quantity, githubID, itemID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("item not found")
	}
	return nil
}

func (d *Database) RemoveUserItem(githubID int, itemID string) error {
	query := `DELETE FROM user_inventory WHERE github_id = $1 AND item_id = $2`
	result, err := d.conn.Exec(query, githubID, itemID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("item not found")
	}
	return nil
}

// Transaction History

func (d *Database) GetUserTransactions(githubID int, limit int) ([]model.Transaction, error) {
	query := `
		SELECT id, github_id, transaction_type, amount, description, balance_before, balance_after, created_at
		FROM account_transactions WHERE github_id = $1
		ORDER BY created_at DESC LIMIT $2
	`
	rows, err := d.conn.Query(query, githubID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var tx model.Transaction
		err := rows.Scan(&tx.ID, &tx.GithubID, &tx.Type, &tx.Amount, &tx.Description,
			&tx.BalanceBefore, &tx.BalanceAfter, &tx.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, rows.Err()
}

// Utility: Update Schema to include user tables
func (d *Database) UpdateSchema() error {
	schemas := []string{
		`CREATE TABLE IF NOT EXISTS user_accounts (
			id SERIAL PRIMARY KEY,
			github_id INTEGER UNIQUE NOT NULL,
			github_login VARCHAR(255) NOT NULL,
			email VARCHAR(255),
			avatar_url TEXT,
			balance DECIMAL(12, 2) DEFAULT 0.0,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS user_pokemons (
			id SERIAL PRIMARY KEY,
			github_id INTEGER NOT NULL,
			pet_id TEXT NOT NULL UNIQUE,
			pet_name TEXT NOT NULL,
			level INTEGER DEFAULT 1,
			species TEXT,
			created_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (github_id) REFERENCES user_accounts(github_id)
		)`,
		`CREATE TABLE IF NOT EXISTS user_inventory (
			id SERIAL PRIMARY KEY,
			github_id INTEGER NOT NULL,
			item_id VARCHAR(255) NOT NULL,
			item_name VARCHAR(255) NOT NULL,
			quantity INTEGER DEFAULT 1,
			acquired_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (github_id) REFERENCES user_accounts(github_id),
			UNIQUE(github_id, item_id)
		)`,
		`CREATE TABLE IF NOT EXISTS account_transactions (
			id SERIAL PRIMARY KEY,
			github_id INTEGER NOT NULL,
			transaction_type VARCHAR(50) NOT NULL,
			amount DECIMAL(12, 2) NOT NULL,
			description TEXT,
			balance_before DECIMAL(12, 2),
			balance_after DECIMAL(12, 2),
			created_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (github_id) REFERENCES user_accounts(github_id)
		)`,
		`CREATE INDEX IF NOT EXISTS idx_user_accounts_github_id ON user_accounts(github_id)`,
		`CREATE INDEX IF NOT EXISTS idx_user_pokemons_github_id ON user_pokemons(github_id)`,
		`CREATE INDEX IF NOT EXISTS idx_user_inventory_github_id ON user_inventory(github_id)`,
		`CREATE INDEX IF NOT EXISTS idx_account_transactions_github_id ON account_transactions(github_id)`,
	}

	for _, schema := range schemas {
		if _, err := d.conn.Exec(schema); err != nil {
			return err
		}
	}
	return nil
}
