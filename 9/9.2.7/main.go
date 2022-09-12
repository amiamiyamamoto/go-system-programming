package main

import "syscall"

func main() {

	// OS固有のファイル属性を取得する
	// windows
	// internalStat := info.Sys().(syscall.Win32FileAttributeDate)

	// windows以外
	internalStat := info.Sys().(*syscall.Stat_t)
}
