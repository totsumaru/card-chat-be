package api_err

import (
	"log"

	"github.com/gin-gonic/gin"
)

// APIのエラーです
func Send(c *gin.Context, status int, err error) {
	var msg string
	switch status {
	case 400:
		msg = "リクエストが不正です"
	case 401:
		msg = "認証できません"
	case 500:
		msg = "エラーが発生しました"
	default:
		msg = "エラーが発生しました"
	}

	log.Println(err)
	c.JSON(status, gin.H{
		"message": msg,
	})
}
