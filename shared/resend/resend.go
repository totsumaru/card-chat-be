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
type sendEmailReq struct {
	to      string
	subject string
	text    string
}

// メールを送信します
//
// host/guest の通知メールです。
func sendEmail(req sendEmailReq) error {
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))

	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("PatChat <%s>", fromEmail),
		To:      []string{req.to},
		Subject: req.subject,
		Text:    req.text,
		Cc:      []string{},
		Bcc:     []string{},
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		return errors.NewError("メールの送信に失敗しました", err)
	}

	return nil
}
