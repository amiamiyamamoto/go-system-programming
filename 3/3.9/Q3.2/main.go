package main

import (
	"crypto/rand"
	"io"
	"os"
)

// ファイルを作成して、ランダムな内容で埋める
func main() {
	randReader := rand.Reader
	// lread := io.LimitReader(randReader, 1024)

	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.CopyN(file, randReader, 1024)

}
