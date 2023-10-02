package domain

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// 自己紹介の文字数上限です
const IntroductionMaxLen = 1000

// 自己紹介です
type Introduction struct {
	value string
}

// 自己紹介を作成します
func NewIntroduction(value string) (Introduction, error) {
	res := Introduction{
		value: value,
	}

	if err := res.validate(); err != nil {
		return Introduction{}, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 文字列で取得します
func (i Introduction) String() string {
	return i.value
}

// 検証します
func (i Introduction) validate() error {
	if len(i.value) > IntroductionMaxLen {
		return errors.NewError("文字数を超過しています")
	}

	return nil
}

// 構造体からJSONに変換します
func (i Introduction) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: i.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (i *Introduction) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	i.value = data.Value

	return nil
}
