package domain

import (
	"time"

	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"github.com/totsumaru/card-chat-be/shared/now"
)

// メッセージです
type Message struct {
	id         id.UUID
	chatID     id.UUID
	fromUserID id.UUID
	content    Content
	created    time.Time
}

// メッセージを作成します
func NewMessage(chatID, fromUserID id.UUID, content Content) (Message, error) {
	mID, err := id.NewUUID()
	if err != nil {
		return Message{}, errors.NewError("IDを作成できません", err)
	}

	res := Message{
		id:         mID,
		chatID:     chatID,
		fromUserID: fromUserID,
		content:    content,
		created:    now.NowJST(),
	}

	if err = res.validate(); err != nil {
		return Message{}, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// メッセージを復元します
func RestoreMessage(
	id, chatID, fromUserID id.UUID,
	content Content,
	created time.Time,
) (Message, error) {
	res := Message{
		id:         id,
		chatID:     chatID,
		fromUserID: fromUserID,
		content:    content,
		created:    created,
	}

	if err := res.validate(); err != nil {
		return Message{}, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// IDを取得します
func (m Message) ID() id.UUID {
	return m.id
}

// チャットIDを取得します
func (m Message) ChatID() id.UUID {
	return m.chatID
}

// 送信者を取得します
func (m Message) FromUserID() id.UUID {
	return m.fromUserID
}

// 送信内容を取得します
func (m Message) Content() Content {
	return m.content
}

// 作成日時を取得します
func (m Message) Created() time.Time {
	return m.created
}

// 検証します
func (m Message) validate() error {
	if m.created.IsZero() {
		return errors.NewError("作成日時が設定されていません")
	}

	return nil
}
