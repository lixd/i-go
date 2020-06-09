package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"i-go/demo/tbk/sdk"
	"i-go/demo/tbk/sdk/examples/model"
	"i-go/utils"
	"strconv"
	"time"
)

const (
	adzone_id = "110386550278"
	AppKey    = "28657749"
	AppSecret = "790273d741f3e988a7e49b046d335968"
	Router    = "http://gw.api.taobao.com/router/rest"
	Timeout   = time.Second * 3
)

func main() {
	id := "614391759359"
	title := FindTitleById(id)
	coupon := FindCouponByTitle(title, id)
	fmt.Println("推广链接:", coupon)
}

// getShareURL 根据id 从多个返回结果中找出需要的那个
func getShareURL(data *model.Material, id string) string {
	numId, err := strconv.Atoi(id)
	if err != nil {
		return ""
	}
	var shareURL string
	for _, v := range data.TbkDgMaterialOptionalResponse.ResultList.MapData {
		if v.CouponShareURL != "" && v.ItemID == int64(numId) {
			shareURL = v.CouponShareURL
		}
	}
	return shareURL
}

func FindTitleById(id string) string {
	tbk := sdk.NewTBK(AppKey, AppSecret, Router, "", Timeout)
	p := sdk.Parameter{
		"num_iids": id,
	}
	bodyBytes, err := tbk.Execute("taobao.tbk.item.info.get", p)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "request"})
	}
	var data model.Item
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "res Unmarshal"})
	}
	fmt.Println("result: ", string(bodyBytes))
	fmt.Println("data: ", data)
	if len(data.TbkItemInfoGetResponse.Results.NTbkItem) != 0 {
		fmt.Println("Title: ", data.TbkItemInfoGetResponse.Results.NTbkItem[0].Title)
		return data.TbkItemInfoGetResponse.Results.NTbkItem[0].Title
	}
	return ""
}

func FindCouponByTitle(title, id string) string {
	tbk := sdk.NewTBK(AppKey, AppSecret, Router, "", Timeout)
	p := sdk.Parameter{
		"q":         title,
		"adzone_id": adzone_id,
	}
	bodyBytes, err := tbk.Execute("taobao.tbk.dg.material.optional", p)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "request"})
	}
	var data model.Material
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "res Unmarshal"})
	}
	if len(data.TbkDgMaterialOptionalResponse.ResultList.MapData) != 0 {
		shareURL := getShareURL(&data, id)
		return shareURL
	}
	return ""
}
