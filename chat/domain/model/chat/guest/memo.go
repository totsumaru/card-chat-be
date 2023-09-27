package guest

import (
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// メモの最大文字数です
const MemoMaxLen = 300

// メモです
type Memo struct {
	value string
}

// メモを算出します
func NewMemo(value string) (Memo, error) {
	res := Memo{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// メモを取得します
func (m Memo) String() string {
	return m.value
}

// メモを検証します
func (m Memo) validate() error {
	if len(m.value) > MemoMaxLen {
		return errors.NewError("文字数を超えています")
	}

	return nil
}
