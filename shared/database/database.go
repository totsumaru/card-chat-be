package database

import "time"

const (
	DBName = "card-chat.db"
)

// ホストのスキーマです
type HostSchema struct {
	ID   string `gorm:"type:uuid;primary_key;"`
	Data []byte `gorm:"type:jsonb"`
}

// チャットのDBです
type ChatSchema struct {
	ID          string `gorm:"type:uuid;primary_key;"`
	Passcode    string `gorm:"index"`
	HostID      string `gorm:"index"`
	DisplayName string
	Memo        string
	Email       *string
	IsRead      bool
	IsClosed    bool
	Created     time.Time
	Updated     time.Time
	LastMessage time.Time
}

// メッセージのスキーマです
type MessageSchema struct {
	ID      string `gorm:"type:uuid;primary_key;"`
	ChatID  string `gorm:"index"`
	FromID  string
	Content string
	Created time.Time
}
