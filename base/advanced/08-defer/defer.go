package main

import "fmt"

// func main() {
// 	a, b := 1, 2
// 	defer func(b int) {
// 		a = a + b
// 		fmt.Println(a, b) // 5 2
// 	}(b)
// 	a = a + b
// 	fmt.Println(a, b) // 3 2
// }

// func main() {
// 	a, b := 1, 2
// 	defer func(a int) {
// 		fmt.Println(a) // 1
// 	}(a)
// 	a = a + b
// 	fmt.Println(a, b) // 3,2
// }

func AA(i int) {
	defer A1(i, 2*i)

	if i > 1 {
		defer A2("hello", "world")
	}

	return
}
func A1(a, b int) {
	fmt.Println(a, b)
}
func A2(m, n string) {
	fmt.Println(m, n)
}
func A() {
	defer B()
}
func B() {

}
