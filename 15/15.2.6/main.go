package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func(task string) {
			// ジョブを実行
			// このサンプルでは出力だけしてる
			fmt.Println(task)

			wg.Done()
		}(task) // goroutineごとにtaskのコピーが作られるようにしている
	}
	wg.Wait()
}
