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
	fmt.Println(a)
	a := 1
	fmt.Println(a)

}
