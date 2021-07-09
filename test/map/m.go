package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

type local struct {
	m  map[string]*int64
	rw sync.RWMutex
}

func Demo() {
	var Local = &local{
		m:  make(map[string]*int64, 1000),
		rw: sync.RWMutex{},
	}
	var count int64
	var wg sync.WaitGroup
	for i := 0; i < 100_0000; i++ {
		wg.Add(1)
		go Local.inc(strconv.Itoa(i), &wg)
		if i%10_0000 == 0 {
			wg.Add(1)
			Local.scan(&wg, &count)
		}
	}
	wg.Wait()
	wg.Add(1)
	Local.scan(&wg, &count)
	fmt.Println("未扫描Key数:", len(Local.m))
	fmt.Println("扫描到的Key数:", count)
}

func deepCopy(src map[string]*int64, dest map[string]int64) {
	for k, v := range src {
		dest[k] = *v
	}
}

func (l *local) inc(key string, wg *sync.WaitGroup) {
	defer wg.Done()
	l.rw.RLock()
	value := l.m[key]
	if value != nil { // 有值则+1后返回
		atomic.AddInt64(value, 1)
		l.rw.RUnlock()
	}
	// 如果value为空则需要写入 释放读锁后加写锁
	l.rw.RUnlock()
	l.rw.Lock()
	r := int64(1)
	l.m[key] = &r
	l.rw.Unlock() // 释放写锁
}

func (l *local) scan(wg *sync.WaitGroup, count *int64) {
	defer wg.Done()
	// 扫描之前加写锁
	l.rw.Lock()
	// 然后复制map
	cp := make(map[string]int64)
	deepCopy(l.m, cp)
	// 在把原来的map重新复制
	ne := make(map[string]*int64, 1000)
	l.m = ne
	l.rw.Unlock() // 直接解锁
	// 	后续慢慢的遍历扫描即可
	atomic.AddInt64(count, int64(len(cp)))
}
