package main

import "fmt"

type Struct struct{ param string }

// 構造体やプリミティブの初期化の方法　→　メモリを確保することをコンパイラやランタイムに指示している

// プリミティブのインスタンスを定義　※プリミティブ→基本の型
var a int = 10

// 構造体のインスタンスをnewして作成
// 変数にはポインタを保存
var b *Struct = new(Struct)

// 構造体を{}でメンバーの初期値を与えて初期化
// 変数にはインスタンスを保存
var c Struct = Struct{param: "param"}

// 構造体を{}でメンバーの初期値を与えて初期化
// 変数にはポインタを保存
var d *Struct = &Struct{param: "param"}

// 配列やスライスの定義の方法

// 固定長配列を定義
func main() {
	// 固定長配列を定義
	a := [4]int{1, 2, 3, 4}

	// サイズを持ったスライスを定義
	b := make([]int, 4)

	// サイズとキャパシティを持ったスライスを定義
	c := make([]int, 4, 16)

	// マップを定義
	d := make(map[string]int)

	// サイズとキャパシティを持ったスライスを定義
	e := make(map[string]int, 100)

	// バッファなしのチャネル
	f := make(chan string)

	// バッファありのチャネル
	g := make(chan string, 10)
	fmt.Println(a)

}
