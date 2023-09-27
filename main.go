package main

import (
	"fmt"
	"os"
)

func init() {
	if _, err := os.Stat(".env"); err != nil {
		
	}
}

func main() {
	fmt.Println("hello")
}
