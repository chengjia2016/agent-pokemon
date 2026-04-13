package main

import (
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/github"
	"judge-server/internal/handler"
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
	http.HandleFunc("/api/dungeons", h.HandleWorldDungeons)
	http.HandleFunc("/api/gyms", h.HandleWorldGyms)
	http.HandleFunc("/api/map/zones", h.HandleWorldMapZones)
	http.HandleFunc("/api/user/levels/progress", h.HandleWorldUserLevelProgress)
	http.HandleFunc("/api/levels", h.HandleWorldLevels)

	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	log.Printf("Judge Server starting on %s", addr)
	log.Printf("Database: %s@%s:%d/%s", config.Database.User, config.Database.Host, config.Database.Port, config.Database.DBName)

	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down...")
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
