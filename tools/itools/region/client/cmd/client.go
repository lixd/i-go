package main

import (
	"fmt"
	"time"

	region "i-go/tools/itools/region/client"
)

func main() {
	bec()
}
func bec() {
	defer func() func() {
		start := time.Now()
		return func() {
			fmt.Println(time.Since(start))
		}
	}()()

	for i := 0; i < 1; i++ {
		_ = region.Ip2Region("183.69.228.166")
	}
}
