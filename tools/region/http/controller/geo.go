package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"i-go/tools/region/core"
)

// Ip2LatLong ip 转经纬度
func (ip2region) Ip2LatLong(c *gin.Context) {
	ip := c.Query("ip")
	if ip == "" {
		ip = c.ClientIP()
	}
	lagLong, err := core.IP2LatLong(ip)
	if err != nil {
		c.JSON(http.StatusOK, "查询失败"+err.Error())
	}
	c.JSON(http.StatusOK, lagLong)
}
