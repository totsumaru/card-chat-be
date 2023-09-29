package start

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/session"
	"github.com/totsumaru/card-chat-be/context/chat/expose/user"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ホストIDを登録します(チャットの開始)
func RegisterHostID(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/start", func(c *gin.Context) {
		chatID := c.Param("chatID")
		displayName := c.PostForm("display_name")

		// 認証
		ok, res := session.Verify(c)
		if !ok {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		tx := db.Begin()
		if tx.Error != nil {
			api_err.Send(c, 500, errors.NewError("Txを開始できません", tx.Error))
			return
		}

		backendErr := func() error {
			_, err := user.StartChat(tx, chatID, res.HostID, displayName)
			if err != nil {
				return errors.NewError("チャットを開始できません", err)
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
