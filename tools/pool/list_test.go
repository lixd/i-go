package pool

import (
	"fmt"
	"testing"
)

func Test_List(t *testing.T) {
	num := 100_0000
	list := NewList(int64(num))
	for i := 0; i < num; i++ {
		list.LPush([]interface{}{"A", "B", "C", "D", "E", "F"})
	}
	for i := 0; i < num; i++ {
		fmt.Println(list.RPopLPush())
	}
}

func BenchmarkName(b *testing.B) {
	num := 100_0000
	list := NewList(int64(num))
	for i := 0; i < num; i++ {
		list.LPush([]interface{}{i})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.RPopLPush()
	}
}
