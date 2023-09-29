package res

// ホストのレスポンスです
type HostRes struct {
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
