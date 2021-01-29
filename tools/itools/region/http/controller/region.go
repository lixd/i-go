package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	region "i-go/tools/itools/region/client"
)

type ip2region struct {
}

var Tools = &ip2region{}

type IP struct {
	IP     string `json:"ip"`
	Region string `json:"region"`
}

func (ip2region) Ip2Region(c *gin.Context) {
	ip := c.Query("ip")
	if ip == "" {
		ip = c.ClientIP()
	}
	reg := region.Ip2Region(ip)
	ret := IP{
		IP:     ip,
		Region: reg,
	}
	c.JSON(http.StatusOK, ret)
}
