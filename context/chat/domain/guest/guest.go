package guest

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/shared/domain_model/email"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// ゲストの情報です
type Guest struct {
	displayName DisplayName
	memo        Memo
	email       email.Email
}

// ゲストを作成します
func NewGuest(d DisplayName, m Memo, e email.Email) (Guest, error) {
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
func (g Guest) Email() email.Email {
	return g.email
}

// ゲストを検証します
func (g Guest) validate() error {
	return nil
}

// 構造体からJSONに変換します
func (g Guest) MarshalJSON() ([]byte, error) {
	data := struct {
		DisplayName DisplayName `json:"display_name"`
		Memo        Memo        `json:"memo"`
		Email       email.Email `json:"email"`
	}{
		DisplayName: g.displayName,
		Memo:        g.memo,
		Email:       g.email,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (g *Guest) UnmarshalJSON(b []byte) error {
	var data struct {
		DisplayName DisplayName `json:"display_name"`
		Memo        Memo        `json:"memo"`
		Email       email.Email `json:"email"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	g.displayName = data.DisplayName
	g.memo = data.Memo
	g.email = data.Email

	return nil
}
