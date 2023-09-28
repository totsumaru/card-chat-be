package user

import "github.com/totsumaru/card-chat-be/context/chat/domain"

// パスコードが正しいかを検証します
func IsVerifyPasscode(id string, passcode string) bool {
	generatedPass := domain.GeneratePasscodeFromUUID(id)
	if passcode == generatedPass {
		return true
	}

	return false
}
