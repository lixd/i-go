package main

import "fmt"

func main() {
	a, b := 1, 2
	defer func(b int) {
		a = a + b
		fmt.Println(a, b) // 5 2
	}(b)
	a = a + b
	fmt.Println(a, b) // 3 2
}

// func main() {
// 	a, b := 1, 2
// 	defer func(a int) {
// 		fmt.Println(a) // 1
// 	}(a)
// 	a = a + b
// 	fmt.Println(a, b) // 3,2
// }
