package main

import (
	"io"
	"os"
)

func main() {
	oldfn := os.Args[1]
	newfn := os.Args[2]
	oldfile, err := os.Open(oldfn)
	if err != nil {
		panic(err)
	}
	defer oldfile.Close()

	newfile, err := os.Create(newfn)
	if err != nil {
		panic(err)
	}
	defer newfile.Close()

	_, err = io.Copy(newfile, oldfile)
	if err != nil {
		panic(err)
	}
}
