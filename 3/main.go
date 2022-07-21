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

	iocopybuffer(file)

}

// バッファを指定してコピー
func iocopybuffer(file *os.File) {
	// io.Copyはデフォルトでは32kbのバッファを使う
	// 8kbのバッファを使う
	buffer := make([]byte, 8*1024)
	io.CopyBuffer(os.Stdout, file, buffer)
}

// 指定したサイズだけコピー
func iocopyn(file *os.File) {
	io.CopyN(os.Stdout, file, 100)
}

// 全てコピー
func iocopy(file *os.File) {
	io.Copy(os.Stdout, file)
}

// 決まったバイト数を読み込む
func readfull(file *os.File) {
	buff := make([]byte, 4)

	// ⇓それ指定のバイト数まで読み込めない場合はエラーが怒る
	// buff := make([]byte, 100)

	_, err := io.ReadFull(file, buff)
	if err != nil {
		panic(err)
	}
	fmt.Println(buff)
}

// 全て読み込む
func readall(file *os.File) {
	buff, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(buff)

}
