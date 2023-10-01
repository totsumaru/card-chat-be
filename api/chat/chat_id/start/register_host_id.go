package start

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ホストIDを登録します(チャットの開始)
func RegisterHostID(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/start", func(c *gin.Context) {
		chatID := c.Param("chatID")

		displayName := c.PostForm("display_name")

		// 認証
		isLogin, res := verify.VerifyToken(c)
		if !isLogin {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		err := db.Transaction(func(tx *gorm.DB) error {
			_, err := chat_expose.StartChat(tx, chatID, res.HostID, displayName)
			if err != nil {
				return errors.NewError("チャットを開始できません", err)
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
