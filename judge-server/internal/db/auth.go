package db

import (
	"database/sql"
	"errors"
	"fmt"
	"judge-server/internal/model"
)

// InitAuthSchema creates authentication-related tables
func (d *Database) InitAuthSchema() error {
	schemas := []string{
		// Player profiles table
		`CREATE TABLE IF NOT EXISTS player_profiles (
			id SERIAL PRIMARY KEY,
			player_id VARCHAR(255) UNIQUE NOT NULL,
			github_id INTEGER,
			github_login VARCHAR(255),
			email VARCHAR(255),
			avatar_url VARCHAR(255),
			is_test_account BOOLEAN DEFAULT false,
			server_token VARCHAR(255) UNIQUE,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW(),
			last_login_at TIMESTAMP,
			token_expire_at TIMESTAMP
		)`,

		// Client sessions table
		`CREATE TABLE IF NOT EXISTS client_sessions (
			id SERIAL PRIMARY KEY,
			player_id VARCHAR(255) NOT NULL REFERENCES player_profiles(player_id),
			client_name VARCHAR(100) NOT NULL,
			access_token VARCHAR(255) UNIQUE NOT NULL,
			client_type VARCHAR(50),
			last_active_at TIMESTAMP DEFAULT NOW(),
			created_at TIMESTAMP DEFAULT NOW(),
			expires_at TIMESTAMP
		)`,

		// Test accounts table
		`CREATE TABLE IF NOT EXISTS test_accounts (
			id SERIAL PRIMARY KEY,
			username VARCHAR(100) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			player_id VARCHAR(255) UNIQUE NOT NULL REFERENCES player_profiles(player_id),
			status VARCHAR(20) DEFAULT 'active',
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,

		// Player tokens table (for GitHub binding)
		`CREATE TABLE IF NOT EXISTS player_tokens (
			id SERIAL PRIMARY KEY,
			player_id VARCHAR(255),
			github_id INTEGER,
			github_login VARCHAR(255),
			token VARCHAR(255) UNIQUE NOT NULL,
			is_used BOOLEAN DEFAULT false,
			used_at TIMESTAMP,
			expires_at TIMESTAMP,
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		// Create indexes for performance
		`CREATE INDEX IF NOT EXISTS idx_player_profiles_github_id ON player_profiles(github_id)`,
		`CREATE INDEX IF NOT EXISTS idx_player_profiles_server_token ON player_profiles(server_token)`,
		`CREATE INDEX IF NOT EXISTS idx_client_sessions_player_id ON client_sessions(player_id)`,
		`CREATE INDEX IF NOT EXISTS idx_client_sessions_access_token ON client_sessions(access_token)`,
		`CREATE INDEX IF NOT EXISTS idx_test_accounts_username ON test_accounts(username)`,
		`CREATE INDEX IF NOT EXISTS idx_player_tokens_token ON player_tokens(token)`,
		`CREATE INDEX IF NOT EXISTS idx_player_tokens_github_id ON player_tokens(github_id)`,
	}

	for _, schema := range schemas {
		if _, err := d.conn.Exec(schema); err != nil {
			return fmt.Errorf("failed to execute schema: %w, query: %s", err, schema)
		}
	}
	return nil
}

// CreatePlayerProfile creates a new player profile
func (d *Database) CreatePlayerProfile(profile *model.PlayerProfile) error {
	query := `
		INSERT INTO player_profiles (player_id, github_id, github_login, email, avatar_url, 
			is_test_account, server_token, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		ON CONFLICT (player_id) DO NOTHING
		RETURNING id
	`
	err := d.conn.QueryRow(query,
		profile.PlayerID, profile.GithubID, profile.GithubLogin, profile.Email,
		profile.AvatarURL, profile.IsTestAccount, profile.ServerToken).Scan(&profile.ID)

	if err == sql.ErrNoRows {
		return errors.New("player profile already exists")
	}
	return err
}

// GetPlayerProfile retrieves a player profile by player_id
func (d *Database) GetPlayerProfile(playerID string) (*model.PlayerProfile, error) {
	query := `
		SELECT id, player_id, github_id, github_login, email, avatar_url, is_test_account,
			server_token, created_at, updated_at, last_login_at, token_expire_at
		FROM player_profiles
		WHERE player_id = $1
	`
	profile := &model.PlayerProfile{}
	err := d.conn.QueryRow(query, playerID).Scan(
		&profile.ID, &profile.PlayerID, &profile.GithubID, &profile.GithubLogin,
		&profile.Email, &profile.AvatarURL, &profile.IsTestAccount, &profile.ServerToken,
		&profile.CreatedAt, &profile.UpdatedAt, &profile.LastLoginAt, &profile.TokenExpireAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("player profile not found")
	}
	return profile, err
}

// GetPlayerProfileByServerToken retrieves a player profile by server token
func (d *Database) GetPlayerProfileByServerToken(serverToken string) (*model.PlayerProfile, error) {
	query := `
		SELECT id, player_id, github_id, github_login, email, avatar_url, is_test_account,
			server_token, created_at, updated_at, last_login_at, token_expire_at
		FROM player_profiles
		WHERE server_token = $1
	`
	profile := &model.PlayerProfile{}
	err := d.conn.QueryRow(query, serverToken).Scan(
		&profile.ID, &profile.PlayerID, &profile.GithubID, &profile.GithubLogin,
		&profile.Email, &profile.AvatarURL, &profile.IsTestAccount, &profile.ServerToken,
		&profile.CreatedAt, &profile.UpdatedAt, &profile.LastLoginAt, &profile.TokenExpireAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("player profile not found")
	}
	return profile, err
}

// GetPlayerProfileByGithubID retrieves a player profile by GitHub ID
func (d *Database) GetPlayerProfileByGithubID(githubID int) (*model.PlayerProfile, error) {
	query := `
		SELECT id, player_id, github_id, github_login, email, avatar_url, is_test_account,
			server_token, created_at, updated_at, last_login_at, token_expire_at
		FROM player_profiles
		WHERE github_id = $1
	`
	profile := &model.PlayerProfile{}
	err := d.conn.QueryRow(query, githubID).Scan(
		&profile.ID, &profile.PlayerID, &profile.GithubID, &profile.GithubLogin,
		&profile.Email, &profile.AvatarURL, &profile.IsTestAccount, &profile.ServerToken,
		&profile.CreatedAt, &profile.UpdatedAt, &profile.LastLoginAt, &profile.TokenExpireAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("player profile not found")
	}
	return profile, err
}

// UpdatePlayerProfile updates a player profile
func (d *Database) UpdatePlayerProfile(profile *model.PlayerProfile) error {
	query := `
		UPDATE player_profiles
		SET github_id = $1, github_login = $2, email = $3, avatar_url = $4,
			server_token = $5, last_login_at = NOW(), updated_at = NOW()
		WHERE player_id = $6
	`
	result, err := d.conn.Exec(query,
		profile.GithubID, profile.GithubLogin, profile.Email, profile.AvatarURL,
		profile.ServerToken, profile.PlayerID)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("player profile not found")
	}

	return nil
}

// CreateClientSession creates a new client session
func (d *Database) CreateClientSession(session *model.ClientSession) error {
	query := `
		INSERT INTO client_sessions (player_id, client_name, access_token, client_type, 
			created_at, expires_at)
		VALUES ($1, $2, $3, $4, NOW(), $5)
		RETURNING id, last_active_at
	`
	err := d.conn.QueryRow(query,
		session.PlayerID, session.ClientName, session.AccessToken, session.ClientType,
		session.ExpiresAt).Scan(&session.ID, &session.LastActiveAt)

	return err
}

// GetClientSession retrieves a client session by access token
func (d *Database) GetClientSession(accessToken string) (*model.ClientSession, error) {
	query := `
		SELECT id, player_id, client_name, access_token, client_type, last_active_at, 
			created_at, expires_at
		FROM client_sessions
		WHERE access_token = $1
	`
	session := &model.ClientSession{}
	err := d.conn.QueryRow(query, accessToken).Scan(
		&session.ID, &session.PlayerID, &session.ClientName, &session.AccessToken,
		&session.ClientType, &session.LastActiveAt, &session.CreatedAt, &session.ExpiresAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("session not found")
	}
	return session, err
}

// GetPlayerSessions retrieves all active sessions for a player
func (d *Database) GetPlayerSessions(playerID string) ([]*model.ClientSession, error) {
	query := `
		SELECT id, player_id, client_name, access_token, client_type, last_active_at, 
			created_at, expires_at
		FROM client_sessions
		WHERE player_id = $1 AND (expires_at IS NULL OR expires_at > NOW())
		ORDER BY last_active_at DESC
	`
	rows, err := d.conn.Query(query, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []*model.ClientSession
	for rows.Next() {
		session := &model.ClientSession{}
		if err := rows.Scan(
			&session.ID, &session.PlayerID, &session.ClientName, &session.AccessToken,
			&session.ClientType, &session.LastActiveAt, &session.CreatedAt, &session.ExpiresAt,
		); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	return sessions, rows.Err()
}

// CreateTestAccount creates a test account
func (d *Database) CreateTestAccount(account *model.TestAccount) error {
	query := `
		INSERT INTO test_accounts (username, password, player_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id
	`
	err := d.conn.QueryRow(query,
		account.Username, account.Password, account.PlayerID, account.Status).Scan(&account.ID)

	return err
}

// GetTestAccount retrieves a test account by username
func (d *Database) GetTestAccount(username string) (*model.TestAccount, error) {
	query := `
		SELECT id, username, password, player_id, status, created_at, updated_at
		FROM test_accounts
		WHERE username = $1
	`
	account := &model.TestAccount{}
	err := d.conn.QueryRow(query, username).Scan(
		&account.ID, &account.Username, &account.Password, &account.PlayerID,
		&account.Status, &account.CreatedAt, &account.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("test account not found")
	}
	return account, err
}

// ListTestAccounts lists all test accounts with optional filtering
func (d *Database) ListTestAccounts(status string) ([]*model.TestAccount, error) {
	var query string
	var rows *sql.Rows
	var err error

	if status == "" {
		query = `
			SELECT id, username, password, player_id, status, created_at, updated_at
			FROM test_accounts
			ORDER BY created_at DESC
		`
		rows, err = d.conn.Query(query)
	} else {
		query = `
			SELECT id, username, password, player_id, status, created_at, updated_at
			FROM test_accounts
			WHERE status = $1
			ORDER BY created_at DESC
		`
		rows, err = d.conn.Query(query, status)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*model.TestAccount
	for rows.Next() {
		account := &model.TestAccount{}
		if err := rows.Scan(
			&account.ID, &account.Username, &account.Password, &account.PlayerID,
			&account.Status, &account.CreatedAt, &account.UpdatedAt,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, rows.Err()
}

// UpdateTestAccountStatus updates a test account's status
func (d *Database) UpdateTestAccountStatus(username string, status string) error {
	query := `
		UPDATE test_accounts
		SET status = $1, updated_at = NOW()
		WHERE username = $2
	`
	result, err := d.conn.Exec(query, status, username)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("test account not found")
	}
	return nil
}

// CreatePlayerToken creates a new player token for GitHub binding
func (d *Database) CreatePlayerToken(token *model.PlayerToken) error {
	query := `
		INSERT INTO player_tokens (player_id, github_id, github_login, token, is_used, 
			expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
		RETURNING id
	`
	err := d.conn.QueryRow(query,
		token.PlayerID, token.GithubID, token.GithubLogin, token.Token,
		token.IsUsed, token.ExpiresAt).Scan(&token.ID)

	return err
}

// GetPlayerToken retrieves a player token
func (d *Database) GetPlayerToken(token string) (*model.PlayerToken, error) {
	query := `
		SELECT id, player_id, github_id, github_login, token, is_used, used_at, 
			expires_at, created_at
		FROM player_tokens
		WHERE token = $1 AND (expires_at IS NULL OR expires_at > NOW())
	`
	pt := &model.PlayerToken{}
	err := d.conn.QueryRow(query, token).Scan(
		&pt.ID, &pt.PlayerID, &pt.GithubID, &pt.GithubLogin, &pt.Token,
		&pt.IsUsed, &pt.UsedAt, &pt.ExpiresAt, &pt.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("player token not found or expired")
	}
	return pt, err
}

// UsePlayerToken marks a player token as used
func (d *Database) UsePlayerToken(token string) error {
	query := `
		UPDATE player_tokens
		SET is_used = true, used_at = NOW()
		WHERE token = $1
	`
	result, err := d.conn.Exec(query, token)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("player token not found")
	}
	return nil
}

// UpdateClientSessionActivity updates the last active time for a session
func (d *Database) UpdateClientSessionActivity(accessToken string) error {
	query := `
		UPDATE client_sessions
		SET last_active_at = NOW()
		WHERE access_token = $1
	`
	_, err := d.conn.Exec(query, accessToken)
	return err
}
