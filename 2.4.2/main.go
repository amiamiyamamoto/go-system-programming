package main

import (
	"os"
)

func main() {
	// os.Stdout.Write([]byte("os.Stdout example\n"))
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()
}
