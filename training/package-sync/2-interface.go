package main

import (
	"fmt"
)

type IceCreamMaker interface {
	// Hello greets a customer
	Hello()
}
type Ben struct {
	id   int
	name string
}

func (b *Ben) Hello() {
	fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
}

type Jerry struct {
	name string
}

// type Jerry struct {
// 	field1 *[5]byte
// 	field2 int
// }

func (j *Jerry) Hello() {
	fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.name)
}

func main() {
	var ben = &Ben{name: "Ben"}
	var jerry = &Jerry{" Jerry"}
	var maker IceCreamMaker = ben
	var loop0, loop1 func()
	// loop0 和 loop1 互相调用，内部只是对maker进行赋值
	// interface类型是8byte按道理说是刚好等于64位CPU的 single machine word,是一个原子操作，整个程序不会出现任何问题。
	// 实际上 interface 结构体有两个字段组成，一个 type 一个 data字段，需要赋值两次，所以在并发情况下可能会出现 type 和data 对应不上的情况
	// 所以程序会出现异常
	loop0 = func() {
		maker = ben
		go loop1()
	}
	loop1 = func() {
		maker = jerry
		go loop0()
	}
	go loop0()
	for {
		maker.Hello()
	}
}
