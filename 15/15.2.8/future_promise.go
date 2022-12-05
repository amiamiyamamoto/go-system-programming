package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(path string) chan string {
	promise := make(chan string)
	go func() {
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("read error %s\n", err.Error())
		} else {
			// 約束を果たした
			promise <- string(content)
		}
	}()
	return promise
}

func printFunc(futureSource chan string) chan []string {
	// 文字列中の関数一覧を返す Futureを返す
	promise := make(chan []string)
	go func() {
		var result []string
		// fetureが解決するまで待って実行
		for _, line := range strings.Split(<-futureSource, "\n") {
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
	// Future/Promiseは1977年の論文で紹介された手法
	// 将来得られるはずの入力(Future)を使ってロジックを作成し
	// 将来値を提供するという約束(Promise)が果たされると必要なデータが揃ったタスクが逐次実行される
	// Goではタスクをgoroutineとして表現し、
	// Futureはバッファなしチャネルの受信、
	// Promiseはそのチャネルへの送信で実現できる
	futureSource := readFile("future_promise.go")
	futureFuncs := printFunc(futureSource)
	fmt.Println(strings.Join(<-futureFuncs, "\n"))
}
