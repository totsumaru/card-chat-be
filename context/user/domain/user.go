package domain

import (
	"github.com/totsumaru/card-chat-be/context/user/domain/company"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/domain_model/url"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// ユーザーです
type User struct {
	id           id.UUID // supabaseのIDと一致します
	name         Name
	avatarURL    url.URL
	headline     Headline
	introduction Introduction
	company      company.Company
}

// ユーザーを作成します
func NewUser(
	id id.UUID,
	name Name,
	avatarURL url.URL,
	headline Headline,
	introduction Introduction,
	company company.Company,
) (User, error) {
	res := User{
		id:           id,
		name:         name,
		avatarURL:    avatarURL,
		headline:     headline,
		introduction: introduction,
		company:      company,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// IDを取得します
func (u User) ID() id.UUID {
	return u.id
}

// 名前を取得します
func (u User) Name() Name {
	return u.name
}

// アバターURLを取得します
func (u User) AvatarURL() url.URL {
	return u.avatarURL
}

// ヘッドラインを取得します
func (u User) Headline() Headline {
	return u.headline
}

// 自己紹介を取得します
func (u User) Introduction() Introduction {
	return u.introduction
}

// 会社情報を取得します
func (u User) Company() company.Company {
	return u.company
}

// 検証します
func (u User) validate() error {
	return nil
}
