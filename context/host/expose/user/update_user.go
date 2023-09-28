package user

import (
	"github.com/totsumaru/card-chat-be/context/host/domain"
	"github.com/totsumaru/card-chat-be/context/host/domain/company"
	"github.com/totsumaru/card-chat-be/context/host/expose"
	"github.com/totsumaru/card-chat-be/context/host/gateway"
	"github.com/totsumaru/card-chat-be/shared/domain_model/email"
	"github.com/totsumaru/card-chat-be/shared/domain_model/id"
	"github.com/totsumaru/card-chat-be/shared/domain_model/tel"
	"github.com/totsumaru/card-chat-be/shared/domain_model/url"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/gorm"
)

// ホストの情報を更新するリクエストです
type UpdateHostReq struct {
	ID           string
	Name         string
	AvatarURL    string
	Headline     string
	Introduction string
	CompanyName  string
	Position     string
	Tel          string
	Email        string
	Website      string
}

// ホストの情報を更新します
func UpdateHost(tx *gorm.DB, req UpdateHostReq) (expose.Res, error) {
	empty := expose.Res{}

	hostID, err := id.RestoreUUID(req.ID)
	if err != nil {
		return empty, errors.NewError("IDを復元できませんん", err)
	}

	// DBからホストを取得します
	gw, err := gateway.NewGateway(tx)
	if err != nil {
		return empty, errors.NewError("Gatewayを作成できません", err)
	}

	h, err := gw.FindByIDForUpdate(hostID)
	if err != nil {
		return empty, errors.NewError("IDでホストを取得できません", err)
	}

	// 構造体を作成します
	name, err := domain.NewName(req.Name)
	if err != nil {
		return empty, errors.NewError("名前を作成できません", err)
	}
	avatar, err := url.NewURL(req.AvatarURL)
	if err != nil {
		return empty, errors.NewError("アバターURLを作成できません", err)
	}
	headline, err := domain.NewHeadline(req.Headline)
	if err != nil {
		return empty, errors.NewError("ヘッドラインを作成できません", err)
	}
	intro, err := domain.NewIntroduction(req.Introduction)
	if err != nil {
		return empty, errors.NewError("自己紹介を作成できません", err)
	}
	// 会社情報を作成します
	companyName, err := company.NewName(req.CompanyName)
	if err != nil {
		return empty, errors.NewError("会社名を作成できません", err)
	}
	position, err := company.NewPosition(req.Position)
	if err != nil {
		return empty, errors.NewError("ポジションを作成できません", err)
	}
	t, err := tel.NewTel(req.Tel)
	if err != nil {
		return empty, errors.NewError("電話番号を作成できません", err)
	}
	mail, err := email.NewEmail(req.Email)
	if err != nil {
		return empty, errors.NewError("メールアドレスを作成できません", err)
	}
	website, err := url.NewURL(req.Website)
	if err != nil {
		return empty, errors.NewError("websiteを作成できません", err)
	}
	comp, err := company.NewCompany(companyName, position, t, mail, website)
	if err != nil {
		return empty, errors.NewError("会社を作成できません", err)
	}

	if err = h.UpdateHost(name, avatar, headline, intro, comp); err != nil {
		return empty, errors.NewError("ホストを更新できません", err)
	}

	if err = gw.Update(h); err != nil {
		return empty, errors.NewError("DBの更新に失敗しました", err)
	}

	return expose.CreateRes(h), nil
}
