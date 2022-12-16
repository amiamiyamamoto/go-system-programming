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

	// 格納するデータの数が決まっている場合に配列を使用する
	// 固定長配列は要素数まで含めて型
	// 3要素の浮動小数点を座標を扱う構造体とする
	type Vector3D [3]float64

	// R,G,B,Aのそれぞれが0-255の範囲で表現できる色構造体
	type Color [4]uint8

	// 配列の範囲外アクセスもエラーになる
	// fmt.Println(c[10])
	// fmt.Println(c[-1])

	// 変数を使用してアクセスすると実行時にエラーになる
	num := 10
	fmt.Println(c[num])
	// panic: runtime error: index out of range [10] with length 4
}
