package api

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/host/host_id"
	"github.com/totsumaru/card-chat-be/api/host/host_id/edit"
	"gorm.io/gorm"
)

// ルートを設定します
func RegisterRouter(e *gin.Engine, db *gorm.DB) {
	Route(e)
	// `/chat`
	// `/chats`
	// `/host`
	host_id.Host(e, db)
	edit.Edit(e, db)
}

// ルートです
//
// Note: この関数は削除しても問題ありません
func Route(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
}
