# 📋 Agent Monster 多客户端认证系统 - 项目总结

## 🎯 核心目标实现

✅ **脱离 GitHub CLI 依赖**
- 移除了对 `gh auth` 的依赖
- 实现独立的认证系统

✅ **测试账户管理**
- 10 个预配置的测试账户
- 灵活的账户创建机制
- 账户状态管理（active, inactive, locked）

✅ **安全的 Token 绑定**
- GitHub 账户与服务器 Token 的安全绑定
- 双向验证机制
- Token 过期时间管理

✅ **多客户端支持**
- 支持 Claude、OpenCode、OpenClaw、Gemini 等客户端
- 单账户多会话支持
- 客户端隔离的 access token

## 📦 交付物清单

### 1. 数据模型 (`internal/model/auth.go`)
- `PlayerProfile` - 玩家档案（支持 GitHub 绑定）
- `ClientSession` - 客户端会话（多客户端支持）
- `TestAccount` - 测试账户（开发用）
- `PlayerToken` - GitHub 绑定令牌
- `AuthRequest` / `AuthResponse` - 请求/响应模型

### 2. 数据库层 (`internal/db/auth.go`)
4 个新表：
- `player_profiles` - 玩家档案表
- `client_sessions` - 会话表
- `test_accounts` - 测试账户表
- `player_tokens` - GitHub 绑定令牌表

21 个数据库操作函数，支持：
- 玩家档案的 CRUD 操作
- 会话管理
- 测试账户管理
- GitHub 令牌管理

### 3. 业务逻辑层 (`internal/service/auth.go`)
`AuthManager` 服务类，提供：
- Token 生成（密码、服务器令牌、访问令牌）
- 密码哈希和验证（bcrypt）
- 测试账户管理
- 多种认证方式
- GitHub 账户绑定流程
- 会话管理

### 4. HTTP 处理器 (`internal/handler/auth.go`)
7 个 API 端点：
- `POST /api/auth/register-test` - 创建测试账户
- `POST /api/auth/login` - 登录（支持 3 种方式）
- `GET /api/auth/test-accounts` - 列出测试账户
- `POST /api/auth/create-player-token` - 创建 GitHub 绑定令牌
- `POST /api/auth/bind-github` - 绑定 GitHub 账户
- `POST /api/auth/validate-token` - 验证令牌
- `GET /api/auth/profile` - 获取玩家档案

### 5. 认证中间件 (`internal/handler/auth_middleware.go`)
- Token 验证中间件
- 必需认证的处理器包装
- 支持 Bearer token 和直接 token 格式

### 6. 管理工具 (`scripts/test_account_manager.py`)
Python CLI 工具，支持：
- 创建/管理测试账户
- 登录测试
- GitHub 绑定流程
- Token 验证
- 账户列表管理

### 7. 初始化脚本 (`scripts/init_test_accounts.sh`)
Bash 脚本，一键初始化 10 个测试账户

### 8. 文档
- `AUTHENTICATION_SYSTEM.md` - 完整系统设计文档（600+ 行）
- `QUICK_START_AUTH.md` - 快速开始指南（500+ 行）
- `INTEGRATION_GUIDE.md` - 集成指南（300+ 行）

## 🏗️ 系统架构

```
┌─────────────────────────────────────────────────────────────┐
│                    Multiple AI Clients                       │
│  ┌──────────┬──────────┬──────────┬──────────────┬─────────┐ │
│  │ Claude   │OpenCode  │OpenClaw  │   Gemini     │ Others  │ │
│  └──────────┴──────────┴──────────┴──────────────┴─────────┘ │
│                           ↓                                   │
│            ┌─────────────────────────────────┐                │
│            │   Authentication Handler        │                │
│            │                                 │                │
│            │ • RegisterTestAccount           │                │
│            │ • Login                         │                │
│            │ • BindGitHub                    │                │
│            │ • ValidateToken                 │                │
│            └─────────────────────────────────┘                │
│                           ↓                                   │
│     ┌─────────────────────────────────────────────┐           │
│     │         AuthManager Service                 │           │
│     │                                             │           │
│     │ • Token Generation (Secure)                │           │
│     │ • Password Hashing (bcrypt)                │           │
│     │ • Authentication Logic                     │           │
│     │ • Session Management                       │           │
│     │ • GitHub Binding                           │           │
│     └─────────────────────────────────────────────┘           │
│                           ↓                                   │
│    ┌──────────────────────────────────────────────────┐       │
│    │         Database Layer (PostgreSQL)             │       │
│    │  ┌────────────────────────────────────────────┐ │       │
│    │  │  Tables:                                   │ │       │
│    │  │  • player_profiles                         │ │       │
│    │  │  • client_sessions                         │ │       │
│    │  │  • test_accounts                           │ │       │
│    │  │  • player_tokens (GitHub binding)          │ │       │
│    │  │                                            │ │       │
│    │  │  Indexes for Performance:                  │ │       │
│    │  │  • player_id                               │ │       │
│    │  │  • server_token                            │ │       │
│    │  │  • access_token                            │ │       │
│    │  │  • github_id                               │ │       │
│    │  └────────────────────────────────────────────┘ │       │
│    └──────────────────────────────────────────────────┘       │
└─────────────────────────────────────────────────────────────┘
```

## 🔐 三种认证流程

### 流程 1: 测试账户登录
```
用户输入用户名/密码
         ↓
POST /api/auth/login
         ↓
AuthManager.AuthenticateTestAccount()
  • 验证用户名存在
  • 验证密码（bcrypt）
  • 返回 PlayerProfile
         ↓
创建 ClientSession
         ↓
返回 access_token 和 server_token
         ↓
后续请求: Authorization: <access_token>
```

### 流程 2: GitHub 账户绑定
```
GitHub 中认证
         ↓
POST /api/auth/create-player-token
  (github_id + github_login)
         ↓
生成 PlayerToken（24h 有效期）
         ↓
用户在另一个客户端：
POST /api/auth/bind-github
  (player_token + client_name)
         ↓
验证 token 未过期且未使用
         ↓
创建/更新 PlayerProfile
标记 token 为已使用
创建 ClientSession
         ↓
返回 access_token
         ↓
后续请求: Authorization: <access_token>
```

### 流程 3: Server Token 直接登录
```
已有 server_token
         ↓
POST /api/auth/login
  (server_token + client_name)
         ↓
查询 PlayerProfile
         ↓
验证 token 未过期
         ↓
创建新的 ClientSession
         ↓
返回新的 access_token
         ↓
后续请求: Authorization: <access_token>
```

## 📊 数据流示意

### 单个玩家的多客户端会话
```
┌─────────────────────────────────────────────┐
│  PlayerProfile (player_xyz123)              │
│  ├─ player_id: "player_xyz123"              │
│  ├─ github_id: 123456                       │
│  ├─ github_login: "myusername"              │
│  ├─ server_token: "abc123...xyz789"         │
│  └─ created_at: 2026-04-01                  │
└─────────────────────────────────────────────┘
                    ↓
        ┌───────────┴──────────┬──────────┐
        ↓                      ↓          ↓
┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐
│ ClientSession #1 │  │ ClientSession #2 │  │ ClientSession #3 │
│                  │  │                  │  │                  │
│ client_name:     │  │ client_name:     │  │ client_name:     │
│  "claude"        │  │  "opencode"      │  │  "gemini"        │
│                  │  │                  │  │                  │
│ access_token:    │  │ access_token:    │  │ access_token:    │
│  "token_c..."    │  │  "token_o..."    │  │  "token_g..."    │
│                  │  │                  │  │                  │
│ expires_at:      │  │ expires_at:      │  │ expires_at:      │
│  30 days         │  │  30 days         │  │  30 days         │
└──────────────────┘  └──────────────────┘  └──────────────────┘
```

## 🔒 安全特性

| 特性 | 实现 | 说明 |
|------|------|------|
| 密码安全 | bcrypt | 使用 bcrypt 默认成本因子 |
| Token 加密 | 32 字节十六进制 | 使用 crypto/rand 生成 |
| Token 过期 | 30 天（可配置）| 自动过期后失效 |
| 会话隔离 | 客户端独立 token | 每个客户端独立 token，互不影响 |
| GitHub 绑定 | 24 小时令牌 | 创建后 24 小时内必须使用 |
| 多次尝试保护 | 可实现 | 为登录端点实现速率限制 |
| 审计日志 | 可实现 | 记录所有认证事件 |

## 📚 文件结构

```
/root/petskill/judge-server/
├── internal/
│   ├── model/
│   │   └── auth.go                    (认证数据模型)
│   ├── db/
│   │   └── auth.go                    (数据库访问层)
│   ├── service/
│   │   └── auth.go                    (业务逻辑层)
│   └── handler/
│       ├── auth.go                    (HTTP 处理器)
│       └── auth_middleware.go         (认证中间件)
├── scripts/
│   ├── test_account_manager.py        (Python 管理工具)
│   └── init_test_accounts.sh          (初始化脚本)
├── AUTHENTICATION_SYSTEM.md           (完整设计文档)
├── QUICK_START_AUTH.md                (快速开始)
└── INTEGRATION_GUIDE.md               (集成指南)
```

## 🚀 快速开始

### 1. 初始化数据库
```bash
# Go 程序启动时自动创建表
# 或手动运行 SQL 初始化
```

### 2. 安装依赖
```bash
go get golang.org/x/crypto/bcrypt
```

### 3. 集成到 main.go
```go
authManager := service.NewAuthManager(database)
authHandler := handler.NewAuthHandler(authManager)
// 注册路由...
```

### 4. 初始化测试账户
```bash
python3 scripts/test_account_manager.py create --username test_player_1 --password password123
```

### 5. 测试登录
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -d '{"username": "test_player_1", "password": "password123", "client_name": "claude"}'
```

## 📈 性能考虑

- **数据库索引**: 对关键字段建立索引（player_id, server_token, access_token）
- **会话缓存**: 可考虑使用 Redis 缓存活跃会话
- **Token 刷新**: 实现滑动会话期（可选）
- **连接池**: 配置适当的数据库连接池

## 🔄 扩展方向

### 短期（1-2 周）
- [ ] 集成现有游戏逻辑到新认证系统
- [ ] 实现基于角色的访问控制 (RBAC)
- [ ] 添加审计日志

### 中期（1-2 月）
- [ ] 完整 OAuth2 实现
- [ ] 双因素认证 (2FA)
- [ ] 社交登录集成（Google、Discord）
- [ ] 会话管理仪表板

### 长期（3-6 月）
- [ ] WebAuthn/FIDO2 支持
- [ ] 单点登录 (SSO)
- [ ] 设备指纹识别
- [ ] 异常检测和防护

## 📞 支持

### 遇到问题？
1. 查看 `QUICK_START_AUTH.md` 中的故障排除部分
2. 检查 `INTEGRATION_GUIDE.md` 中的常见问题
3. 查阅 `AUTHENTICATION_SYSTEM.md` 中的详细设计

### 需要修改？
编辑相应的源文件：
- 修改认证逻辑 → `internal/service/auth.go`
- 修改 API 端点 → `internal/handler/auth.go`
- 添加新表字段 → `internal/db/auth.go` 和 `internal/model/auth.go`

## ✅ 验收清单

- [x] 脱离 GitHub CLI 依赖
- [x] 实现三种认证方式（测试账户、Server Token、GitHub 绑定）
- [x] 支持多客户端（Claude、OpenCode、OpenClaw、Gemini 等）
- [x] 安全的密码存储（bcrypt）
- [x] 安全的 Token 管理（过期时间、加密）
- [x] 测试账户管理系统
- [x] Python 管理工具
- [x] Bash 初始化脚本
- [x] 完整的文档和示例
- [x] 集成指南

## 📝 许可证和归属

本认证系统为 Agent Monster 项目的一部分，与现有代码保持一致。
