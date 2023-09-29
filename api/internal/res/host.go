package res

import host_expose "github.com/totsumaru/card-chat-be/context/host/expose"

// ホストのレスポンスです
type HostAPIRes struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	AvatarURL    string `json:"avatar_url"`
	Headline     string `json:"headline"`
	Introduction string `json:"introduction"`
	Company      struct {
		Name     string `json:"name"`
		Position string `json:"position"`
		Tel      string `json:"tel"`
		Email    string `json:"email"`
		Website  string `json:"website"`
	} `json:"company"`
}

// バックエンドのレスポンスをAPIのレスポンスに変換します
func CastToHostAPIRes(backendResHost host_expose.Res) HostAPIRes {
	res := HostAPIRes{}
	res.ID = backendResHost.ID
	res.Name = backendResHost.Name
	res.AvatarURL = backendResHost.AvatarURL
	res.Headline = backendResHost.Headline
	res.Introduction = backendResHost.Introduction
	res.Company.Name = backendResHost.Company.Name
	res.Company.Position = backendResHost.Company.Position
	res.Company.Tel = backendResHost.Company.Tel
	res.Company.Email = backendResHost.Company.Email
	res.Company.Website = backendResHost.Company.Website

	return res
}
