package cpucache

import "testing"

// 748990 ns/op
func BenchmarkFor1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Row()
	}
}

// 1382133 ns/op
func BenchmarkFor2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Col()
	}
}
