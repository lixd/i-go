package sshutils

import (
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

/*
封装 Conn 对象以实现超时控制

调用 Read/Write 之前通过 SetReadDeadline 设置超时时间，如果指定时间内没能成功 Read/Write 则会超时返回
但是调用 SSHDialTimeout 时会启动后台 goroutine,每两秒发送一次保活消息，因此只要服务器正常（且超时时间大于两秒），连续就会一直保持。
因此不用担心执行耗时命令时触发超时导致被 kill 掉的问题。
这个超时时间只对服务器无响应的情况生效。
默认情况下，服务器无响应后执行命令会一直阻塞，永不退出。
使用封装后的 Conn ，就算有后台保活，但是服务器无响应后会直接退出保活 goroutine,因此达到超时时间后就会退出。
*/

// Conn wraps a net.Conn, and sets a deadline for every read
// and write operation.
type Conn struct {
	net.Conn
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (c *Conn) Read(b []byte) (int, error) {
	err := c.Conn.SetReadDeadline(time.Now().Add(c.ReadTimeout))
	if err != nil {
		return 0, err
	}
	return c.Conn.Read(b)
}

func (c *Conn) Write(b []byte) (int, error) {
	err := c.Conn.SetWriteDeadline(time.Now().Add(c.WriteTimeout))
	if err != nil {
		return 0, err
	}
	return c.Conn.Write(b)
}

func SSHDialTimeout(network, addr string, config *ssh.ClientConfig, timeout time.Duration) (*ssh.Client, error) {
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}

	timeoutConn := &Conn{conn, timeout, timeout}
	c, chans, reqs, err := ssh.NewClientConn(timeoutConn, addr, config)
	if err != nil {
		return nil, err
	}
	client := ssh.NewClient(c, chans, reqs)

	// this sends keepalive packets every 2 seconds
	// there's no useful response from these, so we can just abort if there's an error
	go func() {
		t := time.NewTicker(2 * time.Second)
		defer t.Stop()
		for range t.C {
			_, _, err := client.Conn.SendRequest("keepalive@golang.org", true, nil)
			if err != nil {
				return
			}
		}
	}()
	return client, nil
}
