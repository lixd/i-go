package main

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type Service struct {
	closeChan chan struct{}    //关闭通道
	client    *clientv3.Client //etcd v3 client
	leaseID   clientv3.LeaseID //etcd 租约id
	key       string           //键
	val       string           //值
	wg        sync.WaitGroup
}

// NewService 构造一个注册服务
func NewService(etcdEndpoints []string, key string, val string) (*Service, error) {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdEndpoints,
		DialTimeout: 2 * time.Second,
	})

	if nil != err {
		return nil, err
	}

	s := &Service{
		client:    cli,
		closeChan: make(chan struct{}),
		key:       key,
		val:       val,
	}

	return s, nil
}

// Start 开启注册
// @param - ttlSecond 租期(秒)
func (s *Service) Start(ttlSecond int64) error {

	// minimum lease TTL is 5-second
	resp, err := s.client.Grant(context.TODO(), ttlSecond)
	if err != nil {
		panic(err)
	}

	s.leaseID = resp.ID
	_, err = s.client.Put(context.TODO(), s.key, s.val, clientv3.WithLease(s.leaseID))
	if err != nil {
		panic(err)
	}

	ch, err1 := s.client.KeepAlive(context.TODO(), s.leaseID)
	if nil != err1 {
		panic(err)
	}

	logrus.Info("[discovery] Service Start leaseID:[%d] key:[%s], value:[%s]", s.leaseID, s.key, s.val)

	s.wg.Add(1)
	defer s.wg.Done()

	for {
		select {
		case <-s.closeChan:
			return s.revoke()
		case <-s.client.Ctx().Done():
			return errors.New("server closed")
		case ka, ok := <-ch:
			if !ok {
				logrus.Warn("[discovery] Service Start keep alive channel closed")
				return s.revoke()
			} else {
				logrus.Info("[discovery] Service Start recv reply from Service: %s, ttl:%d", s.key, ka.TTL)
			}
		}
	}

	return nil
}

// Stop 停止
func (s *Service) Stop() {
	close(s.closeChan)
	s.wg.Wait()
	s.client.Close()
}

func (s *Service) revoke() error {

	_, err := s.client.Revoke(context.TODO(), s.leaseID)

	if err != nil {
		logrus.Error("[discovery] Service revoke key:[%s] error:[%s]", s.key, err.Error())
	} else {
		logrus.Info("[discovery] Service revoke successfully key:[%s]", s.key)
	}

	return err
}
