package model

import "time"

// Island 岛屿数据模型
type Island struct {
	ID          string    `json:"id"`
	NameEn      string    `json:"name_en"`
	NameZh      string    `json:"name_zh"`
	Description string    `json:"description"`
	Level       int       `json:"level"`
	TerrainType string    `json:"terrain_type"`
	TownCount   int       `json:"town_count"`
	MaxBases    int       `json:"max_bases"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Town 城镇数据模型
type Town struct {
	ID               string    `json:"id"`
	IslandID         string    `json:"island_id"`
	NameEn           string    `json:"name_en"`
	NameZh           string    `json:"name_zh"`
	Description      string    `json:"description"`
	Type             string    `json:"type"`
	CoordX           float64   `json:"coord_x"`
	CoordY           float64   `json:"coord_y"`
	WildPokemonLevel int       `json:"wild_pokemon_level"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// MapResponse 地图响应数据模型
type MapResponse struct {
	Success bool      `json:"success"`
	Islands []*Island `json:"islands"`
	Towns   []*Town   `json:"towns"`
}

// CreateIslandRequest 创建岛屿请求
type CreateIslandRequest struct {
	ID          string `json:"id" binding:"required"`
	NameEn      string `json:"name_en" binding:"required"`
	NameZh      string `json:"name_zh"`
	Description string `json:"description"`
	Level       int    `json:"level"`
	TerrainType string `json:"terrain_type"`
	MaxBases    int    `json:"max_bases"`
}

// CreateTownRequest 创建城镇请求
type CreateTownRequest struct {
	ID               string  `json:"id" binding:"required"`
	IslandID         string  `json:"island_id" binding:"required"`
	NameEn           string  `json:"name_en" binding:"required"`
	NameZh           string  `json:"name_zh"`
	Description      string  `json:"description"`
	Type             string  `json:"type"`
	CoordX           float64 `json:"coord_x"`
	CoordY           float64 `json:"coord_y"`
	WildPokemonLevel int     `json:"wild_pokemon_level"`
}
