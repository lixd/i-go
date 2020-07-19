package array_linkedlist_skiplist

// https://leetcode-cn.com/problems/climbing-stairs
// 斐波那契数列
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var f1, f2, f3 = 1, 2, 3
	// 动态数组
	for i := 3; i < n+1; i++ {
		f3 = f2 + f1
		f1 = f2
		f2 = f3
	}
	return f3
}

func climbStairs2(n int) int {
	if n <= 3 {
		return n
	}
	var f1, f2, f3 = 1, 2, 3
	for i := 3; i < n+1; i++ {
		// f(n) = f(n-1) + f(n-2)
		f3 = f1 + f2
		// 只需要算最后的f(n) 所以这里每次更新 f(n-1)、f(n-2) 即可
		f1 = f2
		f2 = f3
	}
	return f3
}
