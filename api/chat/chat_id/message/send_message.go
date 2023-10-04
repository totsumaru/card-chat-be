package message

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/cookie"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	host_expose "github.com/totsumaru/card-chat-be/context/host/expose"
	message_expose "github.com/totsumaru/card-chat-be/context/message/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"github.com/totsumaru/card-chat-be/shared/resend"
	"gorm.io/gorm"
)

// -------------------------------------------------------
// このAPIはレスポンスを返しません。
//
// 一定間隔で`GetChat`のAPIがコールされて
// メッセージは最新状態に同期されるため、ここでは保存だけにします。
// -------------------------------------------------------

// メッセージを送信します
//
// 自分がホストの場合は、fromをホストに、
// ヘッダーのパスコードが一致する場合は、fromをチャットID(ゲスト)にします。
func SendMessage(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/chat/:chatID/message", func(c *gin.Context) {
		chatID := c.Param("chatID")

		content := c.PostForm("content")

		var fromID string

		// 認証をします
		isLogin, verifyRes := verify.VerifyToken(c)

		if isLogin {
			// ホストの場合は、fromIDにホストIDを設定します
			isHost, err := verify.IsHost(db, chatID, verifyRes.HostID)
			if err != nil {
				api_err.Send(c, 500, errors.NewError("ホストの確認ができません", err))
				return
			}
			if isHost {
				fromID = verifyRes.HostID
			}
		} else {
			cookiePasscode, _ := c.Cookie(cookie.PassKey(chatID))
			if chat_expose.IsValidPasscode(chatID, cookiePasscode) {
				// パスコードが正しい場合は、fromIDにチャットIDを設定します
				fromID = chatID
			}
		}

		if fromID == "" {
			api_err.Send(c, 401, errors.NewError("送信者が不明です"))
			return
		}

		// Tx
		var (
			messageExposeRes message_expose.Res
			chatExposeRes    chat_expose.Res
		)
		backendErr := db.Transaction(func(tx *gorm.DB) error {
			var err error
			messageExposeRes, err = message_expose.CreateMessage(tx, chatID, fromID, content)
			if err != nil {
				return errors.NewError("メッセージを作成できません", err)
			}

			// 未読処理を行います
			chatExposeRes, err = chat_expose.UpdateIsRead(tx, chatID, false)
			if err != nil {
				return errors.NewError("未読処理に失敗しました", err)
			}

			return nil
		})
		if backendErr != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", backendErr))
			return
		}

		emailSendErr := func() error {
			var fromName string
			var toAddr string

			hostID := verifyRes.HostID
			if hostID == "" {
				// ゲストが送信している場合は、チャットからホストIDを取得
				hostID = chatExposeRes.HostID
			}
			hostExposeRes, err := host_expose.FindByID(db, hostID)
			if err != nil {
				return errors.NewError("ホストを取得できません", err)
			}

			fromIsGuest := fromID == chatID

			// ゲストが送信した場合
			if fromIsGuest {
				fromName = chatExposeRes.Guest.DisplayName
				if fromName == "" {
					fromName = chatExposeRes.ID
				}
				toAddr = hostExposeRes.Email
			} else {
				// ホストが送信した場合

				// ゲストが通知を許可していない場合は終了します
				if chatExposeRes.Guest.Email == "" {
					return nil
				}

				fromName = hostExposeRes.Name
				toAddr = chatExposeRes.Guest.Email
			}

			// メールを送信します
			sendEmailReq := resend.SendEmailReq{
				ChatID:          chatExposeRes.ID,
				ToAddress:       toAddr,
				Message:         messageExposeRes.Content,
				FromDisplayName: fromName,
			}
			if backendErr = resend.SendEmail(sendEmailReq); backendErr != nil {
				return errors.NewError("メールを送信できません", backendErr)
			}

			return nil
		}()
		// メール送信でエラーが発生した場合は、運営だけに通知
		if emailSendErr != nil {
			log.Println("メールの送信に失敗しました", backendErr)
			return
		}

		c.JSON(200, nil)
	})
}
