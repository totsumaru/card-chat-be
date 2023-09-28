package user

import (
	"github.com/totsumaru/card-chat-be/context/user/domain"
	"github.com/totsumaru/card-chat-be/context/user/expose"
	"github.com/totsumaru/card-chat-be/context/user/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ユーザーを新規作成します
func CreateUser(tx *gorm.DB, supabaseID string) (expose.Res, error) {
	empty := expose.Res{}

	uID, err := id.RestoreUUID(supabaseID)
	if err != nil {
		return empty, errors.NewError("IDを復元できません", err)
	}

	u, err := domain.NewUser(uID)
	if err != nil {
		return empty, errors.NewError("ユーザーを作成できません", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return empty, errors.NewError("Gatewayを作成できません", err)
	}

	if err = gw.Create(u); err != nil {
		return empty, errors.NewError("ユーザーのレコードを作成できません", err)
	}

	return expose.CreateRes(u), nil
}
