package user

import (
	"github.com/totsumaru/card-chat-be/context/chat/domain"
	"github.com/totsumaru/card-chat-be/context/chat/domain/guest"
	"github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/context/chat/gateway"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ユーザーの通知用Emailを更新します
func UpdateEmail(tx *gorm.DB, chatID, email string) (expose.Res, error) {
	res := expose.Res{}

	id, err := domain.RestoreID(chatID)
	if err != nil {
		return res, errors.NewError("IDを復元できませんん", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return res, errors.NewError("Gatewayを作成できません", err)
	}

	c, err := gw.FindByIDForUpdate(id)
	if err != nil {
		return res, errors.NewError("IDでチャットを取得できません", err)
	}

	// Emailを作成
	m, err := guest.NewEmail(email)
	if err != nil {
		return res, errors.NewError("Emailを作成できません", err)
	}

	g, err := guest.NewGuest(c.Guest().DisplayName(), c.Guest().Memo(), m)
	if err != nil {
		return res, errors.NewError("ゲストを作成できません", err)
	}

	if err = c.UpdateGuest(g); err != nil {
		return res, errors.NewError("ゲストを更新できません", err)
	}

	if err = gw.Update(c); err != nil {
		return res, errors.NewError("DBを更新できません", err)
	}

	return expose.CreateRes(c), nil
}
