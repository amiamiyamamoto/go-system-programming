package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	// GCを呼ぶと追加された要素が消える
	pool.Put("remove 1")
	pool.Put("remove 2")
	runtime.GC()
	fmt.Println(pool.Get())
}
