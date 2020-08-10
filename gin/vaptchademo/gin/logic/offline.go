package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/lixd/vaptcha-sdk-go"
	vmodel "github.com/lixd/vaptcha-sdk-go/model"
	"i-go/gin/vaptchademo/gin/constant"
	"i-go/gin/vaptchademo/gin/model"
	"net/http"
)

func Offline(c *gin.Context) {
	m := new(model.Offline)
	if err := c.ShouldBindQuery(m); err != nil {
		c.JSON(http.StatusOK, model.Ret{Code: 400, Msg: "params error"})
		return
	}
	option := func(options *vaptcha.Options) {
		options.Vid = constant.Vid
		// options.Vid = "offline" // 如果是测试离线模式 则vid直接传offline即可
		options.SecretKey = constant.SecretKey
		options.Scene = constant.Scene
	}
	v := vaptcha.NewVaptcha(option)
	// invoke sdk offline
	item := vmodel.Offline{
		Action:   m.Action,
		Callback: m.Callback,
		Knock:    m.Knock,
		UserCode: m.UserCode,
	}
	result := v.Offline(item)
	c.Data(http.StatusOK, "application/json", []byte(result))
	return
}
