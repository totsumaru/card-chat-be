package avatar

import (
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/domain_model/url"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// アバターです
type Avatar struct {
	cloudflareImageID id.UUID
	url               url.URL
}

// アバターを作成します
func NewAvatar(imageID id.UUID, url url.URL) (Avatar, error) {
	res := Avatar{
		cloudflareImageID: imageID,
		url:               url,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// 画像IDを取得します
func (a Avatar) CloudflareImageID() id.UUID {
	return a.cloudflareImageID
}

// URLを取得します
func (a Avatar) URL() url.URL {
	return a.url
}

// 検証します
func (a Avatar) validate() error {
	return nil
}
