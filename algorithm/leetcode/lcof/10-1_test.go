package lcof

import (
	"fmt"
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFib(t *testing.T) {
	i := fib(45)
	fmt.Println(i)
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// fib(20)
		fib2(20)
	}
}
