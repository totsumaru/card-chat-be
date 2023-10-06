package expose

import (
	"fmt"
	"mime/multipart"

	"github.com/totsumaru/card-chat-be/context/message/domain"
	"github.com/totsumaru/card-chat-be/context/message/domain/content"
	"github.com/totsumaru/card-chat-be/context/message/gateway"
	"github.com/totsumaru/card-chat-be/shared/cloudflare"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/domain_model/url"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// リクエストです
type Req struct {
	ChatID  string
	FromID  string
	Content struct {
		Kind  content.KindValue
		Text  string
		Image *multipart.FileHeader
	}
}

// メッセージを作成します
func CreateMessage(tx *gorm.DB, req Req) (Res, error) {
	chatID, err := id.RestoreUUID(req.ChatID)
	if err != nil {
		return Res{}, errors.NewError("IDを復元できません", err)
	}

	fromID, err := id.RestoreUUID(req.FromID)
	if err != nil {
		return Res{}, errors.NewError("送信者のIDを復元できません", err)
	}

	messageID, err := id.NewUUID()
	if err != nil {
		return Res{}, errors.NewError("メッセージIDを生成できません", err)
	}

	// Text/ImageによってContentを作成します
	var c content.Content
	switch req.Content.Kind {
	case content.KindText:
		t, err := content.NewText(req.Content.Text)
		if err != nil {
			return Res{}, errors.NewError("テキストを作成できません", err)
		}

		ct, err := content.NewTextContent(t)
		if err != nil {
			return Res{}, errors.NewError("内容を作成できません", err)
		}
		c = ct
	case content.KindImage:
		path := fmt.Sprintf("message/%s/%s", req.ChatID, messageID.String())
		// CloudflareImageに画像をアップロードします
		cloudflareRes, err := cloudflare.UploadImageToCloudflare(path, req.Content.Image)
		if err != nil {
			return Res{}, errors.NewError("ファイルをアップロードできません", err)
		}

		imageURL, err := url.NewURL(cloudflareRes.URL)
		if err != nil {
			return Res{}, errors.NewError("URLを作成できません", err)
		}

		imageContent, err := content.NewImageContent(imageURL)
		if err != nil {
			return Res{}, errors.NewError("内容を作成できません", err)
		}
		c = imageContent
	}

	msg, err := domain.NewMessage(messageID, chatID, fromID, c)
	if err != nil {
		return Res{}, errors.NewError("メッセージを作成できません", err)
	}

	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return Res{}, errors.NewError("Gatewayを作成できません", err)
	}

	if err = gw.Create(msg); err != nil {
		return Res{}, errors.NewError("メッセージのレコードを作成できません", err)
	}

	return CreateRes(msg), nil
}
