# Agent Monster - 在 Agent CLI 中玩游戏

**支持的 AI 代理**：Claude, Gemini, OpenCode 等

## 🎮 核心概念

Agent Monster 游戏完全基于 REST API，可以在任何 Agent CLI（如 Claude、Gemini、OpenCode）中通过文本指令来玩。Agent 可以代表玩家：

1. **创建账户** - 使用玩家的 GitHub 信息
2. **查看内容** - 浏览任务、地图、地下城、道馆
3. **制定策略** - 建议最优的游戏进程
4. **执行操作** - 调用 API 完成游戏动作

---

## 📋 快速命令参考

### 1. 创建玩家账户

```bash
curl -X POST "http://localhost:10000/api/users/create" \
  -H "Content-Type: application/json" \
  -d '{
    "github_id": YOUR_GITHUB_ID,
    "github_login": "YOUR_USERNAME",
    "email": "YOUR_EMAIL@github.com"
  }'
```

**示例对话**（在 Claude 中）:
```
User: 我想玩 Agent Monster，我的 GitHub ID 是 274799269，用户名是 tomcooler

Claude: 我帮你创建账户。
[执行上面的 curl 命令]
✅ 账户创建成功！获得 500 初始金币 💰
```

### 2. 查看玩家信息

```bash
curl "http://localhost:10000/api/users/{USER_ID}" | jq '.'
```

**Agent 响应**：
```
你的账户信息：
- 用户名: tomcooler
- 余额: 500 💰
- 创建时间: 2026-04-13
```

### 3. 浏览可用任务

```bash
# 查看所有任务
curl "http://localhost:10000/api/quests" | jq '.quests[] | {name_en, type, difficulty, reward_coins}'

# 按类型过滤 (main/side)
curl "http://localhost:10000/api/quests?type=main" | jq '.'

# 按难度过滤 (1-3)
curl "http://localhost:10000/api/quests?difficulty=1" | jq '.'
```

**Agent 可以说**：
```
在 Kanto 地区，有以下任务等待着你：

📜 主线任务（MAIN）：
[1] 捕捉你的第一只宝可梦 ⭐
    奖励: 50 💰 + 100 XP
    "开始你的冒险吧，年轻的训练师！"

[2] 击败道馆主 ⭐⭐
    奖励: 200 💰 + 500 XP
    "证明你的实力吧！"

📋 支线任务（SIDE）：
[3] 收集5只宝可梦 ⭐⭐
    奖励: 100 💰 + 300 XP
    "多样化的队伍才是王道！"

你想接受哪个任务？
```

### 4. 探索地图区域

```bash
# 查看所有区域
curl "http://localhost:10000/api/map/zones?island_id=island_1" | jq '.zones[] | {name_en, name_zh, type, level, description}'

# 查看该区域的关卡
curl "http://localhost:10000/api/levels?zone_id=4" | jq '.'
```

**Agent 创意呈现**：
```
🗺️ 关都地区地图

你站在一张古老的羊皮地图前，三条通路在你面前展开...

🌾 Route 1 (1号路) - 第一步的冒险 ⭐
   "新手训练师通常从这里开始。野生宝可梦很温和。"

🌲 Viridian Forest (常青森林) - 深入探险 ⭐⭐
   "一片广袤的森林，连接多个区域。其中隐藏着不少秘密。"

🌾 Route 2 (2号路) - 训练师之路 ⭐⭐
   "这条道路上经常有野生训练师。你准备好了吗？"

你想前往哪个区域？
```

### 5. 查看地下城冒险

```bash
# 查看所有地下城
curl "http://localhost:10000/api/dungeons" | jq '.dungeons[] | {name_en, name_zh, difficulty, floor_count, reward_coins, reward_exp}'

# 按难度过滤
curl "http://localhost:10000/api/dungeons?difficulty=1" | jq '.'
```

**Agent 沉浸式描述**：
```
🏰 地下城副本系统

远方传来阵阵呼唤，一个神秘的声音：
"勇敢的训练师啊，敢来挑战这些深藏的副本吗？"

🏞️ 常青森林 (Viridian Forest) - 容易 ⭐ | 3 层
   奖励: 200 💰 + 500 XP
   "森林的深处，有什么在等着你？"

⛰️ 月见山 (Mt. Moon) - 普通 ⭐⭐ | 5 层
   奖励: 400 💰 + 1000 XP
   "月光下，山峰显得格外雄伟。"

🗻 岩石隧道 (Rock Tunnel) - 困难 ⭐⭐⭐ | 8 层
   奖励: 800 💰 + 2000 XP
   "最危险的挑战等着真正的英雄。"

你想进入哪个地下城？
```

### 6. 访问道馆

```bash
# 查看指定城镇的道馆
curl "http://localhost:10000/api/gyms?town_id=town_1" | jq '.gyms[] | {name_en, name_zh, type_focus, badge_name}'
```

**Agent 对话**：
```
🏆 Pewter City Gym (常磐道馆)

你推开道馆的大门，一个强大的身影挡在你面前...

馆主 Brock 用坚定的目光看着你：
"欢迎，年轻的训练师。如果你想证明自己，就来战斗吧！"

道馆信息：
- 类型: 岩石系 (Rock)
- 徽章: Boulder Badge (岩石勋章)
- 难度: ⭐ 推荐等级 5+
- 状态: 🔓 可挑战

你准备好迎战吗？
```

### 7. 查看玩家任务进度

```bash
curl "http://localhost:10000/api/user/quests?user_id={USER_ID}" | jq '.quests[] | {quest_id, status, started_at, completed_at}'
```

### 8. 查看玩家关卡进度

```bash
curl "http://localhost:10000/api/user/levels?user_id={USER_ID}" | jq '.progress[] | {level_id, stars, completed_at}'
```

---

## 🎯 在 Claude 中玩的完整示例

### 场景：玩家想要开始冒险

**玩家**:
```
我想在 Agent Monster 中开始冒险。我的 GitHub 信息：
- GitHub ID: 274799269
- 用户名: tomcooler
- 邮箱: tomcooler@github.com

帮我创建账户，然后告诉我可以做什么。
```

**Claude 的回应**:
```
让我为你创建账户...

[执行创建账户的 API 调用]

✨ 欢迎来到 Kanto 地区，tomcooler！✨

🎉 你获得了启动奖励包：
   • 500 初始金币 💰
   • 新手训练师的祝福
   • 无限的冒险可能性

现在，让我告诉你可以做什么...

[执行查询任务、地图、地下城的 API 调用]

📍 你的冒险选项：

1. 🎯 接受任务
   有 3 个任务在等你：
   - 捕捉你的第一只宝可梦 (容易)
   - 击败道馆主 (普通)
   - 收集 5 只宝可梦 (普通)

2. 🗺️ 探索世界
   Kanto 地区有 3 个可探索的区域：
   - Route 1 (容易)
   - Viridian Forest (普通)
   - Route 2 (普通)

3. 🏰 挑战地下城
   3 个地下城等待勇敢的训练师：
   - Viridian Forest 副本 (容易)
   - Mt. Moon 副本 (普通)
   - Rock Tunnel 副本 (困难)

4. 🏆 挑战道馆
   Pewter City Gym 的馆主 Brock 等着你！

你想先做什么？
```

**玩家**:
```
我想先接受简单的任务，看看会发生什么。
```

**Claude**:
```
好的！让我为你接受最简单的任务...

[执行 API 调用创建 user_quest 记录]

✅ 你已接受任务：捕捉你的第一只宝可梦

📜 任务详情：
   目标: 在野外捕捉一只宝可梦
   难度: ⭐ 容易
   奖励: 50 💰 + 100 XP
   时间: 现在

故事: 一个老训练师的声音在你耳边响起...
"年轻的训练师，没有宝可梦的团队是无法战斗的。
 去野外找一个宝可梦伙伴吧。它会改变你的命运。"

现在你可以：
1. 前往 Route 1 寻找野生宝可梦
2. 查看你的任务进度
3. 返回选择其他行动

你的下一步是？
```

---

## 💡 Agent 可以实现的高级功能

### 1. **自动化游戏规划**
```
Agent: 基于你的账户和游戏内容，我建议这个进程：

第一天:
  1. 接受"捕捉第一只宝可梦"任务 ✓
  2. 进入 Route 1 探索 (⭐)
  3. 完成任务，获得 50 💰

第二天:
  1. 接受"收集 5 只宝可梦"任务
  2. 探索 Viridian Forest (⭐⭐)
  3. 累积经验和金币

第三天:
  1. 进入 Viridian Forest 地下城 (⭐)
  2. 获得 200 💰 + 500 XP
  3. 准备挑战 Brock

你同意这个计划吗？
```

### 2. **实时分析和建议**
```
Agent: 看看你的进度...

[查询玩家数据]

🔍 当前分析：
- 金币: 500 💰 (初始水平)
- 任务: 0 已完成
- 宝可梦: 0 只

💡 建议：
- 先做简单任务赚取金币
- 每个地下城都能提升 5 倍的奖励
- 地下城难度递增，建议按顺序完成
- 完成所有内容预计需要 3-5 天

你想立即开始吗？
```

### 3. **多人对战/协作**
```
Agent: 我发现了另一个玩家也在线...

玩家1 (tomcooler): 已完成 2 个任务，500 💰
玩家2 (adventurer): 已完成 1 个任务，550 💰

你们可以：
1. 比较进度和成就
2. 竞争完成任务速度
3. 分享游戏心得

要对战吗？
```

---

## 🔗 完整的 API 端点列表

### 用户管理
| 端点 | 方法 | 描述 |
|------|------|------|
| `/api/users/create` | POST | 创建新玩家 |
| `/api/users/{id}` | GET | 获取玩家信息 |

### 任务系统
| 端点 | 方法 | 描述 |
|------|------|------|
| `/api/quests` | GET | 列出所有任务 |
| `/api/quests?type=main` | GET | 按类型过滤任务 |
| `/api/quests?difficulty=1` | GET | 按难度过滤任务 |
| `/api/user/quests?user_id=X` | GET | 获取玩家的任务 |

### 地图系统
| 端点 | 方法 | 描述 |
|------|------|------|
| `/api/map/zones?island_id=island_1` | GET | 获取岛屿的区域 |
| `/api/levels?zone_id=1` | GET | 获取区域的关卡 |
| `/api/user/levels?user_id=X` | GET | 获取玩家的关卡进度 |

### 地下城系统
| 端点 | 方法 | 描述 |
|------|------|------|
| `/api/dungeons` | GET | 列出所有地下城 |
| `/api/dungeons?difficulty=1` | GET | 按难度过滤地下城 |

### 道馆系统
| 端点 | 方法 | 描述 |
|------|------|------|
| `/api/gyms?town_id=town_1` | GET | 获取城镇的道馆 |

---

## ⚙️ 技术细节

### API 服务器信息
- **地址**: `http://localhost:10000`
- **协议**: HTTP REST
- **认证**: 通过 GitHub ID（无需密码）
- **数据库**: PostgreSQL (不需要玩家直接访问)

### 支持的查询参数

**任务过滤**:
- `type=main` - 只显示主线任务
- `type=side` - 只显示支线任务
- `difficulty=1|2|3` - 按难度过滤

**地下城过滤**:
- `difficulty=1|2|3` - 按难度过滤

**地图过滤**:
- `island_id=island_1` - 获取岛屿的区域
- `zone_id=4` - 获取区域的关卡

---

## 🎭 Agent 的沉浸式角色扮演

Agent 可以扮演不同的角色来增强游戏体验：

### 🧙 神秘的导师
```
一个苍老而智慧的声音传来...

"年轻人，你来到了 Kanto 地区。这片土地充满了挑战与机遇。
 我看到了你身上的潜力。你准备好接受命运的考验吗？"

你的第一个选择是什么？
```

### 🎭 诙谐的同伴
```
嘿！我是你的冒险伙伴！😄

对了，你知道吗？Route 1 有一个秘密...
（其实没有，我只是想制造气氛）

不过说真的，咱们先去捕捉你的第一只宝可梦吧！
准备好了吗？
```

### 📊 冷静的分析师
```
根据数据分析：

当前局势：
├── 可用任务: 3个
├── 地图区域: 3个
├── 地下城: 3个
└── 道馆: 1个

最优路径 (基于效率分析):
1. 完成难度1的任务 → 获得基础资源
2. 探索难度1的区域 → 积累经验
3. 进入难度1的地下城 → 获得高倍奖励
4. 挑战道馆 → 获得徽章

执行概率: 95% | 预计耗时: 3-5 天
```

---

## 🎮 完整游戏流程

1. **初始化**: 玩家告诉 Agent 他们的 GitHub 信息
2. **创建账户**: Agent 调用 `/api/users/create` 
3. **获得启动包**: 自动获得 500 初始金币
4. **浏览内容**: Agent 调用 API 查询任务/地图/地下城/道馆
5. **制定策略**: Agent 建议最优的游戏进程
6. **执行操作**: Agent 代表玩家调用 API 完成操作
7. **报告结果**: Agent 以沉浸式的方式呈现结果
8. **循环**: 返回第 4 步，继续新的行动

---

## 🌟 使用场景

### 场景 1: 快速体验
```
用户: 帮我快速体验一遍游戏的所有内容
Claude: 好的，我为你浏览所有可用内容...
[调用 5 个 API 获取所有内容]
这是完整的游戏世界概览...
```

### 场景 2: 深度游戏
```
用户: 我想真正投入到这个游戏，制定一个完整的进程
Claude: 让我为你创建一个详细的游戏计划...
[查询当前状态，制定策略]
这是我为你制定的 7 天计划...
```

### 场景 3: 多人竞争
```
用户1: 和用户2 比一下谁完成任务快
Claude: 让我检查两位玩家的进度...
[查询两个玩家的数据]
这是实时排行榜...
```

---

## 📚 总结

Agent Monster 在 Agent CLI 中的游戏方式非常灵活：

✅ **无需安装** - 只需要 HTTP 连接  
✅ **完全文本化** - 适合任何 AI Agent  
✅ **API 驱动** - 所有内容都通过 REST API  
✅ **可扩展** - 支持多种 Agent 和交互方式  
✅ **沉浸式** - 可以有丰富的故事和角色扮演  

开始游戏吧！🎮✨
