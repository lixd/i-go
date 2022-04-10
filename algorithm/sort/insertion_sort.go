package sort

// InsertionSort 插入排序
/*
时间复杂度： O(n^2)
空间复杂度： O(1) 原地排序
稳定性：稳定
是否基于比较: 是
思想: 将数组分为已排序区间和未排序区间，逐渐将未排序区间元素取出并在已排序区间找到正确的位置插入
*/
func InsertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		j := i - 1
		// 从后往前遍历已排序区间 找到正确插入位置
		for ; j >= 0; j-- {
			// 降序则改成小于即可
			if arr[j] > v {
				// 因为是升序排列 所以把大于当前值的往后移
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = v
	}
	return arr
}

func InsertionSort2(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > v {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = v
	}
	return arr
}
