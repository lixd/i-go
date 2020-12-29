package main

import (
	"context"
	"flag"
	"strconv"
	"sync"

	"github.com/coreos/etcd/clientv3"
	log "github.com/sirupsen/logrus"
	"i-go/core/conf"
	"i-go/core/etcd"
	"i-go/utils"
)

/*
PUT 10000 4.7s
GET 10000 2.9s
DELETE 10000 4.9s
*/
var (
	client *clientv3.Client
	kv     clientv3.KV
	number int
	put    bool
	get    bool
	delete bool
)

func init() {
	conf.Init("D:/lillusory/projects/i-go/conf/config.yml")
	// conf.Init("conf/config.yml")
	etcd.Init()

	client = etcd.CliV3
	kv = clientv3.NewKV(client)

	flag.IntVar(&number, "number", 100, "run op count")
	flag.BoolVar(&put, "put", false, "is run put")
	flag.BoolVar(&get, "get", false, "is run get")
	flag.BoolVar(&delete, "delete", false, "is run delete")
	flag.Parse()
}

func main() {
	defer utils.Trace("main")()
	defer etcd.Release()

	// kv.Put(context.Background(),"/vaptcha/api_system_cfg/ad_random",`{"global":0.6,"enable":1,"mad_ratio":1}`)
	// kv.Put(context.Background(),"/vaptcha/api_system_cfg/cdn_servers",`statics.vaptcha.net`)
	// kv.Put(context.Background(),"/vaptcha/api_system_cfg/guide_version",`3.0.0`)
	// kv.Put(context.Background(),"/vaptcha/api_system_cfg/mobile_referer",`v.vaptcha.com`)
	// kv.Put(context.Background(),"/vaptcha/api_system_cfg/vaptcha_key",`4d7449ac01b44420a38233c52ceb0bf3`)
	kv.Put(context.Background(), "/vaptcha/config/js_config/v3-default", `[{"key":"cssversion","value":"2.3.2"},{"key":"embed","value":"vaptcha-sdk-embed.2.9.2.js"},{"key":"invisible","value":"vaptcha-sdk-invisible.2.9.2.js"},{"key":"mobile","value":"vaptcha-sdk-mobile.2.9.8.js"},{"key":"popup","value":"vaptcha-sdk-popup.2.9.2.js"}]`)

	// ch := generate()

	// var wg sync.WaitGroup
	// for i := 0; i < 6; i++ {
	// 	wg.Add(1)
	// 	go doOP(ch, &wg)
	// }
	// wg.Wait()
}

type T struct {
	Key   string
	Value string
}

func generate() chan T {
	var ch = make(chan T)
	go func() {
		for i := 0; i < number; i++ {
			item := T{
				Key:   "hello" + strconv.Itoa(i),
				Value: strconv.Itoa(i),
			}
			ch <- item
		}
		close(ch)
	}()
	return ch
}

func doOP(ch chan T, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		if put {
			doPUT(v)
		}
		if get {
			doGET(v)
		}
		if delete {
			doDELETE(v)
		}
	}
}

func doPUT(data T) {
	_, err := kv.Put(context.Background(), data.Key, data.Value)
	if err != nil {
		log.Printf("Put error:%v", err)
	}
}
func doGET(data T) {
	_, err := kv.Get(context.Background(), data.Key)
	if err != nil {
		log.Printf("Get error:%v", err)
	}
}
func doDELETE(data T) {
	_, err := kv.Delete(context.Background(), data.Key)
	if err != nil {
		log.Printf("Delete error:%v", err)
	}
}
