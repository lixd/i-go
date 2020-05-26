package main

import "fmt"

// 闭包
func main() {
	for i := 0; i < 10; i++ {
		add(i)
	}

	res := adder()

	for i := 0; i < 10; i++ {
		i2 := res(i)
		fmt.Printf("sum:%d \n", i2)

	}
	counter := counter()
	fmt.Println(counter())
}
func add(num int) int {
	//每次执行时sum都被初始化为0
	sum := 0
	sum += num
	fmt.Printf("sum:%d", sum)
	return sum
}

//使用闭包 内部包含一个匿名函数
/*
内部变量sum被res引用
然后adder()将res传递到外部
这样导致sum一直在被引用就不会被释放掉
于是形成了闭包
*/
func adder() func(int) int {
	sum := 0
	res := func(num int) int {
		sum += num
		return sum
	}
	return res
}

//闭包实现计数器
func counter() func() int {
	i := 0
	res := func() int {
		i++
		return i
	}
	//返回的是res那一段代码的内存地址
	return res
}
