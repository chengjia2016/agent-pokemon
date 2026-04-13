package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Game client configuration
const (
	ServerURL = "http://localhost:10000"
	Timeout   = 10 * time.Second
)

// User represents a game player
type User struct {
	ID          int     `json:"id"`
	GitHubID    int     `json:"github_id"`
	GitHubLogin string  `json:"github_login"`
	Balance     float64 `json:"balance"`
	Email       string  `json:"email"`
}

// Quest represents a game quest
type Quest struct {
	ID          int    `json:"id"`
	QuestCode   string `json:"quest_code"`
	NameEN      string `json:"name_en"`
	NameZH      string `json:"name_zh"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Difficulty  int    `json:"difficulty"`
	RewardExp   *int   `json:"reward_exp"`
	RewardCoins *int   `json:"reward_coins"`
	LevelReq    *int   `json:"level_req"`
}

// UserQuest represents player's quest progress
type UserQuest struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	QuestID     int        `json:"quest_id"`
	Status      string     `json:"status"`
	Progress    int        `json:"progress"`
	StartedAt   time.Time  `json:"started_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

// Dungeon represents a game dungeon
type Dungeon struct {
	ID          int    `json:"id"`
	DungeonCode string `json:"dungeon_code"`
	NameEN      string `json:"name_en"`
	NameZH      string `json:"name_zh"`
	Difficulty  int    `json:"difficulty"`
	FloorCount  int    `json:"floor_count"`
	RewardExp   int    `json:"reward_exp"`
	RewardCoins int    `json:"reward_coins"`
}

// Gym represents a gym
type Gym struct {
	ID        int    `json:"id"`
	GymCode   string `json:"gym_code"`
	NameEN    string `json:"name_en"`
	NameZH    string `json:"name_zh"`
	TownID    string `json:"town_id"`
	TypeFocus string `json:"type_focus"`
	BadgeName string `json:"badge_name"`
}

// MapZone represents a game world zone
type MapZone struct {
	ID             int      `json:"id"`
	ZoneCode       string   `json:"zone_code"`
	NameEN         string   `json:"name_en"`
	NameZH         string   `json:"name_zh"`
	Description    string   `json:"description"`
	Type           string   `json:"type"`
	Level          int      `json:"level"`
	IslandID       string   `json:"island_id"`
	ConnectedZones []string `json:"connected_zones"`
}

// GameClient represents the game CLI client
type GameClient struct {
	httpClient *http.Client
	user       *User
	scanner    *bufio.Scanner
}

// NewGameClient creates a new game client
func NewGameClient() *GameClient {
	return &GameClient{
		httpClient: &http.Client{Timeout: Timeout},
		scanner:    bufio.NewScanner(os.Stdin),
	}
}

// Main game functions

func (gc *GameClient) ShowWelcome() {
	clearScreen()
	fmt.Println(`
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                              ║
║                   🎮 AGENT MONSTER - POKÉMON WORLD ADVENTURE 🎮              ║
║                                                                              ║
║                 Welcome to the Digital Pokémon Revolution!                  ║
║                                                                              ║
║              Your GitHub Repository Becomes Your Pokémon Kingdom            ║
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝
`)
	fmt.Println("\n📖 Legenday lore says: In the land of Kanto, trainers and their Pokémon")
	fmt.Println("   form unbreakable bonds by committing to their digital sanctuary...")
	fmt.Println("\n   Your repository is that sanctuary.")
	fmt.Println("   Your commits are your adventures.")
	fmt.Println("   Your Pokémon await...\n")
}

func (gc *GameClient) AuthenticateWithGitHub() error {
	fmt.Print("🔐 Enter your GitHub ID (e.g., 274799269): ")
	gc.scanner.Scan()
	githubIDStr := strings.TrimSpace(gc.scanner.Text())

	githubID, err := strconv.Atoi(githubIDStr)
	if err != nil {
		fmt.Println("❌ Invalid GitHub ID. Please enter a number.")
		return err
	}

	// Try to fetch existing user
	url := fmt.Sprintf("%s/api/users/%d", ServerURL, githubID)
	user, err := gc.fetchUser(url)
	if err == nil && user.ID > 0 {
		gc.user = user
		fmt.Printf("✅ Welcome back, %s! (Balance: %d 💰)\n", user.GitHubLogin, int(user.Balance))
		time.Sleep(2 * time.Second)
		return nil
	}

	// Create new user
	fmt.Print("👤 Enter your GitHub username (e.g., tomcooler): ")
	gc.scanner.Scan()
	githubLogin := strings.TrimSpace(gc.scanner.Text())

	fmt.Print("📧 Enter your email (e.g., user@github.com): ")
	gc.scanner.Scan()
	email := strings.TrimSpace(gc.scanner.Text())

	user, err = gc.createUser(githubID, githubLogin, email)
	if err != nil {
		fmt.Printf("❌ Failed to create account: %v\n", err)
		return err
	}

	gc.user = user
	fmt.Printf("\n✨ Welcome to the world, %s!\n", githubLogin)
	fmt.Printf("📬 You've received a starter pack bonus: %d 💰\n", int(user.Balance))
	time.Sleep(2 * time.Second)
	return nil
}

func (gc *GameClient) ShowMainMenu() {
	for {
		clearScreen()
		fmt.Printf(`
╔════════════════════════════════════════════════════════════════════════════╗
║ 🌍 POKÉMON ADVENTURE - MAIN MENU                                          ║
╚════════════════════════════════════════════════════════════════════════════╝

Trainer: %s 👤
Balance: %d 💰
Location: Kanto Region 🗺️

═════════════════════════════════════════════════════════════════════════════

What would you like to do?

[1] 📜 View & Accept Quests
[2] 🗺️  Explore Map Zones & Levels
[3] 🏰 Enter Dungeons
[4] 🏆 Visit Pokémon Gyms
[5] 🤖 Talk to NPCs
[6] 👓 View Your Progress
[7] ℹ️  Game Help & Commands
[0] 🚪 Exit Game

═════════════════════════════════════════════════════════════════════════════
`, gc.user.GitHubLogin, int(gc.user.Balance))

		fmt.Print("Enter your choice (0-7): ")
		gc.scanner.Scan()
		choice := strings.TrimSpace(gc.scanner.Text())

		switch choice {
		case "1":
			gc.showQuestMenu()
		case "2":
			gc.showMapMenu()
		case "3":
			gc.showDungeonMenu()
		case "4":
			gc.showGymMenu()
		case "5":
			gc.showNPCMenu()
		case "6":
			gc.showProgressMenu()
		case "7":
			gc.showHelpMenu()
		case "0":
			gc.exitGame()
			return
		default:
			fmt.Println("❌ Invalid choice. Please try again.")
			time.Sleep(1 * time.Second)
		}
	}
}

func (gc *GameClient) showQuestMenu() {
	clearScreen()
	fmt.Println(`
╔════════════════════════════════════════════════════════════════════════════╗
║ 📜 QUEST BOARD                                                            ║
╚════════════════════════════════════════════════════════════════════════════╝

The Quest Master's voice echoes through the hall:
"Many brave trainers seek glory through trials and tribulations.
 Perhaps, young trainer, you are ready for these challenges?"
`)

	// Fetch available quests
	quests, err := gc.fetchQuests("", 0)
	if err != nil {
		fmt.Printf("❌ Error fetching quests: %v\n", err)
		time.Sleep(2 * time.Second)
		return
	}

	if len(quests) == 0 {
		fmt.Println("No quests available at the moment.")
		time.Sleep(2 * time.Second)
		return
	}

	fmt.Println("\n📋 Available Quests:\n")
	for i, quest := range quests {
		stars := strings.Repeat("⭐", quest.Difficulty)
		reward := ""
		if quest.RewardCoins != nil {
			reward = fmt.Sprintf("%d 💰", *quest.RewardCoins)
		}
		if quest.RewardExp != nil {
			if reward != "" {
				reward += fmt.Sprintf(" + %d XP", *quest.RewardExp)
			} else {
				reward = fmt.Sprintf("%d XP", *quest.RewardExp)
			}
		}

		fmt.Printf("[%d] %s (%s)\n", i+1, quest.NameEN, strings.ToUpper(quest.Type))
		fmt.Printf("    Difficulty: %s | Reward: %s\n", stars, reward)
		fmt.Printf("    → %s\n\n", quest.Description)
	}

	fmt.Print("Press Enter to return to main menu...")
	gc.scanner.Scan()
}

func (gc *GameClient) showMapMenu() {
	clearScreen()
	fmt.Println(`
╔════════════════════════════════════════════════════════════════════════════╗
║ 🗺️  WORLD MAP - KANTO REGION                                              ║
╚════════════════════════════════════════════════════════════════════════════╝

An ancient map glows before you. The Kanto region awaits your exploration.
"Choose your path wisely, young trainer..."
`)

	// Fetch map zones
	zones, err := gc.fetchMapZones("island_1")
	if err != nil {
		fmt.Printf("❌ Error fetching map zones: %v\n", err)
		time.Sleep(2 * time.Second)
		return
	}

	if len(zones) == 0 {
		fmt.Println("No zones available.")
		time.Sleep(2 * time.Second)
		return
	}

	fmt.Println("\n🌍 Explorable Zones:\n")
	for i, zone := range zones {
		difficulty := strings.Repeat("⭐", zone.Level)
		fmt.Printf("[%d] 🌳 %s (%s)\n", i+1, zone.NameEN, zone.NameZH)
		fmt.Printf("    Type: %s | Difficulty: %s\n", strings.ToUpper(zone.Type), difficulty)
		fmt.Printf("    → %s\n\n", zone.Description)
	}

	fmt.Print("Press Enter to return to main menu...")
	gc.scanner.Scan()
}

func (gc *GameClient) showDungeonMenu() {
	clearScreen()
	fmt.Println(`
╔════════════════════════════════════════════════════════════════════════════╗
║ 🏰 DUNGEON ADVENTURES                                                     ║
╚════════════════════════════════════════════════════════════════════════════╝

Dark dungeons loom in the distance. Trainers who brave these depths
return with power beyond imagination...

"Do you have what it takes, young one?"
`)

	// Fetch dungeons
	dungeons, err := gc.fetchDungeons(0)
	if err != nil {
		fmt.Printf("❌ Error fetching dungeons: %v\n", err)
		time.Sleep(2 * time.Second)
		return
	}

	if len(dungeons) == 0 {
		fmt.Println("No dungeons available.")
		time.Sleep(2 * time.Second)
		return
	}

	fmt.Println("\n🗝️  Available Dungeons:\n")
	for i, dungeon := range dungeons {
		difficulty := strings.Repeat("⭐", dungeon.Difficulty)
		fmt.Printf("[%d] 🏰 %s (%s)\n", i+1, dungeon.NameEN, dungeon.NameZH)
		fmt.Printf("    Difficulty: %s | Floors: %d\n", difficulty, dungeon.FloorCount)
		fmt.Printf("    Reward: %d 💰 + %d XP\n\n", dungeon.RewardCoins, dungeon.RewardExp)
	}

	fmt.Print("Press Enter to return to main menu...")
	gc.scanner.Scan()
}

func (gc *GameClient) showGymMenu() {
	clearScreen()
	fmt.Println(`
╔════════════════════════════════════════════════════════════════════════════╗
║ 🏆 POKÉMON GYMS                                                           ║
╚════════════════════════════════════════════════════════════════════════════╝

The Gym Master stands before you, arms crossed.
"Show me your strength. Prove you're worthy of a badge!"
`)

	// Fetch gyms
	gyms, err := gc.fetchGyms("town_1")
	if err != nil {
		fmt.Printf("❌ Error fetching gyms: %v\n", err)
		time.Sleep(2 * time.Second)
		return
	}

	if len(gyms) == 0 {
		fmt.Println("No gyms available in this town.")
		time.Sleep(2 * time.Second)
		return
	}

	fmt.Println("\n🥊 Available Gyms:\n")
	for i, gym := range gyms {
		fmt.Printf("[%d] 🏛️  %s (%s)\n", i+1, gym.NameEN, gym.NameZH)
		fmt.Printf("    Type: %s | Badge: %s\n", gym.TypeFocus, gym.BadgeName)
		fmt.Printf("    Status: 🔓 Challengeable\n\n")
	}

	fmt.Print("Press Enter to return to main menu...")
	gc.scanner.Scan()
}

func (gc *GameClient) showNPCMenu() {
	clearScreen()
	fmt.Println(`
╔════════════════════════════════════════════════════════════════════════════╗
║ 🤖 TOWN NPCs                                                              ║
╚════════════════════════════════════════════════════════════════════════════╝

The town bustles with activity. Various NPCs go about their business...

What would you like to know?
`)

	fmt.Print("Enter NPC ID to talk to (or press Enter to go back): ")
	gc.scanner.Scan()
	input := strings.TrimSpace(gc.scanner.Text())
	if input == "" {
		return
	}

	fmt.Println("\n💬 NPC: \"Hail, trainer! What brings you to our town?\"")
	fmt.Println("\n(More NPC interactions coming soon...)")
	time.Sleep(2 * time.Second)
}

func (gc *GameClient) showProgressMenu() {
	clearScreen()
	fmt.Println(`
╔════════════════════════════════════════════════════════════════════════════╗
║ 👓 YOUR PROGRESS                                                          ║
╚════════════════════════════════════════════════════════════════════════════╝
`)

	fmt.Printf("Trainer: %s\n", gc.user.GitHubLogin)
	fmt.Printf("Balance: %d 💰\n", int(gc.user.Balance))
	fmt.Println("\n📊 Statistics:")
	fmt.Println("  • Quests Completed: 0")
	fmt.Println("  • Pokémon Caught: 0")
	fmt.Println("  • Badges Earned: 0")
	fmt.Println("  • Dungeons Cleared: 0")

	fmt.Print("\nPress Enter to return to main menu...")
	gc.scanner.Scan()
}

func (gc *GameClient) showHelpMenu() {
	clearScreen()
	fmt.Println(`
╔════════════════════════════════════════════════════════════════════════════╗
║ ℹ️  GAME HELP & GUIDE                                                      ║
╚════════════════════════════════════════════════════════════════════════════╝

🎮 MAIN MENU OPTIONS:

[1] 📜 QUESTS
    Accept quests to earn coins and experience.
    Types: Main (yellow), Side (purple)
    Complete them to progress in your adventure!

[2] 🗺️  MAP & ZONES
    Explore different areas of Kanto region.
    Each zone has different Pokémon and difficulty levels.
    Discover all 3 zones!

[3] 🏰 DUNGEONS
    Multi-floor dungeons with bosses and treasure.
    Difficulty: ⭐ Easy, ⭐⭐ Normal, ⭐⭐⭐ Hard
    Can you clear them all?

[4] 🏆 GYM BATTLES
    Challenge Gym Leaders to earn prestigious badges!
    Collect all badges to become a Pokémon Master.

[5] 🤖 NPC INTERACTIONS
    Talk to NPCs to learn about the world and get hints.

[6] 👓 PROGRESS TRACKING
    View your achievements and statistics.

💡 GAMEPLAY TIPS:

  • Start with Easy difficulty quests (⭐)
  • Build your Pokémon team before challenging gyms
  • Explore all zones to find rare Pokémon
  • Complete side quests for bonus rewards
  • Your GitHub commits fuel your Pokémon's power!

🎯 MAIN OBJECTIVES:

  1. Catch 5 different Pokémon
  2. Complete 3 Main quests
  3. Earn all 8 Gym badges
  4. Clear all Dungeons
  5. Become a Pokémon Master!

═════════════════════════════════════════════════════════════════════════════
`)

	fmt.Print("Press Enter to return to main menu...")
	gc.scanner.Scan()
}

func (gc *GameClient) exitGame() {
	clearScreen()
	fmt.Println(`
╔════════════════════════════════════════════════════════════════════════════╗
║                         UNTIL NEXT TIME, TRAINER!                         ║
║                                                                            ║
║            Your Pokémon await your return in the digital realm.           ║
║          Remember: Every commit brings you closer to greatness!          ║
╚════════════════════════════════════════════════════════════════════════════╝

✨ Save your progress... ✨
🚀 May your commits be numerous and your bugs be few!
🎮 See you soon, trainer!
`)
	os.Exit(0)
}

// API Helper Methods

func (gc *GameClient) fetchUser(url string) (*User, error) {
	resp, err := gc.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user not found")
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (gc *GameClient) createUser(githubID int, githubLogin, email string) (*User, error) {
	payload := map[string]interface{}{
		"github_id":    githubID,
		"github_login": githubLogin,
		"email":        email,
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := gc.httpClient.Post(
		fmt.Sprintf("%s/api/users/create", ServerURL),
		"application/json",
		strings.NewReader(string(jsonData)),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if !result["success"].(bool) {
		return nil, fmt.Errorf("failed to create user: %v", result["message"])
	}

	// Extract user from result
	userData, _ := json.Marshal(result["user"])
	var user User
	json.Unmarshal(userData, &user)
	return &user, nil
}

func (gc *GameClient) fetchQuests(questType string, difficulty int) ([]*Quest, error) {
	url := fmt.Sprintf("%s/api/quests", ServerURL)
	params := "?"
	if questType != "" {
		params += fmt.Sprintf("type=%s&", questType)
	}
	if difficulty > 0 {
		params += fmt.Sprintf("difficulty=%d&", difficulty)
	}
	if params != "?" {
		url += params[:len(params)-1]
	}

	resp, err := gc.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	questsData, _ := json.Marshal(result["quests"])
	var quests []*Quest
	json.Unmarshal(questsData, &quests)
	return quests, nil
}

func (gc *GameClient) fetchMapZones(islandID string) ([]*MapZone, error) {
	url := fmt.Sprintf("%s/api/map/zones?island_id=%s", ServerURL, islandID)
	resp, err := gc.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	zonesData, _ := json.Marshal(result["zones"])
	var zones []*MapZone
	json.Unmarshal(zonesData, &zones)
	return zones, nil
}

func (gc *GameClient) fetchDungeons(difficulty int) ([]*Dungeon, error) {
	url := fmt.Sprintf("%s/api/dungeons", ServerURL)
	if difficulty > 0 {
		url += fmt.Sprintf("?difficulty=%d", difficulty)
	}

	resp, err := gc.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	dungeonsData, _ := json.Marshal(result["dungeons"])
	var dungeons []*Dungeon
	json.Unmarshal(dungeonsData, &dungeons)
	return dungeons, nil
}

func (gc *GameClient) fetchGyms(townID string) ([]*Gym, error) {
	url := fmt.Sprintf("%s/api/gyms?town_id=%s", ServerURL, townID)
	resp, err := gc.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	gymsData, _ := json.Marshal(result["gyms"])
	var gyms []*Gym
	json.Unmarshal(gymsData, &gyms)
	return gyms, nil
}

// Utility functions

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	client := NewGameClient()
	client.ShowWelcome()

	if err := client.AuthenticateWithGitHub(); err != nil {
		fmt.Printf("Failed to authenticate: %v\n", err)
		os.Exit(1)
	}

	time.Sleep(1 * time.Second)
	client.ShowMainMenu()
}
