package daily

// https://leetcode-cn.com/problems/candy/
// https://leetcode-cn.com/problems/candy/solution/candy-cong-zuo-zhi-you-cong-you-zhi-zuo-qu-zui-da-/
func candy(ratings []int) (ans int) {
	n := len(ratings)
	// 从左往右遍历 找出满足条件时每个学生的糖果数
	left := make([]int, n)
	for i, r := range ratings {
		// 大于前面的则糖果数在前一个基础上+1
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			// 小于等于则先不处理，直接给1个糖果, 由后续从右往左遍历处理
			left[i] = 1
		}
	}
	// 从右往左再遍历一次 取左遍历、右遍历都满足的最大值
	right := 0
	for i := n - 1; i >= 0; i-- {
		// 同样的 大于就+1 这里不需要存储所有的值 所以不用数组 直接++
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		// 取满足左遍历、右遍历的最大值，就能同时满足
		ans += max(left[i], right)
	}
	return
}

func candy2(ratings []int) int {
	n := len(ratings)
	ans, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			// 大于说明是升序队列 把降序队列长度清0
			dec = 0
			if ratings[i] == ratings[i-1] {
				// 等于说明不是升序队列 糖果数置1从新开始
				pre = 1
			} else {
				// 否则糖果数++
				pre++
			}
			ans += pre
			inc = pre
		} else {
			// 小于则降序队列长度++
			dec++
			if dec == inc {
				// 同时注意当当前的递减序列长度和上一个递增序列等长时，需要把最近的递增序列的最后一个同学也并进递减序列中。？
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}
