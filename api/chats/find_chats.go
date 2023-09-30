package chats

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	shared_api "github.com/totsumaru/card-chat-be/api/internal/res"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	message_expose "github.com/totsumaru/card-chat-be/context/message/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	Chats []ChatRes `json:"chats"`
}

// チャットのレスポンスです
type ChatRes struct {
	Chat          shared_api.ChatAPIRes    `json:"chat"`
	LatestMessage shared_api.MessageAPIRes `json:"latest_message"`
}

// 自分がホストの全てのチャットを取得します
func FindChats(e *gin.Engine, db *gorm.DB) {
	e.GET("/api/chats", func(c *gin.Context) {
		// 認証
		isLogin, verifyRes := verify.VerifyToken(c)
		if !isLogin {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		res := Res{}

		backendErr := func() error {
			// 全てのチャットを取得します
			allChats, err := chat_expose.FindByHostID(db, verifyRes.HostID)
			if err != nil {
				return errors.NewError("ホストIDに一致するチャットを取得できません", err)
			}

			// TODO: N+1問題が発生しているため、今後修正
			// 全てのチャットの、それぞれの最新メッセージを取得し、
			// チャットID: メッセージ のmapを作成します。
			messages := map[string]message_expose.Res{}
			for _, chat := range allChats {
				messageRes, err := message_expose.FindLatestByChatID(db, chat.ID)
				if err != nil {
					return errors.NewError("最新のチャットを取得できません", err)
				}

				messages[chat.ID] = messageRes
			}

			// チャットと、その最新メッセージを1つにしてレスポンスを作成します
			chatsRes := make([]ChatRes, 0)
			for _, chat := range allChats {
				r := ChatRes{
					Chat:          shared_api.CastToChatAPIResForHost(chat),
					LatestMessage: shared_api.CastToMessageAPIRes(messages[chat.ID]),
				}
				chatsRes = append(chatsRes, r)
			}

			res.Chats = chatsRes

			return nil
		}()
		if backendErr != nil {
			api_err.Send(c, 500, errors.NewError("バックエンドの処理が失敗しました", backendErr))
			return
		}

		c.JSON(200, res)
	})
}
