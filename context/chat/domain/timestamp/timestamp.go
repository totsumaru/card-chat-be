package timestamp

import (
	"time"

	"github.com/totsumaru/card-chat-be/shared/errors"
	"github.com/totsumaru/card-chat-be/shared/now"
)

// タイムスタンプです
type Timestamp struct {
	created     time.Time
	updated     time.Time
	lastMessage time.Time
}

// タイムスタンプを作成します
//
// チャットが作成された時のみ使用します。
func CreateInitTimestamp() (Timestamp, error) {
	ts := Timestamp{
		created:     now.NowJST(),
		updated:     now.NowJST(),
		lastMessage: time.Time{},
	}

	if err := ts.validate(); err != nil {
		return ts, errors.NewError("検証に失敗しました", err)
	}

	return ts, nil
}

// タイムスタンプを復元します
func RestoreTimestamp(created, updated, lastMessage time.Time) (Timestamp, error) {
	ts := Timestamp{
		created:     created,
		updated:     updated,
		lastMessage: lastMessage,
	}

	if err := ts.validate(); err != nil {
		return ts, errors.NewError("検証に失敗しました", err)
	}

	return ts, nil
}

// 更新日時を更新します
func (t Timestamp) UpdateUpdatedTime() (Timestamp, error) {
	ts := Timestamp{
		created:     t.created,
		updated:     now.NowJST(),
		lastMessage: t.lastMessage,
	}

	if err := ts.validate(); err != nil {
		return ts, errors.NewError("検証に失敗しました", err)
	}

	return ts, nil
}

// メッセージが送られた日時を更新します
func (t Timestamp) UpdateLastMessageAndUpdatedTime() (Timestamp, error) {
	ts := Timestamp{
		created:     t.created,
		updated:     now.NowJST(),
		lastMessage: now.NowJST(),
	}

	if err := ts.validate(); err != nil {
		return ts, errors.NewError("検証に失敗しました", err)
	}

	return ts, nil
}

// 作成日時を取得します
func (t Timestamp) Created() time.Time {
	return t.created
}

// 更新日時を取得します
func (t Timestamp) Updated() time.Time {
	return t.updated
}

// 最新のメッセージの日時を取得します
func (t Timestamp) LastMessage() time.Time {
	return t.lastMessage
}

// 検証します
func (t Timestamp) validate() error {
	if t.created.After(t.updated) {
		return errors.NewError("createdがupdateよりも後になっています")
	}
	if t.created.IsZero() {
		return errors.NewError("createdがゼロ値です")
	}
	if t.updated.IsZero() {
		return errors.NewError("updatedがゼロ値です")
	}

	return nil
}
