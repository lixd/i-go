package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var s = Student{}
	// Speak 方法接收者为 *Student 所以只有 Student的指针才算实现了 People接口
	var peo People = &s
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
