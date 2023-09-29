package edit

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ゲストの情報を編集します
func UpdateGuestInfo(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/edit", func(c *gin.Context) {
		chatID := c.Param("chatID")

		displayName := c.PostForm("display_name")
		memo := c.PostForm("memo")

		// 認証
		ok, verifyRes := verify.VerifyToken(c)
		if !ok {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		err := db.Transaction(func(tx *gorm.DB) error {
			// ホストかどうかを確認
			chatRes, err := chat_expose.FindByID(tx, chatID)
			if err != nil {
				return errors.NewError("IDでチャットを取得できません", err)
			}

			if chatRes.HostID != verifyRes.HostID {
				return errors.NewError("ホストではありません")
			}

			_, err = chat_expose.UpdateGuestInfo(tx, chatID, displayName, memo)
			if err != nil {
				return errors.NewError("ゲストの情報を更新できません", err)
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
