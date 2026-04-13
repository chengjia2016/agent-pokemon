package github

import (
	"context"
	"fmt"
	"judge-server/internal/model"
	"time"

	"github.com/google/go-github/v58/github"
	"golang.org/x/oauth2"
)

type SyncService struct {
	client *github.Client
	owner  string
	repo   string
}

type Config struct {
	Token string
	Owner string
	Repo  string
}

func NewSyncService(cfg Config) *SyncService {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.Token})
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)

	return &SyncService{
		client: client,
		owner:  cfg.Owner,
		repo:   cfg.Repo,
	}
}

func (s *SyncService) SyncLeaderboardToIssue(date string, entries []model.LeaderboardEntry) error {
	ctx := context.Background()

	title := fmt.Sprintf("🏆 Daily Leaderboard - %s", date)
	body := s.formatLeaderboardBody(date, entries)

	issues, _, err := s.client.Issues.ListByRepo(ctx, s.owner, s.repo, &github.IssueListByRepoOptions{
		State: "open",
	})
	if err != nil {
		return err
	}

	for _, issue := range issues {
		if issue.GetTitle() == title {
			return nil
		}
	}

	_, _, err = s.client.Issues.Create(ctx, s.owner, s.repo, &github.IssueRequest{
		Title: github.String(title),
		Body:  github.String(body),
	})
	return err
}

func (s *SyncService) SyncBattleReportToIssue(battle *model.Battle) error {
	ctx := context.Background()

	title := fmt.Sprintf("⚔️ Battle: %s vs %s", battle.AttackerName, battle.DefenderName)
	body := s.formatBattleBody(battle)

	_, _, err := s.client.Issues.Create(ctx, s.owner, s.repo, &github.IssueRequest{
		Title: github.String(title),
		Body:  github.String(body),
	})
	return err
}

func (s *SyncService) TriggerGitHubWorkflow(workflowID string) error {
	ctx := context.Background()

	_, err := s.client.Actions.CreateWorkflowDispatchEventByFileName(ctx, s.owner, s.repo, workflowID, github.CreateWorkflowDispatchEventRequest{
		Ref: "main",
	})
	return err
}

func (s *SyncService) formatLeaderboardBody(date string, entries []model.LeaderboardEntry) string {
	body := fmt.Sprintf("# 🏆 Daily Leaderboard - %s\n\n", date)
	body += "| Rank | Player | Pet Name | Level | Wins | Losses | Rating |\n"
	body += "|------|--------|----------|-------|------|--------|--------|\n"

	for i, e := range entries {
		if i >= 10 {
			break
		}
		body += fmt.Sprintf("| %d | %s | %s | %d | %d | %d | %d |\n",
			e.Rank, e.Player, e.PetName, e.Level, e.Wins, e.Losses, e.Rating)
	}

	body += fmt.Sprintf("\n*Generated at: %s*\n", time.Now().Format("2006-01-02 15:04:05"))
	return body
}

func (s *SyncService) formatBattleBody(battle *model.Battle) string {
	body := fmt.Sprintf("# ⚔️ Battle Report\n\n")
	body += fmt.Sprintf("**Attacker**: %s (%s)\n", battle.AttackerName, battle.AttackerID)
	body += fmt.Sprintf("**Defender**: %s (%s)\n", battle.DefenderName, battle.DefenderID)
	body += fmt.Sprintf("**Winner**: %s\n", battle.Winner)
	body += fmt.Sprintf("**Turns**: %d\n\n", battle.Turns)

	body += "## Battle Log\n\n```\n"
	for _, log := range battle.BattleLog {
		body += log + "\n"
	}
	body += "```\n"

	body += fmt.Sprintf("\n*Battle Time: %s*\n", battle.StartTime.Format("2006-01-02 15:04:05"))
	return body
}
