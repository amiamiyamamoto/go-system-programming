package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buffer bytes.Buffer
	// buffer.Write([]byte("bytes.Buffer example\n"))
	// buffer.WriteString("bytes.Buffer example\n")
	io.WriteString(&buffer, "bytes.Buffer example\n")
	fmt.Println(buffer.String())
}
