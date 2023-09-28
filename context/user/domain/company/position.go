package company

import "github.com/totsumaru/card-chat-be/shared/errors"

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
