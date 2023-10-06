package content

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/shared/domain_model/url"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// 内容です
type Content struct {
	kind Kind
	url  url.URL // 画像URL(ファイルURL)
	text Text    // 文章
}

// テキストの内容を作成します
func NewTextContent(text Text) (Content, error) {
	res := Content{}

	k, err := NewKind(KindText)
	if err != nil {
		return res, errors.NewError("種類を作成できません", err)
	}

	res.kind = k
	res.url = url.URL{}
	res.text = text

	if err = res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 画像の内容を作成します
func NewImageContent(imageURL url.URL) (Content, error) {
	res := Content{}

	k, err := NewKind(KindImage)
	if err != nil {
		return res, errors.NewError("種類を作成できません", err)
	}

	res.kind = k
	res.url = imageURL
	res.text = Text{}

	if err = res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 種類を取得します
func (c Content) Kind() Kind {
	return c.kind
}

// URLを取得します
func (c Content) URL() url.URL {
	return c.url
}

// テキストを取得します
func (c Content) Text() Text {
	return c.text
}

// 検証します
func (c Content) validate() error {
	if c.kind.IsText() {
		if !c.url.IsEmpty() {
			return errors.NewError("URLに値が入っています")
		}
		if c.text.IsEmpty() {
			return errors.NewError("テキストが空です")
		}
	}
	if c.kind.IsImage() {
		if c.url.IsEmpty() {
			return errors.NewError("URLに値が入っていません")
		}
		if !c.text.IsEmpty() {
			return errors.NewError("テキストに値が入っています")
		}
	}

	return nil
}

// 構造体からJSONに変換します
func (c Content) MarshalJSON() ([]byte, error) {
	data := struct {
		Kind Kind    `json:"kind"`
		Url  url.URL `json:"url"`
		Text Text    `json:"text"`
	}{
		Kind: c.kind,
		Url:  c.url,
		Text: c.text,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (c *Content) UnmarshalJSON(b []byte) error {
	var data struct {
		Kind Kind    `json:"kind"`
		Url  url.URL `json:"url"`
		Text Text    `json:"text"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	c.kind = data.Kind
	c.url = data.Url
	c.text = data.Text

	return nil
}
