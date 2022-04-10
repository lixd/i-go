package sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	sort := quickSort2(arr)
	fmt.Println("before:", arr)
	fmt.Println("after:", sort)
}

// 1143 ns/op
func BenchmarkQuickSort(b *testing.B) {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	for i := 0; i < b.N; i++ {
		quickSort(arr)
	}
}
