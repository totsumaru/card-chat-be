package res

import (
	"time"

	messageExpose "github.com/totsumaru/card-chat-be/context/message/expose"
)

// メッセージのレスポンスです
type MessageRes struct {
	ID      string    `json:"id"`
	ChatID  string    `json:"chat_id"`
	FromID  string    `json:"from_id"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

// バックエンドのレスポンスをAPIのレスポンスに変換します
func CastToAPIMessagesRes(messages []messageExpose.Res) []MessageRes {
	res := make([]MessageRes, 0)

	for _, msg := range messages {
		msgRes := MessageRes{
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
