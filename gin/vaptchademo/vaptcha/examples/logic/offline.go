package logic

import (
	"github.com/gin-gonic/gin"
	"i-go/gin/vaptchademo/constant"
	"i-go/gin/vaptchademo/vaptcha"
	"i-go/gin/vaptchademo/vaptcha/model"
	"net/http"
)

type Vaptcha struct {
	Action   string `form:"offline_action"`
	CallBack string `form:"callback"`
	Vid      string `form:"vid"`
	Knock    string `form:"knock"`
	UserCode string `form:"v"`
}

// Offline VAPTCHA离线验证接口
func (*vaptchaDemo) Offline(c *gin.Context) {
	var req Vaptcha
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, ResCode{
			Code: http.StatusOK,
			Data: "",
			Msg:  "params bind error",
		})
		return
	}
	if req.Action == "" {
		c.JSON(http.StatusOK, ResCode{
			Code: http.StatusOK,
			Data: "",
			Msg:  "params error",
		})
		return
	}
	option := func(options *vaptcha.Options) {
		options.Vid = constant.Vid
		//options.Vid = "offline"
		options.SecretKey = constant.SecretKey
		options.Scene = constant.Scene
	}
	v := vaptcha.NewVaptcha(option)
	// 如果是测试离线模式 则vid直接传offline即可
	req.Vid = "offline"
	request := model.Offline{
		Action:   req.Action,
		Callback: req.CallBack,
		Knock:    req.Knock,
		UserCode: req.UserCode,
	}
	result := v.Offline(request)
	c.String(http.StatusOK, result)
}
