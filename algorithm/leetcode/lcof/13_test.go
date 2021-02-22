package lcof

import "testing"

func Test_movingCount(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{m: 1, n: 2, k: 1}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumPos(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 12}, want: 3},
		{name: "2", args: args{n: 120}, want: 3},
		{name: "3", args: args{n: 100}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumPos(tt.args.n); got != tt.want {
				t.Errorf("sumPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
