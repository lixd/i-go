package logic

import (
	"github.com/gin-gonic/gin"
	"i-go/gin/vaptchademo/vaptchasdk"
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
	// 直接调用sdk中提供的方法 用户需要做的就是获取到前端传来的参数
	v := vaptchasdk.New(vaptchasdk.Vid, vaptchasdk.SecretKey, vaptchasdk.Scene)
	// 如果是测试离线模式 则vid直接传offline即可
	req.Vid = "offline"
	result := v.Offline(req.Action, req.CallBack, req.Vid, req.Knock, req.UserCode)
	c.String(http.StatusOK, result)
}
