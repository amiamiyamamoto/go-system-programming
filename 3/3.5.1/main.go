package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

func main() {
	ioSectionReader()
	ioLimitReader()
}

//io.SectionReader
func ioSectionReader() {
	// 読み取りの開始位置を指定できる。
	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)

	// 引数がio.ReaderAt
	// io.ReaderAtを満たさないio.Readerの場合(bytes.Bufferとか)
	buffer := bytes.NewBuffer([]byte("io.ReadAtを満たさないbytes.Bufferを使用する場合"))
	// 文字列やバイト列に書き出し、strings.Readerやbytes.Readerでラップしてから渡す
	data, _ := io.ReadAll(buffer)
	byteRecader := bytes.NewReader(data)
	sr := io.NewSectionReader(byteRecader, 14, 7)

	io.Copy(os.Stdout, sr)

}

//io.LimitReader
func ioLimitReader() {
	reader := strings.NewReader("テストreader")

	lReader := io.LimitReader(reader, 3)
	lReader2 := io.LimitReader(reader, 3)

	io.Copy(os.Stdout, lReader)
	io.Copy(os.Stdout, lReader2)
	// fmt.Println(lReader)
}
