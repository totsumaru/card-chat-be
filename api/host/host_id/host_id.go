package host_id

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/context/host/expose/user"
	shared_api "github.com/totsumaru/card-chat-be/shared/api"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// レスポンスです
type Res struct {
	shared_api.HostRes
}

// ホストの情報を取得します
func Host(e *gin.Engine, db *gorm.DB) {
	e.GET("/api/host/:hostID", func(c *gin.Context) {
		hostID := c.Param("hostID")

		tx := db.Begin()
		if tx.Error != nil {
			log.Println(errors.NewError("トランザクションエラーが発生しました", tx.Error))
			c.JSON(500, "トランザクションエラーが発生しました")
			return
		}

		res := Res{}

		// ホストを取得します
		err := func() error {
			apiRes, err := user.FindByID(tx, hostID)
			if err != nil {
				return errors.NewError("ホストが取得できません", err)
			}

			res.ID = apiRes.ID
			res.Name = apiRes.Name
			res.AvatarURL = apiRes.AvatarURL
			res.Headline = apiRes.Headline
			res.Introduction = apiRes.Introduction
			res.Company.Name = apiRes.Company.Name
			res.Company.Position = apiRes.Company.Position
			res.Company.Tel = apiRes.Company.Tel
			res.Company.Email = apiRes.Company.Email
			res.Company.Website = apiRes.Company.Website

			return nil
		}()
		if err != nil {
			tx.Rollback()
			log.Println(errors.NewError("バックエンドの処理が失敗しました", err))
			c.JSON(500, "エラーが発生しました")
			return
		}

		tx.Commit()

		c.JSON(200, gin.H{
			"hostID": hostID,
		})
	})
}
