package types

import "github.com/jinzhu/gorm"

type Answers struct {
	gorm.Model

	Body string `json:"body"`
	QuestId string `json:"quest_id"`
}
