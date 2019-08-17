package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type GroupManager struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
	once   sync.Once
}

func NewGroupManager() *GroupManager {
	ret := new(GroupManager)
	ret.ctx, ret.cancel = context.WithCancel(context.Background())
	return ret
}

func (this *GroupManager) Close() {
	this.once.Do(this.cancel)
}

func (this *GroupManager) Wait() {
	this.wg.Wait()
}

func (this *GroupManager) Add(delta int) {
	this.wg.Add(delta)
}

func (this *GroupManager) Done() {
	this.wg.Done()
}

func (this *GroupManager) Chan() <-chan struct{} {
	return this.ctx.Done()
}

type Target interface {
	Set(string, string)
	Create(string, string)
	Modify(string, string)
	Delete(string)
}

type Config struct {
	Servers        []string
	DailTimeout    int64
	RequestTimeout int64
	Prefix         bool
	Target         string
}

func Watch(gm *GroupManager, cfg *Config, target Target) {
	defer gm.Done()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Servers,
		DialTimeout: time.Duration(cfg.DailTimeout) * time.Second,
	})
	if err != nil {
		panic(err.Error())
		return
	}
	defer cli.Close() // make sure to close the client

	logrus.Info("[discovery] start watch %s", cfg.Target)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.RequestTimeout)*time.Second)
	var resp *clientv3.GetResponse
	if cfg.Prefix {
		resp, err = cli.Get(ctx, cfg.Target, clientv3.WithPrefix())
	} else {
		resp, err = cli.Get(ctx, cfg.Target)
	}
	cancel()
	if err != nil {
		panic(err.Error())
	}
	for _, ev := range resp.Kvs {
		target.Set(string(string(ev.Key)), string(ev.Value))
	}

	var rch clientv3.WatchChan
	if cfg.Prefix {
		rch = cli.Watch(context.Background(), cfg.Target, clientv3.WithPrefix(), clientv3.WithRev(resp.Header.Revision+1))
	} else {
		rch = cli.Watch(context.Background(), cfg.Target, clientv3.WithRev(resp.Header.Revision+1))
	}
	for {
		select {
		case <-gm.Chan():
			logrus.Info("[discovery] watch %s close", cfg.Target)
			return
		case wresp := <-rch:
			err := wresp.Err()
			if err != nil {
				logrus.Info("[discovery] watch %s response error: %s ", cfg.Target, err.Error())
				gm.Close()
				return
			}
			logrus.Debug("[discovery] watch %s response %+v", cfg.Target, wresp)
			for _, ev := range wresp.Events {
				if ev.IsCreate() {
					target.Create(string(ev.Kv.Key), string(ev.Kv.Value))
				} else if ev.IsModify() {
					target.Modify(string(ev.Kv.Key), string(ev.Kv.Value))
				} else if ev.Type == mvccpb.DELETE {
					target.Delete(string(ev.Kv.Key))
				} else {
					logrus.Error("[discovery] no found watch type: %s %q", ev.Type, ev.Kv.Key)
				}
			}
		}
	}
}
