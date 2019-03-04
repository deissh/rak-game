package types

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model

	FirstName string `json:"first_name"`
	SecondName string `json:"second_name"`
	Nickname string `json:"nickname"`
	Points int `json:"points"`
	QuestCount int `json:"quest_count"`
}