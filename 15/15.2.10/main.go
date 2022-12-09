package main

// protoactor-goのサンプルコード
import (
	"fmt"

	// console "github.com/asynkronIT/goconsole"
	// "github.com/AsynkronIT/protoactor-go/actor"
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

// メッセージの構造体
type Hello struct{ Who string }

// アクターの構造体
type helloActor struct{}

// アクターのメールボックス受診時に呼ばれるメソッド
func (state *helloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}

}

func main() {
	// サンプルコードが書籍と変更されている
	context := actor.RootContext{}
	// props := actor.FromInstance(&helloActor{})
	props := actor.PropsFromProducer(func() actor.Actor { return &helloActor{} })
	// pid:= actor.Spawn(props)
	pid, err := actor.Spawn(props)
	if err != nil {
		panic(err)
	}
	// pid.Tell(&hello{Who: "Roger"})
	context.Send(pid, Hello{Who: "Roger"})
	console.ReadLine()
}
