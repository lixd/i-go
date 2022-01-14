package ratelimit

import (
	"fmt"
	"testing"
)

func Test_circuitBreaker(t *testing.T) {
	circuitBreaker()
}

func TestA(t *testing.T) {
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go count2(i, i+1, ch)
	}
	for {
		select {
		case <-ch:
			fmt.Print("ok")
		}
	}
}

func count2(a, b int, ch chan bool) {

}
