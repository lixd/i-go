package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	URL := "http://www.baidu.com/"
	for i := 0; i < 1; i++ {
		if i%100 == 0 {
			fmt.Printf("已请求%v次\n", i)
		}
		time.Sleep(time.Second * 1)
		resp, err := http.Get(URL)
		if err != nil {
			fmt.Print("Error: ", err)
			continue
		}
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
		_ = resp.Body.Close()
	}
}
