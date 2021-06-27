package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type config struct {
	a []int
}

func (c *config) T() {
	fmt.Printf("%+v\n", c.a)
}

func main() {
	var v atomic.Value
	v.Store(&config{})
	go func() {
		i := 0
		for {
			i++
			cfg := &config{a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				cfg := v.Load().(*config)
				cfg.T()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
