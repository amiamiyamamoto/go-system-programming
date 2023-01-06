package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

// サーバの鍵情報
var hostsKeyString string = "ecdsa-sha2-nistp256 AAAAE...idDI="

// scpをGoで実装する
func main() {

	// ssh
	key, err := os.ReadFile("id_sysorogo")
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}
	hostkey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(hostsKeyString))
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.FixedHostKey(hostkey),
	}

	//接続開始
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

	// scpでファイルをリモートに送信
	go func() {
		w, _ := session.StdinPipe()
		defer w.Close()
		content := []byte("Go言語でシステムプログラミング\n")
		fmt.Fprintln(w, "D0755", 0, "testdir") //mkdir
		fmt.Fprintln(w, "C0644", len(content), "testfile1")
		w.Write(content)
		fmt.Fprint(w, "\x00") // transfer end with \x00
	}()
	err = session.Run("/usr/bin/scp -tr ./")
	if err != nil {
		panic(err)
	}
}
