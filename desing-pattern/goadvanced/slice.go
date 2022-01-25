package goadvanced

import (
	"reflect"
	"sort"
	"unsafe"
)

// Filter 原地删除切片
func Filter(s []int64, fn func(s int64) bool) []int64 {
	// 这里借助传进来的切片的底层数组构建一个空切片来用
	// 优点: 省掉了一次空间分配
	// 缺点: 修改了同一个底层数组，这样也会影响到外层的切片s,使用时需要注意
	b := s[:0]
	for _, v := range s {
		if !fn(v) {
			b = append(b, v)
		}
	}
	return b
}

// SortFloat64Fast []float64切片排序
func SortFloat64Fast(a []float64) {
	var b []int
	// 转成成 reflect.SliceHeader，然后通过更新结构体值的方式把[]float64的数据赋值给[]int
	// 二者共享底层数组，因此对b排序就是对a排序
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	ah := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	*bh = *ah
	sort.Ints(b)
}
