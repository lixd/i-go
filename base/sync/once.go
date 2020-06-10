package main

import (
	"fmt"
	"sync"
	"unsafe"
)

// 协程安全 单例模式
type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

// 只执行一次
func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func main() {
	var wg sync.WaitGroup // 协程安全
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj)) // 输出的结果都是同一个地址
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
*/
