// BLock waiting to either receive from the goroutine 's
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// 本例中同时监听了 8080和 8001 端口，但是两个服务其中一个退出时整个程序都应该退出才行，现在很明显无法做到。
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	go http.ListenAndServe("127.0.0.1: 8001", http.DefaultServeMux) // debug
	http.ListenAndServe("0.0.0.0: 8080", mux)                       // app traffic

}

// 修改后如下
func main3() {
	go serveDebug()
	serveApp()
}

func serveApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	http.ListenAndServe("0.0.0.0: 8080", mux)
}

func serveDebug() {
	http.ListenAndServe("127.0.0.1: 8001", http.DefaultServeMux)
}

// 再次修改如下
func main4() {
	go serveDebug()
	go serveApp()
	select {}
}

func serveApp2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	if err := http.ListenAndServe("0.0.0.0: 8080", mux); err != nil {
		log.Fatal(err) // 通过log.Fatal(err)在出现错误时直接退出整个应用程序
	}
}

func serveDebug2() {
	if err := http.ListenAndServe("127.0.0.1: 8001", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

// 最终版 通过 chan 传值来实现优雅关闭
func mainLatest() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- serveDebug3(stop)
	}()

	go func() {
		done <- serveApp3(stop)
	}()
	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("error: %v\n", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-stop // wait for stop signal
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func serveApp3(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	return serve("0.0.0.0: 8080", mux, stop)
}

func serveDebug3(stop <-chan struct{}) error {
	return serve("127.0.0.1: 8001", http.DefaultServeMux, stop)
}
