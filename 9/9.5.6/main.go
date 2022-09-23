package main

import (
	"fmt"
	"path/filepath"
)

// ファイル名のパターンにマッチするファイルの抽出
func main() {
	// パターンと指定したファイル名が同じかどうか調べる
	fmt.Println(filepath.Match("image-*.png", "image-100.png"))

	//ルールに合致するファイル名の一覧を取得する
	files, err := filepath.Glob("./*.png")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}
