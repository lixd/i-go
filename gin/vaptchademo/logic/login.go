package logic

import (
	"encoding/json"
	"github.com/lixd/vaptcha-sdk-go"
	"github.com/lixd/vaptcha-sdk-go/examples/constant"
	"net/http"
)

type ResCode struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// Login login demo for show vaptcha verify
func Login(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")

	token, username, password, ip := postParams(request)
	option := func(options *vaptcha.Options) {
		options.Vid = constant.Vid
		// options.Vid = "offline" // test offline mode
		options.SecretKey = constant.SecretKey
		options.Scene = constant.Scene
	}
	v := vaptcha.NewVaptcha(option)
	// invoke vaptcha verify
	verify := v.Verify(token, ip)
	if verify.Success != 1 {
		// verify fail return error
		bytes := buildResponse(400, "fail")
		_, _ = writer.Write(bytes)
		return
	}
	isLogin := doLogin(username, password)
	bytes := buildResponse(200, isLogin)
	_, _ = writer.Write(bytes)
}

// postParams parse post params
func postParams(request *http.Request) (token, username, password, ip string) {
	token = request.PostFormValue("token")
	username = request.PostFormValue("username")
	password = request.PostFormValue("password")
	ip = request.RemoteAddr
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
