package main

import (
	"fmt"
	"math/rand"
)

func main() {
	svc := []Server{"s1", "s2", "s3"}
	lb := NewLBRand(svc)
	for i := 0; i < 10; i++ {
		s := lb()
		fmt.Println(s)
	}
}

type Server string
type LB func() Server

func NewLB(svc []Server) LB {
	var i int
	return func() Server {
		i++
		if i >= len(svc) {
			i = 0
		}
		return svc[i]
	}
}

func NewLBRand(svc []Server) LB {
	return func() Server {
		return svc[rand.Int63()%int64(len(svc))]
	}
}
