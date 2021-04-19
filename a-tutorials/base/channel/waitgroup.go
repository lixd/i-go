package channel

type generation struct {
	// 用于让等待者阻塞住的通道
	// 这个通道永远不会用于发送，只用于接收和close。
	wait chan struct{}
	// 计数器，标记需要等待执行完成的job数量
	n int
}

func newGeneration() generation {
	return generation{wait: make(chan struct{})}
}
func (g generation) end() {
	// close通道将释放因为接受通道而阻塞住的goroutine
	close(g.wait)
}

//这里我们使用一个通道来保护当前的generation。
//它基本上是WaitGroup状态的互斥量。
type WaitGroup chan generation

func NewWaitGroup() WaitGroup {
	wg := make(WaitGroup, 1)
	g := newGeneration()
	// 在一个新的WaitGroup上Wait, 因为计数器是0，会立即返回不会阻塞住线程
	// 它表现跟当前世代已经结束了一样, 所以这里先把世代里的wait通道close掉
	// 防止刚创建WaitGroup时调用Wait函数会阻塞线程
	g.end()
	wg <- g
	return wg
}

func (wg WaitGroup) Add(delta int) {
	// 获取当前的世代
	g := <-wg
	if g.n == 0 {
		// 计数器是0，创建一个新的世代
		g = newGeneration()
	}
	g.n += delta
	if g.n < 0 {
		// 跟sync库里的WaitGroup一样，不允许计数器为负数
		panic("negative WaitGroup count")
	}
	if g.n == 0 {
		// 计数器回到0了，关闭wait通道，被WaitGroup的Wait方法
		// 阻塞住的线程会被释放出来继续往下执行
		g.end()
	}
	// 将更新后的世代发送回WaitGroup通道
	wg <- g
}

func (wg WaitGroup) Done() { wg.Add(-1) }

func (wg WaitGroup) Wait() {
	// 获取当前的世代
	g := <-wg
	// 保存一个世代里wait通道的引用
	wait := g.wait
	// 将世代写回WaitGroup通道
	wg <- g
	// 接收世代里的wait通道
	// 因为wait通道里没有值，会把调用Wait方法的goroutine阻塞住
	// 直到WaitGroup的计数器回到0，wait通道被close后才会解除阻塞
	<-wait
}
