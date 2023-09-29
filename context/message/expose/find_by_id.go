package expose

import (
	"github.com/totsumaru/card-chat-be/context/message/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// IDでメッセージを取得します
func FindByID(tx *gorm.DB, messageID string) (Res, error) {
	empty := Res{}

	mID, err := id.RestoreUUID(messageID)
	if err != nil {
		return empty, errors.NewError("IDを復元できませんん", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return empty, errors.NewError("Gatewayを作成できません", err)
	}

	res, err := gw.FindByID(mID)
	if err != nil {
		return empty, errors.NewError("IDでチャットを取得できません", err)
	}

	return CreateRes(res), nil
}
