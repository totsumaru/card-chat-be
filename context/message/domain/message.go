package domain

import (
	"encoding/json"
	"time"

	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"github.com/totsumaru/card-chat-be/shared/now"
)

// メッセージです
type Message struct {
	id      id.UUID
	chatID  id.UUID
	fromID  id.UUID // hostID or chatID が入ります
	content Content
	created time.Time
}

// メッセージを作成します
func NewMessage(chatID, fromID id.UUID, content Content) (Message, error) {
	mID, err := id.NewUUID()
	if err != nil {
		return Message{}, errors.NewError("IDを作成できません", err)
	}

	res := Message{
		id:      mID,
		chatID:  chatID,
		fromID:  fromID,
		content: content,
		created: now.NowJST(),
	}

	if err = res.validate(); err != nil {
		return Message{}, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// メッセージを復元します
func RestoreMessage(
	id, chatID, fromID id.UUID,
	content Content,
	created time.Time,
) (Message, error) {
	res := Message{
		id:      id,
		chatID:  chatID,
		fromID:  fromID,
		content: content,
		created: created,
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
func (m Message) FromID() id.UUID {
	return m.fromID
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

// 構造体からJSONに変換します
func (m Message) MarshalJSON() ([]byte, error) {
	data := struct {
		ID      id.UUID   `json:"id"`
		ChatID  id.UUID   `json:"chat_id"`
		FromID  id.UUID   `json:"from_id"`
		Content Content   `json:"content"`
		Created time.Time `json:"created"`
	}{
		ID:      m.id,
		ChatID:  m.chatID,
		FromID:  m.fromID,
		Content: m.content,
		Created: m.created,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (m *Message) UnmarshalJSON(b []byte) error {
	var data struct {
		ID      id.UUID   `json:"id"`
		ChatID  id.UUID   `json:"chat_id"`
		FromID  id.UUID   `json:"from_id"`
		Content Content   `json:"content"`
		Created time.Time `json:"created"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	m.id = data.ID
	m.chatID = data.ChatID
	m.fromID = data.FromID
	m.content = data.Content
	m.created = data.Created

	return nil
}
