package sort

// BubbleSort 冒泡排序
/*
时间复杂度： O(n^2)
空间复杂度： O(1) 原地排序
稳定性：稳定
是否基于比较: 是
思想: 双重 for loop 比较并交换位置，外层for loop 每完成一次都会排序好一个数字
优化1: 增加标志位 flag 在第i次排序如果没有元素交换（说明已经完成排序），则不进行第i+1次排序
优化2: TODO
*/
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := i; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}
