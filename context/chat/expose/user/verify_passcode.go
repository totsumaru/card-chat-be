package user

import "github.com/totsumaru/card-chat-be/context/chat/domain"

// パスコードが正しいかを検証します
func IsValidPasscode(chatID string, passcode string) bool {
	generatedPass := domain.GeneratePasscodeFromUUID(chatID)
	if passcode == generatedPass {
		return true
	}

	return false
}
