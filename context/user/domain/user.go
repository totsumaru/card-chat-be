package domain

import (
	"time"

	"github.com/totsumaru/card-chat-be/context/user/domain/company"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/domain_model/url"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"github.com/totsumaru/card-chat-be/shared/now"
)

// ユーザーです
type User struct {
	id           id.UUID // supabaseのIDと一致します
	name         Name
	avatarURL    url.URL
	headline     Headline
	introduction Introduction
	company      company.Company
	created      time.Time
	updated      time.Time
}

// ユーザーを作成します
func NewUser(id id.UUID) (User, error) {
	res := User{
		id:      id,
		created: now.NowJST(),
		updated: now.NowJST(),
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// ユーザーを復元します
func RestoreUser(
	id id.UUID,
	name Name,
	avatarURL url.URL,
	headline Headline,
	introduction Introduction,
	company company.Company,
	created time.Time,
	updated time.Time,
) (User, error) {
	res := User{
		id:           id,
		name:         name,
		avatarURL:    avatarURL,
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

// ユーザー情報を更新します
func (u *User) UpdateUser(
	name Name,
	avatarURL url.URL,
	headline Headline,
	introduction Introduction,
	company company.Company,
) error {
	u.name = name
	u.avatarURL = avatarURL
	u.headline = headline
	u.introduction = introduction
	u.company = company
	u.updated = now.NowJST()

	return nil
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

// 作成日時を取得します
func (u User) Created() time.Time {
	return u.created
}

// 更新日時を取得します
func (u User) Updated() time.Time {
	return u.updated
}

// 検証します
func (u User) validate() error {
	if u.created.IsZero() {
		return errors.NewError("作成日時がゼロ値です")
	}

	if u.updated.IsZero() {
		return errors.NewError("更新日時がゼロ値です")
	}

	return nil
}
