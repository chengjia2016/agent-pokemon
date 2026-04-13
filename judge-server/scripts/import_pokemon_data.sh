#!/bin/bash

# Pokemon Database Data Import Script
# Usage: ./import_pokemon_data.sh [database_name] [user] [host] [port]

set -e

# Default values
DB_NAME="${1:-pokemon_game}"
DB_USER="${2:-postgres}"
DB_HOST="${3:-localhost}"
DB_PORT="${4:-5432}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SQL_FILE="$SCRIPT_DIR/INIT_POKEMON_DATA.sql"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}Pokemon Database Data Import${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""
echo -e "${YELLOW}Configuration:${NC}"
echo "  Database: $DB_NAME"
echo "  User: $DB_USER"
echo "  Host: $DB_HOST"
echo "  Port: $DB_PORT"
echo "  SQL File: $SQL_FILE"
echo ""

# Check if SQL file exists
if [ ! -f "$SQL_FILE" ]; then
    echo -e "${RED}Error: SQL file not found at $SQL_FILE${NC}"
    exit 1
fi

# Count SQL statements
STATEMENT_COUNT=$(grep -c "^INSERT INTO" "$SQL_FILE" 2>/dev/null || echo "0")
echo -e "${YELLOW}SQL Statements to Import:${NC} $STATEMENT_COUNT"
echo ""

# Check PostgreSQL connection
echo -e "${YELLOW}Checking PostgreSQL connection...${NC}"
if ! PGPASSWORD="${PGPASSWORD}" psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "SELECT 1;" > /dev/null 2>&1; then
    echo -e "${RED}Error: Cannot connect to database at $DB_HOST:$DB_PORT${NC}"
    echo "Make sure PostgreSQL is running and credentials are correct."
    exit 1
fi
echo -e "${GREEN}✓ Database connection successful${NC}"
echo ""

# Backup existing data (optional)
read -p "Backup existing pokemon tables? (y/n) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    BACKUP_FILE="$SCRIPT_DIR/pokemon_backup_$(date +%Y%m%d_%H%M%S).sql"
    echo -e "${YELLOW}Creating backup...${NC}"
    PGPASSWORD="${PGPASSWORD}" pg_dump -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" \
        --table="pokemon_species" \
        --table="pokemon_moves" \
        --table="pokemon_abilities" \
        --table="pokemon_move_pool" \
        > "$BACKUP_FILE" 2>/dev/null || true
    if [ -s "$BACKUP_FILE" ]; then
        echo -e "${GREEN}✓ Backup created: $BACKUP_FILE${NC}"
    else
        rm -f "$BACKUP_FILE"
    fi
    echo ""
fi

# Import data
echo -e "${YELLOW}Importing Pokemon data...${NC}"
START_TIME=$(date +%s)

if PGPASSWORD="${PGPASSWORD}" psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -f "$SQL_FILE" > /tmp/import.log 2>&1; then
    END_TIME=$(date +%s)
    DURATION=$((END_TIME - START_TIME))
    
    echo -e "${GREEN}✓ Import successful!${NC}"
    echo ""
    echo -e "${YELLOW}Import Statistics:${NC}"
    echo "  Statements: $STATEMENT_COUNT"
    echo "  Duration: ${DURATION}s"
    
    # Verify data
    echo ""
    echo -e "${YELLOW}Verifying imported data...${NC}"
    
    POKEMON_COUNT=$(PGPASSWORD="${PGPASSWORD}" psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT COUNT(*) FROM pokemon_species;" 2>/dev/null | tr -d ' ')
    MOVES_COUNT=$(PGPASSWORD="${PGPASSWORD}" psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT COUNT(*) FROM pokemon_moves;" 2>/dev/null | tr -d ' ')
    ABILITIES_COUNT=$(PGPASSWORD="${PGPASSWORD}" psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT COUNT(*) FROM pokemon_abilities;" 2>/dev/null | tr -d ' ')
    
    echo -e "${GREEN}✓ Pokemon Species: $POKEMON_COUNT${NC}"
    echo -e "${GREEN}✓ Moves: $MOVES_COUNT${NC}"
    echo -e "${GREEN}✓ Abilities: $ABILITIES_COUNT${NC}"
    echo ""
    echo -e "${GREEN}========================================${NC}"
    echo -e "${GREEN}Import completed successfully!${NC}"
    echo -e "${GREEN}========================================${NC}"
else
    END_TIME=$(date +%s)
    DURATION=$((END_TIME - START_TIME))
    
    echo -e "${RED}✗ Import failed!${NC}"
    echo -e "${RED}Error log saved to: /tmp/import.log${NC}"
    cat /tmp/import.log
    exit 1
fi

exit 0
