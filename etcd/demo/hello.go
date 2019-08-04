package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	var (
		config                clientv3.Config
		client                *clientv3.Client
		err                   error
		putResp               *clientv3.PutResponse
		getResp               *clientv3.GetResponse
		delResp               *clientv3.DeleteResponse
		leaseResp, leaseResp1 *clientv3.LeaseGrantResponse
	)
	// 配置客户端
	config = clientv3.Config{
		Endpoints:   []string{"192.168.1.9:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err.Error())
		return
	}
	// 用于读写etcd的键值对
	con := context.TODO()

	if putResp, err = client.Put(con, "test", "haha"); err != nil {
		return
	} else {
		fmt.Println(putResp.Header)
	}

	kv := clientv3.NewKV(client)
	// 用kv设置key
	if putResp, err = kv.Put(con, "/cron/jobs/job2", "hello", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	} else {
		// fmt.Println(putResp.Header)
		if putResp.PrevKv != nil {
			fmt.Println(string(putResp.PrevKv.Value))
		}
	}
	// 用kv获取Key
	if getResp, err = kv.Get(con, "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(getResp.Kvs)
	}
	// 用kv删除key
	if delResp, err = kv.Delete(con, "/cron/jobs/job2", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(len(delResp.PrevKvs))
	}
	// 建立租约
	if leaseResp, err = client.Grant(con, 10); err != nil {
		fmt.Println(err)
		return
	}
	// 用client也可以设置key，kv是client的一个结构，因此可以使用其方法
	if putResp, err = client.Put(con, "/cron/lock/job1", "ok", clientv3.WithLease(leaseResp.ID)); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(putResp.Header.Revision)
	}
	if keepAliveChan, err := client.KeepAlive(con, leaseResp.ID); err != nil { // 有协程来帮自动续租，每秒一次。
		fmt.Println(err)
		return
	} else {
		go func() {
			for {
				select {
				case resp := <-keepAliveChan:
					if resp == nil {
						fmt.Println("续租失败")
						goto END
					} else {
						fmt.Println("续租成功")
					}
				}
			}
		END:
		}()
	}
	k := 8
	for k != 0 {
		if getResp, err = client.Get(con, "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(getResp.Count)
			time.Sleep(2 * time.Second)
		}
		k--
	}

	go func() {
		for {
			client.Put(con, "/cron/watch/job1", "I am watch job1")
			time.Sleep(1 * time.Second)
			client.Delete(con, "/cron/watch/job1")
			time.Sleep(1 * time.Second)
		}

	}()
	conWithTimeout, cancelFunc := context.WithCancel(context.TODO())
	wch := client.Watch(conWithTimeout, "/cron/watch/job1", clientv3.WithRev(getResp.Header.Revision))
	tt := time.After(10 * time.Second)
	go func() {
		select {
		case <-tt:
			cancelFunc()
		}
	}()

	for resp := range wch {
		for _, res := range resp.Events {
			fmt.Println(res.Type, string(res.Kv.Key), string(res.Kv.Value))
		}
	}

	conWithTimeout1, cancelFunc1 := context.WithCancel(context.TODO())
	// 做分布式锁相关,执行事务
	// 建立租约、用租约抢key，抢到后执行业务逻辑，抢失败返回。函数退出时要defer把租约关闭
	client.Grant(conWithTimeout1, 10)
	if leaseResp1, err = client.Grant(con, 10); err != nil {
		fmt.Println(err)
		return
	}
	// defer逻辑可以保证租约被清理，防止长期占用key
	defer client.Revoke(con, leaseResp1.ID)
	defer cancelFunc1()
	if keepAliveChan1, err := client.KeepAlive(conWithTimeout1, leaseResp1.ID); err != nil {
		fmt.Println(err)
		return
	} else {
		go func() {
			for {
				<-keepAliveChan1
			}
		}()
	}
	// 打开下面可以看锁已经被抢占的情况
	// client.Put(con,"/cron/txn/job1","I GET FIRST",clientv3.WithLease(leaseResp1.ID))

	// 执行事务，必须带上租约ID，这样在取消租约的时候key会同时失效。
	txn := client.Txn(con)
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/txn/job1"), "=", 0)).
		Then(clientv3.OpPut("/cron/txn/job1", "my job", clientv3.WithLease(leaseResp1.ID))).
		Else(clientv3.OpGet("/cron/txn/job1"))
	if txnResp, err := txn.Commit(); err != nil {
		fmt.Println(err)
		return
	} else {
		if !txnResp.Succeeded {
			fmt.Println("锁被占用：", txnResp.Responses)
			return
		}
	}

	// 拿到锁执行任务
	fmt.Println("执行业务逻辑")
	time.Sleep(2 * time.Second)
	return
}
