package expose

import (
	"github.com/totsumaru/card-chat-be/context/chat/domain/guest"
	"github.com/totsumaru/card-chat-be/context/chat/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/email"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ゲストの通知用Emailを更新します
func UpdateEmail(tx *gorm.DB, chatID, mail string) (Res, error) {
	empty := Res{}

	cID, err := id.RestoreUUID(chatID)
	if err != nil {
		return empty, errors.NewError("IDを復元できませんん", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return empty, errors.NewError("Gatewayを作成できません", err)
	}

	c, err := gw.FindByIDForUpdate(cID)
	if err != nil {
		return empty, errors.NewError("IDでチャットを取得できません", err)
	}

	// Emailを作成
	m, err := email.NewEmail(mail)
	if err != nil {
		return empty, errors.NewError("Emailを作成できません", err)
	}

	g, err := guest.NewGuest(c.Guest().DisplayName(), c.Guest().Memo(), m)
	if err != nil {
		return empty, errors.NewError("ゲストを作成できません", err)
	}

	if err = c.UpdateGuest(g); err != nil {
		return empty, errors.NewError("ゲストを更新できません", err)
	}

	if err = gw.Update(c); err != nil {
		return empty, errors.NewError("DBを更新できません", err)
	}

	return CreateRes(c), nil
}
