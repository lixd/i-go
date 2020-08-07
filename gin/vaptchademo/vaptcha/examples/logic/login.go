package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"i-go/gin/vaptchademo/constant"
	"i-go/gin/vaptchademo/vaptcha"
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
	fmt.Printf("host:%v url:%v method:%v \n", c.Request.Host, c.Request.URL, c.Request.Method)
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
	option := func(options *vaptcha.Options) {
		options.Vid = constant.Vid
		//options.Vid = "offline" // test offline mode
		options.SecretKey = constant.SecretKey
		options.Scene = constant.Scene
	}
	v := vaptcha.NewVaptcha(option)
	// 二次验证
	verify := v.Verify(req.Token, c.ClientIP())
	if verify.Success == 1 {
		// 二次验证成功了 再去执行真正的登录逻辑
		login := doLogin(req.UserName, req.Password)
		c.JSON(http.StatusOK, ResCode{
			Code: http.StatusOK,
			Data: login,
			Msg:  login,
		})
		return
	}
	c.JSON(http.StatusOK, ResCode{
		Code: http.StatusOK,
		Data: verify,
		Msg:  "登录失败",
	})
}

// doLogin 真正的登录逻辑
func doLogin(userName, password string) string {
	if userName == "admin" && password == "root" {
		return "success"
	}
	return ""
}
