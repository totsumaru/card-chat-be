package id

import (
	"github.com/google/uuid"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// IDです
type UUID struct {
	value string
}

// IDを作成します
func NewUUID() (UUID, error) {
	res := UUID{}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return res, errors.NewError("UUIDの生成に失敗しました", err)
	}

	res.value = newUUID.String()

	return res, nil
}

// IDを復元します
func RestoreUUID(id string) (UUID, error) {
	res := UUID{
		value: id,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// IDを取得します
func (i UUID) String() string {
	return i.value
}

// IDが存在しているか確認します
func (i UUID) IsEmpty() bool {
	return i.value == ""
}

// IDを検証します
func (i UUID) validate() error {
	_, err := uuid.Parse(i.value)
	if err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}
