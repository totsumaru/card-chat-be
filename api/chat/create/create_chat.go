package create

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/session"
	"github.com/totsumaru/card-chat-be/context/chat/expose/admin"
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
		// 管理者であることを確認します
		if !session.IsAdmin(c) {
			api_err.Send(c, 401, errors.NewError("管理者ではありません"))
			return
		}

		tx := db.Begin()
		if tx.Error != nil {
			api_err.Send(c, 500, errors.NewError("Txを開始できません", tx.Error))
			return
		}

		res := Res{}

		// バックエンドの処理を実行します
		apiErr := func() error {
			resp, err := admin.CreateChat(tx)
			if err != nil {
				return errors.NewError("新規チャットを作成できません", err)
			}

			res.ChatID = resp.ID
			res.Passcode = resp.Passcode

			return nil
		}()
		if apiErr != nil {
			tx.Rollback()
			api_err.Send(c, 500, errors.NewError("バックエンドの処理が失敗しました", apiErr))
			return
		}

		tx.Commit()

		c.JSON(200, "")
	})
}