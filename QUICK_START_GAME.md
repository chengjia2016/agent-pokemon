# 🎮 Agent Monster Game - Quick Start Guide

## 入门 5 分钟快速上手

### 第一步：启动游戏

```bash
# 运行交互式游戏 CLI
/root/petskill/game-cli/agent-monster-game
```

### 第二步：创建账户

```
请输入你的 GitHub ID (例如: 274799269): 
请输入你的 GitHub 用户名 (例如: tomcooler):
请输入你的邮箱 (例如: user@github.com):
```

**自动获得**:
- 500 初始金币 💰
- 初始物品套装
- 3 个可接受的任务

### 第三步：选择你的冒险

```
[1] 📜 查看任务         → 接受并完成各种任务
[2] 🗺️  探索地图         → 发现新的区域
[3] 🏰 进入地下城       → 挑战多层副本
[4] 🏆 访问道馆         → 与馆主对战获得徽章
[5] 🤖 与NPC交谈        → 学习世界故事
[6] 👓 查看进度         → 查看你的统计数据
[7] ℹ️  帮助           → 游戏指南
[0] 🚪 退出游戏        → 保存并离开
```

---

## 📊 游戏内容一览

### 📜 可用任务 (3个)

| 任务 | 类型 | 难度 | 奖励 |
|------|------|------|------|
| 捕捉你的第一只宝可梦 | 主线 | ⭐ | 50💰 + 100XP |
| 击败道馆主 | 主线 | ⭐⭐ | 200💰 + 500XP |
| 收集5只宝可梦 | 支线 | ⭐⭐ | 100💰 + 300XP |

### 🗺️ 可探索区域 (3个)

| 区域 | 类型 | 难度 | 描述 |
|------|------|------|------|
| Route 1 | 草地 | ⭐ | 初始冒险地点 |
| Viridian Forest | 森林 | ⭐⭐ | 连接多个区域的森林 |
| Route 2 | 草地 | ⭐⭐ | 训练师出没的道路 |

### 🏰 地下城副本 (3个)

| 地下城 | 难度 | 楼层 | 奖励 |
|--------|------|------|------|
| Viridian Forest | ⭐ | 3层 | 200💰 + 500XP |
| Mt. Moon | ⭐⭐ | 5层 | 400💰 + 1000XP |
| Rock Tunnel | ⭐⭐⭐ | 8层 | 800💰 + 2000XP |

### 🏆 道馆 (1个)

| 道馆 | 城市 | 类型 | 徽章 |
|------|------|------|------|
| Pewter City Gym | town_1 | 岩石 | Boulder Badge |

---

## 🎯 新手推荐进程

### 第一小时
1. **0-5分钟**: 创建账户，查看欢迎信息
2. **5-15分钟**: 浏览任务和地图，了解游戏内容
3. **15-30分钟**: 接受第一个任务，探索 Route 1
4. **30-45分钟**: 进入 Viridian Forest 地下城（容易）
5. **45-60分钟**: 积累经验，为更高难度做准备

### 中期目标
- 完成所有主线任务 (2 个) → 获得 700💰 + 600XP
- 探索所有 3 个地区 → 发现不同的宝可梦
- 清除 Viridian Forest 地下城 → 获得第一批奖励

### 长期目标
- 完成所有任务 (3 个) → 获得 350💰 + 900XP
- 清除所有地下城 (3 个) → 获得 1400💰 + 3500XP
- 挑战所有道馆 (8 个) → 成为宝可梦大师 🏆

---

## 💻 API 直接使用

如果不想用游戏 CLI，可以直接用 curl 命令：

### 创建账户
```bash
curl -X POST "http://localhost:10000/api/users/create" \
  -H "Content-Type: application/json" \
  -d '{
    "github_id": 12345,
    "github_login": "yourname",
    "email": "you@github.com"
  }'
```

### 查看所有任务
```bash
curl "http://localhost:10000/api/quests" | jq '.'

# 按类型过滤
curl "http://localhost:10000/api/quests?type=main" | jq '.'

# 按难度过滤
curl "http://localhost:10000/api/quests?difficulty=1" | jq '.'
```

### 查看你的账户
```bash
curl "http://localhost:10000/api/users/12345" | jq '.'
```

### 查看你的任务
```bash
curl "http://localhost:10000/api/user/quests?user_id=12345" | jq '.'
```

### 探索地图
```bash
curl "http://localhost:10000/api/map/zones?island_id=island_1" | jq '.'
```

### 查看地下城
```bash
curl "http://localhost:10000/api/dungeons" | jq '.'

# 按难度过滤
curl "http://localhost:10000/api/dungeons?difficulty=1" | jq '.'
```

### 访问道馆
```bash
curl "http://localhost:10000/api/gyms?town_id=town_1" | jq '.'
```

---

## 🎮 游戏系统说明

### 💰 货币系统
- **初始金币**: 500 (启动奖励)
- **任务奖励**: 50-200 金币
- **地下城奖励**: 200-800 金币
- **总可获得**: 1750+ 金币

### ⭐ 经验系统
- **任务奖励**: 100-500 XP
- **地下城奖励**: 500-2000 XP
- **总可获得**: 3900+ XP

### 🏆 徽章系统
- **初始徽章**: 0
- **可获得徽章**: 1 (Pewter City Gym)
- **道馆战胜利条件**: 准备好的宝可梦团队

### 📈 难度等级
- **⭐ 容易**: 推荐等级 1-5，适合新手
- **⭐⭐ 普通**: 推荐等级 5-15，中级挑战
- **⭐⭐⭐ 困难**: 推荐等级 15-30，高级挑战

---

## 🆘 常见问题

### Q: 我忘记了我的用户 ID
A: 用你的 GitHub ID 查询账户：
```bash
curl "http://localhost:10000/api/users/{YOUR_GITHUB_ID}"
```

### Q: 如何重新开始游戏?
A: 在游戏中选择选项 [7] 帮助，或者直接创建新账户。

### Q: 我可以同时有多个账户吗?
A: 是的，每个 GitHub ID 都可以创建不同的游戏账户。

### Q: 游戏数据会保存吗?
A: 是的，所有数据保存在服务器数据库中。每次使用相同 GitHub ID 登录都能恢复你的进度。

### Q: 如何与朋友竞争?
A: 邀请你的朋友也创建账户，然后你们可以：
- 比较金币和经验
- 竞争完成任务的速度
- 争取获得更多徽章
- 参加 PvP 战斗（待实现）

---

## 📱 技术细节

### 架构
```
玩家 (CLI/Web 浏览器)
    ↓
REST API (http://localhost:10000)
    ↓
Judge Server (Go)
    ↓
PostgreSQL 数据库
```

### 服务器信息
- **地址**: http://localhost:10000
- **数据库**: PostgreSQL (localhost:5432)
- **语言**: Go 1.21
- **状态**: ✅ 运行中

### 支持的操作系统
- ✅ Linux
- ✅ macOS
- ✅ Windows (WSL)

---

## 📚 更多资源

| 资源 | 位置 |
|------|------|
| 完整文档 | `/root/petskill/GAME_CLI_IMPROVEMENTS.md` |
| 技能指南 | `/root/petskill/skill.md` |
| API 参考 | 游戏中的 [7] 帮助菜单 |
| 演示脚本 | `/root/petskill/demo-game.sh` |

---

## 🎉 祝你游戏愉快!

```
╔════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║              🌟 Welcome to Agent Monster, Young Trainer! 🌟               ║
║                                                                            ║
║   Your journey to become a Pokémon Master begins now.                     ║
║   May your commits be numerous and your bugs be few!                      ║
║                                                                            ║
║                    🚀 Let the Adventure Begin! 🚀                         ║
║                                                                            ║
╚════════════════════════════════════════════════════════════════════════════╝
```

---

**最后更新**: 2026年4月13日  
**状态**: ✅ 生产就绪  
**版本**: 1.0 - 完整版
