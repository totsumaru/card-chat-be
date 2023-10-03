package database

import "time"

// ホストのスキーマです
type HostSchema struct {
	ID   string `gorm:"type:uuid;primary_key;"`
	Data []byte `gorm:"type:jsonb"`
}

// チャットのDBです
type ChatSchema struct {
	ID   string `gorm:"type:uuid;primary_key;"`
	Data []byte `gorm:"type:jsonb"`
}

// メッセージのスキーマです
type MessageSchema struct {
	ID      string `gorm:"type:uuid;primary_key;"`
	ChatID  string `gorm:"index"`
	FromID  string
	Content string
	Created time.Time
}
