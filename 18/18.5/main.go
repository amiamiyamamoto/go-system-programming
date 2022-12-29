package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC822))

	fmt.Println(now.Format("2006/01/02 03:04:05 MST"))
	// 月			1
	// 日			2
	// 時			3
	// 分			4
	// 秒			5
	// 年			6
	// UTCとの時差   Z07
	// タイムゾーン	  MST
	// 午前/午後	 PM
}
