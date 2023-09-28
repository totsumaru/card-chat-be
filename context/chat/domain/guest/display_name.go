package guest

import (
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// 表示名の最大文字数です
const DisplayNameMaxLen = 50

// 表示名です
type DisplayName struct {
	value string
}

// 表示名を算出します
func NewDisplayName(value string) (DisplayName, error) {
	res := DisplayName{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 表示名を取得します
func (d DisplayName) String() string {
	return d.value
}

// 表示名を検証します
func (d DisplayName) validate() error {
	if len(d.value) > DisplayNameMaxLen {
		return errors.NewError("文字数を超えています")
	}

	return nil
}
