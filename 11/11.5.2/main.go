package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/shlex"
)

func parseCmd(cmdStr string) (cmd string, args []string, err error) {
	l := shlex.NewLexer(strings.NewReader(cmdStr))
	cmd, err = l.Next()
	if err != nil {
		return
	}
	for token, err := l.Next(); err != nil; {
		args = append(args, token)
	}
	return
}

// 入力されたテキストの分解はスペースで区切られた文字列？
func main() {
	r := regexp.MustCompile(`\s+`)
	tokens := r.Split(`echo hello`, -1)
	for i, t := range tokens {
		fmt.Println(i, t)
	}
}
