package types

import "github.com/jinzhu/gorm"

type Npcs struct {
	gorm.Model

	Name string `json:"name"`
}
