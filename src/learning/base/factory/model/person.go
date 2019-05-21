package model

import "fmt"

type person struct {
	name   string
	salary float64
}

func NewPerson(name string, salary float64) *person {
	return &person{name, salary}
}

func (person *person) GetName() string {
	return person.name
}
func (person *person) SetName(name string) {
	person.name = name
}
func (person *person) GetSalary() float64 {
	return person.salary
}
func (person *person) SetSalary(salary float64) {
	if salary > 0 {
		person.salary = salary
	} else {
		fmt.Println("Salary 范围不正确...")
	}
}
