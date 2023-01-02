package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

func main() {
	// 種からソースを生成
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	src := rand.NewSource(seed.Int64())
	// ソースから乱数生成器を作成
	rng := rand.New(src)
	fmt.Println(rng.Int())
	fmt.Println(rng.Int())
	fmt.Println(rng.Int())

}
