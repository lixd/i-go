package main

import (
	"fmt"
	model2 "i-go/base/lessons/factory/model"
)

func main() {
	// 创建一个Student实例
	// student首字母小写后 是私有的 只能在 model 包使用
	// var stu=model.student{"illusory",22}
	// 返回的是指针
	var stu = model2.NewStudent("illusory", 22)
	fmt.Println(*stu)
	// 字段私有后也无法访问
	// fmt.Printf("Name %v Age %v",stu.Name,stu.Age)
	fmt.Printf("Name %v Age %v \n", stu.Name, stu.GetAge())

	var per = model2.NewPerson("Azz", 666)
	fmt.Printf("Name %v Salary %v \n", per.GetName(), per.GetSalary())
	per.SetName("Azzz")
	per.SetSalary(888)
	fmt.Printf("Name %v Salary %v \n", per.GetName(), per.GetSalary())
}
