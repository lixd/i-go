package goadvanced

func Random(n int64) <-chan int64 {
	ch := make(chan int64)
	go func() {
		defer close(ch)
		for i := int64(0); i < n; i++ {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	return ch
}
