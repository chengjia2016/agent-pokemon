package model

import "time"

type PokemonSpecies struct {
	ID          int                    `json:"id"`
	PokemonID   string                 `json:"pokemon_id"`
	NameEn      string                 `json:"name_en"`
	NameZh      string                 `json:"name_zh,omitempty"`
	NameJp      string                 `json:"name_jp,omitempty"`
	Types       []string               `json:"types"`
	TotalStats  int                    `json:"total_stats"`
	CaptureRate int                    `json:"capture_rate"`           // 捕捉率
	GenderRatio map[string]int         `json:"gender_ratio,omitempty"` // 性别比例
	EggGroups   []string               `json:"egg_groups,omitempty"`   // 蛋组
	Habitat     string                 `json:"habitat,omitempty"`      // 栖息地
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Stats       map[string]interface{} `json:"stats,omitempty"`
	Genes       map[string]interface{} `json:"genes,omitempty"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}
