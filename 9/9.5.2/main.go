package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// パスを分割する
func main() {
	dir, name := filepath.Split(os.Getenv("PATH"))
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)

	// ファイルパスのセパレータ文字を取得する
	fmt.Println(string(filepath.Separator))
}
