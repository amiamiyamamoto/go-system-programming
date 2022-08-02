package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	oldfile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldfile.Close()
	old, err := io.ReadAll(oldfile)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("new.txt", old, 0666)
	if err != nil {
		panic(err)
	}
}
