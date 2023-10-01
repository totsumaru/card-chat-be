package url

import (
	"net/url"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// URLです
type URL struct {
	value string
}

// URLを作成します
func NewURL(value string) (URL, error) {
	res := URL{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 文字列で取得します
func (u URL) String() string {
	return u.value
}

// 検証します
//
// 空を許容します。
func (u URL) validate() error {
	if u.value == "" {
		return nil
	}

	_, err := url.ParseRequestURI(u.value)
	if err != nil {
		return errors.NewError("URLが不正です")
	}

	return nil
}
