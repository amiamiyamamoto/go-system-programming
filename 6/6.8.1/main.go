package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

// チャンク形式のサーバの実装
// ごんぎつねの文章を100バイトずつぐらいレスポンスとして返すサーバ

var contents = []string{
	"これは、わたしが小さいときに、村の茂平というおじいさんからきいたお話です。",
	"むかしは、私達の村のちかくの、中山というところに小さなお城があって、",
	"中山さまというおとのさまが、おられたそうです",
	"その中山から、少しはなれた山のなかに、「ごん狐」という狐がいました。",
	"ごんは、ひとりぼっちの小狐で、しだのいっぱいしげった森の中に穴をほって住んでいました",
	"そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。",
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		// リクエストを読み込む
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		// レスポンスを書き込む
		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain; charset=UTF-8",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}
func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
