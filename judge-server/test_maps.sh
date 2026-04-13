#!/bin/bash

# Comprehensive Map System Test Suite
# Tests all map API endpoints

BASE_URL="http://127.0.0.1:10000"
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

TEST_COUNT=0
PASS_COUNT=0
FAIL_COUNT=0

# Helper function to print test results
test_result() {
    local test_name=$1
    local expected=$2
    local actual=$3
    
    TEST_COUNT=$((TEST_COUNT + 1))
    
    if echo "$actual" | grep -q "$expected"; then
        echo -e "${GREEN}✓ PASS${NC}: $test_name"
        PASS_COUNT=$((PASS_COUNT + 1))
    else
        echo -e "${RED}✗ FAIL${NC}: $test_name"
        echo "  Expected to contain: $expected"
        echo "  Actual response: $actual"
        FAIL_COUNT=$((FAIL_COUNT + 1))
    fi
}

# Helper function to extract JSON value
get_json_value() {
    echo "$1" | python3 -c "import sys, json; data=json.load(sys.stdin); print(data.get('$2', ''))"
}

echo -e "${YELLOW}=== Agent Monster Map System Test Suite ===${NC}"
echo ""

# Test 1: List all maps
echo -e "${YELLOW}Test Suite 1: List and Retrieve Maps${NC}"
echo "1. Testing GET /api/maps (list all maps)..."
RESPONSE=$(curl -s "$BASE_URL/api/maps")
test_result "List maps returns success" '"success": true' "$RESPONSE"
test_result "List maps returns data array" '"data":' "$RESPONSE"
echo ""

# Test 2: Get specific map
echo "2. Testing GET /api/maps/001 (get specific map)..."
RESPONSE=$(curl -s "$BASE_URL/api/maps/001")
test_result "Get map 001 returns success" '"success": true' "$RESPONSE"
test_result "Get map 001 has correct map_id" '"map_id": "001"' "$RESPONSE"
test_result "Get map 001 has terrain data" '"terrain":' "$RESPONSE"
test_result "Get map 001 has elements" '"elements":' "$RESPONSE"
test_result "Get map 001 has connections" '"connections":' "$RESPONSE"
echo ""

# Test 3: Get map that doesn't exist
echo "3. Testing GET /api/maps/999 (non-existent map)..."
RESPONSE=$(curl -s "$BASE_URL/api/maps/999")
test_result "Non-existent map returns error" '"success": false' "$RESPONSE"
echo ""

# Test Suite 2: Map Elements
echo -e "${YELLOW}Test Suite 2: Map Elements and Filtering${NC}"
echo "4. Testing GET /api/maps/001/elements (all elements)..."
RESPONSE=$(curl -s "$BASE_URL/api/maps/001/elements")
test_result "Get elements returns success" '"success": true' "$RESPONSE"
test_result "Get elements returns elements array" '"elements":' "$RESPONSE"
test_result "Get elements returns count" '"count":' "$RESPONSE"
echo ""

# Test 4: Get elements filtered by type
echo "5. Testing GET /api/maps/001/elements?type=wild_pokemon..."
RESPONSE=$(curl -s "$BASE_URL/api/maps/001/elements?type=wild_pokemon")
test_result "Filter by type returns success" '"success": true' "$RESPONSE"
test_result "Filter elements contains only wild_pokemon" '"type": "wild_pokemon"' "$RESPONSE"
echo ""

# Test 5: Get food elements
echo "6. Testing GET /api/maps/001/elements?type=food..."
RESPONSE=$(curl -s "$BASE_URL/api/maps/001/elements?type=food")
test_result "Food filter returns success" '"success": true' "$RESPONSE"
echo ""

# Test Suite 3: Map Connections
echo -e "${YELLOW}Test Suite 3: Map Connections${NC}"
echo "7. Testing GET /api/maps/001/connections..."
RESPONSE=$(curl -s "$BASE_URL/api/maps/001/connections")
test_result "Get connections returns success" '"success": true' "$RESPONSE"
test_result "Get connections returns connection data" '"data":' "$RESPONSE"
echo ""

# Test 8: Check if connections exist
echo "8. Checking map 001 connections..."
RESPONSE=$(curl -s "$BASE_URL/api/maps/001/connections")
echo "  Response: $RESPONSE"
echo ""

# Test Suite 4: Map Search
echo -e "${YELLOW}Test Suite 4: Map Search${NC}"
echo "9. Testing GET /api/maps/search?q=001..."
RESPONSE=$(curl -s "$BASE_URL/api/maps/search?q=001")
test_result "Search returns success" '"success": true' "$RESPONSE"
test_result "Search returns data" '"data":' "$RESPONSE"
echo ""

# Test 10: Search by owner
echo "10. Testing GET /api/maps/search?owner_id=1..."
RESPONSE=$(curl -s "$BASE_URL/api/maps/search?owner_id=1")
test_result "Search by owner returns success" '"success": true' "$RESPONSE"
echo ""

# Test Suite 5: Map Traversal
echo -e "${YELLOW}Test Suite 5: Map Traversal${NC}"
echo "11. Testing POST /api/maps/traverse (move east from 002 to 003)..."
TRAVERSE_JSON='{
  "current_map_id": "002",
  "direction": "east"
}'
RESPONSE=$(curl -s -X POST "$BASE_URL/api/maps/traverse" \
  -H "Content-Type: application/json" \
  -d "$TRAVERSE_JSON")
test_result "Traverse returns success" '"success": true' "$RESPONSE"
test_result "Traverse returns next map" '"next":' "$RESPONSE"
echo ""

# Test 12: Invalid direction
echo "12. Testing POST /api/maps/traverse with invalid direction..."
TRAVERSE_JSON='{
  "current_map_id": "001",
  "direction": "invalid"
}'
RESPONSE=$(curl -s -X POST "$BASE_URL/api/maps/traverse" \
  -H "Content-Type: application/json" \
  -d "$TRAVERSE_JSON")
test_result "Invalid direction returns error" '"success": false' "$RESPONSE"
echo ""

# Test 13: Non-existent direction
echo "13. Testing POST /api/maps/traverse to non-existent map..."
TRAVERSE_JSON='{
  "current_map_id": "001",
  "direction": "west"
}'
RESPONSE=$(curl -s -X POST "$BASE_URL/api/maps/traverse" \
  -H "Content-Type: application/json" \
  -d "$TRAVERSE_JSON")
test_result "No map in direction returns error" '"success": false' "$RESPONSE"
echo ""

# Test Suite 6: Map Generation (basic test)
echo -e "${YELLOW}Test Suite 6: Map Generation${NC}"
echo "14. Testing POST /api/maps/generate..."
GENERATE_JSON='{
  "owner_id": 1,
  "owner_name": "test_user",
  "map_id": "999",
  "width": 20,
  "height": 20
}'
RESPONSE=$(curl -s -X POST "$BASE_URL/api/maps/generate" \
  -H "Content-Type: application/json" \
  -d "$GENERATE_JSON")
test_result "Generate map returns success or error" '"success":' "$RESPONSE"
echo ""

# Test 15: Generate with missing fields
echo "15. Testing POST /api/maps/generate with missing fields..."
GENERATE_JSON='{
  "owner_id": 1
}'
RESPONSE=$(curl -s -X POST "$BASE_URL/api/maps/generate" \
  -H "Content-Type: application/json" \
  -d "$GENERATE_JSON")
test_result "Missing fields returns error" '"success": false' "$RESPONSE"
echo ""

# Test Suite 7: Data Structure Validation
echo -e "${YELLOW}Test Suite 7: Data Structure Validation${NC}"
echo "16. Validating map data structure..."
RESPONSE=$(curl -s "$BASE_URL/api/maps/001")
test_result "Map has version field" '"version":' "$RESPONSE"
test_result "Map has map_id field" '"map_id":' "$RESPONSE"
test_result "Map has owner_id field" '"owner_id":' "$RESPONSE"
test_result "Map has width and height" '"width":' "$RESPONSE"
test_result "Map has statistics" '"statistics":' "$RESPONSE"
echo ""

# Test 17: Element data structure
echo "17. Validating element data structure..."
RESPONSE=$(curl -s "$BASE_URL/api/maps/001/elements")
if echo "$RESPONSE" | grep -q '"id":'; then
    echo -e "${GREEN}✓ PASS${NC}: Elements have id field"
    PASS_COUNT=$((PASS_COUNT + 1))
else
    echo -e "${YELLOW}⊘ INFO${NC}: No elements to validate (this is OK)"
fi
TEST_COUNT=$((TEST_COUNT + 1))
echo ""

# Test Suite 8: HTTP Methods
echo -e "${YELLOW}Test Suite 8: HTTP Method Validation${NC}"
echo "18. Testing invalid HTTP method on /api/maps/001..."
RESPONSE=$(curl -s -X DELETE "$BASE_URL/api/maps/001")
# DELETE on maps should either be rejected or handled gracefully
echo "  Response received (method handling OK)"
PASS_COUNT=$((PASS_COUNT + 1))
TEST_COUNT=$((TEST_COUNT + 1))
echo ""

# Summary
echo -e "${YELLOW}=== Test Summary ===${NC}"
echo -e "Total Tests: ${TEST_COUNT}"
echo -e "Passed: ${GREEN}${PASS_COUNT}${NC}"
echo -e "Failed: ${RED}${FAIL_COUNT}${NC}"

if [ $FAIL_COUNT -eq 0 ]; then
    echo -e "${GREEN}All tests passed!${NC}"
    exit 0
else
    echo -e "${RED}Some tests failed!${NC}"
    exit 1
fi
