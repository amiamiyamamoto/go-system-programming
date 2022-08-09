package main

import (
	"bytes"
	"mycopyn/mycopyn"
	"os"
	"strings"
)

func main() {

	mycopyn.CopyN(os.Stdout, strings.NewReader("テストio.Reader"), 5)
	mycopyn.CopyN(os.Stdout, bytes.NewReader([]byte("test io.Reader")), 5)
}
