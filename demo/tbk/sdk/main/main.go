package main

import (
	"fmt"
	"github.com/lixd/gotbk"
	"time"
)

const (
	AppKey    = ""
	AppSecret = ""
	Router    = gotbk.ReleaseHttps
	Timeout   = time.Second * 3

	ItemInfo = "taobao.tbk.item.info.get" // 淘宝客 api 名称
)

func main() {
	tbk := gotbk.NewTBK(AppKey, AppSecret, Router, "", Timeout)
	p := gotbk.Parameter{
		"num_iids": "614391759359",
	}
	ret, err := tbk.Execute(ItemInfo, p)
	if err != nil {
		fmt.Printf("err:%v \n", err)
	}
	fmt.Printf("result:%v \n", string(ret))
}
