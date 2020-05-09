package aliyunoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"i-go/utils"
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

func init() {
	defer utils.InitLog("aliyun-oss")()
	// 获取配置文件
	ossConf := initOssConf()
	// 初始化client
	OssClient = newOssClient(ossConf)
}

func initOssConf() *ossConf {
	var c ossConf
	if err := viper.UnmarshalKey("aliyunoss", &c); err != nil {
		panic(err)
	}
	return &ossConf{
		EndPoint:            c.EndPoint,
		EndPointInternal:    c.EndPointInternal,
		AccessKeyID:         c.AccessKeyID,
		AccessKeySecret:     c.AccessKeySecret,
		BucketITest:         c.BucketITest,
		BucketITestInternal: c.BucketITestInternal,
	}
}

// getRunMode 获取运行模式 release/debug
func getRunMode() string {
	return viper.Get("run-mode").(string)
}

func newOssClient(conf *ossConf) *oss.Client {
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
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "创建OSSClient实例"}).Error(err)
		return nil
	}
	return client
}
