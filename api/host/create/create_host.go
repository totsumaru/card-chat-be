package create

import (
	"github.com/gin-gonic/gin"
	"github.com/totsumaru/card-chat-be/api/internal/api_err"
	"github.com/totsumaru/card-chat-be/api/internal/session"
	"github.com/totsumaru/card-chat-be/context/host/expose/user"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ホストを作成します
func CreateHost(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/host/create", func(c *gin.Context) {
		// 認証
		ok, res := session.Verify(c)
		if !ok {
			api_err.Send(c, 401, errors.NewError("認証できません"))
			return
		}

		tx := db.Begin()
		if tx.Error != nil {
			api_err.Send(c, 500, errors.NewError("Txを開始できません", tx.Error))
			return
		}

		// バックエンドの処理を実行します
		apiErr := func() error {
			_, err := user.CreateHost(tx, res.HostID)
			if err != nil {
				return errors.NewError("ホストを作成できません", err)
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
