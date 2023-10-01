package chat_id

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/cookie"
	"github.com/totsumaru/card-chat-be/api/internal/res"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	host_expose "github.com/totsumaru/card-chat-be/context/host/expose"
	message_expose "github.com/totsumaru/card-chat-be/context/message/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// レスポンスのチャットステータスです
type ChatStatus string

const (
	statusHost          ChatStatus = "host"
	statusGuest         ChatStatus = "guest"
	statusVisitor       ChatStatus = "visitor"
	statusFirstIsLogin  ChatStatus = "first-is-login"
	statusFirstNotLogin ChatStatus = "first-not-login"
)

// レスポンスです
type Res struct {
	Status   ChatStatus          `json:"status"`
	Chat     res.ChatAPIRes      `json:"chat"`
	Messages []res.MessageAPIRes `json:"messages"`
	Host     res.HostAPIRes      `json:"host"`
}

// チャットを取得します
func GetChat(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID", func(c *gin.Context) {
		chatID := c.Param("chatID")

		// 認証
		isLogin, verifyRes := verify.VerifyToken(c)

		// チャットを取得
		apiChatRes, err := chat_expose.FindByID(db, chatID)
		if err != nil {
			api_err.Send(c, 500, errors.NewError("チャットを取得できません", err))
			return
		}

		// チャットが開始されているかを確認
		if apiChatRes.HostID == "" {
			// ホスト(現在ログイン中のユーザー)を取得します
			host, err := host_expose.FindByID(db, verifyRes.HostID)
			if err != nil {
				api_err.Send(c, 500, errors.NewError("ホストを取得できません", err))
				return
			}

			// チャットが開始されていない場合
			if isLogin {
				c.JSON(200, Res{
					Status:   statusFirstIsLogin,
					Chat:     res.ChatAPIRes{},
					Messages: make([]res.MessageAPIRes, 0),
					Host:     res.CastToHostAPIRes(host),
				})
				return
			} else {
				c.JSON(200, Res{
					Status:   statusFirstNotLogin,
					Chat:     res.ChatAPIRes{},
					Messages: make([]res.MessageAPIRes, 0),
					Host:     res.HostAPIRes{},
				})
				return
			}
		} else {
			// チャットが開始されている場合

			// 全てのメッセージを取得します
			msgs, err := message_expose.FindByChatID(db, apiChatRes.ID)
			if err != nil {
				api_err.Send(c, 500, errors.NewError("チャットIDでメッセージを取得できません", err))
				return
			}

			// ホストを取得します
			host, err := host_expose.FindByID(db, apiChatRes.HostID)
			if err != nil {
				api_err.Send(c, 500, errors.NewError("ホストを取得できません", err))
				return
			}

			// 自分がホストの場合
			if apiChatRes.HostID == verifyRes.HostID {
				c.JSON(200, Res{
					Status:   statusHost,
					Chat:     res.CastToChatAPIResForHost(apiChatRes),
					Messages: res.CastToMessagesAPIRes(msgs),
					Host:     res.CastToHostAPIRes(host),
				})
				return
			} else {
				// 自分がホストではない(ゲストorビジター)場合

				// cookie or header のパスコードと、チャットのパスコードが一致する場合
				cookiePasscode, err := c.Cookie(cookie.PassKey(apiChatRes.ID))
				if err == nil && (cookiePasscode == apiChatRes.Passcode) {
					c.JSON(200, Res{
						Status:   statusGuest,
						Chat:     res.CastToChatAPIResForGuest(apiChatRes),
						Messages: res.CastToMessagesAPIRes(msgs),
						Host:     res.CastToHostAPIRes(host),
					})
					return
				} else {
					c.JSON(200, Res{
						Status:   statusVisitor,
						Chat:     res.ChatAPIRes{},
						Messages: make([]res.MessageAPIRes, 0),
					})
					return
				}
			}
		}
	})
}
