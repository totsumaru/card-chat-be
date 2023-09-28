package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/totsumaru/card-chat-be/shared/database"
	"github.com/totsumaru/card-chat-be/shared/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	// .envが存在している場合は読み込み
	if _, err := os.Stat(".env"); err == nil {
		if err = godotenv.Load(); err != nil {
			panic(err)
		}
	}
}

func main() {
	f := sqlite.Open(database.DBName)
	db, err := gorm.Open(f, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(errors.NewError("DBに接続できません", err))
	}

	// テーブルが存在していない場合のみテーブルを作成します
	// 存在している場合はスキーマを同期します
	if err = db.AutoMigrate(&database.UserSchema{}, &database.ChatSchema{}); err != nil {
		panic(errors.NewError("テーブルのスキーマが一致しません", err))
	}

	fmt.Println("success!!")
}
