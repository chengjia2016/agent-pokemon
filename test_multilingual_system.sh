#!/bin/bash

# ============================================================
# Agent Monster - Multilingual System Test Suite
# Test with chengjia2016 user account
# ============================================================

BASE_URL="http://pokemon.openx.pro:10000"
TEST_USER="chengjia2016"

echo "╔════════════════════════════════════════════════════════════╗"
echo "║     Agent Monster - Multilingual System Testing            ║"
echo "║     Testing User: $TEST_USER"
echo "║     Server: $BASE_URL"
echo "╚════════════════════════════════════════════════════════════╝"
echo ""

# Color codes for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Counter for tests
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# Test function
run_test() {
    local test_name=$1
    local test_number=$2
    local method=$3
    local endpoint=$4
    local data=$5
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${YELLOW}Test $test_number: $test_name${NC}"
    echo "Method: $method"
    echo "Endpoint: $endpoint"
    
    if [ -z "$data" ]; then
        response=$(curl -s -X "$method" "$BASE_URL$endpoint")
    else
        echo "Data: $data"
        response=$(curl -s -X "$method" "$BASE_URL$endpoint" \
          -H "Content-Type: application/json" \
          -d "$data")
    fi
    
    echo "Response:"
    echo "$response" | jq . 2>/dev/null || echo "$response"
    
    # Check if response contains success: true
    if echo "$response" | jq . 2>/dev/null | grep -q '"success".*true'; then
        echo -e "${GREEN}✅ PASSED${NC}"
        PASSED_TESTS=$((PASSED_TESTS + 1))
    else
        echo -e "${RED}❌ FAILED${NC}"
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi
    echo ""
}

# ============================================================
# Test Suite
# ============================================================

# Test 1: Get available languages
run_test "Get Available Languages" "1" "GET" "/api/language/list"

# Test 2: Get UI strings in English
run_test "Get UI Strings (English)" "2" "GET" "/api/language/strings?language=en"

# Test 3: Get UI strings in Chinese
run_test "Get UI Strings (Chinese)" "3" "GET" "/api/language/strings?language=zh"

# Test 4: Set user language to Chinese
run_test "Set User Language to Chinese" "4" "POST" "/api/language/select" \
  "{\"user_id\": \"$TEST_USER\", \"language_code\": \"zh\"}"

# Test 5: Verify user language preference (should be Chinese)
run_test "Get User Current Language (Chinese)" "5" "GET" "/api/language/current?user_id=$TEST_USER"

# Test 6: Get NPC dialogue in Chinese
run_test "Get NPC Dialogue (Chinese - NPC 1)" "6" "GET" "/api/language/npc-dialogue?npc_id=1&language=zh"

# Test 7: Get NPC dialogue in English
run_test "Get NPC Dialogue (English - NPC 1)" "7" "GET" "/api/language/npc-dialogue?npc_id=1&language=en"

# Test 8: Get NPC dialogue for different NPC (Chinese)
run_test "Get NPC Dialogue (Chinese - NPC 9 Professor Oak)" "8" "GET" "/api/language/npc-dialogue?npc_id=9&language=zh"

# Test 9: Get quest in Chinese
run_test "Get Quest Description (Chinese - Quest 1)" "9" "GET" "/api/language/quest?quest_id=1&language=zh"

# Test 10: Get quest in English
run_test "Get Quest Description (English - Quest 1)" "10" "GET" "/api/language/quest?quest_id=1&language=en"

# Test 11: Get multiple quests in Chinese
run_test "Get Quest Description (Chinese - Quest 5)" "11" "GET" "/api/language/quest?quest_id=5&language=zh"

# Test 12: Switch user language back to English
run_test "Set User Language to English" "12" "POST" "/api/language/select" \
  "{\"user_id\": \"$TEST_USER\", \"language_code\": \"en\"}"

# Test 13: Verify user language preference (should be English)
run_test "Get User Current Language (English)" "13" "GET" "/api/language/current?user_id=$TEST_USER"

# Test 14: Get NPC 2 dialogue in Chinese (after switching back to English)
run_test "Get NPC Dialogue (Chinese - NPC 2)" "14" "GET" "/api/language/npc-dialogue?npc_id=2&language=zh"

# Test 15: Get NPC 10 dialogue (Officer Jenny) in Chinese
run_test "Get NPC Dialogue (Chinese - NPC 10 Officer Jenny)" "15" "GET" "/api/language/npc-dialogue?npc_id=10&language=zh"

# ============================================================
# Test Summary
# ============================================================

echo ""
echo "╔════════════════════════════════════════════════════════════╗"
echo "║                     TEST SUMMARY                           ║"
echo "╠════════════════════════════════════════════════════════════╣"
echo -e "║ Total Tests:  ${YELLOW}$TOTAL_TESTS${NC}"
echo -e "║ Passed:       ${GREEN}$PASSED_TESTS${NC}"
echo -e "║ Failed:       ${RED}$FAILED_TESTS${NC}"
echo "╠════════════════════════════════════════════════════════════╣"

if [ $FAILED_TESTS -eq 0 ]; then
    echo -e "║ ${GREEN}✅ ALL TESTS PASSED!${NC}                                      ║"
else
    echo -e "║ ${RED}❌ SOME TESTS FAILED${NC}                                     ║"
fi

echo "╚════════════════════════════════════════════════════════════╝"
echo ""

# Exit with appropriate code
[ $FAILED_TESTS -eq 0 ] && exit 0 || exit 1
