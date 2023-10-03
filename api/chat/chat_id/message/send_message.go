package message

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/cookie"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	message_expose "github.com/totsumaru/card-chat-be/context/message/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// -------------------------------------------------------
// このAPIはレスポンスを返しません。
//
// 一定間隔で`GetChat`のAPIがコールされて
// メッセージは最新状態に同期されるため、ここでは保存だけにします。
// -------------------------------------------------------

// メッセージを送信します
//
// 自分がホストの場合は、fromをホストに、
// ヘッダーのパスコードが一致する場合は、fromをチャットID(ゲスト)にします。
func SendMessage(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/message", func(c *gin.Context) {
		chatID := c.Param("chatID")

		content := c.PostForm("content")

		var fromID string

		// 認証をします
		isLogin, verifyRes := verify.VerifyToken(c)

		// ホストの場合は、fromIDにホストIDを設定します
		if isLogin {
			// ホストの場合は、fromIDにホストIDを設定します
			isHost, err := verify.IsHost(db, chatID, verifyRes.HostID)
			if err != nil {
				api_err.Send(c, 500, errors.NewError("ホストの確認ができません", err))
				return
			}
			if isHost {
				fromID = verifyRes.HostID
			}
		} else {
			cookiePasscode, _ := c.Cookie(cookie.PassKey(chatID))
			if chat_expose.IsValidPasscode(chatID, cookiePasscode) {
				// パスコードが正しい場合は、fromIDにチャットIDを設定します
				fromID = chatID
			}
		}

		if fromID == "" {
			api_err.Send(c, 401, errors.NewError("送信者が不明です"))
			return
		}

		// Tx
		err := db.Transaction(func(tx *gorm.DB) error {
			_, err := message_expose.CreateMessage(tx, chatID, fromID, content)
			if err != nil {
				return errors.NewError("メッセージを作成できません", err)
			}

			// 未読処理を行います
			_, err = chat_expose.UpdateIsRead(tx, chatID, false)
			if err != nil {
				return errors.NewError("未読処理に失敗しました", err)
			}

			return nil
		})
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		c.JSON(200, nil)
	})
}
