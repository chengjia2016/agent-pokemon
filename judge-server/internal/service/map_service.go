package service

import (
	"database/sql"
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/model"
)

// MapService 地图服务
type MapService struct {
	db *db.Database
}

// NewMapService 创建地图服务
func NewMapService(database *db.Database) *MapService {
	return &MapService{
		db: database,
	}
}

// InitializeMap 初始化地图数据
func (s *MapService) InitializeMap() error {
	// 创建岛屿表（如果不存在）
	islandQuery := `
		CREATE TABLE IF NOT EXISTS islands (
			id VARCHAR(100) PRIMARY KEY,
			name_en VARCHAR(255) NOT NULL,
			name_zh VARCHAR(255),
			description TEXT,
			level INTEGER DEFAULT 1,
			terrain_type VARCHAR(50),
			town_count INTEGER DEFAULT 0,
			max_bases INTEGER DEFAULT 50,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)
	`

	if _, err := s.db.Exec(islandQuery); err != nil {
		return fmt.Errorf("failed to create islands table: %w", err)
	}

	// 创建城镇表
	townQuery := `
		CREATE TABLE IF NOT EXISTS towns (
			id VARCHAR(100) PRIMARY KEY,
			island_id VARCHAR(100) NOT NULL,
			name_en VARCHAR(255) NOT NULL,
			name_zh VARCHAR(255),
			description TEXT,
			type VARCHAR(50),
			coord_x FLOAT,
			coord_y FLOAT,
			wild_pokemon_level INTEGER DEFAULT 5,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (island_id) REFERENCES islands(id)
		)
	`

	if _, err := s.db.Exec(townQuery); err != nil {
		return fmt.Errorf("failed to create towns table: %w", err)
	}

	return nil
}

// CreateIsland 创建岛屿
func (s *MapService) CreateIsland(island *model.Island) (*model.Island, error) {
	query := `
		INSERT INTO islands (id, name_en, name_zh, description, level, terrain_type, town_count, max_bases, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		RETURNING id, name_en, name_zh, description, level, terrain_type, town_count, max_bases, created_at, updated_at
	`

	result := &model.Island{}
	err := s.db.QueryRow(query,
		island.ID, island.NameEn, island.NameZh, island.Description,
		island.Level, island.TerrainType, island.TownCount, island.MaxBases,
	).Scan(
		&result.ID, &result.NameEn, &result.NameZh, &result.Description,
		&result.Level, &result.TerrainType, &result.TownCount, &result.MaxBases,
		&result.CreatedAt, &result.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create island: %w", err)
	}

	return result, nil
}

// GetIsland 获取岛屿信息
func (s *MapService) GetIsland(islandID string) (*model.Island, error) {
	query := `
		SELECT id, name_en, name_zh, description, level, terrain_type, town_count, max_bases, created_at, updated_at
		FROM islands
		WHERE id = $1
	`

	island := &model.Island{}
	err := s.db.QueryRow(query, islandID).Scan(
		&island.ID, &island.NameEn, &island.NameZh, &island.Description,
		&island.Level, &island.TerrainType, &island.TownCount, &island.MaxBases,
		&island.CreatedAt, &island.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("island not found")
		}
		return nil, fmt.Errorf("failed to get island: %w", err)
	}

	return island, nil
}

// ListIslands 列出所有岛屿
func (s *MapService) ListIslands() ([]*model.Island, error) {
	query := `
		SELECT id, name_en, name_zh, description, level, terrain_type, town_count, max_bases, created_at, updated_at
		FROM islands
		ORDER BY level ASC
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list islands: %w", err)
	}
	defer rows.Close()

	var islands []*model.Island
	for rows.Next() {
		island := &model.Island{}
		err := rows.Scan(
			&island.ID, &island.NameEn, &island.NameZh, &island.Description,
			&island.Level, &island.TerrainType, &island.TownCount, &island.MaxBases,
			&island.CreatedAt, &island.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan island: %w", err)
		}
		islands = append(islands, island)
	}

	return islands, nil
}

// CreateTown 创建城镇
func (s *MapService) CreateTown(town *model.Town) (*model.Town, error) {
	query := `
		INSERT INTO towns (id, island_id, name_en, name_zh, description, type, coord_x, coord_y, wild_pokemon_level, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
		RETURNING id, island_id, name_en, name_zh, description, type, coord_x, coord_y, wild_pokemon_level, created_at, updated_at
	`

	result := &model.Town{}
	err := s.db.QueryRow(query,
		town.ID, town.IslandID, town.NameEn, town.NameZh, town.Description,
		town.Type, town.CoordX, town.CoordY, town.WildPokemonLevel,
	).Scan(
		&result.ID, &result.IslandID, &result.NameEn, &result.NameZh, &result.Description,
		&result.Type, &result.CoordX, &result.CoordY, &result.WildPokemonLevel,
		&result.CreatedAt, &result.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create town: %w", err)
	}

	return result, nil
}

// GetTown 获取城镇信息
func (s *MapService) GetTown(townID string) (*model.Town, error) {
	query := `
		SELECT id, island_id, name_en, name_zh, description, type, coord_x, coord_y, wild_pokemon_level, created_at, updated_at
		FROM towns
		WHERE id = $1
	`

	town := &model.Town{}
	err := s.db.QueryRow(query, townID).Scan(
		&town.ID, &town.IslandID, &town.NameEn, &town.NameZh, &town.Description,
		&town.Type, &town.CoordX, &town.CoordY, &town.WildPokemonLevel,
		&town.CreatedAt, &town.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("town not found")
		}
		return nil, fmt.Errorf("failed to get town: %w", err)
	}

	return town, nil
}

// GetTownsByIsland 获取岛屿上的所有城镇
func (s *MapService) GetTownsByIsland(islandID string) ([]*model.Town, error) {
	query := `
		SELECT id, island_id, name_en, name_zh, description, type, coord_x, coord_y, wild_pokemon_level, created_at, updated_at
		FROM towns
		WHERE island_id = $1
		ORDER BY name_en ASC
	`

	rows, err := s.db.Query(query, islandID)
	if err != nil {
		return nil, fmt.Errorf("failed to get towns: %w", err)
	}
	defer rows.Close()

	var towns []*model.Town
	for rows.Next() {
		town := &model.Town{}
		err := rows.Scan(
			&town.ID, &town.IslandID, &town.NameEn, &town.NameZh, &town.Description,
			&town.Type, &town.CoordX, &town.CoordY, &town.WildPokemonLevel,
			&town.CreatedAt, &town.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan town: %w", err)
		}
		towns = append(towns, town)
	}

	return towns, nil
}

// GetMapData 获取地图数据（包括岛屿和城镇）
func (s *MapService) GetMapData() (*model.MapResponse, error) {
	response := &model.MapResponse{Success: true}

	// 获取所有岛屿
	islands, err := s.ListIslands()
	if err != nil {
		return nil, fmt.Errorf("failed to get islands: %w", err)
	}
	response.Islands = islands

	// 获取所有城镇
	query := `
		SELECT id, island_id, name_en, name_zh, description, type, coord_x, coord_y, wild_pokemon_level, created_at, updated_at
		FROM towns
		ORDER BY island_id, name_en ASC
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get towns: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		town := &model.Town{}
		err := rows.Scan(
			&town.ID, &town.IslandID, &town.NameEn, &town.NameZh, &town.Description,
			&town.Type, &town.CoordX, &town.CoordY, &town.WildPokemonLevel,
			&town.CreatedAt, &town.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan town: %w", err)
		}
		response.Towns = append(response.Towns, town)
	}

	return response, nil
}
