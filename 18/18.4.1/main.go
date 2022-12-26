package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 時間（time.Duration）の表現
	// 5秒
	second := 5 * time.Second

	// 10ミリ秒
	milsec := 10 * time.Millisecond

	fmt.Println(second, milsec)

	// 10分30秒
	time.ParseDuration("10m30s")

	// 現在時刻
	time.Now()

	// 指定日時を作成
	time.Date(2022, 12, 26, 13, 58, 30, 0, time.Local)

	// フォーマットを指定してパース
	time.Parse(time.Kitchen, "11:30AM")

	// Epochタイムから作成
	time.Unix(1503673200, 0)

	// 時間に関する演算
	// 3時間後の時間
	fmt.Println(time.Now().Add(3 * time.Hour))

	// ファイル更新日時が何日前か知る
	fileInfo, _ := os.Stat("./main.go")
	datecount := time.Since(fileInfo.ModTime())
	// sub :=
	fmt.Printf("%v日前", datecount)
	// println(datecount)

	// 時間を1時間毎に丸める
	fmt.Println(time.Now().Round(time.Hour))
}
