package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// strings.Readerはio.Closerではないが、
// io.NopCloserでCloseメソッドを持たせることができる
var reader io.Reader = strings.NewReader("テストデータ")
var readCloser io.ReadCloser = io.NopCloser(reader)

// io.Reader とio.Writerをつなげて
// io.ReadWriter型のオブジェクトを作ることができる
var testReader io.Reader = strings.NewReader("テストデータReader")
var r = bufio.NewReader(testReader)
var w = bufio.NewWriter(os.Stdout)
var readWriter io.ReadWriter = bufio.NewReadWriter(r, w)

func main() {
	readCloser.Close() // 何も起きない

	io.Copy(w, readWriter)
}
