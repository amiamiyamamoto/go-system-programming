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

	iocopy(file)

}
func iocopy(file *os.File) {
	io.Copy(os.Stdout, file)
}

func readfull(file *os.File) {
	buff := make([]byte, 4)
	_, err := io.ReadFull(file, buff)
	if err != nil {
		panic(err)
	}
	fmt.Println(buff)
}

func readall(file *os.File) {
	buff, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(buff)

}
