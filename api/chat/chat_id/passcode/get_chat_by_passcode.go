package passcode

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/res"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	message_expose "github.com/totsumaru/card-chat-be/context/message/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	Chat     res.ChatRes         `json:"chat"`
	Messages []res.MessageAPIRes `json:"messages"`
}

// チャットを取得します
func GetChatByPasscode(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/passcode", func(c *gin.Context) {
		chatID := c.Param("chatID")
		passcode := c.GetHeader("Passcode")

		tx := db.Begin()
		if tx.Error != nil {
			api_err.Send(c, 500, errors.NewError("Txを開始できません", tx.Error))
			return
		}

		if !chat_expose.IsValidPasscode(chatID, passcode) {
			api_err.Send(c, 401, errors.NewError("パスコードが一致しません", tx.Error))
			return
		}

		response := Res{}

		backendErr := func() error {
			// チャットを取得
			apiChatRes, err := chat_expose.FindByID(tx, chatID)
			if err != nil {
				return errors.NewError("IDでチャットを取得できません", err)
			}

			if apiChatRes.HostID == "" {
				return errors.NewError("チャットが開始されていません")
			}

			// 全てのメッセージを取得します
			msgs, err := message_expose.FindByChatID(tx, apiChatRes.ID)
			if err != nil {
				return errors.NewError("チャットIDでメッセージを取得できません", err)
			}

			response.Chat = res.ChatResForGuest(apiChatRes)
			response.Messages = res.CastToMessagesAPIRes(msgs)

			return nil
		}()
		if backendErr != nil {
			tx.Rollback()
			api_err.Send(c, 500, errors.NewError("バックエンドの処理が失敗しました", backendErr))
			return
		}

		tx.Commit()

		c.JSON(200, response)
	})
}
