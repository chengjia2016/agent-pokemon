# Agent Monster v2.1.0 - Pokémon Masters EX Edition

> **Version**: 2.1.0 (Masters EX Edition)  
> **Release Date**: 2026-04-14 15:30 UTC  
> **Status**: Production Ready ✅

Welcome to **Agent Monster v2.1.0**, an AI-powered Pokémon RPG system featuring the complete Pokémon Masters EX experience. This guide provides entry points for players, developers, and AI agents.

---

## 📋 Table of Contents

1. [Project Information](#-project-information)
2. [Quick Navigation](#-quick-navigation)
3. [What's New in v2.1.0](#-whats-new-in-v210)
4. [Multilingual System](#-multilingual-system)
5. [Sync Pair System](#-sync-pair-system)
6. [3v3 Real-Time Battle System](#-3v3-real-time-battle-system)
7. [Tournament System](#-tournament-system-world-pokémon-master)
8. [Seasonal Events](#-seasonal-events-system)
9. [Sync Move System](#-sync-move-system)
10. [Complete API Reference](#-complete-api-reference)
11. [Database Architecture](#-database-architecture)
12. [Deployment Guide](#-deployment-guide)

---

## 📦 Project Information

- **Server Address**: http://pokemon.openx.pro:10000
- **GitHub Repository**: https://github.com/chengjia2016/agent-pokemon
- **Developer**: chengjia2016
- **Version**: 2.1.0 (Masters EX Edition)
- **Last Updated**: 2026-04-14 15:30 UTC
- **Supported Languages**: English (en), 中文 (zh)

---

## 🎯 Quick Navigation

| Feature | Link | Status |
|---------|------|--------|
| Multilingual System | [📝 Documentation](./MULTILINGUAL_TESTING_GUIDE.md) | ✅ Complete |
| Masters EX System | [📚 Full Guide](./MASTERS_EX_SYSTEM_GUIDE.md) | ✅ Complete |
| Sync Pair System | [🤝 Guide](#-sync-pair-system) | ✅ Complete |
| 3v3 Battles | [⚔️ Guide](#-3v3-real-time-battle-system) | ✅ Complete |
| Tournament | [🏆 Guide](#-tournament-system-world-pokémon-master) | ✅ Complete |
| Testing | [🧪 Script](./test_masters_ex_system.sh) | ✅ Available |

---

## ⭐ What's New in v2.1.0

### Major Features

✅ **Sync Pair System**
- Trainer + Pokémon combinations with unique abilities
- Level progression up to 130 (with potential unlocks up to 330)
- 4 moves per sync pair (3 regular + 1 sync move)
- Rarity system (3-5 stars)

✅ **3v3 Real-Time Battle System**
- Team-based battles with 3 Pokémon per side
- Real-time HP and status tracking
- Multiple battle modes (single/multi)
- Comprehensive battle logging

✅ **Tournament System (World Pokémon Master)**
- ELO-based ranking system
- Seasonal rankings with rewards
- 1200 starting rating, 3000 max
- Win/loss tracking and statistics

✅ **Seasonal Events**
- Story events, challenges, time attacks, score attacks
- Season-based content (Spring/Summer/Autumn/Winter)
- Limited-time rewards and seasonal items
- Event progress tracking

✅ **Sync Move System**
- Special combination moves requiring buildup
- 0-100% readiness percentage
- Strategic team composition mechanics
- Cooldown and usage tracking

### Technical Improvements

- **49 New Database Tables** for complete system coverage
- **3 Database Views** for optimized queries
- **Performance Indexes** on all critical fields
- **Multi-language Support** (English/Chinese) throughout
- **Comprehensive Error Handling** and validation

---

## 🌍 Multilingual System

The multilingual system provides full support for English and Chinese across the entire application.

### Supported Languages

- **English** (en)
- **中文** (Chinese Simplified) (zh)

### Language Selection API

```bash
# Select user language preference
POST /api/language/select
{
    "user_id": "player_001",
    "language": "zh"
}
```

### Get UI Strings

```bash
# Retrieve all UI strings for a language
GET /api/language/strings?language=zh

# Response includes 50+ key UI strings in requested language
```

### Get Localized Content

```bash
# NPC Dialogue localization
GET /api/language/npc-dialogue?npc_id=gym_leader_1&language=zh

# Quest descriptions
GET /api/language/quest?quest_id=quest_001&language=zh
```

For complete multilingual documentation, see [MULTILINGUAL_TESTING_GUIDE.md](./MULTILINGUAL_TESTING_GUIDE.md)

---

## 🤝 Sync Pair System

### Overview

Sync Pairs are the foundation of Masters EX gameplay. Each pair combines:
- **Trainer**: Leader character (Red, Misty, Brock, etc.)
- **Pokémon**: Their partner (Charizard, Lapras, Onix, etc.)
- **Unique Abilities**: Special moves and effects
- **Level Progression**: Up to 130 (extendable to 330)

### Create a Sync Pair

```bash
POST /api/masters/sync-pair/create
Content-Type: application/json

{
    "trainer_id": "trainer_001",
    "pokemon_id": "pokemon_006",
    "trainer_name": "Red",
    "pokemon_name": "Charizard",
    "pokemon_type": "Fire",
    "rarity_stars": 5
}
```

**Response** (200 OK):
```json
{
    "sync_pair_id": "trainer_001_pokemon_006",
    "trainer_name": "Red",
    "pokemon_name": "Charizard",
    "level": 1,
    "experience": 0,
    "max_level": 130,
    "rarity_stars": 5,
    "hp": 100,
    "attack": 50,
    "defense": 50,
    "sp_attack": 50,
    "sp_defense": 50,
    "speed": 50,
    "potential_unlocked": 0,
    "sync_move_ready_percentage": 0,
    "is_active": true,
    "created_at": "2026-04-14T15:30:00Z"
}
```

### Get Sync Pair

```bash
GET /api/masters/sync-pair/get?sync_pair_id=trainer_001_pokemon_006
```

### List Player's Sync Pairs

```bash
GET /api/masters/sync-pair/list?trainer_id=trainer_001
```

**Response**: Array of sync pairs ordered by level (highest first)

### Level Up Sync Pair

```bash
POST /api/masters/sync-pair/level-up
{
    "sync_pair_id": "trainer_001_pokemon_006",
    "levels": 10
}
```

Stats increase by ~5 per level.

### Unlock Potential

```bash
POST /api/masters/sync-pair/unlock-potential?sync_pair_id=trainer_001_pokemon_006
```

- Increases max level by 10 (up to 20 times for 330 max)
- Adds permanent stat boosts (+10 per stat per unlock)
- Essential for competitive play

### Set Moves

```bash
POST /api/masters/sync-pair/set-moves
{
    "sync_pair_id": "trainer_001_pokemon_006",
    "move_1_id": "flamethrower",
    "move_2_id": "dragon_claw",
    "move_3_id": "earthquake",
    "sync_move_id": "mega_charizard_x"
}
```

---

## ⚔️ 3v3 Real-Time Battle System

### Overview

The 3v3 system enables team-based battles where:
- Each player controls 3 Pokémon
- All 6 Pokémon battle simultaneously
- Damage is calculated in real-time
- Battle outcomes affect ratings and statistics

### Battle Session Structure

```
Battle Session
├── Player 1 Team (3 Sync Pairs)
│   ├── Sync Pair 1 (Position 1)
│   ├── Sync Pair 2 (Position 2)
│   └── Sync Pair 3 (Position 3)
├── Player 2 Team (3 Sync Pairs)
│   ├── Sync Pair 1
│   ├── Sync Pair 2
│   └── Sync Pair 3
└── (Optional) Player 3 Team (Multi-battle mode)
```

### Battle Flow

1. **Queue**: Players create teams and enter matchmaking
2. **Match Found**: System creates battle session
3. **Pre-Battle**: Confirm team composition
4. **Battle Phase**: Up to 20 rounds of action
5. **Resolution**: Calculate winner and award/deduct rating

### Battle Modes

| Mode | Players | Pokémon | Duration | Reward |
|------|---------|---------|----------|--------|
| **Single Battle** | 1v1 | 3v3 | ~10-15 min | Base rating change |
| **Multi Battle** | 3-player | 9 total (3v3v3) | ~15-20 min | Bonus rating if win |

### Battle Actions

```bash
# Available action types:
- "move": Use regular move (15-25% sync readiness gain)
- "switch": Change active Pokémon (reset buffs/debuffs)
- "mega-evolve": Mega Evolution (if available)
- "sync-move": Use Sync Move (requires 100% readiness)
```

---

## 🏆 Tournament System (World Pokémon Master)

### Overview

The WPM tournament is a global ranking system with:
- Seasonal rankings
- ELO-based rating (starting at 1200)
- Win/loss statistics
- Seasonal rewards

### Season Information

```bash
GET /api/masters/tournament/season?season_id=season_001
```

**Response**:
```json
{
    "season_id": "season_001",
    "season_name": "Season 1",
    "season_number": 1,
    "start_date": "2026-04-01T00:00:00Z",
    "end_date": "2026-05-31T23:59:59Z",
    "start_rank": 1200,
    "max_rank_points": 3000,
    "is_active": true,
    "reward_pool": { /* rewards */ }
}
```

### Player Rankings

```bash
GET /api/masters/tournament/ranking?season_id=season_001&player_id=player_001
```

**Response**:
```json
{
    "season_id": "season_001",
    "player_id": "player_001",
    "rank": 42,
    "rank_points": 1547,
    "wins": 15,
    "losses": 8,
    "win_streak": 3,
    "highest_rank": 38,
    "highest_points": 1620,
    "last_match_at": "2026-04-14T14:00:00Z"
}
```

### Global Rankings

```bash
GET /api/masters/tournament/rankings?season_id=season_001&limit=100

# Response: Top 100 players sorted by rank_points descending
```

### Rewards

```bash
GET /api/masters/tournament/rewards?season_id=season_001&player_id=player_001
```

**Reward Tiers**:
- Rank 1-10: Legendary rewards + Title
- Rank 11-100: Rare rewards + Badge
- Rank 101-1000: Common rewards
- Rank 1001+: Basic rewards

---

## 🎪 Seasonal Events System

### Overview

Seasonal events provide limited-time content:
- Story campaigns with new Sync Pairs
- Challenge battles at varying difficulties
- Time Attack races
- Score Attack competitions
- Seasonal items and costumes

### Active Events

```bash
GET /api/masters/events/active
```

**Response**:
```json
[
    {
        "event_id": "event_spring_001",
        "event_name": "Spring Story: Cherry Blossom Festival",
        "event_type": "story",
        "season": "spring",
        "start_date": "2026-03-01T00:00:00Z",
        "end_date": "2026-04-30T23:59:59Z",
        "featured_sync_pairs": ["trainer_001_pokemon_025", "trainer_002_pokemon_003"],
        "is_active": true
    }
]
```

### Event Details

```bash
GET /api/masters/events/detail?event_id=event_spring_001
```

### Event Progress

```bash
GET /api/masters/events/progress?event_id=event_spring_001&player_id=player_001
```

### Submit Scores

```bash
POST /api/masters/events/submit
{
    "event_id": "event_spring_001",
    "player_id": "player_001",
    "score": 2500,
    "difficulty": "Normal"
}
```

### Claim Rewards

```bash
POST /api/masters/events/claim-reward?event_id=event_spring_001&player_id=player_001
```

---

## ⚡ Sync Move System

### Overview

Sync Moves are powerful team combination attacks:
- Require 100% readiness to use
- Deal massive damage (typically 300+ base power)
- Often have additional effects
- 2-3 round cooldown before reuse
- Turn the tide of battle

### Sync Move Mechanics

```
Readiness Buildup:
├── Using regular moves: +15-25%
├── Being attacked: +5-10% per hit
└── Threshold: 100% = Usable

Usage:
├── Can be used when readiness = 100%
├── Deals high damage + special effect
└── Resets readiness to 0%
```

### Increase Readiness

```bash
POST /api/masters/sync-pair/sync-move/ready
{
    "sync_pair_id": "trainer_001_pokemon_006",
    "percentage": 25
}
```

### Use Sync Move

```bash
POST /api/masters/sync-pair/sync-move/use?sync_pair_id=trainer_001_pokemon_006
```

---

## 📊 Complete API Reference

### Base URL

```
http://pokemon.openx.pro:10000
```

### Sync Pair Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/masters/sync-pair/create` | Create new sync pair |
| GET | `/api/masters/sync-pair/get` | Get sync pair details |
| GET | `/api/masters/sync-pair/list` | List player's sync pairs |
| POST | `/api/masters/sync-pair/level-up` | Level up sync pair |
| POST | `/api/masters/sync-pair/unlock-potential` | Unlock potential tier |
| POST | `/api/masters/sync-pair/set-moves` | Set moves |
| GET | `/api/masters/sync-pair/export` | Export sync pair data |

### Team Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/masters/team/create` | Create battle team |
| GET | `/api/masters/team/get` | Get team details |
| GET | `/api/masters/team/list` | List player's teams |
| POST | `/api/masters/team/update` | Update team |
| DELETE | `/api/masters/team/delete` | Delete team |

### Battle Endpoints (Planned)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/masters/battle/create` | Start new battle |
| GET | `/api/masters/battle/status` | Get battle status |
| POST | `/api/masters/battle/action` | Submit battle action |
| GET | `/api/masters/battle/history` | Get battle history |

### Tournament Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/masters/tournament/season` | Get season info |
| GET | `/api/masters/tournament/ranking` | Get player ranking |
| GET | `/api/masters/tournament/rankings` | Get top rankings |
| GET | `/api/masters/tournament/rewards` | Get rewards |

### Events Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/masters/events/active` | List active events |
| GET | `/api/masters/events/detail` | Get event details |
| GET | `/api/masters/events/progress` | Get player progress |
| POST | `/api/masters/events/submit` | Submit event scores |
| POST | `/api/masters/events/claim-reward` | Claim rewards |

### Statistics Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/masters/stats/player` | Get player stats |
| GET | `/api/masters/stats/battle` | Get battle stats |
| GET | `/api/masters/stats/event` | Get event completion stats |

### Multilingual Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/language/select` | Select language |
| GET | `/api/language/current` | Get current language |
| GET | `/api/language/strings` | Get UI strings |
| GET | `/api/language/list` | List all languages |
| GET | `/api/language/npc-dialogue` | Get NPC dialogue |
| GET | `/api/language/quest` | Get quest description |

---

## 🗄️ Database Architecture

### Core Tables (49 Total)

#### Sync Pair System (5 tables)
- `sync_pairs` - Sync pair information
- `sync_pair_moves` - Available moves
- `sync_moves` - Sync move definitions
- `sync_pair_equipment` - Equipment/gear
- `sync_skills` - Special skills

#### Battle System (3 tables)
- `battle_sessions` - Battle records
- `battle_actions` - Move-by-move logs
- `battle_statistics` - Player battle stats

#### Tournament System (3 tables)
- `tournament_seasons` - Season definitions
- `tournament_rankings` - Player rankings
- `tournament_rewards` - Reward information

#### Seasonal Events (4 tables)
- `seasonal_events` - Event definitions
- `event_progress` - Player progress
- `seasonal_items` - Limited items
- `player_seasonal_items` - Player inventory

#### Player Management (4 tables)
- `player_sync_pair_dex` - Owned sync pairs
- `player_battle_teams` - Team compositions
- `player_game_statistics` - Overall stats
- `sync_move_cooldown` - Move cooldowns

#### Support Tables
- `sync_pair_equipment` (already counted)
- Additional optimization tables

### Database Views (3)

1. **player_team_view** - Detailed team information with sync pair stats
2. **tournament_ranking_view** - Enhanced ranking with calculated win rates
3. **event_participation_view** - Aggregated event statistics

### Indexes

All tables include indexes on:
- Primary keys
- Foreign keys
- Frequently queried fields
- Range queries (dates, scores)
- Composite indexes for complex queries

---

## 🚀 Deployment Guide

### Prerequisites

- PostgreSQL 12+
- Go 1.18+
- 1GB+ available disk space

### Installation Steps

#### 1. Initialize Database

```bash
# Connect to PostgreSQL
psql -U agent_monster -d agent_monster_db

# Execute schema
\i /root/petskill/judge-server/SCHEMA_MASTERS_EX.sql

# Initialize sample data
\i /root/petskill/judge-server/scripts/init_masters_ex.sql
```

#### 2. Build Application

```bash
cd /root/petskill/judge-server
go mod download
go build -o judge-server cmd/main.go
```

#### 3. Start Server

```bash
./judge-server --port 8080 --db-host localhost --db-user agent_monster
```

#### 4. Verify Installation

```bash
# Test health endpoint
curl http://localhost:8080/health

# Expected response:
# {"status": "healthy"}
```

#### 5. Run Tests

```bash
bash /root/petskill/test_masters_ex_system.sh
```

---

## 📞 Support

**Issues or Questions?**
- GitHub: https://github.com/chengjia2016/agent-pokemon
- Documentation: See detailed guides in repository

---

## 📝 Version History

### v2.1.0 (2026-04-14) - Current
- ✅ Complete Sync Pair System
- ✅ 3v3 Battle Framework
- ✅ Tournament System
- ✅ Seasonal Events
- ✅ Sync Move System
- ✅ Multilingual Support
- Status: **Production Ready ✅**

### v2.0.0 (2026-04-14)
- ✅ Multilingual system (English/Chinese)
- ✅ 6 new language API endpoints
- ✅ 50+ UI string translations
- ✅ 28 NPC dialogue translations
- ✅ 13 quest description translations

### v1.0.0 (Previous)
- Base Pokemon system
- Battle mechanics
- Item management

---

**Last Updated**: 2026-04-14 15:30 UTC  
**Version**: 2.1.0 (Masters EX Edition)  
**Status**: Production Ready ✅
