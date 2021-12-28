// Package ret 统一返回结构
package ret

import (
	"net/http"
)

const (
	MsgSuccess = "success"
	MsgFail    = "fail"
)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Success(data interface{}, msg ...string) *Result {
	var m = MsgSuccess
	if len(msg) > 0 {
		m = msg[0]
	}
	return &Result{
		Code: http.StatusOK,
		Data: data,
		Msg:  m,
	}
}

func Fail(msg ...string) *Result {
	var m = MsgFail
	if len(msg) > 0 {
		m = msg[0]
	}
	return &Result{
		Code: http.StatusBadRequest,
		Data: "",
		Msg:  m,
	}
}
