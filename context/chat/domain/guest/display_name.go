package guest

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// 表示名の最大文字数です
const DisplayNameMaxLen = 50

// 表示名です
type DisplayName struct {
	value string
}

// 表示名を算出します
func NewDisplayName(value string) (DisplayName, error) {
	res := DisplayName{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 表示名を取得します
func (d DisplayName) String() string {
	return d.value
}

// 表示名を検証します
func (d DisplayName) validate() error {
	if len(d.value) > DisplayNameMaxLen {
		return errors.NewError("文字数を超えています")
	}

	return nil
}

// 構造体からJSONに変換します
func (d DisplayName) Marshal() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: d.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (d *DisplayName) Unmarshal(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	d.value = data.Value

	return nil
}
