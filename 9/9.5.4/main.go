package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// パスのクリーン化
func main() {

	// パスをそのままクリーンにする
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))
	// path/path.go

	// パスを絶対パスに整形
	abspath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println(abspath)
	// /usr/local/go/src/path/filepath/path_unix.go

	// パスを相対パスに整形
	relpath, _ := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/path/filepath/path.go")
	fmt.Println(relpath)
	// path/filepath/path.go

	// 9.5.5 環境変数などの展開
	path := os.ExpandEnv("${GOPATH}/src/github.com/shibukawa/tobubus")
	fmt.Println(path)
}
