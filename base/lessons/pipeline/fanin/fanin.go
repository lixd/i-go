package main

import (
	"fmt"
	"sync"
)

/*
举个例子，设计一个程序：计算一个整数切片中元素的平方值并把它打印出来。非并发的方式是使用for遍历整个切片，然后计算平方，打印结果。

我们使用流水线模型实现这个简单的功能，从流水线的角度，可以分为3个阶段：

1.遍历切片，这是生产者。
2.计算平方值。
3.打印结果，这是消费者。

流水线的特点
每个阶段把数据通过channel传递给下一个阶段。
每个阶段要创建1个goroutine和1个通道，这个goroutine向里面写数据，函数要返回这个通道。
有1个函数来组织流水线，我们例子中是main函数。
*/
func main() {
	in := producer(1, 2, 3, 4)
	c1 := square(in)
	c2 := square(in)
	c3 := square(in)
	for v := range merge(c1, c2, c3) {
		fmt.Println(v)
	}
}

// merge 从多个通道读取值 FAN-IN 模式
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	// 处理数据的 func
	collect := func(in <-chan int) {
		defer wg.Done()
		for v := range in {
			out <- v
		}
	}

	wg.Add(len(cs))
	for _, ch := range cs {
		go collect(ch)
	}

	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// producer 负责生产数据 它会把数据写入通道，并把它写数据的通道返回。
func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v
		}
	}()
	return out
}

// square 负责从某个通道读数字，然后计算平方，将结果写入通道，并把它的输出通道返回。
func square(inChan <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range inChan {
			out <- v * v
		}
	}()
	return out
}
