package main

func sample() string {
	str := "hello" + "world"
	return str
}

func fib(i int) int {
	if i < 2 {
		return i
	}
	return fib(i-1) + fib(i-2)
}

func main() {
	sample()
	fib(10)
}

// 查看当前函数是否可以内联，以及不可以内联的原因：go tool compile -m=2 main.go
