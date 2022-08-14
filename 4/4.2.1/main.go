package main

import (
	"fmt"
	"time"
)

// チャネルの使用方法
func send(tasks chan string) {
	// データを送信
	tasks <- "cmake .."
	tasks <- "cmake. --build Debug"

}

func receipt(tasks chan string) {
	// データを受け取り
	task := <-tasks
	fmt.Println(task)

	// データを受け取り＆クローズ判定
	task2, ok := <-tasks
	fmt.Println(task2, ok)

}

func main() {
	// バッファなし
	tasks := make(chan string)

	// バッファ付き
	// tasks2 := make(chan string, 10)

	// データを送信
	go send(tasks)
	// データを受け取り
	go receipt(tasks)

	// データを読み捨てる場合は代入文も不要
	// <-tasks

	time.Sleep(2 * time.Second)
}
