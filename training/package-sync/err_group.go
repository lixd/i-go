package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// 并发调用多个RPC最容易想到的处理方法，用多个error来分别接收每个RPC的错误，最后通过chan将结果传递回去，这种方案能用但是不好用
/*func main() {
	var (
		a, b       int
		err1, err2 error
		ch         = make(chan result)
	)
	go func() {
		// call rpc1
		a, err1 = rpc()
		ch <- result{res: a, err: err1}
	}()
	go func() {
		// call rpc2
		b, err2 = rpc()
		ch <- result{res: b, err: err2}
	}()
}

func rpc() (int, error) {
	return 1, nil
}

type result struct {
	res int
	err error
}*/

// 使用 errgroup 对上述代码进行改造

func main() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Successfully fetched all URLs.")
}
