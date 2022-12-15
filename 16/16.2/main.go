package main

import "fmt"

func main() {
	// リストを宣言
	var a [4]int

	// リストを生成
	b := [4]int{}

	// リストを生成（初期値付き）
	c := [4]int{0, 1, 2, 3}

	// リストを生成（初期値付き／要素数は自動設定）
	d := [...]int{0, 1, 2, 3}

	fmt.Println(a, b, c, d)

}
