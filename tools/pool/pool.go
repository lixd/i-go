package pool

import (
	"io"

	"github.com/pkg/errors"
)

var (
	// ErrClosed 连接池已经关闭Error
	ErrClosed = errors.New("pool is closed")
	// ErrPoolFull 连接池已满
	ErrPoolFull = errors.New("connection pool is full")
	// ErrConnCreateLimit 连接创建数达到连接池最大限制
	ErrConnCreateLimit = errors.New("connection create limit")
)

// Pool 基本方法 只需要实现了 io.Closer 方法即可使用该连接池
type Pool interface {
	Get() (io.Closer, error)

	Put(io.Closer) error

	Close() error

	Len() int
}
