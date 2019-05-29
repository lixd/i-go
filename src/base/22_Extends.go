package main

import "fmt"

func main() {
	// 初始化时Student3的Name和Age都没赋值 所以是默认值
	var pupil = Pupil{Student3{}, "Azz", "一班"}
	// Name 字段pupil和Student3都有 所以按照就近原则 访问的是pupil中的Name字段 为"Azz"
	// Age 字段 pupil中没有 那么这里就是Student3中的 Age 未初始化 则是默认值 0
	// Class 字段也是pupil中的
	fmt.Printf("Name:%v Age:%v Class:%v \n", pupil.Name, pupil.Age, pupil.Class)
	// 调用的是pupil的read方法
	pupil.read()
	// 调用的是Student3的read方法 其中Student3的Name字段没有初始化为默认值""所以打印出来为""
	pupil.Student3.read()

	var c = C{A{"A", 1}, B{"B", 2}}
	// 错误 由于C中没有字段Name 然后A B 中都要 编译器不知道到底访问哪个 所以报错 如果C中有就会直接访问C中的
	// 这个规则对方法也是一样
	// fmt.Println(c.Name)
	fmt.Println(c.A.Name)
	fmt.Println(c.B.Name)
	var d = D{A{"A", 1}}
	//编译报错
	// fmt.Println(d.Name)
	//组合关系 必须这样写 带上有名结构体名字
	fmt.Println(d.a.Name)
	var e = E{A{"A", 1}, 20}
	fmt.Printf("A.Name %v A.Age %v e.int %v", e.Name, e.Age, e.int)

}

type Student3 struct {
	Name string
	Age  int
}

type Pupil struct {
	Student3
	Name  string
	Class string
}

func (student3 *Student3) read() {
	fmt.Printf("Student3中的方法：%v 在读书 \n", student3.Name)
}
func (pupil *Pupil) read() {
	fmt.Printf("pupil中的方法：%v 在读书 \n", pupil.Name)
}

type A struct {
	// A 中有字段Name
	Name string
	Age  int
}
type B struct {
	// B 中也有字段Name
	Name string
	Id   int
}
type C struct {
	A
	B
	// C 包含A B 然后C 中没有字段Name
	// Name string
}

type D struct {
	a A
}
type E struct {
	A
	int
}
