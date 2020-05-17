package lock

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"sync"
)

type EtcdMutex struct {
	ttl      int64  // 租约时间
	prefix   string // lock的key
	ctx      context.Context
	clientV3 *clientv3.Client
	lease    clientv3.Lease
	leaseID  clientv3.LeaseID
}

// NewEtcdMutex
func NewEtcdMutex(ttl int64, prefix string, ctx context.Context, client *clientv3.Client) *EtcdMutex {
	var em = EtcdMutex{
		prefix:   prefix,
		ttl:      ttl,
		ctx:      ctx,
		clientV3: client,
	}
	return &em
}

// InitMutex 初始化 txn lease和leaseID
func (em *EtcdMutex) InitMutex() error {
	lease := clientv3.NewLease(em.clientV3)
	em.lease = lease
	grantResponse, err := lease.Grant(em.ctx, em.ttl)
	if err != nil {
		return err
	}
	em.leaseID = grantResponse.ID
	return nil
}

// Lock 抢锁
func (em *EtcdMutex) Lock() error {
	// LOCK:
	txn := em.clientV3.Txn(em.ctx)
	txnResp, err := txn.If(clientv3.Compare(clientv3.CreateRevision(em.prefix), "=", 0)).
		Then(clientv3.OpPut(em.prefix, "", clientv3.WithLease(em.leaseID))).
		Else().Commit()
	if err != nil {
		return err
	}
	if !txnResp.Succeeded { // 判断txn.if条件是否成立
		return errors.New("抢锁失败")
	}
	return nil
}

// Lock 抢锁
func (em *EtcdMutex) LockSync() error {
	// LOCK:
	txn := em.clientV3.Txn(em.ctx)
	txnResp, err := txn.If(clientv3.Compare(clientv3.CreateRevision(em.prefix), "=", 0)).
		Then(clientv3.OpPut(em.prefix, "", clientv3.WithLease(em.leaseID))).
		Else().Commit()
	// 事务失败直接返回
	if err != nil {
		return err
	}
	// 条件成立说明获取锁成功
	if txnResp.Succeeded { // 判断txn.if条件是否成立
		return nil
	} else {
		// 否则获取失败 for循环watch 一直阻塞到key被delete才返回
		for {
			wch := em.clientV3.Watch(em.ctx, em.prefix)
			for wr := range wch {
				for _, ev := range wr.Events {
					if ev.Type == clientv3.EventTypeDelete {
						// 这里不能直接返回 应该在走一遍事务 否则阻塞的client都会同时获取到锁
						// return nil
						return em.LockSync()
					}
				}
			}
		}
	}
}
func (em *EtcdMutex) UnLock() error {
	_, err := em.lease.Revoke(context.Background(), em.leaseID)
	if err != nil {
		return err
	}
	return nil
}

func testLockSync(ttl int64, prefix string, num int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	eMutex := NewEtcdMutex(ttl, prefix, context.Background(), client)
	err := eMutex.InitMutex()
	if err != nil {
		fmt.Println("init:", err)
		return
	}
	err = eMutex.LockSync()
	if err != nil {
		fmt.Printf("协程 %v 抢锁错误:%v \n", num, err)
		return
	}
	fmt.Printf("协程 %v 抢锁成功 \n", num)
	doSomething()
	_ = eMutex.UnLock()
	fmt.Printf("协程 %v 释放锁 \n", num)
}
