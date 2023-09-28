package expose

import (
	"time"

	"github.com/totsumaru/card-chat-be/context/chat/domain"
)

// レスポンスです
type Res struct {
	ID       string
	Passcode string
	HostID   string
	Guest    struct {
		DisplayName string
		Memo        string
		Email       string
	}
	IsRead    bool
	IsClosed  bool
	Timestamp struct {
		Created     time.Time
		Updated     time.Time
		LastMessage time.Time
	}
}

// チャットをレスポンスに変換します
func CreateRes(c domain.Chat) Res {
	res := Res{}
	res.ID = c.ID().String()
	res.Passcode = c.Passcode().String()
	res.HostID = c.HostID().String()
	res.Guest.DisplayName = c.Guest().DisplayName().String()
	res.Guest.Memo = c.Guest().Memo().String()
	res.Guest.Email = c.Guest().Email().String()
	res.IsRead = c.IsRead()
	res.IsClosed = c.IsClosed()
	res.Timestamp.Created = c.Timestamp().Created()
	res.Timestamp.Updated = c.Timestamp().Updated()
	res.Timestamp.LastMessage = c.Timestamp().LastMessage()

	return res
}
