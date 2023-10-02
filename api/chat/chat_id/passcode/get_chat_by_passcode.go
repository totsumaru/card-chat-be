package passcode

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/res"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	host_expose "github.com/totsumaru/card-chat-be/context/host/expose"
	message_expose "github.com/totsumaru/card-chat-be/context/message/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	Chat     res.ChatAPIRes      `json:"chat"`
	Messages []res.MessageAPIRes `json:"messages"`
	Host     res.HostAPIRes      `json:"host"`
}

// チャットを取得します
func GetChatByPasscode(e *gin.Engine, db *gorm.DB) {
	e.GET("/api/chat/:chatID/passcode", func(c *gin.Context) {
		chatID := c.Param("chatID")
		passcode := c.GetHeader("Passcode")

		// パスコードが正しいか検証
		if !chat_expose.IsValidPasscode(chatID, passcode) {
			api_err.Send(c, 401, errors.NewError("パスコードが一致しません"))
			return
		}

		response := Res{}
		err := func() error {
			// チャットを取得
			apiChatRes, err := chat_expose.FindByID(db, chatID)
			if err != nil {
				return errors.NewError("IDでチャットを取得できません", err)
			}

			if apiChatRes.HostID == "" {
				return errors.NewError("チャットが開始されていません")
			}

			// 全てのメッセージを取得します
			msgs, err := message_expose.FindByChatID(db, apiChatRes.ID)
			if err != nil {
				return errors.NewError("チャットIDでメッセージを取得できません", err)
			}

			// ホストを取得します
			host, err := host_expose.FindByID(db, apiChatRes.HostID)
			if err != nil {
				return errors.NewError("ホストを取得できません", err)
			}

			response.Chat = res.CastToChatAPIResForGuest(apiChatRes)
			response.Messages = res.CastToMessagesAPIRes(msgs)
			response.Host = res.CastToHostAPIRes(host)

			return nil
		}()
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		//isSecure := true
		//if os.Getenv("ENV") == "dev" {
		//	isSecure = false
		//}

		// cookieを設定
		c.SetCookie(
			"hello",  // cookieのkey
			passcode, // cookieのvalue
			8.64e+6,  // 有効期限(100日)
			"/",      // cookieのパス
			"",       // cookieのドメイン(空文字は現在のドメインのみ)
			false,    // cookieがセキュアであるかどうか
			true,     // cookieがhttp専用であるかどうか
		)

		c.JSON(200, response)
	})
}
