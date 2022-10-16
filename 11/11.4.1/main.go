package main

import (
	"fmt"
	"os"
)

// os.ExpandEnv()は他の環境変数を組み合わせた文字列が欲しい場合に使う
func main() {
	fmt.Println(os.ExpandEnv("${HOME}/gobin"))
}
