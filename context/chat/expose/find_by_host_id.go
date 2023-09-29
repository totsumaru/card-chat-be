package expose

import (
	"github.com/totsumaru/card-chat-be/context/chat/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ホストIDに一致するチャットを全て取得します
//
// 未読 > 最後にメッセージが送信された日時 の優先順位で取得します。
func FindByHostID(tx *gorm.DB, hostID string) ([]Res, error) {
	hID, err := id.RestoreUUID(hostID)
	if err != nil {
		return nil, errors.NewError("ホストIDを復元できません", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return nil, errors.NewError("Gatewayを作成できません", err)
	}

	chats, err := gw.FindByHostID(hID)
	if err != nil {
		return nil, errors.NewError("ホストIDに一致するチャットを取得できません", err)
	}

	res := make([]Res, 0)
	for _, c := range chats {
		res = append(res, CreateRes(c))
	}

	return res, nil
}
