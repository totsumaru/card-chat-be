package chat_id

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/cookie"
	"github.com/totsumaru/card-chat-be/api/internal/res"
	"github.com/totsumaru/card-chat-be/api/internal/session"
	"github.com/totsumaru/card-chat-be/context/chat/expose/user"
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
	Chat   res.ChatRes `json:"chat"`
	Status ChatStatus  `json:"status"`
}

// チャットです
//
// 初回は全ての会話を取得し、認証できたらWebsocketでリアルタイム通信を行います。
func GetChat(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID", func(c *gin.Context) {
		chatID := c.Param("chatID")

		// 認証
		isLogin, verifyRes := session.Verify(c)

		tx := db.Begin()
		if tx.Error != nil {
			api_err.Send(c, 500, errors.NewError("Txを開始できません", tx.Error))
			return
		}

		// チャットを取得
		apiChatRes, err := user.FindByID(tx, chatID)
		if err != nil {
			api_err.Send(c, 404, errors.NewError("チャットを取得できません", tx.Error))
			return
		}

		// チャットが開始されているかを確認
		if apiChatRes.HostID == "" {
			// チャットが開始されていない場合
			if isLogin {
				c.JSON(200, Res{
					Status: statusFirstIsLogin,
					Chat:   res.ChatRes{},
				})
				return
			} else {
				c.JSON(200, Res{
					Status: statusFirstNotLogin,
					Chat:   res.ChatRes{},
				})
				return
			}
		} else {
			// チャットが開始されている場合

			// 自分がホストの場合
			if apiChatRes.HostID == verifyRes.HostID {
				c.JSON(200, Res{
					Status: statusHost,
					Chat:   res.ChatResForHost(apiChatRes),
				})

				// TODO: Websocket通信を開始します

				return
			} else {
				// 自分がホストではない(ゲストorビジター)場合

				// cookieのパスコードとチャットのパスコードが一致する場合
				cookiePasscode, err := c.Cookie(cookie.PassKey(apiChatRes.ID))
				if err == nil && cookiePasscode == apiChatRes.Passcode {
					c.JSON(200, Res{
						Status: statusGuest,
						Chat:   res.ChatResForGuest(apiChatRes),
					})
					return
				} else {
					c.JSON(200, Res{
						Status: statusVisitor,
						Chat:   res.ChatRes{},
					})
					return
				}
			}
		}
	})
}