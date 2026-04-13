-- Judge Server 扩展数据库 Schema
-- 包含农场、食物、Cookie、蛋等系统的完整表定义
-- 这个文件应该在 judge-server/internal/db/database.go 的 InitSchema() 中执行

-- =============================================================================
-- 农场系统表 (Farm System)
-- =============================================================================

CREATE TABLE IF NOT EXISTS farms (
    id SERIAL PRIMARY KEY,
    farm_key VARCHAR(255) UNIQUE NOT NULL,  -- owner/repository
    owner_id VARCHAR(255) NOT NULL,
    repository_name VARCHAR(255) NOT NULL,
    repository_url TEXT NOT NULL,
    planted_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(owner_id, repository_name)
);

-- 农场中的食物 (Foods in Farm)
CREATE TABLE IF NOT EXISTS farm_foods (
    id SERIAL PRIMARY KEY,
    farm_id INTEGER NOT NULL,
    food_id VARCHAR(255) NOT NULL UNIQUE,
    food_type VARCHAR(50) NOT NULL,  -- cookie, donut, apple, gene
    emoji VARCHAR(10),
    current_quantity INTEGER DEFAULT 0,
    max_quantity INTEGER DEFAULT 0,
    regeneration_hours INTEGER DEFAULT 24,
    last_eaten_at TIMESTAMP,
    seed_hash VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (farm_id) REFERENCES farms(id) ON DELETE CASCADE
);

-- 农场食物消费历史 (Food Consumption History)
CREATE TABLE IF NOT EXISTS food_consumption_history (
    id SERIAL PRIMARY KEY,
    food_id VARCHAR(255) NOT NULL,
    eater_player_id VARCHAR(255) NOT NULL,
    eater_pet_id VARCHAR(255) NOT NULL,
    consumed_at TIMESTAMP DEFAULT NOW(),
    nutrition_exp INTEGER DEFAULT 0,
    nutrition_energy INTEGER DEFAULT 0,
    FOREIGN KEY (food_id) REFERENCES farm_foods(food_id) ON DELETE CASCADE
);

-- =============================================================================
-- Cookie 系统表 (Cookie System)
-- =============================================================================

CREATE TABLE IF NOT EXISTS cookies (
    id SERIAL PRIMARY KEY,
    cookie_id VARCHAR(255) NOT NULL UNIQUE,  -- 0x... format
    cookie_type VARCHAR(50) NOT NULL,  -- cookie, donut, apple, gene
    emoji VARCHAR(10),
    source_file TEXT,
    generator_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Cookie 索赔记录 (Cookie Claims)
CREATE TABLE IF NOT EXISTS cookie_claims (
    id SERIAL PRIMARY KEY,
    cookie_id VARCHAR(255) NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    claimed_at TIMESTAMP DEFAULT NOW(),
    exp_reward INTEGER DEFAULT 0,
    energy_reward INTEGER DEFAULT 0,
    FOREIGN KEY (cookie_id) REFERENCES cookies(cookie_id) ON DELETE CASCADE,
    UNIQUE(cookie_id, player_id)
);

-- Cookie 事件历史 (Cookie Events)
CREATE TABLE IF NOT EXISTS cookie_events (
    id SERIAL PRIMARY KEY,
    event_type VARCHAR(50) NOT NULL,  -- generated, scanned, claimed, validated
    cookie_id VARCHAR(255),
    player_id VARCHAR(255),
    event_details TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- =============================================================================
-- 蛋系统表 (Egg System)
-- =============================================================================

CREATE TABLE IF NOT EXISTS eggs (
    id SERIAL PRIMARY KEY,
    egg_id VARCHAR(255) NOT NULL UNIQUE,
    owner_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    incubation_hours INTEGER DEFAULT 72,
    stage INTEGER DEFAULT 0,  -- 0-3: 蛋 -> 孵化中 -> 快孵化 -> 可孵化
    hatched_at TIMESTAMP,
    hatched_pet_id VARCHAR(255),
    attributes TEXT,  -- JSON string for gene data
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 蛋事件历史 (Egg Events)
CREATE TABLE IF NOT EXISTS egg_events (
    id SERIAL PRIMARY KEY,
    egg_id VARCHAR(255) NOT NULL,
    event_type VARCHAR(50) NOT NULL,  -- created, incubating, hatched
    event_details TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (egg_id) REFERENCES eggs(egg_id) ON DELETE CASCADE
);

-- =============================================================================
-- 蛋碎片/Cookie 碎片系统 (Fragment System)
-- =============================================================================

CREATE TABLE IF NOT EXISTS cookie_fragments (
    id SERIAL PRIMARY KEY,
    fragment_id VARCHAR(255) NOT NULL UNIQUE,
    player_id VARCHAR(255) NOT NULL,
    fragment_type VARCHAR(50) NOT NULL,  -- from_cookie, from_egg, from_battle
    source_id VARCHAR(255),  -- cookie_id, egg_id, battle_id
    collected_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);

-- 用户 Cookie 碎片余额 (Player Fragment Balance)
CREATE TABLE IF NOT EXISTS player_fragment_balance (
    id SERIAL PRIMARY KEY,
    player_id VARCHAR(255) NOT NULL UNIQUE,
    total_fragments INTEGER DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW()
);

-- =============================================================================
-- 全局商店系统 (Shop System)
-- =============================================================================

CREATE TABLE IF NOT EXISTS shop_items (
    id SERIAL PRIMARY KEY,
    item_id VARCHAR(255) NOT NULL UNIQUE,
    item_name VARCHAR(255) NOT NULL,
    description TEXT,
    item_type VARCHAR(50) NOT NULL,  -- consumable, equipment, etc
    price INTEGER NOT NULL DEFAULT 0,
    emoji VARCHAR(10),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 商店库存 (Shop Inventory)
CREATE TABLE IF NOT EXISTS shop_inventory (
    id SERIAL PRIMARY KEY,
    item_id VARCHAR(255) NOT NULL,
    current_stock INTEGER DEFAULT 0,
    max_stock INTEGER DEFAULT 999,
    restocked_at TIMESTAMP,
    FOREIGN KEY (item_id) REFERENCES shop_items(item_id) ON DELETE CASCADE,
    UNIQUE(item_id)
);

-- 商店交易 (Shop Transactions)
CREATE TABLE IF NOT EXISTS shop_transactions (
    id SERIAL PRIMARY KEY,
    player_id VARCHAR(255) NOT NULL,
    item_id VARCHAR(255) NOT NULL,
    quantity INTEGER DEFAULT 1,
    unit_price DECIMAL(12, 2),
    total_price DECIMAL(12, 2),
    transaction_type VARCHAR(50),  -- buy, sell
    transaction_date TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (item_id) REFERENCES shop_items(item_id) ON DELETE CASCADE
);

-- =============================================================================
-- 玩家农场访问权限 (Farm Access Control)
-- =============================================================================

CREATE TABLE IF NOT EXISTS farm_permissions (
    id SERIAL PRIMARY KEY,
    farm_id INTEGER NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    permission_type VARCHAR(50) NOT NULL,  -- viewer, visitor, friend
    granted_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (farm_id) REFERENCES farms(id) ON DELETE CASCADE,
    UNIQUE(farm_id, player_id)
);

-- 玩家农场访问历史 (Farm Visit History)
CREATE TABLE IF NOT EXISTS farm_visits (
    id SERIAL PRIMARY KEY,
    farm_id INTEGER NOT NULL,
    visitor_id VARCHAR(255) NOT NULL,
    visited_at TIMESTAMP DEFAULT NOW(),
    action VARCHAR(50),  -- view_farm, eat_food, inspect_plants
    FOREIGN KEY (farm_id) REFERENCES farms(id) ON DELETE CASCADE
);

-- =============================================================================
-- 索引优化 (Indexes for Performance)
-- =============================================================================

CREATE INDEX IF NOT EXISTS idx_farms_owner_id ON farms(owner_id);
CREATE INDEX IF NOT EXISTS idx_farms_farm_key ON farms(farm_key);

CREATE INDEX IF NOT EXISTS idx_farm_foods_farm_id ON farm_foods(farm_id);
CREATE INDEX IF NOT EXISTS idx_farm_foods_food_type ON farm_foods(food_type);
CREATE INDEX IF NOT EXISTS idx_farm_foods_food_id ON farm_foods(food_id);

CREATE INDEX IF NOT EXISTS idx_food_consumption_food_id ON food_consumption_history(food_id);
CREATE INDEX IF NOT EXISTS idx_food_consumption_eater ON food_consumption_history(eater_player_id);
CREATE INDEX IF NOT EXISTS idx_food_consumption_date ON food_consumption_history(consumed_at DESC);

CREATE INDEX IF NOT EXISTS idx_cookies_cookie_id ON cookies(cookie_id);
CREATE INDEX IF NOT EXISTS idx_cookies_cookie_type ON cookies(cookie_type);

CREATE INDEX IF NOT EXISTS idx_cookie_claims_cookie_id ON cookie_claims(cookie_id);
CREATE INDEX IF NOT EXISTS idx_cookie_claims_player_id ON cookie_claims(player_id);
CREATE INDEX IF NOT EXISTS idx_cookie_claims_date ON cookie_claims(claimed_at DESC);

CREATE INDEX IF NOT EXISTS idx_cookie_events_event_type ON cookie_events(event_type);
CREATE INDEX IF NOT EXISTS idx_cookie_events_date ON cookie_events(created_at DESC);

CREATE INDEX IF NOT EXISTS idx_eggs_egg_id ON eggs(egg_id);
CREATE INDEX IF NOT EXISTS idx_eggs_owner_id ON eggs(owner_id);
CREATE INDEX IF NOT EXISTS idx_eggs_stage ON eggs(stage);
CREATE INDEX IF NOT EXISTS idx_eggs_created_at ON eggs(created_at DESC);

CREATE INDEX IF NOT EXISTS idx_egg_events_egg_id ON egg_events(egg_id);
CREATE INDEX IF NOT EXISTS idx_egg_events_type ON egg_events(event_type);

CREATE INDEX IF NOT EXISTS idx_cookie_fragments_player_id ON cookie_fragments(player_id);
CREATE INDEX IF NOT EXISTS idx_cookie_fragments_type ON cookie_fragments(fragment_type);

CREATE INDEX IF NOT EXISTS idx_player_fragment_balance_player_id ON player_fragment_balance(player_id);

CREATE INDEX IF NOT EXISTS idx_shop_items_item_id ON shop_items(item_id);
CREATE INDEX IF NOT EXISTS idx_shop_items_available ON shop_items(is_available);

CREATE INDEX IF NOT EXISTS idx_shop_inventory_item_id ON shop_inventory(item_id);

CREATE INDEX IF NOT EXISTS idx_shop_transactions_player_id ON shop_transactions(player_id);
CREATE INDEX IF NOT EXISTS idx_shop_transactions_date ON shop_transactions(transaction_date DESC);

CREATE INDEX IF NOT EXISTS idx_farm_permissions_farm_id ON farm_permissions(farm_id);
CREATE INDEX IF NOT EXISTS idx_farm_permissions_player_id ON farm_permissions(player_id);

CREATE INDEX IF NOT EXISTS idx_farm_visits_farm_id ON farm_visits(farm_id);
CREATE INDEX IF NOT EXISTS idx_farm_visits_visitor_id ON farm_visits(visitor_id);
CREATE INDEX IF NOT EXISTS idx_farm_visits_date ON farm_visits(visited_at DESC);
