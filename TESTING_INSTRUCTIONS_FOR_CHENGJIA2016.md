# 🎮 Agent Monster 多语言系统 - chengjia2016 测试说明

## 📝 项目完成总结

我已经为 Agent Monster 游戏系统实现了完整的多语言支持系统。所有代码已经提交到 GitHub，你现在可以使用以下指南进行测试。

---

## 🚀 快速开始测试

### 第一步：运行快速测试（推荐用于快速验证）

```bash
cd /root/petskill
bash quick_test_multilingual.sh
```

这将在 5 秒内完成基本功能验证，输出：
- ✅ 可用语言列表
- ✅ 中文 UI 字符串
- ✅ NPC 中文对话
- ✅ 任务中文描述
- ✅ 用户语言偏好切换

### 第二步：运行完整测试套件（推荐用于全面验证）

```bash
cd /root/petskill
bash test_multilingual_system.sh
```

这将执行 15 个完整的测试用例，验证：
- 🌍 所有 6 个 API 端点
- 🗣️ 英文和中文切换
- 📝 用户偏好持久化
- 🎮 NPC 对话本地化
- 📋 任务描述本地化

---

## 📊 已实现的功能

### 1️⃣ 新 API 端点（共 6 个）

| 端点 | 方法 | 功能 |
|------|------|------|
| `/api/language/list` | GET | 获取所有支持的语言 |
| `/api/language/strings` | GET | 获取 UI 字符串翻译（50 项） |
| `/api/language/select` | POST | 设置用户语言偏好 |
| `/api/language/current` | GET | 获取用户当前语言 |
| `/api/language/npc-dialogue` | GET | 获取本地化 NPC 对话 |
| `/api/language/quest` | GET | 获取本地化任务描述 |

### 2️⃣ 多语言覆盖

- ✅ **UI 字符串**: 50 个菜单和界面文本
- ✅ **NPC 对话**: 10 个 NPC，28 条对话
- ✅ **任务描述**: 13 个任务的完整描述
- ✅ **支持语言**: 英文 (en) 和 中文 (zh)

### 3️⃣ 核心特性

- ✅ 用户语言偏好持久化
- ✅ 默认语言设置为英文
- ✅ 完全本地化的 NPC 对话
- ✅ 完全本地化的任务系统
- ✅ 自动回退处理（缺失翻译时显示英文）

---

## 📂 项目文件结构

### 新增文件

```
/root/petskill/
├── skill.md                                    # ✅ 已更新：新增多语言文档
├── MULTILINGUAL_TESTING_GUIDE.md               # ✅ 详细测试指南
├── quick_test_multilingual.sh                  # ✅ 快速测试脚本
├── test_multilingual_system.sh                 # ✅ 完整测试套件
├── complete_chinese_translations.sql           # ✅ SQL 翻译数据
└── judge-server/
    ├── internal/service/language_service.go    # ✅ 新增：语言服务
    └── internal/handler/language_handlers.go   # ✅ 新增：API 处理器
```

### 修改的文件

```
judge-server/
├── internal/handler/handlers.go                # ✅ 修改：添加 languageService
├── cmd/main.go                                 # ✅ 修改：注册 6 个新路由
```

---

## 🔧 数据库设置（可选但推荐）

如果你想添加完整的中文翻译数据到数据库（使 NPC 对话和任务完全本地化）：

```bash
psql -h localhost -U postgres -d agent_monster -f /root/petskill/complete_chinese_translations.sql
```

这将：
- 添加 20 条 NPC 对话的中文翻译
- 添加 13 个任务的中文翻译
- 验证所有数据已正确插入

---

## 🎯 测试场景

### 场景 1：用户切换到中文

```bash
# 第 1 步：设置语言为中文
curl -X POST "http://pokemon.openx.pro:10000/api/language/select" \
  -H "Content-Type: application/json" \
  -d '{"user_id":"chengjia2016", "language_code":"zh"}'

# 第 2 步：获取中文 NPC 对话
curl "http://pokemon.openx.pro:10000/api/language/npc-dialogue?npc_id=1&language=zh"

# 期望结果：
# "dialogue": "欢迎来到华蓝道馆！我是馆主米斯蒂，钢铁属性宝可梦的训练大师。"
```

### 场景 2：查看本地化任务

```bash
curl "http://pokemon.openx.pro:10000/api/language/quest?quest_id=1&language=zh"

# 期望结果：
# {
#   "quest_name": "华蓝道馆挑战",
#   "description": "在华蓝市击败米斯蒂道馆馆主并获得钢铁徽章",
#   "reward_item": "钢铁徽章"
# }
```

### 场景 3：动态语言切换

```bash
# 先设置中文
curl -X POST "http://pokemon.openx.pro:10000/api/language/select" \
  -H "Content-Type: application/json" \
  -d '{"user_id":"chengjia2016", "language_code":"zh"}'

# 验证是中文
curl "http://pokemon.openx.pro:10000/api/language/current?user_id=chengjia2016"
# 返回: {"language_code": "zh"}

# 切换回英文
curl -X POST "http://pokemon.openx.pro:10000/api/language/select" \
  -H "Content-Type: application/json" \
  -d '{"user_id":"chengjia2016", "language_code":"en"}'

# 验证是英文
curl "http://pokemon.openx.pro:10000/api/language/current?user_id=chengjia2016"
# 返回: {"language_code": "en"}
```

---

## ✅ 验证清单

在完成测试后，请确保：

- [ ] 快速测试脚本成功运行 (8 个输出)
- [ ] 完整测试套件通过所有 15 个测试
- [ ] 能够在英文和中文之间切换
- [ ] NPC 对话正确显示中文内容
- [ ] 任务描述正确显示中文内容
- [ ] 用户语言偏好正确保存
- [ ] skill.md 在 GitHub 上包含新的多语言文档

---

## 📊 测试结果示例

### 快速测试输出示例

```
🚀 Agent Monster - Quick Multilingual Test
Testing User: chengjia2016
Server: http://pokemon.openx.pro:10000

1️⃣ Testing: Get available languages
[
  { "code": "en", "name": "English" },
  { "code": "zh", "name": "中文" }
]

2️⃣ Testing: Set language to Chinese
{ "success": true, "language_code": "zh" }

3️⃣ Testing: Get UI strings in Chinese (first 5)
{ "key": "menu.main", "value": "主菜单" }
{ "key": "menu.quest", "value": "任务" }
...

✅ Quick test complete!
```

---

## 🐛 故障排除

### 问题：服务器无法连接
```
解决方案：
1. 确保 Judge Server 在运行
2. 检查地址是否正确：http://pokemon.openx.pro:10000
3. 运行健康检查：curl http://pokemon.openx.pro:10000/health
```

### 问题：收到 404 错误
```
解决方案：
1. 确保 judge-server 已编译并部署最新代码
2. 检查 API 端点名称是否正确（区分大小写）
3. 验证请求方法是否正确 (GET vs POST)
```

### 问题：中文翻译显示为空
```
解决方案：
1. 运行 SQL 脚本以导入翻译数据：
   psql -h localhost -U postgres -d agent_monster -f complete_chinese_translations.sql
2. 验证数据已导入：
   psql -h localhost -U postgres -d agent_monster -c "SELECT COUNT(*) FROM npc_dialogues WHERE dialogue_text_zh IS NOT NULL"
```

---

## 📝 GitHub 提交记录

```
commit 9dec305 - Add SQL script to populate Chinese translations
commit fceddb6 - Add comprehensive multilingual system test suite
commit 4a12860 - Add multilingual system documentation to skill.md
```

所有提交都已推送到 GitHub。你可以在这里查看：
https://github.com/chengjia2016/agent-pokemon/commits/master

---

## 🎮 下一步

测试成功后，你可以：

1. **在游戏中使用多语言**: 让玩家在开始游戏时选择语言
2. **扩展语言支持**: 添加更多语言（日文、西班牙文等）
3. **改进翻译**: 通过社区反馈改进现有翻译
4. **添加动态翻译**: 支持运行时添加新的翻译

---

## 📞 技术支持

如有任何问题，请：

1. 查看 `MULTILINGUAL_TESTING_GUIDE.md` 获取详细说明
2. 检查 `/root/petskill/judge-server/internal/service/language_service.go` 源代码
3. 查看 API 响应中的 `success` 和 `error` 字段

---

## 🏆 项目统计

| 项目 | 数量 |
|------|------|
| 新增 API 端点 | 6 |
| 支持语言 | 2 (English, 中文) |
| UI 字符串翻译 | 50 |
| NPC 对话翻译 | 28 |
| 任务翻译 | 13 |
| 测试用例 | 15 |
| 新增源文件 | 2 |
| 修改源文件 | 2 |
| GitHub 提交 | 3 |

---

## 🎉 完成！

多语言系统已完全实现并准备测试。祝你测试愉快！

**最后更新**: 2026-04-14  
**测试用户**: chengjia2016  
**服务器**: http://pokemon.openx.pro:10000  
**文档**: https://github.com/chengjia2016/agent-pokemon
