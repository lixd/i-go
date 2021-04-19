package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	value := atomic.Value{}
	value.Store("x")
	x := value.Load()
	fmt.Println(x)
}
