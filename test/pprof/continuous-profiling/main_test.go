package main

import "testing"

func Test_fib(t *testing.T) {
	for i := 0; i < 30; i++ {
		f := fib(i)
		t.Logf("fib(%d) = %d", i, f)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(35)
	}
}
