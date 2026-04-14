package main

import (
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/github"
	"judge-server/internal/handler"
	"judge-server/internal/middleware"
	"judge-server/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
	GitHub struct {
		Token string `yaml:"token"`
		Owner string `yaml:"owner"`
		Repo  string `yaml:"repo"`
	} `yaml:"github"`
	Settlement struct {
		DailyTime           string `yaml:"daily_time"`
		SyncToGitHub        bool   `yaml:"sync_to_github"`
		SyncIntervalSeconds int    `yaml:"sync_interval_seconds"`
	} `yaml:"settlement"`
	RateLimit struct {
		RequestsPerMinute int  `yaml:"requests_per_minute"`
		BurstSize         int  `yaml:"burst_size"`
		Enabled           bool `yaml:"enabled"`
	} `yaml:"rate_limit"`
}

func main() {
	configPath := ".config/config.yaml"
	if path := os.Getenv("CONFIG_PATH"); path != "" {
		configPath = path
	}

	config, err := loadConfig(configPath)
	if err != nil {
		log.Printf("Warning: Could not load config: %v, using defaults", err)
		config = getDefaultConfig()
	}

	database, err := db.NewDatabase(db.Config{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		User:     config.Database.User,
		Password: config.Database.Password,
		DBName:   config.Database.DBName,
		SSLMode:  config.Database.SSLMode,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	if err := database.InitSchema(); err != nil {
		log.Fatalf("Failed to initialize schema: %v", err)
	}

	// Initialize user account tables
	if err := database.UpdateSchema(); err != nil {
		log.Fatalf("Failed to update schema with user tables: %v", err)
	}

	// Initialize battle system schema
	if err := database.InitBattleSystemSchema(); err != nil {
		log.Fatalf("Failed to initialize battle system schema: %v", err)
	}

	// Initialize world system schema
	worldService := service.NewWorldService(database)
	if err := worldService.InitializeWorldTables(); err != nil {
		log.Fatalf("Failed to initialize world system tables: %v", err)
	}

	var syncService *github.SyncService
	var batchSync *service.BatchSyncService

	if config.GitHub.Token != "" && config.Settlement.SyncToGitHub {
		syncService = github.NewSyncService(github.Config{
			Token: config.GitHub.Token,
			Owner: config.GitHub.Owner,
			Repo:  config.GitHub.Repo,
		})
		log.Printf("GitHub sync enabled for %s/%s", config.GitHub.Owner, config.GitHub.Repo)

		interval := config.Settlement.SyncIntervalSeconds
		if interval <= 0 {
			interval = 300
		}
		batchSync = service.NewBatchSyncService(database, syncService, interval)
		go batchSync.Start()
		log.Printf("Batch sync interval: %d seconds", interval)
	}

	h := handler.NewHandler(database, syncService, batchSync)

	// Root path - serve skill.md content
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// Read skill.md file
		skillMDPath := "../../skill.md"
		content, err := os.ReadFile(skillMDPath)
		if err != nil {
			// Fallback to alternative paths
			skillMDPath = "/root/petskill/skill.md"
			content, err = os.ReadFile(skillMDPath)
			if err != nil {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Agent Monster - Gameplay Skill & Entry Point</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; line-height: 1.6; }
        a { color: #0066cc; }
    </style>
</head>
<body>
    <h1>Agent Monster - Gameplay Skill & Entry Point</h1>
    <p>Please view the documentation at:</p>
    <p><a href="https://github.com/chengjia2016/agent-pokemon/blob/main/skill.md">https://github.com/chengjia2016/agent-pokemon/blob/main/skill.md</a></p>
    <p>Or use the API endpoints at <strong>http://pokemon.openx.pro:10000/api/</strong></p>
</body>
</html>
`)
				return
			}
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=3600")

		// Build HTML page with skill.md content
		htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Agent Monster - Gameplay Skill & Entry Point</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif;
            line-height: 1.6;
            max-width: 900px;
            margin: 0 auto;
            padding: 20px;
            background-color: #fff;
            color: #333;
        }
        a { color: #0066cc; text-decoration: none; }
        a:hover { text-decoration: underline; }
        code { background-color: #f6f8fa; padding: 2px 6px; border-radius: 3px; font-family: "SFMono-Regular", Consolas, "Liberation Mono", Menlo, monospace; }
        pre {
            background-color: #f6f8fa;
            padding: 16px;
            border-radius: 6px;
            overflow-x: auto;
        }
        h1, h2, h3 { border-bottom: 1px solid #eaecef; padding-bottom: 0.3em; }
        h1 { font-size: 2em; }
        h2 { font-size: 1.5em; margin-top: 1.5em; }
        h3 { font-size: 1.25em; }
        table { border-collapse: collapse; width: 100%%; }
        table, th, td { border: 1px solid #ddd; padding: 12px; }
        th { background-color: #f6f8fa; }
        blockquote { padding: 0 1em; margin: 0; color: #6a737d; border-left: 0.25em solid #dfe2e5; }
    </style>
</head>
<body>
    <pre style="white-space: pre-wrap; word-wrap: break-word;">%s</pre>
</body>
</html>`, string(content))

		w.Write([]byte(htmlContent))
	})

	// Existing endpoints
	http.HandleFunc("/health", h.Health)
	http.HandleFunc("/api/pet/validate", h.ValidatePet)
	http.HandleFunc("/api/battle/validate", h.ValidateBattle)
	http.HandleFunc("/api/leaderboard", h.GetLeaderboard)
	http.HandleFunc("/api/leaderboard/daily", h.GenerateDailyLeaderboard)
	http.HandleFunc("/api/food/record", h.RecordFood)
	http.HandleFunc("/api/food/validate", h.ValidateFood)
	http.HandleFunc("/api/growth/record", h.RecordGrowth)
	http.HandleFunc("/api/egg/incubate", h.ValidateEggIncubation)
	http.HandleFunc("/api/capture/validate", h.ValidateCapture)

	// User Account Management endpoints (register longer paths first!)
	http.HandleFunc("/api/user/api-key/rotate", h.RotateAPIKey)
	http.HandleFunc("/api/user/api-key/history", h.GetAPIKeyHistory)
	http.HandleFunc("/api/user/init-config", h.InitUserConfig)
	http.HandleFunc("/api/user/balance/get", h.GetUserBalance)
	http.HandleFunc("/api/user/balance/update", h.UpdateUserBalance)
	http.HandleFunc("/api/user/pokemons/get", h.GetUserPokemons)
	http.HandleFunc("/api/user/pokemons/add", h.AddUserPokemon)
	http.HandleFunc("/api/user/inventory/get", h.GetUserInventory)
	http.HandleFunc("/api/user/inventory/add", h.AddUserItem)
	http.HandleFunc("/api/user/transactions/get", h.GetUserTransactions)
	http.HandleFunc("/api/users/create", h.CreateUserAccount)
	http.HandleFunc("/api/users/", h.GetUserAccount)

	// Farm Management endpoints (register longer paths first!)
	http.HandleFunc("/api/farms/create", h.CreateFarm)
	http.HandleFunc("/api/farms/search", h.SearchFarms)
	// /api/farms/ catches /api/farms/{id}, /api/farms/{id}/foods, etc.
	http.HandleFunc("/api/farms/", func(w http.ResponseWriter, r *http.Request) {
		// Route to appropriate handler based on path
		if strings.Contains(r.URL.Path, "/foods/consume") {
			h.ConsumeFood(w, r)
		} else if strings.Contains(r.URL.Path, "/foods") && r.Method == http.MethodPost {
			h.AddFoodToFarm(w, r)
		} else if strings.Contains(r.URL.Path, "/statistics") {
			h.GetFarmStatistics(w, r)
		} else if r.Method == http.MethodDelete {
			h.DeleteFarm(w, r)
		} else {
			h.GetFarm(w, r)
		}
	})

	// Cookie Management endpoints
	http.HandleFunc("/api/cookies/register", h.RegisterCookie)
	http.HandleFunc("/api/cookies/claim", h.ClaimCookie)
	http.HandleFunc("/api/cookies/statistics", h.GetCookieStatistics)
	http.HandleFunc("/api/cookies/scan", h.ScanCookies)

	// Egg Management endpoints (register longer paths first!)
	http.HandleFunc("/api/eggs/create", h.CreateEgg)
	http.HandleFunc("/api/eggs/statistics", h.GetEggStatistics)
	// /api/eggs/ catches /api/eggs/{id} and /api/eggs/{id}/hatch
	http.HandleFunc("/api/eggs/", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/hatch") {
			h.HatchEgg(w, r)
		} else {
			h.GetEgg(w, r)
		}
	})

	// Shop Management endpoints
	http.HandleFunc("/api/shop/items", h.ListShopItems)
	http.HandleFunc("/api/shop/buy", h.BuyItem)

	// Pokemon API routes
	http.HandleFunc("/api/pokemons", h.ListPokemonSpecies)
	http.HandleFunc("/api/pokemons/search", h.SearchPokemon)
	http.HandleFunc("/api/pokemons/pokedex", h.GetPokemonPokedex)
	http.HandleFunc("/api/shop/statistics", h.GetShopStatistics)
	http.HandleFunc("/api/shop/transactions", h.GetTransactionHistory)

	// Battle System API routes (register longer paths first!)
	// Battle endpoints
	http.HandleFunc("/api/battles/start", h.StartBattle)
	http.HandleFunc("/api/battles/stats", h.GetBattleStats)
	http.HandleFunc("/api/battles/", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/round") {
			h.ExecuteBattleRound(w, r)
		} else if strings.Contains(r.URL.Path, "/end") {
			h.EndBattle(w, r)
		} else if r.Method == http.MethodGet && strings.Count(r.URL.Path, "/") == 3 {
			// GET /api/battles/{battleId}
			h.GetBattle(w, r)
		} else if r.Method == http.MethodGet {
			// GET /api/battles
			h.GetUserBattles(w, r)
		}
	})

	// Defense System endpoints
	http.HandleFunc("/api/defense/base", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreateBase(w, r)
		} else if r.Method == http.MethodGet {
			h.GetBase(w, r)
		}
	})
	http.HandleFunc("/api/defense/defenders", h.GetDefenders)
	http.HandleFunc("/api/defense/history", h.GetDefenseHistory)
	http.HandleFunc("/api/defense/stats", h.GetDefenseStats)

	// Wild Pokemon endpoints
	http.HandleFunc("/api/wild-pokemon/capture", h.CapturePokemon)
	http.HandleFunc("/api/wild-pokemon/capture-history", h.GetCaptureHistory)
	http.HandleFunc("/api/wild-pokemon", h.ListWildPokemon)

	// Pokemon Management endpoints (register longer paths first!)
	http.HandleFunc("/api/pokemon/recovery-status", h.GetPokemonRecoveryStatus)
	http.HandleFunc("/api/pokemon/", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/train") {
			h.TrainPokemon(w, r)
		} else if r.Method == http.MethodDelete {
			h.ReleasePokemon(w, r)
		} else if r.Method == http.MethodGet && strings.Count(r.URL.Path, "/") == 3 {
			// GET /api/pokemon/{petId}
			h.GetPokemonDetails(w, r)
		} else if r.Method == http.MethodGet {
			// GET /api/pokemon
			h.ListUserPokemon(w, r)
		}
	})

	// Map System endpoints (register longer paths first!)
	http.HandleFunc("/api/maps/search", h.SearchMaps)
	http.HandleFunc("/api/maps/generate", h.GenerateUserMap)
	http.HandleFunc("/api/maps/traverse", h.TraverseMap)
	http.HandleFunc("/api/maps/", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/connections") {
			h.GetMapConnections(w, r)
		} else if strings.Contains(r.URL.Path, "/elements") {
			h.GetMapElements(w, r)
		} else if r.Method == http.MethodGet {
			// GET /api/maps/{mapId}
			h.GetMap(w, r)
		}
	})
	// GET /api/maps - List all maps
	http.HandleFunc("/api/maps", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && len(strings.TrimPrefix(r.URL.Path, "/api/maps")) == 0 {
			h.ListMaps(w, r)
		} else {
			h.GetMap(w, r)
		}
	})

	// World System endpoints (register longer paths first!)
	http.HandleFunc("/api/npcs/talk", h.HandleWorldNPCTalk)
	http.HandleFunc("/api/npcs", h.HandleWorldNPCs)
	http.HandleFunc("/api/user/quests", h.HandleWorldUserQuests)
	http.HandleFunc("/api/quests", h.HandleWorldQuests)
	http.HandleFunc("/api/quests/complete", h.HandleCompleteQuest)
	http.HandleFunc("/api/dungeons", h.HandleWorldDungeons)
	http.HandleFunc("/api/gyms", h.HandleWorldGyms)
	http.HandleFunc("/api/map/zones", h.HandleWorldMapZones)
	http.HandleFunc("/api/user/levels/progress", h.HandleWorldUserLevelProgress)
	http.HandleFunc("/api/levels", h.HandleWorldLevels)

	// Language Management endpoints (register longer paths first!)
	http.HandleFunc("/api/language/select", h.SetUserLanguage)
	http.HandleFunc("/api/language/current", h.GetUserLanguage)
	http.HandleFunc("/api/language/strings", h.GetUIStrings)
	http.HandleFunc("/api/language/list", h.GetAvailableLanguages)
	http.HandleFunc("/api/language/npc-dialogue", h.GetNPCDialogueLocalized)
	http.HandleFunc("/api/language/quest", h.GetQuestLocalized)

	// Masters EX System endpoints (register longer paths first!)
	// Initialize Masters EX database schema
	if err := database.InitMastersExSchema(); err != nil {
		log.Printf("Warning: Failed to initialize Masters EX schema: %v (may already exist)", err)
	}

	// Sync Pair Management endpoints (register longer paths first!)
	http.HandleFunc("/api/masters-ex/sync-pairs/team", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.HandleCreateOrUpdateTeam(w, r)
		} else if r.Method == http.MethodGet {
			h.HandleGetTeam(w, r)
		}
	})
	http.HandleFunc("/api/masters-ex/sync-pairs/stats", h.HandleGetSyncPairStats)
	http.HandleFunc("/api/masters-ex/sync-pairs/sync-move", h.HandleUseSyncMove)
	http.HandleFunc("/api/masters-ex/sync-pairs/dex", h.HandleGetSyncPairDex)
	http.HandleFunc("/api/masters-ex/sync-pairs/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.HandleCreateSyncPair(w, r)
		} else if r.Method == http.MethodGet {
			h.HandleGetSyncPair(w, r)
		} else if r.Method == http.MethodPut {
			h.HandleUpdateSyncPair(w, r)
		}
	})

	// Battle System endpoints (3v3 battles)
	http.HandleFunc("/api/masters-ex/battles/start", h.HandleStartBattle)
	http.HandleFunc("/api/masters-ex/battles/round", h.HandleExecuteBattleRound)
	http.HandleFunc("/api/masters-ex/battles/end", h.HandleEndBattle)
	http.HandleFunc("/api/masters-ex/battles/stats", h.HandleGetPlayerBattleStats)
	http.HandleFunc("/api/masters-ex/battles/history", h.HandleGetPlayerBattles)
	http.HandleFunc("/api/masters-ex/battles/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.HandleGetBattleSession(w, r)
		}
	})

	// Tournament System endpoints (register longer paths first!)
	http.HandleFunc("/api/masters-ex/tournament/season/create", h.HandleCreateTournamentSeason)
	http.HandleFunc("/api/masters-ex/tournament/season/activate", h.HandleActivateTournamentSeason)
	http.HandleFunc("/api/masters-ex/tournament/season/active", h.HandleGetActiveTournamentSeason)
	http.HandleFunc("/api/masters-ex/tournament/season", h.HandleGetTournamentSeason)
	http.HandleFunc("/api/masters-ex/tournament/register", h.HandleRegisterTournamentPlayer)
	http.HandleFunc("/api/masters-ex/tournament/leaderboard", h.HandleGetTournamentLeaderboard)
	http.HandleFunc("/api/masters-ex/tournament/ranking", h.HandleGetTournamentPlayerRanking)
	http.HandleFunc("/api/masters-ex/tournament/rewards", h.HandleGetTournamentSeasonRewards)

	// Seasonal Events endpoints (register longer paths first!)
	http.HandleFunc("/api/masters-ex/events/create", h.HandleCreateSeasonalEvent)
	http.HandleFunc("/api/masters-ex/events/activate", h.HandleActivateSeasonalEvent)
	http.HandleFunc("/api/masters-ex/events/active", h.HandleGetActiveSeasonalEvents)
	http.HandleFunc("/api/masters-ex/events/join", h.HandleJoinSeasonalEvent)
	http.HandleFunc("/api/masters-ex/events/progress/update", h.HandleUpdateEventProgress)
	http.HandleFunc("/api/masters-ex/events/progress", h.HandleGetEventProgress)
	http.HandleFunc("/api/masters-ex/events/items", h.HandleGetEventItems)
	http.HandleFunc("/api/masters-ex/events/buy-item", h.HandleBuyEventItem)
	http.HandleFunc("/api/masters-ex/events/", h.HandleGetSeasonalEvent)
	// Also register without trailing slash for query parameters
	http.HandleFunc("/api/masters-ex/events", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			h.HandleGetSeasonalEvent(w, r)
		} else if r.URL.Path == "/api/masters-ex/events" && r.URL.RawQuery != "" {
			h.HandleGetSeasonalEvent(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	// Enhanced Battle System endpoints
	enhancedBattleEngine := service.NewEnhancedBattleEngine()
	effectsSystem := service.NewBattleEffectsSystem()
	strategySystem := service.NewBattleStrategySystem()
	battleEngineV2 := service.NewBattleEngineV2(nil, database)
	enhancedBattleHandler := handler.NewEnhancedBattleHandler(
		enhancedBattleEngine,
		effectsSystem,
		strategySystem,
		battleEngineV2,
	)
	handler.RegisterEnhancedBattleRoutes(http.DefaultServeMux, enhancedBattleHandler)

	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	log.Printf("Judge Server starting on %s", addr)
	log.Printf("Database: %s@%s:%d/%s", config.Database.User, config.Database.Host, config.Database.Port, config.Database.DBName)

	// Initialize rate limiter if enabled
	var rateLimiter *middleware.RateLimiter
	if config.RateLimit.Enabled {
		rateLimitConfig := middleware.RateLimitConfig{
			RequestsPerMinute: config.RateLimit.RequestsPerMinute,
			BurstSize:         config.RateLimit.BurstSize,
		}
		if rateLimitConfig.RequestsPerMinute <= 0 {
			rateLimitConfig.RequestsPerMinute = 1000 // default 1000 req/min
		}
		if rateLimitConfig.BurstSize <= 0 {
			rateLimitConfig.BurstSize = 50 // default burst size
		}
		rateLimiter = middleware.NewRateLimiter(rateLimitConfig)
		log.Printf("Rate limiting enabled: %d requests/min, burst size: %d",
			rateLimitConfig.RequestsPerMinute, rateLimitConfig.BurstSize)
	}

	// Create HTTP server with optional rate limiting middleware
	var server http.Handler = http.DefaultServeMux
	if rateLimiter != nil {
		server = rateLimiter.Middleware(http.DefaultServeMux)
	}

	go func() {
		if err := http.ListenAndServe(addr, server); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down...")
	if rateLimiter != nil {
		rateLimiter.Stop()
	}
	if batchSync != nil {
		batchSync.Stop()
	}
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func getDefaultConfig() *Config {
	return &Config{
		Server: struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		}{
			Host: "0.0.0.0",
			Port: 8080,
		},
		Database: struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DBName   string `yaml:"dbname"`
			SSLMode  string `yaml:"sslmode"`
		}{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "xiaodudu",
			DBName:   "agent_monster",
			SSLMode:  "disable",
		},
		Settlement: struct {
			DailyTime           string `yaml:"daily_time"`
			SyncToGitHub        bool   `yaml:"sync_to_github"`
			SyncIntervalSeconds int    `yaml:"sync_interval_seconds"`
		}{
			DailyTime:           "00:00",
			SyncToGitHub:        true,
			SyncIntervalSeconds: 300,
		},
	}
}
