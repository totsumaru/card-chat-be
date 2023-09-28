package user

import (
	"github.com/totsumaru/card-chat-be/context/message/expose"
	"github.com/totsumaru/card-chat-be/context/message/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// IDでメッセージを取得します
func FindByID(tx *gorm.DB, messageID string) (expose.Res, error) {
	empty := expose.Res{}

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

	return expose.CreateRes(res), nil
}
