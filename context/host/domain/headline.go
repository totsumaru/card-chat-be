package domain

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// ヘッドラインの文字数上限です
const HeadlineMaxLen = 100

// ヘッドラインです
type Headline struct {
	value string
}

// ヘッドラインを作成します
func NewHeadline(value string) (Headline, error) {
	res := Headline{
		value: value,
	}

	if err := res.validate(); err != nil {
		return Headline{}, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 文字列で取得します
func (h Headline) String() string {
	return h.value
}

// 検証します
func (h Headline) validate() error {
	if len(h.value) > HeadlineMaxLen {
		return errors.NewError("文字数を超過しています")
	}

	return nil
}

// 構造体からJSONに変換します
func (h Headline) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: h.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (h *Headline) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	h.value = data.Value

	return nil
}
