package expose

import (
	"github.com/totsumaru/card-chat-be/context/message/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// チャットIDでメッセージを取得します
//
// メッセージの作成日の降順(最近のチャットが先頭)で取得します
func FindByChatID(tx *gorm.DB, chatID string) ([]Res, error) {
	cID, err := id.RestoreUUID(chatID)
	if err != nil {
		return nil, errors.NewError("IDを復元できません", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return nil, errors.NewError("Gatewayを作成できません", err)
	}

	messages, err := gw.FindByChatID(cID)
	if err != nil {
		return nil, errors.NewError("チャットIDでメッセージを取得できません", err)
	}

	res := make([]Res, 0)
	for _, msg := range messages {
		res = append(res, CreateRes(msg))
	}

	return res, nil
}
