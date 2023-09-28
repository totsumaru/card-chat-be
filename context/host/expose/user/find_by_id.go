package user

import (
	"github.com/totsumaru/card-chat-be/context/host/expose"
	"github.com/totsumaru/card-chat-be/context/host/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// IDでホストを取得します
func FindByID(tx *gorm.DB, hostID string) (expose.Res, error) {
	res := expose.Res{}

	hID, err := id.RestoreUUID(hostID)
	if err != nil {
		return res, errors.NewError("IDを復元できませんん", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return res, errors.NewError("Gatewayを作成できません", err)
	}

	c, err := gw.FindByID(hID)
	if err != nil {
		return res, errors.NewError("IDでチャットを取得できません", err)
	}

	return expose.CreateRes(c), nil
}
