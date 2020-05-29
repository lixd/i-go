package main

import (
	"fmt"
	"math"
)

func main() {
	first := "abcde"
	second := "abdeae"
	fmt.Println(oneEditAway(first, second))
}

func oneEditAway(first string, second string) bool {
	var (
		m     = make(map[byte]int)
		count int
	)
	for i := 0; i < len(first); i++ {
		m[first[i]]++
	}
	fmt.Println(m)
	for i := 0; i < len(second); i++ {
		m[second[i]]--
	}
	fmt.Println(m)

	for _, v := range m {
		c := math.Abs(float64(v))
		if c > 0 {
			count += int(c)
		}
		if count >= 2 {
			return false
		}
	}
	return true
}
