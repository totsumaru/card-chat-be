package verify

import (
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Verifyのレスポンスです
type Res struct {
	HostID string
	Email  string
}

// セッションを検証します
func VerifyToken(c *gin.Context) (bool, Res) {
	authHeader := c.GetHeader("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
		return false, Res{} // ヘッダーが不正またはトークンが存在しない場合は、空文字列を返します
	}

	tokenString := bearerToken[1]

	secret := os.Getenv("SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return false, Res{} // トークンのパースに失敗した場合、またはトークンが無効な場合は、falseと空のResを返します
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, Res{} // Claimsの型が期待どおりでない場合は、falseと空のResを返します
	}

	expiredAt := int64(claims["exp"].(float64))
	if expiredAt <= time.Now().Unix() {
		return false, Res{} // トークンが有効期限切れの場合は、falseと空のResを返します
	}

	id, ok1 := claims["sub"].(string)
	email, ok2 := claims["email"].(string)
	if !ok1 || !ok2 {
		return false, Res{} // IDまたはEmailの抽出に失敗した場合は、falseと空のResを返します
	}

	// トークンが有効であり、IDとEmailを正常に抽出できた場合は、trueとResを返します
	return true, Res{
		HostID: id,
		Email:  email,
	}
}
