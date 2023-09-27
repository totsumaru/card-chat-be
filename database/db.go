package database

const (
	DBName = "card-chat.db"
)

// チャットのDBです
type ChatSchema struct {
	ID          string `gorm:"primaryKey"`
	Passcode    string `gorm:"index"`
	HostID      string
	DisplayName string
	Memo        string
	Email       string
	IsRead      bool
	IsClosed    bool
}

// ユーザーのスキーマです
type UserSchema struct {
	ID int `gorm:"primaryKey"`
}
