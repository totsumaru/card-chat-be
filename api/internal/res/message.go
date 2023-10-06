package res

import (
	"time"

	message_expose "github.com/totsumaru/card-chat-be/context/message/expose"
)

// メッセージのレスポンスです
type MessageAPIRes struct {
	ID      string `json:"id"`
	ChatID  string `json:"chat_id"`
	FromID  string `json:"from_id"`
	Content struct {
		Kind string `json:"kind"`
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"content"`
	Created time.Time `json:"created"`
}

// バックエンドのレスポンスをAPIのレスポンスに変換します
func CastToMessageAPIRes(backendResMessage message_expose.Res) MessageAPIRes {
	res := MessageAPIRes{}

	res.ID = backendResMessage.ID
	res.ChatID = backendResMessage.ChatID
	res.FromID = backendResMessage.FromID
	res.Content.Kind = backendResMessage.Content.Kind
	res.Content.URL = backendResMessage.Content.URL
	res.Content.Text = backendResMessage.Content.Text
	res.Created = backendResMessage.Created

	return res
}

// 複数のバックエンドのレスポンスをAPIのレスポンスに変換します
func CastToMessagesAPIRes(backendResMessages []message_expose.Res) []MessageAPIRes {
	res := make([]MessageAPIRes, 0)

	for _, msg := range backendResMessages {
		msgRes := MessageAPIRes{}
		msgRes.ID = msg.ID
		msgRes.ChatID = msg.ChatID
		msgRes.FromID = msg.FromID
		msgRes.Content.Kind = msg.Content.Kind
		msgRes.Content.URL = msg.Content.URL
		msgRes.Content.Text = msg.Content.Text
		msgRes.Created = msg.Created

		res = append(res, msgRes)
	}

	return res
}
