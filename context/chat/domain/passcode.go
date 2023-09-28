package domain

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
	"regexp"

	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
)

// パスコードです
type Passcode struct {
	value string
}

// パスコードを算出します
func CalcPasscodeFromUUID(chatID id.UUID) (Passcode, error) {
	res := Passcode{
		value: GeneratePasscodeFromUUID(chatID.String()),
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// パスコードを復元します
func RestorePasscode(passcode string) (Passcode, error) {
	res := Passcode{
		value: passcode,
	}

	if err := res.validate(); err != nil {
		return res, errors.NewError("検証に失敗しました", err)
	}

	return res, nil
}

// パスコードを取得します
func (p Passcode) String() string {
	return p.value
}

// パスコードが空か確認します
func (p Passcode) IsEmpty() bool {
	return p.value == ""
}

// パスコードを検証します
func (p Passcode) validate() error {
	matched, err := regexp.MatchString(`^\d{6}$`, p.value)
	if err != nil {
		return errors.NewError("検証に失敗しました", err)
	}
	if !matched {
		return errors.NewError("パスコードが指定の形式ではありません")
	}

	return nil
}

// パスコードを生成します
//
// uuidと秘密鍵から、必ず一意となるパスコードを生成します。
func GeneratePasscodeFromUUID(uuid string) string {
	secretKey := os.Getenv("PASSCODE_SECRET_KEY")

	// HMACを使用してUUIDをハッシュ化
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(uuid))
	hashed := h.Sum(nil)

	// ハッシュ値の最初の4バイトを取得し、それを整数に変換
	var number uint32
	number = binary.BigEndian.Uint32(hashed[:4])

	// 6桁の数字に変換
	sixDigitNumber := number % 1000000

	// 6桁未満の場合も先頭に0をつけて6桁にして返します
	return fmt.Sprintf("%06d", sixDigitNumber)
}
