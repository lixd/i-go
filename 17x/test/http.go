package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	URL := "http://img.lixueduan.com/redis/global-hash-table.png"
	for i := 0; i < 10000; i++ {
		if i%100 == 0 {
			fmt.Printf("已请求%v次\n", i)
		}
		time.Sleep(time.Second * 1)
		resp, err := http.Get(URL)
		if err != nil {
			fmt.Print("Error: ", err)
			continue
		}
		_ = resp.Body.Close()
	}
}
