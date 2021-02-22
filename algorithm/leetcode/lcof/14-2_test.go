package lcof

import (
	"testing"
)

func Test_cuttingRope2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "2", args: args{n: 2}, want: 1},
		{name: "3", args: args{n: 10}, want: 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRope2(tt.args.n); got != tt.want {
				t.Errorf("cuttingRope2() = %v, want %v", got, tt.want)
			}
		})
	}
}
