package main

import (
	"fmt"
	"log"
	"net/http"
)

// Bad
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, GopherCon SG")
	})
	// main goroutine 无法感知 go func() 退出与否
	go func() {
		if err := http.ListenAndServe(" :8080", nil); err != nil {
			log.Fatal(err)
		}
	}()
	select {}
}

// Good
func main2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, GopherCon SG")
	})
	// 直接由 main goroutine 自己处理，可以省去很多麻烦
	if err := http.ListenAndServe(" :8080", nil); err != nil {
		log.Fatal(err)
	}
}

// bad
func main1() {
	doSomething()
	select {}
}
func doSomething() {
	go func() {
		// 	doSomething
	}()
}

// Good
func mainGood() {
	go doSomething2()
	select {}
}
func doSomething2() {
	// 	doSomething
}
