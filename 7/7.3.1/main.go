package main

import (
	"fmt"
	"net"
	"time"
)

// UDP マルチキャスト サーバー側の実装例
const interval = 3 * time.Second

func main() {
	fmt.Println("Start tick server at 244.0.0.2:9999")
	conn, err := net.Dial("udp", "224.0.0.2:9999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	start := time.Now()
	wait := start.Truncate(interval).Add(interval).Sub(start)
	time.Sleep(wait)
	ticker := time.Tick(interval)
	for now := range ticker {
		conn.Write([]byte(now.String()))
		fmt.Println("Tick: ", now.String())
	}
}
