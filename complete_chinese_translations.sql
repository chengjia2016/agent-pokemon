-- ============================================================
-- Agent Monster - Complete Chinese Translations Update Script
-- ============================================================
-- 
-- 此脚本将所有 NPC 对话和任务描述的中文翻译添加到数据库
-- Run after multilingual tables are created
-- Execution: psql -h localhost -U postgres -d agent_monster -f this_file.sql
-- ============================================================

-- ============================================================
-- NPC Dialogues - Chinese Translations
-- ============================================================

-- Misty (Cerulean Gym) - NPC ID: 1
UPDATE npc_dialogues SET dialogue_text_zh = '欢迎来到华蓝道馆！我是馆主米斯蒂，钢铁属性宝可梦的训练大师。' 
WHERE npc_id = 1 AND dialogue_context = 'challenge' AND dialogue_type = 'gym_leader';

UPDATE npc_dialogues SET dialogue_text_zh = '看来你的实力不错！这是你赢得的徽章。' 
WHERE npc_id = 1 AND dialogue_context = 'victory' AND dialogue_type = 'gym_leader';

-- Brock (Pewter Gym) - NPC ID: 2
UPDATE npc_dialogues SET dialogue_text_zh = '我是钢铁城道馆馆主波克，岩石属性宝可梦的大师。你准备好接受挑战了吗？' 
WHERE npc_id = 2 AND dialogue_context = 'challenge' AND dialogue_type = 'gym_leader';

UPDATE npc_dialogues SET dialogue_text_zh = '你赢得了我的尊重。这是岩石徽章，祝你继续努力！' 
WHERE npc_id = 2 AND dialogue_context = 'victory' AND dialogue_type = 'gym_leader';

-- Psychic Master (Saffron Gym) - NPC ID: 3
UPDATE npc_dialogues SET dialogue_text_zh = '欢迎来到琉璃道馆。我是馆主妮姿，超能力属性宝可梦的训练者。' 
WHERE npc_id = 3 AND dialogue_context = 'challenge' AND dialogue_type = 'gym_leader';

UPDATE npc_dialogues SET dialogue_text_zh = '你的宝可梦很强！这是虹徽章。希望你能越来越强大。' 
WHERE npc_id = 3 AND dialogue_context = 'victory' AND dialogue_type = 'gym_leader';

-- Ground Master (Viridian Gym) - NPC ID: 4
UPDATE npc_dialogues SET dialogue_text_zh = '我是常磐市道馆馆主坂木，地面属性宝可梦的真正大师。' 
WHERE npc_id = 4 AND dialogue_context = 'challenge' AND dialogue_type = 'gym_leader';

UPDATE npc_dialogues SET dialogue_text_zh = '你证明了自己的实力。地面徽章是你的了。' 
WHERE npc_id = 4 AND dialogue_context = 'victory' AND dialogue_type = 'gym_leader';

-- Fire Master (Cinnabar Gym) - NPC ID: 5
UPDATE npc_dialogues SET dialogue_text_zh = '红莲镇道馆馆主科特在这里。火属性是最强大的！' 
WHERE npc_id = 5 AND dialogue_context = 'challenge' AND dialogue_type = 'gym_leader';

UPDATE npc_dialogues SET dialogue_text_zh = '不错的表现！这是火焰徽章，继续加油吧！' 
WHERE npc_id = 5 AND dialogue_context = 'victory' AND dialogue_type = 'gym_leader';

-- Grass Master (Viridian Forest) - NPC ID: 6
UPDATE npc_dialogues SET dialogue_text_zh = '我是绿荫镇道馆馆主莉莉，草属性宝可梦的优雅使用者。' 
WHERE npc_id = 6 AND dialogue_context = 'challenge' AND dialogue_type = 'gym_leader';

UPDATE npc_dialogues SET dialogue_text_zh = '你的宝可梦和你一样有天赋。叶绿徽章就是证明。' 
WHERE npc_id = 6 AND dialogue_context = 'victory' AND dialogue_type = 'gym_leader';

-- Electric Master (Vermilion City) - NPC ID: 7
UPDATE npc_dialogues SET dialogue_text_zh = '黄金市道馆馆主阿杜，电属性的使用者。准备好被电到吗？' 
WHERE npc_id = 7 AND dialogue_context = 'challenge' AND dialogue_type = 'gym_leader';

UPDATE npc_dialogues SET dialogue_text_zh = '很好的表现！这是闪电徽章，你赢得了它。' 
WHERE npc_id = 7 AND dialogue_context = 'victory' AND dialogue_type = 'gym_leader';

-- Water Master (Cerulean City) - NPC ID: 8
UPDATE npc_dialogues SET dialogue_text_zh = '我是水系道馆馆主米可利，水之魂的拥有者。' 
WHERE npc_id = 8 AND dialogue_context = 'challenge' AND dialogue_type = 'gym_leader';

UPDATE npc_dialogues SET dialogue_text_zh = '你的实力令人印象深刻。这是海浪徽章。' 
WHERE npc_id = 8 AND dialogue_context = 'victory' AND dialogue_type = 'gym_leader';

-- Professor Oak - NPC ID: 9
UPDATE npc_dialogues SET dialogue_text_zh = '我是大木博士，很高兴见到你，年轻的训练师！你开始你的冒险了吗？' 
WHERE npc_id = 9 AND dialogue_context = 'greeting' AND dialogue_type = 'professor';

-- Officer Jenny - NPC ID: 10
UPDATE npc_dialogues SET dialogue_text_zh = '罪犯逮捕任务！我是詹妮警官，需要你的帮助来追捕逃犯。' 
WHERE npc_id = 10 AND dialogue_context = 'mission' AND dialogue_type = 'officer';

-- ============================================================
-- Quest Descriptions - Chinese Translations
-- ============================================================

UPDATE quests SET 
    description_zh = '在华蓝市击败米斯蒂道馆馆主并获得钢铁徽章', 
    reward_item_zh = '钢铁徽章'
WHERE id = 1;

UPDATE quests SET 
    description_zh = '在钢铁城击败波克道馆馆主并获得岩石徽章', 
    reward_item_zh = '岩石徽章'
WHERE id = 2;

UPDATE quests SET 
    description_zh = '在琉璃市击败妮姿道馆馆主并获得虹徽章', 
    reward_item_zh = '虹徽章'
WHERE id = 3;

UPDATE quests SET 
    description_zh = '在常磐市击败坂木道馆馆主并获得地面徽章', 
    reward_item_zh = '地面徽章'
WHERE id = 4;

UPDATE quests SET 
    description_zh = '在红莲镇击败科特道馆馆主并获得火焰徽章', 
    reward_item_zh = '火焰徽章'
WHERE id = 5;

UPDATE quests SET 
    description_zh = '在绿荫镇击败莉莉道馆馆主并获得叶绿徽章', 
    reward_item_zh = '叶绿徽章'
WHERE id = 6;

UPDATE quests SET 
    description_zh = '在黄金市击败阿杜道馆馆主并获得闪电徽章', 
    reward_item_zh = '闪电徽章'
WHERE id = 7;

UPDATE quests SET 
    description_zh = '在水系道馆击败米可利道馆馆主并获得海浪徽章', 
    reward_item_zh = '海浪徽章'
WHERE id = 8;

UPDATE quests SET 
    description_zh = '收集所有八个道馆徽章', 
    reward_item_zh = '大师证书'
WHERE id = 9;

UPDATE quests SET 
    description_zh = '探索真新镇，认识大木博士', 
    reward_item_zh = '入门宝可梦'
WHERE id = 10;

UPDATE quests SET 
    description_zh = '帮助詹妮警官追捕逃犯', 
    reward_item_zh = '警察勋章'
WHERE id = 11;

UPDATE quests SET 
    description_zh = '探索浅红市', 
    reward_item_zh = '金币'
WHERE id = 12;

UPDATE quests SET 
    description_zh = '击败精英四及冠军', 
    reward_item_zh = '冠军奖杯'
WHERE id = 13;

-- ============================================================
-- Verification
-- ============================================================

-- Verify all translations were updated
SELECT 'NPC Dialogues with Chinese translations:' as verification;
SELECT COUNT(*) as total_with_chinese FROM npc_dialogues WHERE dialogue_text_zh IS NOT NULL AND dialogue_text_zh != '';

SELECT 'Quests with Chinese translations:' as verification;
SELECT COUNT(*) as total_with_chinese FROM quests WHERE description_zh IS NOT NULL AND description_zh != '';

-- Show sample results
SELECT 'Sample NPC Dialogue (Chinese):' as verification;
SELECT npc_id, dialogue_text_zh FROM npc_dialogues WHERE npc_id = 1 AND dialogue_text_zh IS NOT NULL LIMIT 1;

SELECT 'Sample Quest (Chinese):' as verification;
SELECT id as quest_id, description_zh, reward_item_zh FROM quests WHERE id = 1;

-- ============================================================
-- Script completed successfully
-- ============================================================
