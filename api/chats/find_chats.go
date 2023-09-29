package chats

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	shared_api "github.com/totsumaru/card-chat-be/api/internal/res"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	Chats []shared_api.ChatAPIRes `json:"chats"`
}

// 自分がホストの全てのチャットを取得します
func FindChats(e *gin.Engine, db *gorm.DB) {
	e.GET("/api/chats", func(c *gin.Context) {
		// 認証
		ok, verifyRes := verify.VerifyToken(c)
		if !ok {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		res := Res{}

		backendErr := func() error {
			allChats, err := chat_expose.FindByHostID(db, verifyRes.HostID)
			if err != nil {
				return errors.NewError("ホストIDに一致するチャットを取得できません", err)
			}

			resChats := make([]shared_api.ChatAPIRes, 0)
			for _, chat := range allChats {
				resChats = append(resChats, shared_api.CastToChatAPIResForHost(chat))
			}

			res.Chats = resChats

			return nil
		}()
		if backendErr != nil {
			api_err.Send(c, 500, errors.NewError("バックエンドの処理が失敗しました", backendErr))
			return
		}

		c.JSON(200, res)
	})
}
