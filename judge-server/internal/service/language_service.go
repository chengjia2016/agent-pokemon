package service

import (
	"database/sql"
	"fmt"
	"judge-server/internal/db"
)

type LanguageService struct {
	db *db.Database
}

type Language struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type UIString struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

type Translation struct {
	ID             int    `json:"id"`
	UIStringID     int    `json:"ui_string_id"`
	LanguageID     int    `json:"language_id"`
	TranslatedText string `json:"translated_text"`
}

type UserLanguagePreference struct {
	ID         int    `json:"id"`
	UserID     string `json:"user_id"`
	LanguageID int    `json:"language_id"`
}

type NPCDialogue struct {
	ID              int    `json:"id"`
	NPCID           int    `json:"npc_id"`
	DialogueText    string `json:"dialogue_text"`
	DialogueTextZH  string `json:"dialogue_text_zh"`
	DialogueContext string `json:"dialogue_context"`
	DialogueType    string `json:"dialogue_type"`
}

func NewLanguageService(database *db.Database) *LanguageService {
	return &LanguageService{
		db: database,
	}
}

// GetAllLanguages retrieves all active languages
func (ls *LanguageService) GetAllLanguages() ([]Language, error) {
	var languages []Language
	query := `SELECT id, code, name, is_active FROM languages WHERE is_active = true ORDER BY code`

	rows, err := ls.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query languages: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var lang Language
		if err := rows.Scan(&lang.ID, &lang.Code, &lang.Name, &lang.IsActive); err != nil {
			return nil, fmt.Errorf("failed to scan language: %v", err)
		}
		languages = append(languages, lang)
	}

	return languages, rows.Err()
}

// GetLanguageByCode retrieves a language by its code (e.g., 'en', 'zh')
func (ls *LanguageService) GetLanguageByCode(code string) (*Language, error) {
	var lang Language
	query := `SELECT id, code, name, is_active FROM languages WHERE code = $1`

	err := ls.db.QueryRow(query, code).Scan(&lang.ID, &lang.Code, &lang.Name, &lang.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("language not found: %s", code)
		}
		return nil, fmt.Errorf("failed to query language: %v", err)
	}

	return &lang, nil
}

// GetUIStringsForLanguage retrieves all UI strings with translations for a specific language
func (ls *LanguageService) GetUIStringsForLanguage(languageCode string) (map[string]string, error) {
	translations := make(map[string]string)

	query := `
		SELECT us.key, ut.translated_text
		FROM ui_translations ut
		JOIN ui_strings us ON ut.ui_string_id = us.id
		JOIN languages l ON ut.language_id = l.id
		WHERE l.code = $1
	`

	rows, err := ls.db.Query(query, languageCode)
	if err != nil {
		return nil, fmt.Errorf("failed to query UI strings: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var key, text string
		if err := rows.Scan(&key, &text); err != nil {
			return nil, fmt.Errorf("failed to scan UI string: %v", err)
		}
		translations[key] = text
	}

	return translations, rows.Err()
}

// GetTranslation retrieves a single translation for a UI string in a specific language
func (ls *LanguageService) GetTranslation(uiStringKey string, languageCode string) (string, error) {
	query := `
		SELECT ut.translated_text
		FROM ui_translations ut
		JOIN ui_strings us ON ut.ui_string_id = us.id
		JOIN languages l ON ut.language_id = l.id
		WHERE us.key = $1 AND l.code = $2
	`

	var text string
	err := ls.db.QueryRow(query, uiStringKey, languageCode).Scan(&text)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("translation not found for key: %s, language: %s", uiStringKey, languageCode)
		}
		return "", fmt.Errorf("failed to query translation: %v", err)
	}

	return text, nil
}

// SetUserLanguagePreference sets or updates a user's language preference
func (ls *LanguageService) SetUserLanguagePreference(userID string, languageCode string) error {
	// First, get the language ID
	lang, err := ls.GetLanguageByCode(languageCode)
	if err != nil {
		return err
	}

	// Check if preference already exists
	query := `SELECT id FROM user_language_preferences WHERE user_id = $1`
	var existingID int
	err = ls.db.QueryRow(query, userID).Scan(&existingID)

	if err == sql.ErrNoRows {
		// Insert new preference
		insertQuery := `
			INSERT INTO user_language_preferences (user_id, language_id, created_at, updated_at)
			VALUES ($1, $2, NOW(), NOW())
		`
		_, err = ls.db.Exec(insertQuery, userID, lang.ID)
		if err != nil {
			return fmt.Errorf("failed to insert language preference: %v", err)
		}
	} else if err == nil {
		// Update existing preference
		updateQuery := `UPDATE user_language_preferences SET language_id = $1, updated_at = NOW() WHERE user_id = $2`
		_, err = ls.db.Exec(updateQuery, lang.ID, userID)
		if err != nil {
			return fmt.Errorf("failed to update language preference: %v", err)
		}
	} else {
		return fmt.Errorf("failed to check language preference: %v", err)
	}

	return nil
}

// GetUserLanguagePreference retrieves a user's language preference
func (ls *LanguageService) GetUserLanguagePreference(userID string) (string, error) {
	query := `
		SELECT l.code
		FROM user_language_preferences ulp
		JOIN languages l ON ulp.language_id = l.id
		WHERE ulp.user_id = $1
	`

	var code string
	err := ls.db.QueryRow(query, userID).Scan(&code)
	if err != nil {
		if err == sql.ErrNoRows {
			// Default to English if no preference set
			return "en", nil
		}
		return "", fmt.Errorf("failed to query user language preference: %v", err)
	}

	return code, nil
}

// GetNPCDialogue retrieves an NPC dialogue in a specific language
func (ls *LanguageService) GetNPCDialogue(npcID int, languageCode string) (*NPCDialogue, error) {
	query := `
		SELECT nd.id, nd.npc_id, nd.dialogue_text, nd.dialogue_text_zh, nd.dialogue_context, nd.dialogue_type
		FROM npc_dialogues nd
		WHERE nd.npc_id = $1
		LIMIT 1
	`

	var dialogue NPCDialogue
	err := ls.db.QueryRow(query, npcID).Scan(
		&dialogue.ID,
		&dialogue.NPCID,
		&dialogue.DialogueText,
		&dialogue.DialogueTextZH,
		&dialogue.DialogueContext,
		&dialogue.DialogueType,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("dialogue not found for NPC: %d", npcID)
		}
		return nil, fmt.Errorf("failed to query NPC dialogue: %v", err)
	}

	// Select the appropriate text based on language
	if languageCode == "zh" {
		if dialogue.DialogueTextZH == "" {
			dialogue.DialogueTextZH = dialogue.DialogueText // Fallback to English
		}
	}

	return &dialogue, nil
}

// GetAllNPCDialogues retrieves all NPC dialogues in a specific language
func (ls *LanguageService) GetAllNPCDialogues(languageCode string) ([]NPCDialogue, error) {
	var dialogues []NPCDialogue
	query := `
		SELECT id, npc_id, dialogue_text, dialogue_text_zh, dialogue_context, dialogue_type
		FROM npc_dialogues
		ORDER BY npc_id
	`

	rows, err := ls.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query NPC dialogues: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dialogue NPCDialogue
		if err := rows.Scan(
			&dialogue.ID,
			&dialogue.NPCID,
			&dialogue.DialogueText,
			&dialogue.DialogueTextZH,
			&dialogue.DialogueContext,
			&dialogue.DialogueType,
		); err != nil {
			return nil, fmt.Errorf("failed to scan NPC dialogue: %v", err)
		}

		// Select the appropriate text based on language
		if languageCode == "zh" {
			if dialogue.DialogueTextZH == "" {
				dialogue.DialogueTextZH = dialogue.DialogueText // Fallback to English
			}
		}

		dialogues = append(dialogues, dialogue)
	}

	return dialogues, rows.Err()
}

// GetQuestDescription retrieves quest description in a specific language
func (ls *LanguageService) GetQuestDescription(questID int, languageCode string) (map[string]interface{}, error) {
	query := `
		SELECT id, quest_name, description, description_zh, reward_item, reward_item_zh, required_steps
		FROM quests
		WHERE id = $1
	`

	var questName, description, descriptionZH, rewardItem, rewardItemZH string
	var requiredSteps int
	err := ls.db.QueryRow(query, questID).Scan(
		&questID,
		&questName,
		&description,
		&descriptionZH,
		&rewardItem,
		&rewardItemZH,
		&requiredSteps,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("quest not found: %d", questID)
		}
		return nil, fmt.Errorf("failed to query quest: %v", err)
	}

	// Prepare result based on language
	result := make(map[string]interface{})
	result["id"] = questID
	result["quest_name"] = questName
	result["required_steps"] = requiredSteps

	if languageCode == "zh" {
		result["description"] = descriptionZH
		if descriptionZH == "" {
			result["description"] = description // Fallback to English
		}
		result["reward_item"] = rewardItemZH
		if rewardItemZH == "" {
			result["reward_item"] = rewardItem // Fallback to English
		}
	} else {
		result["description"] = description
		result["reward_item"] = rewardItem
	}

	return result, nil
}

// InitializeLanguageTables creates the necessary database tables for multilingual support
func (ls *LanguageService) InitializeLanguageTables() error {
	// Create languages table
	createLanguagesTable := `
		CREATE TABLE IF NOT EXISTS languages (
			id SERIAL PRIMARY KEY,
			code VARCHAR(10) UNIQUE NOT NULL,
			name VARCHAR(50) NOT NULL,
			is_active BOOLEAN DEFAULT true,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	// Create ui_strings table
	createUIStringsTable := `
		CREATE TABLE IF NOT EXISTS ui_strings (
			id SERIAL PRIMARY KEY,
			key VARCHAR(100) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	// Create ui_translations table
	createUITranslationsTable := `
		CREATE TABLE IF NOT EXISTS ui_translations (
			id SERIAL PRIMARY KEY,
			ui_string_id INTEGER NOT NULL REFERENCES ui_strings(id),
			language_id INTEGER NOT NULL REFERENCES languages(id),
			translated_text TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(ui_string_id, language_id)
		);
	`

	// Create user_language_preferences table
	createUserLanguagePreferencesTable := `
		CREATE TABLE IF NOT EXISTS user_language_preferences (
			id SERIAL PRIMARY KEY,
			user_id VARCHAR(100) NOT NULL,
			language_id INTEGER NOT NULL REFERENCES languages(id),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(user_id)
		);
	`

	// Execute all table creations
	if _, err := ls.db.Exec(createLanguagesTable); err != nil {
		return fmt.Errorf("failed to create languages table: %v", err)
	}

	if _, err := ls.db.Exec(createUIStringsTable); err != nil {
		return fmt.Errorf("failed to create ui_strings table: %v", err)
	}

	if _, err := ls.db.Exec(createUITranslationsTable); err != nil {
		return fmt.Errorf("failed to create ui_translations table: %v", err)
	}

	if _, err := ls.db.Exec(createUserLanguagePreferencesTable); err != nil {
		return fmt.Errorf("failed to create user_language_preferences table: %v", err)
	}

	return nil
}
