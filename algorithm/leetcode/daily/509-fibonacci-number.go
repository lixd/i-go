package daily

// https://leetcode-cn.com/problems/fibonacci-number/
func fib(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	// f=n1+n2 类似滑动窗口 不断更新这三个数就行了
	f, n1, n2 := 1, 0, 0
	for i := 2; i <= n; i++ {
		n2 = n1
		n1 = f
		f = n1 + n2
	}
	return f
}
