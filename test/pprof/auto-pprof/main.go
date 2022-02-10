package main

import (
	"net/http"
	"time"

	"github.com/mosn/holmes"
)

// 相关笔记-pprof 自动采样:https://github.com/lixd/daily-notes/blob/master/Golang/%E8%BF%9B%E9%98%B6/%E5%9F%BA%E7%A1%80%E5%BA%93/pprof/03-pprof%E8%87%AA%E5%8A%A8%E9%87%87%E6%A0%B7dump.md
func init() {
	http.HandleFunc("/make1gb", make1gbSlice)
	go http.ListenAndServe(":10003", nil)
}

/*
多次请求 make1gb 后能看到打印的日志和dump信息
[2022-02-07 12:38:08.148][Holmes] pprof mem, config_min : 3, config_diff : 25, config_abs : 80, config_max : 0, previous : [12 12 13 13 13 13 22 12 12 12], current: 22
[2022-02-07 12:38:08.150]heap profile: 5: 2147485136 [21: 16106129168] @ heap/1048576
2: 2147483648 [15: 16106127360] @ 0x78b1c8 0x73cdaf 0x73e6a9 0x73f9db 0x73c108 0x5c7561
#	0x78b1c7	main.make1gbSlice+0x27			D:/lillusory/projects/i-go/test/pprof/auto_pprof/main.go:28
#	0x73cdae	net/http.HandlerFunc.ServeHTTP+0x2e	C:/Go/src/net/http/server.go:2046
#	0x73e6a8	net/http.(*ServeMux).ServeHTTP+0x148	C:/Go/src/net/http/server.go:2424
#	0x73f9da	net/http.serverHandler.ServeHTTP+0x43a	C:/Go/src/net/http/server.go:2878
#	0x73c107	net/http.(*conn).serve+0xb07		C:/Go/src/net/http/server.go:1929
*/
func main() {
	h, _ := holmes.New(
		holmes.WithCollectInterval("2s"),
		holmes.WithCoolDown("1m"),
		holmes.WithDumpPath("./tmp"),
		holmes.WithTextDump(),
		holmes.WithMemDump(3, 25, 80),
	)
	h.EnableCPUDump()
	h.EnableThreadDump()
	h.EnableGoroutineDump()
	h.EnableMemDump().Start()
	time.Sleep(time.Hour)
}

func make1gbSlice(wr http.ResponseWriter, req *http.Request) {
	var a = make([]byte, 1073741824)
	_ = a
}
