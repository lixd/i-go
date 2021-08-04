package slice

import (
	"fmt"
	"io"
	"testing"
	"time"

	"i-go/tools/pool"
)

func TestNewSlicePool(t *testing.T) {
	pc := Config{
		InitialCap: 1,
		MaxCap:     1,
		Factory: func() (io.Closer, error) {
			conn := &pool.Conn{
				Unix: time.Now().UnixNano(),
			}
			return conn, nil
		},
	}
	p, err := NewSlicePool(pc)
	if err != nil {
		t.Fatal("NewSlicePool: ", err)
	}
	c1, err := p.Get()
	if err != nil {
		t.Fatal("Get c1:", err)
	}
	fmt.Println(c1.(*pool.Conn).Unix)
	time.Sleep(time.Second)
	c2, err := p.Get()
	if err != nil {
		t.Fatal("Get c2:", err)
	}
	fmt.Println(c2.(*pool.Conn).Unix)
	_ = p.Put(c1)
	_ = p.Put(c2)
	c3, err := p.Get()
	if err != nil {
		t.Fatal("Get c3:", err)
	}
	fmt.Println(c3.(*pool.Conn).Unix)
	c4, err := p.Get()
	if err != nil {
		t.Fatal("Get c4:", err)
	}
	fmt.Println(c4.(*pool.Conn).Unix)
	_ = p.Close()
}
