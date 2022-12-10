package main

import (
	"fmt"
	"os"
)

func main() {
	// ページ（メモリを管理する単位）のサイズを取得する
	fmt.Printf("Page Size: %d\n", os.Getpagesize()) // Page Size: 4096
}
