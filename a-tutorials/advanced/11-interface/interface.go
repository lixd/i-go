package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

// func main() {
// 	eface()
// }

func eface() {
	var e interface{}
	f, _ := os.Open("17x.txt")
	e = f
	fmt.Printf("%#v \n", e)
}

func iface() {
	var rw io.ReadWriter
	f, _ := os.Open("17x.txt")
	rw = f
	fmt.Printf("%#v \n", rw)
}

type IA interface {
	Foo()
}
type IB interface {
	Foo()
}
type SA struct{}

func (s SA) Foo() {}

type SB struct{}

func (s SB) Foo() {}
func bar(a IA) {
	fmt.Println(reflect.TypeOf(a))
}
func main() {
	var a IA = SA{}
	var b IB = SA{}
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(b)) // true
	// a b 都能调用bar函数说明 go 语言中的接口只要方法一致就行和接口名没关系
	bar(a)
	bar(b)
}
