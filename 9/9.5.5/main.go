package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	// 環境変数などの展開
	path := os.ExpandEnv("${GOPATH}/src/github.com/shibukawa/tobubus")
	fmt.Println(path)

	// ホームディレクトリの取得v1.12以降
	// ※「~」がホームディレクトリを表すのはOSではなくシェルの機能
	fmt.Println(os.UserHomeDir())
	// これまでのホームディレクトリの取得方法
	my, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(my.HomeDir)
}

// ~も展開してパスをクリーン化する関数
func Clean2(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		my, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = my.HomeDir + path[1:]
	}
	path = os.ExpandEnv(path)
	return filepath.Clean(path)
}
