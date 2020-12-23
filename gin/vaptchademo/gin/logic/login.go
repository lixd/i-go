package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lixd/vaptcha-sdk-go"
	"i-go/gin/vaptchademo/constant"
	"i-go/gin/vaptchademo/gin/model"
)

func Login(c *gin.Context) {
	m := new(model.Login)
	if err := c.ShouldBind(m); err != nil {
		c.JSON(http.StatusOK, model.Ret{Code: 400, Msg: "params error"})
		return
	}
	v := vaptcha.NewVaptcha(constant.VID, constant.Key, constant.Scene)

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
