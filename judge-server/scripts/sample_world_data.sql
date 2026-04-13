-- Sample World System Data Initialization
-- This script creates sample NPCs, quests, dungeons, gyms, and levels for testing

-- =========================
-- NPC System Sample Data
-- =========================

-- Sample NPC 1: Quest Giver in first town
INSERT INTO npcs (town_id, name_en, name_zh, type, role, coord_x, coord_y, avatar_url, created_at)
VALUES (
  '1', 
  'Professor Oak', 
  '橡树博士', 
  'NPC', 
  'quest_giver',
  50.0, 
  50.0,
  'https://example.com/oak.png',
  NOW()
) ON CONFLICT DO NOTHING;

-- Sample NPC 2: Gym Leader
INSERT INTO npcs (town_id, name_en, name_zh, type, role, coord_x, coord_y, avatar_url, created_at)
VALUES (
  '1',
  'Brock',
  '布洛克',
  'NPC',
  'gym_leader',
  100.0,
  100.0,
  'https://example.com/brock.png',
  NOW()
) ON CONFLICT DO NOTHING;

-- Sample NPC 3: Merchant
INSERT INTO npcs (town_id, name_en, name_zh, type, role, coord_x, coord_y, avatar_url, created_at)
VALUES (
  '1',
  'Shopkeeper',
  '店主',
  'NPC',
  'merchant',
  75.0,
  75.0,
  'https://example.com/merchant.png',
  NOW()
) ON CONFLICT DO NOTHING;

-- Add NPC Dialogues
INSERT INTO npc_dialogues (npc_id, dialogue_key, dialogue_en, dialogue_zh, dialogue_order, created_at)
SELECT id, 'greeting', 'Welcome, trainer!', '欢迎，训练师！', 1, NOW()
FROM npcs WHERE name_en = 'Professor Oak'
ON CONFLICT DO NOTHING;

INSERT INTO npc_dialogues (npc_id, dialogue_key, dialogue_en, dialogue_zh, dialogue_order, created_at)
SELECT id, 'quest_offer', 'I have a quest for you!', '我有一个任务给你！', 2, NOW()
FROM npcs WHERE name_en = 'Professor Oak'
ON CONFLICT DO NOTHING;

-- =========================
-- Quest System Sample Data
-- =========================

-- Sample Quest 1
INSERT INTO quests (quest_giver_id, title_en, title_zh, description_en, description_zh, quest_type, reward_gold, reward_exp, is_active, created_at)
SELECT id, 
       'Catch 5 Pokemons',
       '捕捉5只宝可梦',
       'Catch and collect 5 different Pokemon species',
       '捕捉并收集5只不同的宝可梦物种',
       'capture',
       100,
       500,
       TRUE,
       NOW()
FROM npcs WHERE name_en = 'Professor Oak'
ON CONFLICT DO NOTHING;

-- Sample Quest 2
INSERT INTO quests (quest_giver_id, title_en, title_zh, description_en, description_zh, quest_type, reward_gold, reward_exp, is_active, created_at)
SELECT id,
       'Train Your Pokemon to Level 10',
       '将你的宝可梦训练到10级',
       'Train your primary Pokemon until it reaches level 10',
       '训练你的主要宝可梦直到达到10级',
       'training',
       200,
       1000,
       TRUE,
       NOW()
FROM npcs WHERE name_en = 'Professor Oak'
ON CONFLICT DO NOTHING;

-- =========================
-- Dungeon System Sample Data
-- =========================

-- Sample Dungeon 1: Mystery Cave
INSERT INTO dungeons (name_en, name_zh, description_en, description_zh, difficulty_level, boss_name_en, boss_name_zh, boss_level, reward_gold, reward_items, is_active, created_at)
VALUES (
  'Mystery Cave',
  '神秘洞穴',
  'A dark cave filled with wild Pokemon',
  '一个充满野生宝可梦的黑暗洞穴',
  'easy',
  'Onix',
  '大岩蛇',
  15,
  500,
  '{"items": [{"name": "Pokedex", "quantity": 1}]}',
  TRUE,
  NOW()
) ON CONFLICT DO NOTHING;

-- Sample Dungeon Floors for Mystery Cave
INSERT INTO dungeon_floors (dungeon_id, floor_number, enemy_config, floor_description, enemy_level_min, enemy_level_max, is_boss_floor)
SELECT id, 1, '{"enemies": ["Zubat", "Paras"]}', 'First floor with weak enemies', 5, 8, FALSE
FROM dungeons WHERE name_en = 'Mystery Cave'
ON CONFLICT DO NOTHING;

INSERT INTO dungeon_floors (dungeon_id, floor_number, enemy_config, floor_description, enemy_level_min, enemy_level_max, is_boss_floor)
SELECT id, 2, '{"enemies": ["Golbat", "Parasect", "Graveler"]}', 'Second floor - stronger enemies', 10, 12, FALSE
FROM dungeons WHERE name_en = 'Mystery Cave'
ON CONFLICT DO NOTHING;

INSERT INTO dungeon_floors (dungeon_id, floor_number, enemy_config, floor_description, enemy_level_min, enemy_level_max, is_boss_floor)
SELECT id, 3, '{"boss": "Onix", "level": 15}', 'Boss floor - face the Onix!', 15, 15, TRUE
FROM dungeons WHERE name_en = 'Mystery Cave'
ON CONFLICT DO NOTHING;

-- =========================
-- Gym System Sample Data
-- =========================

-- Sample Gym: Pewter City Gym
INSERT INTO gyms (gym_leader_id, gym_name_en, gym_name_zh, city_location, gym_badge_en, gym_badge_zh, difficulty_level, reward_gold, created_at)
SELECT id,
       'Pewter City Gym',
       '常磐市健身房',
       'Pewter City',
       'Boulder Badge',
       '岩石勋章',
       'normal',
       1000,
       NOW()
FROM npcs WHERE name_en = 'Brock'
ON CONFLICT DO NOTHING;

-- Add Gym Team members (Brock's team)
INSERT INTO gym_teams (gym_id, pokemon_species, pokemon_level, move_pool)
SELECT id, 'Geodude', 12, '["Tackle", "Defense Curl", "Rock Throw"]'
FROM gyms WHERE gym_name_en = 'Pewter City Gym'
ON CONFLICT DO NOTHING;

INSERT INTO gym_teams (gym_id, pokemon_species, pokemon_level, move_pool)
SELECT id, 'Onix', 14, '["Tackle", "Bind", "Harden", "Rock Throw"]'
FROM gyms WHERE gym_name_en = 'Pewter City Gym'
ON CONFLICT DO NOTHING;

-- =========================
-- Map Zone & Level System Sample Data
-- =========================

-- Sample Region
INSERT INTO regions (region_name_en, region_name_zh, region_description, climate, is_active, created_at)
VALUES (
  'Kanto Region',
  '关都地区',
  'The starting region for trainers',
  'temperate',
  TRUE,
  NOW()
) ON CONFLICT DO NOTHING;

-- Sample Map Zones
INSERT INTO map_zones (zone_name_en, zone_name_zh, region_id, zone_type, difficulty_level, coord_x, coord_y, created_at)
SELECT 
  'Viridian Forest',
  '常青森林',
  id,
  'forest',
  'beginner',
  25.0,
  25.0,
  NOW()
FROM regions WHERE region_name_en = 'Kanto Region'
ON CONFLICT DO NOTHING;

INSERT INTO map_zones (zone_name_en, zone_name_zh, region_id, zone_type, difficulty_level, coord_x, coord_y, created_at)
SELECT 
  'Mt. Moon',
  '月见山',
  id,
  'mountain',
  'intermediate',
  75.0,
  75.0,
  NOW()
FROM regions WHERE region_name_en = 'Kanto Region'
ON CONFLICT DO NOTHING;

-- Sample Levels in Viridian Forest
INSERT INTO levels (zone_id, level_number, level_name_en, level_name_zh, difficulty, pokemon_spawns, reward_gold, reward_exp, is_unlocked)
SELECT id, 1, 'Forest Entrance', '森林入口', 'easy', '["Pidgeotto", "Pikachu"]', 100, 200, TRUE
FROM map_zones WHERE zone_name_en = 'Viridian Forest'
ON CONFLICT DO NOTHING;

INSERT INTO levels (zone_id, level_number, level_name_en, level_name_zh, difficulty, pokemon_spawns, reward_gold, reward_exp, is_unlocked)
SELECT id, 2, 'Forest Depths', '森林深处', 'normal', '["Pidgeot", "Pikachu", "Bulbasaur"]', 200, 400, FALSE
FROM map_zones WHERE zone_name_en = 'Viridian Forest'
ON CONFLICT DO NOTHING;

-- Sample Grass Areas
INSERT INTO grass_areas (zone_id, area_name_en, area_name_zh, grass_density, pokemon_list, encounter_rate, created_at)
SELECT id, 'Tall Grass Patch', '高草丛区域', 'high', '["Pidgeotto", "Pikachu", "Paras"]', 40, NOW()
FROM map_zones WHERE zone_name_en = 'Viridian Forest'
ON CONFLICT DO NOTHING;

-- Sample Wild Pokemon Spawn Config
INSERT INTO wild_pokemon_spawns (grass_area_id, pokemon_species, spawn_chance, level_min, level_max, created_at)
SELECT id, 'Pidgeotto', 30, 5, 8, NOW()
FROM grass_areas WHERE area_name_en = 'Tall Grass Patch'
ON CONFLICT DO NOTHING;

INSERT INTO wild_pokemon_spawns (grass_area_id, pokemon_species, spawn_chance, level_min, level_max, created_at)
SELECT id, 'Pikachu', 50, 3, 7, NOW()
FROM grass_areas WHERE area_name_en = 'Tall Grass Patch'
ON CONFLICT DO NOTHING;

INSERT INTO wild_pokemon_spawns (grass_area_id, pokemon_species, spawn_chance, level_min, level_max, created_at)
SELECT id, 'Paras', 20, 4, 6, NOW()
FROM grass_areas WHERE area_name_en = 'Tall Grass Patch'
ON CONFLICT DO NOTHING;

-- =========================
-- Commit the sample data
-- =========================
COMMIT;
