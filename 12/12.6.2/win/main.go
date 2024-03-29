package main

import (
	"io"
	"os"

	winpty "github.com/iamacarpet/go-winpty"
)

// Windowsの場合
func main() {
	pty, err := winpty.Open("", "check.exe")
	if err != nil {
		panic(err)
	}
	defer pty.Close()
	go func() {
		io.Copy(os.Stdout, pty.StdOut)
	}()
	// press any key to exet
	buffer := make([]byte, 1)
	for {
		_, err = os.Stdin.Read(buffer)
		if err == io.EOF || buffer[0] == '\n' {
			break
		}
	}
}
