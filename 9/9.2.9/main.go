package main

import (
	"os"
	"time"
)

func main() {

	// ファイルモードの変更
	os.Chmod("setting.txt", 0644)

	// ファイルのオーナーを変更
	os.Chown("setting.txt", os.Getuid(), os.Getgid())

	// ファイルの最終アクセス日時と変更日時を変更
	os.Chtimes("setting.txt", time.Now(), time.Now())
}
