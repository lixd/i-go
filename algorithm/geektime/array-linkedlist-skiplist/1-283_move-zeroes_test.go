package array_linkedlist_skiplist

import "testing"

func Test_moveZeroesxx(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{nums: []int{1, 0, 1, 0, 1, 1, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroesxx(tt.args.nums)
		})
	}
}
