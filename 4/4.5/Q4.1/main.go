package main

import (
	"fmt"
	"time"
)

func main() {
	second2 := time.After(10 * time.Second)
	<-second2
	fmt.Println("time up")
}
