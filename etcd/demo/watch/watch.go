package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	etcd2 "i-go/core/etcd"
)

func main() {
	var (
		client  *clientv3.Client
		err     error
		putResp *clientv3.PutResponse
	)
	// 0.建立连接
	client = etcd2.New("etcd-local")

	if putResp, err = client.Put(context.TODO(), "maxProcess", "3"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(putResp.Header)
	}
	watcher := clientv3.NewWatcher(client)
	watch := watcher.Watch(context.Background(), "maxProcess")
	go func() {
		response, err := client.Put(context.TODO(), "maxProcess", "123")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(response)
		}
	}()
	for wresp := range watch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

	select {}
}
