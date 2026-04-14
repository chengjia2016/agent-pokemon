#!/bin/bash
# API Authentication Demo Script
# This script demonstrates the complete authentication flow

set -e

SERVER_URL="${SERVER_URL:-http://localhost:10000}"
CONFIG_DIR="$HOME/.config/petskill"
CONFIG_FILE="$CONFIG_DIR/settings.json"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== PetSkill API Authentication Demo ===${NC}\n"

# Step 1: Create a new user account
echo -e "${YELLOW}Step 1: Creating new user account...${NC}"
USER_DATA=$(curl -s -X POST "$SERVER_URL/api/users/create" \
  -H "Content-Type: application/json" \
  -d '{
    "github_id": 555888777,
    "github_login": "demo_user",
    "email": "demo@example.com",
    "avatar_url": "https://example.com/avatar.jpg"
  }')

echo "$USER_DATA" | jq .

# Extract API key from response
API_KEY=$(echo "$USER_DATA" | jq -r '.api_key')
GITHUB_ID=$(echo "$USER_DATA" | jq -r '.user.github_id')

echo -e "${GREEN}✓ User created with API key: $API_KEY${NC}\n"

# Step 2: Save configuration to local file
echo -e "${YELLOW}Step 2: Saving configuration to $CONFIG_FILE...${NC}"
mkdir -p "$CONFIG_DIR"
cat > "$CONFIG_FILE" << EOF
{
  "github_id": $GITHUB_ID,
  "github_login": "demo_user",
  "api_key": "$API_KEY",
  "email": "demo@example.com"
}
EOF

echo "Configuration saved:"
cat "$CONFIG_FILE" | jq .
echo -e "${GREEN}✓ Configuration saved${NC}\n"

# Step 3: Get user configuration via API
echo -e "${YELLOW}Step 3: Getting user configuration from API...${NC}"
curl -s -X POST "$SERVER_URL/api/user/init-config" \
  -H "Content-Type: application/json" \
  -d "{\"github_id\": $GITHUB_ID}" | jq .
echo -e "${GREEN}✓ Configuration retrieved${NC}\n"

# Step 4: Use API key in request headers
echo -e "${YELLOW}Step 4: Making authenticated API calls with API key...${NC}"

echo "Method 1: Using X-API-Key header"
curl -s "$SERVER_URL/api/user/balance/get?github_id=$GITHUB_ID" \
  -H "X-API-Key: $API_KEY" | jq .

echo ""
echo "Method 2: Using Bearer token (Authorization header)"
curl -s "$SERVER_URL/api/user/balance/get?github_id=$GITHUB_ID" \
  -H "Authorization: Bearer $API_KEY" | jq .

echo -e "${GREEN}✓ Authenticated API calls successful${NC}\n"

echo -e "${BLUE}=== Authentication Demo Complete ===${NC}"
echo -e "Configuration file location: ${YELLOW}$CONFIG_FILE${NC}"
echo -e "You can now use this API key for all protected endpoints!"
