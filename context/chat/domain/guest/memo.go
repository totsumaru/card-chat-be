package guest

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// メモの最大文字数です
const MemoMaxLen = 300

// メモです
type Memo struct {
	value string
}

// メモを算出します
func NewMemo(value string) (Memo, error) {
	res := Memo{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// メモを取得します
func (m Memo) String() string {
	return m.value
}

// メモを検証します
func (m Memo) validate() error {
	if len(m.value) > MemoMaxLen {
		return errors.NewError("文字数を超えています")
	}

	return nil
}

// 構造体からJSONに変換します
func (m Memo) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: m.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (m *Memo) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	m.value = data.Value

	return nil
}
