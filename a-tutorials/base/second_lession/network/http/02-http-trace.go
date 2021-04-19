package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"os"
)

// httptrace
func main() {
	req, err := http.NewRequest("GET", "https://www.lixueduan.com", nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	trace := httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial done")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(context.Background(), &trace))
	http.DefaultTransport.RoundTrip(req)
	response, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	io.Copy(os.Stdout, response.Body)
}
