package res

import (
	"time"

	message_expose "github.com/totsumaru/card-chat-be/context/message/expose"
)

// メッセージのレスポンスです
type MessageAPIRes struct {
	ID      string    `json:"id"`
	ChatID  string    `json:"chat_id"`
	FromID  string    `json:"from_id"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

// バックエンドのレスポンスをAPIのレスポンスに変換します
func CastToMessagesAPIRes(backendResMessages []message_expose.Res) []MessageAPIRes {
	res := make([]MessageAPIRes, 0)

	for _, msg := range backendResMessages {
		msgRes := MessageAPIRes{
			ID:      msg.ID,
			ChatID:  msg.ChatID,
			FromID:  msg.FromID,
			Content: msg.Content,
			Created: msg.Created,
		}
		res = append(res, msgRes)
	}

	return res
}
