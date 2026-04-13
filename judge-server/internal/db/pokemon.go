package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"judge-server/internal/model"

	"github.com/lib/pq"
)

// GetPokemonSpecies 获取Pokemon物种信息
func (d *Database) GetPokemonSpecies(pokemonID string) (*model.PokemonSpecies, error) {
	var ps model.PokemonSpecies
	var metadata, stats, genes sql.NullString
	var types pq.StringArray

	err := d.conn.QueryRow(`
		SELECT id, pokemon_id, name_en, name_zh, name_jp, types, total_stats, metadata, stats, genes, created_at, updated_at
		FROM pokemon_species
		WHERE pokemon_id = $1
	`, pokemonID).Scan(
		&ps.ID, &ps.PokemonID, &ps.NameEn, &ps.NameZh, &ps.NameJp, &types, &ps.TotalStats,
		&metadata, &stats, &genes, &ps.CreatedAt, &ps.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("pokemon species not found")
		}
		return nil, err
	}

	ps.Types = []string(types)

	if metadata.Valid {
		json.Unmarshal([]byte(metadata.String), &ps.Metadata)
	}
	if stats.Valid {
		json.Unmarshal([]byte(stats.String), &ps.Stats)
	}
	if genes.Valid {
		json.Unmarshal([]byte(genes.String), &ps.Genes)
	}

	return &ps, nil
}

// GetAllPokemonSpecies 获取所有Pokemon物种
func (d *Database) GetAllPokemonSpecies(limit int, offset int) ([]*model.PokemonSpecies, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	rows, err := d.conn.Query(`
		SELECT id, pokemon_id, name_en, name_zh, name_jp, types, total_stats, created_at, updated_at
		FROM pokemon_species
		ORDER BY pokemon_id::numeric
		LIMIT $1 OFFSET $2
	`, limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var species []*model.PokemonSpecies
	for rows.Next() {
		var ps model.PokemonSpecies
		var types pq.StringArray
		err := rows.Scan(
			&ps.ID, &ps.PokemonID, &ps.NameEn, &ps.NameZh, &ps.NameJp, &types, &ps.TotalStats,
			&ps.CreatedAt, &ps.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		ps.Types = []string(types)
		species = append(species, &ps)
	}

	return species, rows.Err()
}

// SearchPokemon 搜索Pokemon
func (d *Database) SearchPokemon(keyword string) ([]*model.PokemonSpecies, error) {
	keyword = "%" + keyword + "%"
	rows, err := d.conn.Query(`
		SELECT id, pokemon_id, name_en, name_zh, name_jp, types, total_stats, created_at, updated_at
		FROM pokemon_species
		WHERE name_en ILIKE $1 OR name_zh ILIKE $1 OR name_jp ILIKE $1 OR pokemon_id LIKE $1
		ORDER BY pokemon_id::numeric
		LIMIT 50
	`, keyword)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var species []*model.PokemonSpecies
	for rows.Next() {
		var ps model.PokemonSpecies
		var types pq.StringArray
		err := rows.Scan(
			&ps.ID, &ps.PokemonID, &ps.NameEn, &ps.NameZh, &ps.NameJp, &types, &ps.TotalStats,
			&ps.CreatedAt, &ps.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		ps.Types = []string(types)
		species = append(species, &ps)
	}

	return species, rows.Err()
}

// GetPokemonCount 获取Pokemon总数
func (d *Database) GetPokemonCount() (int, error) {
	var count int
	err := d.conn.QueryRow("SELECT COUNT(*) FROM pokemon_species").Scan(&count)
	return count, err
}

// GetEffortStats 获取宝可梦努力值
func (d *Database) GetEffortStats(userPokemonID int) (*model.PokemonEffortStats, error) {
	var pes model.PokemonEffortStats
	err := d.conn.QueryRow(`
		SELECT id, user_pokemon_id, ev_hp, ev_attack, ev_defense, 
		       ev_sp_atk, ev_sp_def, ev_speed, total_ev, updated_at
		FROM pokemon_effort_stats
		WHERE user_pokemon_id = $1
	`, userPokemonID).Scan(
		&pes.ID, &pes.UserPkmID, &pes.EVHP, &pes.EVAttack, &pes.EVDefense,
		&pes.EVSpAtk, &pes.EVSpDef, &pes.EVSpeed, &pes.TotalEV,
		&pes.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("effort stats not found for user pokemon: %d", userPokemonID)
		}
		return nil, err
	}

	return &pes, nil
}

// UpdateEffortStats 更新宝可梦努力值
func (d *Database) UpdateEffortStats(stats *model.PokemonEffortStats) error {
	_, err := d.conn.Exec(`
		UPDATE pokemon_effort_stats
		SET ev_hp = $1, ev_attack = $2, ev_defense = $3,
		    ev_sp_atk = $4, ev_sp_def = $5, ev_speed = $6,
		    total_ev = $7, updated_at = NOW()
		WHERE user_pokemon_id = $8
	`,
		stats.EVHP, stats.EVAttack, stats.EVDefense,
		stats.EVSpAtk, stats.EVSpDef, stats.EVSpeed,
		stats.TotalEV, stats.UserPkmID,
	)
	return err
}

// GetNature 获取性格信息
func (d *Database) GetNature(natureID string) (*model.PokemonNature, error) {
	var pn model.PokemonNature
	err := d.conn.QueryRow(`
		SELECT id, nature_id, name_en, name_zh, increased_stat, decreased_stat, 
		       favorite_flavor, disliked_flavor, created_at
		FROM pokemon_natures
		WHERE nature_id = $1
	`, natureID).Scan(
		&pn.ID, &pn.NatureID, &pn.NameEn, &pn.NameZh, &pn.IncreasedStat, &pn.DecreasedStat,
		&pn.FavoriteFlavor, &pn.DislikedFlavor, &pn.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("nature not found: %s", natureID)
		}
		return nil, err
	}

	return &pn, nil
}

// UpdatePokemonNature 更新宝可梦性格
func (d *Database) UpdatePokemonNature(userPokemonID int, natureID string) error {
	// First verify the nature exists
	if _, err := d.GetNature(natureID); err != nil {
		return err
	}

	// Update user pokemon with new nature
	_, err := d.conn.Exec(`
		UPDATE user_pokemon
		SET nature_id = $1, updated_at = NOW()
		WHERE id = $2
	`, natureID, userPokemonID)

	return err
}
