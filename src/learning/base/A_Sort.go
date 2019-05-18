package main

import "fmt"

func main() {
	arrs := []int{5, 3, 1, 4, 6, 8, 2}
	fmt.Print(arrs, "\n")
	bubbleSort(arrs)
	fmt.Print(arrs, "\n")

}
func bubbleSort(arr []int) {
	//定义标记 判断本轮循环是否有两两换位，如果没有则说明排序正常可以跳出循环
	flag := true
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				//出现换位 说明排序没有结束 需要继续
				flag = false
			}
		}
		if flag {
			break
		}
	}

}
