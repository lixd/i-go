package vaptcha

import (
	"encoding/json"
	"fmt"
	. "i-go/gin/vaptchademo/vaptcha/constant"
	"i-go/gin/vaptchademo/vaptcha/model"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Verify
func (v *vaptcha) Verify(token, ip string) (ret model.Verify) {
	// check mode by first seven characters
	if len(token) < 7 {
		ret.Msg = "invalid token"
		return
	}
	if token[:7] == OfflineMode {
		// offline mode
		return v.offlineVerify(token)
	}
	// online mode
	return v.onlineVerify(token, ip)
}

// offlineVerify offline mode
func (v *vaptcha) offlineVerify(token string) (ret model.Verify) {
	ret.Msg = "验证失败"
	// token=offline(7)+knock(32)+uuid(32)
	knock := token[7:39]
	uid := token[39:]
	// cacheToken:unix(10)+uuid(32)
	cacheToken := v.options.Cache.Get(knock)
	realToken := cacheToken[10:]
	if realToken == uid {
		// 验证成功则移除token
		v.options.Cache.Delete(knock)
		ret = model.Verify{
			Msg:     "success",
			Success: 1,
			Score:   100,
		}
	}
	return
}

// onlineVerify online mode
func (v *vaptcha) onlineVerify(token, ip string) (ret model.Verify) {
	data := url.Values{
		"id":        {v.options.Vid},
		"secretkey": {v.options.SecretKey},
		"scene":     {v.options.Scene},
		"token":     {token},
		"ip":        {ip},
	}
	response, err := http.PostForm(OnlineVerifyURL, data)
	if err != nil {
		ret.Msg = "http error"
		return
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ret.Msg = "internal error"
		return
	}
	defer response.Body.Close()

	err = json.Unmarshal(bytes, &ret)
	if err != nil {
		ret.Msg = "internal error"
		return
	}
	fmt.Println(ret)
	return
}
