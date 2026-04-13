# 🔗 认证系统集成指南

本文档描述了如何将新的认证系统集成到现有的 Judge Server。

## 实现步骤

### 第 1 步：检查并执行 SQL 初始化

认证系统需要以下表格。验证这些表是否已创建：

```bash
psql -h localhost -U postgres -d agent_monster << 'EOF'

-- 检查表是否存在
SELECT tablename FROM pg_tables 
WHERE tablename IN ('player_profiles', 'client_sessions', 'test_accounts', 'player_tokens');

EOF
```

如果表不存在，Go 程序会在启动时自动创建它们（通过 `database.InitAuthSchema()`）。

### 第 2 步：修改 main.go

编辑 `/root/petskill/judge-server/cmd/main.go`：

#### 2.1 在 imports 中添加认证服务

```go
import (
    // ... existing imports ...
    "judge-server/internal/service"
    // ... rest of imports ...
)
```

#### 2.2 在初始化数据库后添加认证模式初始化

在 `database.InitSchema()` 之后添加：

```go
// 初始化认证系统表
if err := database.InitAuthSchema(); err != nil {
    log.Fatalf("Failed to initialize auth schema: %v", err)
}
```

找到这行：
```go
// Initialize battle system schema
if err := database.InitBattleSystemSchema(); err != nil {
    log.Fatalf("Failed to initialize battle system schema: %v", err)
}
```

在其后面添加：
```go
// Initialize authentication system
if err := database.InitAuthSchema(); err != nil {
    log.Fatalf("Failed to initialize authentication schema: %v", err)
}
```

#### 2.3 创建认证管理器和处理器

在创建 HTTP 服务器之前添加：

```go
// Initialize authentication system
authManager := service.NewAuthManager(database)
authHandler := handler.NewAuthHandler(authManager)
authMiddleware := handler.NewAuthMiddleware(authHandler)
```

#### 2.4 注册认证路由

在设置 HTTP 路由的地方添加（通常在 `http.HandleFunc` 调用之后）：

```go
// Authentication endpoints
http.HandleFunc("/api/auth/register-test", authHandler.RegisterTestAccount)
http.HandleFunc("/api/auth/login", authHandler.Login)
http.HandleFunc("/api/auth/test-accounts", authHandler.ListTestAccounts)
http.HandleFunc("/api/auth/create-player-token", authHandler.CreatePlayerTokenForGitHub)
http.HandleFunc("/api/auth/bind-github", authHandler.BindGitHub)
http.HandleFunc("/api/auth/validate-token", authHandler.ValidateToken)

// Protected endpoint example
http.HandleFunc("/api/user/profile", authMiddleware.RequireAuth(authHandler.GetPlayerProfile))
```

### 第 3 步：修改现有的 UserAccount 相关 API

现在系统有两个并存的用户模型：
- **旧的**: `UserAccount`（在 `users.go`）
- **新的**: `PlayerProfile`（在 `auth.go`）

#### 选项 A: 逐步迁移（推荐）

同时支持两个模型，后续逐步迁移：

```go
// 在现有的 CreateUserAccount handler 中添加认证
func (h *Handler) CreateUserAccount(w http.ResponseWriter, r *http.Request) {
    // ... existing code ...
    
    // 也创建对应的 PlayerProfile
    if authManager != nil {
        profile, _ := authManager.CreateOrUpdatePlayerProfileFromGitHub(
            req.GithubID, req.GithubLogin, req.Email, req.AvatarURL)
        // Store profile.PlayerID for future use
    }
}
```

#### 选项 B: 完全替换

将 `UserAccount` 替换为 `PlayerProfile`。

### 第 4 步：安装依赖

```bash
cd /root/petskill/judge-server
go get golang.org/x/crypto/bcrypt
```

如果遇到版本问题：
```bash
go get -u golang.org/x/crypto
```

### 第 5 步：测试编译

```bash
cd /root/petskill/judge-server
go build ./cmd/main.go
```

## 完整的集成示例

完整的 `cmd/main.go` 片段（相关部分）：

```go
package main

import (
    // ... existing imports ...
    "judge-server/internal/service"
    "judge-server/internal/handler"
    // ... rest of imports ...
)

func main() {
    // ... existing setup code ...

    // Initialize database
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

    // Initialize schemas
    if err := database.InitSchema(); err != nil {
        log.Fatalf("Failed to initialize schema: %v", err)
    }

    if err := database.UpdateSchema(); err != nil {
        log.Fatalf("Failed to update schema with user tables: %v", err)
    }

    if err := database.InitBattleSystemSchema(); err != nil {
        log.Fatalf("Failed to initialize battle system schema: %v", err)
    }

    // NEW: Initialize authentication system
    if err := database.InitAuthSchema(); err != nil {
        log.Fatalf("Failed to initialize authentication schema: %v", err)
    }

    // ... rest of initialization ...

    // NEW: Initialize authentication system
    authManager := service.NewAuthManager(database)
    authHandler := handler.NewAuthHandler(authManager)
    authMiddleware := handler.NewAuthMiddleware(authHandler)

    // NEW: Register authentication endpoints
    http.HandleFunc("/api/auth/register-test", authHandler.RegisterTestAccount)
    http.HandleFunc("/api/auth/login", authHandler.Login)
    http.HandleFunc("/api/auth/test-accounts", authHandler.ListTestAccounts)
    http.HandleFunc("/api/auth/create-player-token", authHandler.CreatePlayerTokenForGitHub)
    http.HandleFunc("/api/auth/bind-github", authHandler.BindGitHub)
    http.HandleFunc("/api/auth/validate-token", authHandler.ValidateToken)
    http.HandleFunc("/api/user/profile", authMiddleware.RequireAuth(authHandler.GetPlayerProfile))

    // ... rest of HTTP handlers ...

    // Start server
    http.ListenAndServe(addr, nil)
}
```

## 验证集成

### 1. 启动服务器

```bash
cd /root/petskill/judge-server
./main
```

### 2. 测试基本认证流程

```bash
# 创建测试账户
curl -X POST http://localhost:8080/api/auth/register-test \
  -H "Content-Type: application/json" \
  -d '{"username": "test_user", "password": "test123"}'

# 应返回 200 Created，包含 player_id 和 server_token

# 登录
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "test_user", "password": "test123", "client_name": "claude"}'

# 应返回 200 OK，包含 access_token

# 列出所有测试账户
curl -X GET http://localhost:8080/api/auth/test-accounts

# 应返回 200 OK，包含账户列表
```

### 3. 检查数据库

```bash
psql -h localhost -U postgres -d agent_monster << 'EOF'

-- 检查玩家档案
SELECT player_id, is_test_account, created_at FROM player_profiles;

-- 检查会话
SELECT * FROM client_sessions;

-- 检查测试账户
SELECT username, status FROM test_accounts;

EOF
```

## 常见问题

### Q: bcrypt 导入失败

```
could not import golang.org/x/crypto/bcrypt
```

**解决方案**:
```bash
go get golang.org/x/crypto/bcrypt
go mod tidy
```

### Q: 表不存在错误

```
ERROR: relation "player_profiles" does not exist
```

**解决方案**: 确保 `database.InitAuthSchema()` 已被调用。可以手动执行 SQL：

```bash
psql -h localhost -U postgres -d agent_monster << 'EOF'
-- Run the CREATE TABLE statements from internal/db/auth.go
EOF
```

### Q: 密钥生成失败

```
failed to hash password: ...
```

**解决方案**: 检查系统的 `/dev/urandom` 是否可读。这通常不会在 Linux 上发生。

## 版本兼容性

- **Go**: 1.16+
- **PostgreSQL**: 12+
- **crypto/bcrypt**: golang.org/x/crypto v0.0.0 或更新

## 下一步

1. **集成到现有 API**: 将认证令牌传播到现有的游戏 API
2. **添加授权**: 实现基于角色的访问控制 (RBAC)
3. **审计日志**: 记录所有认证事件
4. **性能优化**: 缓存会话和档案
5. **扩展功能**: 实现密码重置、2FA 等
