package user

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ModelBind struct {
}

func (modelBind *ModelBind) BindHandler(c *gin.Context) {
	var user = User{}
	if err := c.Bind(&user); err == nil {
		log.Printf("c.Bind() user.Name=%v user.Age=%v user.Address=%v", user.Name, user.Age, user.Address)
	} else {
		log.Printf("c.Bind() err=%v ", err)
	}
}
func (modelBind *ModelBind) BindJSONHandler(c *gin.Context) {
	var user = User{}
	if err := c.BindJSON(&user); err == nil {
		log.Printf("c.BindJSON() user.Name=%v user.Age=%v user.Address=%v", user.Name, user.Age, user.Address)
	} else {
		log.Printf("c.BindJSON() err=%v ", err)
	}
}
func (modelBind *ModelBind) BindQueryHandler(c *gin.Context) {
	var user = User{}
	if err := c.BindQuery(&user); err == nil {
		log.Printf("c.BindQuery() user.Name=%v user.Age=%v user.Address=%v", user.Name, user.Age, user.Address)
	} else {
		log.Printf("c.BindQuery() err=%v ", err)
	}
}

func (modelBind *ModelBind) ShouldBindHandler(c *gin.Context) {
	var user = User{}
	if err := c.ShouldBind(&user); err == nil {
		log.Printf("c.ShouldBind() user.Name=%v user.Age=%v user.Address=%v", user.Name, user.Age, user.Address)
	} else {
		log.Printf("c.ShouldBind() err=%v ", err)
	}
}
func (modelBind *ModelBind) ShouldBindJSONHandler(c *gin.Context) {
	var user = User{}
	if err := c.ShouldBindJSON(&user); err == nil {
		log.Printf("c.ShouldBindJSON() user.Name=%v user.Age=%v user.Address=%v", user.Name, user.Age, user.Address)
	} else {
		log.Printf("c.ShouldBindJSON() err=%v ", err)
	}
}
func (modelBind *ModelBind) ShouldBindQueryHandler(c *gin.Context) {
	var user = User{}
	if err := c.ShouldBindQuery(&user); err == nil {
		log.Printf("c.ShouldBindQuery() user.Name=%v user.Age=%v user.Address=%v", user.Name, user.Age, user.Address)
	} else {
		log.Printf("c.ShouldBindQuery() err=%v ", err)
	}
}
