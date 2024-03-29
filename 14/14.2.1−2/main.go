package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		go func() {
			// goroutineが起動する前にループが回りきって全部のタスクがcpackになってしまう
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}
