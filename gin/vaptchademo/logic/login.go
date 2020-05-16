package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"i-go/gin/vaptchademo/vaptchasdk"
	"net/http"
)

type ResCode struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
type vaptchaDemo struct {
}

var VaptchaDemo = &vaptchaDemo{}

type LoginModel struct {
	UserName string `form:"username"`
	Password string `form:"password"`
	Token    string `form:"token"`
}

// Login 用户登录
// VAPTCHA 演示二次验证配置逻辑
func (*vaptchaDemo) Login(c *gin.Context) {
	var req LoginModel
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, ResCode{
			Code: http.StatusOK,
			Data: "",
			Msg:  "params bind error",
		})
		return
	}
	fmt.Println("params", req)
	if req.UserName == "" || req.Token == "" {
		c.JSON(http.StatusOK, ResCode{
			Code: http.StatusOK,
			Data: "",
			Msg:  "params error",
		})
		c.Data(http.StatusOK, "application/json", []byte("params error"))
		return
	}
	v := vaptchasdk.New(vaptchasdk.Vid, vaptchasdk.SecretKey, vaptchasdk.Scene)
	// 二次验证
	verify := v.Verify(req.Token, c.ClientIP())
	if verify.Success == vaptchasdk.VerifySuccess {
		// 二次验证成功了 再去执行真正的登录逻辑
		// 无效请求虽然也会对服务器增加压力 但是由于不会通过二次验证 所以并不需要真正执行业务逻辑 不会消耗服务器太多资源
		login := doLogin(req.UserName, req.Password)
		c.JSON(http.StatusOK, ResCode{
			Code: http.StatusOK,
			Data: login,
			Msg:  login,
		})
		return
	}
	c.Data(http.StatusOK, "application/json", []byte(""))
}

// doLogin 真正的登录逻辑
func doLogin(userName, password string) string {
	if userName == "admin" && password == "root" {
		return "success"
	}
	return ""
}
