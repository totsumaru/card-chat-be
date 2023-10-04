package email

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/cookie"
	"github.com/totsumaru/card-chat-be/api/internal/res"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"github.com/totsumaru/card-chat-be/shared/resend"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	Chat res.ChatAPIRes `json:"chat"`
}

// 通知用のメールアドレスを変更します
func UpdateNoticeEmail(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/email", func(c *gin.Context) {
		chatID := c.Param("chatID")

		email := c.PostForm("email")

		cookiePasscode, err := c.Cookie(cookie.PassKey(chatID))
		if err != nil {
			api_err.Send(c, 401, errors.NewError("cookieのパスコードを取得できません", err))
			return
		}

		// パスコードが一致するかどうかを確認
		if !chat_expose.IsValidPasscode(chatID, cookiePasscode) {
			api_err.Send(c, 401, errors.NewError("パスコードが一致しません"))
			return
		}

		// Tx
		apiRes := Res{}
		err = db.Transaction(func(tx *gorm.DB) error {
			chatExposeRes, err := chat_expose.UpdateEmail(tx, chatID, email)
			if err != nil {
				return errors.NewError("メールアドレスを更新できません", err)
			}

			apiRes.Chat = res.CastToChatAPIResForGuest(chatExposeRes)

			return nil
		})
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		// メールを送信します
		if err = resend.SendEmailEdited(email); err != nil {
			log.Println("メールの送信に失敗しました", err)
			return
		}

		c.JSON(200, apiRes)
	})
}
