#!/bin/bash

# World System API Integration Test Script
# This script tests the world system endpoints

API_BASE_URL="http://localhost:8080"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "================================"
echo "World System Integration Tests"
echo "================================"

# Function to test endpoint
test_endpoint() {
    local method=$1
    local endpoint=$2
    local data=$3
    local expected_status=$4
    
    echo -e "\n${YELLOW}Testing: $method $endpoint${NC}"
    
    if [ "$method" == "POST" ] || [ "$method" == "PUT" ]; then
        response=$(curl -s -w "\n%{http_code}" -X "$method" \
            -H "Content-Type: application/json" \
            -d "$data" \
            "$API_BASE_URL$endpoint")
    else
        response=$(curl -s -w "\n%{http_code}" -X "$method" \
            "$API_BASE_URL$endpoint")
    fi
    
    status_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n-1)
    
    if [ "$status_code" = "$expected_status" ]; then
        echo -e "${GREEN}✓ Success (HTTP $status_code)${NC}"
        echo "Response: $(echo "$body" | head -c 200)"
    else
        echo -e "${RED}✗ Failed (Expected $expected_status, Got $status_code)${NC}"
        echo "Response: $body"
    fi
}

# ==================== NPC Tests ====================
echo -e "\n\n${YELLOW}=== NPC Endpoints ===${NC}"

# Get NPCs from town 1
test_endpoint "GET" "/api/npcs?town_id=1" "" "200"

# Create NPC
test_endpoint "POST" "/api/npcs" \
    '{
        "town_id": "1",
        "name_en": "Test NPC",
        "name_zh": "测试NPC",
        "type": "NPC",
        "role": "quest_giver",
        "coord_x": 50.0,
        "coord_y": 50.0,
        "avatar_url": "https://example.com/npc.png"
    }' \
    "201"

# ==================== Quest Tests ====================
echo -e "\n\n${YELLOW}=== Quest Endpoints ===${NC}"

# Get quests
test_endpoint "GET" "/api/quests" "" "200"

# Create quest (requires NPC to exist first)
test_endpoint "POST" "/api/quests" \
    '{
        "quest_giver_id": "1",
        "title_en": "Test Quest",
        "title_zh": "测试任务",
        "description_en": "A test quest",
        "description_zh": "一个测试任务",
        "quest_type": "capture",
        "reward_gold": 100,
        "reward_exp": 500
    }' \
    "201"

# Get user quests
test_endpoint "GET" "/api/user/quests?user_id=test_user" "" "200"

# Accept quest
test_endpoint "POST" "/api/user/quests?user_id=test_user&quest_id=1&action=accept" "" "200"

# ==================== Dungeon Tests ====================
echo -e "\n\n${YELLOW}=== Dungeon Endpoints ===${NC}"

# Create dungeon
test_endpoint "POST" "/api/dungeons" \
    '{
        "name_en": "Test Dungeon",
        "name_zh": "测试地牢",
        "description_en": "A test dungeon",
        "description_zh": "一个测试地牢",
        "difficulty_level": "easy",
        "boss_name_en": "Test Boss",
        "boss_name_zh": "测试Boss",
        "boss_level": 15,
        "reward_gold": 500
    }' \
    "201"

# ==================== Gym Tests ====================
echo -e "\n\n${YELLOW}=== Gym Endpoints ===${NC}"

# Get user gym badges
test_endpoint "GET" "/api/gyms?user_id=test_user" "" "200"

# Create gym (requires NPC to exist first)
test_endpoint "POST" "/api/gyms" \
    '{
        "gym_leader_id": "1",
        "gym_name_en": "Test Gym",
        "gym_name_zh": "测试健身房",
        "city_location": "Test City",
        "gym_badge_en": "Test Badge",
        "gym_badge_zh": "测试勋章",
        "difficulty_level": "normal",
        "reward_gold": 1000
    }' \
    "201"

# ==================== Map Zone & Level Tests ====================
echo -e "\n\n${YELLOW}=== Map Zone & Level Endpoints ===${NC}"

# Create map zone
test_endpoint "POST" "/api/map/zones" \
    '{
        "zone_name_en": "Test Zone",
        "zone_name_zh": "测试区域",
        "region_id": "1",
        "zone_type": "forest",
        "difficulty_level": "beginner",
        "coord_x": 50.0,
        "coord_y": 50.0
    }' \
    "201"

# Get levels in a zone
test_endpoint "GET" "/api/levels?zone_id=1" "" "200"

# Create level
test_endpoint "POST" "/api/levels" \
    '{
        "zone_id": "1",
        "level_number": 1,
        "level_name_en": "Test Level",
        "level_name_zh": "测试关卡",
        "difficulty": "easy",
        "reward_gold": 100,
        "reward_exp": 200
    }' \
    "201"

# Complete level
test_endpoint "POST" "/api/user/levels/progress?user_id=test_user&level_id=1" "" "200"

echo -e "\n\n${YELLOW}=== All Tests Completed ===${NC}"
