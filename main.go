package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// pipe()
	aa := Ami{Name: "amiami"}

	// ポインタ型にメソッドを持たせないとフィールドの値を変更することはできない
	// ポインタ型のメソッドでも、struct型変数の状態でメソッドを呼べる
	fmt.Println(aa.changename("yamamoto"))
	fmt.Println(aa.Name)

	fmt.Println(aa.reallyrename("yamamoto"))
	fmt.Println(aa.Name)

	// ポインタ型からも呼べる
	(&aa).pointamethod()
	fmt.Println(aa.Name)
}

type Ami struct {
	Name string
}

func (ami Ami) changename(name string) string {
	ami.Name = name
	return ami.Name
}
func (ami *Ami) reallyrename(name string) string {
	ami.Name = name
	return name
}
func (ami *Ami) pointamethod() {
	ami.Name += "ami"
}

// ポインタ型・構造体のメソッドどちらもフィールドを呼び出せる
func (ami Ami) myname() string {
	name := ami.Name
	return name
}

func (ami *Ami) aminame() string {
	return ami.Name
}

// &Ami // 型名の記述にこの書き方はない
//
// Ami // Ami構造体型
// *Ami // Ami構造体のポインタ型
// ポインタとはアドレスを格納する変数のこと

func point() {
	var abc string = "yamamoto"
	var yama = &abc
	println(abc, *yama, yama)
}

func pipe() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		io.Copy(pw, strings.NewReader("io.Pipeのテスト"))
	}()

	if _, err := io.Copy(os.Stdout, pr); err != nil {
		log.Fatal(err)
	}

}
