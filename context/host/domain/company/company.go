package company

import (
	"encoding/json"

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

// 構造体からJSONに変換します
func (c Company) Marshal() ([]byte, error) {
	data := struct {
		Name     Name        `json:"name"`
		Position Position    `json:"position"`
		Tel      tel.Tel     `json:"tel"`
		Email    email.Email `json:"email"`
		Website  url.URL     `json:"website"`
	}{
		Name:     c.name,
		Position: c.position,
		Tel:      c.tel,
		Email:    c.email,
		Website:  c.website,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (c *Company) Unmarshal(b []byte) error {
	var data struct {
		Name     Name        `json:"name"`
		Position Position    `json:"position"`
		Tel      tel.Tel     `json:"tel"`
		Email    email.Email `json:"email"`
		Website  url.URL     `json:"website"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	c.name = data.Name
	c.position = data.Position
	c.tel = data.Tel
	c.email = data.Email
	c.website = data.Website

	return nil
}
