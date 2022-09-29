package main

import "syscall"

func main() {
	kq, err := syscall.Kqueue()
	if err != nil {
		panic(err)
	}
	// 監視対象のファイルディスクリプタを取得
}
