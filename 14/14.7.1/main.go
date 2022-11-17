package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	// Lock()/Unlock()をペアで呼び出してロックする
	mutex.Lock()
	id++
	result := id
	mutex.Unlock()
	return result
}

func generateId2(mutex *sync.Mutex) int {
	// 多くの場合は次のように連続して書く
	mutex.Lock()
	defer mutex.Unlock()
	id++
	return id
}

func main() {
	//sync.Mutex構造体の変数宣言
	// 次の宣言をしてもポインタ型になるだけで正常に動作する
	// mutex := new(sync.Mutex)
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
		}()
	}
}
