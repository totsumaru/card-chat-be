package expose

import (
	"github.com/totsumaru/card-chat-be/context/chat/domain/guest"
	"github.com/totsumaru/card-chat-be/context/chat/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ゲストの情報を更新します
//
// * 表示名
// * メモ
func UpdateGuestInfo(
	tx *gorm.DB,
	chatID, displayName, memo string,
) (Res, error) {
	res := Res{}

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

	// 表示名を作成
	name, err := guest.NewDisplayName(displayName)
	if err != nil {
		return res, errors.NewError("表示名を作成できません", err)
	}
	// メモを作成
	m, err := guest.NewMemo(memo)
	if err != nil {
		return res, errors.NewError("メモを作成できません", err)
	}
	// ゲストを作成
	g, err := guest.NewGuest(name, m, c.Guest().Email())
	if err != nil {
		return res, errors.NewError("ゲストを作成できません", err)
	}

	if err = c.UpdateGuest(g); err != nil {
		return res, errors.NewError("ゲストを更新できません", err)
	}

	if err = gw.Update(c); err != nil {
		return res, errors.NewError("DBを更新できません", err)
	}

	return CreateRes(c), nil
}
