package channel

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	// "reflect"

	"i-go/tools/pool"
)

// Config 连接池相关配置
type Config struct {
	// 连接池中拥有的最小连接数
	InitialCap int
	// 最大并发存活连接数
	MaxCap int
	// 生成连接的方法
	Factory factory
}

type factory func() (io.Closer, error)

// channelPool 使用 channel 存放连接信息
type channelPool struct {
	mu      sync.Mutex
	conns   chan *idleConn
	closed  bool
	factory factory
	maxCap  int
	running int
}

type idleConn struct {
	conn io.Closer
	t    time.Time
}

// NewChannelPool 初始化连接
func NewChannelPool(conf Config) (pool.Pool, error) {
	if conf.InitialCap <= 0 {
		return nil, errors.New("invalid capacity settings")
	}
	if conf.Factory == nil {
		return nil, errors.New("invalid factory func settings")
	}

	c := &channelPool{
		conns:   make(chan *idleConn, conf.MaxCap),
		factory: conf.Factory,
		maxCap:  conf.MaxCap,
	}

	for i := 0; i < conf.InitialCap; i++ {
		conn, err := c.factory()
		if err != nil {
			return nil, fmt.Errorf("factory is not able to fill the pool: %s", err)
		}
		c.conns <- &idleConn{conn: conn, t: time.Now()}
	}

	return c, nil
}

// Get 从 pool 中取一个连接
func (c *channelPool) Get() (io.Closer, error) {
	if c.closed {
		return nil, pool.ErrClosed
	}
	for {
		select {
		// 优先从池里取
		case wrapConn := <-c.conns:
			if wrapConn == nil {
				return nil, pool.ErrClosed
			}
			c.running++
			return wrapConn.conn, nil
		// 	如果取不到直接建立一个新的连接
		default:
			if c.running >= c.maxCap { // 连接池满后不能再创建新连接
				return nil, pool.ErrConnCreateLimit
			}
			conn, err := c.factory()
			if err != nil {
				return nil, err
			}
			return conn, nil
		}
	}
}

// Put 将连接放回pool中
func (c *channelPool) Put(conn io.Closer) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return pool.ErrClosed
	}
	select {
	case c.conns <- &idleConn{conn: conn, t: time.Now()}:
		c.running--
		return nil
	default:
		return pool.ErrPoolFull
	}
}

// Close 关闭连接池同时关闭池子里的所有连接
func (c *channelPool) Close() error {
	if c.closed {
		return pool.ErrClosed
	}
	// 缩小锁范围，只把 closed 赋值语句锁起来
	c.mu.Lock()
	c.closed = true
	c.mu.Unlock()

	// 关闭通道，不让写入了
	close(c.conns)
	// 依次关闭所有连接
	for conn := range c.conns {
		_ = conn.conn.Close()
	}
	return nil
}

// Len 连接池中已有的连接
func (c *channelPool) Len() int {
	return len(c.conns)
}
