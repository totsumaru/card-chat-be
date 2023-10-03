package edit

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/res"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// APIのレスポンスです
type Res struct {
	Chat res.ChatAPIRes `json:"chat"`
}

// ゲストの情報を編集します
//
// ホストのみが実行できます。
func UpdateGuestInfo(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/edit", func(c *gin.Context) {
		chatID := c.Param("chatID")

		displayName := c.PostForm("display_name")
		memo := c.PostForm("memo")

		// 認証
		isLogin, verifyRes := verify.VerifyToken(c)
		if !isLogin {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		// ホストかどうかを確認します
		isHost, err := verify.IsHost(db, chatID, verifyRes.HostID)
		if err != nil {
			api_err.Send(c, 500, errors.NewError("ホストの確認ができません"))
			return
		}
		if !isHost {
			api_err.Send(c, 401, errors.NewError("ホストではありません"))
			return
		}

		// Tx
		apiRes := Res{}
		err = db.Transaction(func(tx *gorm.DB) error {
			chatExposeRes, err := chat_expose.UpdateGuestInfo(tx, chatID, displayName, memo)
			if err != nil {
				return errors.NewError("ゲストの情報を更新できません", err)
			}

			apiRes.Chat = res.CastToChatAPIResForHost(chatExposeRes)

			return nil
		})
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		c.JSON(200, apiRes)
	})
}
