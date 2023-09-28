package email

import (
	"net/mail"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// メールアドレスの最大文字数です
const EmailMaxLen = 100

// メールアドレスです
type Email struct {
	value string
}

// メールアドレスを算出します
func NewEmail(value string) (Email, error) {
	res := Email{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// メールアドレスを取得します
func (e Email) String() string {
	return e.value
}

// メールアドレスを検証します
func (e Email) validate() error {
	if len(e.value) > EmailMaxLen {
		return errors.NewError("文字数を超えています")
	}
	if _, err := mail.ParseAddress(e.value); err != nil {
		return errors.NewError("メールアドレスが不正な形式です")
	}

	return nil
}
