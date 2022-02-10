package main

import (
	"net/http"
	// _ "net/http/pprof"
	"time"

	"github.com/pyroscope-io/client/pyroscope"
)

// 启动Server：docker run -it -p 4040:4040 pyroscope/pyroscope:latest server
// 启动后pyroscope会采集自身Server端程序的运行信息,浏览器访问http://localhost:4040即可查看
// 采集别的程序只需要通过agent采集程序数据并推送到pyroscope服务器即可
// 也就是在程序中调用pyroscope.Start()启动一个Agent。
func main() {
	push()
	// pull()
}

func pull() {
	// 拉模式只需要对外提供/debug/pprof/路由即可，在pyroscope Server端配置抓取数据
	go func() {
		_ = http.ListenAndServe("localhost:8180", nil)
	}()
	// your code goes here
	doSomething()
}

func push() {
	pyroscope.Start(pyroscope.Config{
		ApplicationName: "simple.golang.app2",
		// replace this with the address of pyroscope server
		ServerAddress: "http://localhost:4040",
		SampleRate:    100, // 采样率，1~100，默认100
		// you can disable logging by setting this to nil
		Logger: pyroscope.StandardLogger,
		// by default all profilers are enabled,
		// but you can select the ones you want to use:
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
		DisableGCRuns: false,
	})

	// your code goes here
	doSomething()
}

func doSomething() {
	ticker := time.NewTicker(time.Millisecond * 200)
	for range ticker.C {
		fib35()
		fib34()
		fib33()
	}
}

func fib35() {
	fib(35)
}
func fib34() {
	fib(34)
}
func fib33() {
	fib(33)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
