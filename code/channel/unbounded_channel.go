package channel

import (
	"sync"
)

// Origin grpc-go
/*
用于多个goroutine之间的消息传递。这是一个非常不错的channel实践，它不用考虑channel的各种阻塞情况（这里主要是channel溢出的情况。方便了channel的应用。
实现来自
/internal/buffer/unbounded.go Unbounded
/internal/transport/transport.go recvBuffer
这两者的实现逻辑是一样的，只是Unbounded包装的interface{} ，而recvBuffer会被高频调用所以使用了具体的类型recvMsg
*/
type Unbounded struct {
	c       chan interface{}
	backlog []interface{}
	sync.Mutex
}

func NewUnbounded() *Unbounded {
	return &Unbounded{c: make(chan interface{}, 1)}
}

// 往管道中写入消息（生产端
func (b *Unbounded) Put(t interface{}) {
	b.Lock()
	// 判断是否有积压消息，如果没有则直接写入管道后退出
	// 如果有，则写入到积压队列中（先进先出队列
	if len(b.backlog) == 0 {
		select {
		case b.c <- t:
			b.Unlock()
			return
		default:
		}
	}
	b.backlog = append(b.backlog, t)
	b.Unlock()
}

func (b *Unbounded) Load() {
	b.Lock()
	// 这里主要是判断积压队列是否有消息，如果有则左移一位
	// 并将移出的消息，写入channel中。
	if len(b.backlog) > 0 {
		select {
		case b.c <- b.backlog[0]:
			b.backlog[0] = nil
			b.backlog = b.backlog[1:]
		default:
		}
	}
	b.Unlock()
}

// 管道的读信号（消费端
func (b *Unbounded) Get() <-chan interface{} {
	return b.c
}
