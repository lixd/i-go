package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
我们发现 Get 方法取出来的对象和上次 Put 进去的对象实际上是同一个，Pool 没有做任何“清空”的处理。但我们不应当对此有任何假设，
因为在实际的并发使用场景中，无法保证这种顺序，最好的做法是在 Put 前，将对象清空。
sync.pool 可以降低分配压力+gc压力，比如 gin 框架，会给每个请求都创建一个 context，就是用的 sync.pool 进行复用。
*/

var defaultGopher, _ = json.Marshal(Gopher{Name: "17x"})

func BenchmarkUnmarshal(b *testing.B) {
	var g *Gopher
	for n := 0; n < b.N; n++ {
		g = new(Gopher)
		json.Unmarshal(defaultGopher, g)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	var g *Gopher
	for n := 0; n < b.N; n++ {
		g = gopherPool.Get().(*Gopher)
		json.Unmarshal(defaultGopher, g)
		g.Reset() // 重置后在放进去
		gopherPool.Put(g)
	}
}

func TestA(t *testing.T) {
	site := 8
	pid := 3
	for i := 0; i < site; i++ {
		fmt.Println((pid + i + 1) % site)
	}
}
