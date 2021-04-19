package main

import (
	"fmt"
	"sync"
	"unicode/utf8"
)

func main() {
	// helloMap()
	fmt.Println("Hello, 世界", len("世界"), utf8.RuneCountInString("世界"))
}

// https://www.jianshu.com/p/5bbe3a1cea61
func helloMap() {
	m := sync.Map{}
	m.Store("key", "value")
	load, ok := m.Load("key")
	if ok {
		fmt.Println(load)
	}
}
