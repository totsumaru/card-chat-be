package expose

import (
	"time"

	"github.com/totsumaru/card-chat-be/context/host/domain"
)

// レスポンスです
type Res struct {
	ID           string
	Name         string
	AvatarURL    string
	Headline     string
	Introduction string
	Company      struct {
		Name     string
		Position string
		Tel      string
		Email    string
		Website  string
	}
	Created time.Time
	Updated time.Time
}

// チャットをレスポンスに変換します
func CreateRes(u domain.Host) Res {
	res := Res{}
	res.ID = u.ID().String()
	res.Name = u.Name().String()
	res.AvatarURL = u.Avatar().URL().String()
	res.Headline = u.Headline().String()
	res.Introduction = u.Introduction().String()
	res.Company.Name = u.Company().Name().String()
	res.Company.Position = u.Company().Position().String()
	res.Company.Tel = u.Company().Tel().String()
	res.Company.Email = u.Company().Email().String()
	res.Company.Website = u.Company().Website().String()
	res.Created = u.Created()
	res.Updated = u.Updated()

	return res
}
