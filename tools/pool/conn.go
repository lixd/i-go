package pool

// Conn 具体连接资源
type Conn struct {
	Unix int64
}

func (c *Conn) Close() error {
	c.Unix = 0
	return nil
}
