# 🚀 多客户端认证系统 - 快速开始指南

## 概览

新的认证系统脱离了对 GitHub CLI 的依赖，支持多个 AI 代理客户端（Claude、OpenCode、OpenClaw、Gemini 等）同时访问同一账户，同时提供完整的安全性和会话管理。

## 三种认证方式

### 1️⃣ 测试账户认证（最简单）
用于开发和测试，无需 GitHub 账户。

```bash
# 创建测试账户
curl -X POST http://localhost:8080/api/auth/register-test \
  -H "Content-Type: application/json" \
  -d '{
    "username": "test_player_1",
    "password": "password123"
  }'

# 返回
{
  "success": true,
  "player_id": "player_xyz123...",
  "server_token": "abc123...xyz789",
  "access_token": "token_xyz...",
  "username": "test_player_1"
}
```

### 2️⃣ Server Token 认证（快速登录）
对于已有账户的直接访问。

```bash
# 登录
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "server_token": "abc123...xyz789",
    "client_name": "claude"
  }'

# 后续请求
curl -H "Authorization: <new_access_token>" \
  http://localhost:8080/api/user/status
```

### 3️⃣ GitHub 账户绑定（安全登录）
将 GitHub 账户绑定到游戏账户。

```bash
# 第一步：创建 GitHub 绑定令牌（在 GitHub App 中）
curl -X POST http://localhost:8080/api/auth/create-player-token \
  -H "Content-Type: application/json" \
  -d '{
    "github_id": 123456,
    "github_login": "myusername"
  }'

# 返回
{
  "success": true,
  "player_token": "binding_token_xyz...",
  "expires_at": "2026-04-14T12:34:56Z"
}

# 第二步：在另一个客户端中绑定账户（如 Claude）
curl -X POST http://localhost:8080/api/auth/bind-github \
  -H "Content-Type: application/json" \
  -d '{
    "player_token": "binding_token_xyz...",
    "client_name": "claude"
  }'

# 返回新的 access_token
{
  "success": true,
  "player_id": "player_xyz123...",
  "server_token": "abc123...xyz789",
  "access_token": "token_new..."
}
```

## 快速设置

### 步骤 1: 初始化数据库

在 Go 代码中调用：

```go
if err := database.InitAuthSchema(); err != nil {
    log.Fatal(err)
}
```

或在 SQL 中手动执行 `/root/petskill/judge-server/internal/db/auth.go` 中的 SQL 语句。

### 步骤 2: 安装依赖

```bash
cd /root/petskill/judge-server
go get golang.org/x/crypto/bcrypt
```

### 步骤 3: 集成到 main.go

```go
// 创建认证管理器
authManager := service.NewAuthManager(database)
authHandler := handler.NewAuthHandler(authManager)

// 注册路由
mux.HandleFunc("/api/auth/register-test", authHandler.RegisterTestAccount)
mux.HandleFunc("/api/auth/login", authHandler.Login)
mux.HandleFunc("/api/auth/test-accounts", authHandler.ListTestAccounts)
mux.HandleFunc("/api/auth/create-player-token", authHandler.CreatePlayerTokenForGitHub)
mux.HandleFunc("/api/auth/bind-github", authHandler.BindGitHub)
mux.HandleFunc("/api/auth/validate-token", authHandler.ValidateToken)
```

### 步骤 4: 初始化测试账户

#### 方法 A: 使用 Python 管理工具

```bash
# 创建单个测试账户
python3 /root/petskill/judge-server/scripts/test_account_manager.py create \
  --username test_player_1 \
  --password password123

# 列出所有测试账户
python3 /root/petskill/judge-server/scripts/test_account_manager.py list

# 登录测试账户
python3 /root/petskill/judge-server/scripts/test_account_manager.py login \
  --username test_player_1 \
  --password password123 \
  --client-name claude
```

#### 方法 B: 使用 Bash 脚本

```bash
# 初始化 10 个预定义的测试账户
bash /root/petskill/judge-server/scripts/init_test_accounts.sh
```

## 预定义的测试账户

脚本会自动创建以下账户（密码见脚本）：

```
test_player_1
test_player_2
test_player_3
trainer_pikachu
trainer_charizard
trainer_blastoise
claude_player
opencode_player
openclaw_player
gemini_player
```

## 多客户端登录示例

### 场景：同一账户从多个客户端登录

```bash
# 玩家首次注册
curl -X POST http://localhost:8080/api/auth/register-test \
  -d '{"username": "trainer_ash", "password": "pikachu123"}'

# 保存返回的 server_token: "token_abc123..."

# Claude 客户端登录
curl -X POST http://localhost:8080/api/auth/login \
  -d '{
    "server_token": "token_abc123...",
    "client_name": "claude"
  }'

# 返回 Claude 的 access_token: "token_claude_xyz..."

# 同时，OpenCode 客户端也登录
curl -X POST http://localhost:8080/api/auth/login \
  -d '{
    "server_token": "token_abc123...",
    "client_name": "opencode"
  }'

# 返回 OpenCode 的 access_token: "token_opencode_abc..."

# 现在可以从两个客户端同时使用各自的 access_token
# Claude: Authorization: token_claude_xyz...
# OpenCode: Authorization: token_opencode_abc...
```

## 客户端集成示例

### Python 客户端

```python
import requests

class AgentMonsterClient:
    def __init__(self, api_url="http://localhost:8080"):
        self.api_url = api_url
        self.access_token = None
        self.player_id = None
        self.server_token = None
    
    def login_with_test_account(self, username, password, client_name):
        """使用测试账户登录"""
        response = requests.post(
            f"{self.api_url}/api/auth/login",
            json={
                "username": username,
                "password": password,
                "client_name": client_name
            }
        )
        
        if response.json()["success"]:
            data = response.json()
            self.access_token = data["access_token"]
            self.player_id = data["player_id"]
            self.server_token = data["server_token"]
            return True
        return False
    
    def get_player_status(self):
        """获取玩家状态"""
        response = requests.get(
            f"{self.api_url}/api/user/status",
            headers={"Authorization": self.access_token}
        )
        return response.json()

# 使用示例
client = AgentMonsterClient()
if client.login_with_test_account("test_player_1", "password123", "claude"):
    print("✓ 登录成功")
    status = client.get_player_status()
    print(f"玩家状态: {status}")
else:
    print("✗ 登录失败")
```

### cURL 客户端

```bash
#!/bin/bash

API_URL="http://localhost:8080"
USERNAME="test_player_1"
PASSWORD="password123"
CLIENT_NAME="claude"

# 1. 登录
echo "正在登录..."
LOGIN_RESPONSE=$(curl -s -X POST $API_URL/api/auth/login \
  -H "Content-Type: application/json" \
  -d "{
    \"username\": \"$USERNAME\",
    \"password\": \"$PASSWORD\",
    \"client_name\": \"$CLIENT_NAME\"
  }")

# 提取 access_token
ACCESS_TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"access_token":"[^"]*"' | cut -d'"' -f4)

echo "✓ 登录成功"
echo "Access Token: $ACCESS_TOKEN"

# 2. 使用 token 调用 API
echo ""
echo "获取玩家状态..."
curl -s -H "Authorization: $ACCESS_TOKEN" \
  $API_URL/api/user/status | jq .
```

## 安全性检查清单

- [ ] 生产环境使用 HTTPS
- [ ] 配置 CORS 白名单
- [ ] 实现登录速率限制
- [ ] 定期轮换 server_token
- [ ] 监控异常登录活动
- [ ] 实现会话超时机制
- [ ] 记录所有认证日志

## 故障排除

### Q: 如何重置测试账户密码？

目前需要手动通过数据库更新。建议添加"密码重置"端点。

### Q: 如何撤销已绑定的 GitHub 账户？

1. 删除 `player_profiles` 中的 `github_id` 和 `github_login`
2. 用户可以重新绑定新的 GitHub 账户

### Q: Access Token 过期后怎么办？

使用 `server_token` 重新登录以获取新的 access_token（推荐）。

### Q: 如何禁用测试账户？

```bash
python3 test_account_manager.py disable --username test_player_1
```

## 相关文件

| 文件 | 功能 |
|------|------|
| `/root/petskill/judge-server/internal/model/auth.go` | 认证数据模型 |
| `/root/petskill/judge-server/internal/db/auth.go` | 数据库访问层 |
| `/root/petskill/judge-server/internal/service/auth.go` | 认证服务业务逻辑 |
| `/root/petskill/judge-server/internal/handler/auth.go` | HTTP 处理器 |
| `/root/petskill/judge-server/internal/handler/auth_middleware.go` | 认证中间件 |
| `/root/petskill/judge-server/AUTHENTICATION_SYSTEM.md` | 完整文档 |
| `/root/petskill/judge-server/scripts/test_account_manager.py` | 管理工具 |
| `/root/petskill/judge-server/scripts/init_test_accounts.sh` | 初始化脚本 |

## 下一步

- [ ] 集成现有的游戏逻辑与新认证系统
- [ ] 实现 OAuth2 完整流程
- [ ] 添加双因素认证 (2FA)
- [ ] 实现会话管理 UI
- [ ] 添加账户恢复机制
- [ ] 集成 GitHub Webhook 用于同步用户信息
