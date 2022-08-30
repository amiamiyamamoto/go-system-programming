package main

import (
	"fmt"
	"net"
)

// UDP マルチキャスト クライアント側の実装例
func main() {
	fmt.Println("Listen tick server 224.0.0.2:9999")
	address, err := net.ResolveUDPAddr("udp", "224.0.0.2:9999")
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := listener.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now    %s\n", string(buffer[:length]))
	}
}
