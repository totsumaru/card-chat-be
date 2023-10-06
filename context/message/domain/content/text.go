package content

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

const TextMaxLen = 500

// テキストです
type Text struct {
	value string
}

// テキストを作成します
func NewText(value string) (Text, error) {
	res := Text{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// テキストを取得します
func (t Text) String() string {
	return t.value
}

// 空か判定します
func (t Text) IsEmpty() bool {
	return t.value == ""
}

// 検証します
func (t Text) validate() error {
	if len(t.value) > TextMaxLen {
		return errors.NewError("文字数を超えています")
	}

	return nil
}

// 構造体からJSONに変換します
func (t Text) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: t.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (t *Text) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	t.value = data.Value

	return nil
}
