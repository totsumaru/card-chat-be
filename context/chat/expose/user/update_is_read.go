package user

import (
	"github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/context/chat/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// 既読/未読のフラグを変更します
func UpdateIsRead(tx *gorm.DB, chatID string, isRead bool) (expose.Res, error) {
	res := expose.Res{}

	cID, err := id.RestoreUUID(chatID)
	if err != nil {
		return res, errors.NewError("IDを復元できませんん", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return res, errors.NewError("Gatewayを作成できません", err)
	}

	c, err := gw.FindByIDForUpdate(cID)
	if err != nil {
		return res, errors.NewError("IDでチャットを取得できません", err)
	}

	if err = c.UpdateIsRead(isRead); err != nil {
		return res, errors.NewError("既読フラグを更新できません", err)
	}

	if err = gw.Update(c); err != nil {
		return res, errors.NewError("DBを更新できません", err)
	}

	return expose.CreateRes(c), nil
}
