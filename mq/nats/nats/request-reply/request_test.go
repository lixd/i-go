package request_reply

import (
	"net"
	"strconv"
	"testing"
)

func TestRequest(t *testing.T) {
	stopChan := make(chan struct{})
	Subscribe("hello", demoHandler, stopChan)
	for i := 0; i < 100; i++ {
		Request("hello", []byte(strconv.Itoa(i)))
	}
	stopChan <- struct{}{}
}

func TestA(t *testing.T) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		t.Fatal(err)
	}
	for _, address := range addresses {
		t.Log(address)
	}
}
