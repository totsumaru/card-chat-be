package domain

import "github.com/totsumaru/card-chat-be/shared/errors"

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
