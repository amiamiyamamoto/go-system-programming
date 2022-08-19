package main

import "net"

// 最小限のTCPサーバの実装
func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	//一度で終了しないためにAccept()を何度も繰り返し呼ぶ
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
			// handle error
		}
		// 1リクエスト処理中に他のリクエストのAccept()が行えるように
		// Gorutineを使って非同期にレスポンスを処理する
		go func() {
			// connを使った読み書き
		}()
	}
}
