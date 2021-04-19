package main

import (
	"fmt"
)

func main() {
	simpleErr()
}

func simpleErr() {
	errorf := fmt.Errorf("X:%v ", "s")
	fmt.Println(errorf)
}
