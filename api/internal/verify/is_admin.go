package verify

import "github.com/gin-gonic/gin"

// 管理者のメールアドレスです
var adminMailAddress = []string{
	"techstart35@gmail.com",
}

// 管理者であることを検証します
func IsAdmin(c *gin.Context) bool {
	_, res := VerifyToken(c)
	for _, email := range adminMailAddress {
		if res.Email == email {
			return true
		}
	}
	return false
}
