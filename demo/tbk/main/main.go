package main

import (
	"fmt"
	"i-go/demo/tbk/core"
	"math/rand"
	"time"
)

// start at 16:08
func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		intn := rand.Intn(300)
		if intn < 150 {
			intn = 150
		}
		fmt.Printf("下次请求时间:%vs \n", intn)
		select {
		case <-time.After(time.Second * time.Duration(intn)):
			core.ReLogin()
			code := core.ShareItem("615675531053")
			if code != 200 {
				fmt.Println("stop at :", time.Now().Unix())
				return
			}
		}
	}
}
