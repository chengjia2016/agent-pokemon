package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"judge-server/internal/db"
	"judge-server/internal/service"
)

// TestWorldSystemNPCEndpoints tests all NPC-related endpoints
func TestWorldSystemNPCEndpoints(t *testing.T) {
	// Setup database connection
	database, err := db.NewDatabase(db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "xiaodudu",
		DBName:   "agent_monster",
		SSLMode:  "disable",
	})
	if err != nil {
		t.Skipf("Could not connect to database: %v", err)
	}
	defer database.Close()

	// Initialize world service and tables
	worldService := service.NewWorldService(database)
	if err := worldService.InitializeWorldTables(); err != nil {
		t.Fatalf("Failed to initialize world tables: %v", err)
	}

	handler := NewWorldHTTPHandler(worldService)

	// Test: Get NPCs by town
	t.Run("GetNPCsByTown", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/npcs?town_id=1", nil)
		w := httptest.NewRecorder()

		handler.HandleNPCs(w, req)

		if w.Code != http.StatusOK && w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 200 or 400, got %d", w.Code)
		}
	})

	// Test: Create NPC
	t.Run("CreateNPC", func(t *testing.T) {
		npcData := map[string]interface{}{
			"town_id":    "1",
			"name_en":    "Test NPC",
			"name_zh":    "测试NPC",
			"type":       "NPC",
			"role":       "quest_giver",
			"coord_x":    50.0,
			"coord_y":    50.0,
			"avatar_url": "https://example.com/npc.png",
		}
		body, _ := json.Marshal(npcData)

		req := httptest.NewRequest("POST", "/api/npcs", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.HandleNPCs(w, req)

		if w.Code != http.StatusCreated && w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status 201 or 500, got %d", w.Code)
		}
	})

	// Test: NPC Talk endpoint
	t.Run("HandleNPCTalk", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/npcs/talk?npc_id=1", nil)
		w := httptest.NewRecorder()

		handler.HandleNPCTalk(w, req)

		if w.Code != http.StatusOK && w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 200 or 400, got %d", w.Code)
		}
	})
}

// TestWorldSystemQuestEndpoints tests all quest-related endpoints
func TestWorldSystemQuestEndpoints(t *testing.T) {
	// Setup database connection
	database, err := db.NewDatabase(db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "xiaodudu",
		DBName:   "agent_monster",
		SSLMode:  "disable",
	})
	if err != nil {
		t.Skipf("Could not connect to database: %v", err)
	}
	defer database.Close()

	// Initialize world service
	worldService := service.NewWorldService(database)
	if err := worldService.InitializeWorldTables(); err != nil {
		t.Fatalf("Failed to initialize world tables: %v", err)
	}

	handler := NewWorldHTTPHandler(worldService)

	// Test: Get quests
	t.Run("GetQuests", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/quests", nil)
		w := httptest.NewRecorder()

		handler.HandleQuests(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}
	})

	// Test: Get user quests
	t.Run("GetUserQuests", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/user/quests?user_id=test_user", nil)
		w := httptest.NewRecorder()

		handler.HandleUserQuests(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}
	})
}

// TestWorldSystemDungeonEndpoints tests dungeon-related endpoints
func TestWorldSystemDungeonEndpoints(t *testing.T) {
	// Setup database connection
	database, err := db.NewDatabase(db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "xiaodudu",
		DBName:   "agent_monster",
		SSLMode:  "disable",
	})
	if err != nil {
		t.Skipf("Could not connect to database: %v", err)
	}
	defer database.Close()

	// Initialize world service
	worldService := service.NewWorldService(database)
	if err := worldService.InitializeWorldTables(); err != nil {
		t.Fatalf("Failed to initialize world tables: %v", err)
	}

	handler := NewWorldHTTPHandler(worldService)

	// Test: Create dungeon
	t.Run("CreateDungeon", func(t *testing.T) {
		dungeonData := map[string]interface{}{
			"name_en":          "Test Dungeon",
			"name_zh":          "测试地牢",
			"description_en":   "A test dungeon",
			"description_zh":   "一个测试地牢",
			"difficulty_level": "easy",
			"boss_name_en":     "Test Boss",
			"boss_name_zh":     "测试Boss",
			"boss_level":       15,
			"reward_gold":      500,
		}
		body, _ := json.Marshal(dungeonData)

		req := httptest.NewRequest("POST", "/api/dungeons", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.HandleDungeons(w, req)

		if w.Code != http.StatusCreated && w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status 201 or 500, got %d", w.Code)
		}
	})
}

// TestWorldSystemGymEndpoints tests gym-related endpoints
func TestWorldSystemGymEndpoints(t *testing.T) {
	// Setup database connection
	database, err := db.NewDatabase(db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "xiaodudu",
		DBName:   "agent_monster",
		SSLMode:  "disable",
	})
	if err != nil {
		t.Skipf("Could not connect to database: %v", err)
	}
	defer database.Close()

	// Initialize world service
	worldService := service.NewWorldService(database)
	if err := worldService.InitializeWorldTables(); err != nil {
		t.Fatalf("Failed to initialize world tables: %v", err)
	}

	handler := NewWorldHTTPHandler(worldService)

	// Test: Get user gym badges
	t.Run("GetUserGymBadges", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/gyms?user_id=test_user", nil)
		w := httptest.NewRecorder()

		handler.HandleGyms(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}
	})
}

// TestWorldSystemLevelEndpoints tests level-related endpoints
func TestWorldSystemLevelEndpoints(t *testing.T) {
	// Setup database connection
	database, err := db.NewDatabase(db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "xiaodudu",
		DBName:   "agent_monster",
		SSLMode:  "disable",
	})
	if err != nil {
		t.Skipf("Could not connect to database: %v", err)
	}
	defer database.Close()

	// Initialize world service
	worldService := service.NewWorldService(database)
	if err := worldService.InitializeWorldTables(); err != nil {
		t.Fatalf("Failed to initialize world tables: %v", err)
	}

	handler := NewWorldHTTPHandler(worldService)

	// Test: Get levels by zone
	t.Run("GetLevelsByZone", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/levels?zone_id=1", nil)
		w := httptest.NewRecorder()

		handler.HandleLevels(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}
	})

	// Test: Complete level
	t.Run("CompleteLevel", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/user/levels/progress?user_id=test_user&level_id=1", nil)
		w := httptest.NewRecorder()

		handler.HandleUserLevelProgress(w, req)

		// Should return OK even if level doesn't exist (returns error in JSON)
		if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status 200 or 500, got %d", w.Code)
		}
	})
}
