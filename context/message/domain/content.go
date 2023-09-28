package domain

import "github.com/totsumaru/card-chat-be/shared/errors"

const ContentMaxLen = 500

// 内容です
type Content struct {
	value string
}

// 内容を作成します
func NewContent(value string) (Content, error) {
	res := Content{
		value: value,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 内容を取得します
func (c Content) String() string {
	return c.value
}

// 空か判定します
func (c Content) IsEmpty() bool {
	return c.value == ""
}

func (c Content) validate() error {
	if len(c.value) > ContentMaxLen {
		return errors.NewError("文字数を超えています")
	}

	return nil
}
