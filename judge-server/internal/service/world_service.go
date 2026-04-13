package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/model"
)

// WorldService 世界系统服务
type WorldService struct {
	database *db.Database
}

// NewWorldService 创建世界系统服务
func NewWorldService(database *db.Database) *WorldService {
	return &WorldService{
		database: database,
	}
}

// InitializeWorldTables 初始化所有世界系统表
func (s *WorldService) InitializeWorldTables() error {
	tables := []string{
		// NPC系统表
		s.getNPCTableSQL(),
		s.getNPCDialogueTableSQL(),
		s.getNPCInteractionTableSQL(),

		// 任务系统表
		s.getQuestTableSQL(),
		s.getUserQuestTableSQL(),
		s.getQuestStepTableSQL(),

		// 地牢系统表
		s.getDungeonTableSQL(),
		s.getDungeonFloorTableSQL(),
		s.getUserDungeonProgressTableSQL(),

		// 体操馆系统表
		s.getGymTableSQL(),
		s.getGymTeamTableSQL(),
		s.getUserGymBadgeTableSQL(),

		// 地区系统表
		s.getRegionTableSQL(),
		s.getGrassAreaTableSQL(),
		s.getWildPokemonSpawnTableSQL(),
		s.getExplorationHistoryTableSQL(),

		// 地图关卡表
		s.getMapZoneTableSQL(),
		s.getLevelTableSQL(),
		s.getUserLevelProgressTableSQL(),
	}

	for _, sql := range tables {
		if _, err := s.database.Exec(sql); err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
	}

	return nil
}

// ==================== SQL表定义 ====================

func (s *WorldService) getNPCTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS npcs (
		id SERIAL PRIMARY KEY,
		town_id VARCHAR(100) NOT NULL,
		name_en VARCHAR(255) NOT NULL,
		name_zh VARCHAR(255),
		type VARCHAR(50) NOT NULL,
		role VARCHAR(255),
		coord_x FLOAT DEFAULT 0,
		coord_y FLOAT DEFAULT 0,
		avatar_url TEXT,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW(),
		FOREIGN KEY (town_id) REFERENCES towns(id)
	)
	`
}

func (s *WorldService) getNPCDialogueTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS npc_dialogues (
		id SERIAL PRIMARY KEY,
		npc_id INTEGER NOT NULL,
		dialogue_text TEXT NOT NULL,
		dialogue_type VARCHAR(50),
		condition VARCHAR(50) DEFAULT 'none',
		condition_val VARCHAR(255),
		order_idx INTEGER DEFAULT 0,
		FOREIGN KEY (npc_id) REFERENCES npcs(id)
	)
	`
}

func (s *WorldService) getNPCInteractionTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS npc_interactions (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		npc_id INTEGER NOT NULL,
		interaction_type VARCHAR(50),
		interaction_data TEXT,
		interacted_at TIMESTAMP DEFAULT NOW(),
		FOREIGN KEY (user_id) REFERENCES user_accounts(github_id),
		FOREIGN KEY (npc_id) REFERENCES npcs(id)
	)
	`
}

func (s *WorldService) getQuestTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS quests (
		id SERIAL PRIMARY KEY,
		quest_code VARCHAR(100) UNIQUE NOT NULL,
		name_en VARCHAR(255) NOT NULL,
		name_zh VARCHAR(255),
		description TEXT,
		type VARCHAR(50) DEFAULT 'side',
		difficulty INTEGER DEFAULT 1,
		giver_npc_id INTEGER,
		reward_exp INTEGER DEFAULT 0,
		reward_coins INTEGER DEFAULT 0,
		reward_item VARCHAR(255),
		level_req INTEGER DEFAULT 1,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW(),
		FOREIGN KEY (giver_npc_id) REFERENCES npcs(id)
	)
	`
}

func (s *WorldService) getUserQuestTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS user_quests (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		quest_id INTEGER NOT NULL,
		status VARCHAR(50) DEFAULT 'not_started',
		progress INTEGER DEFAULT 0,
		started_at TIMESTAMP,
		completed_at TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES user_accounts(github_id),
		FOREIGN KEY (quest_id) REFERENCES quests(id),
		UNIQUE(user_id, quest_id)
	)
	`
}

func (s *WorldService) getQuestStepTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS quest_steps (
		id SERIAL PRIMARY KEY,
		quest_id INTEGER NOT NULL,
		step_order INTEGER NOT NULL,
		step_type VARCHAR(50),
		step_target VARCHAR(255),
		step_count INTEGER DEFAULT 1,
		FOREIGN KEY (quest_id) REFERENCES quests(id)
	)
	`
}

func (s *WorldService) getDungeonTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS dungeons (
		id SERIAL PRIMARY KEY,
		dungeon_code VARCHAR(100) UNIQUE NOT NULL,
		name_en VARCHAR(255) NOT NULL,
		name_zh VARCHAR(255),
		region_id INTEGER,
		difficulty INTEGER DEFAULT 1,
		floor_count INTEGER DEFAULT 1,
		boss_pokemon_id VARCHAR(100),
		reward_exp INTEGER DEFAULT 0,
		reward_coins INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT NOW()
	)
	`
}

func (s *WorldService) getDungeonFloorTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS dungeon_floors (
		id SERIAL PRIMARY KEY,
		dungeon_id INTEGER NOT NULL,
		floor_number INTEGER NOT NULL,
		enemy_data TEXT,
		trap_data TEXT,
		treasure_data TEXT,
		boss_flag BOOLEAN DEFAULT FALSE,
		FOREIGN KEY (dungeon_id) REFERENCES dungeons(id),
		UNIQUE(dungeon_id, floor_number)
	)
	`
}

func (s *WorldService) getUserDungeonProgressTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS user_dungeon_progress (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		dungeon_id INTEGER NOT NULL,
		current_floor INTEGER DEFAULT 1,
		status VARCHAR(50) DEFAULT 'not_started',
		party_composition TEXT,
		started_at TIMESTAMP,
		completed_at TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES user_accounts(github_id),
		FOREIGN KEY (dungeon_id) REFERENCES dungeons(id),
		UNIQUE(user_id, dungeon_id)
	)
	`
}

func (s *WorldService) getGymTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS gyms (
		id SERIAL PRIMARY KEY,
		gym_code VARCHAR(100) UNIQUE NOT NULL,
		name_en VARCHAR(255) NOT NULL,
		name_zh VARCHAR(255),
		town_id VARCHAR(100) NOT NULL,
		leader_npc_id INTEGER,
		type_focus VARCHAR(50),
		badge_name VARCHAR(255),
		badge_icon TEXT,
		created_at TIMESTAMP DEFAULT NOW(),
		FOREIGN KEY (town_id) REFERENCES towns(id),
		FOREIGN KEY (leader_npc_id) REFERENCES npcs(id)
	)
	`
}

func (s *WorldService) getGymTeamTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS gym_teams (
		id SERIAL PRIMARY KEY,
		gym_id INTEGER NOT NULL,
		leader_pokemon_id VARCHAR(100),
		team_size INTEGER DEFAULT 1,
		pokemon_ids TEXT,
		FOREIGN KEY (gym_id) REFERENCES gyms(id)
	)
	`
}

func (s *WorldService) getUserGymBadgeTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS user_gym_badges (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		gym_id INTEGER NOT NULL,
		badge_earned_at TIMESTAMP DEFAULT NOW(),
		FOREIGN KEY (user_id) REFERENCES user_accounts(github_id),
		FOREIGN KEY (gym_id) REFERENCES gyms(id),
		UNIQUE(user_id, gym_id)
	)
	`
}

func (s *WorldService) getRegionTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS regions (
		id SERIAL PRIMARY KEY,
		name_en VARCHAR(255) NOT NULL,
		name_zh VARCHAR(255),
		level_min INTEGER DEFAULT 1,
		level_max INTEGER DEFAULT 30,
		island_count INTEGER DEFAULT 0
	)
	`
}

func (s *WorldService) getGrassAreaTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS grass_areas (
		id SERIAL PRIMARY KEY,
		region_id INTEGER NOT NULL,
		name_en VARCHAR(255) NOT NULL,
		name_zh VARCHAR(255),
		coord_x FLOAT DEFAULT 0,
		coord_y FLOAT DEFAULT 0,
		difficulty INTEGER DEFAULT 1,
		terrain_type VARCHAR(50),
		FOREIGN KEY (region_id) REFERENCES regions(id)
	)
	`
}

func (s *WorldService) getWildPokemonSpawnTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS wild_pokemon_spawns (
		id SERIAL PRIMARY KEY,
		area_id INTEGER NOT NULL,
		pokemon_species_id VARCHAR(100) NOT NULL,
		level_min INTEGER DEFAULT 1,
		level_max INTEGER DEFAULT 10,
		encounter_rate INTEGER DEFAULT 50,
		FOREIGN KEY (area_id) REFERENCES grass_areas(id)
	)
	`
}

func (s *WorldService) getExplorationHistoryTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS exploration_history (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		area_id INTEGER NOT NULL,
		pokemon_encountered TEXT,
		pokemon_captured TEXT,
		last_visited_at TIMESTAMP DEFAULT NOW(),
		FOREIGN KEY (user_id) REFERENCES user_accounts(github_id),
		FOREIGN KEY (area_id) REFERENCES grass_areas(id),
		UNIQUE(user_id, area_id)
	)
	`
}

func (s *WorldService) getMapZoneTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS map_zones (
		id SERIAL PRIMARY KEY,
		zone_code VARCHAR(100) UNIQUE NOT NULL,
		name_en VARCHAR(255) NOT NULL,
		name_zh VARCHAR(255),
		description TEXT,
		type VARCHAR(50),
		level INTEGER DEFAULT 1,
		island_id VARCHAR(100),
		town_id VARCHAR(100),
		connected_zones TEXT,
		created_at TIMESTAMP DEFAULT NOW(),
		FOREIGN KEY (island_id) REFERENCES islands(id),
		FOREIGN KEY (town_id) REFERENCES towns(id)
	)
	`
}

func (s *WorldService) getLevelTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS levels (
		id SERIAL PRIMARY KEY,
		level_code VARCHAR(100) UNIQUE NOT NULL,
		zone_id INTEGER NOT NULL,
		name_en VARCHAR(255) NOT NULL,
		name_zh VARCHAR(255),
		level_num INTEGER DEFAULT 1,
		type VARCHAR(50) DEFAULT 'normal',
		difficulty INTEGER DEFAULT 1,
		enemy_data TEXT,
		trap_data TEXT,
		treasure_data TEXT,
		rewards TEXT,
		created_at TIMESTAMP DEFAULT NOW(),
		FOREIGN KEY (zone_id) REFERENCES map_zones(id)
	)
	`
}

func (s *WorldService) getUserLevelProgressTableSQL() string {
	return `
	CREATE TABLE IF NOT EXISTS user_level_progress (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		level_id INTEGER NOT NULL,
		status VARCHAR(50) DEFAULT 'locked',
		score INTEGER DEFAULT 0,
		completed_at TIMESTAMP,
		attempts INTEGER DEFAULT 0,
		FOREIGN KEY (user_id) REFERENCES user_accounts(github_id),
		FOREIGN KEY (level_id) REFERENCES levels(id),
		UNIQUE(user_id, level_id)
	)
	`
}

// ==================== NPC操作 ====================

// CreateNPC 创建NPC
func (s *WorldService) CreateNPC(npc *model.NPC) (*model.NPC, error) {
	query := `
	INSERT INTO npcs (town_id, name_en, name_zh, type, role, coord_x, coord_y, avatar_url)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, town_id, name_en, name_zh, type, role, coord_x, coord_y, avatar_url, created_at, updated_at
	`

	result := &model.NPC{}
	err := s.database.QueryRow(query, npc.TownID, npc.NameEn, npc.NameZh, npc.Type, npc.Role, npc.CoordX, npc.CoordY, npc.AvatarURL).Scan(
		&result.ID, &result.TownID, &result.NameEn, &result.NameZh, &result.Type, &result.Role, &result.CoordX, &result.CoordY, &result.AvatarURL, &result.CreatedAt, &result.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create NPC: %w", err)
	}

	return result, nil
}

// GetNPC 获取NPC
func (s *WorldService) GetNPC(npcID int) (*model.NPC, error) {
	query := `SELECT id, town_id, name_en, name_zh, type, role, coord_x, coord_y, avatar_url, created_at, updated_at FROM npcs WHERE id = $1`

	npc := &model.NPC{}
	err := s.database.QueryRow(query, npcID).Scan(&npc.ID, &npc.TownID, &npc.NameEn, &npc.NameZh, &npc.Type, &npc.Role, &npc.CoordX, &npc.CoordY, &npc.AvatarURL, &npc.CreatedAt, &npc.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("NPC not found")
		}
		return nil, fmt.Errorf("failed to get NPC: %w", err)
	}

	return npc, nil
}

// GetNPCsByTown 获取城镇中的所有NPC
func (s *WorldService) GetNPCsByTown(townID string) ([]*model.NPC, error) {
	query := `SELECT id, town_id, name_en, name_zh, type, role, coord_x, coord_y, avatar_url, created_at, updated_at FROM npcs WHERE town_id = $1 ORDER BY type, name_en`

	rows, err := s.database.Query(query, townID)
	if err != nil {
		return nil, fmt.Errorf("failed to get NPCs: %w", err)
	}
	defer rows.Close()

	var npcs []*model.NPC
	for rows.Next() {
		npc := &model.NPC{}
		if err := rows.Scan(&npc.ID, &npc.TownID, &npc.NameEn, &npc.NameZh, &npc.Type, &npc.Role, &npc.CoordX, &npc.CoordY, &npc.AvatarURL, &npc.CreatedAt, &npc.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan NPC: %w", err)
		}
		npcs = append(npcs, npc)
	}

	return npcs, nil
}

// AddNPCDialogue 添加NPC对话
func (s *WorldService) AddNPCDialogue(dialogue *model.NPCDialogue) (*model.NPCDialogue, error) {
	query := `
	INSERT INTO npc_dialogues (npc_id, dialogue_text, dialogue_type, condition, condition_val, order_idx)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, npc_id, dialogue_text, dialogue_type, condition, condition_val, order_idx
	`

	result := &model.NPCDialogue{}
	err := s.database.QueryRow(query, dialogue.NPCId, dialogue.DialogueText, dialogue.DialogueType, dialogue.Condition, dialogue.ConditionVal, dialogue.Order).Scan(
		&result.ID, &result.NPCId, &result.DialogueText, &result.DialogueType, &result.Condition, &result.ConditionVal, &result.Order,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to add dialogue: %w", err)
	}

	return result, nil
}

// GetNPCDialogues 获取NPC的所有对话
func (s *WorldService) GetNPCDialogues(npcID int) ([]*model.NPCDialogue, error) {
	query := `SELECT id, npc_id, dialogue_text, dialogue_type, condition, condition_val, order_idx FROM npc_dialogues WHERE npc_id = $1 ORDER BY order_idx`

	rows, err := s.database.Query(query, npcID)
	if err != nil {
		return nil, fmt.Errorf("failed to get dialogues: %w", err)
	}
	defer rows.Close()

	var dialogues []*model.NPCDialogue
	for rows.Next() {
		dialogue := &model.NPCDialogue{}
		if err := rows.Scan(&dialogue.ID, &dialogue.NPCId, &dialogue.DialogueText, &dialogue.DialogueType, &dialogue.Condition, &dialogue.ConditionVal, &dialogue.Order); err != nil {
			return nil, fmt.Errorf("failed to scan dialogue: %w", err)
		}
		dialogues = append(dialogues, dialogue)
	}

	return dialogues, nil
}

// RecordNPCInteraction 记录NPC交互
func (s *WorldService) RecordNPCInteraction(interaction *model.NPCInteraction) (*model.NPCInteraction, error) {
	query := `
	INSERT INTO npc_interactions (user_id, npc_id, interaction_type, interaction_data)
	VALUES ($1, $2, $3, $4)
	RETURNING id, user_id, npc_id, interaction_type, interaction_data, interacted_at
	`

	result := &model.NPCInteraction{}
	err := s.database.QueryRow(query, interaction.UserID, interaction.NPCId, interaction.InteractionType, interaction.InteractionData).Scan(
		&result.ID, &result.UserID, &result.NPCId, &result.InteractionType, &result.InteractionData, &result.InteractedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to record interaction: %w", err)
	}

	return result, nil
}

// ==================== 任务操作 ====================

// CreateQuest 创建任务
func (s *WorldService) CreateQuest(quest *model.Quest) (*model.Quest, error) {
	query := `
	INSERT INTO quests (quest_code, name_en, name_zh, description, type, difficulty, giver_npc_id, reward_exp, reward_coins, reward_item, level_req)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING id, quest_code, name_en, name_zh, description, type, difficulty, giver_npc_id, reward_exp, reward_coins, reward_item, level_req, created_at, updated_at
	`

	result := &model.Quest{}
	err := s.database.QueryRow(query, quest.QuestCode, quest.NameEn, quest.NameZh, quest.Description, quest.Type, quest.Difficulty, quest.GiverNPCId, quest.RewardExp, quest.RewardCoins, quest.RewardItem, quest.LevelReq).Scan(
		&result.ID, &result.QuestCode, &result.NameEn, &result.NameZh, &result.Description, &result.Type, &result.Difficulty, &result.GiverNPCId, &result.RewardExp, &result.RewardCoins, &result.RewardItem, &result.LevelReq, &result.CreatedAt, &result.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create quest: %w", err)
	}

	return result, nil
}

// AcceptQuest 接受任务
func (s *WorldService) AcceptQuest(userID int, questID int) (*model.UserQuest, error) {
	query := `
	INSERT INTO user_quests (user_id, quest_id, status, started_at)
	VALUES ($1, $2, 'in_progress', NOW())
	ON CONFLICT (user_id, quest_id) DO UPDATE SET status = 'in_progress', started_at = NOW()
	RETURNING id, user_id, quest_id, status, progress, started_at, completed_at
	`

	result := &model.UserQuest{}
	err := s.database.QueryRow(query, userID, questID).Scan(&result.ID, &result.UserID, &result.QuestID, &result.Status, &result.Progress, &result.StartedAt, &result.CompletedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to accept quest: %w", err)
	}

	return result, nil
}

// CompleteQuest 完成任务
func (s *WorldService) CompleteQuest(userID int, questID int) (*model.UserQuest, error) {
	query := `
	UPDATE user_quests SET status = 'completed', completed_at = NOW(), progress = 100
	WHERE user_id = $1 AND quest_id = $2
	RETURNING id, user_id, quest_id, status, progress, started_at, completed_at
	`

	result := &model.UserQuest{}
	err := s.database.QueryRow(query, userID, questID).Scan(&result.ID, &result.UserID, &result.QuestID, &result.Status, &result.Progress, &result.StartedAt, &result.CompletedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to complete quest: %w", err)
	}

	return result, nil
}

// GetUserQuests 获取用户的所有任务
func (s *WorldService) GetUserQuests(userID int) ([]*model.UserQuest, error) {
	query := `
	SELECT id, user_id, quest_id, status, progress, started_at, completed_at 
	FROM user_quests 
	WHERE user_id = $1 
	ORDER BY started_at DESC
	`

	rows, err := s.database.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user quests: %w", err)
	}
	defer rows.Close()

	var quests []*model.UserQuest
	for rows.Next() {
		quest := &model.UserQuest{}
		if err := rows.Scan(&quest.ID, &quest.UserID, &quest.QuestID, &quest.Status, &quest.Progress, &quest.StartedAt, &quest.CompletedAt); err != nil {
			return nil, fmt.Errorf("failed to scan quest: %w", err)
		}
		quests = append(quests, quest)
	}

	return quests, nil
}

// ==================== 地牢操作 ====================

// CreateDungeon 创建地牢
func (s *WorldService) CreateDungeon(dungeon *model.Dungeon) (*model.Dungeon, error) {
	query := `
	INSERT INTO dungeons (dungeon_code, name_en, name_zh, region_id, difficulty, floor_count, boss_pokemon_id, reward_exp, reward_coins)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id, dungeon_code, name_en, name_zh, region_id, difficulty, floor_count, boss_pokemon_id, reward_exp, reward_coins, created_at
	`

	result := &model.Dungeon{}
	err := s.database.QueryRow(query, dungeon.DungeonCode, dungeon.NameEn, dungeon.NameZh, dungeon.RegionID, dungeon.Difficulty, dungeon.FloorCount, dungeon.BossPokemonID, dungeon.RewardExp, dungeon.RewardCoins).Scan(
		&result.ID, &result.DungeonCode, &result.NameEn, &result.NameZh, &result.RegionID, &result.Difficulty, &result.FloorCount, &result.BossPokemonID, &result.RewardExp, &result.RewardCoins, &result.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create dungeon: %w", err)
	}

	return result, nil
}

// GetUserDungeonProgress 获取用户地牢进度
func (s *WorldService) GetUserDungeonProgress(userID int, dungeonID int) (*model.UserDungeonProgress, error) {
	query := `
	SELECT id, user_id, dungeon_id, current_floor, status, party_composition, started_at, completed_at
	FROM user_dungeon_progress
	WHERE user_id = $1 AND dungeon_id = $2
	`

	progress := &model.UserDungeonProgress{}
	err := s.database.QueryRow(query, userID, dungeonID).Scan(&progress.ID, &progress.UserID, &progress.DungeonID, &progress.CurrentFloor, &progress.Status, &progress.PartyComposition, &progress.StartedAt, &progress.CompletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			// 创建新进度记录
			insertQuery := `
			INSERT INTO user_dungeon_progress (user_id, dungeon_id, current_floor, status, started_at)
			VALUES ($1, $2, 1, 'in_progress', NOW())
			RETURNING id, user_id, dungeon_id, current_floor, status, party_composition, started_at, completed_at
			`
			err = s.database.QueryRow(insertQuery, userID, dungeonID).Scan(&progress.ID, &progress.UserID, &progress.DungeonID, &progress.CurrentFloor, &progress.Status, &progress.PartyComposition, &progress.StartedAt, &progress.CompletedAt)
			if err != nil {
				return nil, fmt.Errorf("failed to create dungeon progress: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to get dungeon progress: %w", err)
		}
	}

	return progress, nil
}

// ==================== 体操馆操作 ====================

// CreateGym 创建体操馆
func (s *WorldService) CreateGym(gym *model.Gym) (*model.Gym, error) {
	query := `
	INSERT INTO gyms (gym_code, name_en, name_zh, town_id, leader_npc_id, type_focus, badge_name, badge_icon)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, gym_code, name_en, name_zh, town_id, leader_npc_id, type_focus, badge_name, badge_icon, created_at
	`

	result := &model.Gym{}
	err := s.database.QueryRow(query, gym.GymCode, gym.NameEn, gym.NameZh, gym.TownID, gym.LeaderNPCId, gym.TypeFocus, gym.BadgeName, gym.BadgeIcon).Scan(
		&result.ID, &result.GymCode, &result.NameEn, &result.NameZh, &result.TownID, &result.LeaderNPCId, &result.TypeFocus, &result.BadgeName, &result.BadgeIcon, &result.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create gym: %w", err)
	}

	return result, nil
}

// GetUserGymBadges 获取用户获得的体操馆徽章
func (s *WorldService) GetUserGymBadges(userID int) ([]*model.UserGymBadge, error) {
	query := `
	SELECT id, user_id, gym_id, badge_earned_at
	FROM user_gym_badges
	WHERE user_id = $1
	ORDER BY badge_earned_at DESC
	`

	rows, err := s.database.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get gym badges: %w", err)
	}
	defer rows.Close()

	var badges []*model.UserGymBadge
	for rows.Next() {
		badge := &model.UserGymBadge{}
		if err := rows.Scan(&badge.ID, &badge.UserID, &badge.GymID, &badge.BadgeEarnedAt); err != nil {
			return nil, fmt.Errorf("failed to scan badge: %w", err)
		}
		badges = append(badges, badge)
	}

	return badges, nil
}

// AwardGymBadge 颁发体操馆徽章
func (s *WorldService) AwardGymBadge(userID int, gymID int) (*model.UserGymBadge, error) {
	query := `
	INSERT INTO user_gym_badges (user_id, gym_id, badge_earned_at)
	VALUES ($1, $2, NOW())
	ON CONFLICT (user_id, gym_id) DO NOTHING
	RETURNING id, user_id, gym_id, badge_earned_at
	`

	result := &model.UserGymBadge{}
	err := s.database.QueryRow(query, userID, gymID).Scan(&result.ID, &result.UserID, &result.GymID, &result.BadgeEarnedAt)

	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to award badge: %w", err)
	}

	return result, nil
}

// ==================== 地图关卡操作 ====================

// CreateMapZone 创建地图区域
func (s *WorldService) CreateMapZone(zone *model.MapZone) (*model.MapZone, error) {
	query := `
	INSERT INTO map_zones (zone_code, name_en, name_zh, description, type, level, island_id, town_id, connected_zones)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id, zone_code, name_en, name_zh, description, type, level, island_id, town_id, connected_zones, created_at
	`

	result := &model.MapZone{}
	connectedZonesJSON, _ := json.Marshal(zone.ConnectedZones)
	err := s.database.QueryRow(query, zone.ZoneCode, zone.NameEn, zone.NameZh, zone.Description, zone.Type, zone.Level, zone.IslandID, zone.TownID, string(connectedZonesJSON)).Scan(
		&result.ID, &result.ZoneCode, &result.NameEn, &result.NameZh, &result.Description, &result.Type, &result.Level, &result.IslandID, &result.TownID, &result.ConnectedZones, &result.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create map zone: %w", err)
	}

	return result, nil
}

// CreateLevel 创建关卡
func (s *WorldService) CreateLevel(level *model.Level) (*model.Level, error) {
	query := `
	INSERT INTO levels (level_code, zone_id, name_en, name_zh, level_num, type, difficulty, enemy_data, trap_data, treasure_data, rewards)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING id, level_code, zone_id, name_en, name_zh, level_num, type, difficulty, enemy_data, trap_data, treasure_data, rewards, created_at
	`

	result := &model.Level{}
	err := s.database.QueryRow(query, level.LevelCode, level.ZoneID, level.NameEn, level.NameZh, level.LevelNum, level.Type, level.Difficulty, level.EnemyData, level.TrapData, level.TreasureData, level.Rewards).Scan(
		&result.ID, &result.LevelCode, &result.ZoneID, &result.NameEn, &result.NameZh, &result.LevelNum, &result.Type, &result.Difficulty, &result.EnemyData, &result.TrapData, &result.TreasureData, &result.Rewards, &result.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create level: %w", err)
	}

	return result, nil
}

// GetLevelsByZone 获取区域内的所有关卡
func (s *WorldService) GetLevelsByZone(zoneID int) ([]*model.Level, error) {
	query := `
	SELECT id, level_code, zone_id, name_en, name_zh, level_num, type, difficulty, enemy_data, trap_data, treasure_data, rewards, created_at
	FROM levels
	WHERE zone_id = $1
	ORDER BY level_num
	`

	rows, err := s.database.Query(query, zoneID)
	if err != nil {
		return nil, fmt.Errorf("failed to get levels: %w", err)
	}
	defer rows.Close()

	var levels []*model.Level
	for rows.Next() {
		level := &model.Level{}
		if err := rows.Scan(&level.ID, &level.LevelCode, &level.ZoneID, &level.NameEn, &level.NameZh, &level.LevelNum, &level.Type, &level.Difficulty, &level.EnemyData, &level.TrapData, &level.TreasureData, &level.Rewards, &level.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan level: %w", err)
		}
		levels = append(levels, level)
	}

	return levels, nil
}

// CompleteLevel 完成关卡
func (s *WorldService) CompleteLevel(userID int, levelID int, score int) (*model.UserLevelProgress, error) {
	query := `
	INSERT INTO user_level_progress (user_id, level_id, status, score, completed_at, attempts)
	VALUES ($1, $2, 'completed', $3, NOW(), 1)
	ON CONFLICT (user_id, level_id) DO UPDATE SET status = 'completed', score = GREATEST(score, EXCLUDED.score), completed_at = NOW(), attempts = attempts + 1
	RETURNING id, user_id, level_id, status, score, completed_at, attempts
	`

	result := &model.UserLevelProgress{}
	err := s.database.QueryRow(query, userID, levelID, score).Scan(&result.ID, &result.UserID, &result.LevelID, &result.Status, &result.Score, &result.CompletedAt, &result.Attempts)

	if err != nil {
		return nil, fmt.Errorf("failed to complete level: %w", err)
	}

	return result, nil
}

// ==================== 任务查询操作 ====================

// GetAllQuests 获取所有任务（可按类型和难度过滤）
func (s *WorldService) GetAllQuests(questType string, difficulty int) ([]*model.Quest, error) {
	var query string
	var args []interface{}

	query = `SELECT id, quest_code, name_en, name_zh, description, type, difficulty, giver_npc_id, reward_exp, reward_coins, reward_item, level_req, created_at, updated_at FROM quests WHERE 1=1`

	if questType != "" {
		query += ` AND type = $1`
		args = append(args, questType)
	}

	if difficulty > 0 {
		argIdx := len(args) + 1
		query += fmt.Sprintf(` AND difficulty = $%d`, argIdx)
		args = append(args, difficulty)
	}

	query += ` ORDER BY difficulty, created_at DESC`

	rows, err := s.database.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get quests: %w", err)
	}
	defer rows.Close()

	var quests []*model.Quest
	for rows.Next() {
		quest := &model.Quest{}
		if err := rows.Scan(&quest.ID, &quest.QuestCode, &quest.NameEn, &quest.NameZh, &quest.Description, &quest.Type, &quest.Difficulty, &quest.GiverNPCId, &quest.RewardExp, &quest.RewardCoins, &quest.RewardItem, &quest.LevelReq, &quest.CreatedAt, &quest.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan quest: %w", err)
		}
		quests = append(quests, quest)
	}

	return quests, nil
}

// GetQuest 获取单个任务
func (s *WorldService) GetQuest(questID int) (*model.Quest, error) {
	query := `SELECT id, quest_code, name_en, name_zh, description, type, difficulty, giver_npc_id, reward_exp, reward_coins, reward_item, level_req, created_at, updated_at FROM quests WHERE id = $1`

	quest := &model.Quest{}
	err := s.database.QueryRow(query, questID).Scan(&quest.ID, &quest.QuestCode, &quest.NameEn, &quest.NameZh, &quest.Description, &quest.Type, &quest.Difficulty, &quest.GiverNPCId, &quest.RewardExp, &quest.RewardCoins, &quest.RewardItem, &quest.LevelReq, &quest.CreatedAt, &quest.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("quest not found")
		}
		return nil, fmt.Errorf("failed to get quest: %w", err)
	}

	return quest, nil
}

// ==================== 地牢查询操作 ====================

// GetAllDungeons 获取所有地牢（可按难度过滤）
func (s *WorldService) GetAllDungeons(difficulty int) ([]*model.Dungeon, error) {
	var query string
	var args []interface{}

	query = `SELECT id, dungeon_code, name_en, name_zh, region_id, difficulty, floor_count, boss_pokemon_id, reward_exp, reward_coins, created_at FROM dungeons WHERE 1=1`

	if difficulty > 0 {
		query += ` AND difficulty = $1`
		args = append(args, difficulty)
	}

	query += ` ORDER BY difficulty, created_at DESC`

	rows, err := s.database.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get dungeons: %w", err)
	}
	defer rows.Close()

	var dungeons []*model.Dungeon
	for rows.Next() {
		dungeon := &model.Dungeon{}
		if err := rows.Scan(&dungeon.ID, &dungeon.DungeonCode, &dungeon.NameEn, &dungeon.NameZh, &dungeon.RegionID, &dungeon.Difficulty, &dungeon.FloorCount, &dungeon.BossPokemonID, &dungeon.RewardExp, &dungeon.RewardCoins, &dungeon.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan dungeon: %w", err)
		}
		dungeons = append(dungeons, dungeon)
	}

	return dungeons, nil
}

// GetDungeon 获取单个地牢
func (s *WorldService) GetDungeon(dungeonID int) (*model.Dungeon, error) {
	query := `SELECT id, dungeon_code, name_en, name_zh, region_id, difficulty, floor_count, boss_pokemon_id, reward_exp, reward_coins, created_at FROM dungeons WHERE id = $1`

	dungeon := &model.Dungeon{}
	err := s.database.QueryRow(query, dungeonID).Scan(&dungeon.ID, &dungeon.DungeonCode, &dungeon.NameEn, &dungeon.NameZh, &dungeon.RegionID, &dungeon.Difficulty, &dungeon.FloorCount, &dungeon.BossPokemonID, &dungeon.RewardExp, &dungeon.RewardCoins, &dungeon.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("dungeon not found")
		}
		return nil, fmt.Errorf("failed to get dungeon: %w", err)
	}

	return dungeon, nil
}

// ==================== 体操馆查询操作 ====================

// GetGymsByTown 获取城镇中的所有体操馆
func (s *WorldService) GetGymsByTown(townID string) ([]*model.Gym, error) {
	query := `SELECT id, gym_code, name_en, name_zh, town_id, leader_npc_id, type_focus, badge_name, badge_icon, created_at FROM gyms WHERE town_id = $1 ORDER BY created_at`

	rows, err := s.database.Query(query, townID)
	if err != nil {
		return nil, fmt.Errorf("failed to get gyms: %w", err)
	}
	defer rows.Close()

	var gyms []*model.Gym
	for rows.Next() {
		gym := &model.Gym{}
		if err := rows.Scan(&gym.ID, &gym.GymCode, &gym.NameEn, &gym.NameZh, &gym.TownID, &gym.LeaderNPCId, &gym.TypeFocus, &gym.BadgeName, &gym.BadgeIcon, &gym.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan gym: %w", err)
		}
		gyms = append(gyms, gym)
	}

	return gyms, nil
}

// GetGym 获取单个体操馆
func (s *WorldService) GetGym(gymID int) (*model.Gym, error) {
	query := `SELECT id, gym_code, name_en, name_zh, town_id, leader_npc_id, type_focus, badge_name, badge_icon, created_at FROM gyms WHERE id = $1`

	gym := &model.Gym{}
	err := s.database.QueryRow(query, gymID).Scan(&gym.ID, &gym.GymCode, &gym.NameEn, &gym.NameZh, &gym.TownID, &gym.LeaderNPCId, &gym.TypeFocus, &gym.BadgeName, &gym.BadgeIcon, &gym.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("gym not found")
		}
		return nil, fmt.Errorf("failed to get gym: %w", err)
	}

	return gym, nil
}

// ==================== 地图区域查询操作 ====================

// GetMapZonesByIsland 获取岛屿上的所有地图区域
func (s *WorldService) GetMapZonesByIsland(islandID string) ([]*model.MapZone, error) {
	query := `SELECT id, zone_code, name_en, name_zh, description, type, level, island_id, town_id, connected_zones, created_at FROM map_zones WHERE island_id = $1 ORDER BY level, created_at`

	rows, err := s.database.Query(query, islandID)
	if err != nil {
		return nil, fmt.Errorf("failed to get map zones: %w", err)
	}
	defer rows.Close()

	var zones []*model.MapZone
	for rows.Next() {
		zone := &model.MapZone{}
		var connectedZonesStr string
		if err := rows.Scan(&zone.ID, &zone.ZoneCode, &zone.NameEn, &zone.NameZh, &zone.Description, &zone.Type, &zone.Level, &zone.IslandID, &zone.TownID, &connectedZonesStr, &zone.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan map zone: %w", err)
		}

		// Parse connected zones JSON
		if connectedZonesStr != "" {
			if err := json.Unmarshal([]byte(connectedZonesStr), &zone.ConnectedZones); err != nil {
				return nil, fmt.Errorf("failed to parse connected zones: %w", err)
			}
		}

		zones = append(zones, zone)
	}

	return zones, nil
}

// GetMapZone 获取单个地图区域
func (s *WorldService) GetMapZone(zoneID int) (*model.MapZone, error) {
	query := `SELECT id, zone_code, name_en, name_zh, description, type, level, island_id, town_id, connected_zones, created_at FROM map_zones WHERE id = $1`

	zone := &model.MapZone{}
	var connectedZonesStr string
	err := s.database.QueryRow(query, zoneID).Scan(&zone.ID, &zone.ZoneCode, &zone.NameEn, &zone.NameZh, &zone.Description, &zone.Type, &zone.Level, &zone.IslandID, &zone.TownID, &connectedZonesStr, &zone.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("map zone not found")
		}
		return nil, fmt.Errorf("failed to get map zone: %w", err)
	}

	// Parse connected zones JSON
	if connectedZonesStr != "" {
		if err := json.Unmarshal([]byte(connectedZonesStr), &zone.ConnectedZones); err != nil {
			return nil, fmt.Errorf("failed to parse connected zones: %w", err)
		}
	}

	return zone, nil
}

// GetUserLevelProgress 获取用户的关卡进度
func (s *WorldService) GetUserLevelProgress(userID int) ([]*model.UserLevelProgress, error) {
	query := `
	SELECT id, user_id, level_id, status, score, completed_at, attempts
	FROM user_level_progress
	WHERE user_id = $1
	ORDER BY completed_at DESC
	`

	rows, err := s.database.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user level progress: %w", err)
	}
	defer rows.Close()

	var progresses []*model.UserLevelProgress
	for rows.Next() {
		progress := &model.UserLevelProgress{}
		if err := rows.Scan(&progress.ID, &progress.UserID, &progress.LevelID, &progress.Status, &progress.Score, &progress.CompletedAt, &progress.Attempts); err != nil {
			return nil, fmt.Errorf("failed to scan level progress: %w", err)
		}
		progresses = append(progresses, progress)
	}

	return progresses, nil
}
