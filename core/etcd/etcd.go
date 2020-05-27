package etcd

import (
	"errors"
	"i-go/utils"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/spf13/viper"
)

/*
	当前etcd v3包有下面俩个manager
	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3"
	推荐使用	"github.com/coreos/etcd/clientv3" 这个
*/
var (
	CliV3 *clientv3.Client
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

func Init() {
	defer utils.InitLog("Etcd")()

	c, err := parseConf()
	if err != nil {
		panic(err)
	}
	CliV3, err = newClient(c)
	if err != nil {
		panic(err)
	}
}

func parseConf() (*etcdConf, error) {
	var c etcdConf
	if err := viper.UnmarshalKey("etcd", &c); err != nil {
		return &etcdConf{}, err
	}
	// 默认单位为纳秒
	if len(c.Endpoints) == 0 {
		return &etcdConf{}, errors.New("etcd conf nil")
	}
	return &c, nil
}

func newClient(c *etcdConf) (*clientv3.Client, error) {
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
			return &clientv3.Client{}, err
		}
		etctConfig.TLS = tlsConfig
	}
	if c.Username != "" {
		etctConfig.Username = c.Username
		etctConfig.Password = c.Password
	}
	etctConfig.Endpoints = c.Endpoints
	etctConfig.DialTimeout = c.DialTimeout

	client, err := clientv3.New(etctConfig)
	if err != nil {
		return &clientv3.Client{}, err
	}
	return client, nil
}

func Release() {
	if CliV3 != nil {
		_ = CliV3.Close()
	}
}
