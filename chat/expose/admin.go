package expose

import (
	"github.com/totsumaru/card-chat-be/chat/domain"
	"github.com/totsumaru/card-chat-be/chat/gateway"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ===============================
// 管理者のみが操作できる処理です
// ===============================

// チャットを作成します
//
// * idとパスワードは自動で生成されます
func CreateChat(tx *gorm.DB) (Res, error) {
	res := Res{}

	c, err := domain.CreateChat()
	if err != nil {
		return res, errors.NewError("チャットを作成できません", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return res, errors.NewError("Gatewayを作成できません", err)
	}

	if err = gw.Create(c); err != nil {
		return res, errors.NewError("チャットのレコードを作成できません", err)
	}

	return CreateRes(c), nil
}
