package api

import (
	"github.com/gin-gonic/gin"
	chat_create "github.com/totsumaru/card-chat-be/api/chat/create"
	host_create "github.com/totsumaru/card-chat-be/api/host/create"
	"github.com/totsumaru/card-chat-be/api/host/host_id"
	host_edit "github.com/totsumaru/card-chat-be/api/host/host_id/edit"
	"gorm.io/gorm"
)

// ルートを設定します
func RegisterRouter(e *gin.Engine, db *gorm.DB) {
	Route(e)
	// `/chat`
	chat_create.CreateChat(e, db)

	// `/chats`

	// `/host`
	host_id.GetHost(e, db)
	host_edit.EditHostProfile(e, db)
	host_create.CreateHost(e, db)
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
