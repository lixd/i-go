package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"i-go/demo/common/ret"
	"i-go/demo/tbk/dto"
	"i-go/demo/tbk/model"
	"i-go/demo/validate"
	"i-go/utils"
	"strconv"
	"time"
)

const (
	AdZoneId = "110386550278"
	ItemInfo = "taobao.tbk.item.info.get"
	Material = "taobao.tbk.dg.material.optional"
)

type ITBK interface {
	FindURLByKeyWords(req *dto.TBKReq) *ret.Result
}

type tbk struct {
	TBK *sdk.TBK
}

func NewTBK(appKey, appSecret, router, session string, timeout time.Duration) ITBK {
	return &tbk{
		TBK: sdk.NewTBK(appKey, appSecret, router, session, timeout),
	}
}

// FindURLByKeyWords 根据关键字、链接查询商品
func (tbk *tbk) FindURLByKeyWords(req *dto.TBKReq) *ret.Result {
	var keyWords = req.KeyWords
	var ids = make([]string, 0)
	isNumbers := validate.IsNumbers(keyWords)
	if isNumbers {
		//itemId, err := validate.ParseItemId(keyWords)
		//if err != nil {
		//	return ret.Fail("", "输入错误")
		//}
		ids = append(ids, keyWords)
		title, err := tbk.FindTitleById(keyWords)
		if err != nil {
			return ret.Fail("", "查询失败,请稍后重试")
		}
		keyWords = title
	}
	coupon, err := tbk.FindCouponByTitle(keyWords, ids...)
	if err != nil {
		return ret.Fail("", "查询失败,请稍后重试")
	}
	var (
		resp dto.TBKResp
		item dto.TBKItem
	)
	list := make([]dto.TBKItem, 0)

	for _, v := range coupon {
		item = dto.TBKItem{ShareURL: v}
		list = append(list, item)
	}
	resp.List = list
	return ret.Success(resp)
}

// FindTitleById 根据 itemId 查询完整标题
func (tbk *tbk) FindTitleById(id string) (string, error) {
	p := sdk.Parameter{
		"num_iids": id,
	}
	bodyBytes, err := tbk.TBK.Execute(ItemInfo, p)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "request"})
		return "", err
	}
	var data model.Item
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "res Unmarshal"})
		return "", err
	}
	fmt.Println("result: ", string(bodyBytes))
	fmt.Println("data: ", data)
	if len(data.TbkItemInfoGetResponse.Results.NTbkItem) != 0 {
		fmt.Println("Title: ", data.TbkItemInfoGetResponse.Results.NTbkItem[0].Title)
		return data.TbkItemInfoGetResponse.Results.NTbkItem[0].Title, nil
	}
	return "", errors.New("查询失败")
}

// FindCouponByTitle 根据 商品标题 查询推广链接
func (tbk *tbk) FindCouponByTitle(title string, id ...string) ([]string, error) {
	URLs := make([]string, 0)
	p := sdk.Parameter{
		"q":         title,
		"adzone_id": AdZoneId,
	}
	bodyBytes, err := tbk.TBK.Execute(Material, p)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "request"})
		return URLs, err
	}
	var data model.Material
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "res Unmarshal"})
		return URLs, err
	}
	URLs = parseItem(&data, id...)
	/*	if len(id) != 0 && id[0] != "" {
			URLs = getShareURL(&data, id[0])
		} else {
			URLs = getAllShareURL(&data)
		}*/
	return URLs, nil
}

// parseItem
func parseItem(data *model.Material, id ...string) []string {
	if len(id) != 0 && id[0] != "" {
		return getShareURL(data, id[0])
	}
	return getAllShareURL(data)
}

// getShareURL 根据id 从多个返回结果中找出需要的那个
func getShareURL(data *model.Material, id string) []string {
	numId, err := strconv.Atoi(id)
	if err != nil {
		return []string{}
	}
	var URLs = make([]string, 0, 1)
	for _, v := range data.TbkDgMaterialOptionalResponse.ResultList.MapData {
		if v.CouponShareURL != "" && v.ItemID == int64(numId) {
			URLs = append(URLs, v.CouponShareURL)
		}
	}
	return URLs
}

// getAllShareURL
func getAllShareURL(data *model.Material) []string {
	var shareURLs = make([]string, 0, data.TbkDgMaterialOptionalResponse.TotalResults)
	for _, v := range data.TbkDgMaterialOptionalResponse.ResultList.MapData {
		shareURLs = append(shareURLs, v.CouponShareURL)
	}
	return shareURLs
}
