package database

import "time"

const (
	DBName = "card-chat.db"
)

// チャットのDBです
type ChatSchema struct {
	ID          string `gorm:"type:uuid;primary_key;"`
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
	ID           string `gorm:"type:uuid;primary_key;"`
	Name         string
	AvatarURL    string
	Headline     string
	Introduction string
	CompanyName  string
	Position     string
	Tel          string
	Email        string
	Website      string
	Created      time.Time
	Updated      time.Time
}

// メッセージのスキーマです
type MessageSchema struct {
	ID         string `gorm:"type:uuid;primary_key;"`
	ChatID     string `gorm:"index"`
	FromUserID string
	Content    string
	Created    time.Time
}
