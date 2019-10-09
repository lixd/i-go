package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

const (
	localhost    = "127.0.0.1:2379"
	remotehost   = "192.168.1.9:2379"
	remotehost2  = "192.168.0.2:2379"
	clusterhost1 = "192.168.1.9:32769"
	clusterhost2 = "192.168.1.9:32771"
	clusterhost3 = "192.168.1.9:32773"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		putResp *clientv3.PutResponse
		// getResp               *clientv3.GetResponse
		// delResp               *clientv3.DeleteResponse
		// leaseResp, leaseResp1 *clientv3.LeaseGrantResponse
	)

	// 配置客户端
	config = clientv3.Config{
		Endpoints:   []string{clusterhost1, clusterhost2, clusterhost3},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err.Error())
		return
	}

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
