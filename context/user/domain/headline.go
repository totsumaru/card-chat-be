package domain

import "github.com/totsumaru/card-chat-be/shared/errors"

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
