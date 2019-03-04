package handler

import (
	"fmt"
	"github.com/deissh/rak-game/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type Npc struct{
	Name string `json:"name"`
	Questes []Questes `json:"questes"`
}
type Questes struct{
	Name string `json:"name"`
	Body string `json:"body"`
	Answers []string `json:"answers"`
}

func CreateQuests(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Table("npcs").Rows()
		defer rows.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
		}

		result := make(map[string]Npc)

		for rows.Next() {
			var npc types.Npcs
			if err := db.ScanRows(rows, &npc); err != nil {
				log.Println(err)
			}

			rows_quests, _ := db.Table("quests").Select("name, body, id").Where("npc_id == " + fmt.Sprint(npc.ID)).Rows()
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

			result[fmt.Sprint(npc.ID)] = Npc{
				Name: npc.Name,
				Questes: questes,
			}
		}

		c.JSON(http.StatusOK, result)
	}
}