package daily

import (
	"fmt"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{s: "leetcode"}, want: 0},
		{name: "2", args: args{s: "loveleetcode"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstUniqChar("leetcode")
	}
}

func Test1(t *testing.T) {
	a := []int{1, 3, 5, 2, 4, 6}
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
	fmt.Println(a)
}
