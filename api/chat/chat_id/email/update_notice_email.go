package email

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// 通知用のメールアドレスを変更します
func UpdateNoticeEmail(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/email", func(c *gin.Context) {
		chatID := c.Param("chatID")
		passcode := c.GetHeader("Passcode")
		email := c.PostForm("email")

		// パスコードが一致するかどうかを確認
		if !chat_expose.IsValidPasscode(chatID, passcode) {
			api_err.Send(c, 401, errors.NewError("パスコードが一致しません"))
			return
		}

		tx := db.Begin()
		if tx.Error != nil {
			api_err.Send(c, 500, errors.NewError("Txを開始できません", tx.Error))
			return
		}

		backendErr := func() error {
			_, err := chat_expose.UpdateEmail(tx, chatID, email)
			if err != nil {
				return errors.NewError("メールアドレスを更新できません", err)
			}

			return nil
		}
		if backendErr != nil {
			tx.Rollback()
			api_err.Send(c, 500, errors.NewError("バックエンドの処理が失敗しました", backendErr()))
			return
		}

		tx.Commit()

		c.JSON(200, "")
	})
}
