package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	// 设置SSH配置
	config := &ssh.ClientConfig{
		// 服务器用户名
		User: "root",
		Auth: []ssh.AuthMethod{
			// 服务器密码
			ssh.Password("root"),
		},
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// 设置本地监听器，格式：地址:端口
	localListener, err := net.Listen("tcp", "localhost:13306")
	if err != nil {
		fmt.Printf("net.Listen failed: %v\n", err)
	}

	for {
		localConn, err := localListener.Accept()
		if err != nil {
			fmt.Printf("localListener.Accept failed: %v\n", err)
		}
		go forward(localConn, config)
	}
}

// 转发
func forward(localConn net.Conn, config *ssh.ClientConfig) {
	// 设置服务器地址，格式：地址:端口
	sshClientConn, err := ssh.Dial("tcp", "172.20.149.53:22", config)
	if err != nil {
		fmt.Printf("ssh.Dial failed: %s", err)
	}

	// 设置远程地址，格式：地址:端口（请在服务器通过 ifconfig 查看地址）
	sshConn, err := sshClientConn.Dial("tcp", "172.20.149.53:22")

	// 将localConn.Reader复制到sshConn.Writer
	go func() {
		_, err = io.Copy(sshConn, localConn)
		if err != nil {
			fmt.Printf("io.Copy failed: %v", err)
		}
	}()

	// 将sshConn.Reader复制到localConn.Writer
	go func() {
		_, err = io.Copy(localConn, sshConn)
		if err != nil {
			fmt.Printf("io.Copy failed: %v", err)
		}
	}()
}
