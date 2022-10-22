package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"path/filepath"
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

	for token, err := l.Next(); err == nil; token, err = l.Next() {
		args = append(args, token)
	}
	return
}

func main() {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	for {
		if cmdStr, err := line.Prompt("$"); err == nil {
			if cmdStr == "" {
				continue
			}
			// ここでコマンド処理
			fmt.Println(cmdStr)
			cmd, args, err := parseCmd(cmdStr)
			if err != nil {
				log.Print(err)
				break
			}
			_ = cmd
			// 引数の処理
			for _, arg := range args {
				// 絶対パス処理とワイルドカード展開を追加する
				// argがパスじゃない場合も大丈夫？
			}

			reader, writer := io.Pipe()
			c1.Stdout = writer
			c2.Stdin = reader

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

// 相対パスを絶対パスに変換する
func expandPath(dir, workDir string) string {
	if filepath.IsAbs(dir) {
		return dir
	}
	return filepath.Join(workDir, dir)
}

// ワイルドカードなどの展開
func expandWildcard(arg, workDir string) ([]string, error) {
	if !strings.ContainsAny(arg, "*?[") {
		return []string{arg}, nil
	}
	files, err := filepath.Glob(expandPath(arg, workDir))
	if len(files) == 0 {
		return nil, err
	}
	return files, err
}

// 入力されたテキストの分解はスペースで区切られた文字列？
// func main() {
// 	r := regexp.MustCompile(`\s+`)
// 	tokens := r.Split(`echo hello`, -1)
// 	for i, t := range tokens {
// 		fmt.Println(i, t)
// 	}
// }
