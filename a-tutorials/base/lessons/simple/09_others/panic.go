package main

import (
	"errors"
	"fmt"
)

func main() {
	//panicA()
	//panicB()
	panicC()
}

func panicA() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr[5])
}

func panicB() {
	fmt.Println("Enter function main.")
	// 引发 panic。
	panic(errors.New("something wrong"))
	p := recover()
	fmt.Printf("panic: %s\n", p)
	fmt.Println("Exit function main.")
}
func panicC() {
	defer func() {
		p := recover()
		fmt.Printf("panic: %s\n", p)
	}()

	fmt.Println("Enter function main.")
	// 引发 panic。
	panic(errors.New("something wrong"))

	fmt.Println("Exit function main.")
}
