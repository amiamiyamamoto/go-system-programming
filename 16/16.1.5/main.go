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
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

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
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	// goでは変数をヒープに置くかスタックに置くかはコンパイラが判断する
	// メモリがヒープとスタックどちらに確保されているか知りたい場合は
	// ビルド時に -gcflags -m を渡す

	// $ go build -gcflags -m main.go

	// ︙
	//	./main.go:28:13: ... argument does not escape
	// ./main.go:28:13: a escapes to heap
	// ./main.go:29:13: ... argument does not escape
	// ./main.go:30:13: ... argument does not escape
	// ./main.go:30:13: c escapes to heap
	// ./main.go:31:13: ... argument does not escape
	// ./main.go:37:11: make([]int, 4) escapes to heap
	// ./main.go:40:11: make([]int, 4, 16) escapes to heap
	// ./main.go:43:11: make(map[string]int) escapes to heap
	// ./main.go:46:11: make(map[string]int, 100) escapes to heap
	// ./main.go:54:13: ... argument does not escape
	// ./main.go:54:13: a escapes to heap
	// ./main.go:55:13: ... argument does not escape
	// ./main.go:55:13: b escapes to heap
	// ./main.go:56:13: ... argument does not escape
	// ./main.go:56:13: c escapes to heap
	// ./main.go:57:13: ... argument does not escape
	// ./main.go:58:13: ... argument does not escape
	// ./main.go:59:13: ... argument does not escape
	// ./main.go:60:13: ... argument does not escape
	// ./main.go:14:20: new(Struct) escapes to heap

	// スタックのほうが高速なため、基本はスタックを選択する
	// 外部の関数に渡したり、返り値で使おうとすると、宣言した関数のスコープよりも
	// 寿命が長くなる可能性があるため、ヒープに逃がす(escape)

	// データを変更しない、参照のみの場合は、ポインタでなく実態のコピーを受け取る関数のほうが
	// パフォーマンスが上がる場合がある
}
