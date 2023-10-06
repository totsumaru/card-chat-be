package expose

import (
	"time"

	"github.com/totsumaru/card-chat-be/context/message/domain"
)

// レスポンスです
type Res struct {
	ID      string
	ChatID  string
	FromID  string
	Content struct {
		Kind string
		URL  string
		Text string
	}
	Created time.Time
}

// メッセージをレスポンスに変換します
func CreateRes(m domain.Message) Res {
	res := Res{}
	res.ID = m.ID().String()
	res.ChatID = m.ChatID().String()
	res.FromID = m.FromID().String()
	res.Content.Kind = m.Content().Kind().String()
	res.Content.URL = m.Content().URL().String()
	res.Content.Text = m.Content().Text().String()
	res.Created = m.Created()

	return res
}
