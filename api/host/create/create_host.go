package create

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/res"
	"github.com/totsumaru/card-chat-be/api/internal/verify"
	host_expose "github.com/totsumaru/card-chat-be/context/host/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	Host res.HostAPIRes `json:"host"`
}

// ホストを作成します
func CreateHost(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/host/create", func(c *gin.Context) {
		// 認証
		isLogin, verifyRes := verify.VerifyToken(c)
		if !isLogin {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		// Tx
		apiRes := Res{}
		err := db.Transaction(func(tx *gorm.DB) error {
			hostExposeRes, err := host_expose.CreateHost(tx, verifyRes.HostID)
			if err != nil {
				return errors.NewError("ホストを作成できません", err)
			}

			apiRes.Host = res.CastToHostAPIRes(hostExposeRes)

			return nil
		})
		if err != nil {
			api_err.Send(c, 500, errors.NewError("Txエラー", err))
			return
		}

		c.JSON(200, apiRes)
	})
}
