package host_id

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	shared_api "github.com/totsumaru/card-chat-be/api/internal/res"
	"github.com/totsumaru/card-chat-be/context/host/expose/user"
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
		err := func() error {
			apiRes, err := user.FindByID(tx, hostID)
			if err != nil {
				return errors.NewError("ホストが取得できません", err)
			}

			res.Host.ID = apiRes.ID
			res.Host.Name = apiRes.Name
			res.Host.AvatarURL = apiRes.AvatarURL
			res.Host.Headline = apiRes.Headline
			res.Host.Introduction = apiRes.Introduction
			res.Host.Company.Name = apiRes.Company.Name
			res.Host.Company.Position = apiRes.Company.Position
			res.Host.Company.Tel = apiRes.Company.Tel
			res.Host.Company.Email = apiRes.Company.Email
			res.Host.Company.Website = apiRes.Company.Website

			return nil
		}()
		if err != nil {
			tx.Rollback()
			api_err.Send(c, 500, errors.NewError("バックエンドの処理が失敗しました", err))
			return
		}

		tx.Commit()

		c.JSON(200, res)
	})
}
