package aliyunoss

import (
	"errors"
	"i-go/utils"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

type ossConf struct {
	EndPoint            string `json:"EndPoint"`
	EndPointInternal    string `json:"EndPointInternal"`
	AccessKeyID         string `json:"AccessKeyID"`
	AccessKeySecret     string `json:"AccessKeySecret"`
	BucketITest         string `json:"BucketITest"`
	BucketITestInternal string `json:"BucketITestInternal"`
}

const (
	Release = "release"
)

var (
	OssClient *oss.Client
)

// Init 建议显示地调用初始化方法
func Init() {
	defer utils.InitLog("AliyunOSS")()

	ossConf, err := parseConf()
	if err != nil {
		panic(err)
	}
	OssClient, err = newOssClient(ossConf)
	if err != nil {
		panic(err)
	}
}

func parseConf() (*ossConf, error) {
	var c ossConf
	if err := viper.UnmarshalKey("aliyunoss", &c); err != nil {
		return &ossConf{}, err
	}
	if c.AccessKeyID == "" {
		return &ossConf{}, errors.New("aliyunoss conf nil")
	}
	return &c, nil
}

// getRunMode 获取运行模式 release/debug
func getRunMode() string {
	return viper.GetString("run-mode")
}

func newOssClient(conf *ossConf) (*oss.Client, error) {
	var (
		endpoint string
		mode     string
	)

	// 运行模式
	mode = getRunMode()

	// 生产环境使用 oss 内网
	if mode == Release {
		endpoint = conf.EndPointInternal
	} else {
		endpoint = conf.EndPoint
	}
	client, err := oss.New(endpoint, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		return &oss.Client{}, err
	}
	return client, nil
}
