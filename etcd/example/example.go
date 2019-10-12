package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"time"
)

const (
	remote1 = "192.168.0.2:2379"
	remote2 = "192.168.0.2:3379"
	remote3 = "192.168.0.2:4379"
)

func main() {
	// 1.配置
	config := clientv3.Config{
		Endpoints:   []string{remote1, remote2, remote3},
		DialTimeout: 5 * time.Second}
	// 2.建立连接
	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}
	// 3.获取kv对象
	// kv接口包含了etcd所有基本操作
	// 	与client.KV相比clientv3.NewKV(client) 多个自动重连功能 所以一般用这个
	kv := clientv3.NewKV(client)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 4.基本操作2
	// 4.1 Put
	response, err := kv.Put(ctx, "/test/t1", "first1")
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Put"}).Error(err)
	} else {
		fmt.Println("response ", response)
	}
	response, err = kv.Put(ctx, "/test/t1", "second1")
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Put"}).Error(err)
	} else {
		fmt.Println("response ", response)
	}
	// 4.2 Get
	getResponse, err := kv.Get(ctx, "/test/t1")
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Get"}).Error(err)
	} else {
		fmt.Println("response ", getResponse)
	}
	getResponse, err = kv.Get(ctx, "/test", clientv3.WithPrefix())
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Get"}).Error(err)
	} else {
		fmt.Println("response ", getResponse.Kvs)
	}

	// 5. Lease
	lease := clientv3.NewLease(client)
	TenLease, err := lease.Grant(ctx, 10)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Grant"}).Error(err)
	} else {
		fmt.Println("response ", TenLease)
	}
	putResponse, err := kv.Put(ctx, "/test/t3", "third", clientv3.WithLease(TenLease.ID))
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Put WithLease"}).Error(err)
	} else {
		fmt.Println("response ", putResponse)
	}

	// 	6.OP 获取一个op对象 在调用KV.Do()方法执行
	op := clientv3.OpPut("/test/op", "fourth")
	opResponse, err := kv.Do(ctx, op)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "OpPut"}).Error(err)
	} else {
		fmt.Println("response ", opResponse.Put())
	}
	opg := clientv3.OpGet("/test/op")
	opResponse, err = kv.Do(ctx, opg)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "OpGet"}).Error(err)
	} else {
		fmt.Println("response ", opResponse.Get())
	}
	// 	7.Txn
	txn := kv.Txn(ctx)
	txnResponse, err := txn.If(clientv3.Compare(clientv3.Value("/test/op"), "=", "fourth")).
		Then(clientv3.OpPut("/test/txn", "then")).
		Else(clientv3.OpPut("/test/txn", "else")).
		Commit()
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "txn.If"}).Error(err)
	} else {
		fmt.Println("response ", txnResponse)
	}
	if txnResponse.Succeeded {
		fmt.Println("Txn Success")
	} else {
		fmt.Println("Txn Failure")
	}
}
