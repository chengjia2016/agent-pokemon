package service

import (
	"judge-server/internal/db"
	"judge-server/internal/github"
	"judge-server/internal/model"
	"log"
	"sync"
	"time"
)

type BatchSyncService struct {
	db        *db.Database
	github    *github.SyncService
	interval  time.Duration
	batchSize int
	queue     []model.Battle
	mu        sync.Mutex
	stopCh    chan struct{}
}

func NewBatchSyncService(database *db.Database, syncService *github.SyncService, intervalSeconds int) *BatchSyncService {
	return &BatchSyncService{
		db:        database,
		github:    syncService,
		interval:  time.Duration(intervalSeconds) * time.Second,
		batchSize: 10,
		queue:     []model.Battle{},
		mu:        sync.Mutex{},
		stopCh:    make(chan struct{}),
	}
}

func (s *BatchSyncService) Start() {
	if s.github == nil {
		log.Println("BatchSyncService: GitHub sync disabled")
		return
	}

	log.Printf("BatchSyncService: Starting with %v interval", s.interval)

	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.flush()
		case <-s.stopCh:
			s.flush()
			log.Println("BatchSyncService: Stopped")
			return
		}
	}
}

func (s *BatchSyncService) Stop() {
	close(s.stopCh)
}

func (s *BatchSyncService) AddBattle(battle model.Battle) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.queue) >= s.batchSize {
		s.flush()
	}
	s.queue = append(s.queue, battle)
}

func (s *BatchSyncService) flush() {
	s.mu.Lock()
	pending := s.queue
	s.queue = []model.Battle{}
	s.mu.Unlock()

	if len(pending) == 0 {
		return
	}

	log.Printf("BatchSyncService: Syncing %d battles to GitHub", len(pending))

	for _, battle := range pending {
		if err := s.github.SyncBattleReportToIssue(&battle); err != nil {
			log.Printf("BatchSyncService: Failed to sync battle %s: %v", battle.ID, err)
		}
	}

	log.Printf("BatchSyncService: Synced %d battles", len(pending))
}
