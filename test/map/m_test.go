package main

import (
	"strconv"
	"sync"
	"testing"
)

// 42.6 ns/op
func BenchmarkSMap(b *testing.B) {
	var smap sync.Map
	for i := 0; i < b.N; i++ {
		smap.LoadOrStore("key", "value")
	}
}

//  2.42 ns/op
func BenchmarkMap(b *testing.B) {
	m := make(map[string]struct{})
	for i := 0; i < b.N; i++ {
		_ = m["key"]
	}
}

func BenchmarkDeepCopy(b *testing.B) {
	var (
		src       = make(map[string]*int64)
		det       = make(map[string]int64)
		v   int64 = 1
	)
	for i := int64(0); i < 1000; i++ {
		src[strconv.Itoa(int(i))] = &v
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deepCopy(src, det)
	}
}

func BenchmarkDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Demo()
	}
}
