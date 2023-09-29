package message

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	message_expose "github.com/totsumaru/card-chat-be/context/message/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// メッセージを送信します
//
// 自分がホストの場合は、fromをホストに、
// ヘッダーのパスコードが一致する場合は、fromをチャットID(ゲスト)にします。
func SendMessage(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/message", func(c *gin.Context) {
		chatID := c.Param("chatID")
		passcode := c.GetHeader("Passcode")

		content := c.PostForm("content")

		fromID := ""

		err := db.Transaction(func(tx *gorm.DB) error {
			// 認証
			isLogin, verifyRes := verify.VerifyToken(c)
			if isLogin {
				// ホストかどうかを確認
				chatRes, err := chat_expose.FindByID(tx, chatID)
				if err != nil {
					return errors.NewError("IDでチャットを取得できません", err)
				}
				if chatRes.HostID != verifyRes.HostID {
					return errors.NewError("ホストではありません", err)
				}
				// 自分がホストの場合はfromIDにホストIDを入れます
				fromID = verifyRes.HostID
			} else {
				// パスコードが正しいかを確認します
				if !chat_expose.IsValidPasscode(chatID, passcode) {
					return errors.NewError("パスコードが一致しません")
				}
				// チャットIDをfromIDに設定します
				fromID = chatID
			}

			_, err := message_expose.CreateMessage(tx, chatID, fromID, content)
			if err != nil {
				return errors.NewError("メッセージを作成できません", err)
			}

			return nil
		})
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		c.JSON(200, "")
	})
}
