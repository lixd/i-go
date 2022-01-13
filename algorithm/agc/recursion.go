package main

import "fmt"

func Recursion(n int) int {
	if n == 0 {
		return 1
	}

	return n * Recursion(n-1)
}

func main() {
	fmt.Println(Recursion(5))
}
