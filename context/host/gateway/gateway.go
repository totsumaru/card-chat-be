package gateway

import (
	"github.com/totsumaru/card-chat-be/context/host/domain"
	"github.com/totsumaru/card-chat-be/context/host/domain/avatar"
	"github.com/totsumaru/card-chat-be/context/host/domain/company"
	"github.com/totsumaru/card-chat-be/shared/database"
	"github.com/totsumaru/card-chat-be/shared/domain_model/email"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/domain_model/tel"
	"github.com/totsumaru/card-chat-be/shared/domain_model/url"
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
	dbHost := castToDBHost(u)

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
	dbHost := castToDBHost(u)

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
func castToDBHost(u domain.Host) database.HostSchema {
	return database.HostSchema{
		ID:            u.ID().String(),
		Name:          u.Name().String(),
		AvatarImageID: u.Avatar().CloudflareImageID().String(),
		AvatarURL:     u.Avatar().URL().String(),
		Headline:      u.Headline().String(),
		Introduction:  u.Introduction().String(),
		CompanyName:   u.Company().Name().String(),
		Position:      u.Company().Position().String(),
		Tel:           u.Company().Tel().String(),
		Email:         u.Company().Email().String(),
		Website:       u.Company().Website().String(),
	}
}

// DBの構造体からドメインモデルに変換します
func castToDomainModel(dbHost database.HostSchema) (domain.Host, error) {
	res := domain.Host{}

	hID, err := id.RestoreUUID(dbHost.ID)
	if err != nil {
		return res, errors.NewError("IDを復元できません", err)
	}

	name, err := domain.NewName(dbHost.Name)
	if err != nil {
		return res, errors.NewError("名前を作成できません", err)
	}

	// アバター
	imageID, err := id.RestoreUUID(dbHost.AvatarImageID)
	if err != nil {
		return res, errors.NewError("画像IDを作成できません", err)
	}
	avatarURL, err := url.NewURL(dbHost.AvatarURL)
	if err != nil {
		return res, errors.NewError("アバターURLを作成できません", err)
	}
	avt, err := avatar.NewAvatar(imageID, avatarURL)
	if err != nil {
		return res, errors.NewError("アバターを作成できません", err)
	}

	headline, err := domain.NewHeadline(dbHost.Headline)
	if err != nil {
		return res, errors.NewError("ヘッドラインを作成できません", err)
	}

	intro, err := domain.NewIntroduction(dbHost.Introduction)
	if err != nil {
		return res, errors.NewError("自己紹介を作成できません", err)
	}

	// 会社情報を作成
	companyName, err := company.NewName(dbHost.CompanyName)
	if err != nil {
		return res, errors.NewError("会社名を作成できません", err)
	}
	position, err := company.NewPosition(dbHost.Position)
	if err != nil {
		return res, errors.NewError("ポジションを作成できません", err)
	}
	t, err := tel.NewTel(dbHost.Tel)
	if err != nil {
		return res, errors.NewError("電話番号を作成できません", err)
	}
	mail, err := email.NewEmail(dbHost.Email)
	if err != nil {
		return res, errors.NewError("メールアドレスを作成できません", err)
	}
	website, err := url.NewURL(dbHost.Website)
	if err != nil {
		return res, errors.NewError("Websiteを作成できません", err)
	}
	comp, err := company.NewCompany(companyName, position, t, mail, website)
	if err != nil {
		return res, errors.NewError("会社情報を作成できません", err)
	}

	res, err = domain.RestoreHost(
		hID, name, avt, headline, intro, comp, dbHost.Created, dbHost.Updated,
	)
	if err != nil {
		return res, errors.NewError("ホストを復元できません", err)
	}

	return res, nil
}
