package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
)

// データグラム型のUnixドメインソケット、クライアント
func main() {

	clientPath := filepath.Join(os.TempDir(), "unixdomainsocket-client")
	// ⇓エラーが返ってくるが、ファイルが存在しないなら問題ないのでエラーチェックなし
	os.Remove(clientPath)
	conn, err := net.ListenPacket("unixgram", clientPath)
	if err != nil {
		panic(err)
	}
	// conn, err := net.Dial("unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
	// if err != nil {
	// 	panic(err)
	// }

	// 送信先のアドレス
	unixServerAddr, err := net.ResolveUnixAddr("unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
	var serverAddr net.Addr = unixServerAddr
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("Sending to server")
	_, err = conn.WriteTo([]byte("Hello from Client"), serverAddr)
	if err != nil {
		panic(err)
	}
	log.Println("Receiving from server")
	buffer := make([]byte, 1500)
	length, _, err := conn.ReadFrom(buffer)
	if err != nil {
		panic(err)
	}
	log.Printf("Received: %s\n", string(buffer[:length]))
}
