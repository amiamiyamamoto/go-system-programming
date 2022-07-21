package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

var reader io.Reader = strings.NewReader("テストデータ")

// strings.Readerはio.Closerではないが、
// io.NopCloserでCloseメソッドを持たせることができる
var readCloser io.ReadCloser = io.NopCloser(reader)

// io.Reader とio.Writerをつなげて
// io.ReadWriter型のオブジェクトを作ることができる
var buf bytes.Buffer
var r = bufio.NewReader(reader)
var w = bufio.NewWriter(&buf)
var readWriter io.ReadWriter = bufio.NewReadWriter(r, w)

func main() {
	readCloser.Close() // 何も起きない

	// readWriterを読んでreadWriterに書き込む
	io.Copy(readWriter, readWriter)
	// wに値が入っている
	fmt.Println(w)

}
