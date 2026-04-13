-- Judge Server 用户账户管理表
-- 这些SQL应该添加到 judge-server/internal/db/database.go 的 InitSchema() 函数中

-- 用户账户表 (Users Accounts)
CREATE TABLE IF NOT EXISTS user_accounts (
    id SERIAL PRIMARY KEY,
    github_id INTEGER UNIQUE NOT NULL,
    github_login VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    avatar_url TEXT,
    balance DECIMAL(12, 2) DEFAULT 0.0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 用户精灵表 (User Pokemons)
-- 关联到 pets 表，但提供用户友好的查询
CREATE TABLE IF NOT EXISTS user_pokemons (
    id SERIAL PRIMARY KEY,
    github_id INTEGER NOT NULL,
    pet_id TEXT NOT NULL UNIQUE,
    pet_name TEXT NOT NULL,
    level INTEGER DEFAULT 1,
    species TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (github_id) REFERENCES user_accounts(github_id),
    FOREIGN KEY (pet_id) REFERENCES pets(id)
);

-- 用户物品表 (User Inventory)
CREATE TABLE IF NOT EXISTS user_inventory (
    id SERIAL PRIMARY KEY,
    github_id INTEGER NOT NULL,
    item_id VARCHAR(255) NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    quantity INTEGER DEFAULT 1,
    acquired_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (github_id) REFERENCES user_accounts(github_id),
    UNIQUE(github_id, item_id)
);

-- 账户交易记录 (Account Transactions)
CREATE TABLE IF NOT EXISTS account_transactions (
    id SERIAL PRIMARY KEY,
    github_id INTEGER NOT NULL,
    transaction_type VARCHAR(50) NOT NULL,
    amount DECIMAL(12, 2) NOT NULL,
    description TEXT,
    balance_before DECIMAL(12, 2),
    balance_after DECIMAL(12, 2),
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (github_id) REFERENCES user_accounts(github_id)
);

-- 创建索引以提高查询性能
CREATE INDEX IF NOT EXISTS idx_user_accounts_github_id ON user_accounts(github_id);
CREATE INDEX IF NOT EXISTS idx_user_pokemons_github_id ON user_pokemons(github_id);
CREATE INDEX IF NOT EXISTS idx_user_inventory_github_id ON user_inventory(github_id);
CREATE INDEX IF NOT EXISTS idx_account_transactions_github_id ON account_transactions(github_id);
CREATE INDEX IF NOT EXISTS idx_account_transactions_created_at ON account_transactions(created_at DESC);
