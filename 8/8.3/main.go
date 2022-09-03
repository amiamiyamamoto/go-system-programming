// Windows名前付きパイプ
// npipeというパッケージでWindowsの名前付きパイプを利用できる

// サーバー
ln, err := npipe.Listen(`\\.\pipe\mypipename`)
if err != nil {
	panic(err)
}
for {
	conn, err := ln.Accept()
	if err != nil {
		// エラー処理
		continue
	}
	go handleConnection(conn)
}

// クライアント
conn, err := npipe.Dial(`\\.\pipe\mypipename`)
if err != nil {
	// エラー処理
}
fmt.Fprintf(conn, "Hi server!\n")
msg, err := bufio.NewReader(conn).ReadString('\n')