package algorithm

import (
	"github.com/axiomhq/hyperloglog"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type hyper struct {
	hll  *hyperloglog.Sketch
	once sync.Once
}

func NewHyperLL() *hyper {
	return &hyper{hll: hyperloglog.New16()}
}

func (h *hyper) PFAdd(key string) bool {
	h.once.Do(func() {
		go h.expire()
	})
	return h.hll.Insert([]byte(key))
}

func (h *hyper) PFCount() int {
	return int(h.hll.Estimate())
}

func (h *hyper) expire() {
	for {
		time.Sleep(time.Hour)
		hll := hyperloglog.New16()
		addr := unsafe.Pointer(h.hll)
		news := unsafe.Pointer(hll)
		atomic.CompareAndSwapPointer(&addr, addr, news)
		h.hll = (*hyperloglog.Sketch)(addr)
	}
}
