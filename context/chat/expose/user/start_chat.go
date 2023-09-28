package user

import (
	"github.com/totsumaru/card-chat-be/context/chat/domain"
	"github.com/totsumaru/card-chat-be/context/chat/domain/guest"
	"github.com/totsumaru/card-chat-be/context/chat/expose"
	"github.com/totsumaru/card-chat-be/context/chat/gateway"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// チャットを開始します
//
// ホストIDと表示名を設定します。
func StartChat(
	tx *gorm.DB,
	chatID, hostID, displayName string,
) (expose.Res, error) {
	res := expose.Res{}

	id, err := domain.RestoreID(chatID)
	if err != nil {
		return res, errors.NewError("IDを復元できませんん", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return res, errors.NewError("Gatewayを作成できません", err)
	}

	hID, err := domain.RestoreID(hostID)
	if err != nil {
		return res, errors.NewError("ホストIDを作成できません", err)
	}

	c, err := gw.FindByIDForUpdate(id)
	if err != nil {
		return res, errors.NewError("IDでチャットを取得できません", err)
	}

	// ホストIDを設定
	if err = c.SetHostID(hID); err != nil {
		return res, errors.NewError("ホストIDを設定できません", err)
	}

	// 表示名を設定
	name, err := guest.NewDisplayName(displayName)
	if err != nil {
		return res, errors.NewError("表示名を作成できません", err)
	}
	g, err := guest.NewGuest(name, c.Guest().Memo(), c.Guest().Email())
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
