package main

import (
	"fmt"
	"sync"
)

/*
Golang range 内部实现原理:https://juejin.cn/post/6844904016745365518
Go语言设计与实现：https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-for-range/#51-for-%E5%92%8C-range
*/
// 数组是值类型，for range 的时候会拷贝一份数组，所以原数组改变了for range 也不会改变
func arr() {
	var a = [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			// [100 200 3]
			fmt.Println(a)
		}
		fmt.Println(k, v)
		a[k] = 100 + v
	}
	// [101 102 103]
	fmt.Println(a)
}

// 切片是引用类型，，for range 拷贝之后底层指向的还是同一个数组，所以改变原切片(不扩容的情况下)底层数组也会改变，因此 for range 也会被改变
// 如果改变原数组的时候扩容了,那么for range 指向就旧数组就不会被改变
func slice() {
	var a = []int{1, 2, 3}
	fmt.Println(cap(a))
	for k, v := range a {
		if k == 0 {
			a = append(a, 4) // 如果加上这句呢
			a[0], a[1] = 100, 200
			// [100 200 3]
			fmt.Println(a)
		}
		fmt.Println(k, v)
		a[k] = 100 + v
	}
	// [101 300 103]
	fmt.Println(a)
}

// func main() {
// 	slice()
// 	arr()
// }

// 启动3个goroutine，按顺序打印cat、dog、fish各100次
/*
goroutine是无序的，所以需要用channel来控制顺序，这里用3个chan控制顺序,当对应chan能取到值的时候，才会打印
dog 取值后打印一次dog，然后写数据到 catCh，控制打印cat,然后cat打印之后写数据到 fishCh，控制打印fish，然后fish打印之后写数据到dog，形成一个循环
for循环控制只打印100次就退出
由于最后fish打印之后还会往dogCh写数据，所以需要dogCh增加一个缓冲,或者在打印fish的时候判断，如果是最后一次打印，那么就不写数据到dogCh
*/
func printStr(content string, count int, recv, send chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-recv
		fmt.Println(content)
		send <- struct{}{}
	}
}

func main() {
	var (
		dogCh  = make(chan struct{}, 1)
		catCh  = make(chan struct{})
		fishCh = make(chan struct{})
		count  = 100
		wg     sync.WaitGroup
	)
	wg.Add(3)
	go printStr("dog", count, dogCh, catCh, &wg)
	go printStr("cat", count, catCh, fishCh, &wg)
	go printStr("fish", count, fishCh, dogCh, &wg)
	dogCh <- struct{}{}
	wg.Wait()
}
