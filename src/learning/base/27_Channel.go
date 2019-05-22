package main

import (
	"fmt"
)

func main() {
	//1.声明管道
	var intChan chan int
	var intChant chan int
	intChant = make(chan int, 5)
	//只写 channel
	var writeChan chan<- int
	writeChan = make(chan int, 2)
	//只读 channel
	var readChan <-chan int
	readChan = make(chan int, 2)
	for i := 0; i < 5; i++ {
		intChant <- i
	}
	go testPanic()

lable:
	for {
		select {
		//这里 如果 intChan 一直没有关闭，不会一直阻塞二deadlock
		//会自动到下一个 case 匹配
		case v := <-intChan:
			fmt.Printf("取到的数据 %v \n", v)
		case v := <-intChant:
			fmt.Printf("取到的数据 %v \n", v)
		default:
			fmt.Printf("都取不到数据，程序员可以添加自己的逻辑")
			//return
			break lable
		}
	}

	fmt.Println(writeChan, readChan)
	intChan = make(chan int, 3)
	// 2.引用类型 值为地址 0xc000094080
	fmt.Println(intChan)

	//3.向管道写入数据
	intChan <- 10
	intChan <- 11
	intChan <- 12
	//注意：写入数据时不能超过其容量
	fmt.Printf("channel len%v cap %v \n", len(intChan), cap(intChan))

	//4.读取数据 读取后len减少 cap不变 可以继续存数据了
	var num int
	//希望等到第三个数据 直接把前两个推出
	<-intChan
	<-intChan
	num = <-intChan
	fmt.Println(num)

	fmt.Printf("channel len%v cap %v \n", len(intChan), cap(intChan))

	intChan2 := make(chan int, 3)
	intChan2 <- 100
	intChan2 <- 101
	//关闭 channel 不能再写数据 可以继续读取
	close(intChan2)

	intChan3 := make(chan int, 100)
	for i := 1; i <= 100; i++ {
		intChan3 <- i * 2
	}
	close(intChan3)
	for value := range intChan3 {
		fmt.Println(value)
	}

	var intChan4 = make(chan int, 50)
	var exitChan = make(chan bool, 1)

	go writeData(intChan4)
	go readData(intChan4, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("writeData:%v \n", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		data, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData:%v \n", data)
	}
	// 读取完数据后
	exitChan <- true
	close(exitChan)
}
func testPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("testPanic err:%v \n", err)
		}
	}()
	var panicMap map[int]string
	panicMap[0] = "golang"
}
