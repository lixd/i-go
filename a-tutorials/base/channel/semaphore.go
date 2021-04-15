package channel

type Semaphore chan struct{}

func NewSemaphore(size int) Semaphore {
	return make(Semaphore, size)
}

func (s Semaphore) Lock() {
	s <- struct{}{}
}

func (s Semaphore) Unlock() {
	<-s
}
