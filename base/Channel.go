package main

import (
	"fmt"
)

var MyChannel chan int = make(chan int, 10)
var EndChan chan int = make(chan int, 10)

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Send ", i)
	}
}
func recv(ch chan int) {
	for i := 0; i < 10; i++ {
		res := <-ch
		EndChan <- res
		fmt.Println("Recv ", res)
	}
}

func main() {
	go send(MyChannel)
	go recv(MyChannel)
	for i := 0; i < 10; i++ {
		<-EndChan
	}
}
