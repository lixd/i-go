package main

import "testing"

//go:noinline
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var result int

func BenchmarkInline(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = max(-1, i)
	}
	result = r
}
