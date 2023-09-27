package expose

import (
	"github.com/totsumaru/card-chat-be/chat/domain"
)

// レスポンスです
type Res struct {
	ID       string
	Passcode string
	HostID   string
	Guest    GuestRes
	IsRead   bool
	IsClosed bool
}

// ゲストのレスポンスです
type GuestRes struct {
	DisplayName string
	Memo        string
	Email       string
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

	return res
}
