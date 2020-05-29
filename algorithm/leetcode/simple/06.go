package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		select {
		case <-time.After(time.Second * 1):
			fmt.Println(1)
		}
	}
}
