package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{4, 5, 9},
		{6, 7, 13},
		{1, 1, 2},
		{0, 0, 0}}

	for _, tt := range tests {
		if result := add(tt.a, tt.b); result != tt.c {
			t.Errorf("test add %d + %d except %d but got %d", tt.a, tt.b, tt.c, result)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	a := 1
	bb := 2
	res := 3
	// 前面都是在准备数据 所以计算时可以除去这部分时间
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if result := add(a, bb); result != res {
			b.Errorf("test add %d + %d except %d but got %d", a, bb, res, result)
		}
	}
}
