# 🌍 Agent Monster 多语言系统 - 测试指南

## 📋 测试概览

本文档提供了完整的多语言系统测试指南。所有测试将使用 GitHub 用户 `chengjia2016` 进行。

## 🚀 快速开始

### 1. 验证服务器正在运行
```bash
curl -s http://pokemon.openx.pro:10000/health
# 预期输出: {"status": "healthy"}
```

### 2. 运行完整测试套件
```bash
chmod +x test_multilingual_system.sh
./test_multilingual_system.sh
```

## 📊 测试用例详情

### Test 1: 获取可用语言列表
- **端点**: `GET /api/language/list`
- **预期**: 返回英文 (en) 和中文 (zh) 两种语言
- **验证**: 确认两种语言都是活跃的 (is_active: true)

### Test 2 & 3: 获取 UI 字符串
- **端点**: `GET /api/language/strings?language=en|zh`
- **预期**: 返回 50 个 UI 字符串的翻译
- **验证**: 验证两种语言都有完整的菜单翻译

### Test 4: 设置用户语言偏好为中文
- **端点**: `POST /api/language/select`
- **数据**: `{"user_id": "chengjia2016", "language_code": "zh"}`
- **预期**: 返回成功响应，确认语言已设置为中文

### Test 5: 验证用户当前语言 (中文)
- **端点**: `GET /api/language/current?user_id=chengjia2016`
- **预期**: 返回 `language_code: "zh"`

### Test 6 & 7: NPC 对话本地化
- **端点**: `GET /api/language/npc-dialogue?npc_id=1&language=zh|en`
- **预期**: 
  - 中文返回馆主米斯蒂的中文对白
  - 英文返回相应英文对白
- **验证**: 确认对话内容正确翻译

### Test 8: 其他 NPC 对话 (大木博士)
- **端点**: `GET /api/language/npc-dialogue?npc_id=9&language=zh`
- **预期**: 返回大木博士的中文对白
- **示例**: "我是大木博士，很高兴见到你，年轻的训练师！"

### Test 9 & 10: 任务本地化
- **端点**: `GET /api/language/quest?quest_id=1&language=zh|en`
- **预期**:
  - 中文: "在华蓝市击败米斯蒂道馆馆主并获得钢铁徽章"
  - 英文: "Defeat Gym Leader Misty in Cerulean City and earn the Steel Badge"
- **验证**: 任务描述和奖励都正确翻译

### Test 11: 多个任务查询
- **端点**: `GET /api/language/quest?quest_id=5&language=zh`
- **预期**: 返回火焰徽章任务的中文描述

### Test 12: 切换用户语言回英文
- **端点**: `POST /api/language/select`
- **数据**: `{"user_id": "chengjia2016", "language_code": "en"}`
- **预期**: 返回成功响应，确认语言已设置为英文

### Test 13: 验证用户当前语言 (英文)
- **端点**: `GET /api/language/current?user_id=chengjia2016`
- **预期**: 返回 `language_code: "en"`

### Test 14 & 15: 其他 NPC 对话测试
- **端点**: 各种 NPC ID 的对话获取
- **预期**: 验证不同 NPC 的中英文对话都正确工作

## 🎮 手动测试

### 1. 查看所有语言
```bash
curl http://pokemon.openx.pro:10000/api/language/list | jq .
```

### 2. 设置你的语言为中文
```bash
curl -X POST http://pokemon.openx.pro:10000/api/language/select \
  -H "Content-Type: application/json" \
  -d '{"user_id":"chengjia2016", "language_code":"zh"}'
```

### 3. 获取中文 UI 字符串
```bash
curl http://pokemon.openx.pro:10000/api/language/strings?language=zh | jq .
```

### 4. 获取中文 NPC 对话
```bash
curl http://pokemon.openx.pro:10000/api/language/npc-dialogue?npc_id=1&language=zh | jq .
```

### 5. 获取中文任务描述
```bash
curl http://pokemon.openx.pro:10000/api/language/quest?quest_id=1&language=zh | jq .
```

### 6. 切换回英文
```bash
curl -X POST http://pokemon.openx.pro:10000/api/language/select \
  -H "Content-Type: application/json" \
  -d '{"user_id":"chengjia2016", "language_code":"en"}'
```

## 🔍 预期结果

### 成功标准

✅ 所有 15 个测试都通过
✅ 所有响应都包含 `"success": true`
✅ 中文翻译显示正确的汉字内容
✅ 英文翻译显示正确的英文内容
✅ 用户语言偏好正确保存和检索

### NPC 对话示例

#### 米斯蒂道馆馆主 (NPC ID: 1)
- **中文**: "欢迎来到华蓝道馆！我是馆主米斯蒂，钢铁属性宝可梦的训练大师。"
- **英文**: "Welcome to Cerulean Gym! I am Gym Leader Misty, master of steel-type Pokémon."

#### 大木博士 (NPC ID: 9)
- **中文**: "我是大木博士，很高兴见到你，年轻的训练师！你开始你的冒险了吗？"
- **英文**: "I am Professor Oak, delighted to meet you, young trainer! Have you started your adventure?"

#### 詹妮警官 (NPC ID: 10)
- **中文**: "罪犯逮捕任务！我是詹妮警官，需要你的帮助来追捕逃犯。"
- **英文**: "Criminal capture mission! I am Officer Jenny, I need your help catching the criminal."

### 任务示例

#### 任务 1: 华蓝道馆挑战
- **中文**: 
  - 描述: "在华蓝市击败米斯蒂道馆馆主并获得钢铁徽章"
  - 奖励: "钢铁徽章"
- **英文**:
  - 描述: "Defeat Gym Leader Misty in Cerulean City and earn the Steel Badge"
  - 奖励: "Steel Badge"

#### 任务 5: 红莲镇挑战
- **中文**:
  - 描述: "在红莲镇击败科特道馆馆主并获得火焰徽章"
  - 奖励: "火焰徽章"

## 📈 测试指标

- **总测试用例**: 15
- **覆盖的 API 端点**: 6
  - `/api/language/list`
  - `/api/language/strings`
  - `/api/language/select`
  - `/api/language/current`
  - `/api/language/npc-dialogue`
  - `/api/language/quest`
- **覆盖的 NPC**: 3+ (Gym Leaders, Professor, Officer)
- **覆盖的任务**: 2+ (Gym Badge quests)
- **语言覆盖**: 2 (English, Chinese)

## 🛠️ 故障排除

### 问题 1: 服务器不可达
```
解决方案: 确保 Judge Server 在 http://pokemon.openx.pro:10000 正在运行
命令: curl -s http://pokemon.openx.pro:10000/health
```

### 问题 2: 不支持的语言
```
解决方案: 只使用 "en" 或 "zh" 作为语言代码
```

### 问题 3: 用户语言偏好不存在
```
解决方案: 系统会自动返回 "en" 作为默认值
```

### 问题 4: NPC 对话或任务未找到
```
解决方案: 确保使用正确的 NPC ID 和任务 ID
         有效的 NPC IDs: 1-10
         有效的 Quest IDs: 1-13
```

## 📝 测试日志

运行测试后，你应该看到类似以下的输出：

```
╔════════════════════════════════════════════════════════════╗
║     Agent Monster - Multilingual System Testing            ║
║     Testing User: chengjia2016
║     Server: http://pokemon.openx.pro:10000
╚════════════════════════════════════════════════════════════╝

Test 1: Get Available Languages
✅ PASSED

Test 2: Get UI Strings (English)
✅ PASSED

[... more tests ...]

╔════════════════════════════════════════════════════════════╗
║                     TEST SUMMARY                           ║
╠════════════════════════════════════════════════════════════╣
║ Total Tests:  15
║ Passed:       15
║ Failed:       0
║ ✅ ALL TESTS PASSED!
╚════════════════════════════════════════════════════════════╝
```

## ✅ 验证清单

在完成测试后，请确保以下内容都已验证：

- [ ] 服务器可以访问
- [ ] 可以获取语言列表
- [ ] 可以获取英文 UI 字符串
- [ ] 可以获取中文 UI 字符串
- [ ] 可以为用户设置语言偏好
- [ ] 可以检索用户语言偏好
- [ ] NPC 对话在英文中正确显示
- [ ] NPC 对话在中文中正确显示
- [ ] 任务描述在英文中正确显示
- [ ] 任务描述在中文中正确显示
- [ ] 可以在语言之间切换
- [ ] 所有 15 个测试都通过

## 📞 支持

如有任何问题，请联系开发者：
- GitHub: @chengjia2016
- 项目: agent-pokemon

---

**最后更新**: 2026-04-14
**测试用户**: chengjia2016
**服务器**: http://pokemon.openx.pro:10000
