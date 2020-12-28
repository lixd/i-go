package logic

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lixd/vaptcha-sdk-go"
	log "github.com/sirupsen/logrus"
	"i-go/gin/vaptchademo/constant"
)

type ResCode struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// Login login demo for VAPTCHA verify
func Login(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	token, username, password, ip := postParams(request)
	v := vaptcha.NewVaptcha(constant.VID, constant.Key, constant.Scene)
	// invoke vaptcha verify
	ret := v.Verify(token, ip)
	log.Printf("second verify token:%s ip:%s ret:%v", token, ip, ret)
	if ret.Success != 1 {
		bytes := buildResponse(http.StatusBadRequest, "fail")
		_, _ = writer.Write(bytes)
		return
	}
	isLogin := doLogin(username, password)
	bytes := buildResponse(http.StatusOK, isLogin)
	_, _ = writer.Write(bytes)
}

// postParams parse post params
func postParams(request *http.Request) (token, username, password, ip string) {
	token = request.PostFormValue("token")
	username = request.PostFormValue("username")
	password = request.PostFormValue("password")
	ip = request.RemoteAddr
	// fix ipv6
	if strings.Contains(ip, "::1") {
		ip = "127.0.0.1"
	}
	return
}

func buildResponse(code int, msg string) []byte {
	res := ResCode{
		Code: code,
		Data: nil,
		Msg:  msg,
	}
	bytes, _ := json.Marshal(res)
	return bytes
}

// doLogin
func doLogin(userName, password string) string {
	if userName == "admin" && password == "root" {
		return "success"
	}
	return "fail"
}
