package sort

// SelectionSort 选择排序
/*
时间复杂度： O(n^2)
空间复杂度： O(1) 原地排序
稳定性： 不稳定
是否基于比较: 是
思想: 将数组分为已排序区间和未排序区间，每次从未排序区间选择最小值添加到已排序区间末尾
*/
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		// 从 i 开始寻找最小值
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// 最后 和 i 进行交换
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
