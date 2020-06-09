package sdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type TBK struct {
	AppKey    string
	AppSecret string
	Router    string        // 环境请求地址
	Session   string        // Session 用户登录授权成功后，TOP颁发给应用的授权信息。当此API的标签上注明：“需要授权”，则此参数必传；“不需要授权”，则此参数不需要传；“可选授权”，则此参数为可选
	Timeout   time.Duration // api调用超时
}

func NewTBK(appKey, appSecret, router, session string, timeout time.Duration) *TBK {
	return &TBK{
		AppKey:    appKey,
		AppSecret: appSecret,
		Router:    router,
		Session:   session,
		Timeout:   timeout,
	}
}

type CommonParams struct {
	Method       string `json:"method"`         // API接口名称。
	AppKey       string `json:"app_key"`        // TOP分配给应用的AppKey。
	Session      string `json:"session"`        // Session 用户登录授权成功后，TOP颁发给应用的授权信息。当此API的标签上注明：“需要授权”，则此参数必传；“不需要授权”，则此参数不需要传；“可选授权”，则此参数为可选
	Timestamp    string `json:"timestamp"`      // 时间戳
	Format       string `json:"format"`         // 响应格式。 默认为xml格式，可选值：xml，json。
	V            string `json:"v"`              // API协议版本，可选值：2.0。
	PartnerId    string `json:"partner_id"`     // 合作伙伴身份标识。
	TargetAppKey string `json:"target_app_key"` // 被调用的目标AppKey，仅当被调用的API为第三方ISV提供时有效。
	Simplify     string `json:"simplify"`       // 是否采用精简JSON返回格式，仅当format=json时有效，默认值为：false。
	SignMethod   string `json:"sign_method"`    // 签名的摘要算法，可选值为：hmac，md5，hmac-sha256。
	Sign         string `json:"sign"`           // API输入参数签名结果
}

// Execute 执行 http 请求
func (tbk *TBK) Execute(method string, p Parameter) ([]byte, error) {
	p["method"] = method
	tbk.initCommonParams(p)

	bodyBytes, err := tbk.execute(p)
	if err != nil {
		return bodyBytes, err
	}
	return bodyBytes, nil
}

// execute
func (tbk *TBK) execute(p Parameter) ([]byte, error) {
	err := tbk.checkConfig()
	if err != nil {
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequest("POST", tbk.Router, strings.NewReader(tbk.encodeParams(p)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	httpClient := &http.Client{}
	httpClient.Timeout = tbk.Timeout
	var response *http.Response
	response, err = httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return nil, err
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
