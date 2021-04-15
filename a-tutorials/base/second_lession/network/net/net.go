package main

import (
	"fmt"
	"net"
	"strings"
)

func NetworkInterface() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for _, v := range interfaces {
		fmt.Println(v)
	}
	fmt.Println(strings.Repeat("~", 20))
	names, err := net.LookupAddr("192.168.0.1")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for _, v := range names {
		fmt.Println(v)
	}
	fmt.Println(strings.Repeat("~", 20))

}
