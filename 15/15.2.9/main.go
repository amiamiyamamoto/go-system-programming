package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	// observable (監視される側)を作成
	emitter := make(chan interface{})
	source := observable.Observable(emitter)

	// イベントを受け取るobsever(監視する側)を作成
	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			line := item.(string)
			if strings.HasPrefix(line, "func ") {
				fmt.Println(line)
			}
		},
		ErrHandler: func(err error) {
			fmt.Printf("Encountered error: %v\n", err)
		},
		DoneHandler: func() {
			fmt.Println("Done!")
		},
	}

	// observableとobserverを接続（購読）
	sub := source.Subscribe(watcher)

	// observableに値を投入
	go func() {
		content, err := os.ReadFile("main.go")
		if err != nil {
			emitter <- err
		} else {
			for _, line := range strings.Split(string(content), "\n") {
				emitter <- line
			}
		}
		close(emitter)
	}()
	// 終了まち
	<-sub
}
