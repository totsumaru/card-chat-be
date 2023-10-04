package verify

import (
	chat_expose "github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ホストかどうかを検証します
func IsHost(db *gorm.DB, chatID, hostID string) (bool, error) {
	// チャットを取得します
	chatRes, err := chat_expose.FindByID(db, chatID)
	if err != nil {
		return false, errors.NewError("チャットを取得できません", err)
	}

	// ホストIDが一致するかを確認します
	if chatRes.HostID == hostID {
		return true, nil
	}

	return false, nil
}
