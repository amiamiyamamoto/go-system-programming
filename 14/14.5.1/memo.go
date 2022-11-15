package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 現在実行中のOSスレッドでのみgoroutineが実行されるように束縛する
	runtime.LockOSThread()

	// UnlockOSThreadの呼び出しか、goroutine終了時に解除される
	runtime.UnlockOSThread()
	// メインスレッドでの実行が強制されるライブラリを利用する場合などに使う
	// 現在のスレッドがメインスレッドかどうかはわからない
	// main関数のinitは確実にメインスレッドで実行される

	// Goのランタイムでは、シグナルを受け取るスレッドを固定するためにこれらの関数を使っている

	// ---------------------------------

	// 14.5.2
	// 現在実行中のgoroutineを一時中断して他のgoroutineに処理を回す
	runtime.Gosched()

	// ---------------------------------

	// 14.5.3
	// 実行環境のCPUコア数がわかる
	// 物理コア数ではなく、論理コア数を返す
	fmt.Println(runtime.NumCPU())
	// 同時に使用するスレッド数を制御する
	runtime.GOMAXPROCS(1)
	// Go1.5からデフォルトで設定されるのでわざわざ設定する必要はない
	// 最速を狙う場合は論理コア数ではなく物理コア数を設定したほうが早い場合がある
	// 物理コア数を返すAPIはない
	// GOMAXPROCS環境変数で設定することもできる
}
