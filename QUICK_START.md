# Agent Monster - 快速开始卡

## ⚡ 最快的方式：复制这5条命令

### 第1条：检查服务器是否运行
```bash
curl http://localhost:8080/health
```

### 第2条：查看自己的信息
```bash
curl http://localhost:8080/api/users/999
```

### 第3条：查看游戏世界
```bash
curl http://localhost:8080/api/map/zones | jq '.'
```

### 第4条：查看所有NPC
```bash
curl "http://localhost:8080/api/npcs?town_id=1" | jq '.'
```

### 第5条：查看所有任务
```bash
curl http://localhost:8080/api/quests | jq '.'
```

**成功！现在你已经开始游戏了！** 🎮

---

## 📖 更多游戏方式

### 与NPC交谈
```bash
curl "http://localhost:8080/api/npcs/talk?npc_id=1"
```

### 接受一个任务
```bash
curl -X POST "http://localhost:8080/api/user/quests" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "999",
    "quest_id": 1,
    "action": "accept"
  }'
```

### 完成任务
```bash
curl -X POST "http://localhost:8080/api/user/quests" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "999",
    "quest_id": 1,
    "action": "complete"
  }'
```

### 进入地下城
```bash
curl -X POST "http://localhost:8080/api/dungeons" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "999",
    "dungeon_id": 1
  }'
```

### 查看所有体操馆
```bash
curl http://localhost:8080/api/gyms | jq '.'
```

### 探索地图关卡
```bash
curl "http://localhost:8080/api/levels?zone_id=1" | jq '.'
```

### 完成一个关卡
```bash
curl -X POST "http://localhost:8080/api/user/levels/progress" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "999",
    "level_id": 1,
    "stars": 3
  }'
```

---

## 📚 需要帮助？

1. **详细新手指南**：阅读 `/root/petskill/skill.md`
2. **完整系统说明**：查看 `/root/petskill/WORLD_SYSTEM_COMPLETE.md`
3. **技术细节**：参考 `/root/petskill/judge-server/WORLD_SYSTEM_IMPLEMENTATION.md`

---

## 🎯 游戏流程建议

```
1️⃣  检查服务器（第1条命令）
2️⃣  查看世界（第3条命令）
3️⃣  查看NPC（第4条命令）
4️⃣  查看任务（第5条命令）
5️⃣  接受任务（接受任务命令）
6️⃣  完成任务（完成任务命令）
7️⃣  赚取金币和经验
8️⃣  重复2-7步，或尝试挑战地下城
```

---

## ✨ 记住这3个重要概念

| 系统 | 作用 | 主要命令 |
|------|------|---------|
| 🤖 **NPC** | 与角色互动 | GET /api/npcs/talk |
| 📜 **任务** | 主要进度 | POST /api/user/quests |
| 🏰 **地下城** | 副本挑战 | POST /api/dungeons |

---

## 🚀 就是这样！

现在复制上面的5条命令，开始你的Agent Monster冒险吧！

*"Your code is alive. Train it well."* 🎮
