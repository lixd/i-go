package channel

type Once chan struct{}

func NewOnce() Once {
	o := make(Once, 1)
	o <- struct{}{}
	return o
}

func (o Once) Do(f func()) {
	_, ok := <-o
	if !ok {
		return
	}
	f()
	close(o)
}
