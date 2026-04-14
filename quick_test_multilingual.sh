#!/bin/bash

# Quick test for chengjia2016 to verify multilingual system works

echo "🚀 Agent Monster - Quick Multilingual Test"
echo "Testing User: chengjia2016"
echo "Server: http://pokemon.openx.pro:10000"
echo ""

BASE_URL="http://pokemon.openx.pro:10000"
USER="chengjia2016"

echo "1️⃣ Testing: Get available languages"
curl -s "$BASE_URL/api/language/list" | jq '.languages[] | {code, name}'
echo ""

echo "2️⃣ Testing: Set language to Chinese"
curl -s -X POST "$BASE_URL/api/language/select" \
  -H "Content-Type: application/json" \
  -d "{\"user_id\":\"$USER\",\"language_code\":\"zh\"}" | jq '{success, language_code: .language_code}'
echo ""

echo "3️⃣ Testing: Get UI strings in Chinese (first 5)"
curl -s "$BASE_URL/api/language/strings?language=zh" | jq '.strings | to_entries | .[0:5] | .[] | {key, value}'
echo ""

echo "4️⃣ Testing: Get NPC dialogue in Chinese (Gym Leader Misty)"
curl -s "$BASE_URL/api/language/npc-dialogue?npc_id=1&language=zh" | jq '{npc_id, language, dialogue}'
echo ""

echo "5️⃣ Testing: Get quest in Chinese (Quest 1)"
curl -s "$BASE_URL/api/language/quest?quest_id=1&language=zh" | jq '.quest | {quest_name, description, reward_item}'
echo ""

echo "6️⃣ Testing: Get current user language"
curl -s "$BASE_URL/api/language/current?user_id=$USER" | jq '{user_id, language_code}'
echo ""

echo "7️⃣ Testing: Switch to English"
curl -s -X POST "$BASE_URL/api/language/select" \
  -H "Content-Type: application/json" \
  -d "{\"user_id\":\"$USER\",\"language_code\":\"en\"}" | jq '{success, language_code: .language_code}'
echo ""

echo "8️⃣ Testing: Get NPC dialogue in English (Gym Leader Misty)"
curl -s "$BASE_URL/api/language/npc-dialogue?npc_id=1&language=en" | jq '{npc_id, language, dialogue}'
echo ""

echo "✅ Quick test complete!"
