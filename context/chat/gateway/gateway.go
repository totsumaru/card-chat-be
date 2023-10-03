package gateway

import (
	"github.com/totsumaru/card-chat-be/context/chat/domain"
	"github.com/totsumaru/card-chat-be/context/chat/domain/guest"
	"github.com/totsumaru/card-chat-be/context/chat/domain/timestamp"
	"github.com/totsumaru/card-chat-be/shared/database"
	"github.com/totsumaru/card-chat-be/shared/domain_model/email"
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

// チャットを新規作成します
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
//
// レコードが存在していない場合はエラーを返します。
func (g Gateway) FindByID(id id.UUID) (domain.Chat, error) {
	res := domain.Chat{}

	dbChat := database.ChatSchema{}
	if err := g.tx.First(&dbChat, "id = ?", id.String()).Error; err != nil {
		return res, errors.NewError("IDでチャットを取得できません", err)
	}

	// DB->ドメインモデルに変換します
	res, err := castToDomainModelChat(dbChat)
	if err != nil {
		return res, errors.NewError("DBをドメインモデルに変換できません", err)
	}

	return res, nil
}

// 指定されたIDのチャットを取得し、そのチャットに対する排他ロックを取得します。
//
// レコードが存在しない場合はエラーを返します。
func (g Gateway) FindByIDForUpdate(id id.UUID) (domain.Chat, error) {
	res := domain.Chat{}

	dbChat := database.ChatSchema{}
	if err := g.tx.Set("gorm:query_option", "FOR UPDATE").First(
		&dbChat, "id = ?", id.String(),
	).Error; err != nil {
		return res, errors.NewError("IDでチャットを取得できません", err)
	}

	// DB->ドメインモデルに変換します
	res, err := castToDomainModelChat(dbChat)
	if err != nil {
		return res, errors.NewError("DBをドメインモデルに変換できません", err)
	}

	return res, nil
}

// ホストIDに一致するチャットを全て取得します
//
// 取得できない場合は空の値を返し、エラーは発生しません。
//
// 取得する順番
//  1. IsRead=false(未読): メッセージが最近のものから降順
//  2. IsRead=true(既読): メッセージが最近のものから降順
func (g Gateway) FindByHostID(hostID id.UUID) ([]domain.Chat, error) {
	var dbChats []database.ChatSchema

	// ORDER BY句を使ってソート条件を指定
	query := g.tx.Where(
		"host_id = ?",
		hostID.String(),
	).Order(`
		CASE WHEN is_read THEN 1 ELSE 0 END,
		last_message DESC
	`)

	// レコードの取得
	if err := query.Find(&dbChats).Error; err != nil {
		// レコードが存在しない場合、空のスライスを返します
		if err == gorm.ErrRecordNotFound {
			return []domain.Chat{}, nil
		}
		return nil, errors.NewError("取得できません", err)
	}

	domainChats := make([]domain.Chat, 0)
	for _, dbChat := range dbChats {
		domainChat, err := castToDomainModelChat(dbChat)
		if err != nil {
			return nil, errors.NewError("DBをドメインモデルに変換できません", err)
		}
		domainChats = append(domainChats, domainChat)
	}

	return domainChats, nil
}

// ドメインモデルをDBの構造体に変換します
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
		Created:     c.Timestamp().Created(),
		Updated:     c.Timestamp().Updated(),
		LastMessage: c.Timestamp().LastMessage(),
	}
}

// DBのチャットからドメインモデルのチャット構造体に変換します
func castToDomainModelChat(dbChat database.ChatSchema) (domain.Chat, error) {
	res := domain.Chat{}

	cID, err := id.RestoreUUID(dbChat.ID)
	if err != nil {
		return res, errors.NewError("IDを復元できません", err)
	}

	passcode, err := domain.RestorePasscode(dbChat.Passcode)
	if err != nil {
		return res, errors.NewError("パスコードを復元できません", err)
	}

	// 空を許容するため、UUIDのバリデーションに引っかからないため、空の構造体を入れる
	hostID := id.UUID{}
	if dbChat.HostID != "" {
		hostID, err = id.RestoreUUID(dbChat.HostID)
		if err != nil {
			return res, errors.NewError("ホストIDを復元できません", err)
		}
	}

	// ゲスト
	displayName, err := guest.NewDisplayName(dbChat.DisplayName)
	if err != nil {
		return res, errors.NewError("表示名を復元できません", err)
	}
	memo, err := guest.NewMemo(dbChat.Memo)
	if err != nil {
		return res, errors.NewError("メモを復元できません", err)
	}
	mail, err := email.NewEmail(dbChat.Email)
	if err != nil {
		return res, errors.NewError("メールアドレスを復元できません", err)
	}

	g, err := guest.NewGuest(displayName, memo, mail)
	if err != nil {
		return res, errors.NewError("ゲストを復元できません", err)
	}

	ts, err := timestamp.RestoreTimestamp(dbChat.Created, dbChat.Updated, dbChat.LastMessage)
	if err != nil {
		return res, errors.NewError("タイムスタンプを復元できません", err)
	}

	res, err = domain.RestoreChat(cID, passcode, hostID, g, dbChat.IsRead, ts)
	if err != nil {
		return res, errors.NewError("チャットを復元できません", err)
	}

	return res, nil
}
