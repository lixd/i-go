package core

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"i-go/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	BaseURL     = "https://pub.alimama.com/openapi/param2/1/gateway.unionpub/shareitem.json?"
	siteId      = "1382050482"
	adzoneId    = "110387000289"
	TbToken     = "617e4e089eeb"
	UA          = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36"
	LoginCookie = "1e969e67026f4f84c7e1ffce8d3e5b67"
)

var Client http.Client

func init() {
	Client = http.Client{Timeout: 60 * time.Second}
}

// ShareItem
func ShareItem(id string) int {
	// 设置参数
	params := url.Values{}
	params.Set("t", strconv.FormatInt(time.Now().UnixNano(), 10))
	params.Set("_tb_token_", "617e4e089eeb")
	params.Set("shareUserType", "1")
	params.Set("unionBizCode", "union_pub")
	params.Set("shareSceneCode", "item_search")
	params.Set("materialId", id)
	params.Set("tkClickSceneCode", "qtz_pub_search")
	params.Set("siteId", siteId)
	params.Set("adzoneId", adzoneId)
	params.Set("bypage", "1")
	//params.Set("extendMap", `{"qtzParam":{"lensId":"OPT@1591596416@0b579758_99a6_1729289fe4e_676b@01"}}`)
	params.Set("materialType", "1")
	params.Set("needQueryQtz", "true")
	Url, err := url.Parse(BaseURL)
	if err != nil {
		panic(err)
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println("URL: ", urlPath)
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("USER-AGENT", UA)
	loginCookie := getLoginCookie()
	req.AddCookie(loginCookie)
	req.Header.Set("Connection", "keep-alive")
	start := time.Now()
	resp, err := Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("ShareItem Time: ", time.Now().Sub(start))
	// resp解析
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "读取resp.Body"}).Error(err)
		return -1
	}
	var data Result
	err = json.Unmarshal(s, &data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "反序列化"}).Error(err)
		return -1
	}
	fmt.Println("Data: ", data)
	return data.ResultCode
}

func ReLogin() {
	// 设置参数
	params := url.Values{}
	params.Set("t", strconv.FormatInt(time.Now().UnixNano(), 10))
	params.Set("_tb_token_", TbToken)
	Url, err := url.Parse("https://pub.alimama.com/openapi/param2/1/gateway.unionpub/record.queryAccessCondition.json")
	if err != nil {
		panic(err)
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println("URL: ", urlPath)
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("USER-AGENT", UA)
	loginCookie := getLoginCookie()
	req.AddCookie(loginCookie)
	req.Header.Set("Connection", "keep-alive")
	start := time.Now()
	resp, err := Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("ShareItem Time: ", time.Now().Sub(start))
	// resp解析
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "读取resp.Body"}).Error(err)
	}
	fmt.Println("Result:", string(s))
	//var data Result
	//err = json.Unmarshal(s, &data)
	//if err != nil {
	//	logrus.WithFields(logrus.Fields{"Caller": utils.Caller(), "Scenes": "反序列化"}).Error(err)
	//}
	//fmt.Println("Data: ", data)
}

func getLoginCookie() *http.Cookie {
	return &http.Cookie{
		Name:     "cookie2",
		Value:    LoginCookie,
		Path:     "/",
		Domain:   ".alimama.com",
		HttpOnly: true,
	}
}

type shareResult struct {
}
