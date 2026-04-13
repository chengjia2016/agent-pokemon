package db

import (
	"database/sql"
	"fmt"
	"judge-server/internal/model"
	"time"

	"github.com/lib/pq"
)

type Database struct {
	conn *sql.DB
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewDatabase(cfg Config) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{conn: db}, nil
}

func (d *Database) InitSchema() error {
	schemas := []string{
		`CREATE TABLE IF NOT EXISTS pets (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			owner TEXT NOT NULL,
			species TEXT,
			generation INTEGER DEFAULT 1,
			evolution_stage INTEGER DEFAULT 1,
			base_stats JSONB,
			ivs JSONB,
			evs JSONB,
			exp INTEGER DEFAULT 0,
			level INTEGER DEFAULT 1,
			genes JSONB,
			battle_history JSONB,
			signature JSONB,
			is_valid BOOLEAN DEFAULT false,
			validation_msg TEXT,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS battles (
			id TEXT PRIMARY KEY,
			attacker_id TEXT NOT NULL,
			defender_id TEXT NOT NULL,
			attacker_name TEXT,
			defender_name TEXT,
			winner TEXT,
			turns INTEGER,
			attack_stack TEXT[],
			defense_stack TEXT[],
			battle_log TEXT[],
			start_time TIMESTAMP,
			end_time TIMESTAMP,
			is_valid BOOLEAN DEFAULT false,
			validation_msg TEXT,
			created_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS leaderboard (
			id TEXT PRIMARY KEY,
			player TEXT NOT NULL,
			pet_name TEXT,
			level INTEGER DEFAULT 1,
			total_battles INTEGER DEFAULT 0,
			wins INTEGER DEFAULT 0,
			losses INTEGER DEFAULT 0,
			rating INTEGER DEFAULT 1000,
			rank INTEGER,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS food_transactions (
			id TEXT PRIMARY KEY,
			player TEXT NOT NULL,
			pet_id TEXT NOT NULL,
			food_type TEXT NOT NULL,
			effect TEXT,
			amount INTEGER DEFAULT 1,
			timestamp TIMESTAMP DEFAULT NOW(),
			is_valid BOOLEAN DEFAULT false
		)`,
		`CREATE TABLE IF NOT EXISTS growth_records (
			id TEXT PRIMARY KEY,
			pet_id TEXT NOT NULL,
			old_level INTEGER,
			new_level INTEGER,
			old_exp INTEGER,
			new_exp INTEGER,
			source TEXT,
			timestamp TIMESTAMP DEFAULT NOW(),
			is_valid BOOLEAN DEFAULT false
		)`,
		`CREATE TABLE IF NOT EXISTS daily_leaderboards (
			date DATE PRIMARY KEY,
			entries JSONB,
			generated_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE INDEX IF NOT EXISTS idx_pets_owner ON pets(owner)`,
		`CREATE INDEX IF NOT EXISTS idx_battles_attacker ON battles(attacker_id)`,
		`CREATE INDEX IF NOT EXISTS idx_battles_defender ON battles(defender_id)`,
		`CREATE INDEX IF NOT EXISTS idx_leaderboard_rating ON leaderboard(rating DESC)`,
		`CREATE INDEX IF NOT EXISTS idx_food_transactions_pet_id_timestamp ON food_transactions(pet_id, timestamp)`,
	}

	for _, schema := range schemas {
		if _, err := d.conn.Exec(schema); err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) SavePet(pet *model.Pet) error {
	query := `
		INSERT INTO pets (id, name, owner, species, generation, evolution_stage, 
			base_stats, ivs, evs, exp, level, genes, battle_history, signature, 
			is_valid, validation_msg, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, NOW())
		ON CONFLICT (id) DO UPDATE SET
			name = $2, owner = $3, species = $4, generation = $5, evolution_stage = $6,
			base_stats = $7, ivs = $8, evs = $9, exp = $10, level = $11, genes = $12,
			battle_history = $13, signature = $14, is_valid = $15, validation_msg = $16,
			updated_at = NOW()
	`
	_, err := d.conn.Exec(query,
		pet.ID, pet.Name, pet.Owner, pet.Species, pet.Generation, pet.EvolutionStage,
		pet.BaseStats, pet.IVs, pet.EVs, pet.EXP, pet.Level, pet.Genes,
		pet.BattleHistory, pet.Signature, pet.IsValid, pet.ValidationMsg)
	return err
}

func (d *Database) GetPet(id string) (*model.Pet, error) {
	query := `SELECT id, name, owner, species, generation, evolution_stage, 
		base_stats, ivs, evs, exp, level, genes, battle_history, signature,
		is_valid, validation_msg, created_at, updated_at 
		FROM pets WHERE id = $1`

	pet := &model.Pet{}
	err := d.conn.QueryRow(query, id).Scan(
		&pet.ID, &pet.Name, &pet.Owner, &pet.Species, &pet.Generation, &pet.EvolutionStage,
		&pet.BaseStats, &pet.IVs, &pet.EVs, &pet.EXP, &pet.Level, &pet.Genes,
		&pet.BattleHistory, &pet.Signature, &pet.IsValid, &pet.ValidationMsg,
		&pet.CreatedAt, &pet.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return pet, nil
}

func (d *Database) SaveBattle(battle *model.Battle) error {
	query := `
		INSERT INTO battles (id, attacker_id, defender_id, attacker_name, defender_name,
			winner, turns, attack_stack, defense_stack, battle_log, start_time, end_time,
			is_valid, validation_msg)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`
	_, err := d.conn.Exec(query,
		battle.ID, battle.AttackerID, battle.DefenderID, battle.AttackerName, battle.DefenderName,
		battle.Winner, battle.Turns, pq.Array(battle.AttackStack), pq.Array(battle.DefenseStack), pq.Array(battle.BattleLog),
		battle.StartTime, battle.EndTime, battle.IsValid, battle.ValidationMsg)
	return err
}

func (d *Database) GetLeaderboard(limit int) ([]model.LeaderboardEntry, error) {
	query := `SELECT id, player, pet_name, level, total_battles, wins, losses, rating, rank, created_at, updated_at
		FROM leaderboard ORDER BY rating DESC LIMIT $1`

	rows, err := d.conn.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []model.LeaderboardEntry
	for rows.Next() {
		var e model.LeaderboardEntry
		err := rows.Scan(&e.ID, &e.Player, &e.PetName, &e.Level, &e.TotalBattles,
			&e.Wins, &e.Losses, &e.Rating, &e.Rank, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

func (d *Database) SaveFoodTransaction(tx *model.FoodTransaction) error {
	query := `
		INSERT INTO food_transactions (id, player, pet_id, food_type, effect, amount, timestamp, is_valid)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id) DO NOTHING
	`
	_, err := d.conn.Exec(query, tx.ID, tx.PlayerID, tx.PetID, tx.FoodType, "", tx.Quantity, tx.Timestamp, tx.IsValid)
	return err
}

func (d *Database) GetDailyEatCount(petID string, timestamp time.Time) (int, error) {
	// Calculate today's date range
	today := timestamp.Truncate(24 * time.Hour)
	tomorrow := today.AddDate(0, 0, 1)

	query := `
		SELECT COUNT(*) FROM food_transactions
		WHERE pet_id = $1 AND timestamp >= $2 AND timestamp < $3 AND is_valid = true
	`
	var count int
	err := d.conn.QueryRow(query, petID, today, tomorrow).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count, nil
}

func (d *Database) GetRepoVisitCount(farmOwner string, farmRepo string, timestamp time.Time) (int, error) {
	// Calculate today's date range
	today := timestamp.Truncate(24 * time.Hour)
	tomorrow := today.AddDate(0, 0, 1)

	// Note: farm_owner and farm_repo columns don't exist in the current schema
	// This function returns 0 for now to maintain compatibility
	query := `
		SELECT COUNT(*) FROM food_transactions
		WHERE timestamp >= $1 AND timestamp < $2 AND is_valid = true
	`
	var count int
	err := d.conn.QueryRow(query, today, tomorrow).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count, nil
}

func (d *Database) SaveGrowthRecord(gr *model.GrowthRecord) error {
	query := `
		INSERT INTO growth_records (id, pet_id, old_level, new_level, old_exp, new_exp, source, timestamp, is_valid)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := d.conn.Exec(query, gr.ID, gr.PetID, gr.OldLevel, gr.NewLevel, gr.OldEXP, gr.NewEXP, gr.Source, gr.Timestamp, gr.IsValid)
	return err
}

func (d *Database) SaveDailyLeaderboard(dl *model.DailyLeaderboard) error {
	query := `
		INSERT INTO daily_leaderboards (date, entries, generated_at)
		VALUES ($1, $2, $3)
		ON CONFLICT (date) DO UPDATE SET entries = $2, generated_at = $3
	`
	_, err := d.conn.Exec(query, dl.Date, dl.Entries, dl.Generated)
	return err
}

func (d *Database) Close() error {
	return d.conn.Close()
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.conn.Exec(query, args...)
}

// QueryRow 执行单行查询
func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.conn.QueryRow(query, args...)
}

// Query 执行多行查询
func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.conn.Query(query, args...)
}

// InitBattleSystemSchema 初始化战斗系统的Schema
func (d *Database) InitBattleSystemSchema() error {
	schemas := []string{
		// 用户账户表
		`CREATE TABLE IF NOT EXISTS user_accounts (
			id SERIAL PRIMARY KEY,
			github_id INTEGER UNIQUE NOT NULL,
			github_login VARCHAR(255) NOT NULL,
			email VARCHAR(255),
			avatar_url TEXT,
			balance DECIMAL(12, 2) DEFAULT 0.0,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,
		// 用户精灵表
		`CREATE TABLE IF NOT EXISTS user_pokemons (
			id SERIAL PRIMARY KEY,
			github_id INTEGER NOT NULL,
			pet_id TEXT NOT NULL UNIQUE,
			pet_name TEXT NOT NULL,
			level INTEGER DEFAULT 1,
			species TEXT,
			created_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (github_id) REFERENCES user_accounts(github_id),
			FOREIGN KEY (pet_id) REFERENCES pets(id)
		)`,
		// 用户队伍表
		`CREATE TABLE IF NOT EXISTS user_teams (
			id SERIAL PRIMARY KEY,
			github_id INTEGER NOT NULL,
			team_id VARCHAR(255) NOT NULL UNIQUE,
			team_name VARCHAR(255) NOT NULL DEFAULT 'Battle Team',
			description TEXT,
			is_defense_team BOOLEAN DEFAULT FALSE,
			max_members INTEGER DEFAULT 3,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
			UNIQUE(github_id, team_name)
		)`,
		// 用户队伍成员表
		`CREATE TABLE IF NOT EXISTS user_team_members (
			id SERIAL PRIMARY KEY,
			team_id VARCHAR(255) NOT NULL,
			pet_id TEXT NOT NULL,
			slot_position INTEGER NOT NULL,
			joined_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (team_id) REFERENCES user_teams(team_id) ON DELETE CASCADE,
			FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
			UNIQUE(team_id, slot_position),
			UNIQUE(team_id, pet_id)
		)`,
		// 用户拥有的宠物表
		`CREATE TABLE IF NOT EXISTS user_owned_pokemon (
			id SERIAL PRIMARY KEY,
			github_id INTEGER NOT NULL,
			pet_id TEXT NOT NULL UNIQUE,
			is_captured BOOLEAN DEFAULT FALSE,
			status VARCHAR(50) DEFAULT 'active',
			current_hp INTEGER NOT NULL,
			max_hp INTEGER NOT NULL,
			level INTEGER DEFAULT 1,
			experience INTEGER DEFAULT 0,
			captured_at TIMESTAMP,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
			FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE
		)`,
		// 用户基地表
		`CREATE TABLE IF NOT EXISTS user_bases (
			id SERIAL PRIMARY KEY,
			github_id INTEGER NOT NULL UNIQUE,
			base_id VARCHAR(255) NOT NULL UNIQUE,
			repository_url TEXT NOT NULL,
			defense_team_id VARCHAR(255),
			level INTEGER DEFAULT 1,
			prestige INTEGER DEFAULT 0,
			defense_wins INTEGER DEFAULT 0,
			defense_losses INTEGER DEFAULT 0,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
			FOREIGN KEY (defense_team_id) REFERENCES user_teams(team_id) ON DELETE SET NULL
		)`,
		// 基地防守记录表
		`CREATE TABLE IF NOT EXISTS base_defense_records (
			id SERIAL PRIMARY KEY,
			base_id VARCHAR(255) NOT NULL,
			attacker_github_id INTEGER NOT NULL,
			battle_id VARCHAR(255) UNIQUE NOT NULL,
			result VARCHAR(50) NOT NULL,
			attacker_level INTEGER NOT NULL,
			defender_level INTEGER NOT NULL,
			coins_gained DECIMAL(12, 2) DEFAULT 0,
			timestamp TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (base_id) REFERENCES user_bases(base_id) ON DELETE CASCADE,
			FOREIGN KEY (attacker_github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE
		)`,
		// 战斗系统扩展表
		`CREATE TABLE IF NOT EXISTS battles_v2 (
			id SERIAL PRIMARY KEY,
			battle_id VARCHAR(255) NOT NULL UNIQUE,
			attacker_github_id INTEGER NOT NULL,
			defender_github_id INTEGER NOT NULL,
			attacker_team_id VARCHAR(255),
			defender_team_id VARCHAR(255),
			battle_type VARCHAR(50) NOT NULL,
			winner_github_id INTEGER,
			status VARCHAR(50) DEFAULT 'ongoing',
			reward_coins DECIMAL(12, 2) DEFAULT 0,
			started_at TIMESTAMP DEFAULT NOW(),
			ended_at TIMESTAMP,
			FOREIGN KEY (attacker_github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
			FOREIGN KEY (defender_github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE
		)`,
		// 战斗回合表
		`CREATE TABLE IF NOT EXISTS battle_rounds (
			id SERIAL PRIMARY KEY,
			battle_id VARCHAR(255) NOT NULL,
			round_number INTEGER NOT NULL,
			attacker_pet_id TEXT,
			defender_pet_id TEXT,
			attacker_damage INTEGER DEFAULT 0,
			defender_damage INTEGER DEFAULT 0,
			attacker_pet_current_hp INTEGER,
			defender_pet_current_hp INTEGER,
			round_winner VARCHAR(50),
			created_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (battle_id) REFERENCES battles_v2(battle_id) ON DELETE CASCADE
		)`,
		// 宠物战斗统计表
		`CREATE TABLE IF NOT EXISTS pet_battle_stats (
			id SERIAL PRIMARY KEY,
			pet_id TEXT NOT NULL UNIQUE,
			total_battles INTEGER DEFAULT 0,
			wins INTEGER DEFAULT 0,
			losses INTEGER DEFAULT 0,
			draws INTEGER DEFAULT 0,
			total_damage_dealt INTEGER DEFAULT 0,
			total_damage_taken INTEGER DEFAULT 0,
			last_battle_at TIMESTAMP,
			FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE
		)`,
		// 战斗奖励表
		`CREATE TABLE IF NOT EXISTS battle_rewards (
			id SERIAL PRIMARY KEY,
			battle_id VARCHAR(255) NOT NULL,
			winner_github_id INTEGER NOT NULL,
			loser_github_id INTEGER NOT NULL,
			coins_transferred DECIMAL(12, 2) NOT NULL,
			percentage NUMERIC(5, 2) NOT NULL,
			transaction_id VARCHAR(255) UNIQUE,
			completed_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (battle_id) REFERENCES battles_v2(battle_id) ON DELETE CASCADE,
			FOREIGN KEY (winner_github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
			FOREIGN KEY (loser_github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE
		)`,
		// 野生精灵表
		`CREATE TABLE IF NOT EXISTS wild_pokemon (
			id SERIAL PRIMARY KEY,
			wild_id VARCHAR(255) NOT NULL UNIQUE,
			location_id VARCHAR(255) NOT NULL,
			pokemon_species_id VARCHAR(255) NOT NULL,
			level INTEGER NOT NULL,
			status VARCHAR(50) DEFAULT 'active',
			current_hp INTEGER NOT NULL,
			max_hp INTEGER NOT NULL,
			capture_difficulty INTEGER DEFAULT 50,
			caught_by_github_id INTEGER,
			defeated_by_github_id INTEGER,
			created_at TIMESTAMP DEFAULT NOW(),
			captured_at TIMESTAMP,
			defeated_at TIMESTAMP,
			updated_at TIMESTAMP DEFAULT NOW()
		)`,
		// 捕获历史表
		`CREATE TABLE IF NOT EXISTS capture_history (
			id SERIAL PRIMARY KEY,
			github_id INTEGER NOT NULL,
			pet_id TEXT NOT NULL,
			wild_pokemon_id VARCHAR(255),
			capture_type VARCHAR(50) NOT NULL,
			success BOOLEAN DEFAULT TRUE,
			attempt_count INTEGER DEFAULT 1,
			captured_at TIMESTAMP DEFAULT NOW(),
			FOREIGN KEY (github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
			FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE
		)`,
		// 创建索引
		`CREATE INDEX IF NOT EXISTS idx_user_accounts_github_id ON user_accounts(github_id)`,
		`CREATE INDEX IF NOT EXISTS idx_user_teams_github_id ON user_teams(github_id)`,
		`CREATE INDEX IF NOT EXISTS idx_user_team_members_team_id ON user_team_members(team_id)`,
		`CREATE INDEX IF NOT EXISTS idx_user_owned_pokemon_github_id ON user_owned_pokemon(github_id)`,
		`CREATE INDEX IF NOT EXISTS idx_user_owned_pokemon_status ON user_owned_pokemon(status)`,
		`CREATE INDEX IF NOT EXISTS idx_user_bases_github_id ON user_bases(github_id)`,
		`CREATE INDEX IF NOT EXISTS idx_base_defense_records_base_id ON base_defense_records(base_id)`,
		`CREATE INDEX IF NOT EXISTS idx_battles_v2_attacker_id ON battles_v2(attacker_github_id)`,
		`CREATE INDEX IF NOT EXISTS idx_battles_v2_defender_id ON battles_v2(defender_github_id)`,
		`CREATE INDEX IF NOT EXISTS idx_wild_pokemon_location_id ON wild_pokemon(location_id)`,
		`CREATE INDEX IF NOT EXISTS idx_wild_pokemon_status ON wild_pokemon(status)`,
		`CREATE INDEX IF NOT EXISTS idx_capture_history_github_id ON capture_history(github_id)`,
	}

	for _, schema := range schemas {
		if _, err := d.conn.Exec(schema); err != nil {
			// 如果是"already exists"错误，忽略它
			if err.Error() != "pq: relation \"user_accounts\" already exists" {
				// 输出日志但继续
				fmt.Printf("Schema creation notice: %v (continuing)\n", err)
			}
		}
	}
	return nil
}
