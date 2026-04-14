# Agent Monster v2.1.0 - Pokémon Masters EX 系统文档

**版本**: 2.1.0 (Masters EX Edition)  
**发布日期**: 2026-04-14 15:30 UTC  
**状态**: Production Ready ✅

---

## 目录

1. [系统概述](#系统概述)
2. [配对系统 (Sync Pair System)](#配对系统)
3. [3v3 即时对战系统](#3v3-即时对战系统)
4. [赛事系统 (Tournament System)](#赛事系统)
5. [季节活动系统](#季节活动系统)
6. [拍档招式系统 (Sync Move System)](#拍档招式系统)
7. [API 端点参考](#api-端点参考)
8. [数据库架构](#数据库架构)
9. [部署指南](#部署指南)
10. [故障排除](#故障排除)

---

## 系统概述

Agent Monster v2.1.0 包含了完整的 Pokémon Masters EX 系统实现，为玩家提供：

- **配对系统**: 训练家与宝可梦组成"配对"，每个配对有独特的技能和属性
- **3v3 即时对战**: 全新的战斗系统，支持 3 对 3 的团队对战
- **赛事系统**: "世界宝可梦大师赛"排名系统，包含评分和赛季机制
- **季节活动**: 限时活动、季节服装、特殊宝可梦获取
- **拍档招式**: 强大的组合招式，需要配合度和准备

---

## 配对系统

### 概念

配对(Sync Pair)是训练家与宝可梦的组合。每个配对：
- 有独特的名称和属性加成
- 可以学习最多 3 个常规招式 + 1 个拍档招式
- 有稀有度等级(3-5 星)
- 可以升级到最大等级 130
- 支持潜力解锁(最多 20 级)

### 创建配对

```bash
POST /api/masters/sync-pair/create
Content-Type: application/json

{
    "trainer_id": "trainer_001",
    "pokemon_id": "pokemon_001",
    "trainer_name": "Red",
    "pokemon_name": "Charizard",
    "pokemon_type": "Fire",
    "rarity_stars": 5
}
```

响应示例：
```json
{
    "id": 1,
    "sync_pair_id": "trainer_001_pokemon_001",
    "trainer_id": "trainer_001",
    "pokemon_id": "pokemon_001",
    "trainer_name": "Red",
    "pokemon_name": "Charizard",
    "pokemon_type": "Fire",
    "sync_pair_name": "Red & Charizard",
    "rarity_stars": 5,
    "level": 1,
    "max_level": 130,
    "hp": 100,
    "attack": 50,
    "defense": 50,
    "sp_attack": 50,
    "sp_defense": 50,
    "speed": 50,
    "sync_move_ready_percentage": 0,
    "potential_unlocked": 0,
    "is_active": true,
    "created_at": "2026-04-14T15:30:00Z"
}
```

### 升级配对

```bash
POST /api/masters/sync-pair/level-up
Content-Type: application/json

{
    "sync_pair_id": "trainer_001_pokemon_001",
    "levels": 10
}
```

每升 1 级，所有属性增加约 5 点。

### 解锁潜力

```bash
POST /api/masters/sync-pair/unlock-potential?sync_pair_id=trainer_001_pokemon_001
```

解锁潜力会：
- 增加最大等级 10 级
- 永久增加所有属性 10 点
- 支持最多 20 次解锁(最终最大等级达到 330)

### 设置招式

```bash
POST /api/masters/sync-pair/set-moves
Content-Type: application/json

{
    "sync_pair_id": "trainer_001_pokemon_001",
    "move_1_id": "move_flamethrower",
    "move_2_id": "move_dragon_claw",
    "move_3_id": "move_earthquake",
    "sync_move_id": "sync_move_mega_charizard_x"
}
```

---

## 3v3 即时对战系统

### 概念

3v3 对战是 Masters EX 的核心战斗方式：
- 每方派出 3 个配对进行同时对战
- 实时更新伤害和状态
- 支持多种对战模式(单队、多队)

### 对战会话表结构

```sql
battle_sessions
├── session_id (唯一标识)
├── player_1_id, player_2_id, player_3_id
├── battle_type (single/multi)
├── battle_status (waiting/active/finished)
├── player_*_team (JSON: 3个配对信息)
├── current_round (当前回合)
├── player_*_hp (JSONB: 实时 HP 状态)
├── player_*_score (积分战用)
├── winner_id
├── battle_log (详细日志)
└── timestamps
```

### 对战流程

1. **准备阶段**: 玩家确认队伍组成
2. **开始**: 系统生成对战会话
3. **回合制**:
   - 每回合每个玩家选择行动(使用招式、切换宝可梦等)
   - 系统计算伤害和效果
   - 更新状态和 HP
4. **结束**: 一方全灭或回合数达到上限，计算获胜者

### 对战行动类型

- `move`: 使用普通招式
- `switch`: 切换宝可梦
- `mega-evolve`: 使用超级进化(如果可用)
- `sync-move`: 使用拍档招式(准备条件满足时)

### API 端点 (计划)

```bash
# 创建对战会话
POST /api/masters/battle/create

# 获取对战状态
GET /api/masters/battle/status?session_id=...

# 提交行动
POST /api/masters/battle/action

# 获取对战历史
GET /api/masters/battle/history?player_id=...
```

---

## 赛事系统

### 概念

"世界宝可梦大师赛(WPM)"是全球排名系统：
- 按赛季划分(每赛季约 1-2 个月)
- 基于 ELO 评分系统
- 玩家通过胜负获得/失去评分
- 排名前列可获得独特奖励

### 赛季

```sql
tournament_seasons
├── season_id (唯一标识)
├── season_name/season_name_zh
├── season_number (第 N 赛季)
├── start_date, end_date
├── start_rank (初始评分)
├── max_rank_points (最大评分)
├── is_active
├── reward_pool (JSON: 奖励信息)
└── timestamps
```

### 排名系统

```sql
tournament_rankings
├── ranking_id
├── season_id
├── player_id
├── rank (当前排名)
├── rank_points (当前评分,初始 1200)
├── wins/losses (赛季战绩)
├── win_streak/loss_streak (连胜/连败)
├── highest_rank (最高排名)
├── highest_points (最高评分)
└── timestamps
```

### 评分变化规则

```
如果赢了:
  新评分 = 当前评分 + (25 - (当前评分 - 对手评分) / 100)
  
如果输了:
  新评分 = 当前评分 - (25 - (对手评分 - 当前评分) / 100)
```

这确保高评分玩家战胜低评分玩家获得少量分数,而战胜高评分玩家获得更多分数。

### API 端点 (计划)

```bash
# 获取赛季信息
GET /api/masters/tournament/season?season_id=...

# 获取排名
GET /api/masters/tournament/rankings?season_id=...&limit=100

# 获取玩家排名
GET /api/masters/tournament/ranking?season_id=...&player_id=...

# 获取奖励
GET /api/masters/tournament/rewards?season_id=...&player_id=...
```

---

## 季节活动系统

### 概念

季节活动提供限时内容:
- **故事活动**: 有剧情的活动,解锁新配对
- **挑战活动**: 不同难度的战斗关卡
- **计时战**: 在限制时间内完成,根据耗时评分
- **积分战**: 在固定 5 场对战中获得最高积分
- **特殊活动**: 获取季节服装、道具等

### 季节类型

- Spring (春季,3-5 月)
- Summer (夏季,6-8 月)
- Autumn (秋季,9-11 月)
- Winter (冬季,12-2 月)

### 活动表结构

```sql
seasonal_events
├── event_id (唯一标识)
├── event_name/event_name_zh
├── event_type (story/challenge/time-attack/score-attack)
├── season (spring/summer/autumn/winter)
├── description
├── start_date, end_date
├── featured_sync_pairs (特色配对 JSON)
├── difficulty_levels (难度等级)
├── rewards (奖励信息)
├── is_active, is_limited
└── timestamps
```

### 玩家进度

```sql
event_progress
├── progress_id
├── event_id
├── player_id
├── stage_completed (已完成阶段数)
├── progress_percentage
├── score (积分战用)
├── time_spent (计时战用)
├── rewards_claimed
├── completed_at
└── timestamps
```

### 季节物品

```sql
seasonal_items
├── item_id
├── item_name/item_name_zh
├── item_type (clothing/accessory/material)
├── rarity (common/rare/legendary)
├── effect_description
├── available_seasons
└── timestamps
```

### API 端点 (计划)

```bash
# 获取活跃活动
GET /api/masters/events/active

# 获取活动详情
GET /api/masters/events/detail?event_id=...

# 获取玩家进度
GET /api/masters/events/progress?event_id=...&player_id=...

# 提交活动成绩
POST /api/masters/events/submit

# 领取奖励
POST /api/masters/events/claim-reward
```

---

## 拍档招式系统

### 概念

拍档招式(Sync Move)是配对独有的强力招式:
- 需要在对战中积累能量(拍档招式准备度)
- 准备度达到 100% 时才能使用
- 使用后准备度重置为 0
- 有 2-3 回合的冷却期
- 通常效果强大,可能改变战斗局面

### 拍档招式表

```sql
sync_moves
├── sync_move_id (唯一标识)
├── move_name/move_name_zh
├── trainer_id, pokemon_id
├── description/description_zh
├── power (基础威力,通常 300+)
├── effect_description (特殊效果)
├── required_sync_level (需要的拍档等级)
├── cooldown_turns (冷却回合数)
└── timestamps
```

### 拍档技能

某些配对还有被动/主动技能:

```sql
sync_skills
├── skill_id
├── sync_pair_id
├── skill_name/skill_name_zh
├── skill_type (passive/active)
├── description
├── effect_power
├── required_level
├── unlock_condition
└── timestamps
```

### 准备度机制

- 初始准备度: 0%
- 使用招式时增加: 15-25%(取决于招式)
- 被攻击时增加: 5-10%(取决于伤害)
- 达到 100%: 可使用拍档招式
- 使用后: 重置为 0%

### API 端点

```bash
# 提升拍档招式准备度
POST /api/masters/sync-pair/sync-move/ready
{
    "sync_pair_id": "trainer_001_pokemon_001",
    "percentage": 25
}

# 使用拍档招式
POST /api/masters/sync-pair/sync-move/use?sync_pair_id=trainer_001_pokemon_001
```

---

## API 端点参考

### 配对系统 API

#### 创建配对
```bash
POST /api/masters/sync-pair/create
```

#### 获取配对
```bash
GET /api/masters/sync-pair/get?sync_pair_id=<sync_pair_id>
```

#### 列出玩家配对
```bash
GET /api/masters/sync-pair/list?trainer_id=<trainer_id>
```

#### 升级配对
```bash
POST /api/masters/sync-pair/level-up
```

#### 解锁潜力
```bash
POST /api/masters/sync-pair/unlock-potential?sync_pair_id=<sync_pair_id>
```

#### 设置招式
```bash
POST /api/masters/sync-pair/set-moves
```

#### 导出配对数据
```bash
GET /api/masters/sync-pair/export?player_id=<player_id>
```

### 队伍系统 API

#### 创建队伍
```bash
POST /api/masters/team/create
```

#### 获取队伍
```bash
GET /api/masters/team/get?team_id=<team_id>
```

#### 列出玩家队伍
```bash
GET /api/masters/team/list?player_id=<player_id>
```

### 统计 API

#### 获取玩家统计
```bash
GET /api/masters/stats/player?player_id=<player_id>
```

---

## 数据库架构

### 主要表

1. **sync_pairs** - 配对信息
2. **sync_pair_moves** - 配对招式
3. **sync_moves** - 拍档招式
4. **sync_pair_equipment** - 配对装备
5. **battle_sessions** - 对战会话
6. **battle_actions** - 对战行动
7. **battle_statistics** - 对战统计
8. **tournament_seasons** - 赛季信息
9. **tournament_rankings** - 排名信息
10. **tournament_rewards** - 奖励信息
11. **seasonal_events** - 活动信息
12. **event_progress** - 玩家进度
13. **seasonal_items** - 季节物品
14. **player_sync_pair_dex** - 玩家配对图鉴
15. **player_battle_teams** - 玩家队伍
16. **player_game_statistics** - 玩家统计

### 视图

1. **player_team_view** - 玩家队伍详细视图
2. **tournament_ranking_view** - 赛季排名视图
3. **event_participation_view** - 活动参与统计视图

---

## 部署指南

### 1. 创建数据库表

```bash
# 连接到 PostgreSQL
psql -U agent_monster -d agent_monster_db

# 执行 schema 文件
\i /root/petskill/judge-server/SCHEMA_MASTERS_EX.sql
```

### 2. 初始化数据

```bash
# 创建示例赛季
INSERT INTO tournament_seasons (season_id, season_name, season_number, start_date, end_date, is_active)
VALUES ('season_001', 'Season 1', 1, NOW(), NOW() + INTERVAL '2 months', true);

# 创建示例活动
INSERT INTO seasonal_events (event_id, event_name, event_type, season, start_date, end_date, is_active)
VALUES ('event_001', 'Spring Story', 'story', 'spring', NOW(), NOW() + INTERVAL '14 days', true);
```

### 3. 启动服务

```bash
cd /root/petskill/judge-server
go run cmd/main.go
```

### 4. 验证部署

```bash
# 测试 API
curl -X POST http://localhost:8080/api/masters/sync-pair/create \
  -H "Content-Type: application/json" \
  -d '{
    "trainer_id": "test_trainer",
    "pokemon_id": "test_pokemon",
    "trainer_name": "Test",
    "pokemon_name": "Pikachu",
    "pokemon_type": "Electric",
    "rarity_stars": 5
  }'
```

---

## 故障排除

### 问题 1: 配对创建失败

**症状**: 创建配对返回错误
**解决方案**:
1. 检查 sync_pair_id 是否已存在
2. 确保 trainer_id 和 pokemon_id 非空
3. 检查数据库连接

### 问题 2: 对战系统不响应

**症状**: 对战 API 返回 500 错误
**解决方案**:
1. 检查 battle_sessions 表是否存在
2. 验证玩家队伍数据完整性
3. 检查日志文件

### 问题 3: 排名不更新

**症状**: 赛季排名没有更新
**解决方案**:
1. 确保赛季处于激活状态
2. 运行手动排名更新任务
3. 检查对战记录是否正确记录

### 问题 4: 活动无法领取奖励

**症状**: 活动奖励领取失败
**解决方案**:
1. 验证活动是否进行中
2. 检查玩家是否完成了所需阶段
3. 确保奖励池配置正确

---

## 配置参考

### 环境变量

```bash
# 数据库
DB_HOST=localhost
DB_PORT=5432
DB_USER=agent_monster
DB_PASSWORD=...
DB_NAME=agent_monster_db

# API 服务器
API_PORT=8080
API_HOST=0.0.0.0

# 日志
LOG_LEVEL=info
LOG_FORMAT=json
```

### 性能调优

1. **数据库索引**: 已为所有常用查询字段创建索引
2. **缓存**: 考虑为排名信息添加 Redis 缓存
3. **批量操作**: 使用批量 INSERT/UPDATE 提高性能

---

## 更新日志

### v2.1.0 (2026-04-14)

- ✅ 实现配对系统 (Sync Pair System)
- ✅ 实现 3v3 对战数据结构
- ✅ 实现赛事系统 (Tournament System)
- ✅ 实现季节活动系统 (Seasonal Events)
- ✅ 实现拍档招式系统 (Sync Move System)
- ✅ 创建完整数据库 schema
- ✅ 创建 API 处理器
- ✅ 编写测试脚本
- ✅ 编写完整文档
- Status: Production Ready ✅

---

## 支持与反馈

如有问题或建议，请访问: https://github.com/chengjia2016/agent-pokemon

**更新时间**: 2026-04-14 15:30 UTC  
**版本**: 2.1.0 (Masters EX Edition)  
**状态**: Production Ready ✅
