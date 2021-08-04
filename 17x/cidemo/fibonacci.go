package main

// 随变写了一个方法 用于跑 go test

// fibonacci 0 1 1 2 3 5 8
func fibonacci(n int64) int64 {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
