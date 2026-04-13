package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// MapData represents a game map structure
type MapData struct {
	Version     string        `json:"version"`
	MapID       string        `json:"map_id"`
	OwnerID     int           `json:"owner_id"`
	OwnerName   string        `json:"owner_username"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
	Width       int           `json:"width"`
	Height      int           `json:"height"`
	Terrain     [][]int       `json:"terrain"`
	Elements    []MapElement  `json:"elements"`
	Connections MapConnection `json:"connections"`
	Statistics  MapStatistics `json:"statistics"`
}

// MapElement represents an element on the map
type MapElement struct {
	ID   string                 `json:"id"`
	Type string                 `json:"type"` // wild_pokemon, food, obstacle, npc
	X    int                    `json:"x"`
	Y    int                    `json:"y"`
	Data map[string]interface{} `json:"data"`
}

// MapConnection represents connections to adjacent maps
type MapConnection struct {
	North *string `json:"north"`
	South *string `json:"south"`
	East  *string `json:"east"`
	West  *string `json:"west"`
}

// MapStatistics represents map statistics
type MapStatistics struct {
	TotalWildPokemon int    `json:"total_wild_pokemon"`
	TotalFood        int    `json:"total_food"`
	TotalObstacles   int    `json:"total_obstacles"`
	VisitedCount     int    `json:"visited_count"`
	LastVisited      string `json:"last_visited"`
}

// GetMap 获取指定的地图
func (h *Handler) GetMap(w http.ResponseWriter, r *http.Request) {
	// Extract map ID from URL
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/maps/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, `{"error":"map_id is required","success":false}`, http.StatusBadRequest)
		return
	}

	mapID := parts[0]
	mapData, err := h.loadMapByID(mapID)
	if err != nil {
		http.Error(w, `{"error":"map not found","success":false}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    mapData,
	})
}

// ListMaps 获取所有地图列表 (支持分页)
func (h *Handler) ListMaps(w http.ResponseWriter, r *http.Request) {
	// 从URL查询参数获取过滤条件和分页参数
	query := r.URL.Query()
	ownerID := query.Get("owner_id")
	pageStr := query.Get("page")
	limitStr := query.Get("limit")

	// Set default pagination values
	page := 1
	limit := 10
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	mapsList, err := h.loadAllMaps()
	if err != nil {
		http.Error(w, `{"error":"failed to load maps","success":false}`, http.StatusInternalServerError)
		return
	}

	// Filter by owner if specified
	if ownerID != "" {
		ownerIDNum, _ := strconv.Atoi(ownerID)
		filtered := make([]MapData, 0)
		for _, m := range mapsList {
			if m.OwnerID == ownerIDNum {
				filtered = append(filtered, m)
			}
		}
		mapsList = filtered
	}

	// Calculate pagination
	totalCount := len(mapsList)
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex >= totalCount {
		// If page is out of range, return empty results
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":     true,
			"count":       0,
			"total":       totalCount,
			"page":        page,
			"limit":       limit,
			"total_pages": (totalCount + limit - 1) / limit,
			"data":        make([]MapData, 0),
		})
		return
	}

	if endIndex > totalCount {
		endIndex = totalCount
	}

	paginatedMaps := mapsList[startIndex:endIndex]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":     true,
		"count":       len(paginatedMaps),
		"total":       totalCount,
		"page":        page,
		"limit":       limit,
		"total_pages": (totalCount + limit - 1) / limit,
		"data":        paginatedMaps,
	})
}

// GetMapConnections 获取地图的连接信息
func (h *Handler) GetMapConnections(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/maps/"), "/")
	if len(parts) < 2 || parts[1] != "connections" {
		http.Error(w, `{"error":"invalid request","success":false}`, http.StatusBadRequest)
		return
	}

	mapID := parts[0]
	mapData, err := h.loadMapByID(mapID)
	if err != nil {
		http.Error(w, `{"error":"map not found","success":false}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"map_id":  mapID,
		"data":    mapData.Connections,
	})
}

// GetMapElements 获取地图上指定类型的元素
func (h *Handler) GetMapElements(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/maps/"), "/")
	if len(parts) < 2 || parts[1] != "elements" {
		http.Error(w, `{"error":"invalid request","success":false}`, http.StatusBadRequest)
		return
	}

	mapID := parts[0]
	mapData, err := h.loadMapByID(mapID)
	if err != nil {
		http.Error(w, `{"error":"map not found","success":false}`, http.StatusNotFound)
		return
	}

	// Filter by type if specified
	query := r.URL.Query()
	elemType := query.Get("type")

	elements := mapData.Elements
	if elemType != "" {
		filtered := make([]MapElement, 0)
		for _, e := range elements {
			if e.Type == elemType {
				filtered = append(filtered, e)
			}
		}
		elements = filtered
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"map_id":   mapID,
		"elements": elements,
		"count":    len(elements),
	})
}

// SearchMaps 搜索地图
// SearchMaps 搜索地图 (支持分页)
func (h *Handler) SearchMaps(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	search := query.Get("q")
	ownerID := query.Get("owner_id")
	pageStr := query.Get("page")
	limitStr := query.Get("limit")

	// Set default pagination values
	page := 1
	limit := 10
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	mapsList, err := h.loadAllMaps()
	if err != nil {
		http.Error(w, `{"error":"failed to load maps","success":false}`, http.StatusInternalServerError)
		return
	}

	results := make([]MapData, 0)
	for _, m := range mapsList {
		// Search by map ID or owner name
		if search != "" {
			if !strings.Contains(m.MapID, search) && !strings.Contains(m.OwnerName, search) {
				continue
			}
		}

		// Filter by owner ID
		if ownerID != "" {
			ownerIDNum, _ := strconv.Atoi(ownerID)
			if m.OwnerID != ownerIDNum {
				continue
			}
		}

		results = append(results, m)
	}

	// Calculate pagination
	totalCount := len(results)
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex >= totalCount {
		// If page is out of range, return empty results
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":     true,
			"count":       0,
			"total":       totalCount,
			"page":        page,
			"limit":       limit,
			"total_pages": (totalCount + limit - 1) / limit,
			"data":        make([]MapData, 0),
		})
		return
	}

	if endIndex > totalCount {
		endIndex = totalCount
	}

	paginatedResults := results[startIndex:endIndex]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":     true,
		"count":       len(paginatedResults),
		"total":       totalCount,
		"page":        page,
		"limit":       limit,
		"total_pages": (totalCount + limit - 1) / limit,
		"data":        paginatedResults,
	})
}

// Helper functions

// loadMapByID loads a map by ID with caching support
func (h *Handler) loadMapByID(mapID string) (*MapData, error) {
	// Check cache first
	h.mapCacheMutex.RLock()
	if cached, exists := h.mapCache[mapID]; exists && time.Now().Before(cached.ExpiresAt) {
		h.mapCacheMutex.RUnlock()
		return cached.Data, nil
	}
	h.mapCacheMutex.RUnlock()

	// Load from disk
	mapPath := filepath.Join("../maps", fmt.Sprintf("%s.json", mapID))

	data, err := ioutil.ReadFile(mapPath)
	if err != nil {
		return nil, err
	}

	var mapData MapData
	err = json.Unmarshal(data, &mapData)
	if err != nil {
		return nil, err
	}

	// Cache the map data
	h.mapCacheMutex.Lock()
	h.mapCache[mapID] = &MapCacheEntry{
		Data:      &mapData,
		ExpiresAt: time.Now().Add(h.mapCacheTTL),
	}
	h.mapCacheMutex.Unlock()

	return &mapData, nil
}

// loadAllMaps 加载所有地图
func (h *Handler) loadAllMaps() ([]MapData, error) {
	mapsDir := "../maps"

	// Check if maps directory exists
	if _, err := os.Stat(mapsDir); err != nil {
		return []MapData{}, nil
	}

	files, err := ioutil.ReadDir(mapsDir)
	if err != nil {
		return nil, err
	}

	maps := make([]MapData, 0)
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" && file.Name() != "generator.go" {
			mapPath := filepath.Join(mapsDir, file.Name())
			data, err := ioutil.ReadFile(mapPath)
			if err != nil {
				continue
			}

			var mapData MapData
			err = json.Unmarshal(data, &mapData)
			if err != nil {
				continue
			}

			maps = append(maps, mapData)
		}
	}

	return maps, nil
}

// generateUserMap 为用户生成新地图 (调用maps生成器)
func (h *Handler) GenerateUserMap(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, `{"error":"method not allowed","success":false}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		OwnerID   int    `json:"owner_id"`
		OwnerName string `json:"owner_name"`
		MapID     string `json:"map_id"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `{"error":"invalid request","success":false}`, http.StatusBadRequest)
		return
	}

	// Validate inputs
	if req.OwnerID == 0 || req.OwnerName == "" || req.MapID == "" {
		http.Error(w, `{"error":"missing required fields","success":false}`, http.StatusBadRequest)
		return
	}

	// Set defaults for width/height
	if req.Width == 0 {
		req.Width = 20
	}
	if req.Height == 0 {
		req.Height = 20
	}

	// Check if map already exists
	if _, err := h.loadMapByID(req.MapID); err == nil {
		http.Error(w, `{"error":"map already exists","success":false}`, http.StatusConflict)
		return
	}

	mapData, err := h.generateNewMap(req.OwnerID, req.OwnerName, req.MapID, req.Width, req.Height)
	if err != nil {
		http.Error(w, `{"error":"failed to generate map","success":false}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"map_id":  mapData.MapID,
		"data":    mapData,
	})
}

// generateNewMap generates a new map with full random terrain and elements
func (h *Handler) generateNewMap(ownerID int, ownerName, mapID string, width, height int) (*MapData, error) {
	now := time.Now()

	// Validate dimensions
	if width < 10 {
		width = 10
	}
	if width > 100 {
		width = 100
	}
	if height < 10 {
		height = 10
	}
	if height > 100 {
		height = 100
	}

	rand.Seed(time.Now().UnixNano())

	mapData := &MapData{
		Version:     "1.0",
		MapID:       mapID,
		OwnerID:     ownerID,
		OwnerName:   ownerName,
		CreatedAt:   now.Format(time.RFC3339),
		UpdatedAt:   now.Format(time.RFC3339),
		Width:       width,
		Height:      height,
		Terrain:     h.generateTerrain(width, height),
		Elements:    h.generateElements(width, height),
		Connections: MapConnection{}, // Will be updated by updateMapConnections
		Statistics:  MapStatistics{},
	}

	// Calculate statistics
	mapData.Statistics.TotalWildPokemon = 0
	mapData.Statistics.TotalFood = 0
	mapData.Statistics.TotalObstacles = 0
	for _, elem := range mapData.Elements {
		switch elem.Type {
		case "wild_pokemon":
			mapData.Statistics.TotalWildPokemon++
		case "food":
			mapData.Statistics.TotalFood++
		case "obstacle":
			mapData.Statistics.TotalObstacles++
		}
	}

	// Save map to file
	mapsDir := "../maps"
	if _, err := os.Stat(mapsDir); err != nil {
		os.MkdirAll(mapsDir, 0755)
	}

	mapPath := filepath.Join(mapsDir, fmt.Sprintf("%s.json", mapID))
	mapJSON, err := json.MarshalIndent(mapData, "", "  ")
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(mapPath, mapJSON, 0644)
	if err != nil {
		return nil, err
	}

	// Update connections for all adjacent maps
	h.updateMapConnections(mapID, width, height)

	return mapData, nil
}

// generateTerrain generates random terrain for a map
func (h *Handler) generateTerrain(width, height int) [][]int {
	terrain := make([][]int, height)
	for i := 0; i < height; i++ {
		terrain[i] = make([]int, width)
		for j := 0; j < width; j++ {
			terrain[i][j] = h.generateTerrainTile()
		}
	}
	return terrain
}

// generateTerrainTile generates a single terrain tile (0-3)
// 0=grass (70%), 1=forest (15%), 2=water (10%), 3=mountain (5%)
func (h *Handler) generateTerrainTile() int {
	r := rand.Intn(100)
	if r < 70 {
		return 0 // grass
	} else if r < 85 {
		return 1 // forest
	} else if r < 95 {
		return 2 // water
	}
	return 3 // mountain
}

// generateElements generates random elements (Pokemon, food, obstacles) for a map
func (h *Handler) generateElements(width, height int) []MapElement {
	elements := make([]MapElement, 0)

	// Pokemon list for generation
	pokemonList := []map[string]interface{}{
		{"id": "0001", "name": "妙蛙种子", "rarity": "common"},
		{"id": "0004", "name": "小火龙", "rarity": "common"},
		{"id": "0007", "name": "杰尼龟", "rarity": "common"},
		{"id": "0016", "name": "波波", "rarity": "common"},
		{"id": "0025", "name": "皮卡丘", "rarity": "common"},
		{"id": "0027", "name": "小岩鼠", "rarity": "uncommon"},
		{"id": "0029", "name": "尼多兰", "rarity": "uncommon"},
		{"id": "0035", "name": "皮皮", "rarity": "common"},
		{"id": "0041", "name": "小鼠", "rarity": "rare"},
		{"id": "0043", "name": "臭臭花", "rarity": "uncommon"},
		{"id": "0052", "name": "小猫", "rarity": "uncommon"},
		{"id": "0054", "name": "可达鸭", "rarity": "common"},
		{"id": "0058", "name": "小火狐", "rarity": "rare"},
		{"id": "0063", "name": "腹斯", "rarity": "uncommon"},
		{"id": "0066", "name": "小拳石", "rarity": "uncommon"},
		{"id": "0021", "name": "胡地", "rarity": "rare"},
	}

	foodList := []map[string]interface{}{
		{"type": "berry", "name": "树莓", "hp": 20, "qty": 3},
		{"type": "apple", "name": "苹果", "hp": 15, "qty": 2},
		{"type": "blueberry", "name": "蓝莓", "hp": 25, "qty": 5},
		{"type": "orange", "name": "橙子", "hp": 18, "qty": 3},
		{"type": "banana", "name": "香蕉", "hp": 22, "qty": 4},
		{"type": "mango", "name": "芒果", "hp": 28, "qty": 2},
	}

	obstacleList := []map[string]interface{}{
		{"type": "rock", "name": "岩石"},
		{"type": "tree", "name": "大树"},
		{"type": "bush", "name": "灌木"},
		{"type": "fence", "name": "栅栏"},
	}

	// Generate 5-10 wild Pokemon
	numPokemon := 5 + rand.Intn(6)
	for i := 0; i < numPokemon; i++ {
		pokemon := pokemonList[rand.Intn(len(pokemonList))]
		elem := MapElement{
			ID:   fmt.Sprintf("wild_%03d", i+1),
			Type: "wild_pokemon",
			X:    rand.Intn(width),
			Y:    rand.Intn(height),
			Data: map[string]interface{}{
				"level":        3 + rand.Intn(8),
				"pokemon_id":   pokemon["id"],
				"pokemon_name": pokemon["name"],
				"rarity":       pokemon["rarity"],
			},
		}
		elements = append(elements, elem)
	}

	// Generate 3-8 food items
	numFood := 3 + rand.Intn(6)
	for i := 0; i < numFood; i++ {
		food := foodList[rand.Intn(len(foodList))]
		elem := MapElement{
			ID:   fmt.Sprintf("food_%03d", i+1),
			Type: "food",
			X:    rand.Intn(width),
			Y:    rand.Intn(height),
			Data: map[string]interface{}{
				"food_name":   food["name"],
				"food_type":   food["type"],
				"quantity":    food["qty"],
				"restores_hp": food["hp"],
			},
		}
		elements = append(elements, elem)
	}

	// Generate 2-5 obstacles
	numObstacles := 2 + rand.Intn(4)
	for i := 0; i < numObstacles; i++ {
		obstacle := obstacleList[rand.Intn(len(obstacleList))]
		elem := MapElement{
			ID:   fmt.Sprintf("obstacle_%03d", i+1),
			Type: "obstacle",
			X:    rand.Intn(width),
			Y:    rand.Intn(height),
			Data: map[string]interface{}{
				"obstacle_name": obstacle["name"],
				"obstacle_type": obstacle["type"],
				"passable":      false,
			},
		}
		elements = append(elements, elem)
	}

	return elements
}

// updateMapConnections detects and updates connections to adjacent maps
func (h *Handler) updateMapConnections(mapID string, width, height int) error {
	// Parse map ID to numeric value
	mapIDNum, err := strconv.Atoi(mapID)
	if err != nil {
		return nil // Not a numeric ID, skip
	}

	// Check for adjacent maps based on ID convention:
	// East/West: ID differs by 1 (e.g., 001 and 002)
	// North/South: ID differs by 100 (e.g., 001 and 101)

	eastMapID := fmt.Sprintf("%03d", mapIDNum+1)
	westMapID := fmt.Sprintf("%03d", mapIDNum-1)
	southMapID := fmt.Sprintf("%03d", mapIDNum+100)
	northMapID := fmt.Sprintf("%03d", mapIDNum-100)

	// Load current map
	currentMap, err := h.loadMapByID(mapID)
	if err != nil {
		return nil
	}

	// Check each direction and update if map exists
	if mapIDNum > 0 && mapIDNum%100 != 1 { // Don't connect to 0XX or X01 from X02
		if _, err := h.loadMapByID(westMapID); err == nil {
			currentMap.Connections.West = &westMapID
			westMap, _ := h.loadMapByID(westMapID)
			if westMap != nil {
				currentMap.Connections.East = nil // Reset if it was set
				westMap.Connections.East = &mapID
				h.saveMapData(westMapID, westMap)
			}
		}
	}

	if _, err := h.loadMapByID(eastMapID); err == nil {
		currentMap.Connections.East = &eastMapID
		eastMap, _ := h.loadMapByID(eastMapID)
		if eastMap != nil {
			eastMap.Connections.West = &mapID
			h.saveMapData(eastMapID, eastMap)
		}
	}

	if mapIDNum > 100 {
		if _, err := h.loadMapByID(northMapID); err == nil {
			currentMap.Connections.North = &northMapID
			northMap, _ := h.loadMapByID(northMapID)
			if northMap != nil {
				northMap.Connections.South = &mapID
				h.saveMapData(northMapID, northMap)
			}
		}
	}

	if _, err := h.loadMapByID(southMapID); err == nil {
		currentMap.Connections.South = &southMapID
		southMap, _ := h.loadMapByID(southMapID)
		if southMap != nil {
			southMap.Connections.North = &mapID
			h.saveMapData(southMapID, southMap)
		}
	}

	// Save updated map
	return h.saveMapData(mapID, currentMap)
}

// saveMapData saves a map to disk and invalidates cache
func (h *Handler) saveMapData(mapID string, mapData *MapData) error {
	mapsDir := "../maps"
	if _, err := os.Stat(mapsDir); err != nil {
		os.MkdirAll(mapsDir, 0755)
	}

	mapPath := filepath.Join(mapsDir, fmt.Sprintf("%s.json", mapID))
	mapJSON, err := json.MarshalIndent(mapData, "", "  ")
	if err != nil {
		return err
	}

	// Invalidate cache for this map
	h.mapCacheMutex.Lock()
	delete(h.mapCache, mapID)
	h.mapCacheMutex.Unlock()

	return ioutil.WriteFile(mapPath, mapJSON, 0644)
}

// ClearMapCache clears the entire map cache (useful for maintenance)
func (h *Handler) ClearMapCache() {
	h.mapCacheMutex.Lock()
	h.mapCache = make(map[string]*MapCacheEntry)
	h.mapCacheMutex.Unlock()
}

// GetMapCacheStats returns cache statistics
func (h *Handler) GetMapCacheStats() map[string]interface{} {
	h.mapCacheMutex.RLock()
	defer h.mapCacheMutex.RUnlock()

	validCount := 0
	expiredCount := 0
	now := time.Now()

	for _, entry := range h.mapCache {
		if now.Before(entry.ExpiresAt) {
			validCount++
		} else {
			expiredCount++
		}
	}

	return map[string]interface{}{
		"total_cached": len(h.mapCache),
		"valid":        validCount,
		"expired":      expiredCount,
		"ttl":          h.mapCacheTTL.String(),
	}
}

// TraverseMap 玩家穿过边界进入相邻地图
func (h *Handler) TraverseMap(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, `{"error":"method not allowed","success":false}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		CurrentMapID string `json:"current_map_id"`
		Direction    string `json:"direction"` // north, south, east, west
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `{"error":"invalid request","success":false}`, http.StatusBadRequest)
		return
	}

	currentMap, err := h.loadMapByID(req.CurrentMapID)
	if err != nil {
		http.Error(w, `{"error":"current map not found","success":false}`, http.StatusNotFound)
		return
	}

	var nextMapID *string
	switch req.Direction {
	case "north":
		nextMapID = currentMap.Connections.North
	case "south":
		nextMapID = currentMap.Connections.South
	case "east":
		nextMapID = currentMap.Connections.East
	case "west":
		nextMapID = currentMap.Connections.West
	default:
		http.Error(w, `{"error":"invalid direction","success":false}`, http.StatusBadRequest)
		return
	}

	if nextMapID == nil {
		http.Error(w, `{"error":"no map in that direction","success":false}`, http.StatusNotFound)
		return
	}

	nextMap, err := h.loadMapByID(*nextMapID)
	if err != nil {
		http.Error(w, `{"error":"next map not found","success":false}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"current": req.CurrentMapID,
		"next":    *nextMapID,
		"data":    nextMap,
	})
}
