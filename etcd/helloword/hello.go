package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"i-go/core/conf"
	"i-go/core/etcd"

	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
)

var (
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
)

const (
	Prefix = "/hello"
	Suffix = "/2"
)

func init() {
	conf.Load("/conf/config.yml")
	etcd.Init()

	client = etcd.CliV3
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)
}

func main() {
	defer etcd.Release()
	testLease()
	// put()
	// get()
	// delete()
	// leaseFunc()
	// txn()
	// watch()
	// go putForWatch()
	// go watch()
	// select {}
	// compact()
}

func testLease() {
	// 	put+lease
	response, err := lease.Grant(context.Background(), 100)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd delete"}).Error(err)
	}
	leaseID := response.ID
	fmt.Printf("leaseID:%v TTL:%v \n", leaseID, response.TTL)
	_, err = kv.Put(context.Background(), Prefix+"/lease1", "lease1", clientv3.WithLease(leaseID))
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd Put"}).Error(err)
	}
	_, err = kv.Put(context.Background(), Prefix+"/lease1", "noLease") // 更新时若不带leaseID会移除旧Lease
	getResponse, err := kv.Get(context.Background(), Prefix+"/lease1")
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd Put"}).Error(err)
	}
	fmt.Printf("resp:%+v\n", getResponse)
}

func compact() {
	var rev int64 = 10
	// rev:会压缩指定版本之前的记录
	// clientv3.WithCompactPhysical(): RPC 将会等待直 压缩物理性地应用到数据库，之后被压缩的项将完全从后端数据库中移除
	kv.Compact(context.Background(), rev, clientv3.WithCompactPhysical())
}

func putForWatch() {
	for i := 0; i < 9; i++ {
		_, err := kv.Put(context.Background(), Prefix+Suffix, strconv.Itoa(i))
		if err != nil {
			logrus.WithFields(logrus.Fields{"Scenes": "etcd put"}).Error(err)
		}
		time.Sleep(time.Millisecond * 100)
	}
	_, err := kv.Delete(context.Background(), Prefix+Suffix)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd put"}).Error(err)
	}
}
func watch() {
	// In order to prevent a watch stream being stuck in a partitioned node,
	// make sure to wrap context with "WithRequireLeader".
	// 为了保证请求不卡在分区节点上 请一定是要WithRequireLeader来包装ctx
	parentCtx, cancel := context.WithCancel(context.Background())
	ctx := clientv3.WithRequireLeader(parentCtx)
	watchChan := client.Watch(ctx, "mykey")
	for wr := range watchChan {
		for _, e := range wr.Events {
			switch e.Type {
			case clientv3.EventTypePut:
				fmt.Printf("watch event put-current: %#v \n", string(e.Kv.Value))
				// 不再watch后记得cancel掉 ctx 否则etcd资源得不到释放
				if string(e.Kv.Value) == "release" {
					cancel()
				}
			case clientv3.EventTypeDelete:
				fmt.Printf("watch event delete-current: %#v \n", string(e.Kv.Value))
			default:
			}
		}
	}
	cancel()
}

func txn() {
	kv.Put(context.Background(), Prefix+Suffix, "f")
	// 开启事务
	txn := kv.Txn(context.Background())
	getOwner := clientv3.OpGet(Prefix+Suffix, clientv3.WithFirstCreate()...)
	// 如果/illusory/cloud的值为hello则获取/illusory/cloud的值 否则获取/illusory/wind的值
	txnResp, err := txn.If(clientv3.Compare(clientv3.Value(Prefix+Suffix), "=", "hello")).
		Then(clientv3.OpGet(Prefix+"/equal"), getOwner).
		Else(clientv3.OpGet(Prefix+"/unequal"), getOwner).
		Commit()
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd put"}).Error(err)
		return
	}
	fmt.Printf("事务%#v \n", txnResp)
	if txnResp.Succeeded { // If = true
		fmt.Println("true", txnResp.Responses[0].GetResponseRange().Kvs)
	} else { // If =false
		fmt.Println("false", txnResp.Responses[0].GetResponseRange().Kvs)
	}
	kv.Delete(context.Background(), Prefix+Suffix)
}

func put() {
	// hello到hello12的key
	kv.Get(context.Background(), "hello", clientv3.WithRange("hello12"))
	// hello前缀的所有key
	kv.Get(context.Background(), "hello", clientv3.WithPrefix())
	response, err := kv.Put(context.Background(), Prefix+Suffix, "hello")
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd put"}).Error(err)
	}
	fmt.Println(response)
	response, err = kv.Put(context.Background(), Prefix+"/equal", "equal")
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd put"}).Error(err)
	}
	fmt.Println(response)
	response, err = kv.Put(context.Background(), Prefix+"/unequal", "unequal")
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd put"}).Error(err)
	}
	fmt.Println(response)
}

func get() {
	response, err := kv.Get(context.Background(), Prefix+Suffix)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd get"}).Error(err)
	}
	fmt.Println(response)
}

func delete() {
	response, err := kv.Delete(context.Background(), Prefix+Suffix)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd delete"}).Error(err)
	}
	fmt.Println(response)
}

// leaseFunc lease租约
func leaseFunc() {
	response, err := lease.Grant(context.Background(), 10)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd delete"}).Error(err)
	}
	leaseID := response.ID
	fmt.Printf("leaseID:%v TTL:%v \n", leaseID, response.TTL)
	_, err = kv.Put(context.Background(), Prefix+"/lease1", "lease1", clientv3.WithLease(leaseID))
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd Put"}).Error(err)
	}
	leasesResponse, err := lease.Leases(context.Background())
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd Leases"}).Error(err)
	}
	fmt.Printf("Leases:%v \n", leasesResponse.Leases)
	// 主动给Lease进行续约
	keepAliveChan, err := client.KeepAlive(context.TODO(), leaseID)
	if err != nil { // 有协程来帮自动续租,TTL剩余一半时就会续约。
		fmt.Println(err)
		return
	}

	go func() {
		for resp := range keepAliveChan {
			fmt.Println("续租:", resp)
		}
	}()

	for {
		time.Sleep(time.Millisecond * 500)
		liveResponse, err := lease.TimeToLive(context.Background(), leaseID)
		if err != nil {
			logrus.WithFields(logrus.Fields{"Scenes": "etcd TimeToLive"}).Error(err)
		}
		fmt.Printf("leaseID:%v TTL:%v GrantedTTL:%v Keys:%v \n", liveResponse.ID, liveResponse.TTL, liveResponse.GrantedTTL, liveResponse.Keys)
		if liveResponse.TTL == -1 {
			break
		}
	}
}
