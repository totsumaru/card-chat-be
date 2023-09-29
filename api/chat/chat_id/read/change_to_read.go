package read

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/session"
	"github.com/totsumaru/card-chat-be/context/chat/expose/user"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// 既読処理をします
func ChangeToRead(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/read", func(c *gin.Context) {
		chatID := c.Param("chatID")

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
			// ホストかどうかを確認
			chatRes, err := user.FindByID(tx, chatID)
			if err != nil {
				return errors.NewError("IDでチャットを取得できません", err)
			}

			if chatRes.HostID != res.HostID {
				return errors.NewError("ホストではありません")
			}

			_, err = user.UpdateIsRead(tx, chatID, true)
			if err != nil {
				return errors.NewError("既読処理に失敗しました", err)
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
