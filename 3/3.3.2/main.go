package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var reader io.Reader = strings.NewReader("テストデータ")

	// strings.Readerはio.Closerではないが、
	// io.NopCloserでCloseメソッドを持たせることができる
	var readCloser io.ReadCloser = io.NopCloser(reader)
	readCloser.Close() // 何も起きない

	// io.Reader とio.Writerをつなげて
	// io.ReadWriter型のオブジェクトを作ることができる
	var buf bytes.Buffer
	var r = bufio.NewReader(reader)
	var w = bufio.NewWriter(&buf)
	var readWriter io.ReadWriter = bufio.NewReadWriter(r, w)
	// readWriterを読んでreadWriterに書き込む
	io.Copy(readWriter, readWriter)
	// wに値が入っている
	fmt.Println(w)

}
