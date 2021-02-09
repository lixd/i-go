package lcof

func fib(n int) int {
	var mod = 1000000007
	if n < 2 {
		return n
	}
	return (fib(n-1) + fib(n-2)) % mod
}

func fib2(n int) int {
	var mod = 1000000007
	if n < 2 {
		return n
	}
	var f, n1, n2 = 1, 0, 0
	for i := 2; i <= n; i++ {
		n2 = n1
		n1 = f
		f = (n1 + n2) % mod // 题目要求
	}
	return f
}
