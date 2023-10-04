package resend

import (
	"fmt"
	"os"

	"github.com/resendlabs/resend-go"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

const (
	fromEmail = "info@patchat.jp"
)

// リクエストです
type SendEmailReq struct {
	ChatID          string
	ToAddress       string
	Message         string
	FromDisplayName string
}

// メールを送信します
//
// host/guest の通知メールです。
func SendEmail(req SendEmailReq) error {
	fmt.Printf("%+v", req)
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))
	chatURL := fmt.Sprintf(
		"%s/chat/%s",
		os.Getenv("FRONTEND_URL"),
		req.ChatID,
	)

	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("PatChat <%s>", fromEmail),
		To:      []string{req.ToAddress},
		Subject: fmt.Sprintf("[PatChat]%s よりメッセージが届きました", req.FromDisplayName),
		Text:    createContent(req.FromDisplayName, req.Message, chatURL),
		Cc:      []string{},
		Bcc:     []string{},
		ReplyTo: "argate.inc@gmail.com",
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		return errors.NewError("メールの送信に失敗しました", err)
	}

	return nil
}

// ホストに送信するメールを作成します
func createContent(fromDisplayName, content, url string) string {
	hostTmpl := `
## メッセージが届きました

From: %s

内容:
%s

(以下のURLからチャットへ移動できます）
%s
`

	return fmt.Sprintf(
		hostTmpl,
		fromDisplayName,
		truncateString(content, 30),
		url,
	)
}

// 指定の文字数を超えた文字列は`...`に変換します
func truncateString(s string, maxLength int) string {
	runes := []rune(s) // 文字列をルーンスライスに変換して、マルチバイト文字にも対応
	if len(runes) > maxLength {
		return string(runes[:maxLength]) + "..."
	}
	return s
}
