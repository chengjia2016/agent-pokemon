package service

import (
	"errors"
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/model"
	"math"
	"math/rand"
	"time"
)

// RewardCalculator 奖励计算器
type RewardCalculator struct {
	DB *db.Database
}

func NewRewardCalculator(database *db.Database) *RewardCalculator {
	return &RewardCalculator{DB: database}
}

// ==================== 奖励计算 ====================

// CalculateBattleReward 计算战斗奖励
func (rc *RewardCalculator) CalculateBattleReward(loserBalance float64) (float64, float64, error) {
	// 生成随机百分比 1%-5%
	percentage := rc.generateRandomPercentage(1, 5)

	// 计算转账金额
	transferAmount := loserBalance * (percentage / 100.0)

	// 四舍五入到2位小数
	transferAmount = math.Round(transferAmount*100) / 100

	return transferAmount, percentage, nil
}

// generateRandomPercentage 生成随机百分比
func (rc *RewardCalculator) generateRandomPercentage(min, max int) float64 {
	// 生成 min 到 max 之间的随机数（包括小数）
	randomValue := float64(min) + rand.Float64()*float64(max-min)
	// 保留2位小数
	return math.Round(randomValue*100) / 100
}

// ==================== 余额处理 ====================

// HandleInsufficientBalance 处理余额不足情况
func (rc *RewardCalculator) HandleInsufficientBalance(loserBalance, calculatedReward float64) float64 {
	// 如果计算的转账金额大于余额，则转账所有余额
	if calculatedReward > loserBalance {
		return loserBalance
	}
	return calculatedReward
}

// ==================== 精灵币转账 ====================

// TransferCoins 转账精灵币
func (rc *RewardCalculator) TransferCoins(winnerGitHubID, loserGitHubID int, amount float64, battleID string) (*model.BattleReward, error) {
	if amount < 0 {
		return nil, errors.New("transfer amount cannot be negative")
	}

	// 如果金额为0，不进行转账
	if amount == 0 {
		return nil, nil
	}

	// 获取失败者的当前余额
	loser, err := rc.getUser(loserGitHubID)
	if err != nil {
		return nil, fmt.Errorf("failed to get loser: %w", err)
	}

	// 检查余额是否足够
	if loser.Balance < amount {
		// 转账所有余额（根据需求）
		amount = loser.Balance
	}

	// 更新赢家的余额
	winner, err := rc.getUser(winnerGitHubID)
	if err != nil {
		return nil, fmt.Errorf("failed to get winner: %w", err)
	}

	newWinnerBalance := winner.Balance + amount
	newLoserBalance := loser.Balance - amount

	// 更新数据库
	err = rc.updateUserBalance(winnerGitHubID, newWinnerBalance)
	if err != nil {
		return nil, err
	}

	err = rc.updateUserBalance(loserGitHubID, newLoserBalance)
	if err != nil {
		return nil, err
	}

	// 计算百分比
	percentage := 0.0
	if loser.Balance > 0 {
		percentage = (amount / loser.Balance) * 100.0
	}

	// 记录交易
	reward := &model.BattleReward{
		BattleID:         battleID,
		WinnerGitHubID:   winnerGitHubID,
		LoserGitHubID:    loserGitHubID,
		CoinsTransferred: amount,
		Percentage:       percentage,
		TransactionID:    fmt.Sprintf("txn_%d_%d_%d", winnerGitHubID, loserGitHubID, time.Now().Unix()),
		CompletedAt:      time.Now(),
	}

	// 保存到数据库
	err = rc.saveReward(reward)
	if err != nil {
		return nil, err
	}

	return reward, nil
}

// ==================== 用户余额查询 ====================

// GetUserBalance 获取用户余额
func (rc *RewardCalculator) GetUserBalance(githubID int) (float64, error) {
	user, err := rc.getUser(githubID)
	if err != nil {
		return 0, err
	}
	return user.Balance, nil
}

// getUser 获取用户（内部方法）
func (rc *RewardCalculator) getUser(githubID int) (*model.UserAccount, error) {
	// TODO: 实现从数据库查询用户
	// 这需要在数据库层添加查询方法
	return &model.UserAccount{}, errors.New("method not implemented")
}

// updateUserBalance 更新用户余额
func (rc *RewardCalculator) updateUserBalance(githubID int, newBalance float64) error {
	// TODO: 实现到数据库的更新
	// 这需要在数据库层添加更新方法
	return errors.New("method not implemented")
}

// saveReward 保存奖励记录
func (rc *RewardCalculator) saveReward(reward *model.BattleReward) error {
	// TODO: 实现到数据库的保存
	return errors.New("method not implemented")
}

// ==================== 防守基地奖励 ====================

// CalculateDefenseReward 计算防守基地奖励
func (rc *RewardCalculator) CalculateDefenseReward(baseID string, attackerGitHubID int, result string) (*model.BaseDefenseRecord, error) {
	// 获取攻击者信息
	attackerBalance, err := rc.GetUserBalance(attackerGitHubID)
	if err != nil {
		return nil, err
	}

	// 计算奖励
	coinsGained := 0.0
	if result == model.DefenseResultWin {
		// 防守成功，攻击者失去金币
		reward, _, err := rc.CalculateBattleReward(attackerBalance)
		if err != nil {
			return nil, err
		}
		coinsGained = reward
	}

	record := &model.BaseDefenseRecord{
		BaseID:           baseID,
		AttackerGitHubID: attackerGitHubID,
		Result:           result,
		CoinsGained:      coinsGained,
		Timestamp:        time.Now(),
	}

	return record, nil
}

// ==================== 批量奖励处理 ====================

// ProcessBatchRewards 批量处理奖励
func (rc *RewardCalculator) ProcessBatchRewards(rewards []*model.BattleReward) error {
	for _, reward := range rewards {
		if reward == nil {
			continue
		}

		err := rc.saveReward(reward)
		if err != nil {
			return fmt.Errorf("failed to save reward: %w", err)
		}
	}
	return nil
}

// ==================== 辅助方法 ====================

// IsValidRewardAmount 验证奖励金额
func (rc *RewardCalculator) IsValidRewardAmount(amount float64) bool {
	return amount >= 0 && amount <= 10000000 // 最多1000万
}

// FormatCoinAmount 格式化金币金额
func (rc *RewardCalculator) FormatCoinAmount(amount float64) string {
	if amount >= 1000000 {
		return fmt.Sprintf("%.2fM", amount/1000000)
	} else if amount >= 1000 {
		return fmt.Sprintf("%.2fK", amount/1000)
	}
	return fmt.Sprintf("%.2f", amount)
}

// UserAccount 临时数据模型（应该在model中定义）
type UserAccount struct {
	ID        int
	GitHubID  int
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
