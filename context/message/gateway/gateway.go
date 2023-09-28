package gateway

import (
	"github.com/totsumaru/card-chat-be/context/message/domain"
	"github.com/totsumaru/card-chat-be/shared/database"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

type Gateway struct {
	tx *gorm.DB
}

// gatewayを作成します
func NewGateway(tx *gorm.DB) (Gateway, error) {
	if tx == nil {
		return Gateway{}, errors.NewError("引数が空です")
	}

	res := Gateway{
		tx: tx,
	}

	return res, nil
}

// メッセージを新規作成します
func (g Gateway) Create(m domain.Message) error {
	dbMessage := castToDBMessage(m)

	// 新しいレコードをデータベースに保存
	result := g.tx.Create(&dbMessage)
	if result.Error != nil {
		return errors.NewError("レコードを保存できませんでした", result.Error)
	}
	// 主キー制約違反を検出（同じIDのレコードが既に存在する場合）
	if result.RowsAffected == 0 {
		return errors.NewError("既存のレコードが存在しています")
	}

	return nil
}

// IDでメッセージを取得します
func (g Gateway) FindByID(id id.UUID) (domain.Message, error) {
	res := domain.Message{}

	var dbMessage database.MessageSchema
	if err := g.tx.First(&dbMessage, "id = ?", id.String()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, errors.NewError("レコードが見つかりません")
		}
		return res, errors.NewError("IDでチャットを取得できません", err)
	}

	// DB->ドメインモデルに変換します
	res, err := castToDomainModelMessage(dbMessage)
	if err != nil {
		return res, errors.NewError("DBをドメインモデルに変換できません", err)
	}

	return res, nil
}

// チャットIDでメッセージを取得します
//
// createの降順(最新のメッセージが先頭)で取得します。
func (g Gateway) FindByChatID(chatID id.UUID) ([]domain.Message, error) {
	var dbMessages []database.MessageSchema

	// Orderメソッドを使ってcreatedの降順でソート
	if err := g.tx.Where(
		"chat_id = ?",
		chatID.String(),
	).Order("created desc").Find(&dbMessages).Error; err != nil {
		return nil, errors.NewError("取得できません", err)
	}

	domainMessages := make([]domain.Message, 0)
	for _, dbMessage := range dbMessages {
		domainMessage, err := castToDomainModelMessage(dbMessage)
		if err != nil {
			return nil, errors.NewError("DBをドメインモデルに変換できません", err)
		}
		domainMessages = append(domainMessages, domainMessage)
	}

	return domainMessages, nil
}

// ドメインモデルをDBの構造体に変換します
func castToDBMessage(m domain.Message) database.MessageSchema {
	return database.MessageSchema{
		ID:         m.ID().String(),
		ChatID:     m.ChatID().String(),
		FromUserID: m.FromUserID().String(),
		Content:    m.Content().String(),
		Created:    m.Created(),
	}
}

// DBのメッセージからドメインモデルに変換します
func castToDomainModelMessage(dbMessage database.MessageSchema) (domain.Message, error) {
	res := domain.Message{}

	mID, err := id.RestoreUUID(dbMessage.ID)
	if err != nil {
		return res, errors.NewError("IDを復元できません", err)
	}

	chatID, err := id.RestoreUUID(dbMessage.ChatID)
	if err != nil {
		return res, errors.NewError("チャットIDを復元できません", err)
	}

	fromUserID, err := id.RestoreUUID(dbMessage.FromUserID)
	if err != nil {
		return res, errors.NewError("送信者のユーザーIDを復元できません", err)
	}

	content, err := domain.NewContent(dbMessage.Content)
	if err != nil {
		return res, errors.NewError("内容を作成できません", err)
	}

	m, err := domain.RestoreMessage(mID, chatID, fromUserID, content, dbMessage.Created)
	if err != nil {
		return res, errors.NewError("メッセージを作成できません", err)
	}

	return m, nil
}
