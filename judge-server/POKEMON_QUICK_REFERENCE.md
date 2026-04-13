# Pokemon 系统快速参考指南

## 🚀 快速开始

### 编译和测试

```bash
# 进入项目目录
cd /root/petskill/judge-server

# 编译检查
go build ./...

# 运行所有测试
go test ./internal/handler -v
```

## 📋 新注册的API路由

### 养成系统 (4个端点)
- `POST /api/pokemon/effort-values` - 添加努力值
- `POST /api/pokemon/nature` - 改变性格
- `POST /api/pokemon/calculate-stats` - 计算属性
- `POST /api/pokemon/capture` - 捕捉宝可梦

### 地图探索系统 (6个端点)
- `POST /api/map/enter-grass` - 进入草地
- `POST /api/map/exit-grass` - 离开草地
- `GET /api/map/regions` - 获取地区
- `GET /api/map/grass-areas?region_id=...` - 获取草地
- `GET /api/map/exploration-log?user_id=...` - 获取探索日志
- `GET /api/map/status` - 获取地图状态

## 🧪 测试命令

```bash
# 测试路由注册
go test ./internal/handler -run TestRouteRegistration -v

# 测试宝可梦养成
go test ./internal/handler -run TestPokemonBreedingHandlers -v

# 测试地图探索
go test ./internal/handler -run TestMapExplorationHandlers -v

# 端到端测试：捕捉流程
go test ./internal/handler -run TestEndToEndPokemonCatch -v

# 端到端测试：养成流程
go test ./internal/handler -run TestEndToEndPokemonBreeding -v

# 多账户测试
go test ./internal/handler -run TestMultipleAccounts -v

# 错误处理测试
go test ./internal/handler -run TestErrorHandling -v

# 数据导入测试
go test ./internal/handler -run TestPokemonDataImport -v
```

## ✅ 完成的工作

### 路由注册
- ✅ 所有10个新端点已在 `cmd/main.go` 中注册
- ✅ 处理程序方法已在 `internal/handler/*.go` 中实现

### 集成测试
- ✅ `pokemon_integration_test.go` 创建（550+行代码）
- ✅ 8个完整的测试用例
- ✅ TestHelper 辅助类用于快速测试

### 错误处理改进
- ✅ 添加了详细的输入验证
- ✅ 所有响应都包含错误代码
- ✅ 努力值上限验证 (≤510)
- ✅ 等级范围验证 (1-100)
- ✅ HP值范围验证 (0-maxHP)
- ✅ 精灵球类型验证
- ✅ 用户ID和宝可梦ID验证

### 测试数据
- ✅ 可创建4个不同账户
- ✅ 每个账户可创建多只宝可梦
- ✅ 性格和努力值数据创建

## 📊 代码统计

```
新增文件：
- pokemon_integration_test.go    (550+ 行)
- POKEMON_INTEGRATION_TEST_GUIDE.md (400+ 行)

修改文件：
- cmd/main.go                   (+12 行 路由注册)
- pokemon_breeding_handler.go   (+60 行 验证改进)
- map_exploration_handler.go    (+50 行 验证改进)
- internal/db/database.go       (+10 行 查询方法)

总计：1000+ 行新增或改进代码
```

## 🔍 验证规则

### 努力值 (Effort Values)
- 最小: 0
- 最大单个值: 252
- 总计: ≤ 510

### 等级 (Level)
- 最小: 1
- 最大: 100

### 用户ID / 宝可梦ID
- 必须: 正整数 > 0

### 精灵球类型
- 允许值: `pokeball`, `greatball`, `ultraball`, `masterball`

### HP值
- 当前HP: 0 ≤ current_hp ≤ max_hp
- 最大HP: max_hp > 0

## 🛠️ 主要改进

### 1. 路由注册完整性
从 0 个 → 10 个新端点全部注册

### 2. 错误处理详细化
- 添加了唯一的错误代码
- 提供了具体的验证失败原因
- 区分不同的错误类型

### 3. 输入验证强化
- 非空验证
- 范围验证
- 类型验证
- 业务规则验证（如努力值上限）

### 4. 数据库操作扩展
- 添加了 QueryRow() 方法
- 添加了 Query() 方法
- 支持更灵活的查询操作

## 📁 关键文件

```
/root/petskill/judge-server/
├── cmd/main.go 
│   └── 第258-275行: 新的Pokemon系统路由注册
├── internal/handler/
│   ├── pokemon_integration_test.go   [新文件] 测试
│   ├── pokemon_breeding_handler.go   [改进] 完整验证
│   └── map_exploration_handler.go    [改进] 完整验证
├── internal/db/
│   └── database.go                   [改进] 添加查询方法
└── POKEMON_INTEGRATION_TEST_GUIDE.md [新文件] 文档
```

## 🎯 下一步建议

1. **数据库配置**
   ```bash
   # 运行测试数据库初始化
   createdb agent_monster_test
   ```

2. **运行完整测试**
   ```bash
   go test ./internal/handler -v -timeout 120s
   ```

3. **导入Pokemon数据**
   ```bash
   psql agent_monster < INIT_POKEMON_DATA.sql
   ```

4. **性能测试**
   ```bash
   go test -bench=. -benchmem ./internal/handler
   ```

## 💡 使用示例

### 创建和养成宝可梦

```bash
# 1. 添加努力值
curl -X POST http://localhost:8080/api/pokemon/effort-values \
  -H "Content-Type: application/json" \
  -d '{
    "user_pokemon_id": 1,
    "hp_effort": 4,
    "atk_effort": 252,
    "def_effort": 0,
    "spa_effort": 0,
    "spd_effort": 0,
    "spe_effort": 252
  }'

# 2. 改变性格
curl -X POST http://localhost:8080/api/pokemon/nature \
  -H "Content-Type: application/json" \
  -d '{
    "user_pokemon_id": 1,
    "nature_id": "jolly"
  }'

# 3. 计算属性
curl -X POST http://localhost:8080/api/pokemon/calculate-stats \
  -H "Content-Type: application/json" \
  -d '{
    "user_pokemon_id": 1,
    "level": 50
  }'
```

### 地图探索流程

```bash
# 1. 获取地区列表
curl http://localhost:8080/api/map/regions

# 2. 获取特定地区的草地
curl "http://localhost:8080/api/map/grass-areas?region_id=region_001"

# 3. 进入草地
curl -X POST http://localhost:8080/api/map/enter-grass \
  -H "Content-Type: application/json" \
  -d '{
    "area_id": "grass_001",
    "user_id": 1
  }'

# 4. 尝试捕捉
curl -X POST http://localhost:8080/api/pokemon/capture \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "wild_pokemon_id": "25",
    "pokemon_level": 10,
    "current_hp": 5,
    "max_hp": 20,
    "status_condition": "none",
    "ball_type": "pokeball"
  }'

# 5. 离开草地
curl -X POST http://localhost:8080/api/map/exit-grass \
  -H "Content-Type: application/json" \
  -d '{
    "area_id": "grass_001",
    "user_id": 1
  }'
```

## 📞 支持

详细文档请查看：
- `POKEMON_INTEGRATION_TEST_GUIDE.md` - 完整的API和测试文档
- `internal/handler/pokemon_integration_test.go` - 测试实现
- `cmd/main.go` - 路由注册代码
