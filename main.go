package main

import (
	"github.com/deissh/rak-game/pkg/handler"
	"github.com/deissh/rak-game/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	db, err := gorm.Open(os.Args[1], os.Args[2])
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&types.Users{}, &types.Npcs{}, &types.Quests{}, &types.Answers{})

	r := gin.Default()
	r.GET("/quests", handler.CreateQuests(db))

	if err := r.Run(); err != nil {
		panic(err)
	}
}
