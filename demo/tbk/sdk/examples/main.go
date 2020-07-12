package main

import (
	"fmt"
	sdk "i-go/demo/tbk/sdk"
	"time"
)

const (
	AppKey    = ""
	AppSecret = ""
	Router    = sdk.ReleaseHttps
	Timeout   = time.Second * 3

	ItemInfo = "taobao.tbk.item.info.get" // 淘宝客 api 名称
)

func main() {
	tbk := sdk.NewTBK(AppKey, AppSecret, Router, "", Timeout)
	p := sdk.Parameter{
		"num_iids": "614391759359",
	}
	ret, err := tbk.Execute(ItemInfo, p)
	if err != nil {
		fmt.Printf("err:%v \n", err)
	}
	fmt.Printf("result:%v \n", string(ret))
}
