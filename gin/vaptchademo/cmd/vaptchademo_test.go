package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"testing"
	"time"
)

func BenchmarkSpecial(b *testing.B) {
	for i := 0; i < 100; i++ {
		go testOffline()
	}
	select {}
}

const (
	vaptchaURL = "http://www.lixueduan.com:8080/vaptcha/offline?offline_action=get&vid=5e4b559fad980a8810052f9a&callback=VaptchaJsonp1589679108778"
	lixdURL    = "https://www.lixueduan.com"
)

func testOffline() {
	start := time.Now()
	resp, err := http.Get(lixdURL)
	if err != nil {
		logrus.Error(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	// fmt.Printf("%v \n", result.String())
	fmt.Printf("time:%v \n", time.Now().Sub(start))
}
