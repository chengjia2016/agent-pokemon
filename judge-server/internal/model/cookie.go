package model

import "time"

// Cookie Cookie 数据模型
type Cookie struct {
	ID          int        `json:"id"`
	CookieID    string     `json:"cookie_id"` // 0x... 格式
	Type        string     `json:"type"`      // cookie, donut, apple, gene
	Emoji       string     `json:"emoji"`
	SourceFile  string     `json:"source_file"`
	GeneratorID string     `json:"generator_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	ClaimedBy   string     `json:"claimed_by,omitempty"` // 如果已索赔
	ClaimedAt   *time.Time `json:"claimed_at,omitempty"`
}

// CookieClaim Cookie 索赔记录
type CookieClaim struct {
	ID           int       `json:"id"`
	CookieID     string    `json:"cookie_id"`
	PlayerID     string    `json:"player_id"`
	ExpReward    int       `json:"exp_reward"`
	EnergyReward int       `json:"energy_reward"`
	ClaimedAt    time.Time `json:"claimed_at"`
}

// CookieFragment Cookie 碎片
type CookieFragment struct {
	ID           int       `json:"id"`
	FragmentID   string    `json:"fragment_id"`
	PlayerID     string    `json:"player_id"`
	FragmentType string    `json:"fragment_type"` // from_cookie, from_egg, from_battle
	SourceID     string    `json:"source_id"`     // cookie_id, egg_id, battle_id
	CollectedAt  time.Time `json:"collected_at"`
	CreatedAt    time.Time `json:"created_at"`
}

// PlayerFragmentBalance 玩家碎片余额
type PlayerFragmentBalance struct {
	ID             int       `json:"id"`
	PlayerID       string    `json:"player_id"`
	TotalFragments int       `json:"total_fragments"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// CookieResponse API 响应
type CookieResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    *Cookie `json:"data,omitempty"`
	Error   string  `json:"error,omitempty"`
}

// CookieListResponse Cookie 列表响应
type CookieListResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []Cookie `json:"data,omitempty"`
	Total   int      `json:"total"`
	Error   string   `json:"error,omitempty"`
}

// RegisterCookieRequest 注册 Cookie 请求
type RegisterCookieRequest struct {
	CookieID    string `json:"cookie_id" binding:"required"`
	Type        string `json:"type" binding:"required"`
	SourceFile  string `json:"source_file"`
	GeneratorID string `json:"generator_id"`
}

// ClaimCookieRequest 索赔 Cookie 请求
type ClaimCookieRequest struct {
	CookieID string `json:"cookie_id" binding:"required"`
	PlayerID string `json:"player_id" binding:"required"`
}

// ClaimCookieResponse 索赔响应
type ClaimCookieResponse struct {
	Success      bool   `json:"success"`
	CookieID     string `json:"cookie_id"`
	Type         string `json:"type"`
	ExpReward    int    `json:"exp_reward"`
	EnergyReward int    `json:"energy_reward"`
	ClaimedAt    string `json:"claimed_at"`
	Error        string `json:"error,omitempty"`
}

// CookieStatistics Cookie 统计
type CookieStatistics struct {
	TotalCookies int                   `json:"total_cookies"`
	Claimed      int                   `json:"claimed"`
	Unclaimed    int                   `json:"unclaimed"`
	ByType       map[string]int        `json:"by_type"`
	TopPlayers   []PlayerCookieRanking `json:"top_players"`
}

// PlayerCookieRanking 玩家 Cookie 排名
type PlayerCookieRanking struct {
	PlayerID     string `json:"player_id"`
	ClaimedCount int    `json:"claimed_count"`
	TotalExp     int    `json:"total_exp"`
	TotalEnergy  int    `json:"total_energy"`
}

// ScanCookiesRequest 扫描 Cookie 请求
type ScanCookiesRequest struct {
	Directory string `json:"directory" binding:"required"`
	PlayerID  string `json:"player_id" binding:"required"`
}

// ScanCookiesResponse 扫描响应
type ScanCookiesResponse struct {
	Success    bool     `json:"success"`
	FoundCount int      `json:"found_count"`
	NewCount   int      `json:"new_count"`
	NewIDs     []string `json:"new_ids"`
	Error      string   `json:"error,omitempty"`
}
