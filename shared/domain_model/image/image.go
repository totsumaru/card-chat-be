package image

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/domain_model/url"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// 画像です
type Image struct {
	cloudflareID id.UUID // cloudflareの画像IDです
	url          url.URL
}

// 画像を作成します
func NewImage(cloudflareID id.UUID, url url.URL) (Image, error) {
	res := Image{
		cloudflareID: cloudflareID,
		url:          url,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 画像IDを取得します
func (a Image) CloudflareID() id.UUID {
	return a.cloudflareID
}

// URLを取得します
func (a Image) URL() url.URL {
	return a.url
}

// 検証します
func (a Image) validate() error {
	return nil
}

// 構造体からJSONに変換します
func (a Image) MarshalJSON() ([]byte, error) {
	data := struct {
		CloudflareID id.UUID `json:"cloudflare_id"`
		URL          url.URL `json:"url"`
	}{
		CloudflareID: a.cloudflareID,
		URL:          a.url,
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewError("Marshalに失敗しました", err)
	}

	return b, nil
}

// JSONから構造体に変換します
func (a *Image) UnmarshalJSON(b []byte) error {
	var data struct {
		CloudflareID id.UUID `json:"cloudflare_id"`
		URL          url.URL `json:"url"`
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	a.cloudflareID = data.CloudflareID
	a.url = data.URL

	return nil
}
