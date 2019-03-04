package types

import "github.com/jinzhu/gorm"

type Quests struct {
	gorm.Model

	Name string `json:"name"`
	Body string `json:"body"`
	NpcId string `json:"npc_id"`
}

