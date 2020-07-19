package maps

import "github.com/orcaman/concurrent-map"

type ConcurrentMapBenchmarkAdapter struct {
	cm cmap.ConcurrentMap
}

func NewConcurrentMapBenchmarkAdapter() *ConcurrentMapBenchmarkAdapter {
	concurrentMaps := cmap.New()
	return &ConcurrentMapBenchmarkAdapter{cm: concurrentMaps}
}

func (m *ConcurrentMapBenchmarkAdapter) Set(key interface{}, value interface{}) {
	m.cm.Set(key.(string), value)
}

func (m *ConcurrentMapBenchmarkAdapter) Get(key interface{}) (interface{}, bool) {
	return m.cm.Get(key.(string))
}

func (m *ConcurrentMapBenchmarkAdapter) Del(key interface{}) {
	m.cm.Remove(key.(string))
}
