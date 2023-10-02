package res

import (
	"time"

	chatExpose "github.com/totsumaru/card-chat-be/context/chat/expose"
)

// チャットのレスポンスです
type ChatAPIRes struct {
	ID       string `json:"id"`
	Passcode string `json:"passcode"`
	HostID   string `json:"host_id"`
	Guest    struct {
		DisplayName string `json:"display_name"`
		Memo        string `json:"memo"`
		Email       string `json:"email"`
	} `json:"guest"`
	IsRead      bool      `json:"is_read"`
	IsClosed    bool      `json:"is_closed"`
	LastMessage time.Time `json:"last_message"`
}

// ホスト用のチャットのレスポンスです
//
// ゲストが設定した通知用のEmailはホストには送信しません。
func CastToChatAPIResForHost(backendRes chatExpose.Res) ChatAPIRes {
	res := castToChatAPIRes(backendRes)
	res.Guest.Email = ""

	return res
}

// ゲスト用のチャットのレスポンスです
//
// ホストが設定した`表示名`,`メモ`はゲストには送信しません。
func CastToChatAPIResForGuest(backendRes chatExpose.Res) ChatAPIRes {
	res := castToChatAPIRes(backendRes)
	res.Guest.DisplayName = ""
	res.Guest.Memo = ""

	return res
}

// バックエンドのレスポンスをAPIのレスポンスに変換します
func castToChatAPIRes(backendResChat chatExpose.Res) ChatAPIRes {
	res := ChatAPIRes{}
	res.ID = backendResChat.ID
	res.Passcode = backendResChat.Passcode
	res.HostID = backendResChat.HostID
	res.Guest.DisplayName = backendResChat.Guest.DisplayName
	res.Guest.Memo = backendResChat.Guest.Memo
	res.Guest.Email = backendResChat.Guest.Email
	res.IsRead = backendResChat.IsRead
	res.IsClosed = backendResChat.IsClosed
	res.LastMessage = backendResChat.Timestamp.LastMessage

	return res
}
