package resend

import (
	"fmt"
	"os"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// リクエストです
type SendMessageEmailReq struct {
	ChatID          string
	ToAddress       string
	Message         string
	FromDisplayName string
}

// メッセージの通知を送信します
func SendMessageEmail(req SendMessageEmailReq) error {
	chatURL := fmt.Sprintf(
		"%s/chat/%s",
		os.Getenv("FRONTEND_URL"),
		req.ChatID,
	)

	r := sendEmailReq{
		to:      req.ToAddress,
		subject: fmt.Sprintf("[PatChat]%s よりメッセージが届きました", req.FromDisplayName),
		text:    createContent(req.FromDisplayName, req.Message, chatURL),
	}

	if err := sendEmail(r); err != nil {
		return errors.NewError("メールを送信できません", err)
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
