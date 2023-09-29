package host_id

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	shared_api "github.com/totsumaru/card-chat-be/api/internal/res"
	host_expose "github.com/totsumaru/card-chat-be/context/host/expose"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	Host shared_api.HostAPIRes `json:"host"`
}

// ホストの情報を取得します
func GetHost(e *gin.Engine, db *gorm.DB) {
	e.GET("/api/host/:hostID", func(c *gin.Context) {
		hostID := c.Param("hostID")

		res := Res{}

		// ホストを取得します
		err := func() error {
			backendHost, err := host_expose.FindByID(db, hostID)
			if err != nil {
				return errors.NewError("ホストが取得できません", err)
			}

			res.Host = shared_api.CastToHostAPIRes(backendHost)

			return nil
		}()
		if err != nil {
			api_err.Send(c, 500, errors.NewError("バックエンドの処理が失敗しました", err))
			return
		}

		c.JSON(200, res)
	})
}
