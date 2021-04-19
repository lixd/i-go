package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	for i := 0; i < 10; i++ {
		m.Store(i, i)
	}
	fmt.Println(Len(m))
	fmt.Println(LenSimple(m))
}

func LenSimple(sm sync.Map) int {
	length := 0
	f := func(key, value interface{}) bool {
		length++
		return true
	}
	sm.Range(f)

	return length
}

func Len(sm sync.Map) int {
	length := 0
	f := func(key, value interface{}) bool {
		length++
		return true
	}
	one := length
	length = 0
	sm.Range(f)
	// 下面这段是干什么呢？
	if one != length {
		one = length
		length = 0
		sm.Range(f)
		if one < length {
			return length
		}

	}
	return one
}
