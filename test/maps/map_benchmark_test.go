package maps

import (
	"strconv"
	"sync"
	"testing"
)

const (
	NumOfReader = 100
	NumOfWriter = 10
)

type Map interface {
	Set(key interface{}, val interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}

func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < NumOfWriter; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Set(strconv.Itoa(i), i*i)
					hm.Set(strconv.Itoa(i), i*i)
					hm.Del(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		for i := 0; i < NumOfReader; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Get(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

// 读多写少的情况下 Sync.Map 性能好
// 否则 concurrent map 性能好
func BenchmarkSyncMap(b *testing.B) {
	b.Run("map with RWLock", func(b *testing.B) {
		hm := NewRWLockMap()
		benchmarkMap(b, hm)
	})

	b.Run("sync.map", func(b *testing.B) {
		hm := NewSyncMapBenchmarkAdapter()
		benchmarkMap(b, hm)
	})

	b.Run("concurrent map", func(b *testing.B) {
		superman := NewConcurrentMapBenchmarkAdapter()
		benchmarkMap(b, superman)
	})
}
