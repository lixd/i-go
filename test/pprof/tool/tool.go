package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	// record cpu info to file
	file, err := os.Create("./cpu.pprof")
	if err != nil {
		fmt.Printf("create cpu pprof failed, err:%v\n", err)
		return
	}

	if err := pprof.StartCPUProfile(file); err != nil {
		fmt.Printf("could not start CPU profile :%v\n", err)
		return
	}
	defer func() {
		pprof.StopCPUProfile()
		file.Close()
	}()

	for i := 0; i < 10; i++ {
		logic()
	}
	//  record memory info to file
	fileMem, err := os.Create("./mem.pprof")
	if err != nil {
		fmt.Printf("create mem pprof failed, err:%v\n", err)
		return
	}
	// runtime.GC()
	if err := pprof.WriteHeapProfile(fileMem); err != nil {
		fmt.Printf("could not start Heap profile :%v\n", err)
		return
	}
	fileMem.Close()
	// 指定获取某一项的 profile
	// goroutine、threadcreate、heap、allocs、block、mutex
	fileG, err := os.Create("./goroutine.pprof")
	if err != nil {
		fmt.Printf("create mem pprof failed, err:%v\n", err)
		return
	}
	err = pprof.Lookup("goroutine").WriteTo(fileG, 1)

}

// logic logic code with some bug for test
func logic() {
	// normal logic
	fmt.Println("logic")
	// bad logic loop
	for i := 0; i < 1000000000; i++ {

	}
}
