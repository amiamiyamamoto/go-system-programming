package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// ----------------------------------------------
	// bafferの初期化
	// ----------------------------------------------
	// ポインタではなく実態が返ってくる
	var buffer1 bytes.Buffer
	buffer1.Read([]byte{0x10})

	// バイト文字列で初期化
	buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
	// 文字列で初期化
	buffer3 := bytes.NewBufferString("初期化文字列")

	fmt.Println(buffer1, buffer2, buffer3)

	// ----------------------------------------------
	// byte.Readerとstrings.Reader
	// ----------------------------------------------
	bReader1 := bytes.NewReader([]byte{0xff, 0x20, 0x30})
	bReader2 := bytes.NewReader([]byte("文字列をバイト配列にキャスト"))
	fmt.Println(bReader1, bReader2)

	sReader := strings.NewReader("Readerの出力内容は文字列で渡す")
	fmt.Println(sReader)

	// 1ビットが8個で1バイト
	// 2^8＝256通り,最大255
	// ffは、１５＊１６＋１５＝255　になる！
	// 16進数のf →　2進数の1111
	// ffは11111111　→　2^8　-1
	// 10000000 = 01111111 + 1
	// 128 = 127 + 1
	// 16^2 - 1
}
