package domain

import (
	"encoding/json"
	"time"

	"github.com/totsumaru/card-chat-be/context/host/domain/avatar"
	"github.com/totsumaru/card-chat-be/context/host/domain/company"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"github.com/totsumaru/card-chat-be/shared/now"
)

// ホストです
type Host struct {
	id           id.UUID // supabaseのIDと一致します
	name         Name
	avatar       avatar.Avatar
	headline     Headline
	introduction Introduction
	company      company.Company
	created      time.Time
	updated      time.Time
}

// ホストを作成します
func NewHost(id id.UUID) (Host, error) {
	res := Host{
		id:      id,
		created: now.NowJST(),
		updated: now.NowJST(),
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// ホストを復元します
func RestoreHost(
	id id.UUID,
	name Name,
	avatarURL avatar.Avatar,
	headline Headline,
	introduction Introduction,
	company company.Company,
	created time.Time,
	updated time.Time,
) (Host, error) {
	res := Host{
		id:           id,
		name:         name,
		avatar:       avatarURL,
		headline:     headline,
		introduction: introduction,
		company:      company,
		created:      created,
		updated:      updated,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// ホスト情報を更新します
func (h *Host) UpdateHost(
	name Name,
	avatar avatar.Avatar,
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

// アバターURLを取得します
func (h Host) Avatar() avatar.Avatar {
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
func (h Host) Marshal() ([]byte, error) {
	data := struct {
		ID           id.UUID         `json:"id"`
		Name         Name            `json:"name"`
		Avatar       avatar.Avatar   `json:"avatar"`
		Headline     Headline        `json:"headline"`
		Introduction Introduction    `json:"introduction"`
		Company      company.Company `json:"company"`
		Created      time.Time       `json:"created"`
		Updated      time.Time       `json:"updated"`
	}{
		ID:           h.id,
		Name:         h.name,
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
func (h *Host) Unmarshal(b []byte) error {
	var data struct {
		ID           id.UUID         `json:"id"`
		Name         Name            `json:"name"`
		Avatar       avatar.Avatar   `json:"avatar"`
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
	h.avatar = data.Avatar
	h.headline = data.Headline
	h.introduction = data.Introduction
	h.company = data.Company
	h.created = data.Created
	h.updated = data.Updated

	return nil
}
