package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	d := 1234
	s := "テキストです"
	f := 1.1

	fmt.Fprintf(file, "入力した数値：%d\n", d)
	fmt.Fprintf(file, "入力した文字列：%s\n", s)
	fmt.Fprintf(file, "入力した少数点有り数値：%f\n", f)
}
