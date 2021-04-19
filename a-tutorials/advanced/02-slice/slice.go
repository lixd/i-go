package main

import "fmt"

func main() {
	Slice1()
	// SliceArray()
	// SliceAppend()
}

func Slice1() {
	// 定义
	define()
	// 访问
	access()
	// 添加元素
	add()
	// 移除元素
	del()
}
func define() {
	// 只定义 未初始化 Len Cap 和底层数组
	var s1 []int
	fmt.Println(s1)
	// 和 var 定义一样 只不过 new 关键字返回的是指针
	s2 := new([]int)
	fmt.Println(s2)
	// 指定 Len Cap 同时初始化底层数组
	s3 := make([]int, 2, 4)
	fmt.Println(s3)
}
func access() {
	s1 := make([]int, 4, 8)
	// 切片访问下标不能超过 Len
	// panic: runtime error: index out of range [4] with length 4
	fmt.Println(s1[4])
}

func add() {
	s1 := []int{1, 2}
	s1 = append(s1, 3, 4, 5, 6, 7)
	fmt.Println(cap(s1)) // 扩容后容量为 8
}

func del() {
	s1 := []int{1, 2, 3, 4, 5}
	before := s1[:1] // 1
	after := s1[2:]  // 3 4 5
	s2 := append(before, after...)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2) // 1 3 4 5
}

func StringSlice() {
	ss := new([]string)
	*ss = append(*ss, "指月")
}
func SliceArray() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[:3] // 索引范围[)左闭右开
	s2 := arr[6:]
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	s3 := append(s2, 10)
	fmt.Println("s2:", s2) // s2:[6 7 8 9]
	fmt.Println("s3:", s3) // s3:[6 7 8 9 10]
}
