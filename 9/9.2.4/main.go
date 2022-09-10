package main

import (
	"io"
	"os"
)

func main() {
	/* ディレクトリの作成 */
	// フォルダを一階層だけ作成
	os.Mkdir("setting", 0755)
	// 深いフォルダを一階で作成
	os.MkdirAll("setting/myapp/networlsettings", 0755)

	/* ファイルの削除、移動、リネーム */
	// ファイルや空のディレクトリの削除
	// ファイルの削除のシステムコールを呼び出し、失敗したらディレクトリ削除を呼び出す仕組み
	os.Remove("server.log")
	// ディレクトリを中身ごと削除
	os.RemoveAll("wirkdir")

	// 先頭100バイトで切る
	os.Truncate("server.log", 100)
	// Truncateメソッドを利用する場合
	file, _ := os.Open("server.log")
	file.Truncate(100)

	// リネーム
	os.Rename("old_name.txt", "new_name.txt")
	// 移動
	os.Rename("olddir/file.txt", "newdir/file.txt")
	// 移動先はディレクトリではだめ
	os.Rename("olddir/file.txt", "newdir/") // エラー発生

	// マウントされて元のデバイスが異なる場合にはrenameシステムコールでの移動はできない
	err := os.Rename("sample.rst", "/tmp/sample.rst")
	if err != nil {
		panic(err)
		// 以下のようなエラーが表示される
		// rename sample.rst /tmp/sample.rst: cross-device link
	}
	// デバイスやドライブが異なる場合、ファイルを開いてコピーを作り
	// その後にソースファイルを削除する
	// ※mvコマンドもそのような挙動をしている
	{
		oldFile, err := os.Open("old_name.txt")
		if err != nil {
			panic(err)
		}
		newFile, err := os.Create("/other_device/new_file.txt")
		if err != nil {
			panic(err)
		}
		defer newFile.Close()
		_, err = io.Copy(newFile, oldFile)
		if err != nil {
			panic(err)
		}
		oldFile.Close()
		os.Remove("old_name.txt")
	}
}
