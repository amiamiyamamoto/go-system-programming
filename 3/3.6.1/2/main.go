package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1行目です
2行目　です
3行目！です`

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	scanner.Split(bufio.ScanWords) //区切り位置をデフォルトの行ごとから単語ごとに設定変更
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}

}
