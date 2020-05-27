package main

import (
	"fmt"
	"reflect"
)

// nil
//func main() {
//	var i interface{}
//	fmt.Printf("%#v \n",i)
//	if i == nil {
//		fmt.Println("The interface is nil.")
//	}
//}
// nil interface
//func main() {
//	var o *int = nil
//	var i interface{} = o
//	fmt.Printf("%#v \n",i)
//	if i == nil {
//		fmt.Println("The interface is nil.")
//	}
//}

// reflect检查nil接口
func main() {
	var o *int = nil
	var i interface{} = o
	var j interface{}
	fmt.Printf("i==nil:%t,j==nil:%t \n", i == nil, j == nil)
	v := reflect.ValueOf(i)
	if v.IsValid() {
		fmt.Println(v.IsNil())
	}
}
