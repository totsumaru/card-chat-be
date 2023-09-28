package company

import (
	"github.com/totsumaru/card-chat-be/shared/domain_model/email"
	"github.com/totsumaru/card-chat-be/shared/domain_model/tel"
	"github.com/totsumaru/card-chat-be/shared/domain_model/url"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// 会社の情報です
type Company struct {
	name     Name
	position Position
	tel      tel.Tel
	email    email.Email
	website  url.URL
}

// 会社情報を作成します
func NewCompany(
	name Name,
	position Position,
	tel tel.Tel,
	email email.Email,
	website url.URL,
) (Company, error) {
	res := Company{
		name:     name,
		position: position,
		tel:      tel,
		email:    email,
		website:  website,
	}

	if err := res.validate(); err != nil {
		return Company{}, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 会社名を取得します
func (c Company) Name() Name {
	return c.name
}

// 役職を取得します
func (c Company) Position() Position {
	return c.position
}

// 電話番号を取得します
func (c Company) Tel() tel.Tel {
	return c.tel
}

// メールアドレスを取得します
func (c Company) Email() email.Email {
	return c.email
}

// websiteを取得します
func (c Company) Website() url.URL {
	return c.website
}

// 検証します
func (c Company) validate() error {
	return nil
}
