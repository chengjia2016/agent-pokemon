#!/bin/bash

BASE_URL="http://127.0.0.1:10000"

echo "=== Agent Monster Judge Server Test ==="
echo ""

echo "1. Testing health endpoint..."
curl -s "$BASE_URL/health" | python3 -m json.tool
echo ""

echo "2. Validating pet..."
PET_JSON='{
  "monster_id": "test_pet_002",
  "name": "战斗宠物",
  "owner": "fighter@example.com",
  "species": "Dragon",
  "generation": 1,
  "evolution_stage": 1,
  "base_stats": {"hp": 100, "attack": 80, "defense": 60, "speed": 70, "armor": 50, "quota": 120},
  "exp": 1000,
  "level": 10,
  "genes": {"logic": {"weight": 0.5}, "speed": {"weight": 0.5}},
  "signature": {"algorithm": "RSA-SHA256", "value": ""}
}'
curl -s -X POST "$BASE_URL/api/pet/validate" \
  -H "Content-Type: application/json" \
  -d "$PET_JSON" | python3 -m json.tool
echo ""

echo "3. Recording food consumption..."
FOOD_JSON='{
  "player": "fighter@example.com",
  "pet_id": "test_pet_002",
  "food_type": "cookie",
  "effect": "+10 EXP",
  "amount": 1
}'
curl -s -X POST "$BASE_URL/api/food/record" \
  -H "Content-Type: application/json" \
  -d "$FOOD_JSON" | python3 -m json.tool
echo ""

echo "4. Recording growth..."
GROWTH_JSON='{
  "pet_id": "test_pet_002",
  "old_level": 9,
  "new_level": 10,
  "old_exp": 900,
  "new_exp": 1000,
  "source": "battle_win"
}'
curl -s -X POST "$BASE_URL/api/growth/record" \
  -H "Content-Type: application/json" \
  -d "$GROWTH_JSON" | python3 -m json.tool
echo ""

echo "5. Getting leaderboard..."
curl -s "$BASE_URL/api/leaderboard" | python3 -m json.tool
echo ""

echo "=== Test Complete ==="
