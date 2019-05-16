package main

import "fmt"

func main() {
	iSum := addSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(iSum)
}

//不定参数 其实是一个切片
func addSum(data ...int) int {
	fmt.Printf("%T \n", data)
	result := 0
	for _, value := range data {
		result += value
	}
	return result
}
