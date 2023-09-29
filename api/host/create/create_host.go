package create

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	host_expose "github.com/totsumaru/card-chat-be/context/host/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ホストを作成します
func CreateHost(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/host/create", func(c *gin.Context) {
		// 認証
		ok, verifyRes := verify.VerifyToken(c)
		if !ok {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		err := db.Transaction(func(tx *gorm.DB) error {
			_, err := host_expose.CreateHost(tx, verifyRes.HostID)
			if err != nil {
				return errors.NewError("ホストを作成できません", err)
			}

			return nil
		})
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		c.JSON(200, "")
	})
}
