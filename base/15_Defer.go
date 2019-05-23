package main

import "fmt"

func main() {
	testDefer(1, 2)
}

func testDefer(a int, b int) int {
	//当执行到defer时，暂时不执行 会将defer后的语句压入到独立的栈中(defer栈)
	//当函数执行完毕后在从defer栈中 按照先入后出的方式出栈执行
	res := 0
	defer fmt.Printf("deferA res: %d \n", res)
	defer fmt.Printf("deferB res: %d \n", res)
	defer fmt.Println("deferB")
	res = a + b
	fmt.Printf("a: %d \n", a)
	fmt.Printf("b: %d \n", b)
	return res
}
