package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var id int64

func generateId(mutex *sync.Mutex) int64 {
	mutex.Lock()
	defer mutex.Unlock()
	return atomic.AddInt64(&id, 1)
}

func generateId2() int64 {
	// アトミックの関数を使えばmutexやcondが不要でコンテキストスイッチが発生せず用途が合えば速度改善になる
	return atomic.AddInt64(&id, 1)
}

func main() {

	var mutex sync.Mutex
	_ = mutex
	for _ = range []int64{1, 2, 3} {
		fmt.Println(generateId2())
	}
}
