package gateway

import (
	"github.com/totsumaru/card-chat-be/chat/domain"
	guest2 "github.com/totsumaru/card-chat-be/chat/domain/guest"
	"github.com/totsumaru/card-chat-be/database"
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

// レコードを新規作成します
//
// 同じIDのレコードが存在する場合はエラーを返します。
func (g Gateway) Create(c domain.Chat) error {
	dbChat := castToDBChat(c)

	// 新しいレコードをデータベースに保存
	result := g.tx.Create(&dbChat)
	if result.Error != nil {
		return errors.NewError("レコードを保存できませんでした", result.Error)
	}
	// 主キー制約違反を検出（同じIDのレコードが既に存在する場合）
	if result.RowsAffected == 0 {
		return errors.NewError("既存のレコードが存在しています")
	}

	return nil
}

// 更新します
func (g Gateway) Update(c domain.Chat) error {
	dbChat := castToDBChat(c)

	// IDに基づいてレコードを更新
	result := g.tx.Model(&database.ChatSchema{}).Where(
		"id = ?",
		dbChat.ID,
	).Updates(&dbChat)
	if result.Error != nil {
		return errors.NewError("更新できません", result.Error)
	}

	// 主キー制約違反を検出（指定されたIDのレコードが存在しない場合）
	if result.RowsAffected == 0 {
		return errors.NewError("レコードが存在しません")
	}

	return nil
}

// IDでチャットを取得します
func (g Gateway) FindByID(id domain.ID) (domain.Chat, error) {
	res := domain.Chat{}

	var dbChat database.ChatSchema
	if err := g.tx.First(&dbChat, "id = ?", id.String()).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, errors.NewError("レコードが見つかりません")
		}
		return res, errors.NewError("IDでチャットを取得できません", err)
	}

	// DB->ドメインモデルに変換します
	res, err := castToDomainModelChat(dbChat)
	if err != nil {
		return res, errors.NewError("DBをドメインモデルに変換できません", err)
	}

	return res, nil
}

// チャット構造体をDBのチャット構造体に変換します
func castToDBChat(c domain.Chat) database.ChatSchema {
	return database.ChatSchema{
		ID:          c.ID().String(),
		Passcode:    c.Passcode().String(),
		HostID:      c.HostID().String(),
		DisplayName: c.Guest().DisplayName().String(),
		Memo:        c.Guest().Memo().String(),
		Email:       c.Guest().Email().String(),
		IsRead:      c.IsRead(),
		IsClosed:    c.IsClosed(),
	}
}

// DBのチャットからドメインモデルのチャット構造体に変換します
func castToDomainModelChat(dbChat database.ChatSchema) (domain.Chat, error) {
	res := domain.Chat{}

	id, err := domain.RestoreID(dbChat.ID)
	if err != nil {
		return res, errors.NewError("IDを復元できません", err)
	}

	passcode, err := domain.RestorePasscode(dbChat.Passcode)
	if err != nil {
		return res, errors.NewError("パスコードを復元できません", err)
	}

	hostID, err := domain.RestoreID(dbChat.HostID)
	if err != nil {
		return res, errors.NewError("ホストIDを復元できません", err)
	}

	// ゲスト
	displayName, err := guest2.NewDisplayName(dbChat.DisplayName)
	if err != nil {
		return res, errors.NewError("表示名を復元できません", err)
	}
	memo, err := guest2.NewMemo(dbChat.Memo)
	if err != nil {
		return res, errors.NewError("メモを復元できません", err)
	}
	email, err := guest2.NewEmail(dbChat.Email)
	if err != nil {
		return res, errors.NewError("メールアドレスを復元できません", err)
	}

	g, err := guest2.NewGuest(displayName, memo, email)
	if err != nil {
		return res, errors.NewError("ゲストを復元できません", err)
	}

	res, err = domain.RestoreChat(id, passcode, hostID, g, dbChat.IsRead, dbChat.IsClosed)
	if err != nil {
		return res, errors.NewError("チャットを復元できません", err)
	}

	return res, nil
}
