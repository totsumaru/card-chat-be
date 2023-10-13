package create

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	ChatID   string `json:"chat_id"`
	Passcode string `json:"passcode"`
}

// チャットを作成します
func CreateChat(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/create", func(c *gin.Context) {
		// 管理者でない場合はエラーを返します
		if !verify.IsAdmin(c) {
			api_err.Send(c, 401, errors.NewError("管理者ではありません"))
			return
		}

		res := Res{}
		err := db.Transaction(func(tx *gorm.DB) error {
			resp, err := chat_expose.CreateChatForAdmin(tx)
			if err != nil {
				return errors.NewError("新規チャットを作成できません", err)
			}

			res.ChatID = resp.ID
			res.Passcode = resp.Passcode

			return nil
		})
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		c.JSON(200, res)
	})
}
