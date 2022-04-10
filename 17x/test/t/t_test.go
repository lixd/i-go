package t

import (
	"sync"
	"testing"
)

func BenchmarkName(b *testing.B) {
	var s sync.Map
	s.Store("foo", "bar")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.LoadOrStore("foo", "bar")
	}
}
