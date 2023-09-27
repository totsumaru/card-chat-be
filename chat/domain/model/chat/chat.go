package chat

import (
	"github.com/totsumaru/card-chat-be/chat/domain/model/chat/guest"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// チャットです
type Chat struct {
	id       ID
	passcode Passcode
	hostID   ID
	guest    guest.Guest
	isRead   bool
}

// チャットを作成します
//
// チャットカード発行時の処理です。
// (運営が実行する操作です)
func CreateChat() (Chat, error) {
	res := Chat{}

	id, err := NewID()
	if err != nil {
		return res, errors.NewError("IDを作成できません", err)
	}

	passcode, err := CalcPasscodeFromUUID(id)
	if err != nil {
		return res, errors.NewError("パスコードを算出できません", err)
	}

	res.id = id
	res.passcode = passcode

	if err = res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// チャットを復元します
func RestoreChat(
	id ID,
	pass Passcode,
	hostID ID,
	g guest.Guest,
	isRead bool,
) (Chat, error) {
	res := Chat{
		id:       id,
		passcode: pass,
		hostID:   hostID,
		guest:    g,
		isRead:   isRead,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// ホストIDを登録します
//
// チャット開始時の処理です。
func (c *Chat) InitHostID(hostID ID) error {
	if hostID.IsEmpty() {
		return errors.NewError("ホストIDがすでに設定されています")
	}

	c.hostID = hostID

	if err := c.validate(); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}

// ゲストの情報を変更します
func (c *Chat) UpdateGuest(g guest.Guest) error {
	c.guest = g

	if err := c.validate(); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}

// 既読/未読処理をします
func (c *Chat) UpdateIsRead(isRead bool) error {
	c.isRead = isRead

	if err := c.validate(); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}

// チャットを検証します
func (c Chat) validate() error {
	if c.id.IsEmpty() {
		return errors.NewError("IDが設定されていません")
	}
	if c.passcode.IsEmpty() {
		return errors.NewError("パスコードが設定されていません")
	}

	return nil
}
