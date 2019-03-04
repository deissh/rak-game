package main

import (
	"fmt"
	"github.com/deissh/rak-game/pkg/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type Questes struct{
	Name string `json:"name"`
	Body string `json:"body"`
	Answers []string `json:"answers"`
}

func main() {
	db, err := gorm.Open("sqlite3", "data.sqlite")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&types.Users{}, &types.Npcs{}, &types.Quests{}, &types.Answers{})

	rows_quests, _ := db.Table("quests").Select("name, body").Where("npc_id == 2").Rows()
	var questes []Questes
	for rows_quests.Next() {
		var quest types.Quests
		if err := db.ScanRows(rows_quests, &quest); err != nil {
			log.Println(err)
		}

		rows_answer, _ := db.Table("answers").Select("body").Where("quest_id == " + fmt.Sprint(quest.ID)).Rows()
		var answs []string
		for rows_answer.Next() {
			var answ types.Answers
			if err := db.ScanRows(rows_answer, &answ); err != nil {
				log.Println(err)
			}
			answs = append(answs, answ.Body)
		}

		questes = append(questes, Questes{
			Name: quest.Name,
			Body: quest.Body,
			Answers: answs,
		})
	}
	log.Println(questes)
}
