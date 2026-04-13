package model

import (
	"database/sql"
	"time"
)

// ShopItem 商店物品
type ShopItem struct {
	ID          int            `json:"id"`
	ItemID      string         `json:"item_id"`
	ItemName    string         `json:"item_name"`
	Description string         `json:"description"`
	ItemType    string         `json:"item_type"` // consumable, equipment, etc
	BasePrice   float64        `json:"base_price"`
	IsAvailable bool           `json:"is_available"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Stock       *ShopInventory `json:"stock,omitempty"`
}

// ShopInventory 商店库存
type ShopInventory struct {
	ID           int        `json:"id"`
	ItemID       string     `json:"item_id"`
	CurrentStock int        `json:"current_stock"`
	MaxStock     int        `json:"max_stock"`
	RestockedAt  *time.Time `json:"restocked_at"`
}

// ShopTransaction 商店交易
type ShopTransaction struct {
	ID              int       `json:"id"`
	PlayerID        string    `json:"player_id"`
	ItemID          string    `json:"item_id"`
	Quantity        int       `json:"quantity"`
	UnitPrice       float64   `json:"unit_price"`
	TotalPrice      float64   `json:"total_price"`
	TransactionType string    `json:"transaction_type"` // buy, sell
	TransactionDate time.Time `json:"transaction_date"`
}

// ShopResponse 商店 API 响应
type ShopResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    *ShopItem `json:"data,omitempty"`
	Error   string    `json:"error,omitempty"`
}

// ShopListResponse 商店列表响应
type ShopListResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    []ShopItem `json:"data,omitempty"`
	Total   int        `json:"total"`
	Error   string     `json:"error,omitempty"`
}

// CreateItemRequest 创建物品请求
type CreateItemRequest struct {
	ItemID      string  `json:"item_id" binding:"required"`
	ItemName    string  `json:"item_name" binding:"required"`
	Description string  `json:"description"`
	ItemType    string  `json:"item_type" binding:"required"`
	BasePrice   float64 `json:"base_price" binding:"required,gt=0"`
	MaxStock    int     `json:"max_stock" binding:"required,gt=0"`
}

// UpdateStockRequest 更新库存请求
type UpdateStockRequest struct {
	ItemID   string `json:"item_id" binding:"required"`
	NewStock int    `json:"new_stock" binding:"required,gte=0"`
}

// BuyItemRequest 购买物品请求
type BuyItemRequest struct {
	PlayerID string `json:"player_id" binding:"required"`
	ItemID   string `json:"item_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required,gt=0"`
}

// BuyItemResponse 购买响应
type BuyItemResponse struct {
	Success        bool    `json:"success"`
	ItemName       string  `json:"item_name"`
	Quantity       int     `json:"quantity"`
	UnitPrice      float64 `json:"unit_price"`
	TotalPrice     float64 `json:"total_price"`
	RemainingCoins float64 `json:"remaining_coins"`
	Error          string  `json:"error,omitempty"`
}

// ShopCatalog 商店目录
type ShopCatalog struct {
	ID            int       `json:"id"`
	CatalogName   string    `json:"catalog_name"`
	ItemCount     int       `json:"item_count"`
	LastRestocked time.Time `json:"last_restocked"`
}

// ShopStatistics 商店统计
type ShopStatistics struct {
	TotalItems          int       `json:"total_items"`
	AvailableItems      int       `json:"available_items"`
	UnavailableItems    int       `json:"unavailable_items"`
	TotalInventoryValue float64   `json:"total_inventory_value"`
	TotalSales          float64   `json:"total_sales"`
	TopSellingItems     []TopItem `json:"top_selling_items"`
}

// TopItem 热销物品
type TopItem struct {
	ItemID       string  `json:"item_id"`
	ItemName     string  `json:"item_name"`
	UnitsSold    int     `json:"units_sold"`
	TotalRevenue float64 `json:"total_revenue"`
}

// RestockRequest 补货请求
type RestockRequest struct {
	ItemID        string `json:"item_id" binding:"required"`
	RestockAmount int    `json:"restock_amount" binding:"required,gt=0"`
}

// RestockResponse 补货响应
type RestockResponse struct {
	Success     bool   `json:"success"`
	ItemID      string `json:"item_id"`
	NewStock    int    `json:"new_stock"`
	RestockedAt string `json:"restocked_at"`
	Error       string `json:"error,omitempty"`
}

// ItemSearchRequest 物品搜索请求
type ItemSearchRequest struct {
	Keyword  string `form:"keyword"`
	ItemType string `form:"type"`
	Page     int    `form:"page,default=1"`
	Limit    int    `form:"limit,default=20"`
}

// ItemSearchResponse 搜索响应
type ItemSearchResponse struct {
	Success bool       `json:"success"`
	Items   []ShopItem `json:"items"`
	Total   int        `json:"total"`
	Page    int        `json:"page"`
	Error   string     `json:"error,omitempty"`
}

// TransactionHistoryRequest 交易历史请求
type TransactionHistoryRequest struct {
	PlayerID string `form:"player_id" binding:"required"`
	Page     int    `form:"page,default=1"`
	Limit    int    `form:"limit,default=20"`
}

// TransactionHistoryResponse 交易历史响应
type TransactionHistoryResponse struct {
	Success      bool              `json:"success"`
	Transactions []ShopTransaction `json:"transactions"`
	Total        int               `json:"total"`
	TotalSpent   float64           `json:"total_spent"`
	Error        string            `json:"error,omitempty"`
}

// Nullable 辅助函数
func NullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func NullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}
