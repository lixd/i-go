package cpucache

import "testing"

func BenchmarkFor1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Row()
	}
}

func BenchmarkFor2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Col()
	}
}
