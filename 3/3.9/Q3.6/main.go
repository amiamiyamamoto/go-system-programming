package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SUSTEM")
	programming = strings.NewReader("ORIGRAMMING")
)

//　上記文字を使って、「ASCII」という文字を標準出力に表示させる
func main() {
	var stream io.Reader
	data := make([]byte, 5)
	b := make([]byte, 1)

	programming.ReadAt(b, 5)
	data = append(data, b...)

	system.Read(b)
	data = append(data, b...)

	computer.Read(b)
	data = append(data, b...)

	programming.ReadAt(b, 8)
	data = append(data, b...)
	programming.ReadAt(b, 8)
	data = append(data, b...)

	// fmt.Fprintln(os.Stdout, string(data))
	stream = bytes.NewReader(data)

	io.Copy(os.Stdout, stream)
}
