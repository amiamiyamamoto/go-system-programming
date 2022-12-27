package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("waiting 5 second")
	for now := range time.Tick(5 * time.Second) {
		fmt.Println("now:", now)
	}
}
