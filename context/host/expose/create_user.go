package expose

import (
	"github.com/totsumaru/card-chat-be/context/host/domain"
	"github.com/totsumaru/card-chat-be/context/host/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/email"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ホストを新規作成します
func CreateHost(tx *gorm.DB, supabaseID, mailAddress, name string) (Res, error) {
	empty := Res{}

	hostID, err := id.RestoreUUID(supabaseID)
	if err != nil {
		return empty, errors.NewError("IDを復元できません", err)
	}

	addr, err := email.NewEmail(mailAddress)
	if err != nil {
		return empty, errors.NewError("メールアドレスを作成できません", err)
	}

	n, err := domain.NewName(name)
	if err != nil {
		return empty, errors.NewError("名前を作成できません", err)
	}

	h, err := domain.NewHost(hostID, addr, n)
	if err != nil {
		return empty, errors.NewError("ホストを作成できません", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return empty, errors.NewError("Gatewayを作成できません", err)
	}

	if err = gw.Create(h); err != nil {
		return empty, errors.NewError("ホストのレコードを作成できません", err)
	}

	return CreateRes(h), nil
}
