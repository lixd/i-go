package main

import "fmt"

// 数据类型
func main() {
	a := 100
	//uint8
	var b byte = 200
	//int32
	var c rune = 2000
	fmt.Printf("%T %v \n", a, a)
	fmt.Printf("%T %v \n", b, b)
	fmt.Printf("%T %v \n", c, c)
	dataTypeConv()
}

func dataTypeConv() {
	a := 90
	b := 90.5
	avg := ((float64(a)) + b) / 2
	fmt.Printf("%T %v \n", avg, avg)
}
