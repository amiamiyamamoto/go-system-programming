package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

// パイプライニングのクライアント実装
func main() {
	sendMessage := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	var conn net.Conn = nil
	var err error
	requests := make([]*http.Request, 0, len(sendMessage))

	conn, err = net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Access\n")
	defer conn.Close()

	// リクエストだけ先に送る
	for i := 0; i < len(sendMessage); i++ {
		lastMessage := i == len(sendMessage)-1
		request, err := http.NewRequest("GET", "http://localhost:8888?message="+sendMessage[i], nil)
		if lastMessage {
			request.Header.Add("Connection", "close")
		} else {
			request.Header.Add("connection", "keep-alive")
		}
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		fmt.Println("send: ", sendMessage[i])
		requests = append(requests, request)
	}
	// レスポンスをまとめて受信
	reader := bufio.NewReader(conn)
	for _, request := range requests {
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
	}
}
