package chats

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	shared_api "github.com/totsumaru/card-chat-be/api/internal/res"
	"github.com/totsumaru/card-chat-be/api/internal/session"
	"github.com/totsumaru/card-chat-be/context/chat/expose/user"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	Chats []shared_api.ChatRes `json:"chats"`
}

// 自分がホストの全てのチャットを取得します
func FindChats(e *gin.Engine, db *gorm.DB) {
	e.GET("/api/chats", func(c *gin.Context) {
		// 認証
		ok, verifyRes := session.Verify(c)
		if !ok {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		tx := db.Begin()
		if tx.Error != nil {
			api_err.Send(c, 500, errors.NewError("Txを開始できません", tx.Error))
			return
		}

		res := Res{}

		backendErr := func() error {
			allChats, err := user.FindByHostID(tx, verifyRes.HostID)
			if err != nil {
				return errors.NewError("ホストIDに一致するチャットを取得できません", err)
			}

			resChats := make([]shared_api.ChatRes, 0)
			for _, chat := range allChats {
				resChats = append(resChats, shared_api.ChatResForHost(chat))
			}

			res.Chats = resChats

			return nil
		}()
		if backendErr != nil {
			tx.Rollback()
			api_err.Send(c, 500, errors.NewError("バックエンドの処理が失敗しました", backendErr))
			return
		}

		tx.Commit()

		c.JSON(200, res)
	})
}
