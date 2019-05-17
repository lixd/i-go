package main

import "fmt"

//声明切片1 未初始化时为nil
var s1 []int

//声明切片2.1 make方式声明的切片不为nil
var s2 []int = make([]int, 5)
var s3 []int = make([]int, 5, 7)

func main() {
	s4 := make([]int, 5)
	s5 := make([]int, 5, 7)

	printMsg(s1)
	printMsg(s2)
	printMsg(s3)
	printMsg(s4)
	printMsg(s5)
}
func printMsg(s []int) {
	fmt.Printf("len=%d cap=%d slice=%v \n", len(s), cap(s), s)
}
