package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"i-go/demo/conf"
	"os"
)

type ossConf struct {
	EndPoint         string `json:"EndPoint"`
	Bucket           string `json:"Bucket"`
	BucketInternal   string `json:"BucketInternal"`
	EndPointInternal string `json:"EndPointInternal"`
	AccessKeyID      string `json:"AccessKeyID"`
	AccessKeySecret  string `json:"AccessKeySecret"`
}

var OssConf = &ossConf{}

func init() {
	path := "conf/config.json"
	if err := conf.Init(path); err != nil {
		logrus.Panic(err)
	}
	OssConf = NewOssConf()
}
func NewOssConf() *ossConf {
	var c ossConf
	if err := viper.UnmarshalKey("oss", &c); err != nil {
		logrus.Panic(err)
	}
	return &ossConf{
		EndPoint:         c.EndPoint,
		Bucket:           c.Bucket,
		EndPointInternal: c.EndPointInternal,
		BucketInternal:   c.BucketInternal,
		AccessKeyID:      c.AccessKeyID,
		AccessKeySecret:  c.AccessKeySecret}
}

func main() {
	fmt.Printf("endpoint =%v \n AccessKeyID=%v \n AccessKeySecret=%v \n Bucket=%v \n", OssConf.EndPoint, OssConf.AccessKeyID, OssConf.AccessKeySecret, OssConf.Bucket)
	// 创建OSSClient实例。
	client, err := oss.New(OssConf.EndPoint, OssConf.AccessKeyID, OssConf.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(OssConf.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 读取本地文件。
	fd, err := os.Open("d:/a.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer fd.Close()

	// 上传文件流。
	err = bucket.PutObject("123.jpg", fd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
