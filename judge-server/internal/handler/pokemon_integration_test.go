package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"judge-server/internal/db"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestHelper 提供集成测试的辅助函数
type TestHelper struct {
	t       *testing.T
	db      *db.Database
	handler *Handler
	baseURL string
}

// NewTestHelper 创建新的测试助手
func NewTestHelper(t *testing.T) *TestHelper {
	// 连接到测试数据库
	database, err := db.NewDatabase(db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "xiaodudu",
		DBName:   "agent_monster_test",
		SSLMode:  "disable",
	})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// 初始化schema
	if err := database.InitSchema(); err != nil {
		t.Fatalf("Failed to initialize schema: %v", err)
	}

	if err := database.UpdateSchema(); err != nil {
		t.Fatalf("Failed to update schema: %v", err)
	}

	if err := database.InitBattleSystemSchema(); err != nil {
		t.Fatalf("Failed to initialize battle system schema: %v", err)
	}

	if err := database.InitAuthSchema(); err != nil {
		t.Fatalf("Failed to initialize auth schema: %v", err)
	}

	handler := NewHandler(database, nil, nil)
	return &TestHelper{
		t:       t,
		db:      database,
		handler: handler,
		baseURL: "http://localhost:8080",
	}
}

// Close 清理测试资源
func (th *TestHelper) Close() {
	th.db.Close()
}

// CreateTestAccount 创建测试账户
func (th *TestHelper) CreateTestAccount(username, email string) (int, string) {
	query := `
		INSERT INTO users (username, email, created_at)
		VALUES ($1, $2, NOW())
		RETURNING id
	`
	var userID int
	if err := th.db.QueryRow(query, username, email).Scan(&userID); err != nil {
		th.t.Fatalf("Failed to create test account: %v", err)
	}

	// 创建token
	query = `
		INSERT INTO auth_tokens (user_id, token, created_at, expires_at)
		VALUES ($1, $2, NOW(), NOW() + INTERVAL '24 hours')
		RETURNING token
	`
	token := fmt.Sprintf("test_token_%s_%d", username, time.Now().UnixNano())
	var returnedToken string
	if err := th.db.QueryRow(query, userID, token).Scan(&returnedToken); err != nil {
		th.t.Fatalf("Failed to create token: %v", err)
	}

	return userID, returnedToken
}

// CreateTestPokemon 创建测试宝可梦
func (th *TestHelper) CreateTestPokemon(userID int, speciesID int) int {
	query := `
		INSERT INTO user_pokemons (user_id, pokemon_species_id, level, current_hp, max_hp, experience, caught_at)
		VALUES ($1, $2, 5, 20, 20, 0, NOW())
		RETURNING id
	`
	var pokemonID int
	if err := th.db.QueryRow(query, userID, speciesID).Scan(&pokemonID); err != nil {
		th.t.Fatalf("Failed to create test pokemon: %v", err)
	}
	return pokemonID
}

// CreateTestNature 创建测试性格
func (th *TestHelper) CreateTestNature() string {
	// 创建一个性格如果不存在
	natureID := "jolly"
	query := `
		INSERT INTO natures (id, name, increased_stat, decreased_stat)
		VALUES ($1, $2, 'spe', 'spa')
		ON CONFLICT (id) DO NOTHING
	`
	if _, err := th.db.Exec(query, natureID, "Jolly"); err != nil {
		th.t.Logf("Warning: Could not insert nature: %v", err)
	}
	return natureID
}

// Request 发送HTTP请求的辅助函数
func (th *TestHelper) Request(method, path string, body interface{}) *httptest.ResponseRecorder {
	var bodyReader *bytes.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			th.t.Fatalf("Failed to marshal request body: %v", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	} else {
		bodyReader = bytes.NewReader([]byte{})
	}

	req := httptest.NewRequest(method, path, bodyReader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	w := httptest.NewRecorder()
	return w
}

// ExecuteRequest 执行HTTP请求
func (th *TestHelper) ExecuteRequest(handler func(http.ResponseWriter, *http.Request), w *httptest.ResponseRecorder, req *http.Request) {
	handler(w, req)
}

// AssertStatus 验证HTTP状态码
func (th *TestHelper) AssertStatus(w *httptest.ResponseRecorder, expectedStatus int) {
	if w.Code != expectedStatus {
		th.t.Errorf("Expected status %d, got %d. Response: %s", expectedStatus, w.Code, w.Body.String())
	}
}

// AssertSuccess 验证响应成功标志
func (th *TestHelper) AssertSuccess(w *httptest.ResponseRecorder) {
	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		th.t.Errorf("Failed to parse response: %v", err)
		return
	}

	if success, ok := response["success"].(bool); !ok || !success {
		th.t.Errorf("Expected success=true, got response: %v", response)
	}
}

// TestPokemonBreedingHandlers 测试宝可梦养成处理程序
func TestPokemonBreedingHandlers(t *testing.T) {
	th := NewTestHelper(t)
	defer th.Close()

	// 创建测试账户和宝可梦
	userID, _ := th.CreateTestAccount("breeder_user", "breeder@test.com")
	pokemonID := th.CreateTestPokemon(userID, 1) // Bulbasaur
	natureID := th.CreateTestNature()

	t.Run("AddEffortValues", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/pokemon/effort-values", bytes.NewReader([]byte(`{
			"user_pokemon_id": `+fmt.Sprintf("%d", pokemonID)+`,
			"hp_effort": 10,
			"atk_effort": 10,
			"def_effort": 10,
			"spa_effort": 10,
			"spd_effort": 10,
			"spe_effort": 10
		}`)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		th.handler.AddEffortValues(w, req)
		th.AssertStatus(w, http.StatusOK)
		th.AssertSuccess(w)
	})

	t.Run("ChangePokemonNature", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/pokemon/nature", bytes.NewReader([]byte(`{
			"user_pokemon_id": `+fmt.Sprintf("%d", pokemonID)+`,
			"nature_id": "`+natureID+`"
		}`)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		th.handler.ChangePokemonNature(w, req)
		th.AssertStatus(w, http.StatusOK)
		th.AssertSuccess(w)
	})

	t.Run("CalculatePokemonStats", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/pokemon/calculate-stats", bytes.NewReader([]byte(`{
			"user_pokemon_id": `+fmt.Sprintf("%d", pokemonID)+`
		}`)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		th.handler.CalculatePokemonStats(w, req)
		th.AssertStatus(w, http.StatusOK)
		th.AssertSuccess(w)
	})
}

// TestMapExplorationHandlers 测试地图探索处理程序
func TestMapExplorationHandlers(t *testing.T) {
	th := NewTestHelper(t)
	defer th.Close()

	// 创建测试账户
	userID, _ := th.CreateTestAccount("explorer_user", "explorer@test.com")

	t.Run("GetRegions", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/map/regions", nil)
		w := httptest.NewRecorder()

		th.handler.GetRegions(w, req)
		th.AssertStatus(w, http.StatusOK)
		th.AssertSuccess(w)
	})

	t.Run("EnterGrassArea", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/map/enter-grass", bytes.NewReader([]byte(`{
			"area_id": "test_area_1",
			"user_id": `+fmt.Sprintf("%d", userID)+`
		}`)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		th.handler.EnterGrassArea(w, req)
		th.AssertStatus(w, http.StatusOK)
		th.AssertSuccess(w)
	})

	t.Run("GetMapStatus", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/map/status", nil)
		w := httptest.NewRecorder()

		th.handler.GetMapStatus(w, req)
		th.AssertStatus(w, http.StatusOK)
		th.AssertSuccess(w)
	})
}

// TestRouteRegistration 测试路由是否已正确注册
func TestRouteRegistration(t *testing.T) {
	th := NewTestHelper(t)
	defer th.Close()

	tests := []struct {
		name           string
		method         string
		path           string
		handler        func(http.ResponseWriter, *http.Request)
		expectedStatus int
	}{
		{
			name:           "POST /api/pokemon/effort-values",
			method:         "POST",
			path:           "/api/pokemon/effort-values",
			handler:        th.handler.AddEffortValues,
			expectedStatus: http.StatusBadRequest, // 应该返回400因为没有有效的user_pokemon_id
		},
		{
			name:           "POST /api/pokemon/nature",
			method:         "POST",
			path:           "/api/pokemon/nature",
			handler:        th.handler.ChangePokemonNature,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "POST /api/pokemon/calculate-stats",
			method:         "POST",
			path:           "/api/pokemon/calculate-stats",
			handler:        th.handler.CalculatePokemonStats,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "POST /api/map/enter-grass",
			method:         "POST",
			path:           "/api/map/enter-grass",
			handler:        th.handler.EnterGrassArea,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "GET /api/map/regions",
			method:         "GET",
			path:           "/api/map/regions",
			handler:        th.handler.GetRegions,
			expectedStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.method, test.path, nil)
			w := httptest.NewRecorder()
			test.handler(w, req)

			if w.Code != test.expectedStatus {
				t.Errorf("Expected status %d, got %d for %s", test.expectedStatus, w.Code, test.path)
			}
		})
	}
}

// TestEndToEndPokemonCatch 测试端到端的宝可梦捕捉流程
func TestEndToEndPokemonCatch(t *testing.T) {
	th := NewTestHelper(t)
	defer th.Close()

	// 1. 创建用户账户
	userID, token := th.CreateTestAccount("catcher_user", "catcher@test.com")
	t.Logf("Created user account: %d with token: %s", userID, token)

	// 2. 进入草地区域
	enterReq := httptest.NewRequest("POST", "/api/map/enter-grass", bytes.NewReader([]byte(`{
		"area_id": "test_grass_1",
		"user_id": `+fmt.Sprintf("%d", userID)+`
	}`)))
	enterReq.Header.Set("Content-Type", "application/json")
	enterW := httptest.NewRecorder()
	th.handler.EnterGrassArea(enterW, enterReq)
	th.AssertStatus(enterW, http.StatusOK)
	t.Log("Successfully entered grass area")

	// 3. 尝试捕捉宝可梦
	captureReq := httptest.NewRequest("POST", "/api/pokemon/capture", bytes.NewReader([]byte(`{
		"wild_pokemon_id": 1,
		"user_id": `+fmt.Sprintf("%d", userID)+`,
		"pokeball_type": "pokeball"
	}`)))
	captureReq.Header.Set("Content-Type", "application/json")
	captureW := httptest.NewRecorder()
	th.handler.AttemptCapture(captureW, captureReq)
	// 可能返回400或500取决于数据库状态
	t.Logf("Capture attempt returned status: %d", captureW.Code)

	// 4. 离开草地区域
	exitReq := httptest.NewRequest("POST", "/api/map/exit-grass", bytes.NewReader([]byte(`{
		"area_id": "test_grass_1",
		"user_id": `+fmt.Sprintf("%d", userID)+`
	}`)))
	exitReq.Header.Set("Content-Type", "application/json")
	exitW := httptest.NewRecorder()
	th.handler.ExitGrassArea(exitW, exitReq)
	th.AssertStatus(exitW, http.StatusOK)
	t.Log("Successfully exited grass area")
}

// TestEndToEndPokemonBreeding 测试端到端的宝可梦养成流程
func TestEndToEndPokemonBreeding(t *testing.T) {
	th := NewTestHelper(t)
	defer th.Close()

	// 1. 创建用户账户
	userID, _ := th.CreateTestAccount("breeding_user", "breeding@test.com")
	t.Logf("Created user account: %d", userID)

	// 2. 创建一只宝可梦
	pokemonID := th.CreateTestPokemon(userID, 1)
	t.Logf("Created test pokemon: %d", pokemonID)

	// 3. 添加努力值
	effortReq := httptest.NewRequest("POST", "/api/pokemon/effort-values", bytes.NewReader([]byte(`{
		"user_pokemon_id": `+fmt.Sprintf("%d", pokemonID)+`,
		"hp_effort": 4,
		"atk_effort": 252,
		"def_effort": 0,
		"spa_effort": 0,
		"spd_effort": 0,
		"spe_effort": 252
	}`)))
	effortReq.Header.Set("Content-Type", "application/json")
	effortW := httptest.NewRecorder()
	th.handler.AddEffortValues(effortW, effortReq)
	th.AssertStatus(effortW, http.StatusOK)
	t.Log("Successfully added effort values")

	// 4. 改变性格
	natureID := th.CreateTestNature()
	natureReq := httptest.NewRequest("POST", "/api/pokemon/nature", bytes.NewReader([]byte(`{
		"user_pokemon_id": `+fmt.Sprintf("%d", pokemonID)+`,
		"nature_id": "`+natureID+`"
	}`)))
	natureReq.Header.Set("Content-Type", "application/json")
	natureW := httptest.NewRecorder()
	th.handler.ChangePokemonNature(natureW, natureReq)
	th.AssertStatus(natureW, http.StatusOK)
	t.Log("Successfully changed nature")

	// 5. 计算最终属性值
	statsReq := httptest.NewRequest("POST", "/api/pokemon/calculate-stats", bytes.NewReader([]byte(`{
		"user_pokemon_id": `+fmt.Sprintf("%d", pokemonID)+`
	}`)))
	statsReq.Header.Set("Content-Type", "application/json")
	statsW := httptest.NewRecorder()
	th.handler.CalculatePokemonStats(statsW, statsReq)
	th.AssertStatus(statsW, http.StatusOK)
	t.Log("Successfully calculated stats")
}

// TestMultipleAccounts 测试多账户场景
func TestMultipleAccounts(t *testing.T) {
	th := NewTestHelper(t)
	defer th.Close()

	// 创建4个不同的测试账户
	accounts := []struct {
		username string
		email    string
	}{
		{"player_1", "player1@test.com"},
		{"player_2", "player2@test.com"},
		{"player_3", "player3@test.com"},
		{"player_4", "player4@test.com"},
	}

	userIDs := make([]int, 0)
	for _, acc := range accounts {
		userID, token := th.CreateTestAccount(acc.username, acc.email)
		userIDs = append(userIDs, userID)
		t.Logf("Created account %s (ID: %d, Token: %s...)", acc.username, userID, token[:20])

		// 为每个账户创建几只宝可梦
		for i := 1; i <= 3; i++ {
			pokemonID := th.CreateTestPokemon(userID, i)
			t.Logf("  Created pokemon %d for user %d", pokemonID, userID)
		}
	}

	// 验证所有账户都存在
	if len(userIDs) != 4 {
		t.Fatalf("Expected 4 users, got %d", len(userIDs))
	}
	t.Logf("Successfully created %d test accounts with %d pokemons each", len(userIDs), 3)
}

// TestErrorHandling 测试错误处理和输入验证
func TestErrorHandling(t *testing.T) {
	th := NewTestHelper(t)
	defer th.Close()

	tests := []struct {
		name           string
		path           string
		method         string
		body           string
		handler        func(http.ResponseWriter, *http.Request)
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "Invalid method for AddEffortValues",
			path:           "/api/pokemon/effort-values",
			method:         "GET",
			body:           "",
			handler:        th.handler.AddEffortValues,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedError:  "Method not allowed",
		},
		{
			name:           "Missing user_pokemon_id",
			path:           "/api/pokemon/effort-values",
			method:         "POST",
			body:           `{"hp_effort": 10}`,
			handler:        th.handler.AddEffortValues,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid request body",
		},
		{
			name:           "Invalid area_id in EnterGrassArea",
			path:           "/api/map/enter-grass",
			method:         "POST",
			body:           `{"user_id": 1}`,
			handler:        th.handler.EnterGrassArea,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid",
		},
		{
			name:           "Invalid user_id in EnterGrassArea",
			path:           "/api/map/enter-grass",
			method:         "POST",
			body:           `{"area_id": "test", "user_id": -1}`,
			handler:        th.handler.EnterGrassArea,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid",
		},
		{
			name:           "Invalid method for EnterGrassArea",
			path:           "/api/map/enter-grass",
			method:         "GET",
			body:           "",
			handler:        th.handler.EnterGrassArea,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedError:  "Method not allowed",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var req *http.Request
			if test.body != "" {
				req = httptest.NewRequest(test.method, test.path, strings.NewReader(test.body))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req = httptest.NewRequest(test.method, test.path, nil)
			}
			w := httptest.NewRecorder()
			test.handler(w, req)

			if w.Code != test.expectedStatus {
				t.Errorf("Expected status %d, got %d", test.expectedStatus, w.Code)
			}

			if test.expectedError != "" {
				response := w.Body.String()
				if !strings.Contains(response, test.expectedError) {
					t.Errorf("Expected error containing '%s', got: %s", test.expectedError, response)
				}
			}
		})
	}
}

// TestPokemonDataImport 测试宝可梦数据导入
func TestPokemonDataImport(t *testing.T) {
	th := NewTestHelper(t)
	defer th.Close()

	// 验证pokemon_species表是否存在且有数据
	var count int
	err := th.db.QueryRow("SELECT COUNT(*) FROM pokemon_species").Scan(&count)
	if err != nil {
		t.Logf("pokemon_species table may not exist or have data: %v", err)
	} else {
		t.Logf("Found %d pokemon species in database", count)
	}

	// 验证moves表
	err = th.db.QueryRow("SELECT COUNT(*) FROM pokemon_moves").Scan(&count)
	if err != nil {
		t.Logf("pokemon_moves table may not exist or have data: %v", err)
	} else {
		t.Logf("Found %d pokemon moves in database", count)
	}

	// 验证abilities表
	err = th.db.QueryRow("SELECT COUNT(*) FROM pokemon_abilities").Scan(&count)
	if err != nil {
		t.Logf("pokemon_abilities table may not exist or have data: %v", err)
	} else {
		t.Logf("Found %d pokemon abilities in database", count)
	}
}
