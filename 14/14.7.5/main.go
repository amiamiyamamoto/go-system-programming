package main

import (
	"fmt"
	"sync"
)

func main() {
	// 初期化
	smap := &sync.Map{}

	// なんでも入れられる
	smap.Store("Hello", "world")
	smap.Store(1, 2)

	// 削除
	smap.Delete("test")

	// 取り出し方法
	value, ok := smap.Load("Hello")
	fmt.Printf("key=%v value=%v exists?=%v\n", "hello", value, ok)
}
