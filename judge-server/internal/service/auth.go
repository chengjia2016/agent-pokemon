package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"judge-server/internal/db"
	"judge-server/internal/model"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// AuthManager handles authentication operations
type AuthManager struct {
	db *db.Database
}

// NewAuthManager creates a new authentication manager
func NewAuthManager(database *db.Database) *AuthManager {
	return &AuthManager{db: database}
}

// GeneratePlayerID generates a unique player ID
func (am *AuthManager) GeneratePlayerID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return "player_" + hex.EncodeToString(b)
}

// GenerateServerToken generates a secure server token
func (am *AuthManager) GenerateServerToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// GenerateAccessToken generates a client access token
func (am *AuthManager) GenerateAccessToken() string {
	b := make([]byte, 24)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// GeneratePlayerToken generates a GitHub binding token
func (am *AuthManager) GeneratePlayerToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// HashPassword hashes a password using bcrypt
func (am *AuthManager) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// VerifyPassword verifies a password against a hash
func (am *AuthManager) VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// CreateTestAccount creates a new test account with an associated player profile
func (am *AuthManager) CreateTestAccount(username, password string) (*model.PlayerProfile, *model.TestAccount, error) {
	// Generate credentials
	playerID := am.GeneratePlayerID()
	serverToken := am.GenerateServerToken()
	hashedPassword, err := am.HashPassword(password)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create player profile
	profile := &model.PlayerProfile{
		PlayerID:      playerID,
		IsTestAccount: true,
		ServerToken:   serverToken,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := am.db.CreatePlayerProfile(profile); err != nil {
		return nil, nil, fmt.Errorf("failed to create player profile: %w", err)
	}

	// Create test account
	testAccount := &model.TestAccount{
		Username:  username,
		Password:  hashedPassword,
		PlayerID:  playerID,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := am.db.CreateTestAccount(testAccount); err != nil {
		return nil, nil, fmt.Errorf("failed to create test account: %w", err)
	}

	return profile, testAccount, nil
}

// AuthenticateTestAccount authenticates a test account
func (am *AuthManager) AuthenticateTestAccount(username, password string) (*model.PlayerProfile, error) {
	// Get test account
	testAccount, err := am.db.GetTestAccount(username)
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	// Check if account is active
	if testAccount.Status != "active" {
		return nil, fmt.Errorf("account is %s", testAccount.Status)
	}

	// Verify password
	if !am.VerifyPassword(testAccount.Password, password) {
		return nil, errors.New("invalid username or password")
	}

	// Get player profile
	profile, err := am.db.GetPlayerProfile(testAccount.PlayerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get player profile: %w", err)
	}

	return profile, nil
}

// CreateOrUpdatePlayerProfileFromGitHub creates or updates a player profile from GitHub credentials
func (am *AuthManager) CreateOrUpdatePlayerProfileFromGitHub(githubID int, githubLogin, email, avatarURL string) (*model.PlayerProfile, error) {
	// Try to find existing profile by GitHub ID
	profile, err := am.db.GetPlayerProfileByGithubID(githubID)

	if err == nil {
		// Update existing profile
		profile.GithubLogin = githubLogin
		profile.Email = email
		profile.AvatarURL = avatarURL
		profile.UpdatedAt = time.Now()
		now := time.Now()
		profile.LastLoginAt = &now

		if err := am.db.UpdatePlayerProfile(profile); err != nil {
			return nil, fmt.Errorf("failed to update profile: %w", err)
		}
		return profile, nil
	}

	// Create new profile
	playerID := am.GeneratePlayerID()
	serverToken := am.GenerateServerToken()

	now := time.Now()
	profile = &model.PlayerProfile{
		PlayerID:      playerID,
		GithubID:      githubID,
		GithubLogin:   githubLogin,
		Email:         email,
		AvatarURL:     avatarURL,
		IsTestAccount: false,
		ServerToken:   serverToken,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		LastLoginAt:   &now,
	}

	if err := am.db.CreatePlayerProfile(profile); err != nil {
		return nil, fmt.Errorf("failed to create profile: %w", err)
	}

	return profile, nil
}

// CreateClientSession creates a new client session
func (am *AuthManager) CreateClientSession(playerID, clientName, clientType string) (*model.ClientSession, error) {
	accessToken := am.GenerateAccessToken()
	expiresAt := time.Now().Add(30 * 24 * time.Hour) // 30 days

	session := &model.ClientSession{
		PlayerID:     playerID,
		ClientName:   clientName,
		AccessToken:  accessToken,
		ClientType:   clientType,
		CreatedAt:    time.Now(),
		ExpiresAt:    expiresAt,
		LastActiveAt: time.Now(),
	}

	if err := am.db.CreateClientSession(session); err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return session, nil
}

// AuthenticateWithServerToken authenticates using a server token
func (am *AuthManager) AuthenticateWithServerToken(serverToken string) (*model.PlayerProfile, error) {
	profile, err := am.db.GetPlayerProfileByServerToken(serverToken)
	if err != nil {
		return nil, errors.New("invalid server token")
	}

	// Check if token is expired
	if profile.TokenExpireAt != nil && profile.TokenExpireAt.Before(time.Now()) {
		return nil, errors.New("server token expired")
	}

	// Update last login
	now := time.Now()
	profile.LastLoginAt = &now
	if err := am.db.UpdatePlayerProfile(profile); err != nil {
		return nil, fmt.Errorf("failed to update last login: %w", err)
	}

	return profile, nil
}

// AuthenticateWithAccessToken authenticates using an access token
func (am *AuthManager) AuthenticateWithAccessToken(accessToken string) (*model.PlayerProfile, error) {
	session, err := am.db.GetClientSession(accessToken)
	if err != nil {
		return nil, errors.New("invalid access token")
	}

	// Check if session is expired
	if session.ExpiresAt != (time.Time{}) && session.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("access token expired")
	}

	// Update session activity
	if err := am.db.UpdateClientSessionActivity(accessToken); err != nil {
		return nil, fmt.Errorf("failed to update session activity: %w", err)
	}

	// Get player profile
	profile, err := am.db.GetPlayerProfile(session.PlayerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get player profile: %w", err)
	}

	return profile, nil
}

// CreatePlayerTokenForGitHub creates a binding token for GitHub users
func (am *AuthManager) CreatePlayerTokenForGitHub(githubID int, githubLogin string) (*model.PlayerToken, error) {
	token := am.GeneratePlayerToken()
	expiresAt := time.Now().Add(24 * time.Hour) // Token valid for 24 hours

	pt := &model.PlayerToken{
		GithubID:    githubID,
		GithubLogin: githubLogin,
		Token:       token,
		IsUsed:      false,
		ExpiresAt:   expiresAt,
		CreatedAt:   time.Now(),
	}

	if err := am.db.CreatePlayerToken(pt); err != nil {
		return nil, fmt.Errorf("failed to create player token: %w", err)
	}

	return pt, nil
}

// UsePlayerTokenToBindGitHub uses a player token to bind GitHub account
func (am *AuthManager) UsePlayerTokenToBindGitHub(playerToken string, clientName string) (*model.AuthResponse, error) {
	// Get the player token
	pt, err := am.db.GetPlayerToken(playerToken)
	if err != nil {
		return nil, fmt.Errorf("invalid or expired player token")
	}

	// Check if token already used
	if pt.IsUsed {
		return nil, errors.New("player token already used")
	}

	// Mark token as used
	if err := am.db.UsePlayerToken(playerToken); err != nil {
		return nil, fmt.Errorf("failed to use player token: %w", err)
	}

	// Get or create player profile
	profile, err := am.CreateOrUpdatePlayerProfileFromGitHub(pt.GithubID, pt.GithubLogin, "", "")
	if err != nil {
		return nil, fmt.Errorf("failed to create or update profile: %w", err)
	}

	// Create client session
	session, err := am.CreateClientSession(profile.PlayerID, clientName, "agent")
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &model.AuthResponse{
		Success:       true,
		PlayerID:      profile.PlayerID,
		ServerToken:   profile.ServerToken,
		AccessToken:   session.AccessToken,
		ClientSession: session,
		Message:       "GitHub account bound successfully",
	}, nil
}

// ListAllTestAccounts lists all test accounts
func (am *AuthManager) ListAllTestAccounts() ([]*model.TestAccount, error) {
	return am.db.ListTestAccounts("")
}

// ListActiveTestAccounts lists active test accounts
func (am *AuthManager) ListActiveTestAccounts() ([]*model.TestAccount, error) {
	return am.db.ListTestAccounts("active")
}

// DisableTestAccount disables a test account
func (am *AuthManager) DisableTestAccount(username string) error {
	return am.db.UpdateTestAccountStatus(username, "inactive")
}

// LockTestAccount locks a test account
func (am *AuthManager) LockTestAccount(username string) error {
	return am.db.UpdateTestAccountStatus(username, "locked")
}
