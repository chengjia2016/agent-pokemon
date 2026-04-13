package db

import (
	"database/sql"
	"errors"
	"judge-server/internal/model"
)

// ===== Shop Item Operations =====

// CreateShopItem creates a new shop item
func (d *Database) CreateShopItem(item *model.DatabaseShopItem) (*model.DatabaseShopItem, error) {
	query := `
		INSERT INTO shop_items (item_id, item_name, item_type, description, price, emoji)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, item_id, item_name, item_type, description, price, emoji, created_at, updated_at
	`

	err := d.conn.QueryRow(query,
		item.ItemID, item.ItemName, item.ItemType, item.Description, item.Price, item.Emoji).
		Scan(&item.ID, &item.ItemID, &item.ItemName, &item.ItemType, &item.Description,
			&item.Price, &item.Emoji, &item.CreatedAt, &item.UpdatedAt)

	return item, err
}

// GetShopItem retrieves a shop item by ID
func (d *Database) GetShopItem(itemID string) (*model.DatabaseShopItem, error) {
	query := `
		SELECT id, item_id, item_name, item_type, description, price, emoji, created_at, updated_at
		FROM shop_items WHERE item_id = $1
	`

	item := &model.DatabaseShopItem{}
	err := d.conn.QueryRow(query, itemID).Scan(
		&item.ID, &item.ItemID, &item.ItemName, &item.ItemType, &item.Description,
		&item.Price, &item.Emoji, &item.CreatedAt, &item.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("item not found")
	}

	return item, err
}

// GetAllShopItems retrieves all shop items
func (d *Database) GetAllShopItems() ([]model.DatabaseShopItem, error) {
	query := `
		SELECT id, item_id, item_name, item_type, description, price, emoji, created_at, updated_at
		FROM shop_items
		ORDER BY created_at DESC
		LIMIT 100
	`

	rows, err := d.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.DatabaseShopItem
	for rows.Next() {
		item := model.DatabaseShopItem{}
		err := rows.Scan(
			&item.ID, &item.ItemID, &item.ItemName, &item.ItemType, &item.Description,
			&item.Price, &item.Emoji, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

// ===== Shop Inventory Operations =====

// UpdateStockLevel updates the stock level for an item
func (d *Database) UpdateStockLevel(shopItemID int64, quantity int) error {
	query := `
		INSERT INTO shop_inventory (shop_item_id, stock_level, min_stock)
		VALUES ($1, $2, 0)
		ON CONFLICT (shop_item_id) DO UPDATE
		SET stock_level = $2
	`

	_, err := d.conn.Exec(query, shopItemID, quantity)
	return err
}

// GetStockLevel retrieves stock level for an item
func (d *Database) GetStockLevel(shopItemID int64) (int, error) {
	query := "SELECT COALESCE(stock_level, 0) FROM shop_inventory WHERE shop_item_id = $1"

	var level int
	err := d.conn.QueryRow(query, shopItemID).Scan(&level)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	return level, err
}

// ===== Shop Transaction Operations =====

// BuyItem records a purchase transaction
func (d *Database) BuyItem(shopItemID int64, playerID string, quantity int, totalPrice int) (*model.DatabaseShopTransaction, error) {
	tx, err := d.conn.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Check stock
	var stock int
	err = tx.QueryRow("SELECT COALESCE(stock_level, 0) FROM shop_inventory WHERE shop_item_id = $1", shopItemID).Scan(&stock)
	if err != nil {
		return nil, err
	}

	if stock < quantity {
		return nil, errors.New("insufficient stock")
	}

	// Deduct stock
	_, err = tx.Exec("UPDATE shop_inventory SET stock_level = stock_level - $1 WHERE shop_item_id = $2", quantity, shopItemID)
	if err != nil {
		return nil, err
	}

	// Record transaction
	transQuery := `
		INSERT INTO shop_transactions (shop_item_id, player_id, trans_type, quantity, total_price, transacted_at)
		VALUES ($1, $2, 'buy', $3, $4, NOW())
		RETURNING id, shop_item_id, player_id, trans_type, quantity, total_price, transacted_at
	`

	transaction := &model.DatabaseShopTransaction{}
	err = tx.QueryRow(transQuery, shopItemID, playerID, quantity, totalPrice).Scan(
		&transaction.ID, &transaction.ShopItemID, &transaction.PlayerID, &transaction.TransType,
		&transaction.Quantity, &transaction.TotalPrice, &transaction.TransactedAt)

	if err != nil {
		return nil, err
	}

	return transaction, tx.Commit()
}

// GetTransactionHistory retrieves transaction history for a player
func (d *Database) GetTransactionHistory(playerID string, limit int) ([]model.DatabaseShopTransaction, error) {
	if limit <= 0 {
		limit = 100
	}

	query := `
		SELECT id, item_id, player_id, transaction_type, quantity, total_price, created_at
		FROM shop_transactions
		WHERE player_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := d.conn.Query(query, playerID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.DatabaseShopTransaction
	for rows.Next() {
		trans := model.DatabaseShopTransaction{}
		err := rows.Scan(
			&trans.ID, &trans.ShopItemID, &trans.PlayerID, &trans.TransType,
			&trans.Quantity, &trans.TotalPrice, &trans.TransactedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, trans)
	}

	return transactions, rows.Err()
}

// GetShopStats retrieves shop statistics
func (d *Database) GetShopStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	var totalItems int
	err := d.conn.QueryRow("SELECT COUNT(*) FROM shop_items").Scan(&totalItems)
	if err != nil {
		return nil, err
	}
	stats["total_items"] = totalItems

	var totalTransactions int
	err = d.conn.QueryRow("SELECT COUNT(*) FROM shop_transactions").Scan(&totalTransactions)
	if err != nil {
		return nil, err
	}
	stats["total_transactions"] = totalTransactions

	var totalRevenue int
	err = d.conn.QueryRow("SELECT COALESCE(SUM(total_price), 0) FROM shop_transactions WHERE transaction_type = 'buy'").Scan(&totalRevenue)
	if err != nil {
		return nil, err
	}
	stats["total_revenue"] = totalRevenue

	return stats, nil
}
