package cookie

import "fmt"

// パスコードのCookieのkeyです
func PassKey(chatID string) string {
	return fmt.Sprintf("chat_passcode_%s", chatID)
}
