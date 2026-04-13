package db

import (
	"database/sql"
	"errors"
	"judge-server/internal/model"
)

// ===== Cookie Operations =====

// RegisterCookie registers a new cookie
func (d *Database) RegisterCookie(cookie *model.DatabaseCookie) (*model.DatabaseCookie, error) {
	query := `
		INSERT INTO cookies (cookie_id, cookie_type, emoji, source_file, generator_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, cookie_id, cookie_type, emoji, source_file, generator_id, created_at, updated_at
	`

	err := d.conn.QueryRow(query,
		cookie.CookieID, cookie.CookieType, cookie.Emoji, cookie.SourceFile, cookie.GeneratorID).
		Scan(&cookie.ID, &cookie.CookieID, &cookie.CookieType, &cookie.Emoji,
			&cookie.SourceFile, &cookie.GeneratorID, &cookie.CreatedAt, &cookie.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return cookie, nil
}

// GetCookieByID retrieves a cookie by its ID
func (d *Database) GetCookieByID(cookieID string) (*model.DatabaseCookie, error) {
	query := `
		SELECT id, cookie_id, cookie_type, emoji, source_file, generator_id, created_at, updated_at
		FROM cookies
		WHERE cookie_id = $1
	`

	cookie := &model.DatabaseCookie{}
	err := d.conn.QueryRow(query, cookieID).Scan(
		&cookie.ID, &cookie.CookieID, &cookie.CookieType, &cookie.Emoji,
		&cookie.SourceFile, &cookie.GeneratorID, &cookie.CreatedAt, &cookie.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("cookie not found")
	}

	return cookie, err
}

// ClaimCookie claims a cookie for a player
func (d *Database) ClaimCookie(cookieID string, playerID string, expReward int, energyReward int) (*model.DatabaseCookieClaim, error) {
	tx, err := d.conn.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Check if cookie exists
	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM cookies WHERE cookie_id = $1)", cookieID).Scan(&exists)
	if err != nil || !exists {
		return nil, errors.New("cookie not found")
	}

	// Check if already claimed by this player
	var claimExists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM cookie_claims WHERE cookie_id = $1 AND player_id = $2)", cookieID, playerID).Scan(&claimExists)
	if err != nil {
		return nil, err
	}

	if claimExists {
		return nil, errors.New("cookie already claimed by this player")
	}

	// Insert claim
	claimQuery := `
		INSERT INTO cookie_claims (cookie_id, player_id, claimed_at, exp_reward, energy_reward)
		VALUES ($1, $2, NOW(), $3, $4)
		RETURNING id, cookie_id, player_id, claimed_at, exp_reward, energy_reward
	`

	claim := &model.DatabaseCookieClaim{}
	err = tx.QueryRow(claimQuery, cookieID, playerID, expReward, energyReward).Scan(
		&claim.ID, &claim.CookieID, &claim.PlayerID, &claim.ClaimedAt, &claim.ExpReward, &claim.EnergyReward)

	if err != nil {
		return nil, err
	}

	// Log event
	eventQuery := `
		INSERT INTO cookie_events (event_type, cookie_id, player_id, event_details, created_at)
		VALUES ('claimed', $1, $2, $3, NOW())
	`
	_, err = tx.Exec(eventQuery, cookieID, playerID, `{"exp_reward": `+string(rune(expReward))+`, "energy_reward": `+string(rune(energyReward))+`}`)
	if err != nil {
		return nil, err
	}

	return claim, tx.Commit()
}

// GetCookieStats retrieves statistics for cookies
func (d *Database) GetCookieStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Total cookies
	var totalCookies int
	err := d.conn.QueryRow("SELECT COUNT(*) FROM cookies").Scan(&totalCookies)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	stats["total_cookies"] = totalCookies

	// Total claims
	var totalClaims int
	err = d.conn.QueryRow("SELECT COUNT(*) FROM cookie_claims").Scan(&totalClaims)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	stats["total_claims"] = totalClaims

	// Breakdown by type
	typeQuery := `
		SELECT cookie_type, COUNT(*) as count
		FROM cookies
		GROUP BY cookie_type
	`

	rows, err := d.conn.Query(typeQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	typeBreakdown := make(map[string]int)
	for rows.Next() {
		var cookieType string
		var count int
		if err := rows.Scan(&cookieType, &count); err != nil {
			return nil, err
		}
		typeBreakdown[cookieType] = count
	}
	stats["type_breakdown"] = typeBreakdown

	return stats, rows.Err()
}

// ===== Cookie Fragment Operations =====

// AddCookieFragment adds a fragment to a player's collection
func (d *Database) AddCookieFragment(playerID string, fragmentType string, fragmentID string, quantity int) (*model.DatabaseCookieFragment, error) {
	query := `
		INSERT INTO cookie_fragments (fragment_id, fragment_type, emoji, created_at)
		VALUES ($1, $2, $3, NOW())
		ON CONFLICT (fragment_id) DO NOTHING
		RETURNING id, fragment_id, fragment_type, emoji, created_at
	`

	fragment := &model.DatabaseCookieFragment{}
	err := d.conn.QueryRow(query, fragmentID, fragmentType, "").Scan(
		&fragment.ID, &fragment.FragmentID, &fragment.FragmentType, &fragment.Emoji, &fragment.CreatedAt)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Update player's balance
	balanceQuery := `
		INSERT INTO player_fragment_balance (player_id, fragment_id, balance)
		VALUES ($1, $2, $3)
		ON CONFLICT (player_id, fragment_id) DO UPDATE
		SET balance = player_fragment_balance.balance + $3
	`

	_, err = d.conn.Exec(balanceQuery, playerID, fragmentID, quantity)
	if err != nil {
		return nil, err
	}

	return fragment, nil
}

// GetPlayerFragmentBalance retrieves a player's fragment balance by type
func (d *Database) GetPlayerFragmentBalance(playerID string) (map[string]int, error) {
	query := `
		SELECT fragment_type, COUNT(*) as count
		FROM cookie_fragments
		WHERE player_id = $1
		GROUP BY fragment_type
	`

	rows, err := d.conn.Query(query, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	balance := make(map[string]int)
	for rows.Next() {
		var fragmentType string
		var count int
		if err := rows.Scan(&fragmentType, &count); err != nil {
			return nil, err
		}
		balance[fragmentType] = count
	}

	return balance, rows.Err()
}
