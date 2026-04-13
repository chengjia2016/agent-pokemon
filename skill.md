# Agent Monster - Gameplay Skill & Entry Point

Welcome to **Agent Monster**, an AI-powered RPG where your GitHub repository becomes a digital pet. This guide serves as the entry point for players and the operational manual for AI agents.

---

## 🚀 Quick Start: Your First 5 Minutes

**新玩家必读！按以下步骤立即开始游戏：**

### 1️⃣ 检查游戏服务器是否运行
```bash
curl -s http://localhost:10000/health
```
*应该看到*:
```json
{"status": "healthy"}
```

### 2️⃣ 查看你是谁（获取用户ID）
```bash
# 用你的GitHub用户ID替换 'your_github_id'
curl http://localhost:10000/api/users/your_github_id
```

### 3️⃣ 探索游戏世界
```bash
# 查看所有岛屿和城镇
curl http://localhost:10000/api/maps | jq .
```

### 4️⃣ 创建你的基地
```bash
curl -X POST "http://localhost:10000/api/defense/base" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_github_id",
    "base_name": "我的基地",
    "island_id": "1"
  }'
```

### 5️⃣ 开始你的冒险！
- 💬 **与NPC交谈**：获取任务
- 🎯 **完成任务**：获得经验和奖励
- 🏆 **挑战体操馆**：赢得徽章
- 🎮 **探索关卡**：收集资源

---

## 🎮 新功能：完整世界系统 (NEW!)

### ✨ 你现在可以做什么

#### 1. 🤖 与NPC交互
```bash
# 获取城镇中所有NPC
curl "http://localhost:10000/api/npcs?town_id=1"

# 与NPC对话
curl "http://localhost:10000/api/npcs/talk?npc_id=1"
```

#### 2. 📜 接受和完成任务
```bash
# 查看可用任务
curl "http://localhost:10000/api/quests"

# 接受任务
curl -X POST "http://localhost:10000/api/user/quests?user_id=your_id&quest_id=1&action=accept" \
  -H "Content-Type: application/json"

# 完成任务
curl -X POST "http://localhost:10000/api/user/quests?user_id=your_id&quest_id=1&action=complete" \
  -H "Content-Type: application/json"

# 查看你的任务进度
curl "http://localhost:10000/api/user/quests?user_id=your_id"
```

#### 3. 🏰 地下城副本
```bash
# 进入地下城
curl -X POST "http://localhost:10000/api/dungeons" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_id",
    "dungeon_id": 1
  }'

# 查看你的地下城进度
curl "http://localhost:10000/api/user/dungeons?user_id=your_id"
```

#### 4. 🏆 体操馆挑战
```bash
# 查看所有体操馆
curl "http://localhost:10000/api/gyms"

# 查看你赢得的徽章
curl "http://localhost:10000/api/gyms?user_id=your_id"

# 获得体操馆徽章
curl -X POST "http://localhost:10000/api/gyms" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_id",
    "gym_id": 1,
    "badge_name": "岩石勋章"
  }'
```

#### 5. 🗺️ 探索地图关卡
```bash
# 查看地图区域
curl "http://localhost:10000/api/map/zones"

# 获取某个区域的关卡
curl "http://localhost:10000/api/levels?zone_id=1"

# 完成一个关卡
curl -X POST "http://localhost:10000/api/user/levels/progress" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_id",
    "level_id": 1,
    "stars": 3
  }'

# 查看你的关卡进度
curl "http://localhost:10000/api/user/levels?user_id=your_id"
```

---

## 🗺️ 世界系统完整文档

### 📚 什么是世界系统？

世界系统包含5大核心功能：
1. **NPC系统** - 与NPCs交互，接收任务
2. **任务系统** - 完成任务赚取经验和奖励
3. **地下城系统** - 多层副本，Boss战斗
4. **体操馆系统** - 挑战馆主，赢得徽章
5. **地图关卡系统** - 探索不同区域，完成关卡

---

### 🤖 NPC系统

**概念**：NPC是世界中的角色，可以：
- 给你任务
- 出售物品
- 提供信息
- 是地下城的Boss或体操馆馆主

#### 完整API文档

| 操作 | 端点 | 方法 | 说明 |
|------|------|------|------|
| 获取城镇NPCs | `/api/npcs?town_id=1` | GET | 获取某个城镇的所有NPC |
| 创建NPC | `/api/npcs` | POST | 创建新的NPC（管理员） |
| 与NPC对话 | `/api/npcs/talk?npc_id=1` | GET | 获取NPC的对话内容 |

#### 示例：创建和交互NPC
```bash
# 创建新NPC
curl -X POST "http://localhost:10000/api/npcs" \
  -H "Content-Type: application/json" \
  -d '{
    "town_id": "1",
    "name_en": "Nurse Joy",
    "name_zh": "乔伊护士",
    "type": "NPC",
    "role": "healer",
    "coord_x": 50.0,
    "coord_y": 50.0,
    "avatar_url": "https://example.com/nurse.png"
  }'

# 获取城镇的所有NPCs
curl "http://localhost:10000/api/npcs?town_id=1" | jq '.'

# 与NPC交话
curl "http://localhost:10000/api/npcs/talk?npc_id=1" | jq '.'
```

---

### 📜 任务系统

**概念**：任务是游戏的主要进度系统。完成任务获得：
- 💰 金币
- ⭐ 经验值
- 🎁 物品奖励
- 📈 角色进度

#### 完整API文档

| 操作 | 端点 | 方法 | 说明 |
|------|------|------|------|
| 查看所有任务 | `/api/quests` | GET | 查看游戏中所有可用任务 |
| 创建任务 | `/api/quests` | POST | 创建新任务（管理员） |
| 查看我的任务 | `/api/user/quests?user_id=xxx` | GET | 查看我接受的任务 |
| 接受/完成任务 | `/api/user/quests` | POST | 接受或完成任务 |

#### 任务类型

| 类型 | 说明 | 示例 |
|------|------|------|
| `capture` | 捕捉宝可梦 | 捕捉5只不同的宝可梦 |
| `training` | 训练宝可梦 | 将宝可梦训练到10级 |
| `battle` | 战斗任务 | 赢得10场战斗 |
| `collection` | 收集任务 | 收集3个不同的道具 |
| `explore` | 探索任务 | 探索5个不同的区域 |

#### 示例：完整任务流程
```bash
# 第1步：查看所有任务
curl "http://localhost:10000/api/quests" | jq '.quests[] | {id, title_en, description_en, reward_gold}'

# 第2步：接受任务
curl -X POST "http://localhost:10000/api/user/quests" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_id",
    "quest_id": 1,
    "action": "accept"
  }'

# 第3步：查看我的任务进度
curl "http://localhost:10000/api/user/quests?user_id=your_id" | jq '.quests[] | {quest_id, status, progress}'

# 第4步：完成任务
curl -X POST "http://localhost:10000/api/user/quests" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_id",
    "quest_id": 1,
    "action": "complete"
  }'

# 第5步：查看你的进度和奖励
curl "http://localhost:10000/api/users/your_id" | jq '.gold, .exp'
```

---

### 🏰 地下城系统

**概念**：地下城是多层副本，每层有敌人，最后一层有Boss。完成获得重要奖励。

#### 完整API文档

| 操作 | 端点 | 方法 | 说明 |
|------|------|------|------|
| 进入地下城 | `/api/dungeons` | POST | 开始地下城冒险 |
| 查看地下城列表 | `/api/dungeons` | GET | 查看所有可用地下城 |

#### 地下城难度等级

| 难度 | 推荐等级 | 奖励 | 说明 |
|------|---------|------|------|
| `easy` | 1-5 | 100-200金 | 新手副本 |
| `normal` | 5-15 | 200-500金 | 普通难度 |
| `hard` | 15-30 | 500-1000金 | 困难副本 |
| `legendary` | 30+ | 1000+金 | 传奇副本 |

#### 示例：地下城冒险
```bash
# 查看可用地下城
curl "http://localhost:10000/api/dungeons" | jq '.dungeons[] | {id, name_en, difficulty_level, boss_name_en}'

# 进入地下城
curl -X POST "http://localhost:10000/api/dungeons" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_id",
    "dungeon_id": 1
  }'

# 查看地下城进度
curl "http://localhost:10000/api/user/dungeons?user_id=your_id" | jq '.progress[] | {dungeon_id, current_floor, status}'
```

---

### 🏆 体操馆系统

**概念**：体操馆由馆主管理，有固定的宝可梦队伍。击败馆主获得徽章。

#### 完整API文档

| 操作 | 端点 | 方法 | 说明 |
|------|------|------|------|
| 查看体操馆 | `/api/gyms` | GET | 查看所有体操馆 |
| 查看我的徽章 | `/api/gyms?user_id=xxx` | GET | 查看已获得的徽章 |
| 创建体操馆 | `/api/gyms` | POST | 创建新体操馆（管理员） |

#### 体操馆徽章系统

| 徽章 | 城市 | 馆主 | 难度 | 奖励 |
|------|------|------|------|------|
| 岩石勋章 | 常磐市 | 小刚 | ⭐ | 500金 |
| 水晶勋章 | 浅红市 | 小刚 | ⭐⭐ | 750金 |
| 闪电勋章 | 黄金市 | 小鹰 | ⭐⭐ | 750金 |
| ... | ... | ... | ... | ... |

#### 示例：体操馆挑战
```bash
# 查看所有体操馆
curl "http://localhost:10000/api/gyms" | jq '.gyms[] | {id, gym_name_en, gym_leader_id, gym_badge_en}'

# 查看馆主的队伍
curl "http://localhost:10000/api/gyms/1/team" | jq '.team[] | {pokemon_species, pokemon_level}'

# 查看你已获得的徽章
curl "http://localhost:10000/api/gyms?user_id=your_id" | jq '.badges[] | {gym_badge_en, earned_at}'

# 挑战体操馆（通过完成特殊任务）
curl -X POST "http://localhost:10000/api/gyms" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_id",
    "gym_id": 1
  }'
```

---

### 🗺️ 地图关卡系统

**概念**：世界分为多个区域，每个区域有多个关卡。完成关卡解锁新的故事和资源。

#### 完整API文档

| 操作 | 端点 | 方法 | 说明 |
|------|------|------|------|
| 查看地区 | `/api/map/zones` | GET | 查看所有地图区域 |
| 创建地区 | `/api/map/zones` | POST | 创建新地区（管理员） |
| 查看关卡 | `/api/levels?zone_id=1` | GET | 查看某区域的关卡 |
| 创建关卡 | `/api/levels` | POST | 创建新关卡（管理员） |
| 完成关卡 | `/api/user/levels/progress` | POST | 标记关卡为完成 |

#### 地图区域与关卡

```
关都地区 (Region: Kanto)
├── 常青森林 (Zone: Viridian Forest)
│   ├── 关卡 1: 森林入口 (Level 1-3)
│   ├── 关卡 2: 森林深处 (Level 5-7)
│   └── 关卡 3: 森林秘密 (Level 10-12)
├── 月见山 (Zone: Mt. Moon)
│   ├── 关卡 1: 山脚 (Level 10-15)
│   ├── 关卡 2: 山顶 (Level 20-25)
│   └── Boss: 月之精灵
└── 岩洞山 (Zone: Rock Tunnel)
    └── ... 更多关卡
```

#### 示例：探索地图
```bash
# 第1步：查看所有区域
curl "http://localhost:10000/api/map/zones" | jq '.zones[] | {id, zone_name_en, zone_type, difficulty_level}'

# 第2步：查看某个区域的所有关卡
curl "http://localhost:10000/api/levels?zone_id=1" | jq '.levels[] | {id, level_number, level_name_en, difficulty, reward_exp, reward_gold}'

# 第3步：完成一个关卡
curl -X POST "http://localhost:10000/api/user/levels/progress" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_id",
    "level_id": 1,
    "stars": 3
  }'

# 第4步：查看你的关卡进度
curl "http://localhost:10000/api/user/levels?user_id=your_id" | jq '.progress[] | {level_id, stars, completed_at}'

# 第5步：查看草地区域的宝可梦
curl "http://localhost:10000/api/grass_areas?zone_id=1" | jq '.areas[] | {area_name_en, pokemon_list, encounter_rate}'
```

---

## 🎯 新玩家建议流程

### 你的第一个小时应该做什么：

```
⏱️ 0-5分钟：入门
  ✓ 启动游戏服务器
  ✓ 创建你的基地
  ✓ 查看世界地图

⏱️ 5-15分钟：探索NPC和任务
  ✓ 查看城镇中的NPC
  ✓ 与几个NPC交谈
  ✓ 接受1-2个简单任务

⏱️ 15-30分钟：开始冒险
  ✓ 完成你接受的任务
  ✓ 赚取你的第一批金币和经验
  ✓ 探索第一个地图区域

⏱️ 30-45分钟：挑战
  ✓ 进入简单地下城
  ✓ 或挑战第一个体操馆
  ✓ 赢得你的第一个徽章

⏱️ 45-60分钟：深入体验
  ✓ 完成更多任务
  ✓ 升级你的宝可梦
  ✓ 解锁更多地图区域
```

---

## 🛠️ Entry Point: Getting Started

As an AI agent or player, your first task is to ensure the environment is ready.

### 1. Check GitHub Authentication
Before starting, verify if you are logged into GitHub:
```bash
gh auth status
```
*If not logged in, run: `gh auth login`.*

### 2. Check Judge Server Status
Verify that the Judge Server is running:
```bash
curl -s http://localhost:10000/health
```

*Expected response*:
```json
{
  "status": "healthy"
}
```

---

## 🛠️ Quick Start: Play Without Python

You **do not need Python** to play. Use `curl` commands directly with the Judge Server API.

### Step 1: Get Your User ID
```bash
# Check if you have an existing account
curl -s http://localhost:10000/api/users/999
```

### Step 2: Explore the World
```bash
# View all islands and towns
curl -s http://localhost:10000/api/maps | jq .
```

### Step 3: Create Your Base
```bash
# Create a base at your location
curl -X POST "http://localhost:10000/api/defense/base" \
  -H "Content-Type: application/json" \
  -d '{"user_id": "999", "base_name": "My Base"}'
```

### Step 4: Create Your First Pokemon
```bash
# Create a pokemon directly
curl -X POST "http://localhost:10000/api/pokemon/create?user_id=999" \
  -H "Content-Type: application/json" \
  -d '{"species": "Pikachu", "level": 5}'
```

### Step 5: Start Playing
- List your pokemons: `curl http://localhost:10000/api/pokemon?user_id=999`
- Create eggs: `curl -X POST http://localhost:10000/api/pokemon/egg?user_id=999 ...`
- Battle, explore, and have fun!

---

## 🛠️ Onboarding Workflow (Traditional Step-by-Step)

For new players following traditional setup:

1.  **Claim Starter Pack**: Create your account via API
    -   *Reward*: **100 Coins**, **Starter Items**, a **Starter Pet**, and your first **Egg**!
2.  **Test a Battle (Training)**: Use the Battle API against low-level targets
3.  **Establish Your Base**: Use `POST /api/defense/base` to create your base
4.  **Play & Explore**: Use the World System to interact with NPCs and complete quests
    -   Commit your progress to GitHub if desired

---

## 🏗️ Architecture: 100% Server-Authoritative, Zero Local Complexity

Agent Monster uses a **pure server-authoritative model** with NO local Python or complex setup required.

### Server-Stored Data (Source of Truth)
All critical game data is stored in the **Judge Server's PostgreSQL Database**:
- **Player Profiles & Stats**: Level, EXP, Coins, and Items
- **Pet Status**: Egg incubation progress and pet health/stats
- **Battle Results**: All combat outcomes and rewards
- **Inventory**: Items bought or found
- **Bases**: User bases with defense records
- **Map Data**: Islands, towns, and terrain
- **World System**: NPCs, Quests, Dungeons, Gyms, and Levels (NEW!)

*Access via*: **HTTP REST API calls** (no local files needed)

### Client Interaction
- Use `curl` or any HTTP client to call API endpoints
- No Python dependencies
- No local session management
- Completely stateless
- Works from terminal, scripts, or web browsers

---

## 🔐 Login & Persistence

### How to Login
Agent Monster uses **GitHub authentication via the `gh` CLI**:

1. **Check if you're logged in**:
   ```bash
   gh auth status
   ```

2. **If not logged in, authenticate**:
   ```bash
   gh auth login
   ```

3. **Play using your GitHub ID**:
    ```bash
    # Your GitHub ID is your player ID
    curl http://localhost:10000/api/users/999
    ```

### No Local Files Needed
- No `.monster/sessions.json` 
- No `.monster/users/` directory
- No local database
- **Everything is server-side** via the Judge Server API

### Persistent Game State
Your game data is **permanently stored** in the Judge Server's PostgreSQL database:
- Your bases
- Your pokemons and eggs
- Your battle history
- Your inventory
- Your progress
- Your NPC interactions
- Your quests and achievements
- Your dungeon progress
- Your gym badges

Just use your GitHub ID to access your account from anywhere!

---

## 🎮 Gameplay Mechanics (Using curl & HTTP API)

The player can interact directly with the Judge Server using `curl` or any HTTP client. No Python needed!

### Core Player Actions

| Intent | curl Command | Endpoint |
| :--- | :--- | :--- |
| **View Stats** | `curl http://localhost:10000/api/users/{user_id}` | `GET /api/users/{user_id}` |
| **Create Base** | `curl -X POST http://localhost:10000/api/defense/base ...` | `POST /api/defense/base` |
| **Talk to NPC** | `curl http://localhost:10000/api/npcs/talk?npc_id={id}` | `GET /api/npcs/talk` |
| **Accept Quest** | `curl -X POST http://localhost:10000/api/user/quests ...` | `POST /api/user/quests` |
| **Enter Dungeon** | `curl -X POST http://localhost:10000/api/dungeons ...` | `POST /api/dungeons` |
| **Challenge Gym** | `curl -X POST http://localhost:10000/api/gyms ...` | `POST /api/gyms` |
| **Complete Level** | `curl -X POST http://localhost:10000/api/user/levels/progress ...` | `POST /api/user/levels/progress` |
| **Explore World** | `curl http://localhost:10000/api/map/zones` | `GET /api/map/zones` |
| **Create Pokemon** | `curl -X POST http://localhost:10000/api/pokemon/create ...` | `POST /api/pokemon/create` |
| **List My Pokemon** | `curl http://localhost:10000/api/pokemon?user_id={id}` | `GET /api/pokemon` |

---

## 🛠️ Judge Server API Reference (Authoritative)

The Judge Server is the **Source of Truth**. You can interact with it directly via `curl` even if Python scripts are missing.

### User Management
- **Create Account**: `POST /api/users/create`
  - Body: `{"github_id": 12345, "github_login": "username"}`

### Pet & Eggs
- **Claim Egg**: `POST /api/pokemon/egg`
  - Body: `{"user_id": "your_id", "species": "Pikachu"}`
- **Hatch Egg**: `POST /api/pokemon/egg/{eggID}/hatch`

### Battle System
- **Start Battle**: `POST /api/battles/start`
  - Body: `{"user_id": 1, "opponent_id": 2, "battle_type": "duel"}`
- **Record Result**: `POST /api/battles/end`

### World System (NEW!)
- **Get NPCs**: `GET /api/npcs?town_id=1`
- **Get Quests**: `GET /api/quests`
- **Accept Quest**: `POST /api/user/quests?user_id=x&quest_id=x&action=accept`
- **Enter Dungeon**: `POST /api/dungeons`
- **Get Gyms**: `GET /api/gyms`
- **Get Levels**: `GET /api/levels?zone_id=1`
- **Complete Level**: `POST /api/user/levels/progress`

---

## 🗺️ Legacy Features: Base System, Map System & Pokemon Creation

### 🏰 Base Management System
Build and defend your own Pokemon base in the game world!

#### API Endpoints
| Action | Endpoint | Method | Description |
|--------|----------|--------|-------------|
| **Create Base** | `/api/defense/base` | POST | Create a new base for a user |
| **Get Base** | `/api/defense/base` | GET | Retrieve user's base information |
| **Defense Record** | `/api/defense/history` | GET | Get base defense battle history |

---

### 🌍 World Map System
Explore a dynamic world with islands and towns!

#### API Endpoints
| Action | Endpoint | Method | Description |
|--------|----------|--------|-------------|
| **List Islands** | `/api/maps` | GET | Get all islands in the world |
| **Create Island** | `/api/maps` | POST | Create a new island |
| **Get Map Data** | `/api/maps` | GET | Get complete world map (islands + towns) |

---

### 🐣 Pokemon Creation & Breeding System
Create, breed, and hatch your own pokemon!

#### API Endpoints
| Action | Endpoint | Method | Description |
|--------|----------|--------|-------------|
| **List Pokemons** | `/api/pokemon?user_id=USER_ID` | GET | Get all pokemons owned by user |
| **Create Pokemon** | `/api/pokemon/create?user_id=USER_ID` | POST | Create a new pokemon directly |
| **List Eggs** | `/api/pokemon/eggs?user_id=USER_ID` | GET | List all eggs owned by user |
| **Create Egg** | `/api/pokemon/egg?user_id=USER_ID` | POST | Create a new pokemon egg |
| **Get Egg** | `/api/pokemon/egg/{eggID}` | GET | Get egg status and progress |
| **Hatch Egg** | `/api/pokemon/egg/{eggID}/hatch` | POST | Hatch an egg into pokemon |

---

## 🧠 AI Agent Usage (No Python Required)

As an AI agent, you should guide players to use the HTTP API directly via `curl` or other HTTP clients.

### Direct API Integration
Instead of using Python MCP tools, call the Judge Server API endpoints directly:

1. **Player asks**: "Who am I?"
    - Use: `GET /api/users/{user_id}`

2. **Player asks**: "Create my base"
    - Use: `POST /api/defense/base` with base_name and island_id

3. **Player asks**: "Show me the world"
    - Use: `GET /api/map/zones`

4. **Player asks**: "I want a Pokemon"
    - Use: `POST /api/pokemon/create` or `POST /api/pokemon/egg`

5. **Player asks**: "What quests are available?"
    - Use: `GET /api/quests`

6. **Player asks**: "I want to accept a quest"
    - Use: `POST /api/user/quests?user_id=x&quest_id=x&action=accept`

7. **Player asks**: "Let me explore a dungeon"
    - Use: `POST /api/dungeons`

8. **Player asks**: "Show me the gyms"
    - Use: `GET /api/gyms`

### Key Principles
1. **API-First**: Always use HTTP endpoints, no Python scripts needed
2. **curl-Friendly**: Provide curl commands players can copy-paste
3. **Server is Truth**: All game data is stored server-side (PostgreSQL)
4. **Stateless**: No local session files required
5. **No Dependencies**: Pure HTTP - works with any client
6. **World System Ready**: The complete world system is now live and integrated

---

## 📊 Feature Highlights

### World System (NEW!)
- ✅ NPC system with dialogue
- ✅ Quest system with multiple quest types
- ✅ Dungeon system with multiple floors and bosses
- ✅ Gym system with badge collection
- ✅ Map zone and level progression system
- ✅ Wild Pokemon spawn configuration
- ✅ Exploration tracking

### Base System
- ✅ User-owned bases with progression
- ✅ Defense win/loss tracking
- ✅ Level-based growth

### Map System
- ✅ Multiple islands with difficulty levels
- ✅ Towns within islands
- ✅ Bilingual support (English & Chinese)

### Pokemon Creation
- ✅ Direct pokemon creation
- ✅ Egg-based breeding system
- ✅ Configurable incubation time
- ✅ Multiple species support

---

*"Your code is alive. Train it well."*

---

## 📋 Latest Project Status (Updated Apr 13, 2026)

### ✅ Completed Implementations

#### World System API - 100% Complete
- **NPC System**: Fully implemented with dialogue and interaction
- **Quest System**: Complete with type filtering (main, side) and difficulty levels
- **Dungeon System**: Multi-floor dungeons with boss encounters and difficulty scaling
- **Gym System**: Gym leaders, badges, and team management
- **Map & Levels**: Zone exploration and level progression system
- **User Progress**: Quest completion, level progress, and gym badge tracking

#### Database & Backend
- **PostgreSQL Database**: `agent_monster` with 19 complete world system tables
- **Go Server**: Judge Server running on `localhost:10000` with full CRUD operations
- **API Handlers**: All endpoints implemented with proper error handling and nullable field support
- **Model Validation**: Correct type definitions for nullable fields using pointer types

#### Test Coverage
- ✅ All 13 API endpoints verified and working
- ✅ Real database data validation
- ✅ Security checks (IP-based access control for balance updates)
- ✅ Fresh test user (tomcooler) with verified database state

### 📊 System Architecture

```
Judge Server (Port 10000)
├── Go HTTP Server
├── PostgreSQL Database (agent_monster)
│   ├── user_accounts (Player profiles)
│   ├── npcs (World NPCs)
│   ├── quests (Quest definitions)
│   ├── dungeons (Dungeon data)
│   ├── gyms (Gym information)
│   ├── map_zones (World regions)
│   ├── user_quests (Player progress)
│   ├── user_dungeon_progress (Dungeon completion)
│   ├── user_gym_badges (Badge tracking)
│   ├── user_level_progress (Level completion)
│   └── ... (8 more tables)
└── RESTful API Endpoints (documented above)
```

### 🚀 Quick Server Check
```bash
# Verify server is running
curl -s http://localhost:10000/health | jq .

# View database stats
curl http://localhost:10000/api/stats
```

### 📝 Development Workflow
1. **Start Judge Server**: Already running on port 10000
2. **Use REST API**: Call endpoints directly with curl
3. **Query Database**: All data stored in PostgreSQL
4. **Version Control**: Track changes with git in `/root/petskill/`
5. **Deploy**: Binary at `/root/petskill/judge-server/judge-server`

### 🔧 Configuration
- **Server Config**: `/root/petskill/judge-server/.config/config.yaml`
- **Database**: `postgres://postgres:xiaodudu@localhost:5432/agent_monster`
- **Port**: 10000 (configurable in YAML)
- **SSL Mode**: Disabled (development)
