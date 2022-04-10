package sort

/*
时间复杂度： O(nlogn)
空间复杂度：
quickSort  步骤
1）选取一个 pivot（基准元素）
2）将数组分为两个子数组
	比 pivot 大的装到右侧数组
    比 pivot 小(或相等)的装到左则数组
3）对两个子数组分别递归调用 quickSort
4）返回结果为 左侧数组+pivot+右侧数组
*/
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0] // 暂时简单的把第一个元素作为基准元素
	var left, right []int
	for _, v := range arr[1:] {
		if v > pivot {
			right = append(right, v)
		} else if v <= pivot {
			left = append(left, v)
		}
	}
	// 最后对左右两边的数组分别递归调用 quickSort
	// 返回结果为 左侧数组+pivot+右侧数组
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)
}

func quickSort2(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else if v >= pivot {
			right = append(right, v)
		}
	}
	l := quickSort2(left)
	r := quickSort2(right)
	return append(append(l, pivot), r...)
}
