package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFor(t *testing.T) {
	var i int
	for {
		rand.Seed(time.Now().UnixNano())
		i = rand.Intn(4)
		if i > 3 {
			fmt.Println(i)
			continue
		}
		if i > 2 {
			fmt.Println(i)
			continue
		}
		if i > 1 {
			fmt.Println(i)
			continue
		}
	}
}
