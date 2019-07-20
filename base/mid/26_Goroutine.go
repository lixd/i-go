package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		// 注意 这里的i是通过参数传递进去的
		// 虽然不写 里面也可以直接使用外面for循环的i但是
		// 这样会形成闭包 即 fun中的i就是for循环的i
		// for循环出i加到10之后退出循环执行tiem.sleep
		// 此时goroutine还是执行 a[10]++ 会出错
		go func(i int) {
			// goroutine是非抢占式的
			// 但是 这里 a[i]++ 无法交出控制权
			// 会导致一直阻塞在当前goroutine
			// 无法执行到main方法中的其他代码
			a[i]++
			// Gosched 手动交出控制权
			runtime.Gosched()
		}(i)
	}
	time.Sleep(time.Millisecond)
}

// func main() {
// 	// go testg()
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println(" main hello golang",strconv.Itoa(i))
// 	// }
// 	// cpu()
// 	// for i := 1; i <= 20; i++ {
// 	// 	testAdd(i)
// 	// }
// 	// lock.Lock()
// 	// for i, value := range myMap {
// 	// 	fmt.Printf("map[%d]=%d \n", i, value)
// 	// }
// 	// lock.Unlock()
//
// 	datachan := make(chan int, 100)
// 	var wg sync.WaitGroup
// 	for i := 0; i < 2; i++ {
// 		go Productor(datachan, rand.Intn(999), &wg)
// 		wg.Add(1)
// 	}
// 	for i := 0; i < 2; i++ {
// 		go Comnsumer(datachan, &wg)
// 		wg.Add(1)
// 	}
// 	wg.Wait()
//
// }
func Productor(datachan chan int, data int, wg *sync.WaitGroup) {
	for {
		datachan <- data
		fmt.Printf("Productor data= %d \n", data)
		time.Sleep(time.Second)
	}
	wg.Done()
}
func Comnsumer(datachan chan int, wg *sync.WaitGroup) {
	for {
		data := <-datachan
		fmt.Printf("Consumer data %d \n", data)
		time.Sleep(time.Second)
	}
	wg.Done()
}

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局互斥锁
	lock sync.Mutex
)

// 阶乘
func testAdd(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func testg() {
	for i := 0; i < 10; i++ {
		fmt.Println("test hello world", strconv.Itoa(i))
	}
}

func cpu() {
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
}
