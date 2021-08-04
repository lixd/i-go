package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/health", health)
	server := &http.Server{
		Addr: ":8080",
	}
	fmt.Println("server startup...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}

func hello(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("hello dockerfile ! \n"))
}

func health(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte(time.Now().Format(time.RFC3339)))
}
