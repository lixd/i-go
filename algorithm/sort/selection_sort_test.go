package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSS(t *testing.T) {
	arr := []int{1, 3, 5, 2, 4, 6}
	fmt.Printf("before sort:%v \n", arr)
	sort := SelectionSort(arr)
	fmt.Printf("after  sort:%v \n", sort)
}

func TestSelectionSort(t *testing.T) {
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
			if got := SelectionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSelectionSort  24.0 ns/op
func BenchmarkSelectionSort(b *testing.B) {
	arr := []int{1, 3, 5, 2, 4, 6}
	for i := 0; i < b.N; i++ {
		SelectionSort(arr)
	}
}
