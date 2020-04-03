package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
	"i-go/etcd"
	"time"
)

func main() {
	client := etcd.New("etcd-local")
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
		fmt.Println("Put response ", response.OpResponse().Put())
	}
	response, err = kv.Put(ctx, "/test/t1", "second1")
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Put"}).Error(err)
	} else {
		fmt.Println("Put response ", response.OpResponse().Put())
	}
	// 4.2 Get
	getResponse, err := kv.Get(ctx, "/test/t1")
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Get"}).Error(err)
	} else {
		fmt.Println("Get response ", getResponse.OpResponse().Get())
	}
	getResponse, err = kv.Get(ctx, "/test", clientv3.WithPrefix())
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Get"}).Error(err)
	} else {
		fmt.Println("Get response ", getResponse.Kvs)
	}

	// 5. Lease
	lease := clientv3.NewLease(client)
	TenLease, err := lease.Grant(ctx, 10)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Grant"}).Error(err)
	} else {
		fmt.Println("Grant response ", TenLease)
	}
	putResponse, err := kv.Put(ctx, "/test/t3", "third", clientv3.WithLease(TenLease.ID))
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Put WithLease"}).Error(err)
	} else {
		fmt.Println("Put response ", putResponse.OpResponse().Put())
	}
	// 获取剩余TTL时间
	liveResponse, err := lease.TimeToLive(ctx, TenLease.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "TimeToLive"}).Error(err)
	} else {
		fmt.Println("response ", liveResponse.TTL)
	}
	// 列举所有etcd中的租约。
	leasesResponse, err := lease.Leases(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Leases"}).Error(err)
	} else {
		fmt.Println("Leases response ", leasesResponse.Leases)
	}
	// 为某个租约续约一次
	aliveResponse, err := lease.KeepAliveOnce(ctx, TenLease.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "KeepAliveOnce"}).Error(err)
	} else {
		fmt.Println("KeepAliveOnce response ", aliveResponse.TTL)
	}
	// 自动定时的续约某个租约
	aliveResponseC, err := lease.KeepAlive(ctx, TenLease.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "KeepAlive"}).Error(err)
	} else {
		fmt.Println("KeepAlive response ", <-aliveResponseC)
	}
	// 释放一个租约
	revokeResponse, err := lease.Revoke(ctx, TenLease.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "Revoke"}).Error(err)
	} else {
		fmt.Println("Revoke response ", revokeResponse)
	}

	// 貌似是关闭当前客户端建立的所有租约。
	err = lease.Close()
	// 	6.OP 获取一个op对象 在调用KV.Do()方法执行
	op := clientv3.OpPut("/test/op", "fourth")
	opResponse, err := kv.Do(ctx, op)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "OpPut"}).Error(err)
	} else {
		fmt.Println("Do response ", opResponse.Put())
	}
	opg := clientv3.OpGet("/test/op")
	opResponse, err = kv.Do(ctx, opg)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "OpGet"}).Error(err)
	} else {
		fmt.Println("Do response ", opResponse.Get())
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
		fmt.Println("txn response ", txnResponse)
	}
	if txnResponse.Succeeded {
		fmt.Println("Txn Success")
	} else {
		fmt.Println("Txn Failure")
	}
}
