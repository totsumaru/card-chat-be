package content

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// Valueの型です
type KindValue string

const (
	KindText  KindValue = "text"
	KindImage KindValue = "image"
)

// 内容の種類です
type Kind struct {
	value KindValue
}

// 種類を作成します
func NewKind(value KindValue) (Kind, error) {
	res := Kind{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// テキストを取得します
func (k Kind) String() string {
	return string(k.value)
}

// テキストかどうかを判定します
func (k Kind) IsText() bool {
	return k.value == KindText
}

// 画像かどうかを判定します
func (k Kind) IsImage() bool {
	return k.value == KindImage
}

// 空か判定します
func (k Kind) IsEmpty() bool {
	return k.value == ""
}

func (k Kind) validate() error {
	if k.value == "" {
		return errors.NewError("種類が空です")
	}

	return nil
}

// 構造体からJSONに変換します
func (k Kind) MarshalJSON() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: string(k.value),
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (k *Kind) UnmarshalJSON(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	k.value = KindValue(data.Value)

	return nil
}
