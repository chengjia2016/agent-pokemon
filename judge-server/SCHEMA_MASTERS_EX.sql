-- Agent Monster v2.1.0 - Masters EX 系统数据库 Schema
-- 包含配对系统、3v3 对战、赛事系统、季节活动、拍档招式等
-- 执行时间：2026-04-14 15:30 UTC

-- =============================================================================
-- DROP OLD TABLES (清理旧表)
-- =============================================================================

-- Drop all Masters EX related tables (if they exist from old schema)
DROP TABLE IF EXISTS player_seasonal_items CASCADE;
DROP TABLE IF EXISTS seasonal_items CASCADE;
DROP TABLE IF EXISTS event_progress CASCADE;
DROP TABLE IF EXISTS seasonal_events CASCADE;
DROP TABLE IF EXISTS tournament_rewards CASCADE;
DROP TABLE IF EXISTS tournament_rankings CASCADE;
DROP TABLE IF EXISTS tournament_seasons CASCADE;
DROP TABLE IF EXISTS battle_statistics CASCADE;
DROP TABLE IF EXISTS battle_actions CASCADE;
DROP TABLE IF EXISTS battle_sessions CASCADE;
DROP TABLE IF EXISTS sync_pair_equipment CASCADE;
DROP TABLE IF EXISTS sync_moves CASCADE;
DROP TABLE IF EXISTS sync_pair_moves CASCADE;
DROP TABLE IF EXISTS sync_pairs CASCADE;

-- =============================================================================
-- 配对系统表 (Sync Pair System)
-- 基于 Pokémon Masters EX 的 Sync Pair 概念
-- =============================================================================

-- 1. 配对信息表 (Sync Pairs)
CREATE TABLE IF NOT EXISTS sync_pairs (
    id SERIAL PRIMARY KEY,
    sync_pair_id VARCHAR(255) NOT NULL UNIQUE,  -- 格式: trainer_id_pokemon_id
    trainer_id VARCHAR(255) NOT NULL,           -- 训练家 ID
    pokemon_id VARCHAR(255) NOT NULL,           -- 宝可梦 ID
    trainer_name VARCHAR(100) NOT NULL,         -- 训练家名称
    pokemon_name VARCHAR(100) NOT NULL,         -- 宝可梦名称
    pokemon_type VARCHAR(20) NOT NULL,          -- 宝可梦属性
    sync_pair_name VARCHAR(100),                -- 配对组合的昵称
    rarity_stars INTEGER DEFAULT 5,             -- 稀有度 (3-5 星)
    level INTEGER DEFAULT 1,                    -- 当前等级
    experience BIGINT DEFAULT 0,                -- 经验值
    max_level INTEGER DEFAULT 130,              -- 最大等级
    hp INTEGER DEFAULT 100,                     -- 生命值
    attack INTEGER DEFAULT 50,                  -- 攻击
    defense INTEGER DEFAULT 50,                 -- 防御
    sp_attack INTEGER DEFAULT 50,               -- 特攻
    sp_defense INTEGER DEFAULT 50,              -- 特防
    speed INTEGER DEFAULT 50,                   -- 速度
    move_1_id VARCHAR(255),                     -- 招式 1
    move_2_id VARCHAR(255),                     -- 招式 2
    move_3_id VARCHAR(255),                     -- 招式 3
    sync_move_id VARCHAR(255),                  -- 拍档招式 ID
    sync_move_ready_percentage INTEGER DEFAULT 0,  -- 拍档招式准备百分比
    potential_unlocked INTEGER DEFAULT 0,       -- 解锁潜力等级 (0-20)
    is_active BOOLEAN DEFAULT TRUE,             -- 是否当前激活
    obtained_at TIMESTAMP DEFAULT NOW(),        -- 获得时间
    last_used_at TIMESTAMP,                     -- 最后使用时间
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 索引
CREATE INDEX IF NOT EXISTS idx_sync_pairs_trainer ON sync_pairs(trainer_id);
CREATE INDEX IF NOT EXISTS idx_sync_pairs_pokemon ON sync_pairs(pokemon_id);

-- 2. 配对招式表 (Sync Pair Moves)
CREATE TABLE IF NOT EXISTS sync_pair_moves (
    id SERIAL PRIMARY KEY,
    move_id VARCHAR(255) NOT NULL UNIQUE,       -- 招式 ID
    move_name VARCHAR(100) NOT NULL,            -- 招式名称
    move_type VARCHAR(20) NOT NULL,             -- 招式属性
    power INTEGER,                              -- 威力
    accuracy INTEGER,                           -- 命中率
    pp INTEGER,                                 -- 使用次数
    category VARCHAR(20),                       -- 物理/特殊
    effect_description TEXT,                    -- 效果说明 (支持多语言)
    priority INTEGER DEFAULT 0,                 -- 优先度
    created_at TIMESTAMP DEFAULT NOW()
);

-- 3. 拍档招式表 (Sync Moves)
CREATE TABLE IF NOT EXISTS sync_moves (
    id SERIAL PRIMARY KEY,
    sync_move_id VARCHAR(255) NOT NULL UNIQUE,
    move_name VARCHAR(100) NOT NULL,
    move_name_zh VARCHAR(100),                  -- 中文名称
    trainer_id VARCHAR(255) NOT NULL,
    pokemon_id VARCHAR(255) NOT NULL,
    description TEXT,
    description_zh TEXT,                        -- 中文说明
    power INTEGER DEFAULT 300,                  -- 基础威力
    effect_description TEXT,                    -- 特殊效果
    effect_description_zh TEXT,
    required_sync_level INTEGER DEFAULT 1,      -- 需要的拍档等级
    cooldown_turns INTEGER DEFAULT 2,           -- 冷却回合数
    created_at TIMESTAMP DEFAULT NOW()
);

-- 4. 配对装备表 (Sync Pair Equipment/Gear)
CREATE TABLE IF NOT EXISTS sync_pair_equipment (
    id SERIAL PRIMARY KEY,
    equipment_id VARCHAR(255) NOT NULL UNIQUE,
    sync_pair_id VARCHAR(255) NOT NULL,
    equipment_name VARCHAR(100) NOT NULL,
    equipment_type VARCHAR(50),                 -- 鞋子、帽子、套装等
    level INTEGER DEFAULT 1,                    -- 装备等级
    hp_bonus INTEGER DEFAULT 0,
    attack_bonus INTEGER DEFAULT 0,
    defense_bonus INTEGER DEFAULT 0,
    sp_attack_bonus INTEGER DEFAULT 0,
    sp_defense_bonus INTEGER DEFAULT 0,
    speed_bonus INTEGER DEFAULT 0,
    obtained_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (sync_pair_id) REFERENCES sync_pairs(sync_pair_id)
);

-- =============================================================================
-- 3v3 即时对战系统表 (3v3 Real-Time Battle System)
-- =============================================================================

-- 1. 对战会话表 (Battle Sessions)
CREATE TABLE IF NOT EXISTS battle_sessions (
    id SERIAL PRIMARY KEY,
    session_id VARCHAR(255) NOT NULL UNIQUE,    -- 对战会话 ID
    player_1_id VARCHAR(255) NOT NULL,          -- 玩家 1 ID
    player_2_id VARCHAR(255) NOT NULL,          -- 玩家 2 ID
    player_3_id VARCHAR(255),                   -- 玩家 3 ID (可选，多队模式)
    battle_type VARCHAR(50) NOT NULL,           -- single, multi (3v3)
    battle_status VARCHAR(50) DEFAULT 'waiting', -- waiting, active, finished
    player_1_team TEXT NOT NULL,                -- 玩家 1 队伍 JSON (3 个配对)
    player_2_team TEXT NOT NULL,                -- 玩家 2 队伍 JSON
    player_3_team TEXT,                         -- 玩家 3 队伍 JSON
    current_round INTEGER DEFAULT 0,            -- 当前回合
    max_rounds INTEGER DEFAULT 20,              -- 最大回合数
    player_1_hp JSONB,                          -- 玩家 1 队伍当前 HP
    player_2_hp JSONB,                          -- 玩家 2 队伍当前 HP
    player_3_hp JSONB,                          -- 玩家 3 队伍当前 HP
    player_1_score INTEGER DEFAULT 0,           -- 玩家 1 得分 (积分战用)
    player_2_score INTEGER DEFAULT 0,           -- 玩家 2 得分
    player_3_score INTEGER DEFAULT 0,           -- 玩家 3 得分
    winner_id VARCHAR(255),                     -- 获胜者 ID
    winner_team TEXT,                           -- 获胜者队伍信息
    battle_log TEXT,                            -- 对战日志 JSON
    started_at TIMESTAMP,
    ended_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_battles_player1 ON battle_sessions(player_1_id);
CREATE INDEX IF NOT EXISTS idx_battles_player2 ON battle_sessions(player_2_id);
CREATE INDEX IF NOT EXISTS idx_battles_status ON battle_sessions(battle_status);

-- 2. 对战行动表 (Battle Actions)
CREATE TABLE IF NOT EXISTS battle_actions (
    id SERIAL PRIMARY KEY,
    session_id VARCHAR(255) NOT NULL,
    round_number INTEGER NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    action_type VARCHAR(50),                    -- move, switch, mega-evolve, sync-move
    action_data JSONB,                          -- 行动详细信息
    timestamp TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (session_id) REFERENCES battle_sessions(session_id)
);

-- 3. 对战统计表 (Battle Statistics)
CREATE TABLE IF NOT EXISTS battle_statistics (
    id SERIAL PRIMARY KEY,
    player_id VARCHAR(255) NOT NULL UNIQUE,
    total_battles INTEGER DEFAULT 0,
    wins INTEGER DEFAULT 0,
    losses INTEGER DEFAULT 0,
    win_rate DECIMAL(5, 2),
    total_damage_dealt BIGINT DEFAULT 0,
    total_damage_taken BIGINT DEFAULT 0,
    sync_moves_used INTEGER DEFAULT 0,
    most_used_pokemon VARCHAR(255),
    last_battle_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- =============================================================================
-- 赛事系统表 (World Pokémon Master Tournament System)
-- =============================================================================

-- 1. 赛季表 (Tournament Seasons)
CREATE TABLE IF NOT EXISTS tournament_seasons (
    id SERIAL PRIMARY KEY,
    season_id VARCHAR(255) NOT NULL UNIQUE,
    season_name VARCHAR(100) NOT NULL,
    season_name_zh VARCHAR(100),
    season_number INTEGER NOT NULL,
    description TEXT,
    description_zh TEXT,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    start_rank INTEGER DEFAULT 1200,            -- 初始评分
    max_rank_points INTEGER DEFAULT 3000,       -- 最大评分
    is_active BOOLEAN DEFAULT FALSE,
    reward_pool JSONB,                          -- 奖励池 JSON
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 2. 赛季排名表 (Tournament Rankings)
CREATE TABLE IF NOT EXISTS tournament_rankings (
    id SERIAL PRIMARY KEY,
    ranking_id VARCHAR(255) NOT NULL UNIQUE,
    season_id VARCHAR(255) NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    rank INTEGER,                               -- 当前排名
    rank_points INTEGER DEFAULT 1200,           -- 当前评分
    wins INTEGER DEFAULT 0,                     -- 赛季获胜
    losses INTEGER DEFAULT 0,                   -- 赛季失败
    win_streak INTEGER DEFAULT 0,               -- 连胜数
    loss_streak INTEGER DEFAULT 0,              -- 连败数
    matches_played INTEGER DEFAULT 0,
    highest_rank INTEGER,                       -- 最高排名
    highest_points INTEGER DEFAULT 1200,        -- 最高评分
    last_match_at TIMESTAMP,
    last_rank_update TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (season_id) REFERENCES tournament_seasons(season_id),
    UNIQUE(season_id, player_id)
);

CREATE INDEX IF NOT EXISTS idx_rankings_season ON tournament_rankings(season_id);
CREATE INDEX IF NOT EXISTS idx_rankings_player ON tournament_rankings(player_id);
CREATE INDEX IF NOT EXISTS idx_rankings_rank ON tournament_rankings(rank);

-- 3. 赛事奖励表 (Tournament Rewards)
CREATE TABLE IF NOT EXISTS tournament_rewards (
    id SERIAL PRIMARY KEY,
    reward_id VARCHAR(255) NOT NULL UNIQUE,
    season_id VARCHAR(255) NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    reward_type VARCHAR(50),                    -- gems, coins, items, pokemon
    reward_name VARCHAR(100),
    reward_amount INTEGER,
    rank_earned_at INTEGER,
    claimed BOOLEAN DEFAULT FALSE,
    claimed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (season_id) REFERENCES tournament_seasons(season_id)
);

-- =============================================================================
-- 季节活动系统表 (Seasonal Event System)
-- =============================================================================

-- 1. 活动表 (Events)
CREATE TABLE IF NOT EXISTS seasonal_events (
    id SERIAL PRIMARY KEY,
    event_id VARCHAR(255) NOT NULL UNIQUE,
    event_name VARCHAR(100) NOT NULL,
    event_name_zh VARCHAR(100),
    event_type VARCHAR(50),                     -- story, challenge, time-attack, score-attack
    season VARCHAR(20),                         -- spring, summer, autumn, winter
    description TEXT,
    description_zh TEXT,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    featured_sync_pairs TEXT,                   -- 特色配对 JSON
    difficulty_levels JSONB,                    -- 难度等级
    rewards JSONB,                              -- 奖励信息
    is_active BOOLEAN DEFAULT FALSE,
    is_limited BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 2. 活动进度表 (Event Progress)
CREATE TABLE IF NOT EXISTS event_progress (
    id SERIAL PRIMARY KEY,
    progress_id VARCHAR(255) NOT NULL UNIQUE,
    event_id VARCHAR(255) NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    stage_completed INTEGER DEFAULT 0,
    progress_percentage DECIMAL(5, 2),
    score INTEGER DEFAULT 0,                    -- 积分战用
    time_spent BIGINT DEFAULT 0,                -- 时间战用(毫秒)
    rewards_claimed JSONB,
    last_attempt_at TIMESTAMP,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (event_id) REFERENCES seasonal_events(event_id),
    UNIQUE(event_id, player_id)
);

-- 3. 季节物品表 (Seasonal Items)
CREATE TABLE IF NOT EXISTS seasonal_items (
    id SERIAL PRIMARY KEY,
    item_id VARCHAR(255) NOT NULL UNIQUE,
    event_id VARCHAR(255),
    item_name VARCHAR(100) NOT NULL,
    item_name_zh VARCHAR(100),
    item_type VARCHAR(50),                      -- clothing, accessory, material
    rarity VARCHAR(20),                         -- common, rare, legendary
    effect_description TEXT,
    effect_description_zh TEXT,
    available_seasons TEXT,                     -- JSON 数组: ["spring", "summer"]
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (event_id) REFERENCES seasonal_events(event_id)
);

-- 4. 玩家季节物品表 (Player Seasonal Items)
CREATE TABLE IF NOT EXISTS player_seasonal_items (
    id SERIAL PRIMARY KEY,
    player_id VARCHAR(255) NOT NULL,
    item_id VARCHAR(255) NOT NULL,
    quantity INTEGER DEFAULT 1,
    obtained_at TIMESTAMP DEFAULT NOW(),
    equipped_to_sync_pair VARCHAR(255),         -- 装备在哪个配对上
    FOREIGN KEY (item_id) REFERENCES seasonal_items(item_id),
    UNIQUE(player_id, item_id)
);

-- =============================================================================
-- 拍档招式机制表 (Sync Move System)
-- =============================================================================

-- 1. 拍档技能表 (Sync Skills)
CREATE TABLE IF NOT EXISTS sync_skills (
    id SERIAL PRIMARY KEY,
    skill_id VARCHAR(255) NOT NULL UNIQUE,
    sync_pair_id VARCHAR(255) NOT NULL,
    skill_name VARCHAR(100) NOT NULL,
    skill_name_zh VARCHAR(100),
    skill_type VARCHAR(50),                     -- passive, active
    description TEXT,
    description_zh TEXT,
    effect_power INTEGER,
    required_level INTEGER DEFAULT 1,
    unlock_condition TEXT,                      -- 解锁条件
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (sync_pair_id) REFERENCES sync_pairs(sync_pair_id)
);

-- 2. 拍档招式冷却表 (Sync Move Cooldown)
CREATE TABLE IF NOT EXISTS sync_move_cooldown (
    id SERIAL PRIMARY KEY,
    player_id VARCHAR(255) NOT NULL,
    sync_pair_id VARCHAR(255) NOT NULL,
    last_used_at TIMESTAMP,
    cooldown_until TIMESTAMP,
    times_used_today INTEGER DEFAULT 0,
    max_uses_per_day INTEGER DEFAULT 10,
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(player_id, sync_pair_id)
);

-- =============================================================================
-- 玩家配对管理表 (Player Sync Pair Management)
-- =============================================================================

-- 1. 玩家配对库表 (Player Dex - 玩家已有的配对)
CREATE TABLE IF NOT EXISTS player_sync_pair_dex (
    id SERIAL PRIMARY KEY,
    dex_id VARCHAR(255) NOT NULL UNIQUE,
    player_id VARCHAR(255) NOT NULL,
    sync_pair_id VARCHAR(255) NOT NULL,
    sync_pair_number_owned INTEGER DEFAULT 1,   -- 拥有数量 (支持重复获得)
    is_favorite BOOLEAN DEFAULT FALSE,
    obtained_date TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (sync_pair_id) REFERENCES sync_pairs(sync_pair_id),
    UNIQUE(player_id, sync_pair_id)
);

-- 2. 玩家当前队伍表 (Player Battle Teams)
CREATE TABLE IF NOT EXISTS player_battle_teams (
    id SERIAL PRIMARY KEY,
    team_id VARCHAR(255) NOT NULL UNIQUE,
    player_id VARCHAR(255) NOT NULL,
    team_name VARCHAR(100),
    sync_pair_1_id VARCHAR(255),                -- 队伍位置 1
    sync_pair_2_id VARCHAR(255),                -- 队伍位置 2
    sync_pair_3_id VARCHAR(255),                -- 队伍位置 3
    is_main_team BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(player_id, team_id)
);

-- =============================================================================
-- 辅助统计表 (Support Statistics)
-- =============================================================================

-- 1. 玩家游戏统计表 (Player Game Statistics)
CREATE TABLE IF NOT EXISTS player_game_statistics (
    id SERIAL PRIMARY KEY,
    player_id VARCHAR(255) NOT NULL UNIQUE,
    total_sync_pairs INTEGER DEFAULT 0,         -- 拥有的配对总数
    total_level_sum BIGINT DEFAULT 0,           -- 所有配对等级总和
    average_sync_pair_level DECIMAL(5, 2),
    sync_pairs_with_max_level INTEGER DEFAULT 0,
    highest_rarity_obtained INTEGER DEFAULT 3,  -- 最高稀有度
    total_power_index BIGINT DEFAULT 0,         -- 总战力指数
    last_updated TIMESTAMP DEFAULT NOW()
);

-- 2. 活动完成度表 (Event Completion Statistics)
CREATE TABLE IF NOT EXISTS event_completion_stats (
    id SERIAL PRIMARY KEY,
    player_id VARCHAR(255) NOT NULL,
    event_id VARCHAR(255) NOT NULL,
    completion_percentage DECIMAL(5, 2),
    difficulty_max_cleared VARCHAR(50),         -- Normal, Hard, Very Hard, Expert
    rewards_max_tier INTEGER,
    first_cleared_at TIMESTAMP,
    last_cleared_at TIMESTAMP,
    total_clears INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(player_id, event_id)
);

-- =============================================================================
-- 创建所有必需的索引
-- =============================================================================

-- 性能优化索引
CREATE INDEX IF NOT EXISTS idx_sync_pairs_level ON sync_pairs(level);
CREATE INDEX IF NOT EXISTS idx_sync_pairs_active ON sync_pairs(is_active);
CREATE INDEX IF NOT EXISTS idx_tournament_rankings_points ON tournament_rankings(rank_points);
CREATE INDEX IF NOT EXISTS idx_seasonal_events_active ON seasonal_events(is_active);
CREATE INDEX IF NOT EXISTS idx_seasonal_events_dates ON seasonal_events(start_date, end_date);
CREATE INDEX IF NOT EXISTS idx_event_progress_player ON event_progress(player_id);
CREATE INDEX IF NOT EXISTS idx_player_dex_player ON player_sync_pair_dex(player_id);
CREATE INDEX IF NOT EXISTS idx_battle_teams_player ON player_battle_teams(player_id);

-- 外键索引 (自动创建)

-- =============================================================================
-- 创建视图 (Views) 用于简化复杂查询
-- =============================================================================

-- 1. 玩家队伍视图 (Player Team View)
CREATE OR REPLACE VIEW player_team_view AS
SELECT 
    pbt.team_id,
    pbt.player_id,
    pbt.team_name,
    sp1.sync_pair_name as position_1_name,
    sp2.sync_pair_name as position_2_name,
    sp3.sync_pair_name as position_3_name,
    COALESCE(sp1.level, 0) + COALESCE(sp2.level, 0) + COALESCE(sp3.level, 0) as total_level,
    pbt.is_main_team
FROM player_battle_teams pbt
LEFT JOIN sync_pairs sp1 ON pbt.sync_pair_1_id = sp1.sync_pair_id
LEFT JOIN sync_pairs sp2 ON pbt.sync_pair_2_id = sp2.sync_pair_id
LEFT JOIN sync_pairs sp3 ON pbt.sync_pair_3_id = sp3.sync_pair_id;

-- 2. 赛季排名视图 (Season Rankings View)
CREATE OR REPLACE VIEW tournament_ranking_view AS
SELECT 
    tr.rank,
    tr.rank_points,
    tr.player_id,
    tr.wins,
    tr.losses,
    tr.matches_played,
    ts.season_name,
    tr.season_id,
    CASE 
        WHEN tr.matches_played > 0 THEN ROUND((tr.wins::NUMERIC / tr.matches_played) * 100, 2)
        ELSE 0
    END as win_rate
FROM tournament_rankings tr
JOIN tournament_seasons ts ON tr.season_id = ts.season_id
ORDER BY tr.rank_points DESC;

-- 3. 活动参与视图 (Event Participation View)
CREATE OR REPLACE VIEW event_participation_view AS
SELECT 
    se.event_id,
    se.event_name,
    se.event_name_zh,
    COUNT(DISTINCT ep.player_id) as total_participants,
    AVG(ep.progress_percentage) as avg_completion,
    MAX(ep.score) as highest_score,
    se.start_date,
    se.end_date,
    se.is_active
FROM seasonal_events se
LEFT JOIN event_progress ep ON se.event_id = ep.event_id
GROUP BY se.event_id, se.event_name, se.event_name_zh, se.start_date, se.end_date, se.is_active;

-- =============================================================================
-- 完成标记 (Schema Version)
-- =============================================================================
-- Schema Version: 2.1.0
-- Created: 2026-04-14 15:30 UTC
-- Status: Production Ready ✅
-- Features: Sync Pair System, 3v3 Real-Time Battle, Tournament, Seasonal Events, Sync Moves
