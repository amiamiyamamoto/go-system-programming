package main

import "os"

func main() {
	// ハードリンク
	os.Link("oldfile.txt", "newfile.txt")
	// シンボリックリンク
	os.Symlink("oldfile.txt", "newfile-symlink.txt")
	// シンボリックリンクのリンク先を取得
	link, err := os.Readlink(("newfile-symlink.txt"))
}
