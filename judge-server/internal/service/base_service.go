package service

import (
	"database/sql"
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/model"
	"time"
)

// BaseService 基地服务
type BaseService struct {
	db *db.Database
}

// NewBaseService 创建基地服务
func NewBaseService(database *db.Database) *BaseService {
	return &BaseService{
		db: database,
	}
}

// CreateBase 创建用户的基地
func (s *BaseService) CreateBase(req *model.CreateBaseRequest) (*model.UserBase, error) {
	// Note: githubID should be extracted from authentication context by caller
	// For now, we'll return an error since githubID is required but not provided
	return nil, fmt.Errorf("githubID must be provided in authentication context")
}

// CreateBaseForUser 创建用户的基地(带githubID)
func (s *BaseService) CreateBaseForUser(githubID int, req *model.CreateBaseRequest) (*model.UserBase, error) {
	// 生成基地ID
	baseID := fmt.Sprintf("base_%d_%d", githubID, time.Now().Unix())

	query := `
		INSERT INTO user_bases (
			github_id, base_id, repository_url, defense_team_id, 
			level, prestige, defense_wins, defense_losses, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		RETURNING id, github_id, base_id, repository_url, defense_team_id, 
				  level, prestige, defense_wins, defense_losses, created_at, updated_at
	`

	base := &model.UserBase{}
	err := s.db.QueryRow(query,
		githubID,
		baseID,
		req.RepositoryURL,
		nil, // defense_team_id
		1,   // level
		0,   // prestige
		0,   // defense_wins
		0,   // defense_losses
	).Scan(
		&base.ID,
		&base.GitHubID,
		&base.BaseID,
		&base.RepositoryURL,
		&base.DefenseTeamID,
		&base.Level,
		&base.Prestige,
		&base.DefenseWins,
		&base.DefenseLosses,
		&base.CreatedAt,
		&base.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create base: %w", err)
	}

	return base, nil
}

// GetUserBase 获取用户的基地
func (s *BaseService) GetUserBase(githubID int) (*model.UserBase, error) {
	query := `
		SELECT id, github_id, base_id, repository_url, defense_team_id, 
			   level, prestige, defense_wins, defense_losses, created_at, updated_at
		FROM user_bases
		WHERE github_id = $1
		LIMIT 1
	`

	base := &model.UserBase{}
	err := s.db.QueryRow(query, githubID).Scan(
		&base.ID,
		&base.GitHubID,
		&base.BaseID,
		&base.RepositoryURL,
		&base.DefenseTeamID,
		&base.Level,
		&base.Prestige,
		&base.DefenseWins,
		&base.DefenseLosses,
		&base.CreatedAt,
		&base.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("base not found")
		}
		return nil, fmt.Errorf("failed to get base: %w", err)
	}

	return base, nil
}

// UpdateBase 更新基地信息
func (s *BaseService) UpdateBase(baseID string, updates map[string]interface{}) (*model.UserBase, error) {
	query := `
		UPDATE user_bases
		SET defense_team_id = COALESCE($1, defense_team_id),
			updated_at = NOW()
		WHERE base_id = $2
		RETURNING id, github_id, base_id, repository_url, defense_team_id, 
				  level, prestige, defense_wins, defense_losses, created_at, updated_at
	`

	defenseTeamID := updates["defense_team_id"]

	base := &model.UserBase{}
	err := s.db.QueryRow(query, defenseTeamID, baseID).Scan(
		&base.ID,
		&base.GitHubID,
		&base.BaseID,
		&base.RepositoryURL,
		&base.DefenseTeamID,
		&base.Level,
		&base.Prestige,
		&base.DefenseWins,
		&base.DefenseLosses,
		&base.CreatedAt,
		&base.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update base: %w", err)
	}

	return base, nil
}

// AddDefenseWin 增加防守胜场
func (s *BaseService) AddDefenseWin(baseID string, coinsGained float64) error {
	query := `
		UPDATE user_bases
		SET defense_wins = defense_wins + 1,
			prestige = prestige + 10,
			updated_at = NOW()
		WHERE base_id = $1
	`

	_, err := s.db.Exec(query, baseID)
	if err != nil {
		return fmt.Errorf("failed to add defense win: %w", err)
	}

	return nil
}

// AddDefenseLoss 增加防守负场
func (s *BaseService) AddDefenseLoss(baseID string) error {
	query := `
		UPDATE user_bases
		SET defense_losses = defense_losses + 1,
			prestige = MAX(0, prestige - 5),
			updated_at = NOW()
		WHERE base_id = $1
	`

	_, err := s.db.Exec(query, baseID)
	if err != nil {
		return fmt.Errorf("failed to add defense loss: %w", err)
	}

	return nil
}

// DeleteBase 删除基地
func (s *BaseService) DeleteBase(baseID string) error {
	query := `DELETE FROM user_bases WHERE base_id = $1`

	_, err := s.db.Exec(query, baseID)
	if err != nil {
		return fmt.Errorf("failed to delete base: %w", err)
	}

	return nil
}

// ListUserBases 列出用户的所有基地（通常每用户一个）
func (s *BaseService) ListUserBases(githubID int) ([]*model.UserBase, error) {
	query := `
		SELECT id, github_id, base_id, repository_url, defense_team_id, 
			   level, prestige, defense_wins, defense_losses, created_at, updated_at
		FROM user_bases
		WHERE github_id = $1
		ORDER BY created_at DESC
	`

	rows, err := s.db.Query(query, githubID)
	if err != nil {
		return nil, fmt.Errorf("failed to list bases: %w", err)
	}
	defer rows.Close()

	var bases []*model.UserBase
	for rows.Next() {
		base := &model.UserBase{}
		err := rows.Scan(
			&base.ID,
			&base.GitHubID,
			&base.BaseID,
			&base.RepositoryURL,
			&base.DefenseTeamID,
			&base.Level,
			&base.Prestige,
			&base.DefenseWins,
			&base.DefenseLosses,
			&base.CreatedAt,
			&base.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan base: %w", err)
		}
		bases = append(bases, base)
	}

	return bases, nil
}
