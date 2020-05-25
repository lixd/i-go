package main

import "fmt"

func main() {
	a := 100
	//uint8
	var b byte = 200
	//int32
	var c rune = 2000
	fmt.Printf("%T %v \n", a, a)
	fmt.Printf("%T %v \n", b, b)
	fmt.Printf("%T %v \n", c, c)
}
