-- Judge Server 战斗系统 Schema
-- 包含战斗、防守、队伍管理等功能的表定义
-- 这个文件应该在 judge-server/internal/db/database.go 的 InitSchema() 中执行

-- =============================================================================
-- 用户队伍管理表 (User Team Management)
-- =============================================================================

-- 用户的战斗队伍（一个用户最多可以有多个队伍）
CREATE TABLE IF NOT EXISTS user_teams (
    id SERIAL PRIMARY KEY,
    github_id INTEGER NOT NULL,
    team_id VARCHAR(255) NOT NULL UNIQUE,
    team_name VARCHAR(255) NOT NULL DEFAULT 'Battle Team',
    description TEXT,
    is_defense_team BOOLEAN DEFAULT FALSE,  -- 是否是防守队伍（基地防御）
    max_members INTEGER DEFAULT 3,  -- 最多成员数
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
    UNIQUE(github_id, team_name)
);

-- 用户队伍成员（关联宠物到队伍）
CREATE TABLE IF NOT EXISTS user_team_members (
    id SERIAL PRIMARY KEY,
    team_id VARCHAR(255) NOT NULL,
    pet_id VARCHAR(255) NOT NULL,
    slot_position INTEGER NOT NULL,  -- 队伍位置 1-3
    joined_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (team_id) REFERENCES user_teams(team_id) ON DELETE CASCADE,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    UNIQUE(team_id, slot_position),
    UNIQUE(team_id, pet_id)
);

-- 用户拥有的宠物（追踪用户最多10个宠物的限制）
CREATE TABLE IF NOT EXISTS user_owned_pokemon (
    id SERIAL PRIMARY KEY,
    github_id INTEGER NOT NULL,
    pet_id VARCHAR(255) NOT NULL UNIQUE,
    is_captured BOOLEAN DEFAULT FALSE,  -- TRUE=捕获的野生宠物, FALSE=孵化的宠物
    status VARCHAR(50) DEFAULT 'active',  -- active, fainted, training
    current_hp INTEGER NOT NULL,
    max_hp INTEGER NOT NULL,
    level INTEGER DEFAULT 1,
    experience INTEGER DEFAULT 0,
    captured_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE
);

-- =============================================================================
-- 防守基地系统表 (Defense Base System)
-- =============================================================================

-- 用户的基地（Fork后才能创建）
CREATE TABLE IF NOT EXISTS user_bases (
    id SERIAL PRIMARY KEY,
    github_id INTEGER NOT NULL UNIQUE,
    base_id VARCHAR(255) NOT NULL UNIQUE,
    repository_url TEXT NOT NULL,  -- Fork来源的repo URL
    defense_team_id VARCHAR(255),  -- 防守队伍ID（最强3只）
    level INTEGER DEFAULT 1,
    prestige INTEGER DEFAULT 0,  -- 威望值，根据防守成功次数增加
    defense_wins INTEGER DEFAULT 0,  -- 防守成功次数
    defense_losses INTEGER DEFAULT 0,  -- 防守失败次数
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
    FOREIGN KEY (defense_team_id) REFERENCES user_teams(team_id) ON DELETE SET NULL
);

-- 基地防守记录（记录所有防守战斗）
CREATE TABLE IF NOT EXISTS base_defense_records (
    id SERIAL PRIMARY KEY,
    base_id VARCHAR(255) NOT NULL,
    attacker_github_id INTEGER NOT NULL,
    battle_id VARCHAR(255) UNIQUE NOT NULL,
    result VARCHAR(50) NOT NULL,  -- 'win' or 'lose'
    attacker_level INTEGER NOT NULL,
    defender_level INTEGER NOT NULL,
    coins_gained DECIMAL(12, 2) DEFAULT 0,
    timestamp TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (base_id) REFERENCES user_bases(base_id) ON DELETE CASCADE,
    FOREIGN KEY (attacker_github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE
);

-- =============================================================================
-- 战斗系统表 (Battle System)
-- =============================================================================

-- 战斗记录（记录每一场战斗）
CREATE TABLE IF NOT EXISTS battles (
    id SERIAL PRIMARY KEY,
    battle_id VARCHAR(255) NOT NULL UNIQUE,
    attacker_github_id INTEGER NOT NULL,
    defender_github_id INTEGER NOT NULL,
    attacker_team_id VARCHAR(255),  -- 攻击者队伍ID
    defender_team_id VARCHAR(255),  -- 防守者队伍ID（可能是防守队伍）
    battle_type VARCHAR(50) NOT NULL,  -- 'pvp' or 'base_defense'
    winner_github_id INTEGER,  -- 赢家的ID
    status VARCHAR(50) DEFAULT 'ongoing',  -- 'ongoing', 'completed', 'abandoned'
    reward_coins DECIMAL(12, 2) DEFAULT 0,  -- 赢家获得的硬币
    started_at TIMESTAMP DEFAULT NOW(),
    ended_at TIMESTAMP,
    FOREIGN KEY (attacker_github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
    FOREIGN KEY (defender_github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE
);

-- 战斗回合记录（记录战斗的每一回合）
CREATE TABLE IF NOT EXISTS battle_rounds (
    id SERIAL PRIMARY KEY,
    battle_id VARCHAR(255) NOT NULL,
    round_number INTEGER NOT NULL,
    attacker_pet_id VARCHAR(255),
    defender_pet_id VARCHAR(255),
    attacker_damage INTEGER DEFAULT 0,
    defender_damage INTEGER DEFAULT 0,
    attacker_pet_current_hp INTEGER,  -- 本回合后攻击者宠物HP
    defender_pet_current_hp INTEGER,  -- 本回合后防守者宠物HP
    round_winner VARCHAR(50),  -- 'attacker', 'defender', or 'draw'
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (battle_id) REFERENCES battles(battle_id) ON DELETE CASCADE
);

-- 宠物战斗统计（追踪宠物的战斗历史）
CREATE TABLE IF NOT EXISTS pet_battle_stats (
    id SERIAL PRIMARY KEY,
    pet_id VARCHAR(255) NOT NULL UNIQUE,
    total_battles INTEGER DEFAULT 0,
    wins INTEGER DEFAULT 0,
    losses INTEGER DEFAULT 0,
    draws INTEGER DEFAULT 0,
    total_damage_dealt INTEGER DEFAULT 0,
    total_damage_taken INTEGER DEFAULT 0,
    last_battle_at TIMESTAMP,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE
);

-- =============================================================================
-- 精灵币交易记录（用于战斗奖励）
-- =============================================================================

-- 战斗奖励交易
CREATE TABLE IF NOT EXISTS battle_rewards (
    id SERIAL PRIMARY KEY,
    battle_id VARCHAR(255) NOT NULL,
    winner_github_id INTEGER NOT NULL,
    loser_github_id INTEGER NOT NULL,
    coins_transferred DECIMAL(12, 2) NOT NULL,
    percentage NUMERIC(5, 2) NOT NULL,  -- 1-5% 
    transaction_id VARCHAR(255) UNIQUE,
    completed_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (battle_id) REFERENCES battles(battle_id) ON DELETE CASCADE,
    FOREIGN KEY (winner_github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
    FOREIGN KEY (loser_github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE
);

-- =============================================================================
-- 野生精灵系统表 (Wild Pokemon System)
-- =============================================================================

-- 地图上的野生宠物（NPC）
CREATE TABLE IF NOT EXISTS wild_pokemon (
    id SERIAL PRIMARY KEY,
    wild_id VARCHAR(255) NOT NULL UNIQUE,
    location_id VARCHAR(255) NOT NULL,  -- 所在位置ID
    pokemon_species_id VARCHAR(255) NOT NULL,
    level INTEGER NOT NULL,
    status VARCHAR(50) DEFAULT 'active',  -- 'active', 'captured', 'defeated'
    current_hp INTEGER NOT NULL,
    max_hp INTEGER NOT NULL,
    capture_difficulty INTEGER DEFAULT 50,  -- 1-100, 难度越高越难捕获
    caught_by_github_id INTEGER,  -- 捕获的玩家ID
    defeated_by_github_id INTEGER,  -- 击败的玩家ID
    created_at TIMESTAMP DEFAULT NOW(),
    captured_at TIMESTAMP,
    defeated_at TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (pokemon_species_id) REFERENCES pokemon_species(pokemon_id) ON DELETE CASCADE
);

-- =============================================================================
-- 精灵捕获历史 (Pokemon Capture History)
-- =============================================================================

-- 用户捕获精灵的历史记录
CREATE TABLE IF NOT EXISTS capture_history (
    id SERIAL PRIMARY KEY,
    github_id INTEGER NOT NULL,
    pet_id VARCHAR(255) NOT NULL,
    wild_pokemon_id VARCHAR(255),  -- 来自哪只野生宠物
    capture_type VARCHAR(50) NOT NULL,  -- 'wild', 'hatched', 'traded'
    success BOOLEAN DEFAULT TRUE,
    attempt_count INTEGER DEFAULT 1,  -- 尝试次数
    captured_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (github_id) REFERENCES user_accounts(github_id) ON DELETE CASCADE,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE
);

-- =============================================================================
-- 索引优化 (Indexes for Performance)
-- =============================================================================

CREATE INDEX IF NOT EXISTS idx_user_teams_github_id ON user_teams(github_id);
CREATE INDEX IF NOT EXISTS idx_user_teams_team_id ON user_teams(team_id);
CREATE INDEX IF NOT EXISTS idx_user_teams_is_defense ON user_teams(is_defense_team);

CREATE INDEX IF NOT EXISTS idx_user_team_members_team_id ON user_team_members(team_id);
CREATE INDEX IF NOT EXISTS idx_user_team_members_pet_id ON user_team_members(pet_id);

CREATE INDEX IF NOT EXISTS idx_user_owned_pokemon_github_id ON user_owned_pokemon(github_id);
CREATE INDEX IF NOT EXISTS idx_user_owned_pokemon_pet_id ON user_owned_pokemon(pet_id);
CREATE INDEX IF NOT EXISTS idx_user_owned_pokemon_status ON user_owned_pokemon(status);
CREATE INDEX IF NOT EXISTS idx_user_owned_pokemon_level ON user_owned_pokemon(level);

CREATE INDEX IF NOT EXISTS idx_user_bases_github_id ON user_bases(github_id);
CREATE INDEX IF NOT EXISTS idx_user_bases_base_id ON user_bases(base_id);

CREATE INDEX IF NOT EXISTS idx_base_defense_records_base_id ON base_defense_records(base_id);
CREATE INDEX IF NOT EXISTS idx_base_defense_records_attacker ON base_defense_records(attacker_github_id);
CREATE INDEX IF NOT EXISTS idx_base_defense_records_timestamp ON base_defense_records(timestamp DESC);

CREATE INDEX IF NOT EXISTS idx_battles_battle_id ON battles(battle_id);
CREATE INDEX IF NOT EXISTS idx_battles_attacker_id ON battles(attacker_github_id);
CREATE INDEX IF NOT EXISTS idx_battles_defender_id ON battles(defender_github_id);
CREATE INDEX IF NOT EXISTS idx_battles_status ON battles(status);
CREATE INDEX IF NOT EXISTS idx_battles_started_at ON battles(started_at DESC);

CREATE INDEX IF NOT EXISTS idx_battle_rounds_battle_id ON battle_rounds(battle_id);
CREATE INDEX IF NOT EXISTS idx_battle_rounds_round_number ON battle_rounds(round_number);

CREATE INDEX IF NOT EXISTS idx_pet_battle_stats_pet_id ON pet_battle_stats(pet_id);
CREATE INDEX IF NOT EXISTS idx_pet_battle_stats_wins ON pet_battle_stats(wins DESC);

CREATE INDEX IF NOT EXISTS idx_battle_rewards_battle_id ON battle_rewards(battle_id);
CREATE INDEX IF NOT EXISTS idx_battle_rewards_winner_id ON battle_rewards(winner_github_id);
CREATE INDEX IF NOT EXISTS idx_battle_rewards_loser_id ON battle_rewards(loser_github_id);
CREATE INDEX IF NOT EXISTS idx_battle_rewards_completed_at ON battle_rewards(completed_at DESC);

CREATE INDEX IF NOT EXISTS idx_wild_pokemon_wild_id ON wild_pokemon(wild_id);
CREATE INDEX IF NOT EXISTS idx_wild_pokemon_location_id ON wild_pokemon(location_id);
CREATE INDEX IF NOT EXISTS idx_wild_pokemon_status ON wild_pokemon(status);
CREATE INDEX IF NOT EXISTS idx_wild_pokemon_level ON wild_pokemon(level);

CREATE INDEX IF NOT EXISTS idx_capture_history_github_id ON capture_history(github_id);
CREATE INDEX IF NOT EXISTS idx_capture_history_pet_id ON capture_history(pet_id);
CREATE INDEX IF NOT EXISTS idx_capture_history_captured_at ON capture_history(captured_at DESC);
