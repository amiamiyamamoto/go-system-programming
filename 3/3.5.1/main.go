package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
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
