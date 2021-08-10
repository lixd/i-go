package main

import (
	"log"
	"net"
	"runtime"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{Port: 20000})
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	defer listen.Close()
	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			log.Printf("read udp error: %v\n", err)
			continue
		}
		runtime.GC()
		data := append([]byte("hello "), buf[:n]...)
		_, _ = listen.WriteToUDP(data, addr)
	}
}
