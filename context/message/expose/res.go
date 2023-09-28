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
	Content string
	Created time.Time
}

// メッセージをレスポンスに変換します
func CreateRes(m domain.Message) Res {
	return Res{
		ID:      m.ID().String(),
		ChatID:  m.ChatID().String(),
		FromID:  m.FromID().String(),
		Content: m.Content().String(),
		Created: m.Created(),
	}
}
