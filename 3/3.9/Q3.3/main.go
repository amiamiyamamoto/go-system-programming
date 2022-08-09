package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

// zipファイルの書き込み
func main() {
	// sr := strings.NewReader("sample text for zip")

	file, err := os.Create("sample.zip")
	// file, err := os.OpenFile("sample.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	newfile, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(newfile, strings.NewReader("zip test"))
}
