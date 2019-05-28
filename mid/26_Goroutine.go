package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	// go testg()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(" main hello golang",strconv.Itoa(i))
	// }
	// cpu()
	// for i := 1; i <= 20; i++ {
	// 	testAdd(i)
	// }
	// lock.Lock()
	// for i, value := range myMap {
	// 	fmt.Printf("map[%d]=%d \n", i, value)
	// }
	// lock.Unlock()

	datachan := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		go Productor(datachan, rand.Intn(999), &wg)
		wg.Add(1)
	}
	for i := 0; i < 2; i++ {
		go Comnsumer(datachan, &wg)
		wg.Add(1)
	}
	wg.Wait()

}
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
