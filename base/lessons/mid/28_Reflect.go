package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var num int = 100
	// reflectTest01(num)
	//
	// fmt.Println("---------------------")
	// stu := StudentR{"illusory", 22}
	// reflectTest02(stu)
	//
	// fmt.Println("---------------------")
	// reflectTest03(num)
	//
	// fmt.Println("---------------------")
	// reflectTest04(stu)
	// ref(int64(1))
	r := StudentR{
		Name: "kernel",
		Age:  12,
	}
	refCall(r)
}
func reflectTest01(i interface{}) {

	// 获取传入变量的类型
	// 1.先获取到 reflect.Type
	rTyp := reflect.TypeOf(i)
	//
	fmt.Println(rTyp)
	// 2. 获取 reflect.Value
	rVal := reflect.ValueOf(i)
	fmt.Println(rVal)

	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println("iV ", iV)
	fmt.Println("num2 ", num2)
}
func ref(x interface{}) {
	of := reflect.TypeOf(x)
	switch of.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Invalid")

	}
}
func refCall(x interface{}) {
	name := reflect.ValueOf(x).FieldByName("Name") // Field 不存在直接返回 Invalid Value
	reflect.TypeOf(x).FieldByName("Name")          // 会返回两个值 可以判断 Field 是否存在
	fmt.Println(name)
	reflect.ValueOf(x).MethodByName("Print").Call(nil)
}

type StudentR struct {
	Name string `test:"name"`
	Age  int
}

func (student StudentR) Print() {
	fmt.Println("start")
	fmt.Printf("Name:%v Age:%v \n", student.Name, student.Age)
	fmt.Println("end")
}
func (student StudentR) GetSum(a int, b int) {
	fmt.Printf("sum: %v\n", a+b)
}
func reflectTest02(i interface{}) {
	// 获取传入变量的类型
	// 1.先获取到 reflect.Type
	rTyp := reflect.TypeOf(i)
	//
	fmt.Println(rTyp)
	// 2. 获取 reflect.Value
	rVal := reflect.ValueOf(i)
	fmt.Println(rVal)
	kind1 := rTyp.Kind()
	kind2 := rVal.Kind()
	fmt.Printf("rTyp kind %v rVal kind %v\n", kind1, kind2)
	iV := rVal.Interface()
	fmt.Printf("iV type %T value %v", iV, iV)
	stu2 := iV.(StudentR)
	fmt.Println("iV ", iV)
	fmt.Println("num2 ", stu2)
}

func reflectTest03(i interface{}) {
	var num int = 100
	rVal := reflect.ValueOf(&num)
	rVal.Elem().SetInt(20)
	fmt.Printf("int=%d \n", num)
}

func reflectTest04(i interface{}) {
	typeOf := reflect.TypeOf(i)
	valueOf := reflect.ValueOf(i)

	// 字段个数
	field := valueOf.NumField()

	for i := 0; i < field; i++ {
		fmt.Printf("字段 %d %v \n", i, typeOf.Field(i))
		get := typeOf.Field(i).Tag.Get("test")
		fmt.Printf("带Tag test 的 字段 %v", get)
	}
	// 方法个数
	method := valueOf.NumMethod()
	fmt.Printf("方法个数%v \n", method)
	// 获取第0个方法并调用 方法排序是根据字母排的 A-Z这样 和写的顺序无关
	// 一个两个方法 按照排序 其中GetSum 为第0个 Print 为第1个
	valueOf.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	valueOf.Method(0).Call(params)

}
