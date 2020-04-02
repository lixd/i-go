package simple

import "fmt"

func main() {
	arrs := []int{5, 3, 1, 4, 6, 8, 2}
	fmt.Print(arrs, "\n")
	bubbleSort(arrs)
	fmt.Print(arrs, "\n")

	arr := [11]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 23}
	fmt.Println("--------------------------")
	binarySerach(&arr, 0, len(arr)-1, 23)

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

//二分查找
func binarySerach(arr *[11]int, min int, max int, value int) {
	if min > max {
		fmt.Println("不存在")
		return
	}
	mid := (min + max) / 2
	if (*arr)[mid] > value {
		binarySerach(arr, min, mid-1, value)
	} else if (*arr)[mid] < value {
		binarySerach(arr, mid+1, max, value)
	} else {
		fmt.Printf("找到了 下标为 %d \n", mid)
	}
}
