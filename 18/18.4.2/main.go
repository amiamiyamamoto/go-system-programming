package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("waiting S seconds")
	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
