package main

import (
	"fmt"
	"path/filepath"
)

// ファイル名のパターンにマッチするファイルの抽出
func main() {
	fmt.Println(filepath.Match("image-*.png", "image-100.png"))
}
