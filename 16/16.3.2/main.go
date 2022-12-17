package main

import "fmt"

func main() {
	// 長さ1、確保された要素2のスライスを作成
	s := make([]int, 1, 2)
	fmt.Println(&s[0], len(s), cap(s))
}
