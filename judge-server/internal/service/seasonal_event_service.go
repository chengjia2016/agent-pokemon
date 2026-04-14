package service

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// SeasonalEvent represents a seasonal event
type SeasonalEvent struct {
	ID                int             `json:"id"`
	EventID           string          `json:"event_id"`
	EventName         string          `json:"event_name"`
	EventNameZh       *string         `json:"event_name_zh,omitempty"`
	EventType         string          `json:"event_type"` // story, challenge, time-attack, score-attack
	Season            *string         `json:"season,omitempty"`
	Description       *string         `json:"description,omitempty"`
	DescriptionZh     *string         `json:"description_zh,omitempty"`
	StartDate         time.Time       `json:"start_date"`
	EndDate           time.Time       `json:"end_date"`
	FeaturedSyncPairs json.RawMessage `json:"featured_sync_pairs,omitempty"` // JSON array
	DifficultyLevels  json.RawMessage `json:"difficulty_levels,omitempty"`   // JSONB
	Rewards           json.RawMessage `json:"rewards,omitempty"`             // JSONB
	IsActive          bool            `json:"is_active"`
	IsLimited         bool            `json:"is_limited"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}

// EventProgress represents a player's progress in an event
type EventProgress struct {
	ID                 int             `json:"id"`
	ProgressID         string          `json:"progress_id"`
	EventID            string          `json:"event_id"`
	PlayerID           string          `json:"player_id"`
	StageCompleted     int             `json:"stage_completed"`
	ProgressPercentage float64         `json:"progress_percentage"`
	Score              int             `json:"score"`
	TimeSpent          int64           `json:"time_spent"`                // milliseconds
	RewardsClaimed     json.RawMessage `json:"rewards_claimed,omitempty"` // JSONB
	LastAttemptAt      *time.Time      `json:"last_attempt_at,omitempty"`
	CompletedAt        *time.Time      `json:"completed_at,omitempty"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

// SeasonalItem represents an item available during a seasonal event
type SeasonalItem struct {
	ID                  int       `json:"id"`
	ItemID              string    `json:"item_id"`
	EventID             *string   `json:"event_id,omitempty"`
	ItemName            string    `json:"item_name"`
	ItemNameZh          *string   `json:"item_name_zh,omitempty"`
	ItemType            string    `json:"item_type"` // clothing, accessory, material
	Rarity              string    `json:"rarity"`    // common, rare, legendary
	EffectDescription   *string   `json:"effect_description,omitempty"`
	EffectDescriptionZh *string   `json:"effect_description_zh,omitempty"`
	AvailableSeasons    *string   `json:"available_seasons,omitempty"` // JSON array
	CreatedAt           time.Time `json:"created_at"`
}

// PlayerSeasonalItem represents a player's seasonal item
type PlayerSeasonalItem struct {
	ID                 int       `json:"id"`
	PlayerID           string    `json:"player_id"`
	ItemID             string    `json:"item_id"`
	Quantity           int       `json:"quantity"`
	ObtainedAt         time.Time `json:"obtained_at"`
	EquippedToSyncPair *string   `json:"equipped_to_sync_pair,omitempty"`
}

// SeasonalEventService handles seasonal events for Masters EX
type SeasonalEventService struct {
	db *sql.DB
}

// NewSeasonalEventService creates a new seasonal event service
func NewSeasonalEventService(db *sql.DB) *SeasonalEventService {
	return &SeasonalEventService{
		db: db,
	}
}

// CreateSeasonalEvent creates a new seasonal event
func (s *SeasonalEventService) CreateSeasonalEvent(eventName, eventNameZh, eventType, season string,
	description, descriptionZh *string, startDate, endDate time.Time,
	featuredSyncPairs, difficultyLevels, rewards json.RawMessage, isLimited bool) (*SeasonalEvent, error) {

	eventID := fmt.Sprintf("event_%s_%d", eventType, time.Now().Unix())

	event := &SeasonalEvent{
		EventID:           eventID,
		EventName:         eventName,
		EventNameZh:       &eventNameZh,
		EventType:         eventType,
		Season:            &season,
		Description:       description,
		DescriptionZh:     descriptionZh,
		StartDate:         startDate,
		EndDate:           endDate,
		FeaturedSyncPairs: featuredSyncPairs,
		DifficultyLevels:  difficultyLevels,
		Rewards:           rewards,
		IsActive:          false,
		IsLimited:         isLimited,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	query := `
		INSERT INTO seasonal_events
		(event_id, event_name, event_name_zh, event_type, season, description, description_zh,
		 start_date, end_date, featured_sync_pairs, difficulty_levels, rewards, is_active, is_limited, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
		RETURNING id, created_at, updated_at`

	err := s.db.QueryRow(query, event.EventID, event.EventName, event.EventNameZh, event.EventType,
		event.Season, event.Description, event.DescriptionZh, event.StartDate, event.EndDate,
		event.FeaturedSyncPairs, event.DifficultyLevels, event.Rewards, event.IsActive, event.IsLimited,
		event.CreatedAt, event.UpdatedAt).Scan(&event.ID, &event.CreatedAt, &event.UpdatedAt)

	if err != nil {
		log.Printf("Error creating seasonal event: %v", err)
		return nil, err
	}

	return event, nil
}

// GetSeasonalEvent retrieves a seasonal event by ID
func (s *SeasonalEventService) GetSeasonalEvent(eventID string) (*SeasonalEvent, error) {
	event := &SeasonalEvent{}

	query := `
		SELECT id, event_id, event_name, event_name_zh, event_type, season, description, description_zh,
		       start_date, end_date, featured_sync_pairs, difficulty_levels, rewards, is_active, is_limited,
		       created_at, updated_at
		FROM seasonal_events WHERE event_id = $1`

	var eventNameZh, season, description, descriptionZh sql.NullString
	var featuredSyncPairs, difficultyLevels, rewards sql.NullString

	err := s.db.QueryRow(query, eventID).Scan(
		&event.ID, &event.EventID, &event.EventName, &eventNameZh, &event.EventType, &season,
		&description, &descriptionZh, &event.StartDate, &event.EndDate, &featuredSyncPairs,
		&difficultyLevels, &rewards, &event.IsActive, &event.IsLimited, &event.CreatedAt, &event.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("seasonal event not found: %s", eventID)
		}
		log.Printf("Error getting seasonal event: %v", err)
		return nil, err
	}

	if eventNameZh.Valid {
		event.EventNameZh = &eventNameZh.String
	}
	if season.Valid {
		event.Season = &season.String
	}
	if description.Valid {
		event.Description = &description.String
	}
	if descriptionZh.Valid {
		event.DescriptionZh = &descriptionZh.String
	}
	if featuredSyncPairs.Valid {
		event.FeaturedSyncPairs = json.RawMessage(featuredSyncPairs.String)
	}
	if difficultyLevels.Valid {
		event.DifficultyLevels = json.RawMessage(difficultyLevels.String)
	}
	if rewards.Valid {
		event.Rewards = json.RawMessage(rewards.String)
	}

	return event, nil
}

// ActivateSeasonalEvent activates a seasonal event
func (s *SeasonalEventService) ActivateSeasonalEvent(eventID string) error {
	query := `UPDATE seasonal_events SET is_active = true, updated_at = $1 WHERE event_id = $2`

	_, err := s.db.Exec(query, time.Now(), eventID)
	if err != nil {
		log.Printf("Error activating seasonal event: %v", err)
		return err
	}

	return nil
}

// DeactivateSeasonalEvent deactivates a seasonal event
func (s *SeasonalEventService) DeactivateSeasonalEvent(eventID string) error {
	query := `UPDATE seasonal_events SET is_active = false, updated_at = $1 WHERE event_id = $2`

	_, err := s.db.Exec(query, time.Now(), eventID)
	if err != nil {
		log.Printf("Error deactivating seasonal event: %v", err)
		return err
	}

	return nil
}

// GetActiveEvents returns all currently active seasonal events
func (s *SeasonalEventService) GetActiveEvents() ([]*SeasonalEvent, error) {
	query := `
		SELECT id, event_id, event_name, event_name_zh, event_type, season, description, description_zh,
		       start_date, end_date, featured_sync_pairs, difficulty_levels, rewards, is_active, is_limited,
		       created_at, updated_at
		FROM seasonal_events 
		WHERE is_active = true AND start_date <= NOW() AND end_date > NOW()
		ORDER BY end_date ASC`

	rows, err := s.db.Query(query)
	if err != nil {
		log.Printf("Error getting active events: %v", err)
		return nil, err
	}
	defer rows.Close()

	var events []*SeasonalEvent
	for rows.Next() {
		event := &SeasonalEvent{}
		var eventNameZh, season, description, descriptionZh sql.NullString
		var featuredSyncPairs, difficultyLevels, rewards sql.NullString

		err := rows.Scan(
			&event.ID, &event.EventID, &event.EventName, &eventNameZh, &event.EventType, &season,
			&description, &descriptionZh, &event.StartDate, &event.EndDate, &featuredSyncPairs,
			&difficultyLevels, &rewards, &event.IsActive, &event.IsLimited, &event.CreatedAt, &event.UpdatedAt)

		if err != nil {
			log.Printf("Error scanning active event row: %v", err)
			continue
		}

		if eventNameZh.Valid {
			event.EventNameZh = &eventNameZh.String
		}
		if season.Valid {
			event.Season = &season.String
		}
		if description.Valid {
			event.Description = &description.String
		}
		if descriptionZh.Valid {
			event.DescriptionZh = &descriptionZh.String
		}
		if featuredSyncPairs.Valid {
			event.FeaturedSyncPairs = json.RawMessage(featuredSyncPairs.String)
		}
		if difficultyLevels.Valid {
			event.DifficultyLevels = json.RawMessage(difficultyLevels.String)
		}
		if rewards.Valid {
			event.Rewards = json.RawMessage(rewards.String)
		}

		events = append(events, event)
	}

	return events, nil
}

// JoinEvent registers a player in an event
func (s *SeasonalEventService) JoinEvent(eventID, playerID string) (*EventProgress, error) {
	progressID := fmt.Sprintf("progress_%s_%s_%d", eventID, playerID, time.Now().Unix())

	progress := &EventProgress{
		ProgressID:         progressID,
		EventID:            eventID,
		PlayerID:           playerID,
		StageCompleted:     0,
		ProgressPercentage: 0.0,
		Score:              0,
		TimeSpent:          0,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	query := `
		INSERT INTO event_progress
		(progress_id, event_id, player_id, stage_completed, progress_percentage, score, time_spent, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at`

	err := s.db.QueryRow(query, progress.ProgressID, progress.EventID, progress.PlayerID,
		progress.StageCompleted, progress.ProgressPercentage, progress.Score, progress.TimeSpent,
		progress.CreatedAt, progress.UpdatedAt).Scan(&progress.ID, &progress.CreatedAt, &progress.UpdatedAt)

	if err != nil {
		log.Printf("Error joining event: %v", err)
		return nil, err
	}

	return progress, nil
}

// GetEventProgress retrieves a player's progress in an event
func (s *SeasonalEventService) GetEventProgress(eventID, playerID string) (*EventProgress, error) {
	progress := &EventProgress{}

	query := `
		SELECT id, progress_id, event_id, player_id, stage_completed, progress_percentage,
		       score, time_spent, rewards_claimed, last_attempt_at, completed_at, created_at, updated_at
		FROM event_progress
		WHERE event_id = $1 AND player_id = $2`

	var rewardsClaimed, lastAttemptAt, completedAt sql.NullString

	err := s.db.QueryRow(query, eventID, playerID).Scan(
		&progress.ID, &progress.ProgressID, &progress.EventID, &progress.PlayerID,
		&progress.StageCompleted, &progress.ProgressPercentage, &progress.Score, &progress.TimeSpent,
		&rewardsClaimed, &lastAttemptAt, &completedAt, &progress.CreatedAt, &progress.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("event progress not found")
		}
		log.Printf("Error getting event progress: %v", err)
		return nil, err
	}

	if rewardsClaimed.Valid {
		progress.RewardsClaimed = json.RawMessage(rewardsClaimed.String)
	}
	if lastAttemptAt.Valid {
		lastTime, _ := time.Parse(time.RFC3339, lastAttemptAt.String)
		progress.LastAttemptAt = &lastTime
	}
	if completedAt.Valid {
		complTime, _ := time.Parse(time.RFC3339, completedAt.String)
		progress.CompletedAt = &complTime
	}

	return progress, nil
}

// UpdateEventProgress updates a player's progress in an event
func (s *SeasonalEventService) UpdateEventProgress(eventID, playerID string,
	stageCompleted int, progressPercentage float64, score int64, timeSpent int64) (*EventProgress, error) {

	progress, err := s.GetEventProgress(eventID, playerID)
	if err != nil {
		return nil, err
	}

	progress.StageCompleted = stageCompleted
	progress.ProgressPercentage = progressPercentage
	progress.Score = int(score)
	progress.TimeSpent = timeSpent

	now := time.Now()
	progress.LastAttemptAt = &now

	if progressPercentage >= 100.0 && progress.CompletedAt == nil {
		progress.CompletedAt = &now
	}

	query := `
		UPDATE event_progress
		SET stage_completed = $1, progress_percentage = $2, score = $3, time_spent = $4,
		    last_attempt_at = $5, completed_at = $6, updated_at = $7
		WHERE event_id = $8 AND player_id = $9`

	_, err = s.db.Exec(query, progress.StageCompleted, progress.ProgressPercentage,
		progress.Score, progress.TimeSpent, progress.LastAttemptAt, progress.CompletedAt,
		time.Now(), eventID, playerID)

	if err != nil {
		log.Printf("Error updating event progress: %v", err)
		return nil, err
	}

	return progress, nil
}

// ClaimEventRewards claims rewards for event progress
func (s *SeasonalEventService) ClaimEventRewards(eventID, playerID string, rewards json.RawMessage) (*EventProgress, error) {
	progress, err := s.GetEventProgress(eventID, playerID)
	if err != nil {
		return nil, err
	}

	progress.RewardsClaimed = rewards

	query := `
		UPDATE event_progress
		SET rewards_claimed = $1, updated_at = $2
		WHERE event_id = $3 AND player_id = $4`

	_, err = s.db.Exec(query, rewards, time.Now(), eventID, playerID)
	if err != nil {
		log.Printf("Error claiming event rewards: %v", err)
		return nil, err
	}

	return progress, nil
}

// CreateSeasonalItem creates a new seasonal item
func (s *SeasonalEventService) CreateSeasonalItem(itemName, itemNameZh, itemType, rarity string,
	eventID *string, effectDescription, effectDescriptionZh, availableSeasons *string) (*SeasonalItem, error) {

	itemID := fmt.Sprintf("item_%s_%d", itemType, time.Now().Unix())

	item := &SeasonalItem{
		ItemID:              itemID,
		EventID:             eventID,
		ItemName:            itemName,
		ItemNameZh:          &itemNameZh,
		ItemType:            itemType,
		Rarity:              rarity,
		EffectDescription:   effectDescription,
		EffectDescriptionZh: effectDescriptionZh,
		AvailableSeasons:    availableSeasons,
		CreatedAt:           time.Now(),
	}

	query := `
		INSERT INTO seasonal_items
		(item_id, event_id, item_name, item_name_zh, item_type, rarity, effect_description,
		 effect_description_zh, available_seasons, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, created_at`

	err := s.db.QueryRow(query, item.ItemID, item.EventID, item.ItemName, item.ItemNameZh,
		item.ItemType, item.Rarity, item.EffectDescription, item.EffectDescriptionZh,
		item.AvailableSeasons, item.CreatedAt).Scan(&item.ID, &item.CreatedAt)

	if err != nil {
		log.Printf("Error creating seasonal item: %v", err)
		return nil, err
	}

	return item, nil
}

// GetEventItems retrieves all items for a seasonal event
func (s *SeasonalEventService) GetEventItems(eventID string) ([]*SeasonalItem, error) {
	query := `
		SELECT id, item_id, event_id, item_name, item_name_zh, item_type, rarity,
		       effect_description, effect_description_zh, available_seasons, created_at
		FROM seasonal_items
		WHERE event_id = $1
		ORDER BY rarity DESC, item_name ASC`

	rows, err := s.db.Query(query, eventID)
	if err != nil {
		log.Printf("Error getting event items: %v", err)
		return nil, err
	}
	defer rows.Close()

	var items []*SeasonalItem
	for rows.Next() {
		item := &SeasonalItem{}
		var itemNameZh, effectDescription, effectDescriptionZh, availableSeasons sql.NullString

		err := rows.Scan(
			&item.ID, &item.ItemID, &item.EventID, &item.ItemName, &itemNameZh,
			&item.ItemType, &item.Rarity, &effectDescription, &effectDescriptionZh,
			&availableSeasons, &item.CreatedAt)

		if err != nil {
			log.Printf("Error scanning item row: %v", err)
			continue
		}

		if itemNameZh.Valid {
			item.ItemNameZh = &itemNameZh.String
		}
		if effectDescription.Valid {
			item.EffectDescription = &effectDescription.String
		}
		if effectDescriptionZh.Valid {
			item.EffectDescriptionZh = &effectDescriptionZh.String
		}
		if availableSeasons.Valid {
			item.AvailableSeasons = &availableSeasons.String
		}

		items = append(items, item)
	}

	return items, nil
}

// Implement Value interface for JSON serialization
func (se *SeasonalEvent) Value() (driver.Value, error) {
	return json.Marshal(se)
}

func (ep *EventProgress) Value() (driver.Value, error) {
	return json.Marshal(ep)
}

func (si *SeasonalItem) Value() (driver.Value, error) {
	return json.Marshal(si)
}
