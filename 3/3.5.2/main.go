package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 32ビットのビッグエンディアンのデータ(10000)
	data := []byte{0x00, 0x00, 0x27, 0x10}
	var j int32
	var i int32

	binary.Read(bytes.NewReader(data), binary.LittleEndian, &j)
	fmt.Printf("litteleEndian: %d\n", j)
	// エンディアンの変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
