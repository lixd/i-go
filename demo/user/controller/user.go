package controller

import (
	"github.com/gin-gonic/gin"
	"i-go/demo/cmodel"
	"i-go/demo/common/ret"
	"i-go/demo/user/dto"
	"i-go/demo/user/server"
	"net/http"
	"strconv"
)

type IUser interface {
	Insert(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	Find(c *gin.Context)
}

type user struct {
	Server server.IUser
}

func NewUser(ser server.IUser) IUser {
	return &user{Server: ser}
}

// Insert
func (u *user) Insert(c *gin.Context) {
	var m dto.UserReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	m.RegisterIP = c.ClientIP()
	m.LoginIP = c.ClientIP()
	result := u.Server.Insert(&m)
	c.JSON(http.StatusOK, result)
}

// Delete
func (u *user) Delete(c *gin.Context) {
	var m dto.UserReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	result := u.Server.DeleteById(&m)
	c.JSON(http.StatusOK, result)
}

// Update
func (u *user) Update(c *gin.Context) {
	var m dto.UserReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	result := u.Server.UpdateById(&m)
	c.JSON(http.StatusOK, result)
}

// Find
func (u *user) FindById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	res := u.Server.FindById(uint(id))
	c.JSON(http.StatusOK, res)
}

// Find
func (u *user) Find(c *gin.Context) {
	var m cmodel.Page
	if err := c.ShouldBindQuery(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	res := u.Server.Find(&m)
	c.JSON(http.StatusOK, res)
}
