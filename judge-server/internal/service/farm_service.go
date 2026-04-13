package service

import (
	"database/sql"
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/model"
	"time"
)

// FarmService 菜园服务
type FarmService struct {
	database *db.Database
}

// NewFarmService 创建菜园服务
func NewFarmService(database *db.Database) *FarmService {
	return &FarmService{
		database: database,
	}
}

// InitializeFarmTables 初始化菜园相关表
func (s *FarmService) InitializeFarmTables() error {
	// 创建 base_farm 表（基地菜园地块）
	baseFarmSQL := `
	CREATE TABLE IF NOT EXISTS base_farm (
		id SERIAL PRIMARY KEY,
		base_id VARCHAR(255) NOT NULL,
		github_id INTEGER NOT NULL,
		plot_number INTEGER NOT NULL,
		crop_type VARCHAR(100),
		planted_at TIMESTAMP,
		ready_at TIMESTAMP,
		quantity INTEGER DEFAULT 0,
		status VARCHAR(20) DEFAULT 'empty',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(base_id, plot_number)
	)
	`

	// 创建 farm_harvest 表（菜园收获）
	farmHarvestSQL := `
	CREATE TABLE IF NOT EXISTS farm_harvest (
		id SERIAL PRIMARY KEY,
		base_id VARCHAR(255) NOT NULL,
		github_id INTEGER NOT NULL,
		crop_type VARCHAR(100) NOT NULL,
		quantity INTEGER NOT NULL,
		harvested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	// 创建 base_visit 表（基地访问）
	baseVisitSQL := `
	CREATE TABLE IF NOT EXISTS base_visit (
		id SERIAL PRIMARY KEY,
		base_id VARCHAR(255) NOT NULL,
		owner_github_id INTEGER NOT NULL,
		visitor_github_id INTEGER NOT NULL,
		visit_type VARCHAR(20) NOT NULL,
		visitor_pet_id VARCHAR(255),
		defender_pet_id VARCHAR(255),
		battle_result VARCHAR(20),
		points_gained INTEGER DEFAULT 0,
		reputation_change INTEGER DEFAULT 0,
		coins_looted INTEGER DEFAULT 0,
		visited_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	// 创建 farm_eating 表（进食记录）
	farmEatingSQL := `
	CREATE TABLE IF NOT EXISTS farm_eating (
		id SERIAL PRIMARY KEY,
		base_id VARCHAR(255) NOT NULL,
		owner_github_id INTEGER NOT NULL,
		visitor_github_id INTEGER NOT NULL,
		visitor_pet_id VARCHAR(255) NOT NULL,
		crop_type VARCHAR(100) NOT NULL,
		hp_recovered INTEGER DEFAULT 0,
		eaten_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	// 创建 user_reputation 表（用户声望）
	userReputationSQL := `
	CREATE TABLE IF NOT EXISTS user_reputation (
		id SERIAL PRIMARY KEY,
		github_id INTEGER NOT NULL UNIQUE,
		total_points INTEGER DEFAULT 0,
		total_reputation INTEGER DEFAULT 0,
		defense_wins INTEGER DEFAULT 0,
		defense_losses INTEGER DEFAULT 0,
		attack_wins INTEGER DEFAULT 0,
		attack_losses INTEGER DEFAULT 0,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	// 执行所有创建表的SQL语句
	tables := []string{
		baseFarmSQL,
		farmHarvestSQL,
		baseVisitSQL,
		farmEatingSQL,
		userReputationSQL,
	}

	for _, sql := range tables {
		if _, err := s.database.Exec(sql); err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
	}

	return nil
}

// PlantCrop 种植作物
func (s *FarmService) PlantCrop(baseID string, githubID int, plotNumber int, cropType string, quantity int, userLevel int) (*model.BaseFarmPlot, error) {
	// 验证作物类型
	crop, exists := model.Crops[cropType]
	if !exists {
		return nil, fmt.Errorf("crop type %s does not exist", cropType)
	}

	// 验证用户等级
	if userLevel < crop.RequiredLevel {
		return nil, fmt.Errorf("user level %d is below required level %d for crop %s", userLevel, crop.RequiredLevel, cropType)
	}

	// 验证地块编号
	if plotNumber < 1 || plotNumber > 8 {
		return nil, fmt.Errorf("plot number must be between 1 and 8")
	}

	// 验证数量
	if quantity <= 0 {
		return nil, fmt.Errorf("quantity must be greater than 0")
	}

	// 计算成熟时间
	now := time.Now()
	readyAt := now.Add(time.Duration(crop.RecoveryTime) * time.Second)

	// 插入或更新地块
	query := `
	INSERT INTO base_farm (base_id, github_id, plot_number, crop_type, planted_at, ready_at, quantity, status, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())
	ON CONFLICT (base_id, plot_number) DO UPDATE SET
		github_id = $2,
		crop_type = $4,
		planted_at = $5,
		ready_at = $6,
		quantity = $7,
		status = $8,
		updated_at = NOW()
	RETURNING id, base_id, github_id, plot_number, crop_type, planted_at, ready_at, quantity, status, created_at, updated_at
	`

	plot := &model.BaseFarmPlot{}
	err := s.database.QueryRow(query, baseID, githubID, plotNumber, cropType, now, readyAt, quantity, "growing").Scan(
		&plot.ID,
		&plot.BaseID,
		&plot.GitHubID,
		&plot.PlotNumber,
		&plot.CropType,
		&plot.PlantedAt,
		&plot.ReadyAt,
		&plot.Quantity,
		&plot.Status,
		&plot.CreatedAt,
		&plot.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to plant crop: %w", err)
	}

	return plot, nil
}

// HarvestCrop 收获作物
func (s *FarmService) HarvestCrop(baseID string, plotNumber int) (*model.FarmHarvest, error) {
	// 首先查询地块信息
	plotQuery := `
	SELECT id, github_id, crop_type, quantity, ready_at, status
	FROM base_farm
	WHERE base_id = $1 AND plot_number = $2
	`

	var plotID int
	var githubID int
	var cropType string
	var quantity int
	var readyAt time.Time
	var status string

	err := s.database.QueryRow(plotQuery, baseID, plotNumber).Scan(&plotID, &githubID, &cropType, &quantity, &readyAt, &status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("plot not found")
		}
		return nil, fmt.Errorf("failed to query plot: %w", err)
	}

	// 检查地块是否准备好收获
	if status != "ready" && time.Now().Before(readyAt) {
		return nil, fmt.Errorf("crop is not ready for harvest yet")
	}

	if cropType == "" || quantity == 0 {
		return nil, fmt.Errorf("plot is empty")
	}

	// 记录收获
	harvestQuery := `
	INSERT INTO farm_harvest (base_id, github_id, crop_type, quantity)
	VALUES ($1, $2, $3, $4)
	RETURNING id, base_id, github_id, crop_type, quantity, harvested_at
	`

	harvest := &model.FarmHarvest{}
	err = s.database.QueryRow(harvestQuery, baseID, githubID, cropType, quantity).Scan(
		&harvest.ID,
		&harvest.BaseID,
		&harvest.GitHubID,
		&harvest.CropType,
		&harvest.Quantity,
		&harvest.HarvestedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to record harvest: %w", err)
	}

	// 清空地块
	updateQuery := `
	UPDATE base_farm
	SET crop_type = NULL, quantity = 0, status = 'empty', updated_at = NOW()
	WHERE base_id = $1 AND plot_number = $2
	`

	if _, err := s.database.Exec(updateQuery, baseID, plotNumber); err != nil {
		return nil, fmt.Errorf("failed to clear plot: %w", err)
	}

	return harvest, nil
}

// GetFarmStatus 查看菜园状态
func (s *FarmService) GetFarmStatus(baseID string) ([]*model.BaseFarmPlot, error) {
	query := `
	SELECT id, base_id, github_id, plot_number, crop_type, planted_at, ready_at, quantity, status, created_at, updated_at
	FROM base_farm
	WHERE base_id = $1
	ORDER BY plot_number
	`

	rows, err := s.database.Query(query, baseID)
	if err != nil {
		return nil, fmt.Errorf("failed to query farm status: %w", err)
	}
	defer rows.Close()

	plots := make([]*model.BaseFarmPlot, 0)
	for rows.Next() {
		plot := &model.BaseFarmPlot{}
		err := rows.Scan(
			&plot.ID,
			&plot.BaseID,
			&plot.GitHubID,
			&plot.PlotNumber,
			&plot.CropType,
			&plot.PlantedAt,
			&plot.ReadyAt,
			&plot.Quantity,
			&plot.Status,
			&plot.CreatedAt,
			&plot.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan farm plot: %w", err)
		}

		// 更新状态：检查是否成熟
		if plot.Status == "growing" && time.Now().After(plot.ReadyAt) {
			plot.Status = "ready"
			// 更新数据库中的状态
			updateQuery := `UPDATE base_farm SET status = 'ready' WHERE id = $1`
			if _, err := s.database.Exec(updateQuery, plot.ID); err != nil {
				fmt.Printf("warning: failed to update plot status: %v\n", err)
			}
		}

		plots = append(plots, plot)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating farm plots: %w", err)
	}

	return plots, nil
}

// GetReadyPlots 获取已成熟的地块
func (s *FarmService) GetReadyPlots(baseID string) ([]*model.BaseFarmPlot, error) {
	query := `
	SELECT id, base_id, github_id, plot_number, crop_type, planted_at, ready_at, quantity, status, created_at, updated_at
	FROM base_farm
	WHERE base_id = $1 AND (status = 'ready' OR (status = 'growing' AND ready_at <= NOW()))
	ORDER BY plot_number
	`

	rows, err := s.database.Query(query, baseID)
	if err != nil {
		return nil, fmt.Errorf("failed to query ready plots: %w", err)
	}
	defer rows.Close()

	plots := make([]*model.BaseFarmPlot, 0)
	for rows.Next() {
		plot := &model.BaseFarmPlot{}
		err := rows.Scan(
			&plot.ID,
			&plot.BaseID,
			&plot.GitHubID,
			&plot.PlotNumber,
			&plot.CropType,
			&plot.PlantedAt,
			&plot.ReadyAt,
			&plot.Quantity,
			&plot.Status,
			&plot.CreatedAt,
			&plot.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ready plot: %w", err)
		}

		plot.Status = "ready"
		plots = append(plots, plot)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating ready plots: %w", err)
	}

	return plots, nil
}

// AllowEating 允许进食
func (s *FarmService) AllowEating(baseID string, visitorGitHubID int, cropType string, visitorPetID string) (*model.FarmEating, error) {
	// 验证作物类型
	crop, exists := model.Crops[cropType]
	if !exists {
		return nil, fmt.Errorf("crop type %s does not exist", cropType)
	}

	// 查询菜园是否有该作物可用
	plotQuery := `
	SELECT github_id FROM base_farm
	WHERE base_id = $1 AND crop_type = $2 AND status = 'ready' AND quantity > 0
	LIMIT 1
	`

	var ownerGitHubID int
	err := s.database.QueryRow(plotQuery, baseID, cropType).Scan(&ownerGitHubID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no %s available for eating at this farm", cropType)
		}
		return nil, fmt.Errorf("failed to check crop availability: %w", err)
	}

	// 记录进食
	hpRecovered := crop.HPRecovery
	insertQuery := `
	INSERT INTO farm_eating (base_id, owner_github_id, visitor_github_id, visitor_pet_id, crop_type, hp_recovered)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, base_id, owner_github_id, visitor_github_id, visitor_pet_id, crop_type, hp_recovered, eaten_at
	`

	eating := &model.FarmEating{}
	err = s.database.QueryRow(insertQuery, baseID, ownerGitHubID, visitorGitHubID, visitorPetID, cropType, hpRecovered).Scan(
		&eating.ID,
		&eating.BaseID,
		&eating.OwnerGitHubID,
		&eating.VisitorGitHubID,
		&eating.VisitorPetID,
		&eating.CropType,
		&eating.HPRecovered,
		&eating.EatenAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to record eating: %w", err)
	}

	// 减少地块的作物数量
	updateQuery := `
	UPDATE base_farm
	SET quantity = quantity - 1
	WHERE base_id = $1 AND crop_type = $2 AND quantity > 0
	`

	if _, err := s.database.Exec(updateQuery, baseID, cropType); err != nil {
		return nil, fmt.Errorf("failed to update crop quantity: %w", err)
	}

	return eating, nil
}

// GetFarmFood 获取可用食物列表
func (s *FarmService) GetFarmFood(baseID string) ([]string, error) {
	query := `
	SELECT DISTINCT crop_type
	FROM base_farm
	WHERE base_id = $1 AND status = 'ready' AND quantity > 0
	`

	rows, err := s.database.Query(query, baseID)
	if err != nil {
		return nil, fmt.Errorf("failed to query farm food: %w", err)
	}
	defer rows.Close()

	foods := make([]string, 0)
	for rows.Next() {
		var cropType string
		err := rows.Scan(&cropType)
		if err != nil {
			return nil, fmt.Errorf("failed to scan crop type: %w", err)
		}
		foods = append(foods, cropType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating farm foods: %w", err)
	}

	return foods, nil
}

// InitializeReputation 初始化声望
func (s *FarmService) InitializeReputation(githubID int) (*model.UserReputation, error) {
	// 检查声望是否已存在
	checkQuery := `SELECT id FROM user_reputation WHERE github_id = $1`
	var existingID int
	err := s.database.QueryRow(checkQuery, githubID).Scan(&existingID)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to check existing reputation: %w", err)
	}

	// 如果已存在，直接返回
	if err == nil {
		return s.GetReputation(githubID)
	}

	// 初始化新的声望记录
	insertQuery := `
	INSERT INTO user_reputation (github_id, total_points, total_reputation, defense_wins, defense_losses, attack_wins, attack_losses, updated_at)
	VALUES ($1, 0, 0, 0, 0, 0, 0, NOW())
	RETURNING id, github_id, total_points, total_reputation, defense_wins, defense_losses, attack_wins, attack_losses, updated_at
	`

	reputation := &model.UserReputation{}
	err = s.database.QueryRow(insertQuery, githubID).Scan(
		&reputation.ID,
		&reputation.GitHubID,
		&reputation.TotalPoints,
		&reputation.TotalReputation,
		&reputation.DefenseWins,
		&reputation.DefenseLosses,
		&reputation.AttackWins,
		&reputation.AttackLosses,
		&reputation.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to initialize reputation: %w", err)
	}

	return reputation, nil
}

// GetReputation 获取声望
func (s *FarmService) GetReputation(githubID int) (*model.UserReputation, error) {
	query := `
	SELECT id, github_id, total_points, total_reputation, defense_wins, defense_losses, attack_wins, attack_losses, updated_at
	FROM user_reputation
	WHERE github_id = $1
	`

	reputation := &model.UserReputation{}
	err := s.database.QueryRow(query, githubID).Scan(
		&reputation.ID,
		&reputation.GitHubID,
		&reputation.TotalPoints,
		&reputation.TotalReputation,
		&reputation.DefenseWins,
		&reputation.DefenseLosses,
		&reputation.AttackWins,
		&reputation.AttackLosses,
		&reputation.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("reputation not found for github_id %d", githubID)
		}
		return nil, fmt.Errorf("failed to get reputation: %w", err)
	}

	return reputation, nil
}

// UpdateReputation 更新声望
func (s *FarmService) UpdateReputation(githubID int, pointsChange int, reputationChange int) error {
	// 首先确保声望记录存在
	_, err := s.GetReputation(githubID)
	if err != nil {
		// 如果不存在，初始化
		_, initErr := s.InitializeReputation(githubID)
		if initErr != nil {
			return fmt.Errorf("failed to initialize reputation: %w", initErr)
		}
	}

	// 更新声望
	updateQuery := `
	UPDATE user_reputation
	SET total_points = total_points + $1,
		total_reputation = total_reputation + $2,
		updated_at = NOW()
	WHERE github_id = $3
	`

	result, err := s.database.Exec(updateQuery, pointsChange, reputationChange, githubID)
	if err != nil {
		return fmt.Errorf("failed to update reputation: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no reputation record found for github_id %d", githubID)
	}

	return nil
}

// RecordBaseVisit 记录基地访问和战斗结果
func (s *FarmService) RecordBaseVisit(baseID string, ownerGitHubID int, visitorGitHubID int, visitType string, visitorPetID string, defenderPetID string, battleResult string, pointsGained int, reputationChange int, coinsLooted int) (*model.BaseVisit, error) {
	// 验证访问类型
	if visitType != model.VisitTypeDuel && visitType != model.VisitTypeAttack {
		return nil, fmt.Errorf("invalid visit type: %s", visitType)
	}

	// 验证战斗结果
	if battleResult != model.BattleResultWin && battleResult != model.BattleResultLose && battleResult != model.BattleResultDraw {
		return nil, fmt.Errorf("invalid battle result: %s", battleResult)
	}

	// 插入访问记录
	query := `
	INSERT INTO base_visit (base_id, owner_github_id, visitor_github_id, visit_type, visitor_pet_id, defender_pet_id, battle_result, points_gained, reputation_change, coins_looted)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id, base_id, owner_github_id, visitor_github_id, visit_type, visitor_pet_id, defender_pet_id, battle_result, points_gained, reputation_change, coins_looted, visited_at
	`

	visit := &model.BaseVisit{}
	err := s.database.QueryRow(query, baseID, ownerGitHubID, visitorGitHubID, visitType, visitorPetID, defenderPetID, battleResult, pointsGained, reputationChange, coinsLooted).Scan(
		&visit.ID,
		&visit.BaseID,
		&visit.OwnerGitHubID,
		&visit.VisitorGitHubID,
		&visit.VisitType,
		&visit.VisitorPetID,
		&visit.DefenderPetID,
		&visit.BattleResult,
		&visit.PointsGained,
		&visit.ReputationChange,
		&visit.CoinsLooted,
		&visit.VisitedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to record visit: %w", err)
	}

	return visit, nil
}

// CalculateRewards 根据访问类型和战斗结果计算奖励/处罚
func CalculateRewards(visitType string, battleResult string) (visitorPoints int, visitorReputation int, ownerPoints int, ownerReputation int, coinsLoot int) {
	// 根据设计：
	// Duel Winner: +2 points, +1 reputation, +50 XP -> +2 points
	// Duel Loser: +20 XP -> +1 point
	// Attack Winner (visitor): +5 points, +3 reputation, +100 XP -> +5 points, 5-10% coin loot
	// Attack Winner (defender): +3 points, +1 reputation, +80 XP -> +3 points
	// Attack Loser (owner): -2 reputation, +1 defense loss

	if visitType == model.VisitTypeDuel {
		if battleResult == model.BattleResultWin {
			return 2, 1, 0, -1, 0 // Visitor wins
		} else if battleResult == model.BattleResultLose {
			return 1, 0, 2, 1, 0 // Owner wins (visitor loses)
		}
	} else if visitType == model.VisitTypeAttack {
		if battleResult == model.BattleResultWin {
			return 5, 3, 0, -2, 0 // Visitor wins (will set coins separately)
		} else if battleResult == model.BattleResultLose {
			return 0, 0, 3, 1, 0 // Owner wins (visitor loses)
		}
	}

	return 0, 0, 0, 0, 0
}
