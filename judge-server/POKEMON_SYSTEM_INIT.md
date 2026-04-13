# Pokemon 游戏系统初始化指南

## 📋 概述

本指南帮助您初始化完整的 Agent Monster Pokemon 游戏系统。系统已实现了以下核心功能：

- ✅ 宝可梦养成系统（能力值计算、个体值、努力值、性格）
- ✅ 战斗引擎（回合制、伤害计算、属性克制、优先级）
- ✅ 捕捉系统（捕捉率计算、异常状态、精灵球种类）
- ✅ 地图探索系统（地区、草地、野生宝可梦生成）
- ✅ 孵蛋系统框架

## 🚀 快速开始

### 1. 环境要求

- Go 1.19+
- PostgreSQL 12+
- Python 3.8+ (可选，用于数据分析)

### 2. 项目结构

```
judge-server/
├── internal/
│   ├── handler/
│   │   ├── pokemon_breeding_handler.go    # 养成系统API
│   │   ├── map_exploration_handler.go     # 地图探索API
│   │   └── ...
│   ├── service/
│   │   ├── pokemon_breeding.go            # 养成业务逻辑
│   │   ├── battle_engine_v2.go            # 改进的战斗引擎
│   │   ├── capture_system.go              # 捕捉系统
│   │   ├── map_exploration.go             # 地图探索
│   │   └── ...
│   ├── model/
│   │   ├── pokemon_stats.go               # 宝可梦数据模型
│   │   └── ...
│   └── db/
│       └── pokemon.go                     # 数据库操作
├── INIT_POKEMON_DATA.sql                 # 宝可梦数据导入脚本
└── scripts/
    └── import_pokemon_data.sh             # 自动导入脚本
```

### 3. 数据库设置

#### 3.1 创建数据库

```bash
# 连接到PostgreSQL
psql -U postgres

# 创建数据库
CREATE DATABASE pokemon_game;
\c pokemon_game

# 创建必要的表
```

#### 3.2 导入Pokemon数据

```bash
# 方法1：使用自动化脚本
cd /root/petskill/judge-server
chmod +x scripts/import_pokemon_data.sh
./scripts/import_pokemon_data.sh pokemon_game postgres localhost 5432

# 方法2：手动导入
psql -U postgres -d pokemon_game < INIT_POKEMON_DATA.sql

# 方法3：从Python生成最新数据
python3 /root/petskill/pokemon_analyzer.py
```

#### 3.3 验证数据导入

```bash
# 连接到数据库
psql -U postgres -d pokemon_game

# 查询宝可梦数据
SELECT COUNT(*) as pokemon_count FROM pokemon_species;

# 查询招式数据
SELECT COUNT(*) as move_count FROM pokemon_moves;

# 查询特性数据
SELECT COUNT(*) as ability_count FROM pokemon_abilities;

# 查询示例数据
SELECT * FROM pokemon_species LIMIT 5;
SELECT * FROM pokemon_moves WHERE name_en = 'Thunderbolt';
```

## 📊 数据统计

从 pokemon-dataset-zh 导入的数据：

- **Pokemon 物种**: 1,025 条
- **招式数据**: 953 条
- **特性数据**: 311 条
- **总计**: 2,289 条 INSERT 语句
- **文件大小**: 358 KB

### 数据来源

- GitHub: https://github.com/42arch/pokemon-dataset-zh
- 包含完整的中文宝可梦数据和官方属性信息

## 🔧 API 端点

### 宝可梦养成 API

#### 增加努力值
```bash
POST /api/pokemon/effort-values
Content-Type: application/json

{
  "user_pokemon_id": 123,
  "hp_effort": 10,
  "atk_effort": 0,
  "def_effort": 0,
  "spa_effort": 0,
  "spd_effort": 0,
  "spe_effort": 0
}
```

#### 改变性格
```bash
POST /api/pokemon/nature
Content-Type: application/json

{
  "user_pokemon_id": 123,
  "nature_id": "hardy"
}
```

#### 获取努力值
```bash
GET /api/pokemon/effort-values?user_pokemon_id=123
```

#### 计算能力值
```bash
POST /api/pokemon/calculate-stats
Content-Type: application/json

{
  "pokemon_species_id": "charizard",
  "level": 50,
  "base_stats": {
    "hp": 78,
    "attack": 84,
    "defense": 78,
    "sp_atk": 109,
    "sp_def": 85,
    "speed": 100
  },
  "individual_stats": {
    "pokemon_id": "charizard",
    "iv_hp": 31,
    "iv_attack": 31,
    "iv_defense": 31,
    "iv_sp_atk": 31,
    "iv_sp_def": 31,
    "iv_speed": 31
  },
  "effort_stats": {
    "id": 1,
    "user_pokemon_id": 123,
    "ev_hp": 0,
    "ev_attack": 0,
    "ev_defense": 0,
    "ev_sp_atk": 252,
    "ev_sp_def": 4,
    "ev_speed": 252
  },
  "nature": {
    "id": 1,
    "nature_id": "timid",
    "name_en": "Timid",
    "name_zh": "胆小",
    "increased_stat": "speed",
    "decreased_stat": "attack"
  }
}
```

#### 尝试捕捉
```bash
POST /api/pokemon/capture
Content-Type: application/json

{
  "user_id": 1,
  "wild_pokemon_id": "pikachu",
  "pokemon_level": 5,
  "current_hp": 10,
  "max_hp": 35,
  "status_condition": "none",
  "ball_type": "pokeball"
}
```

### 战斗 API

#### 计算伤害
```bash
POST /api/battle/calculate-damage-v2
Content-Type: application/json

{
  "attacker_pokemon": { /* BattlePokemon */ },
  "defender_pokemon": { /* BattlePokemon */ },
  "move": { /* PokemonMove */ },
  "weather": "rain",
  "is_critical": false
}
```

#### 获取属性克制
```bash
POST /api/battle/type-matchup
Content-Type: application/json

{
  "attacker_type": "fire",
  "defender_type": "grass"
}
```

### 地图探索 API

#### 进入草地
```bash
POST /api/map/enter-grass
Content-Type: application/json

{
  "area_id": "viridian-forest",
  "user_id": 1
}
```

#### 离开草地
```bash
POST /api/map/exit-grass
Content-Type: application/json

{
  "area_id": "viridian-forest",
  "user_id": 1
}
```

#### 获取所有地区
```bash
GET /api/map/regions
```

#### 获取地区的草地
```bash
GET /api/map/grass-areas?region_id=kanto
```

#### 获取地图状态
```bash
GET /api/map/status
```

## 🎯 核心系统说明

### 1. 能力值计算系统

计算公式（HP）：
```
HP = ((2 * BaseHP + IVHP + EVHP/4) * Level) / 100 + Level + 5
```

计算公式（其他能力值）：
```
Stat = (((2 * BaseStat + IVStat + EVStat/4) * Level) / 100 + 5) * NatureMultiplier
```

其中 NatureMultiplier = 1.1（增加）、1.0（中立）、0.9（减少）

### 2. 战斗伤害计算

```
Damage = (((((2 * AttackStat / 5 + 2) * Power * DefenseStat / 50) / 50) + 2) * Effectiveness * Critical * Weather * Random) % 255
```

### 3. 捕捉概率计算

```
CatchRate = BaseRate * HPMultiplier * StatusMultiplier * BallMultiplier
Success = CatchRate > Random(0-255)
```

### 4. 属性克制系统

支持18种属性的完整克制表：
- 普通、火、水、草、电、冰、格斗、毒、地、飞、超能、虫、岩、幽灵、龙、恶、钢、仙

## 📝 数据库表结构

### pokemon_species
```sql
CREATE TABLE pokemon_species (
    pokemon_id VARCHAR(10) PRIMARY KEY,
    name_en VARCHAR(255) NOT NULL,
    name_zh VARCHAR(255) NOT NULL,
    types TEXT,
    capture_rate INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### pokemon_moves
```sql
CREATE TABLE pokemon_moves (
    move_id VARCHAR(50) PRIMARY KEY,
    name_en VARCHAR(255) NOT NULL,
    name_zh VARCHAR(255) NOT NULL,
    type VARCHAR(50),
    category VARCHAR(20),
    power INT,
    accuracy INT,
    pp INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### pokemon_abilities
```sql
CREATE TABLE pokemon_abilities (
    ability_id VARCHAR(50) PRIMARY KEY,
    name_en VARCHAR(255) NOT NULL,
    name_zh VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### pokemon_effort_stats
```sql
CREATE TABLE pokemon_effort_stats (
    id SERIAL PRIMARY KEY,
    user_pokemon_id INT NOT NULL,
    ev_hp INT DEFAULT 0,
    ev_attack INT DEFAULT 0,
    ev_defense INT DEFAULT 0,
    ev_sp_atk INT DEFAULT 0,
    ev_sp_def INT DEFAULT 0,
    ev_speed INT DEFAULT 0,
    total_ev INT DEFAULT 0,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### pokemon_natures
```sql
CREATE TABLE pokemon_natures (
    id SERIAL PRIMARY KEY,
    nature_id VARCHAR(50) NOT NULL UNIQUE,
    name_en VARCHAR(255) NOT NULL,
    name_zh VARCHAR(255),
    increased_stat VARCHAR(50),
    decreased_stat VARCHAR(50),
    favorite_flavor VARCHAR(50),
    disliked_flavor VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🧪 集成测试

```bash
# 运行所有测试
cd /root/petskill/judge-server
go test ./...

# 运行特定测试
go test ./internal/service -v

# 测试API端点
bash test_api.sh
```

## 📚 文件清单

| 文件 | 行数 | 描述 |
|------|------|------|
| pokemon_breeding.go | 418 | 宝可梦养成服务 |
| battle_engine_v2.go | 376 | 改进的战斗引擎 |
| capture_system.go | 321 | 捕捉系统 |
| map_exploration.go | 296 | 地图探索系统 |
| pokemon_stats.go | 214 | 数据模型 |
| pokemon_breeding_handler.go | 280+ | 养成API处理器 |
| map_exploration_handler.go | 240+ | 地图探索API处理器 |
| INIT_POKEMON_DATA.sql | 4,617 | 宝可梦数据初始化 |
| pokemon_analyzer.py | 408 | 数据分析器 |

**总计：2,500+ 行Go代码 + 4,600+ 行SQL + 400 行Python**

## 🔐 安全性

- ✅ SQL注入防护（参数化查询）
- ✅ 数据验证（范围检查、类型检查）
- ✅ 错误处理（完整的错误返回）
- ✅ 日志记录（所有重要操作）

## 🎓 学习资源

- [Pokemon 官方数据](https://bulbapedia.bulbagarden.net/)
- [Pokemon 数据集](https://github.com/42arch/pokemon-dataset-zh)
- [Judge Server 文档](./README.md)
- [认证系统文档](./AUTHENTICATION_SYSTEM.md)

## 📞 支持

遇到问题？检查以下内容：

1. PostgreSQL 是否正确安装并运行
2. 数据库连接字符串是否正确
3. SQL文件是否完整导入
4. Go 应用是否能访问数据库
5. 日志中是否有错误信息

## 📄 许可证

本项目遵循 MIT 许可证。数据来自 pokemon-dataset-zh，遵循其许可证。

---

**最后更新**: 2026-04-13  
**版本**: 1.0.0  
**状态**: 生产就绪 ✅
