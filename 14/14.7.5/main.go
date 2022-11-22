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

	// 標準のrangeは使えないが、ループを行うメソッドが用意されている
	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})

	// キーが登録されていれば過去のデータを、登録されていなければ新しい値を登録するメソッド
	// これは既に登録されているので無視
	smap.LoadOrStore(1, 3)
	// これは保存されていないので挿入
	smap.LoadOrStore(2, 4)
}
