package company

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// 会社名の文字数上限です
const CompanyNameMaxLen = 100

// 会社名です
type Name struct {
	value string
}

// 会社名を作成します
func NewName(value string) (Name, error) {
	res := Name{
		value: value,
	}

	if err := res.validate(); err != nil {
		return Name{}, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 文字列で取得します
func (n Name) String() string {
	return n.value
}

// 検証します
func (n Name) validate() error {
	if len(n.value) > CompanyNameMaxLen {
		return errors.NewError("文字数を超過しています")
	}

	return nil
}

// 構造体からJSONに変換します
func (n Name) Marshal() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: n.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (n *Name) Unmarshal(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	n.value = data.Value

	return nil
}
