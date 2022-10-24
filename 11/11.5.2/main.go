package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
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
			cmd, args, err := parseCmd(cmdStr)
			if err != nil {
				log.Print(err)
				break
			}

			var exe_args []string // ワイルドカードなどを展開したargs
			// 引数の処理
			for _, arg := range args {
				// argがパスかどうか判定する（最初の文字が"."or"/"）
				if strings.HasPrefix(arg, ".") || strings.HasPrefix(arg, "/") {
					// カレントディレクトリを取得
					wd, err := os.Getwd()
					if err != nil {
						log.Print("Error get working dir: ", err)
					}
					dirs, err := expandWildcard(arg, wd)
					if err != nil {
						log.Print("Error expandWildcard: ", err)
					}
					exe_args = append(exe_args, dirs...)
				} else {
					exe_args = append(exe_args, arg)
				}
			}

			fmt.Println(exe_args)
			_ = cmd

			// reader, writer := io.Pipe()
			// c1.Stdout = writer
			// c2.Stdin = reader

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
		return []string{expandPath(arg, workDir)}, nil
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
