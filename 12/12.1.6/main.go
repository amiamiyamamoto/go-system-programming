package main

import (
	"fmt"
	"os"
)

func main() {
	// プロセスに割り当てられた作業フォルダを表示する
	wd, _ := os.Getwd()
	fmt.Println(wd)
}
