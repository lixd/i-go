package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	originServerHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[origin server] received request at: %s\n", time.Now())
		_, _ = fmt.Fprint(rw, "origin server response")
	})

	log.Fatal(http.ListenAndServe(":8081", originServerHandler))
}

/*
情况一
客户端和目标服务器不再同一网络，无法直接反问。
反向代理服务有两张网卡，分别位于这两个网络，作为中转。
客户端发送请求给反向代理服务，反向代理服务将请求转发到目标服务，并将返回值转发给客户端。

情况二
该情况中，反向代理服务器和客户端处于同一网络，因此也不能直接访问目标服务器。
需要先让客户端访问反向代理服务器，反向代理服务器将该连接保存下来。
后续客户端发送请求给反向代理服务，反向代理服务使用之前保存的连接将请求转发到目标服务，并将返回值转发给客户端。
*/
