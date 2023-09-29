package edit

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/session"
	"github.com/totsumaru/card-chat-be/context/host/expose/user"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ホストのプロフィールを編集します
func EditHostProfile(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/host/:hostID/edit", func(c *gin.Context) {
		hostID := c.Param("hostID")

		// 認証
		ok, res := session.Verify(c)
		if !ok || hostID != res.HostID {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		avatarImageFile, apiErr := c.FormFile("avatar")
		if apiErr != nil {
			api_err.Send(c, 500, errors.NewError("画像ファイルを取得できません", apiErr))
			return
		}

		tx := db.Begin()
		if tx.Error != nil {
			api_err.Send(c, 500, errors.NewError("Txを開始できません", tx.Error))
			return
		}

		// バックエンドの処理を行います
		apiErr = func() error {
			req := user.UpdateHostReq{
				ID:           hostID,
				Name:         c.PostForm("name"),
				AvatarFile:   avatarImageFile,
				Headline:     c.PostForm("headline"),
				Introduction: c.PostForm("introduction"),
				CompanyName:  c.PostForm("company_name"),
				Position:     c.PostForm("position"),
				Tel:          c.PostForm("tel"),
				Email:        c.PostForm("email"),
				Website:      c.PostForm("website"),
			}

			_, apiErr = user.UpdateHost(tx, req)
			if apiErr != nil {
				return errors.NewError("ホストの情報を変更できません", apiErr)
			}

			return nil
		}()
		if apiErr != nil {
			tx.Rollback()
			api_err.Send(c, 500, errors.NewError("バックエンドの処理が失敗しました", apiErr))
			return
		}

		tx.Commit()

		c.JSON(200, "")
	})
}
