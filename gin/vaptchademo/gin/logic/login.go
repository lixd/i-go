package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/lixd/vaptcha-sdk-go"
	"i-go/gin/vaptchademo/gin/constant"
	"i-go/gin/vaptchademo/gin/model"
	"net/http"
)

func Login(c *gin.Context) {
	m := new(model.Login)
	if err := c.ShouldBind(m); err != nil {
		c.JSON(http.StatusOK, model.Ret{Code: 400, Msg: "params error"})
		return
	}
	option := func(options *vaptcha.Options) {
		options.Vid = constant.Vid
		// options.Vid = "offline" // test offline mode
		options.SecretKey = constant.SecretKey
		options.Scene = constant.Scene
	}
	v := vaptcha.NewVaptcha(option)

	ret := v.Verify(m.Token, c.ClientIP())
	if ret.Success != 1 {
		c.JSON(http.StatusOK, model.Ret{Code: 400, Msg: "verify fail"})
		return
	}
	isLogin := doLogin(m.Username, m.Password)
	c.JSON(http.StatusOK, model.Ret{Code: 200, Msg: isLogin})
	return
}

// doLogin
func doLogin(userName, password string) string {
	if userName == "admin" && password == "root" {
		return "success"
	}
	return "fail"
}
