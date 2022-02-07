package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	go func() {
		for { // for循环用于消耗CPU
		}
	}()
	_ = make([]int64, 123456789) // 用于消耗内存

	p, _ := process.NewProcess(int32(os.Getpid()))
	cpuPercent, _ := p.Percent(time.Second)
	cp := cpuPercent / float64(runtime.NumCPU())
	// 获取进程占用内存的比例
	mp, _ := p.MemoryPercent()
	// 创建的线程数
	threadCount := pprof.Lookup("threadcreate").Count()
	// Goroutine数
	gNum := runtime.NumGoroutine()
	// output: cpuPercentTotal: 96.93539875444107 cpuPercentSingle:24.233849688610267 mp:0.4705571 threadCount:7 gNum:2
	fmt.Printf("cpuPercentTotal: %v cpuPercentSingle:%v mp:%v threadCount:%v gNum:%v\n", cpuPercent, cp, mp, threadCount, gNum)
}
