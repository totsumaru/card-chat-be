package user

import (
	"github.com/totsumaru/card-chat-be/context/user/expose"
	"github.com/totsumaru/card-chat-be/context/user/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// IDでユーザーを取得します
func FindByID(tx *gorm.DB, userID string) (expose.Res, error) {
	res := expose.Res{}

	uID, err := id.RestoreUUID(userID)
	if err != nil {
		return res, errors.NewError("IDを復元できませんん", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return res, errors.NewError("Gatewayを作成できません", err)
	}

	c, err := gw.FindByID(uID)
	if err != nil {
		return res, errors.NewError("IDでチャットを取得できません", err)
	}

	return expose.CreateRes(c), nil
}
