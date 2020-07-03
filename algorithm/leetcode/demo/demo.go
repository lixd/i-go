package main

import "fmt"

func main() {
	haystack := "hello world"
	needle := "world"
	fmt.Printf("strStr:%v \n", strStr(haystack, needle))
	nums := []int{1, 2, 3, 4, 5}
	ints := subsets(nums)
	fmt.Printf("subsets:总个数:%v 列表:%v \n", len(ints), ints)
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	// i 只需要到最后一个能容纳 needle 的位置即可
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// 判断字符串长度是否相等
		if len(needle) == j {
			return i
		}
	}
	return -1
}

// subsets 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
func subsets(nums []int) [][]int {
	// 保存最终结果
	result := make([][]int, 0)
	// 保存中间结果
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result
}

// nums 给定的集合
// pos 下次添加到集合中的元素位置索引
// list 临时结果集合(每次需要复制保存)
// result 最终结果
func backtrack(nums []int, pos int, list []int, result *[][]int) {
	// 把临时结果复制出来保存到最终结果
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	// 选择、处理结果、再撤销选择
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}
