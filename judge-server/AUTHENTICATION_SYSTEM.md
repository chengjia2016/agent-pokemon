# 🎮 Agent Monster 认证系统 - 多客户端支持

本文档描述了新的认证系统设计，它脱离了对 GitHub CLI 的依赖，支持多个客户端（Claude、OpenCode、OpenClaw、Gemini 等）。

## 核心设计原则

### 1. **脱离 GitHub CLI 依赖**
- 原系统依赖 `gh auth` 进行认证
- 新系统支持三种认证方式：
  - **测试账户认证**：用于开发和测试
  - **Server Token 认证**：用于已有账户的直接访问
  - **GitHub OAuth 绑定**：用于 GitHub 账户绑定

### 2. **安全性**
- 使用 bcrypt 对密码进行哈希处理
- Server Token 是加密的 32 字节十六进制字符串
- Access Token 是 24 字节十六进制字符串，有效期 30 天
- Player Token（GitHub 绑定）有效期 24 小时

### 3. **多客户端支持**
- 每个客户端（Claude、OpenCode、Gemini 等）都有独立的 session
- 支持从同一账户同时登录多个客户端
- 每个 session 都有独立的 access token

## 系统架构

```
┌─────────────────────────────────────────────────────────────┐
│                    Multiple Clients                          │
│  ┌──────────────┬──────────────┬──────────────┬──────────────┐
│  │    Claude    │   OpenCode   │  OpenClaw    │    Gemini    │
│  └──────────────┴──────────────┴──────────────┴──────────────┘
│                        ↓
│                ┌─────────────────┐
│                │   Auth Handler   │
│                │  - Login         │
│                │  - Register      │
│                │  - Bind GitHub   │
│                └─────────────────┘
│                        ↓
│        ┌───────────────────────────────────┐
│        │        AuthManager Service         │
│        │  - Token generation               │
│        │  - Authentication logic           │
│        │  - Session management             │
│        └───────────────────────────────────┘
│                        ↓
│     ┌──────────────────────────────────────────────┐
│     │          Database Layer (PostgreSQL)         │
│     │  ┌─────────────────────────────────────────┐ │
│     │  │ player_profiles                         │ │
│     │  │ client_sessions                         │ │
│     │  │ test_accounts                           │ │
│     │  │ player_tokens (GitHub binding)          │ │
│     │  └─────────────────────────────────────────┘ │
│     └──────────────────────────────────────────────┘
```

## 数据库表结构

### 1. `player_profiles` - 玩家档案
```sql
CREATE TABLE player_profiles (
    id SERIAL PRIMARY KEY,
    player_id VARCHAR(255) UNIQUE NOT NULL,      -- 唯一玩家ID
    github_id INTEGER,                            -- GitHub ID (可选)
    github_login VARCHAR(255),                    -- GitHub 用户名 (可选)
    email VARCHAR(255),                           -- 邮箱
    avatar_url VARCHAR(255),                      -- 头像 URL
    is_test_account BOOLEAN DEFAULT false,        -- 是否为测试账户
    server_token VARCHAR(255) UNIQUE,             -- 服务器令牌（用于直接认证）
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    last_login_at TIMESTAMP,
    token_expire_at TIMESTAMP                     -- Token 过期时间（可选）
);
```

### 2. `client_sessions` - 客户端会话
```sql
CREATE TABLE client_sessions (
    id SERIAL PRIMARY KEY,
    player_id VARCHAR(255) NOT NULL REFERENCES player_profiles(player_id),
    client_name VARCHAR(100) NOT NULL,            -- "claude", "opencode", "openclaw", "gemini"
    access_token VARCHAR(255) UNIQUE NOT NULL,    -- 客户端访问令牌
    client_type VARCHAR(50),                      -- "web", "cli", "agent", etc.
    last_active_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP                          -- Session 过期时间
);
```

### 3. `test_accounts` - 测试账户
```sql
CREATE TABLE test_accounts (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,        -- 测试账户用户名
    password VARCHAR(255) NOT NULL,               -- 密码（bcrypt 哈希）
    player_id VARCHAR(255) UNIQUE NOT NULL REFERENCES player_profiles(player_id),
    status VARCHAR(20) DEFAULT 'active',          -- "active", "inactive", "locked"
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 4. `player_tokens` - GitHub 绑定令牌
```sql
CREATE TABLE player_tokens (
    id SERIAL PRIMARY KEY,
    player_id VARCHAR(255),                       -- 玩家ID（可能为空）
    github_id INTEGER,                            -- GitHub ID
    github_login VARCHAR(255),                    -- GitHub 用户名
    token VARCHAR(255) UNIQUE NOT NULL,           -- 绑定令牌
    is_used BOOLEAN DEFAULT false,                -- 是否已使用
    used_at TIMESTAMP,                            -- 使用时间
    expires_at TIMESTAMP,                         -- 过期时间
    created_at TIMESTAMP DEFAULT NOW()
);
```

## API 端点

### 认证相关

#### 1. 注册测试账户
```bash
POST /api/auth/register-test
Content-Type: application/json

{
    "username": "test_player_1",
    "password": "password123"
}

Response (201):
{
    "success": true,
    "player_id": "player_a1b2c3d4e5f6g7h8",
    "server_token": "abc123...xyz789",
    "access_token": "token_xyz...",
    "username": "test_player_1",
    "message": "Test account created successfully"
}
```

#### 2. 登录
```bash
POST /api/auth/login
Content-Type: application/json

# 测试账户登录
{
    "username": "test_player_1",
    "password": "password123",
    "client_name": "claude"
}

# 或使用 Server Token
{
    "server_token": "abc123...xyz789",
    "client_name": "opencode"
}

Response (200):
{
    "success": true,
    "player_id": "player_a1b2c3d4e5f6g7h8",
    "server_token": "abc123...xyz789",
    "access_token": "token_xyz...",
    "client_session": {
        "id": 1,
        "player_id": "player_a1b2c3d4e5f6g7h8",
        "client_name": "claude",
        "access_token": "token_xyz...",
        "client_type": "agent",
        "last_active_at": "2026-04-13T12:34:56Z",
        "created_at": "2026-04-13T12:34:56Z",
        "expires_at": "2026-05-13T12:34:56Z"
    },
    "message": "Login successful"
}
```

#### 3. 列出测试账户（管理员）
```bash
GET /api/auth/test-accounts

Response (200):
{
    "success": true,
    "accounts": [
        {
            "id": 1,
            "username": "test_player_1",
            "password": "***",
            "player_id": "player_a1b2c3d4e5f6g7h8",
            "status": "active",
            "created_at": "2026-04-01T10:00:00Z",
            "updated_at": "2026-04-13T12:34:56Z"
        },
        ...
    ],
    "total": 10
}
```

#### 4. 创建 GitHub 绑定令牌
```bash
POST /api/auth/create-player-token
Content-Type: application/json

{
    "github_id": 123456,
    "github_login": "myusername"
}

Response (201):
{
    "success": true,
    "player_token": "binding_token_xyz...",
    "expires_at": "2026-04-14T12:34:56Z",
    "message": "Player token created. Use this token to bind your GitHub account."
}
```

#### 5. 绑定 GitHub 账户
```bash
POST /api/auth/bind-github
Content-Type: application/json

{
    "player_token": "binding_token_xyz...",
    "client_name": "claude"
}

Response (200):
{
    "success": true,
    "player_id": "player_a1b2c3d4e5f6g7h8",
    "server_token": "abc123...xyz789",
    "access_token": "token_xyz...",
    "client_session": {...},
    "message": "GitHub account bound successfully"
}
```

#### 6. 验证令牌
```bash
POST /api/auth/validate-token
Content-Type: application/json

{
    "access_token": "token_xyz..."
}

Response (200):
{
    "success": true,
    "player_id": "player_a1b2c3d4e5f6g7h8",
    "valid": true
}
```

## 认证流程示例

### 流程 1: 测试账户登录
```
1. 用户输入用户名和密码
   ↓
2. POST /api/auth/login (username + password)
   ↓
3. AuthManager.AuthenticateTestAccount()
   - 查找测试账户
   - 验证密码
   - 返回 player_profile
   ↓
4. 创建 client_session
   ↓
5. 返回 access_token 和 server_token
   ↓
6. 后续请求使用 Authorization: <access_token>
```

### 流程 2: GitHub 账户绑定
```
1. GitHub 用户在 GitHub App 中认证
   ↓
2. 获取 github_id 和 github_login
   ↓
3. POST /api/auth/create-player-token (github_id + github_login)
   ↓
4. 返回 player_token（24小时有效）
   ↓
5. 用户在另一个客户端（如 Claude）中：
   POST /api/auth/bind-github (player_token + client_name)
   ↓
6. AuthManager.UsePlayerTokenToBindGitHub()
   - 验证 player_token
   - 创建或更新 player_profile
   - 标记 token 为已使用
   - 创建 client_session
   ↓
7. 返回 access_token
   ↓
8. 后续在该客户端使用 access_token 认证
```

### 流程 3: Server Token 直接认证
```
1. 已有 server_token（来自之前的登录或创建）
   ↓
2. POST /api/auth/login (server_token + client_name)
   ↓
3. AuthManager.AuthenticateWithServerToken()
   - 验证 server_token
   - 返回 player_profile
   ↓
4. 创建新的 client_session
   ↓
5. 返回新的 access_token
```

## 使用示例

### Python 管理工具

```bash
# 创建测试账户
python3 test_account_manager.py create --username test_player_1 --password password123

# 列出所有测试账户
python3 test_account_manager.py list

# 登录测试账户
python3 test_account_manager.py login --username test_player_1 --password password123 --client-name claude

# 创建 GitHub 绑定令牌
python3 test_account_manager.py create-github-token --github-id 123456 --github-login myusername

# 绑定 GitHub 账户
python3 test_account_manager.py bind-github --player-token <token> --client-name opencode

# 验证 access token
python3 test_account_manager.py validate --token <access_token>
```

### Bash 初始化脚本

```bash
# 初始化 10 个测试账户
bash init_test_accounts.sh
```

### 直接 API 调用

```bash
# 1. 创建测试账户
curl -X POST http://localhost:8080/api/auth/register-test \
  -H "Content-Type: application/json" \
  -d '{"username": "player1", "password": "pass123"}'

# 2. 登录
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "player1", "password": "pass123", "client_name": "claude"}'

# 3. 在后续请求中使用 access_token
curl -H "Authorization: <access_token>" \
  http://localhost:8080/api/user/profile
```

## 集成到 Judge Server

### 1. 更新 main.go

在 `cmd/main.go` 中添加路由：

```go
// 初始化认证系统
authManager := service.NewAuthManager(database)
authHandler := handler.NewAuthHandler(authManager)
authMiddleware := handler.NewAuthMiddleware(authHandler)

// 注册路由
mux.HandleFunc("/api/auth/register-test", authHandler.RegisterTestAccount)
mux.HandleFunc("/api/auth/login", authHandler.Login)
mux.HandleFunc("/api/auth/test-accounts", authHandler.ListTestAccounts)
mux.HandleFunc("/api/auth/create-player-token", authHandler.CreatePlayerTokenForGitHub)
mux.HandleFunc("/api/auth/bind-github", authHandler.BindGitHub)
mux.HandleFunc("/api/auth/validate-token", authHandler.ValidateToken)
mux.HandleFunc("/api/auth/profile", authMiddleware.RequireAuth(authHandler.GetPlayerProfile))
```

### 2. 初始化数据库

在 `database.InitSchema()` 之后调用：

```go
if err := database.InitAuthSchema(); err != nil {
    log.Fatalf("Failed to initialize auth schema: %v", err)
}
```

### 3. 依赖安装

```bash
go get golang.org/x/crypto/bcrypt
```

## 安全建议

1. **HTTPS**: 在生产环境中总是使用 HTTPS
2. **Token 管理**:
   - 定期轮换 server_token
   - 自动清理过期的 session
   - 实现令牌黑名单机制
3. **速率限制**: 在登录端点实现速率限制防止暴力破解
4. **日志审计**: 记录所有认证相关事件
5. **密码策略**: 强制测试账户使用强密码

## 故障排除

### 问题 1: Token 已过期

```
Error: access token expired
```

**解决方案**: 使用 server_token 重新登录以获取新的 access_token

### 问题 2: Player Token 无法使用

```
Error: player token not found or expired
```

**解决方案**: Player token 有效期为 24 小时，需要重新生成

### 问题 3: 测试账户已存在

```
Error: player profile already exists
```

**解决方案**: 直接登录已有账户，或使用不同的用户名

## 未来扩展

- OAuth2 全流程集成
- WebAuthn/FIDO2 支持
- 双因素认证 (2FA)
- SSO 集成
- 社交登录（Google、Discord 等）
