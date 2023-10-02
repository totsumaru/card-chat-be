package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// APIのレスポンスです
type ImageUploadResponse struct {
	Result struct {
		ID       string   `json:"id"`
		Variants []string `json:"variants"`
	} `json:"result"`
	Success bool `json:"success"`
}

// この関数のレスポンスです
type Res struct {
	ImageID string
	URL     string
}

// CloudflareImageに画像をアップロードします
//
// 画像URLを返します。
// 画像が空の場合はアップロードせず、空の文字列""を返します。
func UploadImageToCloudflare(hostID id.UUID, image *multipart.FileHeader) (Res, error) {
	res := Res{}

	if image == nil || image.Size == 0 {
		fmt.Println("imageが入っていません")
		return res, nil
	}

	url := fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/accounts/%s/images/v1",
		os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
	)

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// ファイルをオープン
	file, err := image.Open()
	if err != nil {
		return res, err
	}
	defer func(file multipart.File) {
		if err = file.Close(); err != nil {
			log.Println(errors.NewError("Closeに失敗しました", err))
		}
	}(file)

	// ファイルをmultipartに追加
	fileWriter, err := writer.CreateFormFile("file", hostID.String())
	if err != nil {
		return res, err
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return res, err
	}

	// Multipart writerをクローズしてバウンダリーを書き込む
	if err = writer.Close(); err != nil {
		return res, err
	}

	// HTTPリクエストを作成
	req, err := http.NewRequest(http.MethodPost, url, &buffer)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf(
		"Bearer %s", os.Getenv("CLOUDFLARE_IMAGE_TOKEN"),
	))

	// HTTPクライアントを作成し、リクエストを実行
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			log.Println(errors.NewError("Closeに失敗しました", err))
		}
	}(resp.Body)

	// レスポンスをチェック
	if resp.StatusCode == http.StatusOK {
		var imageUploadResponse ImageUploadResponse
		if err = json.NewDecoder(resp.Body).Decode(&imageUploadResponse); err != nil {
			return res, err
		}
		if imageUploadResponse.Success && len(imageUploadResponse.Result.Variants) > 0 {
			res = Res{
				ImageID: imageUploadResponse.Result.ID,
				URL:     imageUploadResponse.Result.Variants[0],
			}
			fmt.Println(res)
			return res, nil
		}
	}

	return Res{}, errors.NewError("ファイルのアップロードに失敗しました")
}
