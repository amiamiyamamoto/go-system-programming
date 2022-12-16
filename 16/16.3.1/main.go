package main

import "fmt"

func main() {
	// 既存の配列を参照するスライス
	a := [4]int{1, 2, 3, 4}
	b := a[:]
	// fmt.Println(&b[0], &a[0])
	fmt.Println(&b[0], len(b), cap(b))
	a[0] = 10
	fmt.Println(b[0], a[0])

	// 既存の配列の一部を参照するスライス

	c := a[1:3]
	fmt.Println(&c[0], len(c), cap(c))
	// 0xc00001c048 2 3

	// 何も参照していないスライス
	var d []int
	fmt.Println(len(d), cap(d))
	d = append(d, 1)
	fmt.Println(len(d), cap(d))

	d = append(d, []int{1}...)
	fmt.Println(len(d), cap(d))

	d = append(d, []int{1}...)
	fmt.Println(len(d), cap(d))

	// 初期の配列とリンクされているスライス
	e := []int{1, 2, 3, 4}
	fmt.Println(&e[0], len(e), cap(e))
	fmt.Println(&e[0], &e[1], &e[2], &e[3])

	// サイズを持ったスライスを定義
	f := make([]int, 4)
	fmt.Println(&f[0], len(f), cap(f))

	// サイズと容量を持ったスライスを定義
	g := make([]int, 4, 8)
	fmt.Println(&g[0], len(g), cap(g))
}
