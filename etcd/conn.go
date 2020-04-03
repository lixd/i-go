package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
	"time"
)

/*
	当前etcd v3包有下面俩个
	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3"
	推荐使用	"github.com/coreos/etcd/clientv3" 这个
*/
// 将所有连接存在connect中 不存在时才去连建立新连接
type connect map[string]*clientv3.Client

var conn = make(connect)

func New(key string) *clientv3.Client {
	return conn.get(key)
}

func Release() {
	for _, cli := range conn {
		if cli != nil {
			_ = cli.Close()
		}
	}
}

func (c *connect) get(key string) *clientv3.Client {
	if client, ok := conn[key]; ok {
		return client
	}
	client := c.new(key)
	return client
}

func (c *connect) new(key string) *clientv3.Client {
	//endpoints := endpoints(key)
	endpoints := endpointsMock()

	config := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(config)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "conn etcd error"}).Error(err)
		return nil
	}
	(*c)[key] = client

	return client
}

// endpoints 查配置文件 获取连接地址
func endpoints(key string) []string {
	endpoints := viper.GetString(key)
	if endpoints == "" {
		panic("cannot find endpoints")
	}
	return strings.Split(endpoints, ",")
}

func endpointsMock() []string {
	return []string{"http://192.168.1.2:2379", "http://192.168.1.3:2379", "http://192.168.1.4:2379"}
}
