package database

import "time"

const (
	DBName = "card-chat.db"
)

// チャットのDBです
type ChatSchema struct {
	ID          string `gorm:"primaryKey"`
	Passcode    string `gorm:"index"`
	HostID      string `gorm:"index"`
	DisplayName string
	Memo        string
	Email       string
	IsRead      bool
	IsClosed    bool
	Created     time.Time
	Updated     time.Time
	LastMessage time.Time
}

// ユーザーのスキーマです
type UserSchema struct {
	ID int `gorm:"primaryKey"`
}
