package recursion

// https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
func climbStairs2(n int) int {
	if n <= 3 {
		return n
	}
	f1, f2, f3 := 1, 2, 3
	// 不需要完整的 Fibonacci 数列，所以只需要保存中间 3 个变量即可
	for i := 3; i < n+1; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}
