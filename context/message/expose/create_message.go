package expose

import (
	"github.com/totsumaru/card-chat-be/context/message/domain"
	"github.com/totsumaru/card-chat-be/context/message/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// メッセージを作成します
func CreateMessage(
	tx *gorm.DB,
	chatID, fromID, content string,
) (Res, error) {
	empty := Res{}

	cID, err := id.RestoreUUID(chatID)
	if err != nil {
		return empty, errors.NewError("IDを復元できません", err)
	}

	from, err := id.RestoreUUID(fromID)
	if err != nil {
		return empty, errors.NewError("送信者のIDを復元できません", err)
	}

	c, err := domain.NewContent(content)
	if err != nil {
		return empty, errors.NewError("内容を作成できません", err)
	}

	m, err := domain.NewMessage(cID, from, c)
	if err != nil {
		return empty, errors.NewError("メッセージを作成できません", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return empty, errors.NewError("Gatewayを作成できません", err)
	}

	if err = gw.Create(m); err != nil {
		return empty, errors.NewError("メッセージのレコードを作成できません", err)
	}

	return CreateRes(m), nil
}
