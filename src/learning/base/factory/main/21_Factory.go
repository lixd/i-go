package main

import (
	"fmt"
	"learning/base/factory/model"
)

func main() {
	//创建一个Student实例
	//student首字母小写后 是私有的 只能在 model 包使用
	//var stu=model.student{"illusory",22}
	//返回的是指针
	var stu = model.NewStudent("illusory", 22)
	fmt.Println(*stu)
	//字段私有后也无法访问
	//fmt.Printf("Name %v Age %v",stu.Name,stu.Age)
	fmt.Printf("Name %v Age %v", stu.Name, stu.GetAge())
}
