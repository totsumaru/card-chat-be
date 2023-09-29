package res

import (
	"time"

	chatExpose "github.com/totsumaru/card-chat-be/context/chat/expose"
)

// チャットのレスポンスです
type ChatRes struct {
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
func ChatResForHost(backendRes chatExpose.Res) ChatRes {
	res := castToAPIChatRes(backendRes)
	res.Guest.Email = ""

	return res
}

// ゲスト用のチャットのレスポンスです
//
// ホストが設定した`表示名`,`メモ`はゲストには送信しません。
func ChatResForGuest(backendRes chatExpose.Res) ChatRes {
	res := castToAPIChatRes(backendRes)
	res.Guest.DisplayName = ""
	res.Guest.Memo = ""

	return res
}

// バックエンドのレスポンスをAPIのレスポンスに変換します
func castToAPIChatRes(backendRes chatExpose.Res) ChatRes {
	res := ChatRes{}
	res.ID = backendRes.Passcode
	res.Passcode = backendRes.Passcode
	res.HostID = backendRes.HostID
	res.Guest.DisplayName = backendRes.Guest.DisplayName
	res.Guest.Memo = backendRes.Guest.Memo
	res.Guest.Email = backendRes.Guest.Email
	res.IsRead = backendRes.IsRead
	res.IsClosed = backendRes.IsClosed
	res.LastMessage = backendRes.Timestamp.LastMessage

	return res
}
