package gateway

import (
	"encoding/json"

	"github.com/totsumaru/card-chat-be/context/host/domain"
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

// ホストを新規作成します
//
// 同じIDのレコードが存在する場合はエラーを返します。
func (g Gateway) Create(u domain.Host) error {
	dbHost, err := castToDBHost(u)
	if err != nil {
		return errors.NewError("ドメインモデルをDBの構造体に変換できません", err)
	}

	// 新しいレコードをデータベースに保存
	result := g.tx.Create(&dbHost)
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
func (g Gateway) Update(u domain.Host) error {
	dbHost, err := castToDBHost(u)
	if err != nil {
		return errors.NewError("ドメインモデルをDBの構造体に変換できません", err)
	}

	// IDに基づいてレコードを更新
	result := g.tx.Model(&database.HostSchema{}).Where(
		"id = ?",
		dbHost.ID,
	).Updates(&dbHost)
	if result.Error != nil {
		return errors.NewError("更新できません", result.Error)
	}

	// 主キー制約違反を検出（指定されたIDのレコードが存在しない場合）
	if result.RowsAffected == 0 {
		return errors.NewError("レコードが存在しません")
	}

	return nil
}

// IDでホストを取得します
//
// レコードが存在しない場合はエラーを返します。
func (g Gateway) FindByID(id id.UUID) (domain.Host, error) {
	res := domain.Host{}

	var dbHost database.HostSchema
	if err := g.tx.First(&dbHost, "id = ?", id.String()).Error; err != nil {
		return res, errors.NewError("IDでホストを取得できません", err)
	}

	// DB->ドメインモデルに変換します
	res, err := castToDomainModel(dbHost)
	if err != nil {
		return res, errors.NewError("DBをドメインモデルに変換できません", err)
	}

	return res, nil
}

// FOR UPDATEでホストを取得します
//
// レコードが存在しない場合はエラーを返します。
func (g Gateway) FindByIDForUpdate(id id.UUID) (domain.Host, error) {
	res := domain.Host{}

	var dbHost database.HostSchema
	if err := g.tx.Set("gorm:query_option", "FOR UPDATE").First(
		&dbHost, "id = ?", id.String(),
	).Error; err != nil {
		return res, errors.NewError("IDでホストを取得できません", err)
	}

	// DB->ドメインモデルに変換します
	res, err := castToDomainModel(dbHost)
	if err != nil {
		return res, errors.NewError("DBをドメインモデルに変換できません", err)
	}

	return res, nil
}

// ドメインモデルをDBの構造体に変換します
func castToDBHost(domainHost domain.Host) (database.HostSchema, error) {
	res := database.HostSchema{}

	b, err := json.Marshal(&domainHost)
	if err != nil {
		return res, errors.NewError("Marshalに失敗しました", err)
	}

	res.ID = domainHost.ID().String()
	res.Data = b

	return res, nil
}

// DBの構造体からドメインモデルに変換します
func castToDomainModel(dbHost database.HostSchema) (domain.Host, error) {
	res := domain.Host{}
	if err := json.Unmarshal(dbHost.Data, &res); err != nil {
		return res, errors.NewError("Unmarshalに失敗しました", err)
	}

	return res, nil
}
