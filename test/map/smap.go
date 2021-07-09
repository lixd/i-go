package main

import (
	"sync"
	"sync/atomic"
)

var smap sync.Map

func Inc(userID string, inc int64) {
	store, loaded := smap.LoadOrStore(userID, &inc)
	if loaded {
		count := store.(*int64)
		atomic.AddInt64(count, inc)
	}
}

func Collect() {
	smap.Range(func(key, value interface{}) bool {
		// xxx
		return true
	})
}

// go tool pprof -http=:8081 C:\Users\AIM\pprof\pprof.gw.samples.cpu.012.pb.gz

type localNew struct {
	m  map[string]*int64
	rw sync.RWMutex
}

func (l *localNew) inc(key string, inc int64) {
	l.rw.RLock()
	value := l.m[key]
	if value != nil { // 有值则+1后返回
		atomic.AddInt64(value, inc)
		l.rw.RUnlock()
		return
	}
	// 如果value为空则需要写入 这里需要加写锁
	// 由于不能锁升级(比如直接将读锁升级为写锁)，所以只能释放读锁后加写锁
	l.rw.RUnlock()
	l.rw.Lock()
	value2 := l.m[key] // double check 防止获取写锁这段时间里其他 goroutine 一直把值写入了
	if value2 != nil {
		atomic.AddInt64(value2, inc)
		l.rw.Unlock()
		return
	}
	l.m[key] = &inc
	l.rw.Unlock() // 释放写锁
}

func (l *localNew) scan() {
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
}
