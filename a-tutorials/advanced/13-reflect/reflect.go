package main

import (
	"fmt"
	"reflect"
)

func main() {
	// reflectType()
	// reflectValue1()
	reflectValue2()
}

type i17x struct {
	Name string
}

func (i i17x) A() {
	fmt.Println("A")
}
func reflectType() {
	i := i17x{Name: "17x"}
	t := reflect.TypeOf(i)
	println(t.Name(), t.NumMethod())
}

func reflectValue1() {
	a := "17x"
	v := reflect.ValueOf(a)
	v.SetString("i17x") // 会panic
	println(a)
}
func reflectValue2() {
	a := "17x"
	v := reflect.ValueOf(&a)
	v = v.Elem()
	v.SetString("i17x")
	println(a)
}
