package model

type student struct {
	Name string
	age  int
}

//student首字母小写后 是私有的 只能在 model 包使用
//使用工厂模式来解决
func NewStudent(name string, age int) *student {
	return &student{name, age}
}

//如果字段是私有的 也可以提供一个获取字段的方法 类似Java中的Get方法
func (student *student) GetAge() int {
	return student.age
}
