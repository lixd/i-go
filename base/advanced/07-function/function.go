package main

import (
	"fmt"
	"reflect"
)

type A struct {
	name string
}

func (a A) GetName() string {
	a.name = "Hi " + a.name
	return a.name
}
func (a *A) SetName() string {
	a.name = "Hi " + a.name
	return a.name
}

func first() {
	a := A{name: "17x"}
	fmt.Println(a.GetName())  // Hi 17x
	fmt.Println(A.GetName(a)) // Hi 17x
}
func second() {
	t1 := reflect.TypeOf(A.GetName)
	t2 := reflect.TypeOf(NameOfA)
	// Go 语言中函数类型只与参数和返回值有关
	// 所以这两个类型值相等就可以证明:方法本质上就是普通的函数
	fmt.Println(t1 == t2) // true
}
func Third() {
	a := A{name: "17x"}
	pa := &a
	// 通过值调用指针接收者的方法
	fmt.Println(a.SetName())
	// 通过指针调用值接收者的方法
	fmt.Println(pa.GetName())
	// Go 语言提供的语法糖，编译阶段会进行如下转换
	// pa.GetName()-->(*pa).GetName()
	// a.SetName()-->(&a).SetName()
	// 字面量 编译期间无法获取地址 所以语法糖无法生效
	// A{name: "17x"}.SetName()

}

func Fourth() {
	a := A{name: "17x"}
	// 这样赋值后 f1 叫做 方法表达式
	f1 := A.GetName
	f1(a)
	f2 := a.GetName
	f2()
}

func Five() {
	a := A{name: "17x in main"}
	f2 := a.GetName
	fmt.Println(f2()) // 17x in main

	f3 := GetFunc()
	fmt.Println(f3()) // 17x in GetFunc
}
func main() {
	Five()
}
func NameOfA(a A) string {
	a.name = "Hi " + a.name
	return a.name
}

func GetFunc() func() string {
	a := A{name: "17x in GetFunc"}
	return a.GetName
}
func GetFunc2() func() string {
	a := A{name: "17x in GetFunc"}

	return func() string {
		return A.GetName(a)
	}
}
