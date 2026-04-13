# Pokemon 系统集成测试与API文档

## 目录

1. [快速开始](#快速开始)
2. [路由注册](#路由注册)
3. [API端点文档](#api端点文档)
4. [集成测试](#集成测试)
5. [测试数据](#测试数据)
6. [错误处理](#错误处理)

---

## 快速开始

### 编译项目

```bash
cd /root/petskill/judge-server
go build ./...
```

### 运行测试

```bash
# 运行所有集成测试
go test ./internal/handler -v

# 运行特定测试
go test ./internal/handler -run TestRouteRegistration -v
go test ./internal/handler -run TestEndToEndPokemonCatch -v
```

### 启动服务器

```bash
# 确保数据库已配置
export CONFIG_PATH=".config/config.yaml"
./judge-server
```

---

## 路由注册

### 新注册的Pokemon系统路由

以下路由已在 `cmd/main.go` 中注册：

#### 宝可梦养成系统 (Breeding System)

| 方法 | 路由 | 描述 |
|------|------|------|
| POST | `/api/pokemon/effort-values` | 添加宝可梦努力值 |
| POST | `/api/pokemon/nature` | 改变宝可梦性格 |
| POST | `/api/pokemon/calculate-stats` | 计算宝可梦能力值 |
| POST | `/api/pokemon/capture` | 尝试捕捉宝可梦 |

#### 地图探索系统 (Map Exploration System)

| 方法 | 路由 | 描述 |
|------|------|------|
| POST | `/api/map/enter-grass` | 进入草地区域 |
| POST | `/api/map/exit-grass` | 离开草地区域 |
| GET | `/api/map/regions` | 获取所有地区 |
| GET | `/api/map/grass-areas` | 获取特定地区的草地 |
| GET | `/api/map/exploration-log` | 获取用户的探索日志 |
| GET | `/api/map/status` | 获取地图系统状态 |

---

## API端点文档

### 养成系统 API

#### 1. 添加努力值

**请求：**
```
POST /api/pokemon/effort-values
Content-Type: application/json

{
  "user_pokemon_id": 123,
  "hp_effort": 4,
  "atk_effort": 252,
  "def_effort": 0,
  "spa_effort": 0,
  "spd_effort": 0,
  "spe_effort": 252
}
```

**响应成功：** (HTTP 200)
```json
{
  "success": true,
  "data": {
    "user_pokemon_id": 123,
    "hp_effort": 4,
    "atk_effort": 252,
    "def_effort": 0,
    "spa_effort": 0,
    "spd_effort": 0,
    "spe_effort": 252,
    "total_effort": 508
  }
}
```

**错误响应：** (HTTP 400)
```json
{
  "success": false,
  "error": "Total effort values cannot exceed 510",
  "code": "EFFORT_LIMIT_EXCEEDED"
}
```

**验证规则：**
- `user_pokemon_id`: 必须为正整数
- 所有努力值: 必须 >= 0
- 总努力值: 最多510

---

#### 2. 改变性格

**请求：**
```
POST /api/pokemon/nature
Content-Type: application/json

{
  "user_pokemon_id": 123,
  "nature_id": "jolly"
}
```

**响应成功：** (HTTP 200)
```json
{
  "success": true,
  "data": {
    "id": "jolly",
    "name": "Jolly",
    "increased_stat": "spe",
    "decreased_stat": "spa"
  }
}
```

**错误响应：** (HTTP 404)
```json
{
  "success": false,
  "error": "Nature not found",
  "code": "NATURE_NOT_FOUND"
}
```

**验证规则：**
- `user_pokemon_id`: 必须为正整数
- `nature_id`: 不能为空，必须存在于数据库

---

#### 3. 计算能力值

**请求：**
```
POST /api/pokemon/calculate-stats
Content-Type: application/json

{
  "user_pokemon_id": 123,
  "level": 50,
  "base_stats": {
    "hp": 45,
    "attack": 49,
    "defense": 49,
    "sp_atk": 65,
    "sp_def": 65,
    "speed": 45
  },
  "individual_stats": {
    "hp": 31,
    "attack": 31,
    "defense": 31,
    "sp_atk": 31,
    "sp_def": 31,
    "speed": 31
  },
  "effort_stats": {
    "hp": 0,
    "attack": 252,
    "defense": 0,
    "sp_atk": 0,
    "sp_def": 0,
    "speed": 252
  },
  "nature": {
    "id": "jolly",
    "increased_stat": "spe",
    "decreased_stat": "spa"
  }
}
```

**响应成功：** (HTTP 200)
```json
{
  "success": true,
  "data": {
    "user_pokemon_id": 123,
    "level": 50,
    "stats": {
      "hp": 71,
      "attack": 116,
      "defense": 65,
      "sp_atk": 75,
      "sp_def": 85,
      "speed": 121
    }
  }
}
```

**验证规则：**
- `level`: 必须在 1-100 之间
- `user_pokemon_id`: 必须为正整数

---

#### 4. 尝试捕捉宝可梦

**请求：**
```
POST /api/pokemon/capture
Content-Type: application/json

{
  "user_id": 1,
  "wild_pokemon_id": "1",
  "pokemon_level": 10,
  "current_hp": 5,
  "max_hp": 20,
  "status_condition": "none",
  "ball_type": "pokeball"
}
```

**响应成功：** (HTTP 200)
```json
{
  "success": true,
  "captured": true,
  "capture_rate": 85,
  "message": "Pokemon captured successfully!"
}
```

**错误响应：** (HTTP 400)
```json
{
  "success": false,
  "error": "Invalid ball_type (must be: pokeball, greatball, ultraball, masterball)",
  "code": "INVALID_BALL_TYPE"
}
```

**验证规则：**
- `user_id`: 必须为正整数
- `wild_pokemon_id`: 不能为空
- `pokemon_level`: 必须在 1-100 之间
- `current_hp`: 必须在 0 到 `max_hp` 之间
- `max_hp`: 必须为正整数
- `ball_type`: 必须为 "pokeball", "greatball", "ultraball" 或 "masterball"

---

### 地图探索系统 API

#### 1. 进入草地区域

**请求：**
```
POST /api/map/enter-grass
Content-Type: application/json

{
  "area_id": "test_grass_1",
  "user_id": 1
}
```

**响应成功：** (HTTP 200)
```json
{
  "success": true,
  "encounter": {
    "pokemon_id": "25",
    "pokemon_name": "Pikachu",
    "level": 15
  },
  "data": {
    "encounter": {...}
  }
}
```

**验证规则：**
- `area_id`: 不能为空
- `user_id`: 必须为正整数

---

#### 2. 离开草地区域

**请求：**
```
POST /api/map/exit-grass
Content-Type: application/json

{
  "area_id": "test_grass_1",
  "user_id": 1
}
```

**响应成功：** (HTTP 200)
```json
{
  "success": true,
  "message": "Successfully exited grass area",
  "code": "EXIT_SUCCESS"
}
```

---

#### 3. 获取所有地区

**请求：**
```
GET /api/map/regions
```

**响应成功：** (HTTP 200)
```json
{
  "success": true,
  "data": [
    {
      "id": "region_001",
      "name_en": "Kanto",
      "name_zh": "关都地区",
      "description": "The starting region",
      "level_range_min": 1,
      "level_range_max": 25
    },
    {
      "id": "region_002",
      "name_en": "Johto",
      "name_zh": "城都地区",
      "description": "The second region",
      "level_range_min": 15,
      "level_range_max": 40
    }
  ],
  "code": "GET_REGIONS_SUCCESS"
}
```

---

#### 4. 获取地区的草地

**请求：**
```
GET /api/map/grass-areas?region_id=region_001
```

**响应成功：** (HTTP 200)
```json
{
  "success": true,
  "data": [
    {
      "id": "grass_001",
      "region_id": "region_001",
      "name_en": "Viridian Forest",
      "name_zh": "常青森林",
      "x_coord": 10,
      "y_coord": 20,
      "difficulty": 2
    }
  ],
  "code": "GET_GRASS_AREAS_SUCCESS"
}
```

---

#### 5. 获取探索日志

**请求：**
```
GET /api/map/exploration-log?user_id=1
```

**响应成功：** (HTTP 200)
```json
{
  "success": true,
  "user_id": 1,
  "data": [],
  "code": "GET_LOG_SUCCESS"
}
```

---

#### 6. 获取地图系统状态

**请求：**
```
GET /api/map/status
```

**响应成功：** (HTTP 200)
```json
{
  "success": true,
  "data": {
    "total_regions": 8,
    "total_grass_areas": 42
  },
  "code": "GET_STATUS_SUCCESS"
}
```

---

## 集成测试

### 测试文件位置

`internal/handler/pokemon_integration_test.go`

### 可用的测试

#### 1. 路由注册测试 `TestRouteRegistration`

验证所有新的处理程序路由是否正确注册。

```bash
go test ./internal/handler -run TestRouteRegistration -v
```

#### 2. 宝可梦养成处理程序测试 `TestPokemonBreedingHandlers`

测试努力值、性格和属性计算。

```bash
go test ./internal/handler -run TestPokemonBreedingHandlers -v
```

#### 3. 地图探索处理程序测试 `TestMapExplorationHandlers`

测试进入/离开草地、获取地区等功能。

```bash
go test ./internal/handler -run TestMapExplorationHandlers -v
```

#### 4. 端到端宝可梦捕捉测试 `TestEndToEndPokemonCatch`

测试完整的捕捉流程：
- 创建用户账户
- 进入草地
- 尝试捕捉
- 离开草地

```bash
go test ./internal/handler -run TestEndToEndPokemonCatch -v
```

#### 5. 端到端宝可梦养成测试 `TestEndToEndPokemonBreeding`

测试完整的养成流程：
- 创建用户账户
- 创建宝可梦
- 添加努力值
- 改变性格
- 计算属性

```bash
go test ./internal/handler -run TestEndToEndPokemonBreeding -v
```

#### 6. 多账户测试 `TestMultipleAccounts`

创建4个不同的测试账户，每个账户包含3只宝可梦。

```bash
go test ./internal/handler -run TestMultipleAccounts -v
```

#### 7. 错误处理测试 `TestErrorHandling`

验证所有错误处理和输入验证。

```bash
go test ./internal/handler -run TestErrorHandling -v
```

#### 8. 宝可梦数据导入测试 `TestPokemonDataImport`

验证宝可梦数据导入流程和数据库表。

```bash
go test ./internal/handler -run TestPokemonDataImport -v
```

### 运行所有测试

```bash
go test ./internal/handler -v
```

### 测试辅助函数

测试文件提供了以下辅助函数：

- `NewTestHelper()` - 创建测试助手
- `CreateTestAccount()` - 创建测试账户
- `CreateTestPokemon()` - 创建测试宝可梦
- `CreateTestNature()` - 创建测试性格
- `Request()` - 发送HTTP请求
- `AssertStatus()` - 验证HTTP状态码
- `AssertSuccess()` - 验证响应成功标志

---

## 测试数据

### 测试账户结构

```go
type TestAccount struct {
    Username string
    Email    string
    UserID   int
    Token    string
}
```

### 创建的测试账户

测试中会创建4个不同的账户：

1. **player_1** - player1@test.com
2. **player_2** - player2@test.com
3. **player_3** - player3@test.com
4. **player_4** - player4@test.com

每个账户会创建3只宝可梦。

### 测试数据库配置

```yaml
Host: localhost
Port: 5432
User: postgres
Password: xiaodudu
DBName: agent_monster_test
SSLMode: disable
```

---

## 错误处理

### HTTP 状态码

| 状态码 | 含义 |
|-------|------|
| 200 | 成功 |
| 400 | 错误的请求（验证失败） |
| 404 | 未找到 |
| 405 | 方法不允许 |
| 500 | 内部服务器错误 |

### 错误代码 (Error Codes)

#### 通用错误

| 代码 | 含义 |
|-----|------|
| METHOD_NOT_ALLOWED | HTTP方法不允许 |
| INVALID_REQUEST | 无效的请求体 |
| INVALID_POKEMON_ID | 无效的宝可梦ID |
| INVALID_USER_ID | 无效的用户ID |
| INVALID_LEVEL | 无效的等级 |

#### 养成系统错误

| 代码 | 含义 |
|-----|------|
| INVALID_EFFORT_VALUES | 无效的努力值 |
| EFFORT_LIMIT_EXCEEDED | 努力值超过上限（最多510） |
| INVALID_NATURE_ID | 无效的性格ID |
| NATURE_NOT_FOUND | 性格未找到 |
| ADD_EFFORT_FAILED | 添加努力值失败 |
| RETRIEVE_STATS_FAILED | 获取属性失败 |

#### 捕捉系统错误

| 代码 | 含义 |
|-----|------|
| INVALID_BALL_TYPE | 无效的精灵球类型 |
| INVALID_HP | 无效的HP值 |
| POKEMON_NOT_FOUND | 宝可梦未找到 |
| CAPTURE_FAILED | 捕捉失败 |

#### 地图探索错误

| 代码 | 含义 |
|-----|------|
| INVALID_AREA_ID | 无效的区域ID |
| INVALID_REGION_ID | 无效的地区ID |

### 错误响应格式

所有错误响应都包含以下字段：

```json
{
  "success": false,
  "error": "错误描述",
  "code": "ERROR_CODE"
}
```

---

## 快速测试命令

```bash
# 编译整个项目
cd /root/petskill/judge-server && go build ./...

# 运行所有集成测试
go test ./internal/handler -v -timeout 60s

# 运行特定测试套件
go test ./internal/handler -run TestEndToEndPokemonBreeding -v

# 运行测试并显示覆盖率
go test ./internal/handler -v -cover

# 运行测试并生成覆盖率报告
go test ./internal/handler -v -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## 项目结构

```
/root/petskill/judge-server/
├── cmd/
│   └── main.go                    # 主程序，包含路由注册
├── internal/
│   ├── db/
│   │   └── database.go            # 数据库连接和查询方法
│   ├── handler/
│   │   ├── pokemon_breeding_handler.go      # 养成系统处理程序
│   │   ├── map_exploration_handler.go       # 地图探索处理程序
│   │   └── pokemon_integration_test.go      # 集成测试
│   ├── service/
│   │   ├── pokemon_breeding.go             # 养成系统服务
│   │   ├── capture_system.go               # 捕捉系统服务
│   │   └── map_exploration.go              # 地图探索服务
│   └── model/
│       └── pokemon_stats.go                # 宝可梦数据模型
└── INIT_POKEMON_DATA.sql          # Pokemon官方数据导入脚本
```

---

## 状态摘要

✅ **已完成：**
- ✅ 所有处理程序已在路由中注册
- ✅ 创建了全面的集成测试框架
- ✅ 实现了8个集成测试用例
- ✅ 改进了所有处理程序的错误处理和输入验证
- ✅ 创建了测试数据工厂函数
- ✅ 编写了完整的API文档

⚠️ **待完成：**
- 在测试数据库上运行完整的集成测试
- 导入Pokemon官方数据
- 性能测试和优化

---

## 联系与支持

如有问题或需要进一步说明，请参考：
- 测试文件：`/root/petskill/judge-server/internal/handler/pokemon_integration_test.go`
- API处理程序：`/root/petskill/judge-server/internal/handler/pokemon_breeding_handler.go`
- API处理程序：`/root/petskill/judge-server/internal/handler/map_exploration_handler.go`
