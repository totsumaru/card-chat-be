package guest

import "github.com/totsumaru/card-chat-be/shared/errors"

// ゲストの情報です
type Guest struct {
	displayName DisplayName
	memo        Memo
	email       Email
}

// ゲストを作成します
func NewGuest(d DisplayName, m Memo, e Email) (Guest, error) {
	res := Guest{
		displayName: d,
		memo:        m,
		email:       e,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 表示名を取得します
func (g Guest) DisplayName() DisplayName {
	return g.displayName
}

// メモを取得します
func (g Guest) Memo() Memo {
	return g.memo
}

// メールアドレスを取得します
func (g Guest) Email() Email {
	return g.email
}

// ゲストを検証します
func (g Guest) validate() error {
	return nil
}
