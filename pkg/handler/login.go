package handler

import (
	"github.com/deissh/rak-game/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func CreateLogin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.Query("nick")

		if login == "" {
			c.JSON(http.StatusUnauthorized, gin.H{})
		}

		var user types.Users
		db.Where("nickname = ?", login).First(&user)

		if user.Nickname != login {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"message": "not finded",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"user": user,
			})
		}
	}
}
