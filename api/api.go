package api

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/chat/chat_id"
	"github.com/totsumaru/card-chat-be/api/chat/chat_id/edit"
	"github.com/totsumaru/card-chat-be/api/chat/chat_id/email"
	"github.com/totsumaru/card-chat-be/api/chat/chat_id/passcode"
	"github.com/totsumaru/card-chat-be/api/chat/chat_id/read"
	"github.com/totsumaru/card-chat-be/api/chat/chat_id/start"
	chat_create "github.com/totsumaru/card-chat-be/api/chat/create"
	"github.com/totsumaru/card-chat-be/api/chats"
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
	edit.UpdateGuestInfo(e, db)
	chat_id.GetChat(e, db)
	email.UpdateNoticeEmail(e, db)
	passcode.GetChatByPasscode(e, db)
	read.ChangeToRead(e, db)
	start.RegisterHostID(e, db)

	// `/chats`
	chats.FindChats(e, db)

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
