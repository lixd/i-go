package slice

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

// channelPool 使用 slice 存放连接信息
type channelPool struct {
	mu      sync.Mutex
	conns   []*idleConn
	closed  bool
	factory factory
	running int // 使用中的连接数
	maxCap  int // 最大连接数
}

type idleConn struct {
	conn io.Closer
	t    time.Time
}

// NewSlicePool 初始化连接
func NewSlicePool(conf Config) (pool.Pool, error) {
	if conf.InitialCap <= 0 {
		return nil, errors.New("invalid capacity settings")
	}
	if conf.Factory == nil {
		return nil, errors.New("invalid factory func settings")
	}

	c := &channelPool{
		conns:   make([]*idleConn, 0, conf.MaxCap),
		factory: conf.Factory,
		maxCap:  conf.MaxCap,
	}

	for i := 0; i < conf.InitialCap; i++ {
		conn, err := c.factory()
		if err != nil {
			return nil, fmt.Errorf("factory is not able to fill the pool: %s", err)
		}
		c.conns = append(c.conns, &idleConn{conn: conn, t: time.Now()})
	}

	return c, nil
}

// Get 从 pool 中取一个连接
func (c *channelPool) Get() (io.Closer, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return nil, pool.ErrClosed
	}

	if c.Len() > 0 { // 池子里有则直接提取
		ic := c.conns[0]
		c.conns = c.conns[1:]
		c.running++
		return ic.conn, nil
	}
	// 没有就创建一个新的
	if c.running >= c.maxCap { // 如果达到池子大小就不让创建了
		return nil, pool.ErrConnCreateLimit
	}
	conn, err := c.factory()
	if err != nil {
		return nil, err
	}
	c.running++
	return conn, nil
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
	if c.Len() >= c.maxCap {
		return pool.ErrPoolFull
	}
	c.running--
	c.conns = append(c.conns, &idleConn{conn: conn, t: time.Now()})
	return nil
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

	// 依次关闭所有连接
	for _, conn := range c.conns {
		_ = conn.conn.Close()
	}
	return nil
}

// Len 连接池中已有的连接
func (c *channelPool) Len() int {
	return len(c.conns)
}
