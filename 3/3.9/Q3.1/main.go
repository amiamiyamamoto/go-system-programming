package main

import (
	"io"
	"os"
)

func main() {
	oldfile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldfile.Close()

	newfile, err := os.Create("new2.txt")
	if err != nil {
		panic(err)
	}
	defer newfile.Close()

	_, err = io.Copy(newfile, oldfile)
	if err != nil {
		panic(err)
	}
}
