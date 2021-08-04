package main

import (
	"testing"
)

func Test_fibonacci(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "0", args: args{n: 0}, want: 0},
		{name: "1", args: args{n: 1}, want: 1},
		{name: "1", args: args{n: 2}, want: 1},
		{name: "2", args: args{n: 3}, want: 2},
		{name: "2", args: args{n: 4}, want: 3},
		{name: "2", args: args{n: 5}, want: 5},
		{name: "2", args: args{n: 6}, want: 8},
		{name: "2", args: args{n: 7}, want: 13},
		{name: "2", args: args{n: 8}, want: 21},
		{name: "2", args: args{n: 9}, want: 34},
		{name: "2", args: args{n: 10}, want: 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args.n); got != tt.want {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(10)
	}
}
