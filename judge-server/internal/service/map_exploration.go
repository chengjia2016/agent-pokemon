package service

import (
	"fmt"
	"math/rand"
	"time"
)

// MapExplorationSystem 地图和探索系统
type MapExplorationSystem struct {
	db Database
}

// NewMapExplorationSystem 创建新的探索系统
func NewMapExplorationSystem(db Database) *MapExplorationSystem {
	return &MapExplorationSystem{db: db}
}

// ==================== 地图结构 ====================

// Region 地区
type Region struct {
	ID            string `json:"id"`
	NameEn        string `json:"name_en"`
	NameZh        string `json:"name_zh"`
	Description   string `json:"description"`
	LevelRangeMin int    `json:"level_range_min"`
	LevelRangeMax int    `json:"level_range_max"`
	Order         int    `json:"order"` // 在地图中的顺序
}

// GrassArea 草地区域
type GrassArea struct {
	ID         string `json:"id"`
	RegionID   string `json:"region_id"`
	NameEn     string `json:"name_en"`
	NameZh     string `json:"name_zh"`
	XCoord     int    `json:"x_coord"`
	YCoord     int    `json:"y_coord"`
	Difficulty int    `json:"difficulty"` // 1-10
}

// WildPokemonSpawn 野生宝可梦生成规则
type WildPokemonSpawn struct {
	ID            int    `json:"id"`
	AreaID        string `json:"area_id"`
	SpeciesID     string `json:"species_id"`
	LevelMin      int    `json:"level_min"`
	LevelMax      int    `json:"level_max"`
	EncounterRate int    `json:"encounter_rate"` // 百分比
	TimeOfDay     string `json:"time_of_day"`    // morning/day/evening/night/any
}

// ==================== 地区初始化 ======================

// GetAllRegions 获取所有地区
func (mes *MapExplorationSystem) GetAllRegions() []Region {
	regions := []Region{
		{
			ID:            "region_kanto",
			NameEn:        "Kanto",
			NameZh:        "关都",
			Description:   "起始地区，有着四季分明的气候",
			LevelRangeMin: 1,
			LevelRangeMax: 30,
			Order:         1,
		},
		{
			ID:            "region_johto",
			NameEn:        "Johto",
			NameZh:        "城都",
			Description:   "与关都相邻的地区，充满历史气息",
			LevelRangeMin: 15,
			LevelRangeMax: 50,
			Order:         2,
		},
		{
			ID:            "region_hoenn",
			NameEn:        "Hoenn",
			NameZh:        "丰缘",
			Description:   "以水而闻名的地区，有大量的海洋区域",
			LevelRangeMin: 20,
			LevelRangeMax: 55,
			Order:         3,
		},
		{
			ID:            "region_sinnoh",
			NameEn:        "Sinnoh",
			NameZh:        "神奥",
			Description:   "寒冷的山区地区",
			LevelRangeMin: 30,
			LevelRangeMax: 65,
			Order:         4,
		},
		{
			ID:            "region_unova",
			NameEn:        "Unova",
			NameZh:        "合众",
			Description:   "都市化的地区",
			LevelRangeMin: 35,
			LevelRangeMax: 70,
			Order:         5,
		},
		{
			ID:            "region_kalos",
			NameEn:        "Kalos",
			NameZh:        "卡洛斯",
			Description:   "浪漫的地区",
			LevelRangeMin: 40,
			LevelRangeMax: 75,
			Order:         6,
		},
	}
	return regions
}

// GetGrassAreasForRegion 获取某地区的草地区域
func (mes *MapExplorationSystem) GetGrassAreasForRegion(regionID string) []GrassArea {
	areas := map[string][]GrassArea{
		"region_kanto": {
			{ID: "area_viridian_forest", RegionID: "region_kanto", NameEn: "Viridian Forest", NameZh: "常青森林", XCoord: 1, YCoord: 1, Difficulty: 1},
			{ID: "area_pewter_city", RegionID: "region_kanto", NameEn: "Pewter City", NameZh: "浅红市", XCoord: 2, YCoord: 1, Difficulty: 2},
			{ID: "area_rock_tunnel", RegionID: "region_kanto", NameEn: "Rock Tunnel", NameZh: "岩山隧道", XCoord: 3, YCoord: 2, Difficulty: 3},
			{ID: "area_route_4", RegionID: "region_kanto", NameEn: "Route 4", NameZh: "4号路线", XCoord: 3, YCoord: 1, Difficulty: 2},
			{ID: "area_cerulean_city", RegionID: "region_kanto", NameEn: "Cerulean City", NameZh: "华蓝市", XCoord: 4, YCoord: 2, Difficulty: 3},
			{ID: "area_route_24", RegionID: "region_kanto", NameEn: "Route 24", NameZh: "24号路线", XCoord: 4, YCoord: 3, Difficulty: 2},
			{ID: "area_route_25", RegionID: "region_kanto", NameEn: "Route 25", NameZh: "25号路线", XCoord: 5, YCoord: 3, Difficulty: 2},
			{ID: "area_vermilion_city", RegionID: "region_kanto", NameEn: "Vermilion City", NameZh: "华蓝市", XCoord: 5, YCoord: 2, Difficulty: 3},
		},
		"region_johto": {
			{ID: "area_new_bark_town", RegionID: "region_johto", NameEn: "New Bark Town", NameZh: "新树镇", XCoord: 1, YCoord: 1, Difficulty: 1},
			{ID: "area_cherrygrove_city", RegionID: "region_johto", NameEn: "Cherrygrove City", NameZh: "酒馆镇", XCoord: 2, YCoord: 1, Difficulty: 2},
			{ID: "area_azalea_town", RegionID: "region_johto", NameEn: "Azalea Town", NameZh: "方可镇", XCoord: 2, YCoord: 2, Difficulty: 2},
		},
	}

	if regionAreas, exists := areas[regionID]; exists {
		return regionAreas
	}
	return []GrassArea{}
}

// ==================== 野生宝可梦生成 ======================

// InitializeWildPokemonSpawns 初始化野生宝可梦生成规则
func (mes *MapExplorationSystem) InitializeWildPokemonSpawns() map[string][]WildPokemonSpawn {
	spawns := map[string][]WildPokemonSpawn{
		// 常青森林
		"area_viridian_forest": {
			{ID: 1, AreaID: "area_viridian_forest", SpeciesID: "pidgeotto", LevelMin: 2, LevelMax: 5, EncounterRate: 40, TimeOfDay: "any"},
			{ID: 2, AreaID: "area_viridian_forest", SpeciesID: "pidgeot", LevelMin: 3, LevelMax: 6, EncounterRate: 35, TimeOfDay: "any"},
			{ID: 3, AreaID: "area_viridian_forest", SpeciesID: "caterpie", LevelMin: 3, LevelMax: 5, EncounterRate: 25, TimeOfDay: "any"},
		},
		// 浅红市
		"area_pewter_city": {
			{ID: 4, AreaID: "area_pewter_city", SpeciesID: "mankey", LevelMin: 5, LevelMax: 8, EncounterRate: 50, TimeOfDay: "day"},
			{ID: 5, AreaID: "area_pewter_city", SpeciesID: "growlithe", LevelMin: 6, LevelMax: 9, EncounterRate: 50, TimeOfDay: "day"},
		},
		// 岩山隧道
		"area_rock_tunnel": {
			{ID: 6, AreaID: "area_rock_tunnel", SpeciesID: "geodude", LevelMin: 8, LevelMax: 12, EncounterRate: 60, TimeOfDay: "any"},
			{ID: 7, AreaID: "area_rock_tunnel", SpeciesID: "onix", LevelMin: 10, LevelMax: 15, EncounterRate: 40, TimeOfDay: "any"},
		},
	}
	return spawns
}

// ==================== 探索逻辑 ======================

// ExploreArea 探索区域
type ExploreResult struct {
	AreaID        string `json:"area_id"`
	Pokemon       string `json:"pokemon"`   // 遇到的宝可梦
	Level         int    `json:"level"`     // 宝可梦的等级
	Encounter     bool   `json:"encounter"` // 是否遇到宝可梦
	ItemFound     string `json:"item_found,omitempty"`
	EncounterRate int    `json:"encounter_rate"`
}

// EnterGrassArea 进入草地区域探索
func (mes *MapExplorationSystem) EnterGrassArea(
	areaID string,
	spawns []WildPokemonSpawn,
) *ExploreResult {
	result := &ExploreResult{
		AreaID: areaID,
	}

	// 计算遇到宝可梦的概率
	encountRate := 50 // 默认50%遇到宝可梦
	if rand.Intn(100) < encountRate {
		result.Encounter = true

		// 从生成表中随机选择宝可梦
		if len(spawns) > 0 {
			spawn := spawns[rand.Intn(len(spawns))]
			result.Pokemon = spawn.SpeciesID
			result.Level = spawn.LevelMin + rand.Intn(spawn.LevelMax-spawn.LevelMin+1)
			result.EncounterRate = spawn.EncounterRate
		}
	} else {
		// 可能找到物品
		items := []string{"potion", "antidote", "paralyze_heal", "awaken", "full_heal"}
		if rand.Intn(100) < 30 { // 30%概率找到物品
			result.ItemFound = items[rand.Intn(len(items))]
		}
	}

	return result
}

// ==================== 探索日志 ======================

// ExplorationLogEntry 探索日志条目
type ExplorationLogEntry struct {
	ID                 int       `json:"id"`
	UserID             int       `json:"user_id"`
	AreaID             string    `json:"area_id"`
	PokemonEncountered string    `json:"pokemon_encountered,omitempty"`
	VisitedAt          time.Time `json:"visited_at"`
}

// LogExploration 记录探索
func (mes *MapExplorationSystem) LogExploration(
	userID int,
	areaID string,
	pokemonEncountered string,
) error {
	// TODO: 保存到数据库
	// INSERT INTO exploration_log (user_id, area_id, pokemon_encountered, visited_at)
	// VALUES (userID, areaID, pokemonEncountered, NOW())

	return nil
}

// GetExplorationHistory 获取用户的探索历史
func (mes *MapExplorationSystem) GetExplorationHistory(userID int) []ExplorationLogEntry {
	// TODO: 从数据库查询

	return []ExplorationLogEntry{}
}

// ==================== 地图连接 ======================

// MapConnection 地图区域之间的连接
type MapConnection struct {
	FromAreaID string `json:"from_area_id"`
	ToAreaID   string `json:"to_area_id"`
	Direction  string `json:"direction"` // north/south/east/west
	Distance   int    `json:"distance"`  // 距离或消耗
}

// GetMapConnections 获取地图连接关系
func (mes *MapExplorationSystem) GetMapConnections(areaID string) []MapConnection {
	connections := map[string][]MapConnection{
		"area_viridian_forest": {
			{FromAreaID: "area_viridian_forest", ToAreaID: "area_pewter_city", Direction: "north", Distance: 5},
		},
		"area_pewter_city": {
			{FromAreaID: "area_pewter_city", ToAreaID: "area_viridian_forest", Direction: "south", Distance: 5},
			{FromAreaID: "area_pewter_city", ToAreaID: "area_route_4", Direction: "east", Distance: 3},
		},
		"area_route_4": {
			{FromAreaID: "area_route_4", ToAreaID: "area_pewter_city", Direction: "west", Distance: 3},
			{FromAreaID: "area_route_4", ToAreaID: "area_rock_tunnel", Direction: "north", Distance: 4},
		},
		"area_rock_tunnel": {
			{FromAreaID: "area_rock_tunnel", ToAreaID: "area_route_4", Direction: "south", Distance: 4},
			{FromAreaID: "area_rock_tunnel", ToAreaID: "area_cerulean_city", Direction: "east", Distance: 5},
		},
	}

	if areaConnections, exists := connections[areaID]; exists {
		return areaConnections
	}
	return []MapConnection{}
}

// CanTraverse 检查是否可以穿过区域
func (mes *MapExplorationSystem) CanTraverse(
	fromAreaID, toAreaID string,
	playerLevel int,
	allowedMoves []string, // HM技能
) (bool, string) {
	// TODO: 检查是否需要特殊技能或等级

	// 检查是否在连接列表中
	connections := mes.GetMapConnections(fromAreaID)
	for _, conn := range connections {
		if conn.ToAreaID == toAreaID {
			return true, ""
		}
	}

	return false, fmt.Sprintf("cannot reach %s from %s", toAreaID, fromAreaID)
}
