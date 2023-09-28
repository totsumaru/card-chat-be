package user

import (
	"github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/context/chat/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// IDでチャットを取得します
func FindByID(tx *gorm.DB, chatID string) (expose.Res, error) {
	res := expose.Res{}

	cID, err := id.RestoreUUID(chatID)
	if err != nil {
		return res, errors.NewError("IDを復元できませんん", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return res, errors.NewError("Gatewayを作成できません", err)
	}

	c, err := gw.FindByID(cID)
	if err != nil {
		return res, errors.NewError("IDでチャットを取得できません", err)
	}

	return expose.CreateRes(c), nil
}
