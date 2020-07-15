package channel

type RWMutex struct {
	write   chan struct{}
	readers chan int
}

func NewLock() RWMutex {
	return RWMutex{
		// 用来做一个普通的互斥锁
		write: make(chan struct{}, 1),
		// 用来保护读锁的数量，获取读锁时通过接受通道里的值确保
		// 其他goroutine不会在同一时间更改读锁的数量。
		readers: make(chan int, 1),
	}
}

func (l RWMutex) Lock()   { l.write <- struct{}{} }
func (l RWMutex) Unlock() { <-l.write }

func (l RWMutex) RLock() {
	// 统计当前读锁的数量，默认为0
	var rs int
	select {
	case l.write <- struct{}{}:
	// 如果write通道能发送成功，证明现在没有读锁
	// 向write通道发送一个值，防止出现并发的读-写
	case rs = <-l.readers:
		// 能从通道里接收到值，证明RWMutex上已经有读锁了，下面会更新读锁数量
	}
	// 如果执行了l.write <- struct{}{}, rs的值会是0
	rs++
	// 更新RWMutex读锁数量
	l.readers <- rs
}

func (l RWMutex) RUnlock() {
	// 读出读锁数量然后减一
	rs := <-l.readers
	rs--
	// 如果释放后读锁的数量变为0了，抽空write通道，让write通道变为可用
	if rs == 0 {
		<-l.write
		return
	}
	// 如果释放后读锁的数量减一后不是0，把新的读锁数量发送给readers通道
	l.readers <- rs
}
