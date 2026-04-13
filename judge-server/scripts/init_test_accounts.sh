#!/bin/bash
# Initialize test accounts for Agent Monster Game

set -e

# Configuration
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-postgres}"
DB_PASSWORD="${DB_PASSWORD:-}"
DB_NAME="${DB_NAME:-agent_monster}"
API_URL="${API_URL:-http://localhost:8080}"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🎮 Agent Monster - Test Account Initialization${NC}"
echo "================================================"
echo ""

# Test database connection
echo -e "${BLUE}Testing database connection...${NC}"
if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "SELECT 1" > /dev/null 2>&1; then
    echo -e "${GREEN}✓ Database connection successful${NC}"
else
    echo "✗ Failed to connect to database"
    exit 1
fi

echo ""
echo -e "${BLUE}Creating test accounts...${NC}"

# Test account definitions
declare -a TEST_ACCOUNTS=(
    "test_player_1:password123"
    "test_player_2:password123"
    "test_player_3:password123"
    "trainer_pikachu:pokemon_trainer"
    "trainer_charizard:dragon_master"
    "trainer_blastoise:water_master"
    "claude_player:claude_test123"
    "opencode_player:opencode_test123"
    "openclaw_player:openclaw_test123"
    "gemini_player:gemini_test123"
)

# Create accounts via API
for account in "${TEST_ACCOUNTS[@]}"; do
    IFS=':' read -r username password <<< "$account"
    
    echo -e "${BLUE}Creating account: ${NC}$username"
    
    RESPONSE=$(curl -s -X POST $API_URL/api/auth/register-test \
        -H "Content-Type: application/json" \
        -d "{\"username\": \"$username\", \"password\": \"$password\"}")
    
    # Check if creation was successful
    if echo "$RESPONSE" | grep -q '"success":true'; then
        PLAYER_ID=$(echo "$RESPONSE" | grep -o '"player_id":"[^"]*"' | cut -d'"' -f4)
        SERVER_TOKEN=$(echo "$RESPONSE" | grep -o '"server_token":"[^"]*"' | cut -d'"' -f4)
        echo -e "${GREEN}✓ Account created${NC}"
        echo "  Username: $username"
        echo "  Player ID: $PLAYER_ID"
        echo "  Server Token: $SERVER_TOKEN"
    else
        # Account might already exist, try to get it
        echo "  (Account may already exist, continuing...)"
    fi
    
    echo ""
done

echo -e "${BLUE}Listing all test accounts...${NC}"
ACCOUNTS=$(curl -s -X GET $API_URL/api/auth/test-accounts)

if echo "$ACCOUNTS" | grep -q '"success":true'; then
    TOTAL=$(echo "$ACCOUNTS" | grep -o '"total":[0-9]*' | cut -d':' -f2)
    echo -e "${GREEN}✓ Total test accounts: $TOTAL${NC}"
    echo ""
    echo "Test accounts:"
    echo "$ACCOUNTS" | grep -o '"username":"[^"]*"' | cut -d'"' -f4 | while read username; do
        echo "  - $username"
    done
else
    echo "✗ Failed to list accounts"
fi

echo ""
echo -e "${GREEN}================================================"
echo "✓ Test account initialization complete!"
echo "================================================${NC}"
echo ""
echo "Next steps:"
echo "1. Try logging in with any test account:"
echo "   curl -X POST http://localhost:8080/api/auth/login \\"
echo "     -H 'Content-Type: application/json' \\"
echo "     -d '{\"username\": \"test_player_1\", \"password\": \"password123\", \"client_name\": \"claude\"}'"
echo ""
echo "2. Use the server_token for direct authentication:"
echo "   curl -H 'Authorization: <server_token>' http://localhost:8080/api/user/profile"
echo ""
