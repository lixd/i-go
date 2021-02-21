package main

import "fmt"

func A() {
	defer A1()
	defer A2()
	panic("panic A")
}

func A1() {
	defer func() {
		fmt.Println("recover A1")
		defer func() {
			fmt.Println("recover A2")
			recover()
		}()
		recover()
		panic("Panic A2")
	}()
	panic("panic A1")
}
func A2() {

}

func main() {
	A()
}
