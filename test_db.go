package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	connStr := "user=postgres password=xiaodudu dbname=agent_monster sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer db.Close()

	query := `
		INSERT INTO battles (id, attacker_id, defender_id, attacker_name, defender_name,
			winner, turns, attack_stack, defense_stack, battle_log, start_time, end_time,
			is_valid, validation_msg)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`
	_, err = db.Exec(query,
		"test_battle_1", "pet1", "pet2", "Pet 1", "Pet 2",
		"pet1", 5, "{}", "{}", "{\"log1\"}", time.Now(), time.Now(), true, "")
	if err != nil {
		fmt.Println("Error inserting:", err)
	} else {
		fmt.Println("Success!")
	}
}
