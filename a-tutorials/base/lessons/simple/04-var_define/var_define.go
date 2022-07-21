package main

import "fmt"

// 变量定义
func main() {
	simple()
	constant()
	arithmetic()
}

func simple() {
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
	// 10 20
	fmt.Println(x, y)
	x, y = y, x
	// 20 10
	fmt.Println(x, y)
}

type user struct {
	Age  int
	Name string
}

func constant() {
	const NAME string = "illusory"
	u := user{30, NAME}
	const (
		A = 10
		B
		C
		D = iota
		E = iota
	)
	fmt.Println(u, A, B, C, D, E)
}
func arithmetic() {
	var a = 21.0
	var b = 5.0
	var c float64

	c = a + b
	fmt.Printf("c %.2f \n", c)
	c = a - b
	fmt.Printf("c %.2f \n", c)

	c = a / b
	fmt.Printf("c %.2f \n", c)

	c = a * b
	fmt.Printf("c %.2f \n", c)

	a++
	fmt.Printf("a %.2f \n", a)
	num := 20
	if num%2 == 0 {
		fmt.Printf("num %v \n", num)
	} else {
		fmt.Println("奇数")
	}
}
