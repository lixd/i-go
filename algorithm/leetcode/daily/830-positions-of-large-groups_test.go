package daily

import (
	"reflect"
	"testing"
)

func Test_largeGroupPositions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "1", args: args{s: "abbxxxxzzy"}, want: [][]int{{3, 6}}},
		{name: "2", args: args{s: "aaa"}, want: [][]int{{0, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largeGroupPositions2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}

//   361 ns/op  168 B/op 11 allocs/op
func BenchmarkLargeGroupPositions(b *testing.B) {
	s := "abbxxxxzzy"
	for i := 0; i < b.N; i++ {
		largeGroupPositions(s)
	}
}

// 100 ns/op 48 B/op 2 allocs/op
func BenchmarkLargeGroupPositions2(b *testing.B) {
	s := "abbxxxxzzy"
	for i := 0; i < b.N; i++ {
		largeGroupPositions2(s)
	}
}
