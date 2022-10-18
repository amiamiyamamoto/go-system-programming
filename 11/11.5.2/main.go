package main

import (
	"errors"
	"io"
	"log"
	"strings"

	"github.com/google/shlex"
	"github.com/peterh/liner"
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

func main() {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	for {
		if cmd, err := line.Prompt(" "); err == nil {
			if cmd == "" {
				continue
			}
			// ここでコマンド処理
		} else if errors.Is(err, io.EOF) {
			break
		} else if err == liner.ErrPromptAborted {
			log.Print("Aborted")
			break
		} else {
			log.Print("Error reading line: ", err)
		}
	}
}

// 入力されたテキストの分解はスペースで区切られた文字列？
// func main() {
// 	r := regexp.MustCompile(`\s+`)
// 	tokens := r.Split(`echo hello`, -1)
// 	for i, t := range tokens {
// 		fmt.Println(i, t)
// 	}
// }
