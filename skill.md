# Agent Monster - Gameplay Skill & Entry Point

> **Version**: 2.0.0 (Multilingual Edition)  
> **Last Updated**: 2026-04-14 15:30 UTC  
> **Status**: Production Ready ✅

Welcome to **Agent Monster**, an AI-powered RPG where your GitHub repository becomes a digital pet. This guide serves as the entry point for players and the operational manual for AI agents.

---

## 📦 项目信息

- **服务器地址**: http://pokemon.openx.pro:10000
- **开发者**: chengjia2016
- **版本**: 2.0.0 (Multilingual Edition)
- **上次更新**: 2026年4月14日 15:30

---

## 🎯 快速导航

- [快速开始 (Quick Start)](#-快速开始-your-first-5-minutes)
- [🌍 多语言系统 (NEW!)](#-多语言系统-new)
- [世界系统 (World System)](#-新功能完整世界系统-new)
- [API 完整列表](#api-完整参考)
- [AI Agent 集成](#-entry-point-agent-cli-integration)

---

## 🚀 Quick Start: Your First 5 Minutes

**新玩家必读！按以下步骤立即开始游戏：**

> **📌 重要信息**
> **服务器地址**：`http://pokemon.openx.pro:10000`
> （不使用 `localhost`，使用公网地址）

### 1️⃣ 检查游戏服务器是否运行
```bash
curl -s http://pokemon.openx.pro:10000/health
```
*应该看到*:
```json
{"status": "healthy"}
```

### 2️⃣ 查看你是谁（获取用户ID）
```bash
# 用你的GitHub用户ID替换 'your_github_id'
curl http://pokemon.openx.pro:10000/api/users/your_github_id
```

### 3️⃣ 探索游戏世界
```bash
# 查看所有岛屿和城镇
curl http://pokemon.openx.pro:10000/api/maps | jq .
```

### 4️⃣ 创建你的基地
```bash
curl -X POST "http://pokemon.openx.pro:10000/api/defense/base" \
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

## 🌍 多语言系统 (NEW!) ⭐

**Version 2.0.0 新增功能** - Agent Monster 现在完全支持多语言游戏体验！所有菜单、NPC对话和任务都可以用英文或中文显示。

### ✨ 多语言功能概览

| 功能 | 英文支持 | 中文支持 | 说明 |
|------|---------|---------|------|
| UI菜单 | ✅ 50项 | ✅ 50项 | 所有游戏菜单和按钮 |
| NPC对话 | ✅ 28条 | ✅ 28条 | 8位道馆主 + 大木博士 + 詹妮警官 |
| 任务描述 | ✅ 13项 | ✅ 13项 | 所有13个任务的完整描述 |
| 系统消息 | ✅ 完整 | ✅ 完整 | 战斗、任务、奖励等所有提示 |

### 1️⃣ 🌐 选择游戏语言

```bash
# 设置用户语言偏好为中文
curl -X POST "http://pokemon.openx.pro:10000/api/language/select" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_github_id",
    "language_code": "zh"
  }'

# 响应示例
{
  "success": true,
  "user_id": "your_github_id",
  "language_code": "zh",
  "message": "Language preference updated"
}
```

### 2️⃣ 📝 获取所有UI字符串翻译

```bash
# 获取中文UI字符串（菜单、按钮等）
curl "http://pokemon.openx.pro:10000/api/language/strings?language=zh"

# 获取英文UI字符串
curl "http://pokemon.openx.pro:10000/api/language/strings?language=en"

# 响应示例
{
  "success": true,
  "language": "zh",
  "strings": {
    "menu.main": "主菜单",
    "menu.quest": "任务",
    "menu.battle": "战斗",
    "menu.inventory": "背包",
    ...
  }
}
```

### 3️⃣ 🤖 获取本地化NPC对话

```bash
# 获取NPC用中文说话
curl "http://pokemon.openx.pro:10000/api/language/npc-dialogue?npc_id=1&language=zh"

# 英文对话
curl "http://pokemon.openx.pro:10000/api/language/npc-dialogue?npc_id=1&language=en"

# 响应示例 (中文)
{
  "success": true,
  "npc_id": 1,
  "language": "zh",
  "dialogue": "欢迎来到华蓝道馆！我是馆主米斯蒂，钢铁属性宝可梦的训练大师。",
  "context": "challenge",
  "type": "gym_leader"
}

# 响应示例 (English)
{
  "success": true,
  "npc_id": 1,
  "language": "en",
  "dialogue": "Welcome to Cerulean Gym! I am Gym Leader Misty, master of steel-type Pokémon.",
  "context": "challenge",
  "type": "gym_leader"
}
```

### 4️⃣ 📜 获取本地化任务描述

```bash
# 获取中文任务描述
curl "http://pokemon.openx.pro:10000/api/language/quest?quest_id=1&language=zh"

# 英文任务描述
curl "http://pokemon.openx.pro:10000/api/language/quest?quest_id=1&language=en"

# 响应示例 (中文)
{
  "success": true,
  "quest_id": 1,
  "language": "zh",
  "quest": {
    "id": 1,
    "quest_name": "华蓝道馆挑战",
    "description": "在华蓝市击败米斯蒂道馆馆主并获得钢铁徽章",
    "reward_item": "钢铁徽章",
    "required_steps": 1
  }
}

# 响应示例 (English)
{
  "success": true,
  "quest_id": 1,
  "language": "en",
  "quest": {
    "id": 1,
    "quest_name": "Cerulean Gym Challenge",
    "description": "Defeat Gym Leader Misty in Cerulean City and earn the Steel Badge",
    "reward_item": "Steel Badge",
    "required_steps": 1
  }
}
```

### 5️⃣ 🔄 查询当前用户语言偏好

```bash
# 获取用户当前设置的语言
curl "http://pokemon.openx.pro:10000/api/language/current?user_id=your_github_id"

# 响应示例
{
  "success": true,
  "user_id": "your_github_id",
  "language_code": "zh"
}
# 如果未设置，默认返回 "en" (英文)
```

### 6️⃣ 📋 查看所有支持的语言

```bash
# 列出所有可用语言
curl "http://pokemon.openx.pro:10000/api/language/list"

# 响应示例
{
  "success": true,
  "languages": [
    {
      "id": 1,
      "code": "en",
      "name": "English",
      "is_active": true
    },
    {
      "id": 2,
      "code": "zh",
      "name": "中文",
      "is_active": true
    }
  ],
  "count": 2
}
```

### 🎮 多语言快速开始

```bash
# 第1步：为你的账户设置语言
curl -X POST "http://pokemon.openx.pro:10000/api/language/select" \
  -H "Content-Type: application/json" \
  -d '{"user_id": "chengjia2016", "language_code": "zh"}'

# 第2步：获取中文UI字符串
curl "http://pokemon.openx.pro:10000/api/language/strings?language=zh" | jq .

# 第3步：用中文获取NPC对话
curl "http://pokemon.openx.pro:10000/api/language/npc-dialogue?npc_id=1&language=zh" | jq .

# 第4步：用中文查看任务
curl "http://pokemon.openx.pro:10000/api/language/quest?quest_id=1&language=zh" | jq .

# 第5步：切换回英文（随时可以切换）
curl -X POST "http://pokemon.openx.pro:10000/api/language/select" \
  -H "Content-Type: application/json" \
  -d '{"user_id": "chengjia2016", "language_code": "en"}'
```

### 🌟 多语言API端点完整列表

| 端点 | 方法 | 说明 |
|------|------|------|
| `/api/language/list` | GET | 获取所有支持的语言 |
| `/api/language/strings` | GET | 获取UI字符串翻译 |
| `/api/language/select` | POST | 设置用户语言偏好 |
| `/api/language/current` | GET | 获取用户当前语言 |
| `/api/language/npc-dialogue` | GET | 获取本地化NPC对话 |
| `/api/language/quest` | GET | 获取本地化任务描述 |

### 🛠️ 多语言系统技术架构

- **数据库表**: `languages`, `ui_strings`, `ui_translations`, `user_language_preferences`, `npc_dialogues` (扩展), `quests` (扩展)
- **支持语言**: English (en), 中文 (zh)
- **UI字符串**: 50个关键菜单和界面文本
- **翻译覆盖**: 100%完整双语支持
- **用户偏好**: 持久化存储每个用户的语言选择
- **自动回退**: 若某语言翻译缺失，自动使用英文

---

## 🎯 新玩家建议流程

### 你的第一个小时应该做什么：

```
⏱️ 0-5分钟：入门
  ✓ 启动游戏服务器
  ✓ 创建你的基地
  ✓ 选择你喜欢的语言（英文/中文）

⏱️ 5-15分钟：探索NPC和任务
  ✓ 查看城镇中的NPC
  ✓ 获取你选定语言的NPC对话
  ✓ 接受1-2个简单任务（用你选定的语言）

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

## 🎮 新功能：完整世界系统

### ✨ 你现在可以做什么

#### 1. 🤖 与NPC交互
```bash
# 获取城镇中所有NPC
curl "http://pokemon.openx.pro:10000/api/npcs?town_id=1"

# 与NPC对话
curl "http://pokemon.openx.pro:10000/api/npcs/talk?npc_id=1"
```

#### 2. 📜 接受和完成任务
```bash
# 查看可用任务
curl "http://pokemon.openx.pro:10000/api/quests"

# 接受任务
curl -X POST "http://pokemon.openx.pro:10000/api/user/quests?user_id=your_id&quest_id=1&action=accept" \
  -H "Content-Type: application/json"

# 完成任务
curl -X POST "http://pokemon.openx.pro:10000/api/user/quests?user_id=your_id&quest_id=1&action=complete" \
  -H "Content-Type: application/json"

# 查看你的任务进度
curl "http://pokemon.openx.pro:10000/api/user/quests?user_id=your_id"
```

#### 3. 🏰 地下城副本
```bash
# 进入地下城
curl -X POST "http://pokemon.openx.pro:10000/api/dungeons" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_id",
    "dungeon_id": 1
  }'
```

#### 4. 🏆 道馆系统
```bash
# 查看所有道馆
curl "http://pokemon.openx.pro:10000/api/gyms"

# 查看你已获得的徽章
curl "http://pokemon.openx.pro:10000/api/gyms?user_id=your_id"
```

#### 5. 🗺️ 地图探索
```bash
# 查看所有地图区域
curl "http://pokemon.openx.pro:10000/api/maps"

# 探索特定地图
curl -X POST "http://pokemon.openx.pro:10000/api/maps/traverse" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your_id",
    "map_id": 1
  }'
```

---

## 🛠️ Entry Point: Agent CLI Integration

As an AI agent, you should guide players to play Agent Monster directly through this CLI using REST API calls.

### Key Points for AI Agents:

1. **No Installation Required** - Just HTTP API calls
2. **Text-Based Gameplay** - Perfect for AI/LLM interaction
3. **Full API Support** - Every game action is an API endpoint
4. **Rich Storytelling** - Add narrative context around API results
5. **Multi-Agent Support** - Works with Claude, Gemini, OpenCode, etc.

### Agent Workflow:

```
1. User asks to play Agent Monster
2. Agent creates account via /api/users/create
3. Agent queries game content (quests, map, dungeons, gyms)
4. Agent presents content with immersive storytelling
5. Agent executes player actions via API calls
6. Agent updates world state and continues narrative
```

### Example Agent Interaction Flow:

```bash
# Step 1: Create user account
curl -X POST "http://pokemon.openx.pro:10000/api/users/create" \
  -H "Content-Type: application/json" \
  -d '{"github_id": "player_name"}'

# Step 2: Get user info
curl "http://pokemon.openx.pro:10000/api/users/player_name"

# Step 3: Get available quests
curl "http://pokemon.openx.pro:10000/api/quests"

# Step 4: Present quest to user and get their choice
# (Agent narrates the quest and waits for player input)

# Step 5: Accept chosen quest
curl -X POST "http://pokemon.openx.pro:10000/api/user/quests?user_id=player_name&quest_id=1&action=accept"

# Step 6: Execute quest actions and update narrative
```

---

## 📊 API 完整参考

### 用户管理
| 操作 | 端点 | 方法 | 说明 |
|------|------|------|------|
| 创建用户 | `/api/users/create` | POST | 创建新用户账户 |
| 获取用户信息 | `/api/users/{user_id}` | GET | 获取用户详情 |

### 多语言系统 (NEW!)
| 操作 | 端点 | 方法 | 说明 |
|------|------|------|------|
| 获取语言列表 | `/api/language/list` | GET | 所有支持的语言 |
| 获取UI字符串 | `/api/language/strings` | GET | UI翻译 |
| 设置语言偏好 | `/api/language/select` | POST | 用户语言选择 |
| 获取当前语言 | `/api/language/current` | GET | 用户语言查询 |
| 获取NPC对话 | `/api/language/npc-dialogue` | GET | 本地化对话 |
| 获取任务信息 | `/api/language/quest` | GET | 本地化任务 |

### NPC与任务
| 操作 | 端点 | 方法 | 说明 |
|------|------|------|------|
| 获取NPC | `/api/npcs` | GET | 获取城镇NPC列表 |
| 与NPC交谈 | `/api/npcs/talk` | POST | 触发NPC对话 |
| 获取任务 | `/api/quests` | GET | 获取所有任务 |
| 用户任务 | `/api/user/quests` | GET/POST | 用户任务管理 |
| 完成任务 | `/api/quests/complete` | POST | 标记任务完成 |

### 地下城与道馆
| 操作 | 端点 | 方法 | 说明 |
|------|------|------|------|
| 获取地下城 | `/api/dungeons` | GET | 地下城列表 |
| 进入地下城 | `/api/dungeons` | POST | 开始副本 |
| 获取道馆 | `/api/gyms` | GET | 道馆列表 |

### 地图探索
| 操作 | 端点 | 方法 | 说明 |
|------|------|------|------|
| 获取地图 | `/api/maps` | GET | 地图列表 |
| 生成用户地图 | `/api/maps/generate` | POST | 生成新地图 |
| 探索地图 | `/api/maps/traverse` | POST | 地图遍历 |

### 战斗系统
| 操作 | 端点 | 方法 | 说明 |
|------|------|------|------|
| 开始战斗 | `/api/battles/start` | POST | 开始新战斗 |
| 执行回合 | `/api/battles/{id}/round` | POST | 战斗回合 |
| 结束战斗 | `/api/battles/{id}/end` | POST | 战斗结束 |
| 获取战斗统计 | `/api/battles/stats` | GET | 战斗数据 |

---

## 📈 系统架构

### 核心组件

```
Agent Monster (v2.0.0)
│
├── 🌍 多语言系统 (NEW!)
│   ├── LanguageService
│   ├── 6个API端点
│   ├── 双语UI支持 (50项)
│   ├── 双语NPC对话 (28条)
│   └── 双语任务系统 (13项)
│
├── 🎮 游戏系统
│   ├── 用户管理
│   ├── NPC系统
│   ├── 任务系统
│   ├── 地图探索
│   ├── 战斗引擎
│   ├── 地下城副本
│   └── 道馆挑战
│
└── 🗄️ 数据库
    ├── 用户数据
    ├── 游戏内容
    ├── 语言数据
    └── 游戏状态
```

### 技术栈

- **后端**: Go (golang)
- **数据库**: PostgreSQL
- **API**: RESTful HTTP
- **部署**: Docker (可选)

---

## 🚀 版本历史

### v2.0.0 (2026-04-14) - 多语言版本 ⭐
- ✅ 添加完整的多语言支持系统
- ✅ 6个新的API端点
- ✅ 50个UI字符串翻译
- ✅ 28条NPC对话翻译
- ✅ 13个任务描述翻译
- ✅ 用户语言偏好管理
- ✅ 英文/中文完全支持

### v1.0.0 (之前)
- 用户系统
- NPC与任务
- 地下城副本
- 道馆挑战
- 地图探索
- 战斗系统

---

## ✅ 功能完成度

| 功能 | 状态 | 说明 |
|------|------|------|
| 用户系统 | ✅ 完成 | 用户创建和管理 |
| NPC系统 | ✅ 完成 | 10个NPC配置 |
| 任务系统 | ✅ 完成 | 13个任务 |
| 地图系统 | ✅ 完成 | 12个城镇/地区 |
| 战斗系统 | ✅ 完成 | 完整战斗引擎 |
| 地下城副本 | ✅ 完成 | 3个地下城 |
| 道馆系统 | ✅ 完成 | 8个道馆 |
| 多语言系统 | ✅ 完成 | 英文/中文支持 |

---

## 📞 支持和文档

### 测试和文档
- `TESTING_INSTRUCTIONS_FOR_CHENGJIA2016.md` - 完整测试指南
- `MULTILINGUAL_TESTING_GUIDE.md` - 多语言系统测试
- `test_multilingual_system.sh` - 自动化测试脚本
- `quick_test_multilingual.sh` - 快速验证脚本

### SQL脚本
- `complete_chinese_translations.sql` - 中文翻译数据导入

---

## 🎓 学习资源

### 快速学习
1. 阅读 Quick Start 部分（5分钟）
2. 运行快速测试脚本（1分钟）
3. 尝试一个简单的API调用（2分钟）

### 深入学习
1. 研究 NPC 和任务系统
2. 理解战斗和道馆系统
3. 学习多语言API的使用

### 高级用法
1. 通过AI agent与游戏交互
2. 创建自己的游戏内容
3. 构建自定义游戏客户端

---

## 📝 许可证

Agent Monster 是由 chengjia2016 开发的开源项目。

---

## 🙏 致谢

感谢所有参与开发、测试和改进这个项目的人！

---

## 📊 项目统计

| 项目 | 数值 |
|------|------|
| 版本 | 2.0.0 |
| 总行数 | 1000+ |
| API 端点 | 30+ |
| 支持语言 | 2 |
| NPC数量 | 10 |
| 任务数量 | 13 |
| 城镇/地区 | 12 |
| UI字符串 | 50+ |

---

**🎮 现在就开始游戏吧！**

祝你在 Agent Monster 的世界中冒险愉快！

> Version 2.0.0 | Last Updated: 2026-04-14 15:30 UTC | Status: Production Ready ✅
