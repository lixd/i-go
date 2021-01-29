// Server层返回值
package svc

import (
	"i-go/core/http/ret"
)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// Unauthorized 未授权
func Unauthorized() *Result {
	return &Result{
		Code: ret.Unauthorized,
		Data: "",
		Msg:  ret.UnauthorizedMsg,
	}
}

// BadRequest 请求体语法错误，参数错误
func BadRequest(msg ...string) *Result {
	var m = ret.BadRequest
	if len(msg) > 0 {
		m = msg[0]
	}
	return &Result{
		Code: ret.Fail,
		Data: "",
		Msg:  m,
	}
}

// Forbidden 禁止访问 封禁
func Forbidden() *Result {
	return &Result{
		Code: ret.Forbidden,
		Data: "",
		Msg:  ret.ForbiddenMsg,
	}
}

// Limit 频率限制
func Limit() *Result {
	return &Result{
		Code: ret.Limit,
		Msg:  ret.LimitMsg,
	}
}

func Success(data interface{}, msg ...string) *Result {
	var m = ret.SuccessMsg
	if len(msg) > 0 {
		m = msg[0]
	}
	return &Result{
		Code: ret.Success,
		Data: data,
		Msg:  m,
	}
}

func Fail(msg ...string) *Result {
	var m = ret.FailMsg
	if len(msg) > 0 {
		m = msg[0]
	}
	return &Result{
		Code: ret.Fail,
		Data: "",
		Msg:  m,
	}
}
