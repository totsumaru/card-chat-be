package edit

import (
	"net/http"

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

// ホストのプロフィールを編集します
func EditHostProfile(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/host/:hostID/edit", func(c *gin.Context) {
		hostID := c.Param("hostID")

		// 認証
		isLogin, verifyRes := verify.VerifyToken(c)
		if !isLogin || hostID != verifyRes.HostID {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		// Tx
		apiRes := Res{}
		err := db.Transaction(func(tx *gorm.DB) error {
			// ファイルが添付されていない場合はエラーにならない
			avatarImageFile, err := c.FormFile("avatar")
			if err != nil && err != http.ErrMissingFile {
				return errors.NewError("ファイルを取得できません")
			}

			req := host_expose.UpdateHostReq{
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

			hostExposeRes, err := host_expose.UpdateHost(tx, req)
			if err != nil {
				return errors.NewError("ホストの情報を変更できません", err)
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
