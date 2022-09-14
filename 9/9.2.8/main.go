package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo1, _ := os.Stat("main.go")
	fileInfo2, _ := os.Stat("main.go")

	// ファイルの同一性チェック
	if os.SameFile(fileInfo1, fileInfo2) {
		fmt.Println("同じファイル")
	}
}
