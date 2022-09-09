package main

import (
	"fmt"
	"os"
	"time"
)

// ファイル出力のシステムコール時は一旦バッファメモリに書き込む
// 実際にストレージに書き込まれたことを確認するにはsyncメソッドを利用する
// ストレージへの書き込みは、メモリへの書き込みに比べて500~1000倍近く遅い

// ▼出力結果
// Write: 38.653µs
// Sync: 23.330627ms
// Close: 43.33µs
func main() {
	f, _ := os.Create("file.txt")
	a := time.Now()
	f.Write([]byte("緑の怪獣"))
	b := time.Now()
	f.Sync()
	c := time.Now()
	f.Close()
	d := time.Now()

	fmt.Printf("Write: %v\n", b.Sub(a))
	fmt.Printf("Sync: %v\n", c.Sub(b))
	fmt.Printf("Close: %v\n", d.Sub(c))
}
