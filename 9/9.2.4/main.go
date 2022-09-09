package main

import "os"

func main() {
	/* ディレクトリの作成 */
	// フォルダを一階層だけ作成
	os.Mkdir("setting", 0755)

	// 深いフォルダを一階で作成
	os.MkdirAll("setting/myapp/networlsettings", 0755)

	/* ディレクトリの作成 */

}
