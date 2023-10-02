package domain

import (
	"encoding/json"
	"fmt"

	"github.com/totsumaru/card-chat-be/context/chat/domain/guest"
	"github.com/totsumaru/card-chat-be/context/chat/domain/timestamp"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// チャットです
type Chat struct {
	id        id.UUID
	passcode  Passcode
	hostID    id.UUID
	guest     guest.Guest
	isRead    bool
	isClosed  bool // 使うかどうかは不明
	timestamp timestamp.Timestamp
}

// チャットを作成します
//
// チャットカード発行時の処理です。
// (運営が実行する操作です)
func NewChat() (Chat, error) {
	res := Chat{}

	cID, err := id.NewUUID()
	if err != nil {
		return res, errors.NewError("IDを作成できません", err)
	}

	passcode, err := CalcPasscodeFromUUID(cID)
	if err != nil {
		return res, errors.NewError("パスコードを算出できません", err)
	}

	ts, err := timestamp.NewTimestamp()
	if err != nil {
		return res, errors.NewError("タイムスタンプを作成できません", err)
	}

	res.id = cID
	res.passcode = passcode
	res.timestamp = ts

	if err = res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// チャットを復元します
func RestoreChat(
	id id.UUID,
	pass Passcode,
	hostID id.UUID,
	g guest.Guest,
	isRead bool,
	ts timestamp.Timestamp,
) (Chat, error) {
	res := Chat{
		id:        id,
		passcode:  pass,
		hostID:    hostID,
		guest:     g,
		isRead:    isRead,
		isClosed:  false, // 使用するまでは必ずfalse
		timestamp: ts,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// ホストIDを登録します
//
// チャット開始時の処理です。
func (c *Chat) SetHostID(hostID id.UUID) error {
	fmt.Println("ホストID: ", hostID)
	if !c.hostID.IsEmpty() {
		return errors.NewError("ホストIDがすでに設定されています")
	}

	newTimeStamp, err := c.timestamp.UpdateUpdatedTime()
	if err != nil {
		return errors.NewError("更新日時を更新できません", err)
	}

	c.hostID = hostID
	c.timestamp = newTimeStamp

	if err = c.validate(); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}

// ゲストの情報を変更します
func (c *Chat) UpdateGuest(g guest.Guest) error {
	newTimestamp, err := c.timestamp.UpdateUpdatedTime()
	if err != nil {
		return errors.NewError("更新日時を更新できません", err)
	}

	c.guest = g
	c.timestamp = newTimestamp

	fmt.Println("UpdateGuestのEmail: ", c.guest.Email().String())

	if err = c.validate(); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}

// 既読/未読処理をします
func (c *Chat) UpdateIsRead(isRead bool) error {
	newTimestamp, err := c.timestamp.UpdateLastMessageAndUpdatedTime()
	if err != nil {
		return errors.NewError("更新日時を更新できません", err)
	}

	c.isRead = isRead
	c.timestamp = newTimestamp

	if err = c.validate(); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}

// IDを取得します
func (c Chat) ID() id.UUID {
	return c.id
}

// パスコードを取得します
func (c Chat) Passcode() Passcode {
	return c.passcode
}

// ホストIDを取得します
func (c Chat) HostID() id.UUID {
	return c.hostID
}

// ゲストを取得します
func (c Chat) Guest() guest.Guest {
	return c.guest
}

// 既読フラグを取得します
func (c Chat) IsRead() bool {
	return c.isRead
}

// Closeフラグを取得します
func (c Chat) IsClosed() bool {
	return c.isClosed
}

// タイムスタンプを取得します
func (c Chat) Timestamp() timestamp.Timestamp {
	return c.timestamp
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

// 構造体からJSONに変換します
func (c Chat) Marshal() ([]byte, error) {
	data := struct {
		ID        id.UUID             `json:"id"`
		Passcode  Passcode            `json:"passcode"`
		HostID    id.UUID             `json:"host_id"`
		Guest     guest.Guest         `json:"guest"`
		IsRead    bool                `json:"is_read"`
		IsClosed  bool                `json:"is_closed"` // 使うかどうかは不明
		Timestamp timestamp.Timestamp `json:"timestamp"`
	}{
		ID:        c.id,
		Passcode:  c.passcode,
		HostID:    c.hostID,
		Guest:     c.guest,
		IsRead:    c.isRead,
		IsClosed:  c.isClosed,
		Timestamp: c.timestamp,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (c *Chat) Unmarshal(b []byte) error {
	var data struct {
		ID        id.UUID             `json:"id"`
		Passcode  Passcode            `json:"passcode"`
		HostID    id.UUID             `json:"host_id"`
		Guest     guest.Guest         `json:"guest"`
		IsRead    bool                `json:"is_read"`
		IsClosed  bool                `json:"is_closed"` // 使うかどうかは不明
		Timestamp timestamp.Timestamp `json:"timestamp"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	c.id = data.ID
	c.passcode = data.Passcode
	c.hostID = data.HostID
	c.guest = data.Guest
	c.isRead = data.IsRead
	c.isClosed = data.IsClosed
	c.timestamp = data.Timestamp

	return nil
}
