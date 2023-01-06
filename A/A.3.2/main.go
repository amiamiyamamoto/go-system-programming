package main

import (
	"fmt"
	"os"
	"time"

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
	if err != nil {
		panic(err)
	}

	// サーバーの鍵の準備
	hostkey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(hostsKeyString))
	if err != nil {
		panic(err)
	}

	// 接続設定
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.FixedHostKey(hostkey),
	}

	// 通信開始
	conn, err := ssh.Dial("tcp", "localhost:1222", config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// コマンドを実行して出力結果を取得
	output, err := session.CombinedOutput("ssh -V")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}

// CombineOutputメソッドの内部
// func (s *Session) CombinedOutput(cmd string) ([]byte, error) {
// 	var b singleWriter
// 	s.Stdout = &b
// 	s.Stderr = &b
// 	err := s.Run(cmd)
// 	return b.b.Byte(), err
// }
