package shared_api

// ホストのレスポンスです
type HostRes struct {
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
}
