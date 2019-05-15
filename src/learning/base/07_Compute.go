package main

import (
	"fmt"
)

var a = 21.0
var b = 5.0
var c float64

func main() {
	Arithmetic()
}
func Arithmetic() {
	c = a + b
	fmt.Printf("c %.2f \n",c)
	c = a - b
	fmt.Printf("c %.2f \n",c)

	c = a / b
	fmt.Printf("c %.2f \n",c)

	c = a * b
	fmt.Printf("c %.2f \n",c)

	a++
	fmt.Printf("a %.2f \n",a)
	num :=20
	if num%2==0 {
		fmt.Printf("num %v \n",num)
	}else {
		fmt.Println("奇数")
	}
}
