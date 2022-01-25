package goadvanced

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestFilter(t *testing.T) {
	s := make([]int64, 0, 10)
	for i := int64(0); i < 10; i++ {
		s = append(s, i)
	}
	filter := Filter(s, func(s int64) bool {
		return s%2 == 0
	})

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fh := (*reflect.SliceHeader)(unsafe.Pointer(&filter))
	// 根据输出可知，确实是共用的底层数组
	fmt.Printf("sh:%+v\n", sh) // sh:&{Data:824633827728 Len:10 Cap:10}
	fmt.Printf("fh:%+v\n", fh) // fh:&{Data:824633827728 Len:5 Cap:10}
	// 然后s也确实是被影响到了
	fmt.Printf("s:%+v\n", s)      // s:[1 3 5 7 9 5 6 7 8 9]
	fmt.Printf("f:%+v\n", filter) // f:[1 3 5 7 9]
}

func TestSortFloat64Fast(t *testing.T) {
	rand.Seed(time.Now().Unix())
	a := []float64{1.1, 11.11, 2.2, 22, 22, 3.3, 33.33}
	fmt.Printf("before :%+v\n", a) // before :[1.1 11.11 2.2 22 22 3.3 33.33]
	SortFloat64Fast(a)
	fmt.Printf("after :%+v\n", a) // after :[1.1 2.2 3.3 11.11 22 22 33.33]
}
