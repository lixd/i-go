package main

import "fmt"

func main() {
	var (
		a int
		b string
		c float32
		d []int
		e int32 = 8
	)
	fmt.Println(a, b, c)
	fmt.Printf("%T ,%v \n", a, a)
	fmt.Printf("%T ,%v \n", b, b)
	fmt.Printf("%T ,%v \n", c, c)
	fmt.Printf("%T ,%v \n", d, d)
	fmt.Printf("%T ,%v \n", e, e)

	x := 10
	y := 20
	//10 20
	fmt.Println(x, y)
	x, y = y, x
	//20 10
	fmt.Println(x, y)
}
