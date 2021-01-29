package daily

import "testing"

func Test_numEquivDominoPairs(t *testing.T) {
	type args struct {
		dominoes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: struct{ dominoes [][]int }{dominoes: [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numEquivDominoPairs(tt.args.dominoes); got != tt.want {
				t.Errorf("numEquivDominoPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
