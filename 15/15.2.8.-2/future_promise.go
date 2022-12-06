package main

import (
	"fmt"
	"os"
	"strings"
)

type StringFuture struct {
	receiver chan string
	cache    string
}

func NewStringFuture() (*StringFuture, chan string) {
	f := &StringFuture{
		receiver: make(chan string),
	}
	return f, f.receiver
}

func (f *StringFuture) Get() string {
	r, ok := <-f.receiver
	if ok {
		close(f.receiver)
		f.cache = r
	}
	return f.cache
}

func (f *StringFuture) Close() {
	close(f.receiver)
}

func readFile(path string) *StringFuture {
	future, promise := NewStringFuture()
	go func() {
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("read error %s \n", err.Error())
		} else {
			// 約束を果たした
			promise <- string(content)
		}
	}()
	return future
}
func printFunc(futureSource *StringFuture) chan []string {
	// 文字列中の関数一覧を返す
	promise := make(chan []string)
	go func() {
		var result []string
		// futureが解決するまで待って実行
		for _, line := range strings.Split(futureSource.Get(), "\n") {
			if strings.HasPrefix(line, "func ") {
				result = append(result, line)
			}
		}
		// 約束を果たした
		promise <- result
	}()
	return promise

}

func main() {
	// ファイルを読み込む
	futureSource := readFile("future_promise.go")
	defer futureSource.Close()
	futureFuncs := printFunc(futureSource)
	fmt.Println(strings.Join(<-futureFuncs, "\n"))
	fmt.Println("複数回参照可能", "\n-------------------------------------------------\n", futureSource.Get())
	// fmt.Println("複数回参照可能")
}
