package domain

import (
	"github.com/google/uuid"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// IDです
type ID struct {
	value string
}

// IDを作成します
func NewID() (ID, error) {
	res := ID{}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return res, errors.NewError("UUIDの生成に失敗しました", err)
	}

	res.value = newUUID.String()

	return res, nil
}

// IDを復元します
func RestoreID(id string) (ID, error) {
	res := ID{
		value: id,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// IDを取得します
func (i ID) String() string {
	return i.value
}

// IDが存在しているか確認します
func (i ID) IsEmpty() bool {
	return i.value == ""
}

// IDを検証します
func (i ID) validate() error {
	_, err := uuid.Parse(i.value)
	if err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}
