package main

import (
	"os"

	"golang.org/x/crypto/ssh"
)

// 取得してきたサーバの鍵情報
var hostsKeyString string = "ecdsa-sha2-nistp256 AAAAE...idDI="

func main() {
	// 秘密鍵の準備
	key, err := os.ReadFile("id_sysprogo")
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
}
