package main

import (
	"fmt"
	"io"
	"os"
)

// 3.2.1 読み込みの補助関数

func main() {
	file, _ := os.Open("text.txt")
	defer file.Close()

	buff, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(buff)
}

func readall() {

}
