package database

// ホストのスキーマです
type Host struct {
	ID   string `gorm:"type:uuid;primary_key;"`
	Data []byte `gorm:"type:jsonb"`
}

// チャットのDBです
type Chat struct {
	ID   string `gorm:"type:uuid;primary_key;"`
	Data []byte `gorm:"type:jsonb"`
}

// メッセージのスキーマです
type Message struct {
	ID     string `gorm:"type:uuid;primary_key;"`
	ChatID string `gorm:"index"`
	Data   []byte `gorm:"type:jsonb"`
}
