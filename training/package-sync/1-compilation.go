package main

import "fmt"

var counter int

// go tool compile -S 1-compilation.go
func main() {
	counter++
	fmt.Println(counter)
}
