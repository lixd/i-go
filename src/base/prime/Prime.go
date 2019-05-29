package main

import (
	"fmt"
	"time"
)

//计算素数
func main() {
	//存放1-8000的数
	intChanp := make(chan int, 8000)
	//存放素数
	primeChanp := make(chan int, 2000)
	//标识退出
	exitChanp := make(chan bool, 4)
	start := time.Now().UnixNano()
	//1.开启一个协程向intChanp 写入数据
	go putNum(intChanp)
	//2.开启四个协程 从intChanp取数据并判断是否为素数 是则放入primeChanp
	for i := 0; i < 4; i++ {
		go primeNum(intChanp, primeChanp, exitChanp)
	}
	//3.
	go func() {
		// 主线程从exitChanp中取4个true 没有取到则阻塞
		for i := 0; i < 4; i++ {
			<-exitChanp
		}
		end := time.Now().UnixNano()
		fmt.Printf("time:%v", end-start)
		//最后关闭primeChanp
		close(primeChanp)
	}()

	//4.遍历primeChanp
	//for value := range primeChanp {
	//	fmt.Println(value)
	//}
	for {
		res, ok := <-primeChanp
		if !ok {
			break
		}
		fmt.Println("素数=", res)
	}
}

func putNum(intChanp chan int) {
	for i := 1; i <= 8000; i++ {
		intChanp <- i
	}
	//关闭 channel
	close(intChanp)
}
func primeNum(intChanp chan int, primeChanp chan int, exitChanp chan bool) {
	//使用for循环 取
	//var num int
	var flag bool
	for {
		num, ok := <-intChanp
		if !ok { //intChanp 取不到数据则退出
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明不是素数
				flag = false
				break
			}
		}
		//是素数则放入primeChanp
		if flag {
			primeChanp <- num
		}
	}
	fmt.Println("有一个协程完成了。。。")
	exitChanp <- true
}
