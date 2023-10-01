package read

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// 既読処理をします
func ChangeToRead(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/read", func(c *gin.Context) {
		chatID := c.Param("chatID")

		// 認証
		isLogin, verifyRes := verify.VerifyToken(c)
		if !isLogin {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		// ホストかどうか検証します
		isHost, err := verify.IsHost(db, chatID, verifyRes.HostID)
		if err != nil || !isHost {
			api_err.Send(c, 401, errors.NewError("ホストを認証できません", err))
			return
		}

		err = db.Transaction(func(tx *gorm.DB) error {
			_, err = chat_expose.UpdateIsRead(tx, chatID, true)
			if err != nil {
				return errors.NewError("既読処理に失敗しました", err)
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
