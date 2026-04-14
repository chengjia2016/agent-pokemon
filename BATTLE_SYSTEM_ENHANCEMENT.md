# 战斗系统增强文档

## 概述

PetSkill 战斗系统已全面升级，包含三个主要改进方向：

1. **增强的战斗机制** - 更复杂的伤害计算、属性成长系统、特性系统和物品加成
2. **丰富的战斗效果** - 招式动画、状态异常效果、粒子特效、屏幕震动等
3. **完整的战斗UI** - 详细的战斗日志、战略建议、类型相克提示、伤害预览等

---

## 1. 增强的战斗机制

### 1.1 属性成长系统

#### 基础成长率
```
HP:       1.1倍 (每级增长最快)
攻击:     1.05倍
防御:     1.05倍
特攻:     1.08倍
特防:     1.08倍
速度:     1.07倍
```

#### 等级奖励
- 每10级获得额外属性奖励
- HP: +5 (10级) → +50 (100级)
- 攻击: +3 (10级) → +30 (100级)
- 防御: +3 (10级) → +30 (100级)

#### 计算公式
```
属性值 = 基础值 × (成长率 ^ (等级-1)) + 等级奖励
```

### 1.2 特性系统

支持的特性及其效果：

| 特性ID | 特性名 | 效果类型 | 效果描述 |
|--------|--------|---------|---------|
| static | 静电 | 状态应用 | 30%概率使接触者陷入麻痹 |
| speed_boost | 加速 | 属性提升 | 每回合速度+10% |
| filter | 过滤 | 伤害减轻 | 减轻属性弱点伤害至50% |
| natural_cure | 自然恢复 | 治愈 | 更换精灵时治愈异常状态 |

### 1.3 物品加成系统

支持的物品加成：

| 物品ID | 物品名 | 目标属性 | 加成倍数 | 副作用 |
|--------|--------|---------|---------|--------|
| choice_band | 选择围巾 | 攻击 | 1.5x | 只能用同一招式 |
| life_orb | 生命宝珠 | 特攻 | 1.3x | 每回合损失1/10 HP |
| assault_vest | 突击背心 | 特防 | 1.5x | 无 |

### 1.4 增强的伤害计算

新的伤害计算公式包含：

```
最终伤害 = 基础伤害 × 类型克制倍数 × 会心一击倍数 × 天气倍数 × 场地倍数 × 随机浮动(85-100%)
```

#### 示例计算
- **基础伤害**: 皮卡丘使用十万伏特(威力90)攻击喷火龙
- **类型克制**: 电-&gt;飞行 = 2.0倍
- **会心一击**: 无 = 1.0倍
- **天气加成**: 晴天，无加成 = 1.0倍
- **场地加成**: 电气场地 = 1.5倍
- **随机浮动**: 随机 85-100%

---

## 2. 战斗效果系统

### 2.1 招式动画系统

每个招式都有详细的动画数据：

```json
{
  "move_id": "thunderbolt",
  "move_name": "十万伏特",
  "animation_type": "beam",
  "duration": 800,
  "color": "#FFD700",
  "particle_effect": {
    "type": "aura",
    "count": 20,
    "speed": 4.0,
    "lifetime": 600,
    "spread": 60,
    "color_start": "#FFD700",
    "color_end": "#FFA500"
  },
  "sound_effect": "thunderbolt.mp3",
  "screen_effect": {
    "effect": "bloom",
    "intensity": 0.8,
    "duration": 400
  }
}
```

### 2.2 支持的动画类型

| 类型 | 描述 | 示例 |
|-----|------|------|
| projectile | 投掷类 | 飞行系招式、龙系招式 |
| melee | 近战类 | 格斗系、毒系招式 |
| beam | 射线类 | 十万伏特、日光束 |
| status | 变化类 | 睡眠粉、毒粉 |
| particle | 粒子类 | 火焰爆发、水之波动 |

### 2.3 状态异常效果

每个状态异常都有完整的视觉和数据效果：

```json
{
  "condition_id": "status_burn",
  "condition_name": "灼伤",
  "icon": "🔥",
  "badge_color": "#FF4500",
  "overlay_color": "#FFB6C1",
  "animation": "flame_loop",
  "sound_effect": "burn.mp3",
  "stat_modifiers": {
    "attack": 0.5
  },
  "max_duration": 999,
  "can_stack": false,
  "priority": 3
}
```

### 2.4 支持的状态异常

| 状态ID | 状态名 | 效果 | 优先级 |
|--------|--------|------|--------|
| status_freeze | 冻结 | 无法行动 | 9 (最高) |
| status_sleep | 睡眠 | 无法行动 | 8 |
| status_paralysis | 麻痹 | 25%无法行动, 速度-25% | 5 |
| status_burn | 灼伤 | 每回合损失1/8HP, 攻击-50% | 3 |
| status_poison | 中毒 | 每回合损失1/8HP | 2 |
| status_confusion | 混乱 | 33%无法行动并自伤 | 4 |

---

## 3. 战斗UI系统

### 3.1 增强的战斗状态响应

```json
{
  "battle_id": "battle_123",
  "status": "ongoing",
  "turn": 5,
  "weather": "harsh_sunlight",
  "attacker_state": {
    "pokemon_id": "pikachu",
    "name": "Pikachu",
    "level": 50,
    "current_hp": 120,
    "max_hp": 150,
    "hp_percentage": 0.8,
    "hp_bar_color": "green",
    "moves": [...],
    "status": [...],
    "type_weaknesses": {...}
  },
  "battle_log": [...],
  "action_options": {...},
  "ui_hints": {...},
  "statistics": {...}
}
```

### 3.2 详细的战斗日志

每条日志包含：

```json
{
  "round": 3,
  "turn": 1,
  "actor_name": "Pikachu",
  "action_type": "move",
  "action_detail": "使用十万伏特",
  "target_name": "Charizard",
  "damage": 85,
  "effective_index": 2.0,
  "is_critical": false,
  "message_zh": "皮卡丘使用了十万伏特！对喷火龙造成了85点伤害！",
  "color": "#FFD700",
  "icon": "⚡",
  "animation": {
    "type": "damage_float",
    "duration": 500,
    "magnitude": 0.5
  }
}
```

### 3.3 战斗建议系统

实时提供智能建议：

```json
{
  "best_move": {
    "move": {...},
    "score": 150,
    "reasons": ["超级有效！", "威力很强", "目标HP较低"],
    "priority": 10,
    "icon": "🔴"
  },
  "should_switch": true,
  "switch_recommendation": {...},
  "status_threats": [...],
  "overall_tip": "这一回合有强力的招式可用，应该马上使用！"
}
```

### 3.4 类型相克提示

清晰显示属性优劣：

```json
{
  "typing_advantage": {
    "your_type": "electric",
    "opponent_type": "water/flying",
    "advantage": "你有优势",
    "multiplier": 2.0,
    "icon": "⚡",
    "description": "电系招式对水/飞行型造成2倍伤害"
  }
}
```

---

## 4. API 端点

### 4.1 获取增强的战斗状态
```
GET /api/battles/enhanced
响应: EnhancedBattleState (包含所有UI元素)
```

### 4.2 获取招式动画
```
GET /api/moves/animation?move_id=thunderbolt
响应: MoveAnimation
```

### 4.3 获取状态异常效果
```
GET /api/status/effects?condition_id=status_burn
响应: StatusEffectData
```

### 4.4 获取战斗建议
```
POST /api/battles/recommendation
请求体:
{
  "attacker": BattlePokemon,
  "attacker_types": ["electric"],
  "defender": BattlePokemon,
  "defender_types": ["water", "flying"],
  "weather": "harsh_sunlight",
  "available_pokemon": [...]
}
响应: BattleStrategyRecommendation
```

### 4.5 计算增强的伤害
```
POST /api/battles/damage-calculation
请求体:
{
  "attacker": BattlePokemon,
  "defender": BattlePokemon,
  "move": PokemonMove,
  "weather": "harsh_sunlight",
  "terrain": "grassy_terrain"
}
响应: DamageCalculationResponse
```

### 4.6 获取类型相克信息
```
GET /api/types/matchup?attack_type=electric&defend_type=water
响应: 类型相克详细信息
```

---

## 5. 实现文件位置

### 5.1 核心服务
- `internal/service/battle_effects_system.go` - 效果和动画系统
- `internal/service/battle_strategy_system.go` - 策略建议系统
- `internal/service/battle_engine_v2.go` - 改进的战斗引擎

### 5.2 模型定义
- `internal/model/enhanced_battle_ui.go` - 增强UI模型
- `internal/model/pokemon_stats.go` - 宝可梦属性模型

### 5.3 处理器
- `internal/handler/enhanced_battle_handler.go` - 增强战斗处理器

---

## 6. 使用示例

### 6.1 计算伤害并获取推荐

```bash
# 1. 首先获取当前战斗状态
curl http://localhost:8080/api/battles/battle_123/enhanced

# 2. 计算特定招式的伤害
curl -X POST http://localhost:8080/api/battles/damage-calculation \
  -H "Content-Type: application/json" \
  -d '{
    "attacker": {...},
    "defender": {...},
    "move": {"id": "thunderbolt"},
    "weather": "harsh_sunlight"
  }'

# 3. 获取整体战斗建议
curl -X POST http://localhost:8080/api/battles/recommendation \
  -H "Content-Type: application/json" \
  -d '{
    "attacker": {...},
    "attacker_types": ["electric"],
    "defender": {...},
    "defender_types": ["water", "flying"],
    "weather": "harsh_sunlight"
  }'
```

### 6.2 获取招式动画效果

```bash
curl http://localhost:8080/api/moves/animation?move_id=thunderbolt
```

---

## 7. 性能优化

### 7.1 缓存策略
- 招式动画数据在启动时加载，不频繁变化
- 状态异常效果数据静态存储
- 类型相克表使用哈希表快速查询

### 7.2 计算优化
- 伤害计算使用浮点数避免精度丢失
- 会心一击采用快速随机判定
- 招式推荐使用评分系统而非穷举

---

## 8. 未来扩展方向

1. **天气系统扩展** - 支持更多天气类型和组合效果
2. **场地系统** - 扩展场地类型及其交互效果
3. **队伍加成** - 支持队伍级别的统计加成
4. **AI对手** - 基于策略系统实现智能对手决策
5. **动画库扩展** - 为所有招式添加详细动画数据
6. **特性交互** - 支持特性之间的相互作用

---

## 9. 常见问题

**Q: 如何自定义招式动画?**
A: 编辑 `initializeMoveAnimations()` 函数，添加新的招式动画配置即可。

**Q: 如何添加新的特性?**
A: 在 `initializeAbilities()` 函数中添加新的 AbilityData 条目。

**Q: 伤害计算是否考虑所有修正?**
A: 是的，包括天气、场地、特性、物品和随机浮动。

**Q: 战斗建议的优先级如何计算?**
A: 综合考虑类型相克、目标HP、招式优先级和天气加成等因素。

---

## 版本历史

- **v2.0.0** (2024年) - 全面增强战斗系统，新增效果、UI、建议系统
- **v1.0.0** (2023年) - 初始战斗系统实现
