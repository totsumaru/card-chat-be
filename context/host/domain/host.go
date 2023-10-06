package domain

import (
	"encoding/json"
	"time"

	"github.com/totsumaru/card-chat-be/context/host/domain/company"
	"github.com/totsumaru/card-chat-be/shared/domain_model/email"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/domain_model/image"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"github.com/totsumaru/card-chat-be/shared/now"
)

// ホストです
type Host struct {
	id           id.UUID // supabaseのIDと一致します
	name         Name
	email        email.Email
	avatar       image.Image
	headline     Headline
	introduction Introduction
	company      company.Company
	created      time.Time
	updated      time.Time
}

// ホストを作成します
func NewHost(
	id id.UUID,
	email email.Email,
	name Name,
) (Host, error) {
	res := Host{
		id:      id,
		email:   email,
		name:    name,
		created: now.NowJST(),
		updated: now.NowJST(),
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// ホスト情報を更新します
func (h *Host) UpdateHost(
	name Name,
	avatar image.Image,
	headline Headline,
	introduction Introduction,
	company company.Company,
) error {
	h.name = name
	h.avatar = avatar
	h.headline = headline
	h.introduction = introduction
	h.company = company
	h.updated = now.NowJST()

	if err := h.validate(); err != nil {
		return errors.NewError("検証に失敗しました")
	}

	return nil
}

// IDを取得します
func (h Host) ID() id.UUID {
	return h.id
}

// 名前を取得します
func (h Host) Name() Name {
	return h.name
}

// Emailを取得します
func (h Host) Email() email.Email {
	return h.email
}

// アバターを取得します
func (h Host) Avatar() image.Image {
	return h.avatar
}

// ヘッドラインを取得します
func (h Host) Headline() Headline {
	return h.headline
}

// 自己紹介を取得します
func (h Host) Introduction() Introduction {
	return h.introduction
}

// 会社情報を取得します
func (h Host) Company() company.Company {
	return h.company
}

// 作成日時を取得します
func (h Host) Created() time.Time {
	return h.created
}

// 更新日時を取得します
func (h Host) Updated() time.Time {
	return h.updated
}

// 検証します
func (h Host) validate() error {
	if h.created.IsZero() {
		return errors.NewError("作成日時がゼロ値です")
	}

	if h.updated.IsZero() {
		return errors.NewError("更新日時がゼロ値です")
	}

	return nil
}

// 構造体からJSONに変換します
func (h Host) MarshalJSON() ([]byte, error) {
	data := struct {
		ID           id.UUID         `json:"id"`
		Name         Name            `json:"name"`
		Email        email.Email     `json:"email"`
		Avatar       image.Image     `json:"avatar"`
		Headline     Headline        `json:"headline"`
		Introduction Introduction    `json:"introduction"`
		Company      company.Company `json:"company"`
		Created      time.Time       `json:"created"`
		Updated      time.Time       `json:"updated"`
	}{
		ID:           h.id,
		Name:         h.name,
		Email:        h.email,
		Avatar:       h.avatar,
		Headline:     h.headline,
		Introduction: h.introduction,
		Company:      h.company,
		Created:      h.created,
		Updated:      h.updated,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (h *Host) UnmarshalJSON(b []byte) error {
	var data struct {
		ID           id.UUID         `json:"id"`
		Name         Name            `json:"name"`
		Email        email.Email     `json:"email"`
		Avatar       image.Image     `json:"avatar"`
		Headline     Headline        `json:"headline"`
		Introduction Introduction    `json:"introduction"`
		Company      company.Company `json:"company"`
		Created      time.Time       `json:"created"`
		Updated      time.Time       `json:"updated"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	h.id = data.ID
	h.name = data.Name
	h.email = data.Email
	h.avatar = data.Avatar
	h.headline = data.Headline
	h.introduction = data.Introduction
	h.company = data.Company
	h.created = data.Created
	h.updated = data.Updated

	return nil
}
