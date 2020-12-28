package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBS(t *testing.T) {
	arr := []int{1, 3, 5, 2, 4, 6}
	fmt.Printf("before sort:%v \n", arr)
	sort := BubbleSort(arr)
	fmt.Printf("after  sort:%v \n", sort)
}

func TestBubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{arr: []int{1, 3, 5, 2, 4, 6}}, want: []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkBubbleSort 22.6 ns/op
func BenchmarkBubbleSort(b *testing.B) {
	arr := []int{1, 3, 5, 2, 4, 6}
	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}
}
