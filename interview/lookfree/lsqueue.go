/**
 * 线程安全的队列，使用轻量级的 CAS 锁
 */
package lookfree

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

type esCache struct {
	putNo uint32
	getNo uint32
	value interface{}
}

// LFQueue lock free queue
type LFQueue struct {
	cap    uint32
	capMod uint32
	putPos uint32
	getPos uint32
	cache  []esCache
}

func NewQueue(cap uint32) *LFQueue {
	q := new(LFQueue)
	q.cap = minQuantity(cap)
	q.capMod = q.cap - 1
	q.putPos = 0
	q.getPos = 0
	q.cache = make([]esCache, q.cap)
	for i := range q.cache {
		cache := &q.cache[i]
		cache.getNo = uint32(i)
		cache.putNo = uint32(i)
	}
	cache := &q.cache[0]
	cache.getNo = q.cap
	cache.putNo = q.cap
	return q
}

func (q *LFQueue) String() string {
	getPos := atomic.LoadUint32(&q.getPos)
	putPos := atomic.LoadUint32(&q.putPos)
	return fmt.Sprintf("Queue{cap: %v, capMod: %v, putPos: %v, getPos: %v}",
		q.cap, q.capMod, putPos, getPos)
}

func (q *LFQueue) Cap() uint32 {
	return q.cap
}

func (q *LFQueue) Quantity() uint32 {
	var putPos, getPos uint32
	var quantity uint32
	getPos = atomic.LoadUint32(&q.getPos)
	putPos = atomic.LoadUint32(&q.putPos)

	if putPos >= getPos {
		quantity = putPos - getPos
	} else {
		quantity = q.capMod + (putPos - getPos)
	}

	return quantity
}

// Put put queue functions
func (q *LFQueue) Put(val interface{}) (ok bool, quantity uint32) {
	var putPos, putPosNew, getPos, posCnt uint32
	var cache *esCache
	capMod := q.capMod

	getPos = atomic.LoadUint32(&q.getPos)
	putPos = atomic.LoadUint32(&q.putPos)

	if putPos >= getPos {
		posCnt = putPos - getPos
	} else {
		posCnt = capMod + (putPos - getPos)
	}

	if posCnt >= capMod-1 {
		runtime.Gosched()
		return false, posCnt
	}

	putPosNew = putPos + 1
	if !atomic.CompareAndSwapUint32(&q.putPos, putPos, putPosNew) {
		runtime.Gosched()
		return false, posCnt
	}

	cache = &q.cache[putPosNew&capMod]

	for {
		getNo := atomic.LoadUint32(&cache.getNo)
		putNo := atomic.LoadUint32(&cache.putNo)
		if putPosNew == putNo && getNo == putNo {
			cache.value = val
			atomic.AddUint32(&cache.putNo, q.cap)
			return true, posCnt + 1
		} else {
			runtime.Gosched()
		}
	}
}

// Puts puts queue functions
func (q *LFQueue) Puts(values []interface{}) (puts, quantity uint32) {
	var putPos, putPosNew, getPos, posCnt, putCnt uint32
	capMod := q.capMod

	getPos = atomic.LoadUint32(&q.getPos)
	putPos = atomic.LoadUint32(&q.putPos)

	if putPos >= getPos {
		posCnt = putPos - getPos
	} else {
		posCnt = capMod + (putPos - getPos)
	}

	if posCnt >= capMod-1 {
		runtime.Gosched()
		return 0, posCnt
	}

	if capPuts, size := q.cap-posCnt, uint32(len(values)); capPuts >= size {
		putCnt = size
	} else {
		putCnt = capPuts
	}
	putPosNew = putPos + putCnt

	if !atomic.CompareAndSwapUint32(&q.putPos, putPos, putPosNew) {
		runtime.Gosched()
		return 0, posCnt
	}

	for posNew, v := putPos+1, uint32(0); v < putCnt; posNew, v = posNew+1, v+1 {
		var cache = &q.cache[posNew&capMod]
		for {
			getNo := atomic.LoadUint32(&cache.getNo)
			putNo := atomic.LoadUint32(&cache.putNo)
			if posNew == putNo && getNo == putNo {
				cache.value = values[v]
				atomic.AddUint32(&cache.putNo, q.cap)
				break
			} else {
				runtime.Gosched()
			}
		}
	}
	return putCnt, posCnt + putCnt
}

// Get get queue functions
func (q *LFQueue) Get() (val interface{}, ok bool, quantity uint32) {
	var putPos, getPos, getPosNew, posCnt uint32
	var cache *esCache
	capMod := q.capMod

	putPos = atomic.LoadUint32(&q.putPos)
	getPos = atomic.LoadUint32(&q.getPos)

	if putPos >= getPos {
		posCnt = putPos - getPos
	} else {
		posCnt = capMod + (putPos - getPos)
	}

	if posCnt < 1 {
		runtime.Gosched()
		return nil, false, posCnt
	}

	getPosNew = getPos + 1
	if !atomic.CompareAndSwapUint32(&q.getPos, getPos, getPosNew) {
		runtime.Gosched()
		return nil, false, posCnt
	}

	cache = &q.cache[getPosNew&capMod]

	for {
		getNo := atomic.LoadUint32(&cache.getNo)
		putNo := atomic.LoadUint32(&cache.putNo)
		if getPosNew == getNo && getNo == putNo-q.cap {
			val = cache.value
			cache.value = nil
			atomic.AddUint32(&cache.getNo, q.cap)
			return val, true, posCnt - 1
		} else {
			runtime.Gosched()
		}
	}
}

// Gets gets queue functions
func (q *LFQueue) Gets(values []interface{}) (gets, quantity uint32) {
	var putPos, getPos, getPosNew, posCnt, getCnt uint32
	capMod := q.capMod

	putPos = atomic.LoadUint32(&q.putPos)
	getPos = atomic.LoadUint32(&q.getPos)

	if putPos >= getPos {
		posCnt = putPos - getPos
	} else {
		posCnt = capMod + (putPos - getPos)
	}

	if posCnt < 1 {
		runtime.Gosched()
		return 0, posCnt
	}

	if size := uint32(len(values)); posCnt >= size {
		getCnt = size
	} else {
		getCnt = posCnt
	}
	getPosNew = getPos + getCnt

	if !atomic.CompareAndSwapUint32(&q.getPos, getPos, getPosNew) {
		runtime.Gosched()
		return 0, posCnt
	}

	for posNew, v := getPos+1, uint32(0); v < getCnt; posNew, v = posNew+1, v+1 {
		var cache *esCache = &q.cache[posNew&capMod]
		for {
			getNo := atomic.LoadUint32(&cache.getNo)
			putNo := atomic.LoadUint32(&cache.putNo)
			if posNew == getNo && getNo == putNo-q.cap {
				values[v] = cache.value
				cache.value = nil
				getNo = atomic.AddUint32(&cache.getNo, q.cap)
				break
			} else {
				runtime.Gosched()
			}
		}
	}

	return getCnt, posCnt - getCnt
}

// minQuantity 到最近的2的倍数
func minQuantity(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

func Delay(z int) {
	for x := z; x > 0; x-- {
	}
}
