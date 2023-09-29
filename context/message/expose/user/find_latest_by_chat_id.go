package user

import (
	"github.com/totsumaru/card-chat-be/context/message/expose"
	"github.com/totsumaru/card-chat-be/context/message/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// チャットIDの最新のメッセージを取得します
func FindLatestByChatID(tx *gorm.DB, chatID string) (expose.Res, error) {
	empty := expose.Res{}

	cID, err := id.RestoreUUID(chatID)
	if err != nil {
		return empty, errors.NewError("IDを復元できません", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return empty, errors.NewError("Gatewayを作成できません", err)
	}

	msg, err := gw.FindLatestByChatID(cID)
	if err != nil {
		return empty, errors.NewError("チャットIDで最新のメッセージを取得できません", err)
	}

	return expose.CreateRes(msg), nil
}
