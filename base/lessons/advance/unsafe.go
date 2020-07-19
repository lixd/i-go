package main

import (
	"fmt"
	"unsafe"
)

func main() {
	unsafe1()
}

func unsafe1() {
	var i int = 1
	f := *(*float64)(unsafe.Pointer(&i))
	fmt.Println(unsafe.Pointer(&i))
	fmt.Println(f)
}
