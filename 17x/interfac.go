package main

import "fmt"

type IUser interface {
	Greet()
	Run()
	GetName()
}
type Student struct {
	Name string
}

func (s Student) Greet() {
	fmt.Println("student hello~")
}
func (s Student) GetName() string {
	return fmt.Sprintf("s-%s", s.Name)
}
func (s Student) Run() {
	fmt.Printf("%s is running~\n", s.GetName())
}

type Teacher struct {
	Name string
}

func (t Teacher) Greet() {
	fmt.Println("teacher hello~")
}
func (t Teacher) GetName() string {
	return fmt.Sprintf("t-%s", t.Name)
}
func (t Teacher) Run() {
	fmt.Printf("%s is running~\n", t.GetName())
}

type Man struct {
	Student
	Name string
}

func (m Man) GetName() string {
	return fmt.Sprintf("t-%s", m.Name)
}

func main() {
	var man = Man{
		Student: Student{
			Name: "Student",
		},
		Name: "man",
	}
	man.Run()
}

type People interface {
	Speak(string) string
}
