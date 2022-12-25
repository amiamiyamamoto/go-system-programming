package main

import "time"

func main() {
	// 時間（time.Duration）の表現
	// 5秒
	5 * time.Second

	// 10ミリ秒
	10 * time.Millisecond

	// 10分30秒
	time.ParseDuration("10m30s")
}
