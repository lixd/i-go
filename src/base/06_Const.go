package main

import "fmt"

type user struct {
	age  int
	name string
}

func main() {
	const NAME string = "illusory"
	user := user{30, "illusory"}
	const (
		A = 10
		B
		C
		D = iota
		E = iota
	)
	fmt.Println(user, A, B, C, D, E)
}
