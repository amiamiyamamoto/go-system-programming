// 8.2 Unixドメインソケットの使い方
// ストリーム型のUnixドメインソケット

// クライアント側
conn, err := net.Dial("unix", "socketfile")
if err != nil {
	panic(err)
}
// connを使った読み書き

// サーバ側
listenner, err := net.Listen("unix", "socketfile")
if err != nil {
	panic(err)
}
defer listenner.Close()
conn, err := listener.Accept()
if err != nil {
	// エラー処理
}
// connを使った読み書き