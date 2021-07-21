package main

import (
	"fmt"
	"reflect"
)

/*
https://blog.golang.org/laws-of-reflection
reflect 3大定律
1. Reflection goes from interface value to reflection object.
	interface{}类型的值到反射reflecton对象.
根源上来说， reflection的原理就是检查interface中保存的一对值和类型, 所以在reflect包中，有两个类型我们需要记住， Type和Value两个类型. 通过这两个类型，我们可以访问一个interface变量的内容. 调用reflect.ValueOf和reflect.TypeOf可以检索出一个interface的值和具体类型. 当然通过reflect.Value我们也可以获得reflect.Type。
2.Reflection goes from reflection object to interface value.
	反射reflection对象到interface{}类型的值.
通过reflect.Value的Interface方法，我们可以获得一个Interface值。实际上这个方法将一个type和value打包回interface
3.To modify a reflection object, the value must be settable.
	当修改一个反射reflection时, 其值必须是settable. settable类似于addressable但是更严格

汇编源码分析
https://blog.csdn.net/breaksoftware/article/details/85995767 上
https://blog.csdn.net/breaksoftware/article/details/86068788 下
*/

type order struct {
	ordId      int
	customerId int
}

func query(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
}

func createQuery(q interface{}) {
	// v := reflect.ValueOf(q).MethodByName("123").Call([]reflect.Value{reflect.ValueOf(q)})
	v := reflect.ValueOf(q)
	if v.Kind() == reflect.Struct {
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func call1() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	query(o)
	createQuery(o)
}

func settable() {
	// var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// v.SetFloat(7.1) // Error: will panic.

	var x float64 = 3.4
	v := reflect.ValueOf(&x) // give address
	v.Elem().SetFloat(7.1)   // Ok
}

func main() {
	// settable()

	// call1()

	// r := StudentR{
	// 	Name: "kernel",
	// 	Age:  12,
	// }
	// reflectCall(r)
	fmt.Println("-----------test01----------")

	var num = 100
	reflectTest01(num)

	fmt.Println("-----------test02----------")
	stu := StudentR{"illusory", 22}
	reflectTest02(stu)

	fmt.Println("---------test03------------")
	reflectTest03()

	fmt.Println("----------test04-----------")
	reflectTest04(stu)

	fmt.Println("----------test05-----------")
	reflectTest05(int64(1))

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

func reflectTest05(x interface{}) {
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

func reflectCall(x interface{}) {
	rt := reflect.ValueOf(x)
	name := rt.FieldByName("Name")
	age := rt.FieldByName("Age")
	fmt.Printf("name:%s age:%v \n", name, age)
	p := "参数1"
	// 需要参数就必须穿一个 不需要时可以传 nil
	rt.MethodByName("PrintWithParam").Call([]reflect.Value{reflect.ValueOf(p)})
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
func (student StudentR) PrintWithParam(p string) {
	fmt.Println("p:", p)
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

func reflectTest03() {
	var num = 100
	rVal := reflect.ValueOf(&num)
	rVal.Elem().SetInt(20)
	fmt.Printf("int=%d \n", num)
}

func reflectTest04(data interface{}) {
	typeOf := reflect.TypeOf(data)
	valueOf := reflect.ValueOf(data)

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
