package cloudflare

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// CloudflareImageから画像を削除します
func DeleteImageFromCloudflare(cloudflareImageID id.UUID) error {
	url := fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/accounts/%s/images/v1/%s",
		os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		cloudflareImageID.String(),
	)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return errors.NewError("httpリクエストを作成できません", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf(
		"Bearer %s",
		os.Getenv("CLOUDFLARE_IMAGE_TOKEN"),
	))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.NewError("httpリクエストに失敗しました", err)
	}
	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			log.Println(errors.NewError("Closeに失敗しました", err))
		}
	}(res.Body)

	return nil
}
