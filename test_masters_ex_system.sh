#!/bin/bash
# Agent Monster v2.1.0 - Masters EX System Test Suite
# Tests for Sync Pair System, 3v3 Battles, Tournament System, Seasonal Events, and Sync Moves
# Created: 2026-04-14 15:30 UTC

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
API_URL="${API_URL:-http://localhost:8080}"
TEST_DIR="/tmp/masters_ex_tests"
TIMESTAMP=$(date +%s)

# Create test directory
mkdir -p "$TEST_DIR"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}Agent Monster v2.1.0 - Masters EX Tests${NC}"
echo -e "${BLUE}========================================${NC}\n"

# Test counters
TESTS_PASSED=0
TESTS_FAILED=0

# Helper function to make API calls
api_call() {
    local method=$1
    local endpoint=$2
    local data=$3
    local output_file="$TEST_DIR/response_${TIMESTAMP}.json"

    if [ -z "$data" ]; then
        curl -s -X "$method" "$API_URL$endpoint" \
            -H "Content-Type: application/json" \
            -o "$output_file"
    else
        curl -s -X "$method" "$API_URL$endpoint" \
            -H "Content-Type: application/json" \
            -d "$data" \
            -o "$output_file"
    fi

    cat "$output_file"
}

# Helper function for test assertions
test_case() {
    local test_name=$1
    local test_cmd=$2
    
    echo -ne "${YELLOW}Testing:${NC} $test_name ... "
    
    if eval "$test_cmd" &>/dev/null; then
        echo -e "${GREEN}✓ PASSED${NC}"
        ((TESTS_PASSED++))
    else
        echo -e "${RED}✗ FAILED${NC}"
        ((TESTS_FAILED++))
    fi
}

# =============================================================================
# Test 1: Sync Pair System Tests
# =============================================================================
echo -e "${BLUE}## Test Suite 1: Sync Pair System${NC}\n"

test_case "Create Sync Pair" \
    "api_call POST '/api/masters/sync-pair/create' '{
        \"trainer_id\": \"trainer_001\",
        \"pokemon_id\": \"pokemon_001\",
        \"trainer_name\": \"Red\",
        \"pokemon_name\": \"Charizard\",
        \"pokemon_type\": \"Fire\",
        \"rarity_stars\": 5
    }' | grep -q 'sync_pair_id'"

test_case "Get Sync Pair" \
    "api_call GET '/api/masters/sync-pair/get?sync_pair_id=trainer_001_pokemon_001' | grep -q 'trainer_name'"

test_case "Get Player Sync Pairs" \
    "api_call GET '/api/masters/sync-pair/list?trainer_id=trainer_001' | grep -q 'pokemon_name'"

test_case "Level Up Sync Pair" \
    "api_call POST '/api/masters/sync-pair/level-up' '{
        \"sync_pair_id\": \"trainer_001_pokemon_001\",
        \"levels\": 10
    }' | grep -q 'level'"

test_case "Unlock Potential" \
    "api_call POST '/api/masters/sync-pair/unlock-potential?sync_pair_id=trainer_001_pokemon_001' | grep -q 'potential_unlocked'"

test_case "Set Moves" \
    "api_call POST '/api/masters/sync-pair/set-moves' '{
        \"sync_pair_id\": \"trainer_001_pokemon_001\",
        \"move_1_id\": \"move_001\",
        \"move_2_id\": \"move_002\",
        \"move_3_id\": \"move_003\",
        \"sync_move_id\": \"sync_move_001\"
    }' | grep -q 'move_1_id'"

# =============================================================================
# Test 2: Battle Team Management Tests
# =============================================================================
echo -e "\n${BLUE}## Test Suite 2: Battle Team Management${NC}\n"

test_case "Create Battle Team" \
    "api_call POST '/api/masters/team/create' '{
        \"player_id\": \"player_001\",
        \"team_name\": \"Champion Team\",
        \"sync_pair_1_id\": \"trainer_001_pokemon_001\",
        \"sync_pair_2_id\": \"trainer_002_pokemon_002\",
        \"sync_pair_3_id\": \"trainer_003_pokemon_003\",
        \"is_main_team\": true
    }' | grep -q 'team_id'"

test_case "Get Battle Team" \
    "api_call GET '/api/masters/team/get?team_id=team_player_001_*' | grep -q 'sync_pairs'"

test_case "Get Player Teams" \
    "api_call GET '/api/masters/team/list?player_id=player_001' | grep -q 'team_name'"

# =============================================================================
# Test 3: Game Statistics Tests
# =============================================================================
echo -e "\n${BLUE}## Test Suite 3: Game Statistics${NC}\n"

test_case "Get Player Game Statistics" \
    "api_call GET '/api/masters/stats/player?player_id=player_001' | grep -q 'total_sync_pairs'"

test_case "Player Statistics Contains Power Index" \
    "api_call GET '/api/masters/stats/player?player_id=player_001' | grep -q 'total_power_index'"

# =============================================================================
# Test 4: Sync Move System Tests
# =============================================================================
echo -e "\n${BLUE}## Test Suite 4: Sync Move System${NC}\n"

test_case "Increase Sync Move Ready" \
    "api_call POST '/api/masters/sync-pair/sync-move/ready' '{
        \"sync_pair_id\": \"trainer_001_pokemon_001\",
        \"percentage\": 25
    }' | grep -q 'status'"

test_case "Use Sync Move" \
    "api_call POST '/api/masters/sync-pair/sync-move/use?sync_pair_id=trainer_001_pokemon_001' | grep -q 'status'"

# =============================================================================
# Test 5: 3v3 Battle System Tests (Placeholder for battle engine)
# =============================================================================
echo -e "\n${BLUE}## Test Suite 5: 3v3 Battle System${NC}\n"

test_case "3v3 Battle System Schema Exists" \
    "[ -f '/root/petskill/judge-server/SCHEMA_MASTERS_EX.sql' ] && grep -q 'battle_sessions' '/root/petskill/judge-server/SCHEMA_MASTERS_EX.sql'"

# =============================================================================
# Test 6: Tournament System Tests (Placeholder)
# =============================================================================
echo -e "\n${BLUE}## Test Suite 6: Tournament System${NC}\n"

test_case "Tournament System Schema Exists" \
    "grep -q 'tournament_seasons' '/root/petskill/judge-server/SCHEMA_MASTERS_EX.sql'"

test_case "Tournament Rankings Schema Exists" \
    "grep -q 'tournament_rankings' '/root/petskill/judge-server/SCHEMA_MASTERS_EX.sql'"

# =============================================================================
# Test 7: Seasonal Events System Tests (Placeholder)
# =============================================================================
echo -e "\n${BLUE}## Test Suite 7: Seasonal Events System${NC}\n"

test_case "Seasonal Events Schema Exists" \
    "grep -q 'seasonal_events' '/root/petskill/judge-server/SCHEMA_MASTERS_EX.sql'"

test_case "Seasonal Items Schema Exists" \
    "grep -q 'seasonal_items' '/root/petskill/judge-server/SCHEMA_MASTERS_EX.sql'"

# =============================================================================
# Test 8: Export & Import Tests
# =============================================================================
echo -e "\n${BLUE}## Test Suite 8: Data Export & Import${NC}\n"

test_case "Export Sync Pair Data" \
    "api_call GET '/api/masters/sync-pair/export?player_id=player_001' | grep -q '\\['"

# =============================================================================
# Test Results Summary
# =============================================================================
echo -e "\n${BLUE}========================================${NC}"
echo -e "${BLUE}Test Results Summary${NC}"
echo -e "${BLUE}========================================${NC}\n"

TOTAL_TESTS=$((TESTS_PASSED + TESTS_FAILED))
PASS_RATE=$((TESTS_PASSED * 100 / TOTAL_TESTS))

echo -e "Total Tests:  ${BLUE}$TOTAL_TESTS${NC}"
echo -e "Passed:       ${GREEN}$TESTS_PASSED${NC}"
echo -e "Failed:       ${RED}$TESTS_FAILED${NC}"
echo -e "Pass Rate:    ${BLUE}${PASS_RATE}%${NC}\n"

if [ $TESTS_FAILED -eq 0 ]; then
    echo -e "${GREEN}✓ All tests passed!${NC}\n"
    exit 0
else
    echo -e "${RED}✗ Some tests failed. Check logs above.${NC}\n"
    exit 1
fi
