package unittest

import (
	"testing"
)

func Test_Sum(t *testing.T) {
	s := Sum(1, 1)
	if s != 2 {
		t.Fatalf("sum(1,1) failed. Got %d, expected 2.", s)
	}
}

func BenchmarkFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 差不过50ms
		for j := 0; j < 1_0000_0000; j++ {
		}
	}
}

func TestA(t *testing.T) {
	// 差不过50ms
	for {
		for j := 0; j < 1_0000_0000; j++ {
			if j%100_0000 == 0 {
				// time.Sleep(time.Millisecond)
			}
		}
	}
}
