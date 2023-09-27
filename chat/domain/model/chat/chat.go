package chat

import "github.com/totsumaru/card-chat-be/chat/domain/model/chat/guest"

// チャットです
type Chat struct {
	id       ID
	passcode Passcode
	hostID   ID
	guest    guest.Guest
	isRead   bool
}
