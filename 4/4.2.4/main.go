package main

import (
	"fmt"
	"time"
)

func main() {
	reader := make(chan string, 5)
	exit := make(chan bool, 1)

	go func() {
		reader <- "test1"
		reader <- "test2"
		reader <- "test3"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		exit <- true
	}()

L:
	for {
		select {
		case data := <-reader:
			fmt.Println(data)
		case <-exit:
			// ループを抜ける
			break L
		}
	}

}
