package resend

import (
	"fmt"

	"github.com/totsumaru/card-chat-be/shared/errors"
)

// ゲストのメールアドレスが登録/編集された時に通知します
func SendEmailEdited(newEmail string) error {
	contactURL := "https://forms.gle/Vy38edW46DrXTAq2A"

	text := `
[PatChat]通知用メールアドレスが登録されました。

メッセージが届いた場合、こちらのメールアドレスに通知されます。
※メールアドレスは通知以外に使用されることはありません。

このメールに心当たりがない場合は、恐れ入りますが問い合わせフォームからご連絡ください。

[問い合わせフォーム]
%s
`

	req := sendEmailReq{
		to:      newEmail,
		subject: "[PatChat]通知用メールアドレスが登録されました",
		text:    fmt.Sprintf(text, contactURL),
	}

	if err := sendEmail(req); err != nil {
		return errors.NewError("メールを送信できません", err)
	}

	return nil
}
