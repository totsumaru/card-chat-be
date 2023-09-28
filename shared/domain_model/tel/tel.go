package tel

import (
	"regexp"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// 電話番号の最大文字列です
const TelMaxLength = 20

// 電話番号です
type Tel struct {
	value string
}

// 電話番号を作成します
func NewTel(value string) (Tel, error) {
	res := Tel{
		value: value,
	}

	if err := res.validate(); err != nil {
		return Tel{}, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 文字列で取得します
func (t Tel) String() string {
	return t.value
}

// 検証します
func (t Tel) validate() error {
	// 空の値を許容します
	if t.value == "" {
		return nil
	}

	if len(t.value) > TelMaxLength {
		return errors.NewError("文字数が超過しています")
	}

	// 数字とハイフンのみを許容する正規表現
	regex := regexp.MustCompile(`^[0-9-]+$`)
	if !regex.MatchString(t.value) {
		return errors.NewError("数字とハイフン以外が入っています")
	}

	return nil
}
