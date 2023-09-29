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
	Host shared_api.HostRes `json:"host"`
}

// ホストの情報を取得します
func GetHost(e *gin.Engine, db *gorm.DB) {
	e.GET("/api/host/:hostID", func(c *gin.Context) {
		hostID := c.Param("hostID")

		tx := db.Begin()
		if tx.Error != nil {
			api_err.Send(c, 500, errors.NewError("Txを開始できません", tx.Error))
			return
		}

		res := Res{}

		// ホストを取得します
		backendErr := func() error {
			backendHost, err := host_expose.FindByID(tx, hostID)
			if err != nil {
				return errors.NewError("ホストが取得できません", err)
			}

			res.Host = shared_api.CastToAPIHostRes(backendHost)

			return nil
		}()
		if backendErr != nil {
			tx.Rollback()
			api_err.Send(c, 500, errors.NewError("バックエンドの処理が失敗しました", backendErr))
			return
		}

		tx.Commit()

		c.JSON(200, res)
	})
}
