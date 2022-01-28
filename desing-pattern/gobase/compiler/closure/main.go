package main

import "fmt"

func main() {
	a := 1
	b := 2
	go func() {
		fmt.Println(a, b)
	}()
	a = 100
}

// 查看当前程序闭包变量捕获的情况：go tool compile -m=2 main.go|grep capturing

func do() {
	a := 1
	func() {
		fmt.Println(a)
		a = 2
	}()
}

func do2() {
	a := 1
	func1(&a)
}

func func1(a *int) {
	fmt.Println(*a)
	*a = 2
}
