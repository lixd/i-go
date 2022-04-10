package sort

import "math"

// SelectionSort 选择排序
/*
时间复杂度： O(n^2) 每个只能排序一个元素所以要遍历n次，每次遍历都是O(n) 所以最终为 O(n^2) 当然了每次遍历的n是递减的，准确值为O(1/2n^2) 不过一般复杂度都是省略系数的
空间复杂度： O(1) 原地排序
稳定性： 不稳定
是否基于比较: 是
思想: 将数组分为已排序区间和未排序区间，每次从未排序区间选择最小值添加到已排序区间末尾
	第一次时，已排序区间为空，全是未排序的，然后从未排序中找出最小值，放入已排序区间
	第二次，同样从未排序区间找最小值，放入已排序区间（放入尾部）
	直到未排序区间为空。
具体实现：具体实现上可以不需要额外的空间，直接在原数组上交换即可，第一个招到的值就和数组第一个元素交换，以此类推。
*/

func SelectionSort1(arr []int) []int {
	// 最容易理解的思路 用额外数组来接收
	var ret = make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		// 选出最小值并存入 ret 中
		idx := Minimum(arr)
		ret = append(ret, arr[idx])
		// 从 arr 中将本次选中的元素移除
		// NOTE: 这里有大量元素拷贝 会拖慢运行效率
		arr = append(arr[:idx], arr[idx+1:]...)
	}
	return arr
}

func Minimum(arr []int) int {
	var (
		min      = math.MaxInt
		minIndex int
	)
	for i, v := range arr {
		if v < min {
			min = v
			minIndex = i
		}
	}
	return minIndex
}

// SelectionSort 选择排序-优化版本 原地排序 不借助外部数组
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		// 每轮都从 i 开始寻找最小值 因为i之前的排好序了
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 最后将最小值和 i 进行交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func SelectionSort2(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
