package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"i-go/core/conf"
	"i-go/utils"
	"time"
)

/*
	当前etcd v3包有下面俩个manager
	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3"
	推荐使用	"github.com/coreos/etcd/clientv3" 这个
*/
var (
	Cli *clientv3.Client
)

type etcdConf struct {
	Endpoints   []string      `json:"endpoints"`
	DialTimeout time.Duration `json:"dialTimeout"`
	// auth相关 没有也可以不填
	Username string `json:"username"`
	Password string `json:"password"`
	// TLS相关配置可以不填
	CertFile      string `json:"certFile"`      //  "/tmp/test-certs/test-name-1.pem"
	KeyFile       string `json:"keyFile"`       //   "/tmp/test-certs/test-name-1-key.pem"
	TrustedCAFile string `json:"trustedCAFile"` // "/tmp/test-certs/trusted-ca.pem"
}

func init() {
	defer utils.InitLog("Etcd")()
	conf.Init("conf/config.json")

	c := readConf()
	Cli = newConn(c)
}

func readConf() *etcdConf {
	var c etcdConf
	if err := viper.UnmarshalKey("etcd", &c); err != nil {
		panic(err)
	}
	if len(c.Endpoints) == 0 {
		panic("etcd conf nil")
	}
	return &c
}

func newConn(c *etcdConf) *clientv3.Client {
	var etctConfig clientv3.Config
	// 有TLS配置就设置
	if c.CertFile != "" {
		tlsInfo := transport.TLSInfo{
			CertFile:      c.CertFile,
			KeyFile:       c.KeyFile,
			TrustedCAFile: c.TrustedCAFile,
		}
		tlsConfig, err := tlsInfo.ClientConfig()
		if err != nil {
			logrus.WithFields(logrus.Fields{"Scenes": "etcd tls conf error"}).Error(err)
		} else {
			etctConfig.TLS = tlsConfig
		}
	}
	if c.Username != "" {
		etctConfig.Username = c.Username
		etctConfig.Password = c.Password
	}
	etctConfig.Endpoints = c.Endpoints
	etctConfig.DialTimeout = c.DialTimeout

	client, err := clientv3.New(etctConfig)
	if err != nil {
		panic(err)
		return nil
	}
	return client
}

func Release() {
	if Cli != nil {
		_ = Cli.Close()
	}
}
