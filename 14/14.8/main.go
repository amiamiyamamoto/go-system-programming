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

func main() {

	var mutex sync.Mutex
	for _ = range []int64{1, 2, 3} {
		fmt.Println(generateId(&mutex))
	}
}
