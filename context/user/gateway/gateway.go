package gateway

import (
	"github.com/totsumaru/card-chat-be/context/user/domain"
	"github.com/totsumaru/card-chat-be/context/user/domain/company"
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

// ユーザーを新規作成します
//
// 同じIDのレコードが存在する場合はエラーを返します。
func (g Gateway) Create(u domain.User) error {
	dbUser := castToDBUser(u)

	// 新しいレコードをデータベースに保存
	result := g.tx.Create(&dbUser)
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
func (g Gateway) Update(u domain.User) error {
	dbUser := castToDBUser(u)

	// IDに基づいてレコードを更新
	result := g.tx.Model(&database.UserSchema{}).Where(
		"id = ?",
		dbUser.ID,
	).Updates(&dbUser)
	if result.Error != nil {
		return errors.NewError("更新できません", result.Error)
	}

	// 主キー制約違反を検出（指定されたIDのレコードが存在しない場合）
	if result.RowsAffected == 0 {
		return errors.NewError("レコードが存在しません")
	}

	return nil
}

// ドメインモデルをDBの構造体に変換します
func castToDBUser(u domain.User) database.UserSchema {
	return database.UserSchema{
		ID:           u.ID().String(),
		Name:         u.Name().String(),
		AvatarURL:    u.AvatarURL().String(),
		Headline:     u.Headline().String(),
		Introduction: u.Introduction().String(),
		CompanyName:  u.Company().Name().String(),
		Position:     u.Company().Position().String(),
		Tel:          u.Company().Tel().String(),
		Email:        u.Company().Email().String(),
		Website:      u.Company().Website().String(),
	}
}

// DBの構造体からドメインモデルに変換します
func castToDomainModel(dbUser database.UserSchema) (domain.User, error) {
	res := domain.User{}

	uID, err := id.RestoreUUID(dbUser.ID)
	if err != nil {
		return res, errors.NewError("IDを復元できません", err)
	}

	name, err := domain.NewName(dbUser.Name)
	if err != nil {
		return res, errors.NewError("名前を作成できません", err)
	}

	avatar, err := url.NewURL(dbUser.AvatarURL)
	if err != nil {
		return res, errors.NewError("アバターURLを作成できません", err)
	}

	headline, err := domain.NewHeadline(dbUser.Headline)
	if err != nil {
		return res, errors.NewError("ヘッドラインを作成できません", err)
	}

	intro, err := domain.NewIntroduction(dbUser.Introduction)
	if err != nil {
		return res, errors.NewError("自己紹介を作成できません", err)
	}

	// 会社情報を作成
	companyName, err := company.NewName(dbUser.CompanyName)
	if err != nil {
		return res, errors.NewError("会社名を作成できません", err)
	}
	position, err := company.NewPosition(dbUser.Position)
	if err != nil {
		return res, errors.NewError("ポジションを作成できません", err)
	}
	t, err := tel.NewTel(dbUser.Tel)
	if err != nil {
		return res, errors.NewError("電話番号を作成できません", err)
	}
	mail, err := email.NewEmail(dbUser.Email)
	if err != nil {
		return res, errors.NewError("メールアドレスを作成できません", err)
	}
	website, err := url.NewURL(dbUser.Website)
	if err != nil {
		return res, errors.NewError("Websiteを作成できません", err)
	}
	comp, err := company.NewCompany(companyName, position, t, mail, website)
	if err != nil {
		return res, errors.NewError("会社情報を作成できません", err)
	}

	res, err = domain.RestoreUser(
		uID, name, avatar, headline, intro, comp, dbUser.Created, dbUser.Updated,
	)
	if err != nil {
		return res, errors.NewError("ユーザーを作成できません", err)
	}

	return res, nil
}
