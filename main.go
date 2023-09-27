package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// .envが存在している場合は読み込み
	if _, err := os.Stat(".env"); err != nil {
		if err = godotenv.Load(); err != nil {
			panic(err)
		}
	}
}

func main() {
	fmt.Println("hello")
}
