package expose

import (
	"github.com/totsumaru/card-chat-be/context/host/domain"
	"github.com/totsumaru/card-chat-be/context/host/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ホストを新規作成します
func CreateHost(tx *gorm.DB, supabaseID string) (Res, error) {
	empty := Res{}

	hostID, err := id.RestoreUUID(supabaseID)
	if err != nil {
		return empty, errors.NewError("IDを復元できません", err)
	}

	u, err := domain.NewHost(hostID)
	if err != nil {
		return empty, errors.NewError("ホストを作成できません", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return empty, errors.NewError("Gatewayを作成できません", err)
	}

	if err = gw.Create(u); err != nil {
		return empty, errors.NewError("ホストのレコードを作成できません", err)
	}

	return CreateRes(u), nil
}
