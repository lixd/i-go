package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lixd/vaptcha-sdk-go"
	vmodel "github.com/lixd/vaptcha-sdk-go/model"
	"i-go/gin/vaptchademo/constant"
	"i-go/gin/vaptchademo/gin/model"
)

func Offline(c *gin.Context) {
	m := new(model.Offline)
	if err := c.ShouldBindQuery(m); err != nil {
		c.JSON(http.StatusOK, model.Ret{Code: 400, Msg: "params error"})
		return
	}
	v := vaptcha.NewVaptcha(constant.VID, constant.Key, constant.Scene)

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
