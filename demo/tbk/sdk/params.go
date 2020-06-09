package sdk

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Parameter map[string]string

// initCommonParams 初始化公共参数
func (tbk *TBK) initCommonParams(p Parameter) {
	p["app_key"] = tbk.AppKey
	if tbk.Session != "" {
		p["session"] = tbk.Session
	}
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	p["timestamp"] = strconv.FormatInt(loc.Unix(), 10)
	p["format"] = "json"
	p["v"] = "2.0"
	p["partner_id"] = "lixd"
	p["Simplify"] = "false"
	p["sign_method"] = "md5"
	// 设置签名
	p["sign"] = tbk.signTopRequest(p)
}

// encodeParams encode为url参数
func (tbk *TBK) encodeParams(p Parameter) string {
	args := url.Values{}
	for k, v := range p {
		args.Set(k, v)
	}
	return args.Encode()
}

// SignTopRequest 获取签名
func (tbk *TBK) signTopRequest(p Parameter) string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	// 第一步：把字典按Key的字母顺序排序
	sort.Strings(keys)
	// 第二步：把所有参数名和参数值串在一起
	query := bytes.NewBufferString(tbk.AppSecret)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(p[k])
	}
	query.WriteString(tbk.AppSecret)
	// 第三步：使用MD5/HMAC加密
	h := md5.New()
	_, _ = io.Copy(h, query)
	/// 第四步：把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// checkConfig 检查配置
func (tbk *TBK) checkConfig() error {
	if tbk.AppKey == "" {
		return AppKeyNil
	}
	if tbk.AppSecret == "" {
		return AppSecretNil
	}
	if tbk.Router == "" {
		return RouterNil
	}
	return nil
}
