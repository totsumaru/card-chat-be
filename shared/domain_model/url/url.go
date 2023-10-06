package url

import (
	"encoding/json"
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

// 空かどうか判定します
func (u URL) IsEmpty() bool {
	return u.value == ""
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

// 構造体からJSONに変換します
func (u URL) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: u.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (u *URL) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	u.value = data.Value

	return nil
}
