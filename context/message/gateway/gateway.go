package gateway

import (
	"encoding/json"

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
	dbMessage, err := castToDBMessage(m)
	if err != nil {
		return errors.NewError("ドメインモデルをDBの構造体に変換できません", err)
	}

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
//
// レコードが存在しない場合はエラーを返します。
func (g Gateway) FindByID(id id.UUID) (domain.Message, error) {
	res := domain.Message{}

	var dbMessage database.Message
	if err := g.tx.First(&dbMessage, "id = ?", id.String()).Error; err != nil {
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
// 取得できない場合は空の値を返し、エラーは発生しません。
//
// createの降順(最新のメッセージが先頭)で取得します。
func (g Gateway) FindByChatID(chatID id.UUID) ([]domain.Message, error) {
	dbMessages := make([]database.Message, 0)

	// Orderメソッドを使ってcreatedの降順でソート
	if err := g.tx.Where(
		"chat_id = ?",
		chatID.String(),
	).Order(
		"(data->>'created')::timestamp desc",
	).Find(&dbMessages).Error; err != nil {
		// レコードが存在しない場合、空のスライスを返します
		if err == gorm.ErrRecordNotFound {
			return []domain.Message{}, nil
		}
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

// チャットIDの最新のメッセージを取得します
//
// 取得できない場合は空の値を返し、エラーは発生しません。
func (g Gateway) FindLastByChatID(chatID id.UUID) (domain.Message, error) {
	empty := domain.Message{}
	var dbMessage database.Message

	// Orderメソッドを使ってcreatedの降順でソートし、Firstメソッドで最新のメッセージを取得
	if err := g.tx.Where(
		"chat_id = ?",
		chatID.String(),
	).Order("created desc").First(&dbMessage).Error; err != nil {
		// レコードが見つからない場合は空のメッセージを返す
		if err == gorm.ErrRecordNotFound {
			return empty, nil
		}
		return empty, errors.NewError("取得できません", err)
	}

	// DBのスキーマをドメインモデルに変換
	domainMessage, err := castToDomainModelMessage(dbMessage)
	if err != nil {
		return empty, errors.NewError("DBをドメインモデルに変換できません", err)
	}

	return domainMessage, nil
}

// ドメインモデルをDBの構造体に変換します
func castToDBMessage(domainMessage domain.Message) (database.Message, error) {
	res := database.Message{}

	b, err := json.Marshal(&domainMessage)
	if err != nil {
		return res, errors.NewError("Marshalに失敗しました", err)
	}

	res.ID = domainMessage.ID().String()
	res.ChatID = domainMessage.ChatID().String()
	res.Data = b

	return res, nil
}

// DBのメッセージからドメインモデルに変換します
func castToDomainModelMessage(dbMessage database.Message) (domain.Message, error) {
	res := domain.Message{}
	if err := json.Unmarshal(dbMessage.Data, &res); err != nil {
		return res, errors.NewError("Unmarshalに失敗しました", err)
	}

	return res, nil
}
