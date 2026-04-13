#!/bin/bash

# Agent Monster - 在 Agent CLI 中的示例游戏流程
# 这个脚本展示了 Claude/Gemini/OpenCode 如何游戏

echo "╔════════════════════════════════════════════════════════════════════════════╗"
echo "║           🎮 Agent Monster - Agent CLI 游戏演示 🎮                        ║"
echo "╚════════════════════════════════════════════════════════════════════════════╝"
echo ""

SERVER="http://localhost:10000"
GITHUB_ID=333444555
USERNAME="agent_player"
EMAIL="agent@github.com"

# 模拟 Agent 的思考过程
echo "🤖 Agent 思考过程："
echo "玩家想要在 Agent Monster 中冒险。"
echo "我需要:"
echo "1. 创建账户"
echo "2. 查询游戏内容"
echo "3. 以沉浸式的方式呈现内容"
echo ""
sleep 2

# 第一步：创建账户
echo "════════════════════════════════════════════════════════════════════════════"
echo "📍 步骤 1: 创建玩家账户"
echo "════════════════════════════════════════════════════════════════════════════"
echo ""

echo "API 调用:"
echo "curl -X POST \"$SERVER/api/users/create\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{\"github_id\": $GITHUB_ID, \"github_login\": \"$USERNAME\", \"email\": \"$EMAIL\"}'"
echo ""

ACCOUNT=$(curl -s -X POST "$SERVER/api/users/create" \
  -H "Content-Type: application/json" \
  -d "{\"github_id\": $GITHUB_ID, \"github_login\": \"$USERNAME\", \"email\": \"$EMAIL\"}")

echo "API 响应:"
echo "$ACCOUNT" | jq '.'
echo ""

USER_ID=$(echo "$ACCOUNT" | jq '.user.id')
BALANCE=$(echo "$ACCOUNT" | jq '.user.balance')

echo "✨ Agent 的沉浸式呈现:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "        欢迎来到 Kanto 地区，年轻的训练师！"
echo ""
echo "        一个年长的教授对你说："
echo "        '我感受到了你身上的力量。这是我为新手准备的礼物。'"
echo ""
echo "✨ 你获得了启动奖励包！"
echo "   • 500 初始金币 💰"
echo "   • 新手训练师的祝福"
echo "   • 这个 Kanto 地区的完整地图"
echo ""
echo "   用户 ID: $USER_ID"
echo "   余额: $BALANCE 💰"
echo ""
sleep 3

# 第二步：查询任务
echo "════════════════════════════════════════════════════════════════════════════"
echo "📍 步骤 2: Agent 查询可用的任务"
echo "════════════════════════════════════════════════════════════════════════════"
echo ""

echo "API 调用:"
echo "curl \"$SERVER/api/quests\" | jq '.'"
echo ""

QUESTS=$(curl -s "$SERVER/api/quests")

echo "✨ Agent 的沉浸式呈现:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "        传说的任务公告板出现在你面前..."
echo ""
echo "📜 主线任务 (MAIN):"
echo ""
echo "$QUESTS" | jq -r '.quests[] | select(.type=="main") | "   [\(.id)] 🎯 \(.name_en)\n       ⭐\(.difficulty) 难度 | 奖励: \(.reward_coins)💰 + \(.reward_exp)XP"' | head -20
echo ""
echo "📋 支线任务 (SIDE):"
echo ""
echo "$QUESTS" | jq -r '.quests[] | select(.type=="side") | "   [\(.id)] 📋 \(.name_en)\n       ⭐\(.difficulty) 难度 | 奖励: \(.reward_coins)💰 + \(.reward_exp)XP"'
echo ""
sleep 3

# 第三步：查询地图
echo "════════════════════════════════════════════════════════════════════════════"
echo "📍 步骤 3: Agent 查询可探索的地图"
echo "════════════════════════════════════════════════════════════════════════════"
echo ""

echo "API 调用:"
echo "curl \"$SERVER/api/map/zones?island_id=island_1\" | jq '.'"
echo ""

ZONES=$(curl -s "$SERVER/api/map/zones?island_id=island_1")

echo "✨ Agent 的沉浸式呈现:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "        一张古老的羊皮地图展开在你面前..."
echo ""
echo "🗺️  Kanto 地区 - 可探索的区域:"
echo ""
echo "$ZONES" | jq -r '.zones[] | "   🌍 \(.name_en) (\(.name_zh))\n      类型: \(.type) | ⭐\(.level) 难度\n      \(.description)\n"'
echo ""
sleep 3

# 第四步：查询地下城
echo "════════════════════════════════════════════════════════════════════════════"
echo "📍 步骤 4: Agent 查询可挑战的地下城"
echo "════════════════════════════════════════════════════════════════════════════"
echo ""

echo "API 调用:"
echo "curl \"$SERVER/api/dungeons\" | jq '.'"
echo ""

DUNGEONS=$(curl -s "$SERVER/api/dungeons")

echo "✨ Agent 的沉浸式呈现:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "        远处传来神秘的声音..."
echo ""
echo "🏰 地下城冒险 - 多层副本等待你:"
echo ""
echo "$DUNGEONS" | jq -r '.dungeons[] | "   🏞️  \(.name_en) (\(.name_zh))\n      难度: ⭐\(.difficulty) | \(.floor_count) 层\n      奖励: \(.reward_coins)💰 + \(.reward_exp)XP\n"'
echo ""
sleep 3

# 第五步：查询道馆
echo "════════════════════════════════════════════════════════════════════════════"
echo "📍 步骤 5: Agent 查询可挑战的道馆"
echo "════════════════════════════════════════════════════════════════════════════"
echo ""

echo "API 调用:"
echo "curl \"$SERVER/api/gyms?town_id=town_1\" | jq '.'"
echo ""

GYMS=$(curl -s "$SERVER/api/gyms?town_id=town_1")

echo "✨ Agent 的沉浸式呈现:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "        一座宏伟的建筑矗立在你面前..."
echo ""
echo "🏆 道馆挑战 - 成为宝可梦大师的途径:"
echo ""
echo "$GYMS" | jq -r '.gyms[] | "   🏛️  \(.name_en) (\(.name_zh))\n      类型: \(.type_focus) | 徽章: \(.badge_name)\n      状态: 🔓 可挑战\n"'
echo ""
sleep 3

# 第六步：Agent 的建议
echo "════════════════════════════════════════════════════════════════════════════"
echo "📍 步骤 6: Agent 分析并建议"
echo "════════════════════════════════════════════════════════════════════════════"
echo ""

echo "✨ Agent 的分析:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "基于我查询的游戏数据，我为你制定了一个优化的游戏进程："
echo ""
echo "🎯 第一天："
echo "   1. 接受任务 [1] '捕捉你的第一只宝可梦' (容易)"
echo "   2. 前往 Route 1 探索"
echo "   3. 完成任务 → 获得 50 💰 + 100 XP"
echo "   预计耗时: 10 分钟"
echo ""
echo "🎯 第二天："
echo "   1. 接受任务 [3] '收集 5 只宝可梦' (普通)"
echo "   2. 探索 Viridian Forest"
echo "   3. 积累经验和金币"
echo "   预计耗时: 20 分钟"
echo ""
echo "🎯 第三天："
echo "   1. 进入 Viridian Forest 地下城 (容易)"
echo "   2. 获得 200 💰 + 500 XP (5 倍效率！)"
echo "   3. 准备挑战馆主"
echo "   预计耗时: 15 分钟"
echo ""
echo "💡 长期目标："
echo "   • 完成所有 3 个任务 → 总共 350 💰 + 900 XP"
echo "   • 清除所有 3 个地下城 → 总共 1400 💰 + 3500 XP"
echo "   • 挑战 Pewter City Gym → 获得 Boulder Badge 🏆"
echo ""
echo "你同意这个计划吗？你想先做什么？"
echo ""
sleep 3

# 最后的总结
echo "════════════════════════════════════════════════════════════════════════════"
echo "✅ Agent CLI 游戏演示完成"
echo "════════════════════════════════════════════════════════════════════════════"
echo ""
echo "📚 这演示展示了:"
echo "  ✓ 如何创建账户 (POST /api/users/create)"
echo "  ✓ 如何查询任务 (GET /api/quests)"
echo "  ✓ 如何查询地图 (GET /api/map/zones)"
echo "  ✓ 如何查询地下城 (GET /api/dungeons)"
echo "  ✓ 如何查询道馆 (GET /api/gyms)"
echo "  ✓ 如何以沉浸式的方式呈现内容"
echo "  ✓ 如何基于数据制定游戏策略"
echo ""
echo "🤖 Agent 的工作流程:"
echo "  1. 接收用户输入"
echo "  2. 调用相关 API 获取数据"
echo "  3. 解析 JSON 响应"
echo "  4. 用故事性的语言呈现"
echo "  5. 建议下一步行动"
echo "  6. 等待用户响应"
echo ""
echo "📖 更详细的指南："
echo "  👉 /root/petskill/AGENT_CLI_GAMEPLAY.md"
echo ""
echo "🎮 现在你可以在任何 Agent CLI (Claude, Gemini, OpenCode) 中玩这个游戏了！"
echo ""
