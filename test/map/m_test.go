package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"

	"google.golang.org/grpc"
)

// 42.6 ns/op
func BenchmarkSMap(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.LoadOrStore("key", "value")
	}
}

//  2.42 ns/op
func BenchmarkMap(b *testing.B) {
	m := make(map[string]struct{})
	for i := 0; i < b.N; i++ {
		_ = m["key"]
	}
}

func BenchmarkDeepCopy(b *testing.B) {
	var (
		src       = make(map[string]*int64)
		det       = make(map[string]int64)
		v   int64 = 1
	)
	for i := int64(0); i < 1000; i++ {
		src[strconv.Itoa(int(i))] = &v
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deepCopy(src, det)
	}
}

func BenchmarkDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Demo()
	}
}

// serverConfig 相关文档：https://github.com/grpc/grpc/blob/master/doc/service_config.md
// 具体配置内容定义：https://github.com/grpc/grpc-proto/blob/master/grpc/service_config/service_config.proto
// RetryableStatusCodes 见 google.golang.org/grpc/codes
const serverConfig = `{
		"loadBalancingPolicy":"round_robin",
		"methodConfig": [{
		  "name": [],
		  "retryPolicy": {
			  "MaxAttempts": 3,
			  "InitialBackoff": "0.1s",
			  "MaxBackoff": "1s",
			  "BackoffMultiplier": 2.0,
			  "RetryableStatusCodes": [ "UNAVAILABLE","INTERNAL","ABORTED","DEADLINE_EXCEEDED" ]
		  }
		}]}`

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("xxx"),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), // 指定轮询负载均衡算法
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(serverConfig),
	)
	if err != nil {
		log.Fatal(err)
	}
	_ = conn
}
