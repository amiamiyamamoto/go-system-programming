package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `１行目
２行目
３行目`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}
