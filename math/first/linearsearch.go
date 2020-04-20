package main

import (
	"fmt"
	"i-go/utils"
)

func main() {
	const maxLen = 10000000
	var arrs = make([]int, 0, maxLen)
	for i := 0; i < maxLen; i++ {
		arrs = append(arrs, i)
	}
	search := linearSearch(maxLen-1, arrs)
	fmt.Println(search)

}

func linearSearch(key int, arrs []int) int {
	defer utils.Trace("linearSearch")()

	for i, v := range arrs {
		if v == key {
			return i
		}
	}
	return -1
}
