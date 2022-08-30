package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	localAddrMQ    = flag.String("lmq", ":9889", "host:port to listen on")
	localAddrHTTP  = flag.String("lhttp", ":8081", "host:port to listen on")
	remoteAddrMQ   = flag.String("rmq", "172.20.149.197:4222", "host:port to forward to")
	remoteAddrHTTP = flag.String("rhttp", "172.20.149.197:8888", "host:port to forward to")
	prefix         = flag.String("p", "tcpforward: ", "String to prefix log output")
)

func forward(conn net.Conn, remoteAddr string) {
	client, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Printf("Dial failed: %v", err)
		defer conn.Close()
		return
	}
	log.Printf("Forwarding from %v to %v\n", conn.LocalAddr(), client.RemoteAddr())
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(client, conn)
	}()
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(conn, client)
	}()
}

// 监听两个端口 分别转发 MQ和 http 的流量
func main() {
	flag.Parse()
	log.SetPrefix(*prefix + ": ")

	go func() {
		listener, err := net.Listen("tcp", *localAddrMQ)
		if err != nil {
			log.Fatalf("Failed to setup listener: %v", err)
		}

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatalf("ERROR: failed to accept listener: %v", err)
			}
			log.Printf("Accepted connection from %v\n", conn.RemoteAddr().String())
			go forward(conn, *remoteAddrMQ)
		}
	}()

	go func() {
		listener, err := net.Listen("tcp", *localAddrHTTP)
		if err != nil {
			log.Fatalf("Failed to setup listener: %v", err)
		}

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatalf("ERROR: failed to accept listener: %v", err)
			}
			log.Printf("Accepted connection from %v\n", conn.RemoteAddr().String())
			go forward(conn, *remoteAddrHTTP)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}
