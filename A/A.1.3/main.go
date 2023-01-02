package main

import (
	// "math/randでも動く"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	a := make([]byte, 20)
	rand.Read(a) //セキュリティの目的でrandを使ってはいけない
	fmt.Println(a)
	fmt.Println(hex.EncodeToString(a))
	//セキュリティ目的の場合はcrypto/randを使う
}
