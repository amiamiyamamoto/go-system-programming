package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "examplee.com:80")
	if err != nil {
		panic(err)
	}
	// io.Copy(os.Stdout, conn)

	req, err := http.NewRequest("GET", "example.com", nil)
	req.Write(conn)

	if err != nil {
		panic(err)
	}

	// io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")
	io.Copy(os.Stdout, conn)
	// req.Write(os.Stdout)

}
