package company

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// ポジションの文字数上限です
const PositionMaxLen = 50

// ポジションです
type Position struct {
	value string
}

// ポジションを作成します
func NewPosition(value string) (Position, error) {
	res := Position{
		value: value,
	}

	if err := res.validate(); err != nil {
		return Position{}, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 文字列で取得します
func (p Position) String() string {
	return p.value
}

// 検証します
func (p Position) validate() error {
	if len(p.value) > PositionMaxLen {
		return errors.NewError("文字数を超過しています")
	}

	return nil
}

// 構造体からJSONに変換します
func (p Position) Marshal() ([]byte, error) {
	data := struct {
		Value string `json:"value"`
	}{
		Value: p.value,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (p *Position) Unmarshal(b []byte) error {
	var data struct {
		Value string `json:"value"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	p.value = data.Value

	return nil
}
